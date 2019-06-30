package server

import (
	"context"
	"net"
	"runtime/debug"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/go-ignite/ignite-agent/config"
	"github.com/go-ignite/ignite-agent/protos"
	"github.com/go-ignite/ignite-agent/service"
)

type Server struct {
	config  *config.Config
	service *service.Service
	server  *grpc.Server
}

func New(config *config.Config, service *service.Service) *Server {
	authFunc := func(ctx context.Context) (context.Context, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
		}

		authorization := md["authorization"]
		if len(authorization) < 1 || authorization[0] != config.Token {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		return ctx, nil
	}

	logEntry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	recoverOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logrus.Errorf("recover from panic: %v\n%s\n", p, debug.Stack())
			return status.Errorf(codes.Internal, "%v", p)
		}),
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(authFunc),
			grpc_logrus.StreamServerInterceptor(logEntry, opts...),
			grpc_recovery.StreamServerInterceptor(recoverOpts...),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(authFunc),
			grpc_logrus.UnaryServerInterceptor(logEntry, opts...),
			grpc_recovery.UnaryServerInterceptor(recoverOpts...),
		)),
	)
	protos.RegisterAgentServiceServer(server, service)

	// register health check service
	hs := health.NewServer()
	hs.SetServingStatus(protos.ServiceName, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(server, hs)

	return &Server{
		config:  config,
		service: service,
		server:  server,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return err
	}

	logrus.WithField("address", s.config.Address).Info("start agent gRPC server")
	return s.server.Serve(lis)
}
