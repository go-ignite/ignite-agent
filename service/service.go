package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/utils"
)

type Service struct {
	cli *client.Client
}

func Init() (*Service, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Service{
		cli: cli,
	}, nil
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

	for {
		var services []*pb.ServiceInfo
		if err := func() error {
			containers, err := s.cli.ContainerList(context.Background(), types.ContainerListOptions{
				All: true,
			})

			if err != nil {
				return err
			}

			for _, c := range containers {
				if err := func() error {
					svc := &pb.ServiceInfo{
						ServiceId:   c.Names[0],
						ContainerId: c.ID,
						Port:        int32(c.Ports[0].PublicPort),
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

func (s *Service) Init(ctx context.Context, req *pb.GeneralRequest) (*pb.GeneralResponse, error) {
	wg := new(sync.WaitGroup)
	images := []string{pb.ServiceType_SS_LIBEV.ImageName(), pb.ServiceType_SSR.ImageName()}

	for _, image := range images {
		wg.Add(1)
		reader, err := s.cli.ImagePull(ctx, image, types.ImagePullOptions{})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		go func() {
			defer wg.Done()
			_, _ = io.Copy(os.Stdout, reader)
		}()
	}

	wg.Wait()
	return &pb.GeneralResponse{}, nil
}

func (s *Service) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	// first find an available port
	containers, err := s.cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	usedPortMap := map[int32]bool{}
	for _, c := range containers {
		usedPortMap[int32(c.Ports[0].PublicPort)] = true
	}

	port, err := utils.GetAvailablePort(req.PortFrom, req.PortTo, usedPortMap)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// create container
	config := &container.Config{
		Image: req.Type.ImageName(),
		ExposedPorts: nat.PortSet{
			"3389/tcp": struct{}{},
		},
		Cmd: []string{"-k", req.Password, "-m", req.EncryptionMethod.ValidMethod()},
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
	result, err := s.cli.ContainerCreate(context.Background(), config, hostConfig, nil, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.cli.ContainerStart(context.Background(), result.ID, types.ContainerStartOptions{}); err != nil {
		// remove the container
		_ = s.cli.ContainerRemove(context.Background(), result.ID, types.ContainerRemoveOptions{Force: true})
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateServiceResponse{
		ServiceId:   req.Name,
		ContainerId: result.ID,
		Port:        port,
	}, nil
}

func (s *Service) StopService(ctx context.Context, req *pb.StopServiceRequest) (*pb.GeneralResponse, error) {
	if err := s.cli.ContainerStop(context.Background(), req.ServiceId, nil); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GeneralResponse{}, nil
}

func (s *Service) RemoveService(ctx context.Context, req *pb.RemoveServiceRequest) (*pb.GeneralResponse, error) {
	if err := s.cli.ContainerRemove(context.Background(), req.ServiceId, types.ContainerRemoveOptions{Force: true}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GeneralResponse{}, nil
}
