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

	BattleRoomList[room.InstId] = &room
	fmt.Println("CreateBattleRoom", &room)

	return &room
}

func FindBattle(battleId int64) *BattleRoom {
	if room, ok := BattleRoomList[battleId]; ok {
		return room
	}
	return nil
}

////////////////////////////////////////////////////////////////////////
////销毁部分
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) BattleRoomOver() {
	this.Status = kIdle
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

	for _, u := range this.Units {
		report.UnitList = append(report.UnitList, u.GetBattleUnitCOM())

		u.CastSkill(this)
	}
}

func (this *BattleRoom) SendReport(report prpc.COM_BattleReport) {
	for _, p := range this.PlayerList {
		p.session.BattleReport(report)
	}
}

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SetupPosition(p *GamePlayer, posList []prpc.COM_BattlePosition) {
	if this.Round == 0 { //第一回合 必须设置主角卡
		for _, pos := range posList {
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
		this.Units[pos.Position] = p.GetBattleUnit(pos.InstId)
		this.Units[pos.Position].Position = pos.Position
	}
}

//func (this *BattleRoom) BattleAction(Player *BattlePlayer, AllPlayer map[int64]*BattlePlayer) {
//	fmt.Println("BattleAcction start", Player.Player.MyUnit.InstId)
//	report := prpc.COM_BattleReport{}
//	fmt.Println("BattleAcction report start",report.UnitList, "ssssss", report.ActionList)
//
//	for _, bp := range AllPlayer {				//遍历所有参与战斗的
//		for _, unit := range bp.BattlePosition {		//遍历本人所有战斗单位
//			instid := unit.InstId
//			//u := bp.Player.GetBattleUnit(instid)
//
//			var u *GameUnit
//			u = bp.Player.GetUnit(instid)
//
//			if u == nil {
//				continue
//			}
//
//			skill := u.Skill[0]
//
//			targetPlayer, _ := this.Target[bp.Player.MyUnit.InstId]
//			targetList := []*GameUnit{}
//
//			for _, pos := range targetPlayer.BattlePosition {
//				e_u := targetPlayer.Player.GetBattleUnit(pos.InstId)
//				if e_u == nil{
//					continue
//				}
//				targetList = append(targetList, e_u)
//			}
//
//			if len(targetList) == 0{
//				targetList = append(targetList, targetPlayer.Player.MyUnit)
//			}		//测试用代码
//
//			unit_con := prpc.COM_BattleUnit{}
//
//			unit_con.InstId = unit.InstId
//			unit_con.Position = unit.Position
//			unit_con.Camp = bp.Camp
//			unit_con.UnitId = u.UnitId
//			unit_con.Name = u.InstName
//			unit_con.HP = int32(u.CProperties[prpc.CPT_HP])
//
//			report.UnitList = append(report.UnitList, unit_con)
//
//			fmt.Println("UnitId", unit_con.UnitId, "Instid ", unit_con.InstId)
//
//			al, dl := skill.Action(u, targetList, this.Bout)
//
//			if len(dl) > 0{
//				//这里需要在战斗单位中删除某些角色
//				for _, iid := range dl {
//					delete(bp.BattlePosition, iid)
//				}
//				if len(bp.BattlePosition) == 0{
//				}
//			}
//
//			fmt.Println("BattleAcction", bp.Player.MyUnit.InstId, "target", targetPlayer.Player.MyUnit.InstId, "acction", al)
//			report.ActionList = append(report.ActionList, al)
//			fmt.Println("Battle report", report)
//		}
//	}
//
//	fmt.Println("Battle end, unitlist is ", report.UnitList, "		action is ", report.ActionList)
//	this.SendReport(report)
//
//}
