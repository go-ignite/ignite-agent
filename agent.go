package agent

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const (
	agentLableKey   = "org.label-schema.url"
	agentLableValue = "https://github.com/go-ignite"
)

func GetDockerClient() (*client.Client, error) {
	return client.NewEnvClient()
}

func PullImage(image string) (io.ReadCloser, error) {
	cli, err := GetDockerClient()
	if err != nil {
		return nil, fmt.Errorf("GetDockerClient error: %v", err)
	}

	return cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
}

func ListImages() ([]types.ImageSummary, error) {
	cli, err := GetDockerClient()
	if err != nil {
		return nil, fmt.Errorf("GetDockerClient error: %v", err)
	}
	filters := filters.NewArgs()
	filters.Add("label", fmt.Sprintf("%s=%s", agentLableKey, agentLableValue))
	filters.Add("dangling", "false")
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All:     false,
		Filters: filters,
	})
	if err != nil {
		return nil, fmt.Errorf("cli.ImageList error: %v", err)
	}
	return images, nil
}
