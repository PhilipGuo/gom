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
	lins *net.Listener
	// 用户连接
	sessions *SessionBucket
	// 退出标志
	exit chan int

	m gom.Igom
}

// NewTCPServer...
//
func (server *TCPServer) NewTCPServer(ip string, port int) *TCPServer {
	return &TCPServer{
		ip:       ip,
		port:     port,
		sessions: NewSessionBucket(),
		//m:        gom.Newgom(),
		exit: make(chan int, 1),
	}
}

func (server *TCPServer) Start() (bool, error) {
	if nil == server {
		return false, errors.New("server is nil")
	}
	l, err := net.Listen("tcp", server.ip+":"+string(server.port))
	if nil != err {
		return false, err
	}
	server.m = gom.Newgom()
	server.lins = &l
	for {
		c, err := (*server.lins).Accept()
		if nil != err {
			// todo
		}

		se := NewSession(&c, server.m)
		server.sessions.Put(c.RemoteAddr().String(), se)
		// todo
		// se.start()

	}

	return true, nil
}
