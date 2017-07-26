package game

import (
	"fmt"
	"logic/prpc"
	"sync/atomic"
)

const (
	kIdle = 0 // 无效状态
	kUsed = 1 // 使用状态

	kTurn    = 1 //回合數
	kMaxUnit = 6 //雙方最多上陣卡牌
	kMaxMove = 2 //行动结束
)

var roomInstId int64 = 1

type BattleRoom struct {
	InstId     int64         //房间ID
	Status     int32         //战斗房间状态
	Round      int32         //回合计数
	Units      []*GameUnit   //当前战斗中牌 数组索引跟下面玩家对应
	PlayerList []*GamePlayer //房间中玩家信息
	Turn       int32
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
	room.Units = make([]*GameUnit, prpc.BP_MAX)
	room.PlayerList = append(room.PlayerList, p0, p1)
	p0.BattleId = room.InstId
	p1.BattleId = room.InstId

	p0.BattleCamp = prpc.CT_RED
	p1.BattleCamp = prpc.CT_BLUE

	BattleRoomList[room.InstId] = &room
	fmt.Println("CreateBattleRoom", &room)

	room.BattleStart()

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
		p.session.JoinBattleOk()
	}
}

////////////////////////////////////////////////////////////////////////
////销毁部分
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) BattleRoomOver(camp int) {
	for _, p := range this.PlayerList {
		var money int32
		if p.BattleCamp == camp {
			money = 1000
		} else {
			money = 2000
		}

		result := prpc.COM_BattleResult{}

		result.Money = money
		p.session.BattleExit(result)
	}

	fmt.Println("BattleRoomOver, winner is ", camp)
}

////////////////////////////////////////////////////////////////////////
////回合操作
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) Update() {
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
		if u.IsDead() {			// 非主角死亡跳過
			continue
		}

		report.UnitList = append(report.UnitList, u.GetBattleUnitCOM())

		ac, ownerdead := u.CastSkill(this)
		report.ActionList = append(report.ActionList, ac)

		fmt.Println("BattleAcction", u.InstId, "acction", ac)
		if ownerdead {
			this.Status = kIdle
			WinCamp = u.Owner.BattleCamp
			break
		}
	}
	fmt.Println("Battle report", report)

	for _, p := range this.PlayerList {		//戰鬥結束之後要重置屬性
		p.IsActive = false
	}

	this.Round += 1
	this.SendReport(report)

	if this.Status == kIdle {
		this.BattleRoomOver(WinCamp)
	}
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

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SetupPosition(p *GamePlayer, posList []prpc.COM_BattlePosition) {
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

	this.Update()
}
