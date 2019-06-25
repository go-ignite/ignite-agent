package main

import (
	"context"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/go-ignite/ignite-agent/config"
	pb "github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/service"
	"github.com/go-ignite/ignite-agent/utils"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
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
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor), grpc.StreamInterceptor(streamServerInterceptor))
	pb.RegisterAgentServiceServer(grpcServer, &service.AgentService{})
	grpcServer.Serve(lis)
}

func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	if !ensureValidToken(md) {
		return nil, errInvalidToken
	}

	return handler(ctx, req)
}

func streamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errMissingMetadata
	}

	if !ensureValidToken(md) {
		return errInvalidToken
	}

	return handler(srv, ss)
}

func ensureValidToken(md metadata.MD) bool {
	authorization := md["authorization"]
	if len(authorization) < 1 {
		return false
	}

	return authorization[0] == config.C.App.Token
}
