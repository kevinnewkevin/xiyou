package game

import (
	"fmt"
	"logic/prpc"
	"logic/socket"
	"encoding/json"
)

type Session struct {
	prpc.COM_ServerToClientStub
	peer *socket.Peer
	username string
	player *GamePlayer
}

func (this *Session) Login(info prpc.COM_LoginInfo) error {
	fmt.Println("Login", info)
	infoext := prpc.COM_AccountInfo{}
	infoext.SessionCode = info.Username + info.Password

	this.username = info.Username

	p := prpc.SGE_DBPlayer{Username:info.Username}

	if QueryPlayer(&p){
		this.player = &GamePlayer{}
		this.player.SetSession(this)

		this.player.SetPlayerSGE(p)
		infoext.MyPlayer = p.COM_Player
		fmt.Println(infoext.MyPlayer.UnitGroup)
		fmt.Println(p.Employees)
	}

	this.LoginOK(infoext)

	if this.player != nil{
		this.player.SyncBag()
	}

	return nil
} // 0
func (this *Session) CreatePlayer(tempId int32, playerName string) error {

	this.player = CreatePlayer(tempId, playerName)
	this.player.SetSession(this)
	this.player.Username = this.username

	r := this.player.GetPlayerSGE()

	InsertPlayer(r)

	this.CreatePlayerOK(r.COM_Player)

	this.player.SyncBag()
	fmt.Println(tempId, "CreatePlayer", &r)
	b,_ := json.Marshal(r)
	fmt.Println(string(b))

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

func (this *Session) SetupBattle(positionList []prpc.COM_BattlePosition, skillid int32) error {
	//fmt.Println("SetupBattle", positionList)
	r := this.player.SetupBattle(positionList, skillid)

	if r != nil {
		return nil
	}

	this.SetupBattleOK()
	//fmt.Println("SetupBattleOK", positionList)

	return nil
} // 5

func (this *Session)RequestChapterData(chapterId int32) error {
	if this.player == nil {
		return nil
	}
	this.player.GetMyChapterDataById(chapterId)

	return nil
}

func (this *Session)ChallengeSmallChapter(smallChapterId int32) error {
	if this.player == nil {
		return nil
	}
	this.player.AttackChapter(smallChapterId)

	return nil
}

func (this *Session)SetBattleUnit(instId int64, groupId int32, isBattle bool) error {
	if this.player == nil {
		return nil
	}

	this.player.SetBattleUnitGroup(instId,groupId,isBattle)

	return nil
}

func (this *Session)DelUnitGroup(groupId int32) error {
	if this.player == nil {
		return nil
	}
	this.player.DeleteUnitGroup(groupId)
	return nil
}

func  (this *Session)StartMatching(groupId int32 ) error  {
	if this.player == nil {
		return nil
	}
	StartMatching(this.player,groupId)
	return nil
}

func (this *Session)StopMatching() error {
	if this.player == nil {
		return nil
	}
	StopMatching(this.player)
	return nil
}

func (this *Session)DeleteItem(instId int64, stack int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.DelItemByInstId(instId,stack)
	return nil
}

func (this *Session)PromoteUnit(instId int64) error  {
	if this.player == nil {
		return nil
	}
	this.player.PromoteUnit(instId)
	return nil
}

func (this *Session)RequestChapterStarReward(chapterId int32, star int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.GetChapterStarReward(chapterId,star)
	return nil
}

func (this *Session)EquipSkill(skillinfo prpc.COM_LearnSkill) error  {
	if this.player == nil {
		return nil
	}
	this.player.EquipSkill(skillinfo)
	return nil
}

func (this *Session)SkillUpdate(skillindex int32, skillId int32) error  {
	if this.player == nil {
		return nil
	}
	this.player.SkillUpdate(skillindex, skillId)
	return nil
}

func (this *Session)BuyShopItem(shopId int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.BuyShopItem(shopId)
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

	if this.player != nil {
		this.player.Save()
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
