package server

import (
	bufm "Go-util/bufmanager"
	"errors"
	gom "gom/gom"
	proto "gom/protocol"
	"net"
)

type Session struct {
	// 连接
	con *net.Conn

	// 接收缓存
	recvbuf bufm.Ibufmanager
	// 接收包
	recvpacks chan *proto.Packet
	// 发送包
	sendpacks chan *proto.Packet

	// 数据存储
	m gom.Igom

	exitchan chan byte
}

// NewSession
//
func NewSession(c *net.Conn, gm gom.Igom) *Session {
	return &Session{
		con:       c,
		recvbuf:   bufm.NewBufManager(),
		recvpacks: make(chan *proto.Packet),
		sendpacks: make(chan *proto.Packet),
		m:         gm,
	}
}

func (s *Session) Start() (bool, error) {
	if nil == s {
		return false, errors.New("s is nil")
	}


	return true, nil
}

func (s *Session) ReadPack() {
	if nil == s {
		return
	}

	for {
		select {
			case <- s.
		}
	}
}
