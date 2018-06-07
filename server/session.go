package server

import (
	bufm "Go-util/bufmanager"
	gom "gom/gom"
	proto "gom/protocol"
	"net"
)

type Session struct {
	// 连接
	con *net.Conn

	// 接收缓存
	recvbuf *bufm.Bufmanager
	// 接收包
	recvpacks chan *proto.Packet
	// 发送包
	sendpacks chan *proto.Packet

	// 数据存储
	m gom.Igom
}