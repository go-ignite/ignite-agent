package agent

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	agentLabelKey   = "org.label-schema.url"
	agentLabelValue = "https://github.com/go-ignite"
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
	filters.Add("label", fmt.Sprintf("%s=%s", agentLabelKey, agentLabelValue))
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

func CreateContainer(image, name, method, password string, port int32) (string, error) {
	cli, err := GetDockerClient()
	if err != nil {
		return "", fmt.Errorf("GetDockerClient error: %v", err)
	}
	config := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			"3389/tcp": struct{}{},
		},
		Cmd: []string{"-k", password, "-m", method},
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
	result, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, name)

	return result.ID, err
}

func StartContainer(id string) error {
	cli, err := GetDockerClient()
	if err != nil {
		return fmt.Errorf("GetDockerClient error: %v", err)
	}
	return cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
}

func RemoveContainer(id string) error {
	cli, err := GetDockerClient()
	if err != nil {
		return fmt.Errorf("GetDockerClient error: %v", err)
	}
	return cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{Force: true})
}
