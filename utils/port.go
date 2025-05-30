package utils

import (
	"fmt"
	"net"
)

// FindAvailablePort mencari port yang tersedia dari daftar prioritas
func FindAvailablePort(ports []int) (int, error) {
	for i := 0; i < len(ports); i++ {
		port := ports[i]
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return port, nil
		}
	}
	return 0, fmt.Errorf("does not find available port on the list")
}