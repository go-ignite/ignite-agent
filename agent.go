package agent

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const (
	agentLableKey   = "org.label-schema.url"
	agentLableValue = "https://github.com/go-ignite/ignite-agent"
)

func GetDockerClient() (*client.Client, error) {
	return client.NewEnvClient()
}

func BuildImage(d Dockerfile, logWriter io.Writer) error {
	cli, err := GetDockerClient()
	if err != nil {
		return fmt.Errorf("GetDockerClient error: %v", err)
	}

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	if err := tw.WriteHeader(&tar.Header{
		Name: "Dockerfile",
		Size: int64(len(d.Content)),
	}); err != nil {
		return fmt.Errorf("tw.WriteHeader error: %v", err)
	}
	if _, err := tw.Write(d.Content); err != nil {
		return fmt.Errorf("tw.Write error: %v", err)
	}

	buildContext := bytes.NewReader(buf.Bytes())
	buildOptions := types.ImageBuildOptions{
		Tags: []string{d.Image},
		Labels: map[string]string{
			agentLableKey: agentLableValue,
			"version":     d.Version,
		},
	}
	resp, err := cli.ImageBuild(context.Background(), buildContext, buildOptions)
	if err != nil {
		return fmt.Errorf("cli.ImageBuild error: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(logWriter, resp.Body)
	if err != nil {
		return fmt.Errorf("io.Copy error: %v", err)
	}
	return nil
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
