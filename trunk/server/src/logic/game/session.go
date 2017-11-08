package game

import (
	"logic/prpc"
	"logic/log"
	"fmt"
	"strings"
	"bytes"
	"runtime"
	"net"
	"sync"
)

type Session struct {
	sync.Mutex
	prpc.COM_ServerToClientStub

	username string
	player *GamePlayer

	TotalIncoming, TotalOutgoing   int
	IncomingBuffer, OutgoingBuffer *bytes.Buffer
	Connection                     *net.TCPConn

}

func (this *Session) Login(info prpc.COM_LoginInfo) error {
	log.Println("Login ", info)
	infoext := prpc.COM_AccountInfo{}
	infoext.SessionCode = info.Username + info.Password

	this.username = info.Username

	this.player = FindPlayerByUsername(info.Username)

	if this.player == nil {
		p := prpc.SGE_DBPlayer{Username:info.Username}

		if QueryPlayer(&p) {
			this.player = &GamePlayer{}
			this.player.SetSession(this)
			this.player.SetPlayerSGE(p)
			if FindPlayerByUsername(info.Username) != nil{
				for k, n:=range PlayerStore{
					if n == nil {
						continue
					}
					if n.Username == info.Username {
						PlayerStore[k] = this.player
					}
				}
			}else {
				PlayerStore = append(PlayerStore,this.player)
			}
			infoext.MyPlayer = p.COM_Player

			log.Println("Query player ", p)
		}
	}else{
		this.player.SetSession(this)
		infoext.MyPlayer = this.player.GetPlayerCOM()
	}

	this.LoginOK(infoext)
	if this.player != nil{
		this.player.PlayerLogin()
	}

	return nil
} // 0
func (this *Session) CreatePlayer(tempId int32, playerName string) error {

	if FindPlayerByInstName(playerName) != nil{
		return nil
	}

	this.player = CreatePlayer(tempId, playerName)
	this.player.SetSession(this)
	this.player.Username = this.username

	if this.player != nil{
		this.player.PlayerLogin()
	}

	r := this.player.GetPlayerSGE()

	InsertPlayer(r)

	this.CreatePlayerOK(r.COM_Player)

	log.Println("CreatePlayer ", r)

	return nil
} // 1
func (this *Session) AddBattleUnit(instId int64, groupId int32) error {
	log.Info("SetBattleUnit", instId)
	this.player.SetBattleUnit(instId)

	r := this.player.GetBattleUnit(instId)

	this.SetBattleUnitOK(r.InstId)

	log.Info("SetBattleUnitOK")

	return nil
} // 2

func (this *Session) PopBattleUnit(instId int64, groupId int32) error {
	log.Info("SetBattleUnit", instId)
	this.player.SetBattleUnit(instId)

	r := this.player.GetBattleUnit(instId)

	this.SetBattleUnitOK(r.InstId)

	log.Info("SetBattleUnitOK")

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
	log.Info("SetupBattle", positionList)
	r := this.player.SetupBattle(positionList, skillid)

	if r != nil {
		return nil
	}

	this.SetupBattleOK()

	return nil
} // 5

func (this *Session) RequestChapterData(chapterId int32) error {
	if this.player == nil {
		return nil
	}
	this.player.GetMyChapterDataById(chapterId)

	return nil
}

func (this *Session) ChallengeSmallChapter(smallChapterId int32) error {
	if this.player == nil {
		return nil
	}
	log.Info("1");
	this.player.AttackChapter(smallChapterId)

	return nil
}

func (this *Session) SetBattleUnit(instId int64, groupId int32, isBattle bool) error {
	if this.player == nil {
		return nil
	}

	this.player.SetBattleUnitGroup(instId,groupId,isBattle)

	return nil
}

func (this *Session) DelUnitGroup(groupId int32) error {
	if this.player == nil {
		return nil
	}
	this.player.DeleteUnitGroup(groupId)
	return nil
}

func  (this *Session) StartMatching(groupId int32 ) error  {
	if this.player == nil {
		return nil
	}
	StartMatching(this.player,groupId)
	return nil
}

func (this *Session) StopMatching() error {
	if this.player == nil {
		return nil
	}
	StopMatching(this.player)
	return nil
}

func (this *Session) DeleteItem(instId int64, stack int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.DelItemByInstId(instId,stack)
	return nil
}

func (this *Session) PromoteUnit(instId int64) error  {
	if this.player == nil {
		return nil
	}
	this.player.PromoteUnit(instId)
	return nil
}

func (this *Session) RequestChapterStarReward(chapterId int32, star int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.GetChapterStarReward(chapterId,star)
	return nil
}

func (this *Session) EquipSkill(skillinfo prpc.COM_LearnSkill) error  {
	if this.player == nil {
		return nil
	}
	this.player.EquipSkill(skillinfo)
	return nil
}

func (this *Session) SkillUpdate(skillindex int32, skillId int32) error  {
	if this.player == nil {
		return nil
	}
	this.player.SkillUpdate(skillindex, skillId)
	return nil
}

func (this *Session) BuyShopItem(shopId int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.BuyShopItem(shopId)
	return nil
}

func (this *Session)ResolveItem(instId int64, num int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.CardDebrisResolve(instId,num)
	return nil
}

func (this *Session)RefreshBlackMarkte() error {
	if this.player == nil {
		return nil
	}
	this.player.RefreshMyBlackMarket(true)
	return nil
}

func (this *Session)NewPlayerGuide(Step uint64) error {
	if this.player == nil {
		return nil
	}
	this.player.NewPlayerGuide(Step)
	return nil
}

func (this *Session)SendChat(content prpc.COM_Chat ) error  {
	if this.player == nil {
		return nil
	}
	this.player.SendChat(content)
	return nil
}

func (this *Session)RequestAudio(audioId int64 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.RequestAudio(audioId)
	return nil
}

func (this *Session) AllTopByPage (page int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.AllTopByPage(page)
	return nil
}

func (this *Session) FriendTopByPage (page int32 ) error  {
	if this.player == nil {
		return nil
	}
	this.player.AllTopByPage(page)
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *Session) Update() {

	for {
		defer func() {

			if r := recover(); r != nil {
				log.Error("main panic %s",fmt.Sprint(r))
			}

		}()
		err := this.HandleSocket()
		if err != nil {
			log.Info("%s", err.Error())
			goto endLoop
		}
		if this.IncomingBuffer.Len() >= 2 {
			err := prpc.COM_ClientToServerDispatch(this.IncomingBuffer, this)
			if err != nil {
				log.Error("err", err)
				goto endLoop
			}
		}
	}
endLoop:

	//do clean

	if this.player != nil {
		this.player.Logout()
		this.player.SetSession(nil)
		this.player = nil
		this.Sender = nil

		log.Info("Socket close ")
	}

}

//////////////////////////////////////////////////////////////
func (this *Session) MethodBegin() *bytes.Buffer {
	this.Lock()
	return this.OutgoingBuffer
}

func (this *Session) MethodEnd() error {
	defer this.Unlock()
	c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
	if e != nil {
		return e
	}
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		file = "???"
		line = 0
	}else {

		i := strings.LastIndex(file, "/")
		if i == -1 {
			i = strings.LastIndex(file, "\\")
		}
		if i != -1{
			file = file[i+1:]
		}

	}
	log.Info("CALL %s:%d MethodEnd() %d ",file,line, c)
	this.OutgoingBuffer.Reset()
	this.TotalOutgoing += c
	return nil
}

func (this *Session) HandleSocket() error {
	{
		bs := make([]byte, 2048)

		c, e := this.Connection.Read(bs)
		if e != nil {
			log.Error("Incoming error %s",e.Error())
			return e
		}
		this.IncomingBuffer.Write(bs[:c])
		this.TotalIncoming += c
	}

	{
		c, e := this.Connection.Write(this.OutgoingBuffer.Bytes())
		if e != nil {
			log.Error("Outgoing error %s",e.Error())

			return e
		}
		this.TotalOutgoing += c
	}
	return nil
}

func NewClient(conn *net.TCPConn) *Session {
	c := Session{}
	conn.SetNoDelay(true)
	c.Connection = conn
	c.OutgoingBuffer =  bytes.NewBuffer(nil)
	c.IncomingBuffer =  bytes.NewBuffer(nil)
	c.Sender = &c
	return &c
}
