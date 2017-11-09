package network

import (
	"net"
	"context"
	"jimny/logs"
)



type acceptor struct {
	ln        net.Listener
	ctx 	  context.Context
	queue     chan *Conn
	cancelFunc 	func()
}

func (a *acceptor) Cancel() {
	a.cancelFunc()
	if a.ln != nil {
		a.ln.Close()
	}
}

func (a *acceptor) Accept() <-chan *Conn {
	return a.queue
}

func NewAcceptor(c context.Context, family, host string) *acceptor {
	a := &acceptor{}
	l, e := net.Listen(family, host)
	if e != nil {
		logs.Errorf("Listen %s %s %s", family, host, e.Error())
		return a
	}
	a.ln = l
	a.ctx,a.cancelFunc= context.WithCancel(c)
	a.queue = make(chan *Conn, defaultChannelBufferMax)

	ctx, _ := context.WithCancel(a.ctx)
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("Acceptor accept routine panic %s",r)
			}

		}()
		for {
			select {
			case <- ctx.Done():
				close(a.queue)
				return
			default:
				c, e := a.ln.Accept()
				if e == nil {
					a.queue <- newConn(a.ctx, c)
				}else{
					logs.Errorf("Accept %s %s %s", family, host, e.Error())
				}
			}
		}
	}(ctx)

	return a
}

func NewAcceptorC(family, host string)*acceptor{
	return NewAcceptor(context.Background(),family,host)
}