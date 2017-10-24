package socket

import (
	"bytes"
	"net"
	"sync"
	"logic/log"
)

type Peer struct {
	sync.Mutex
	TotalIncoming, TotalOutgoing   int
	IncomingBuffer, OutgoingBuffer *bytes.Buffer
	Connection                     *net.TCPConn
}

func (this *Peer) MethodBegin() *bytes.Buffer {
	this.Lock()
	return this.OutgoingBuffer
}

func (this *Peer) MethodEnd() error {
	defer this.Unlock()
	c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
	if e != nil {
		return e
	}

	log.Info("func (this *Peer) MethodEnd() %d",c)
	this.OutgoingBuffer.Reset()
	this.TotalOutgoing += c
	return nil
}

func (this *Peer) HandleSocket() error {
	{
		bs := make([]byte, 2048)

		c, e := this.Connection.Read(bs)
		if e != nil {
			return e
		}
		this.IncomingBuffer.Write(bs[:c])
		this.TotalIncoming += c
	}

	{
		c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
		if e != nil {
			return e
		}
		this.TotalOutgoing += c
	}
	return nil
}


func NewPeer(conn *net.TCPConn) *Peer {
	return &Peer{IncomingBuffer: bytes.NewBuffer(nil), OutgoingBuffer: bytes.NewBuffer(nil), Connection: conn}
}
