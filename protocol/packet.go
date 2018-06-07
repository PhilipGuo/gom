package protocol

//
// 信息包
//
type Packet struct {
	// 头帧类型
	HeadType uint16
	// 子帧类型
	SubType uint16
	// 包内容
	context []byte
}

// NewPacket
//
func NewPacket(headtype, subtype uint16, cxt []byte) *Packet {
	return &Packet{
		HeadType: headtype,
		SubType:  subtype,
		context:  cxt,
	}
}

// get context
//
func (p *Packet) GetContext() []byte {
	if nil == p {
		return nil
	}

	return p.context
}
