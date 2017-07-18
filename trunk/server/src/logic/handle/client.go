package handle

import (
	"logic/prpc"
	"logic/socket"
	"fmt"
)



type Client struct {
	prpc.COM_ServerToClientStub
	peer *socket.Peer
}

func (this *Client) Login(info prpc.COM_LoginInfo) error {
	fmt.Println(info)
	return nil
} // 0
func (this *Client) CreatePlayer(tempId int32, playerName string) error {
	fmt.Println(tempId,playerName)

	p := prpc.COM_Player{}
	u := prpc.COM_Unit{}
	u.UnitId = tempId
	u.InstId = 1

	p.InstId = 1
	p.Unit = u
	p.Name = playerName

	u.UnitId = tempId
	u.InstId = 2
	p.Employees = append(p.Employees,u)


	this.CreatePlayerOK(p)

	return nil
} // 1
func (this *Client) SetBattleUnit(instId int64) error {
	return nil
} // 2
func (this *Client) JoinBattle() error {
	return nil
} // 3
func (this *Client) SetupBattle(positionList []prpc.COM_BattlePosition) error {
	return nil
} // 4

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func(this*Client) Update(){
	for{
		//fmt.Println("aaaaaaaaaaaaaaaaaaa")
		err := this.peer.HandleSocket()
		if err != nil {
			fmt.Println(err)
		}
		if this.peer.IncomingBuffer.Len() >= 2{
			err := prpc.COM_ClientToServerDispatch(this.peer.IncomingBuffer,this)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func NewClient(peer *socket.Peer) *Client {
	c := Client{}
	c.Sender = peer
	c.peer = peer
	return  &c
}
