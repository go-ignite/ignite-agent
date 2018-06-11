package utils

import (
	"errors"
	"fmt"
	"net"
)

func GetAvailablePort(portFrom, portTo int32, usedPorts []int32) (int32, error) {
	portMap := map[int32]bool{}
	for _, port := range usedPorts {
		portMap[port] = true
	}

	for port := portFrom; port <= portTo; port++ {
		if portMap[port] {
			continue
		}
		conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return port, nil
		}
		conn.Close()
	}
	return 0, errors.New("no available port")
}

func GetImageName(name string) string {
	return fmt.Sprintf("goignite/%s:latest", name)
}
