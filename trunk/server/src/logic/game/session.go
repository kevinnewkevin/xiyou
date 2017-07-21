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

	this.player = CreatePlayer(tempId,playerName)
	this.player.SetSession(this)

	r := this.player.GetPlayerCOM()

	this.CreatePlayerOK(r)

	return nil
} // 1
func (this *Session) SetBattleUnit(instId int64) error {
	this.player.SetBattleUnit(instId)
	return nil
} // 2
//dont care mutli thread
var battlePlayerList = []*GamePlayer{}
func (this *Session) JoinBattle() error {
	fmt.Println("JoinBattle", battlePlayerList)
	battlePlayerList = append(battlePlayerList, this.player)
	if len(battlePlayerList) == 2{
		//把他俩都拉到战斗力去			这里还要加一个判断,不能重复加入战斗
		CreateBattleRoom(battlePlayerList[0], battlePlayerList[1])

		battlePlayerList = battlePlayerList[:0]
	}
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
			goto endLoop
		}
		if this.peer.IncomingBuffer.Len() >= 2 {
			err := prpc.COM_ClientToServerDispatch(this.peer.IncomingBuffer, this)
			if err != nil {
				fmt.Println("err", err)
				goto endLoop
			}
		}
	}
	endLoop:

	//do clean
	this.player.SetSession(nil)
	this.player = nil
	this.peer = nil

	fmt.Println("Socket close ")
}

func NewClient(peer *socket.Peer) *Session {
	c := Session{}
	c.Sender = peer
	c.peer = peer
	return &c
}
