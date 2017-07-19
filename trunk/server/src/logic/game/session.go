package game

import (
	"fmt"
	"logic/prpc"
	"logic/socket"
)

type Session struct {
	prpc.COM_ServerToClientStub
	peer *socket.Peer

	player *GamePlayer
}

func (this *Session) Login(info prpc.COM_LoginInfo) error {
	fmt.Println("Login", info)
	return nil
} // 0
func (this *Session) CreatePlayer(tempId int32, playerName string) error {

	fmt.Println("CreatePlayer11", tempId,playerName )
	this.player = CreatePlayer(tempId,playerName)

	r := this.player.GetPlayerCOM()

	this.CreatePlayerOK(r)

	return nil
} // 1
func (this *Session) SetBattleUnit(instId int64) error {
	return nil
} // 2
func (this *Session) JoinBattle() error {
	return nil
} // 3
func (this *Session) SetupBattle(positionList []prpc.COM_BattlePosition) error {
	return nil
} // 4

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *Session) Update() {
	for {
		err := this.peer.HandleSocket()
		if err != nil {
			fmt.Println(err)
		}
		if this.peer.IncomingBuffer.Len() >= 2 {
			err := prpc.COM_ClientToServerDispatch(this.peer.IncomingBuffer, this)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func NewClient(peer *socket.Peer) *Session {
	c := Session{}
	c.Sender = peer
	c.peer = peer
	return &c
}
