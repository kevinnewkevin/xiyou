package network

import (
	"jimny/logs"
	"testing"

	"context"
	"time"
	"fmt"
	"runtime"
)

//func TestAcceptor_Open(t *testing.T) {
//
//	done := make(chan struct{})
//
//	socks0 := make([]*Conn, 0)
//	socks1 := make([]*Conn, 0)
//	acceptor := NewAcceptorC("tcp", "127.0.0.1:12999")
//
//	go func() {
//		for {
//			select {
//			case sock := <-acceptor.Accept():
//				//logs.Info("0:", sock.I().RemoteAddr(),sock.I().LocalAddr())
//				socks0 = append(socks0, sock)
//				//sock.Close()
//			case <-done:
//				close(done)
//				return
//			}
//		}
//	}()
//
//	go func() {
//		connector := NewConnectorC()
//
//		for i := 0; i < 10000; i++ {
//			sock := connector.Connect("tcp", "127.0.0.1:12999")
//			socks1 = append(socks1, sock)
//			//sock.Close()
//		}
//		time.Sleep(time.Second)
//		connector.Cancel()
//	}()
//
//	//time.Sleep(time.Second * 17)
//
//	logs.Info(len(socks0), " ", len(socks1))
//
//	done <- struct {
//	}{}
//
//	acceptor.Cancel()
//
//}

func TestConn_Select(t *testing.T) {

	var (
		conn0 *Conn
		conn1 *Conn

	)
	ctx, _  := context.WithCancel(context.Background())
	acceptor := NewAcceptorC("tcp", "127.0.0.1:18999")
	go func() {

		for {
			select {
			case sock := <-acceptor.Accept():
				//logs.Info("0:", sock.I().RemoteAddr(),sock.I().LocalAddr())
				conn0 = sock
			//sock.Close()
			case <-ctx.Done():
				return
			}
		}
	}()

	connector := NewConnectorC()

	conn1 = connector.Connect("tcp", "127.0.0.1:18999")

	time.Sleep(time.Second)


	go func() {

		b,o := conn0.Open()
		for i := 0 ; i< 100; i++ {
			b <- []byte(fmt.Sprint("Conn1 said ", i))
			logs.Debug(string(<-o))
		}
	}()


	go func() {

		b,o := conn1.Open()
		for i := 0 ; i< 100; i++ {
			b <- []byte(fmt.Sprint("Conn0 said ", i))
			logs.Debug(string(<-o))
		}
	}()

	time.Sleep(time.Second * 10)


	logs.Debug(runtime.NumGoroutine(),conn0.Status(),conn1.Status())
	logs.Debug(conn0.TotalIncoming(),conn0.TotalOutgoing())


	acceptor.Cancel()
	connector.Cancel()

}
