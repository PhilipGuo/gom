package protocol

import (
	bufm "go-util/bufmanager"
)

type IProto interface {
	// read a packet from buf
	ReadAPacket(buf bufm.Ibufmanager) Packet

	// write a packet into buf
	WriteAPacket(buf bufm.Ibufmanager, pac Packet) int
}

type Proto struct{}

func (p Proto) ReadAPacket(buf bufm.Ibufmanager) Packet {

}
