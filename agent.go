package agent

import (
	"github.com/go-ignite/ignite-agent/config"
	"github.com/go-ignite/ignite-agent/server"
	"github.com/go-ignite/ignite-agent/service"
)

//go:generate protoc --go_out=plugins=grpc,paths=source_relative:. protos/agent.proto

type Agent struct {
	s *server.Server
}

func Init() (*Agent, error) {
	c, err := config.Init()
	if err != nil {
		return nil, err
	}

	s, err := service.Init()
	if err != nil {
		return nil, err
	}

	return &Agent{
		s: server.New(c, s),
	}, nil
}

func (a *Agent) Start() error {
	return a.s.Start()
}
