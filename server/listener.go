package server

import "net"

func NewTCPListener(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}
