package server

import (
	"errors"
	gom "gom/gom"
	"net"
)

type TCPServer struct {
	//server ip
	ip string
	// 监听端口
	port int

	// 监听socket
	con *net.Conn

	// 退出标志
	exit chan int

	m gom.Igom
}

// NewTCPServer...
//
func (server *TCPServer) NewTCPServer(ip string, port int) *TCPServer {
	return &TCPServer{
		ip:   ip,
		port: port,
		exit: make(chan int, 1),
	}
}

func (server *TCPServer) Start() (bool, error) {
	if nil == server {
		return false, errors.New("server is nil")
	}

	return true, nil
}
