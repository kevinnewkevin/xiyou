package socket

import (
	"net"
	"bytes"
)

type Peer struct {
	TotalIncoming, TotalOutgoing int
	IncomingBuffer, OutgoingBuffer *bytes.Buffer
	Connection                   net.Conn
}

func(this *Peer) MethodBegin() *bytes.Buffer{
	return this.OutgoingBuffer
}

func(this *Peer) MethodEnd() error {
	c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
	if e != nil{
		return e
	}
	this.TotalOutgoing += c
	return nil
}

func(this *Peer) HandleSocket() error {
	{
		bs := make([]byte, 2048)

		c, e := this.Connection.Read(bs)
		if e != nil {
			return  e
		}
		this.IncomingBuffer.Write(bs[:c])
		this.TotalIncoming += c

	}

	{
		c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
		if e != nil{
			return e
		}
		this.TotalOutgoing += c
	}
	return  nil
}

func NewPeer(conn net.Conn)* Peer{
	return &Peer{IncomingBuffer:bytes.NewBuffer(nil),OutgoingBuffer:bytes.NewBuffer(nil),Connection:conn}
}