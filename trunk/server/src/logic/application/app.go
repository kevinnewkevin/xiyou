package application

import (
	"net"
	"suzuki/logs"
	"logic/socket"
	"logic/handle"
	"fmt"
)



type App struct {
	l net.Listener
}

func (this *App) Run() {

	var (
		err        error
		conn       net.Conn
		endRunning = make(chan bool, 1)
	)

	logs.Init()
	this.l, err = net.Listen("tcp", "0.0.0.0:10999")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go func() {
		for {




			conn, err = this.l.Accept()
			if err != nil {
				fmt.Println(err.Error())
				endRunning <- true
			}
			fmt.Println("Has one connect ")
			peer := socket.NewPeer(conn)
			client := game.NewClient(peer)
			//
			go client.Update()


		}
	}()
	logs.Fini()
	<-endRunning
}

func NewApp() *App {
	return &App{}
}
