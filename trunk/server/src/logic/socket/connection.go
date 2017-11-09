package socket

import (
	"jimny/logs"
	"net"
	"sync"
)

type Connection interface {
	RdMessage() ([]byte, error)
	WrMessage(msg ...[]byte) error
	LocalAddress() net.Addr
	RemoteAddress() net.Addr
	Close()
	Destroy()
}

type TcpConnection struct {
	sync.Mutex
	base      *net.TCPConn
	writeChan chan []byte
	closeFlag bool
}

func (this *TcpConnection) doDestroy() {
	this.base.SetLinger(0)
	this.base.Close()
	if !this.closeFlag {
		close(this.writeChan)
		this.closeFlag = true
	}
}

func (this *TcpConnection) Destroy() {
	this.Lock()
	defer this.Unlock()
	this.doDestroy()
}

func (this *TcpConnection) Close() {
	this.Lock()
	defer this.Unlock()
	if this.closeFlag {
		return
	}
	this.doWrite(nil)
	this.closeFlag = true

}

func (this *TcpConnection) doWrite(b []byte) {

	if len(this.writeChan) == cap(this.writeChan) {
		logs.Debug("close conn: channel full")
		this.doDestroy()
		return
	}

	this.writeChan <- b
}

// b must not be modified by the others goroutines
func (this *TcpConnection) Write(b []byte) {
	this.Lock()
	defer this.Unlock()
	if this.closeFlag || b == nil {
		return
	}

	this.doWrite(b)
}

func (this *TcpConnection) Read(b []byte) (int, error) {
	return this.base.Read(b)
}

func (this *TcpConnection) LocalAddr() net.Addr {
	return this.base.LocalAddr()
}

func (this *TcpConnection) RemoteAddr() net.Addr {
	return this.base.RemoteAddr()
}

//
//func (tcpConn *TcpConnection) ReadMsg() ([]byte, error) {
//	return tcpConn.msgParser.Read(tcpConn)
//}
//
//func (tcpConn *TcpConnection) WriteMsg(args ...[]byte) error {
//	return tcpConn.msgParser.Write(tcpConn, args...)
//}

func (this *TcpConnection) Run() {

}

func NewTcpConnection(base *net.TCPConn, cap int) *TcpConnection {
	return &TcpConnection{sync.Mutex{}, base, make(chan []byte, cap), false}
}
