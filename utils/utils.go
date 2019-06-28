package utils

import (
	"errors"
	"fmt"
	"net"
)

func GetAvailablePort(portFrom, portTo int32, usedPortMap map[int32]bool) (int32, error) {
	for port := portFrom; port <= portTo; port++ {
		if usedPortMap[port] {
			continue
		}
		conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return port, nil
		}
		_ = conn.Close()
	}

	return 0, errors.New("no available port")
}

func CopyMap(m map[string]string) map[string]string {
	cp := make(map[string]string)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
