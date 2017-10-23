package socket

import (
	"net"
	"logic/log"
	"fmt"
)

type TCPServer struct {
	Addr string
	MaxConnNum int
	PendingWriteNum int
	running bool
	ln net.TCPListener
	connList []*TcpConnection
}

func (this *TCPServer)Open(addr string){

}

func (this *TCPServer)Close(){

}

func (this *TCPServer)Run(){
	this.running = true

	go func() {
		for this.running {

			defer func() {

				if r := recover(); r != nil {
					log.Error("Tcp server panic %s",fmt.Sprint(r))
				}

			}()
			con, err := this.ln.AcceptTCP()
			if err != nil {
				log.Debug("Tcp socket error %s ", err.Error())
				this.running = false
			}
			log.Debug("One socket connect %s <==> %s", con.LocalAddr(), con.RemoteAddr())

			tcpcon := NewTcpConnection(con,this.PendingWriteNum)

			this.connList = append(this.connList,tcpcon)

			tcpcon.Run()
		}
	}()
}
