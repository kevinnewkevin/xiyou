package socket

import (
	"bytes"
	"net"
	"sync"
	"logic/log"
	"runtime"
	"strings"
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
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		file = "???"
		line = 0
	}else {

			i := strings.LastIndex(file, "/")
			if i == -1 {
				i = strings.LastIndex(file, "\\")
			}
			if i != -1{
				file = file[i+1:]
			}

	}
	log.Info("CALL %s:%d MethodEnd() %d ",file,line, c)
	this.OutgoingBuffer.Reset()
	this.TotalOutgoing += c
	return nil
}

func (this *Peer) HandleSocket() error {
	{
		bs := make([]byte, 2048)

		c, e := this.Connection.Read(bs)
		if e != nil {
			log.Error("Incoming error %s",e.Error())
			return e
		}
		this.IncomingBuffer.Write(bs[:c])
		this.TotalIncoming += c
	}

	{
		c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
		if e != nil {
			log.Error("Outgoing error %s",e.Error())

			return e
		}
		this.TotalOutgoing += c
	}
	return nil
}


func NewPeer(conn *net.TCPConn) *Peer {
	conn.SetNoDelay(true)
	return &Peer{IncomingBuffer: bytes.NewBuffer(nil), OutgoingBuffer: bytes.NewBuffer(nil), Connection: conn}
}
