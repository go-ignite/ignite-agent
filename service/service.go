package service

import (
	"context"
	"io"
	"io/ioutil"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/utils"
)

var (
	imageMap = map[pb.ServiceType_Enum]string{
		pb.ServiceType_SS_LIBEV: utils.GetImageName("ss-libev"),
		pb.ServiceType_SSR:      utils.GetImageName("ssr"),
	}
)

type AgentService struct{}

func (s *AgentService) Heartbeat(req *pb.HeartbeatRequest, stream pb.AgentService_HeartbeatServer) error {
	logrus.Info("node heartbeat starts")

	interval, err := ptypes.Duration(req.Interval)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "interval is invalid")
	}

	for {
		if err := stream.Send(&pb.HeartbeatStreamServer{}); err != nil {
			break
		}
		time.Sleep(interval)
	}
	logrus.Info("node heartbeat end")
	return nil
}

func (s *AgentService) Sync(req *pb.SyncRequest, stream pb.AgentService_SyncServer) error {
	logrus.Info("sync service starts")

	// init tht service list
	services := []*pb.ServiceInfo{}

	for {
		// sync service data
		interval, err := ptypes.Duration(req.SyncInterval)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "interval is invalid")
		}

		time.Sleep(interval)

		// load detail info for every services
		containers, err := utils.ListContainers()
		if err != nil {
			logrus.Error("failed to list container:", err)
			continue
		}

		for _, c := range containers {
			svc := &pb.ServiceInfo{
				ServiceId:   c.Names[0],
				UserId:      c.Names[0],
				ContainerId: c.ID,
				Port:        int32(c.Ports[0].PublicPort),
			}

			// get net data of container
			statResult, err := utils.GetContainerStatsOutNet(svc.ContainerId)
			if err != nil {
				logrus.Error(err)
			}

			svc.StatsResult = int64(statResult)
			services = append(services, svc)
		}

		if err := stream.Send(&pb.SyncStreamServer{
			Services: services,
		}); err != nil {
			logrus.Error("sync service stream is unavailable")
			break
		}
	}

	logrus.Info("sync service end")
	return nil
}

func (s *AgentService) Init(ctx context.Context, req *pb.GeneralRequest) (*pb.GeneralResponse, error) {
	logrus.Info("agent init")

	wg := new(sync.WaitGroup)
	for _, image := range imageMap {
		reader, err := utils.PullImage(image)
		if err != nil {
			return nil, err
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = io.Copy(ioutil.Discard, reader)
		}()
	}
	wg.Wait()
	return &pb.GeneralResponse{}, nil
}

func (s *AgentService) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	logrus.WithFields(logrus.Fields{
		"name":   req.Name,
		"type":   req.Type,
		"image":  imageMap[req.Type],
		"method": req.EncryptionMethod.String(),
	}).Info("create service")

	// TODO: refactor this part
	serviceID, err := utils.CreateContainer(imageMap[req.Type], req.Name, req.EncryptionMethod.ValidMethod(), req.Password, 0)
	if err != nil {
		return nil, err
	}
	err = utils.StartContainer(serviceID)
	if err != nil {
		_ = utils.RemoveContainer(serviceID)
		return nil, err
	}
	return &pb.CreateServiceResponse{ServiceId: serviceID}, nil
}

func (s *AgentService) StopService(ctx context.Context, req *pb.StopServiceRequest) (*pb.GeneralResponse, error) {
	logrus.WithField("serviceID", req.ServiceId).Info("stop service")

	err := utils.StopContainer(req.ServiceId)
	if err != nil {
		return nil, err
	}
	return &pb.GeneralResponse{}, nil
}

func (s *AgentService) RemoveService(ctx context.Context, req *pb.RemoveServiceRequest) (*pb.GeneralResponse, error) {
	logrus.WithField("serviceID", req.ServiceId).Info("remove service")

	err := utils.RemoveContainer(req.ServiceId)
	if err != nil {
		return nil, err
	}
	return &pb.GeneralResponse{}, nil
}
