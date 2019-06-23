package main

import (
	"github.com/go-ignite/ignite-agent/utils"
	"log"
	"net"

	"github.com/go-ignite/ignite-agent/config"
	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func checkEnv() {
	client, err := utils.GetDockerClient()
	if err != nil {
		logrus.Fatalf("agent.GetDockerClient error: %v", err)
	}
	defer client.Close()
}

func main() {
	config.Init()
	checkEnv()

	lis, err := net.Listen("tcp", config.C.App.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logrus.WithField("addr", lis.Addr()).Info("grpc server start")
	grpcServer := grpc.NewServer()
	pb.RegisterAgentServiceServer(grpcServer, &service.AgentService{})
	grpcServer.Serve(lis)
}
