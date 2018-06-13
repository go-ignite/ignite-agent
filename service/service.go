package service

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"sync"
	"time"

	agent "github.com/go-ignite/ignite-agent"
	"github.com/go-ignite/ignite-agent/config"
	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var (
	imageMap = map[pb.ServiceType_Enum]string{
		pb.ServiceType_SS_LIBEV: utils.GetImageName("ss-libev"),
		pb.ServiceType_SSR:      utils.GetImageName("ssr"),
	}
)

type AgentService struct{}

func verifyToken(tokenString string, isAdmin *bool) bool {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method is invalid")
		}
		return []byte(config.C.App.Secret), nil
	})
	if err != nil {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	if !token.Valid {
		return false
	}
	if isAdmin != nil {
		id, ok := claims["id"].(float64)
		if !ok {
			return false
		}
		if (*isAdmin && id != -1) || (!*isAdmin && id <= 0) {
			return false
		}
	}
	return claims.VerifyExpiresAt(time.Now().Unix(), true)
}

func (s *AgentService) NodeHeartbeat(req *pb.GeneralRequest, stream pb.AgentService_NodeHeartbeatServer) error {
	logrus.Info("node heartbeat start")
	for {
		if err := stream.Send(&pb.HeartbeatStreamServer{}); err != nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	logrus.Info("node heartbeat end")
	return nil
}

func (s *AgentService) ServiceHeartbeat(req *pb.GeneralRequest, stream pb.AgentService_ServiceHeartbeatServer) error {
	return nil
}

func (s *AgentService) Init(ctx context.Context, req *pb.GeneralRequest) (*pb.GeneralResponse, error) {
	logrus.Info("agent init")
	if isAdmin := true; !verifyToken(req.Token, &isAdmin) {
		return nil, errors.New("request token is invalid")
	}
	wg := new(sync.WaitGroup)
	for _, image := range imageMap {
		reader, err := agent.PullImage(image)
		if err != nil {
			return nil, err
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			io.Copy(ioutil.Discard, reader)
		}()
	}
	wg.Wait()
	return &pb.GeneralResponse{}, nil
}

func (s *AgentService) GetAvailablePort(ctx context.Context, req *pb.GetAvailablePortRequest) (*pb.GetAvailablePortResponse, error) {
	logrus.WithFields(logrus.Fields{
		"usedPorts": req.UsedPorts,
		"portFrom":  req.PortFrom,
		"portTo":    req.PortTo,
	}).Info("get available port")
	if isAdmin := false; !verifyToken(req.Token, &isAdmin) {
		return nil, errors.New("request token is invalid")
	}
	port, err := utils.GetAvailablePort(req.PortFrom, req.PortTo, req.UsedPorts)
	if err != nil {
		return nil, err
	}
	logrus.WithField("port", port).Info("get available port")
	return &pb.GetAvailablePortResponse{Port: port}, nil
}

func (s *AgentService) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	logrus.WithFields(logrus.Fields{
		"name":   req.Name,
		"type":   req.Type,
		"image":  imageMap[req.Type],
		"method": req.Method,
		"port":   req.Port,
	}).Info("create service")

	if isAdmin := false; !verifyToken(req.Token, &isAdmin) {
		return nil, errors.New("request token is invalid")
	}
	serviceID, err := agent.CreateContainer(imageMap[req.Type], req.Name, req.Method, req.Password, req.Port)
	if err != nil {
		return nil, err
	}
	err = agent.StartContainer(serviceID)
	if err != nil {
		agent.RemoveContainer(serviceID)
		return nil, err
	}
	return &pb.CreateServiceResponse{ServiceId: serviceID}, nil
}

func (s *AgentService) StopService(ctx context.Context, req *pb.StopServiceRequest) (*pb.GeneralResponse, error) {
	logrus.WithField("serviceID", req.ServiceId).Info("stop service")

	if isAdmin := false; !verifyToken(req.Token, &isAdmin) {
		return nil, errors.New("request token is invalid")
	}
	err := agent.StopContainer(req.ServiceId)
	if err != nil {
		return nil, err
	}
	return &pb.GeneralResponse{}, nil
}

func (s *AgentService) RemoveService(ctx context.Context, req *pb.RemoveServiceRequest) (*pb.GeneralResponse, error) {
	logrus.WithField("serviceID", req.ServiceId).Info("remove service")

	if !verifyToken(req.Token, nil) {
		return nil, errors.New("request token is invalid")
	}
	err := agent.RemoveContainer(req.ServiceId)
	if err != nil {
		return nil, err
	}
	return &pb.GeneralResponse{}, nil
}
