package agent

type Dockerfile struct {
	Image   string
	Content []byte
	Version string
}

func NewDockerfile(image, content, version string) Dockerfile {
	return Dockerfile{
		Image:   image,
		Content: []byte(content),
		Version: version,
	}
}
