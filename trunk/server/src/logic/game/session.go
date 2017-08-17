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

	this.player = CreatePlayer(tempId, playerName)
	this.player.SetSession(this)

	r := this.player.GetPlayerCOM()

	this.CreatePlayerOK(r)

	fmt.Println(tempId, "CreatePlayer", &r)

	return nil
} // 1
func (this *Session) AddBattleUnit(instId int64, groupId int32) error {
	fmt.Println("SetBattleUnit", instId)
	this.player.SetBattleUnit(instId)

	r := this.player.GetBattleUnit(instId)

	this.SetBattleUnitOK(r.InstId)

	fmt.Println("SetBattleUnitOK")

	return nil
} // 2

func (this *Session) PopBattleUnit(instId int64, groupId int32) error {
	fmt.Println("SetBattleUnit", instId)
	this.player.SetBattleUnit(instId)

	r := this.player.GetBattleUnit(instId)

	this.SetBattleUnitOK(r.InstId)

	fmt.Println("SetBattleUnitOK")

	return nil
} // 3

func (this *Session) JoinBattle() error {

	if this.player == nil {
		return nil
	}

	this.player.JoinBattle()

	return nil
} // 4

func (this *Session) SetupBattle(positionList []prpc.COM_BattlePosition) error {
	//fmt.Println("SetupBattle", positionList)
	r := this.player.SetupBattle(positionList)

	if r != nil {
		return nil
	}

	this.SetupBattleOK()
	//fmt.Println("SetupBattleOK", positionList)

	return nil
} // 5


func (this* Session) JoinBattlePvE(a,b int32)error{
	return nil
}
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

	if this.player != nil && this != nil {
		this.player.SetSession(nil)
		this.player = nil
		this.peer = nil

		fmt.Println("Socket close ")
	}

}

func NewClient(peer *socket.Peer) *Session {
	c := Session{}
	c.Sender = peer
	c.peer = peer
	return &c
}
