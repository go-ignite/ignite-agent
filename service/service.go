package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	agent "github.com/go-ignite/ignite-agent"
	"github.com/go-ignite/ignite-agent/config"
	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/sirupsen/logrus"
)

type AgentService struct{}

func verifyToken(tokenString, secret string) bool {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(secret), nil
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
	logrus.Info("init start")
	if !verifyToken(req.Token, config.C.Secret.Admin) {
		return nil, errors.New("request token is invalid")
	}
	services := agent.GetServices()
	wg := new(sync.WaitGroup)
	for _, service := range services {
		reader, err := agent.PullImage(service.Image)
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
	logrus.Info("init end")
	return &pb.GeneralResponse{}, nil
}
