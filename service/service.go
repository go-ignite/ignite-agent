package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/utils"
)

var labels = map[string]string{
	"org.label-schema.url": "https://github.com/go-ignite",
	"maintainer":           "go-ignite",
}

type Service struct {
	cli *client.Client
}

func newLabels(nodeID string) map[string]string {
	nl := map[string]string{}
	for k, v := range labels {
		nl[k] = v
	}
	nl["nodeID"] = nodeID
	return nl
}

func Init() (*Service, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cli: cli,
	}

	// pull the required images
	if err := svc.pullImages(); err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *Service) pullImages() error {
	wg := new(sync.WaitGroup)
	images := []string{pb.ServiceType_SS_LIBEV.ImageName(), pb.ServiceType_SSR.ImageName()}

	for _, image := range images {
		wg.Add(1)
		reader, err := s.cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
		if err != nil {
			return err
		}

		go func() {
			defer wg.Done()
			_, _ = io.Copy(os.Stdout, reader)
		}()
	}

	wg.Wait()
	return nil
}

func (s *Service) Heartbeat(req *pb.HeartbeatRequest, stream pb.AgentService_HeartbeatServer) error {
	interval, err := ptypes.Duration(req.Interval)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "interval is invalid")
	}

	for {
		if err := stream.Send(new(pb.HeartbeatStreamServer)); err != nil {
			break
		}

		select {
		case <-stream.Context().Done():
			break
		case <-time.After(interval):
		}
	}

	return nil
}

func (s *Service) Sync(req *pb.SyncRequest, stream pb.AgentService_SyncServer) error {
	interval, err := ptypes.Duration(req.SyncInterval)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "interval is invalid")
	}
	nl := newLabels(req.NodeId)

	for {
		var services []*pb.ServiceInfo
		if err := func() error {
			containers, err := s.containerList(nl)
			if err != nil {
				return err
			}

			for _, c := range containers {
				if err := func() error {
					state := pb.ServiceStatus_RUNNING
					if c.State != "running" {
						state = pb.ServiceStatus_STOPPED
					}
					svc := &pb.ServiceInfo{
						ContainerName: strings.TrimPrefix(c.Names[0], "/"),
						ContainerId:   c.ID,
						Status:        state,
					}

					stat, err := s.cli.ContainerStats(context.Background(), c.ID, false)
					if err != nil {
						return err
					}

					bytes, err := ioutil.ReadAll(stat.Body)
					if err != nil {
						return err
					}

					sj := types.StatsJSON{}
					if err := json.Unmarshal(bytes, &sj); err != nil {
						return err
					}

					svc.StatsResult = int64(sj.Networks["eth0"].TxBytes)
					services = append(services, svc)
					return nil
				}(); err != nil {
					logrus.WithError(err).WithField("containerID", c.ID).Error("get container info error")
				}
			}

			return nil
		}(); err != nil {
			logrus.WithError(err).Error("get containers error")
			continue
		}

		if len(services) > 0 {
			resp := &pb.SyncStreamServer{
				Services: services,
			}
			if err := stream.Send(resp); err != nil {
				break
			}
		}

		select {
		case <-stream.Context().Done():
			break
		case <-time.After(interval):
		}
	}

	return nil
}

func (s *Service) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	// first find an available port
	nl := newLabels(req.NodeId)
	containers, err := s.containerList(nl)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	usedPortMap := map[int32]bool{}
	for _, c := range containers {
		port, err := strconv.Atoi(c.Labels["port"])
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		usedPortMap[int32(port)] = true
	}

	port, err := utils.GetAvailablePort(req.PortFrom, req.PortTo, usedPortMap)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	nl["port"] = fmt.Sprintf("%d", port)

	// create container
	config := &container.Config{
		Image: req.Type.ImageName(),
		ExposedPorts: nat.PortSet{
			"3389/tcp": struct{}{},
		},
		Cmd:    []string{"-k", req.Password, "-m", req.EncryptionMethod.ValidMethod()},
		Labels: nl,
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3389/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("%d", port),
				},
			},
		},
		RestartPolicy: container.RestartPolicy{Name: "always"},
	}
	result, err := s.cli.ContainerCreate(context.Background(), config, hostConfig, nil, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.cli.ContainerStart(context.Background(), result.ID, types.ContainerStartOptions{}); err != nil {
		// remove the container
		_ = s.cli.ContainerRemove(context.Background(), result.ID, types.ContainerRemoveOptions{Force: true})
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateServiceResponse{
		ContainerName: req.UserId,
		ContainerId:   result.ID,
		Port:          port,
	}, nil
}

func (s *Service) StopService(ctx context.Context, req *pb.StopServiceRequest) (*pb.EmptyResponse, error) {
	if err := s.cli.ContainerStop(context.Background(), req.ContainerId, nil); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.EmptyResponse{}, nil
}

func (s *Service) StartService(ctx context.Context, req *pb.StartServiceRequest) (*pb.EmptyResponse, error) {
	if err := s.cli.ContainerStart(context.Background(), req.ContainerId, types.ContainerStartOptions{}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.EmptyResponse{}, nil
}

func (s *Service) RemoveService(ctx context.Context, req *pb.RemoveServiceRequest) (*pb.EmptyResponse, error) {
	if err := s.cli.ContainerRemove(context.Background(), req.ContainerId, types.ContainerRemoveOptions{Force: true}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.EmptyResponse{}, nil
}

func (s *Service) containerList(nl map[string]string) ([]types.Container, error) {
	args := filters.NewArgs()
	for k, v := range nl {
		args.Add("label", fmt.Sprintf("%s=%s", k, v))
	}

	return s.cli.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: args,
		All:     true,
	})
}
