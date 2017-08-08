package game

import (
	"fmt"
	"logic/prpc"
	"sync"
	"sync/atomic"
	"time"
)

const (
	kIdle = 0 // 无效状态
	kUsed = 1 // 使用状态

	kTurn    = 1 //回合數
	kMaxUnit = 6 //雙方最多上陣卡牌
	kMaxMove = 2 //行动结束

	kTimeSleep = 5   //檢測間隔
	kTimeMax   = 500 //戰鬥持續時間
)

var roomInstId int64 = 1

type BattleRoom struct {
	sync.Mutex
	InstId     int64         //房间ID
	Status     int32         //战斗房间状态
	Round      int32         //回合计数
	Units      []*GameUnit   //当前战斗中牌 数组索引跟下面玩家对应
	PlayerList []*GamePlayer //房间中玩家信息
	Turn       int32
	Winner     int //获胜者
}

var BattleRoomList = map[int64]*BattleRoom{} //所有房间

////////////////////////////////////////////////////////////////////////
////创建部分
////////////////////////////////////////////////////////////////////////

func CreateBattle(p0 *GamePlayer, p1 *GamePlayer) *BattleRoom {

	if p0 == p1 {
		return nil
	}

	room := BattleRoom{}
	room.Status = kUsed
	room.InstId = atomic.AddInt64(&roomInstId, 1)
	room.Round = 0
	room.Winner = 0
	room.Units = make([]*GameUnit, prpc.BP_MAX)
	room.PlayerList = append(room.PlayerList, p0, p1)
	p0.BattleId = room.InstId
	p1.BattleId = room.InstId

	p0.BattleCamp = prpc.CT_RED
	p1.BattleCamp = prpc.CT_BLUE

	BattleRoomList[room.InstId] = &room
	fmt.Println("CreateBattleRoom", &room)

	room.BattleStart()
	go room.BattleUpdate()

	return &room
}

func FindBattle(battleId int64) *BattleRoom {
	if room, ok := BattleRoomList[battleId]; ok {
		return room
	}
	return nil
}

func (this *BattleRoom) BattleStart() {
	for _, p := range this.PlayerList {

		if p == nil || p.session == nil {
			continue
		}
		p.session.JoinBattleOk(int32(p.BattleCamp))
	}
}

//臨時用
func (this *BattleRoom) BattleUpdate() {
	start := time.Now().Unix()
	now_start := time.Now().Unix()
	checkindex := 0
	for this.Status == kUsed {

		now := time.Now().Unix()

		if now-kTimeSleep < now_start { //每隔5S檢測一次
			continue
		}

		if now-kTimeMax >= start { //超時直接結束 並且沒有勝負方
			this.Status = kIdle
			this.BattleRoomOver(prpc.CT_MAX)
			continue
		}

		fmt.Println("BattleUpdate, roomId is ", this.InstId, "index is", checkindex)
		this.Update()
		now_start = time.Now().Unix()
		checkindex += 1
	}
}

////////////////////////////////////////////////////////////////////////
////销毁部分
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) BattleRoomOver(camp int) {
	for _, p := range this.PlayerList {

		if p == nil || p.session == nil {
			continue
		}

		var money int32
		var win int32
		if p.BattleCamp == camp {
			money = 2000
			win = 1
		} else {
			money = 1000
			win = 0
		}

		result := prpc.COM_BattleResult{}

		result.Money = money
		result.Win = win
		p.session.BattleExit(result)
		fmt.Println("BattleRoomOver, result is ", result, "player is ", p.MyUnit.InstId)
	}

	fmt.Println("BattleRoomOver, winner is ", camp)
}

////////////////////////////////////////////////////////////////////////
////回合操作
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) Update() {
	this.Lock()
	defer this.Unlock()
	for _, p := range this.PlayerList {
		if !p.IsActive {
			return
		}
	}

	//顺序排序

	report := prpc.COM_BattleReport{}

	WinCamp := 0
	for _, u := range this.Units {
		if u == nil {
			continue
		}

		fmt.Println("report.UnitList, append", u)
		report.UnitList = append(report.UnitList, u.GetBattleUnitCOM())

		if u.IsDead() { // 非主角死亡跳過
			continue
		}

		ac, ownerdead := u.CastSkill(this)
		report.ActionList = append(report.ActionList, ac)

		fmt.Println("BattleAcction", u.InstId, "acction", ac)
		if ownerdead {
			this.Status = kIdle
			WinCamp = u.Owner.BattleCamp
			this.Winner = u.Owner.BattleCamp
			for _, a := range ac.TargetList {
				unit := this.SelectOneUnit(a.InstId)
				if unit == nil {
					continue
				}
				fmt.Println("append", unit)
				report.UnitList = append(report.UnitList, unit.GetBattleUnitCOM())
			}

		}

		if this.calcWinner() == true {
			this.Round += 1
			this.SendReport(report)
			this.BattleRoomOver(WinCamp)
			this.Status = kIdle
			break
		}
	}
	fmt.Println("Battle report", report)

	for _, p := range this.PlayerList { //戰鬥結束之後要重置屬性
		p.IsActive = false
	}

	this.Round += 1
	this.SendReport(report)

}

func (this *BattleRoom) calcWinner() bool {
	return this.Winner != 0
}

func (this *BattleRoom) SendReport(report prpc.COM_BattleReport) {
	for _, p := range this.PlayerList {
		p.session.BattleReport(report)
	}
}

////////////////////////////////////////////////////////////////////////
////數據處理
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SelectAllTarget(camp int) []*GameUnit {
	targets := []*GameUnit{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.Owner.BattleCamp == camp {
			continue
		}
		targets = append(targets, u)
	}

	return targets
}

func (this *BattleRoom) SelectOneUnit(instid int64) *GameUnit {
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.InstId == instid {
			return u
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SetupPosition(p *GamePlayer, posList []prpc.COM_BattlePosition) {
	this.Lock()
	defer this.Unlock()
	fmt.Println("SetupPosition.start")
	if this.Round == 0 { //第一回合 必须设置主角卡
		for _, pos := range posList {
			fmt.Println("SetupPosition, set ", pos.InstId, p.MyUnit.InstId)
			if pos.InstId == p.MyUnit.InstId {
				goto setup_check_success
			}
			return //没有主角卡
		}
	}
setup_check_success:
	for i := 0; i < len(posList)-1; i++ {
		for j := i + 1; j < len(posList); j++ {
			if posList[i].InstId == posList[j].InstId {
				return //有重复卡牌设置
			}
			if posList[i].Position == posList[j].Position {
				return //有重复位置设置
			}
		}
	}
	for _, pos := range posList {
		for _, u := range this.Units {
			if u == nil {
				continue
			}
			if u.InstId == pos.InstId {
				return //已经上场
			}
		}

		if this.Units[pos.Position] != nil {
			return //这个位置上有人
		}
	}

	//处理数据
	for _, pos := range posList {
		this.Units[pos.Position] = p.GetUnit(pos.InstId)
		this.Units[pos.Position].Position = pos.Position
	}
	p.IsActive = true

	fmt.Println("SetupPosition", this.Units)
}
