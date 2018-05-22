package service

import (
	"fmt"

	agent "github.com/go-ignite/ignite-agent"
)

type Service struct {
	Name string
	agent.Dockerfile

	ImageExist   bool
	ImageVersion string
}

func GetServices() []*Service {
	return []*Service{
		NewService("ss-libev", SSLibev, "3.1.3"),
		NewService("ssr", SSR, "3.1.2"),
	}
}

func NewService(name, content, version string) *Service {
	return &Service{
		Name:       name,
		Dockerfile: agent.NewDockerfile(getImageName(name), content, version),
	}
}

func getImageName(name string) string {
	return fmt.Sprintf("goignite/agent-%s:latest", name)
}
