package agent

import "fmt"

type Service struct {
	Name  string
	Image string

	ImageExist   bool
	ImageVersion string
}

func GetServices() []*Service {
	return []*Service{
		NewService("ss-libev"),
		NewService("ssr"),
	}
}

func NewService(name string) *Service {
	return &Service{
		Name:  name,
		Image: getImage(name),
	}
}

func getImage(name string) string {
	return fmt.Sprintf("goignite/%s:latest", name)
}
