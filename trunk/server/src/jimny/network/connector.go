package network

import (
	"net"
	"context"
)

type Connector struct {
	ctx context.Context
	lastError error
	Cancel func()
}

func (c *Connector) Connect(family, host string) *Conn {
	s, e := net.Dial(family, host)
	if e == nil {
		return newConn(c.ctx,s)
	}
	c.lastError = e

	return nil
}

func NewConnector(ctx context.Context)*Connector{
	c := &Connector{}
	c.ctx, c.Cancel = context.WithCancel(ctx)
	return c
}

func NewConnectorC()*Connector{
	return NewConnector(context.Background())
}