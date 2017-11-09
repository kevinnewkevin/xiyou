package network

import (
	"context"
	"jimny/logs"
	"net"

	"time"
)

type (
	Conn struct {
		conn        net.Conn
		ctx         context.Context
		recvChannel chan []byte
		sendChannel chan []byte
		recvBytes   int
		sendBytes   int
		openTime    time.Time
		makeTime    time.Time
		cancelFunc  func()
	}
)

func (c *Conn) Handle() net.Conn {
	return c.conn
}

func (c *Conn) Cancel() {
	c.cancelFunc()
	c.conn.Close()
}

func (c *Conn) Status() map[string]interface{} {
	stat := map[string]interface{}{}
	stat["recvBytes"] = c.recvBytes
	stat["sendBytes"] = c.sendBytes
	stat["recvChannelSize"] = len(c.recvChannel)
	stat["sendChannelSize"] = len(c.sendChannel)
	stat["openTime"] = c.openTime
	stat["makeTime"] = c.makeTime
	return stat
}

func (c *Conn) TotalOutgoing() int {
	return c.sendBytes
}

func (c *Conn) TotalIncoming() int {
	return c.recvBytes
}

func (c *Conn) RemoteAddress() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) LocalAddress() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) Open() (chan<- []byte, <-chan []byte) {
	c.openTime = time.Now()
	return c.sendChannel, c.recvChannel
}

func newConn(ctx context.Context, conn net.Conn) *Conn {

	c := &Conn{}
	c.conn = conn
	c.ctx, c.cancelFunc = context.WithCancel(ctx)

	c.recvChannel = make(chan []byte, defaultChannelBufferMax)
	c.sendChannel = make(chan []byte, defaultChannelBufferMax)
	c.makeTime = time.Now()
	//read routine
	rdContext, _ := context.WithCancel(c.ctx)
	go func(ctx context.Context) {

		defer func() {
			if r := recover(); r != nil {
				logs.Errorf("Conn read routine panic %s", r)
			}

		}()

		b := make([]byte, defaultByteBufferMax)
		for {
			select {
			case <-ctx.Done():
				close(c.recvChannel)

				return
			default:

				i, e := c.conn.Read(b)

				if e != nil {
					logs.Errorf("Socket read net(%s) local(%s) remote(%s) error(%s) ", c.conn.LocalAddr().Network(), c.conn.LocalAddr(), c.conn.RemoteAddr(), e.Error())
					c.Cancel()

				} else {
					c.recvBytes += i
					c.recvChannel <- b[:i]
				}
			}
			time.Sleep(time.Millisecond)
		}
	}(rdContext)

	//write routine
	wrContext, _ := context.WithCancel(c.ctx)
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Errorf("Conn write routine panic %s", r)
			}

		}()

		for {
			select {
			case <-ctx.Done():
				close(c.sendChannel)
				return
			case b := <-c.sendChannel:
				o, e := c.conn.Write(b)
				if e != nil {
					logs.Errorf("Socket write net(%s) local(%s) remote(%s) error(%s) ", c.conn.LocalAddr().Network(), c.conn.LocalAddr(), c.conn.RemoteAddr(), e.Error())
					c.Cancel()
				} else {
					logs.Debugf("write routine bytes %d", o)
					c.sendBytes += o
				}
			}
			time.Sleep(time.Millisecond)
		}
	}(wrContext)

	return c
}
