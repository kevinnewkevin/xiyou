package game

import (
	"encoding/json"
	"jimny/logs"
	"logic/prpc"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
	//"fmt"
	"fmt"
)

type UnitList []*GameUnit

func (a UnitList) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a UnitList) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a UnitList) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].CProperties[prpc.CPT_AGILE] < a[i].CProperties[prpc.CPT_AGILE]
}

const (
	kIdle = 0 // 无效状态
	kUsed = 1 // 使用状态

	kTurn    = 1 //回合數
	kMaxUnit = 6 //雙方最多上陣卡牌
	kMaxMove = 2 //行动结束

	kTimeSleep = 3   //檢測間隔
	kTimeMax   = 600 //戰鬥持續時間
)

var roomInstId int64 = 1

type Monster struct {
	sync.Locker
	MainUnit       *GameUnit   //自己的卡片
	BattleUnitList []*GameUnit //拥有的卡片

	//战斗相关辅助信息
	BattleCamp int //阵营 //prpc.CompType
}

type BattleRoom struct {
	sync.Mutex
	Type        int32         //战斗类型 1是pvp 2是pve
	InstId      int64         //房间ID
	BattleID    int32         //戰鬥ID
	Status      int32         //战斗房间状态
	Round       int32         //回合计数
	Point       int32         //本场战斗的能量点
	Units       []*GameUnit   //当前战斗中牌 数组索引跟下面玩家对应
	Dead        []*GameUnit   //本回合死亡的人数
	PlayerList  []*GamePlayer //房间中玩家信息
	Monster     *Monster
	Turn        int32
	Winner      int                         //获胜者
	ReportAll   []prpc.COM_BattleReport     //整场战斗的所有战报
	ReportOne   prpc.COM_BattleReport       //每回合的战报
	AcctionList prpc.COM_BattleAction       //行动单元
	TargetCOM   prpc.COM_BattleActionTarget //行动单元中的每个子元素
	NewAction   bool                        //是否行动过
}

var BattleRoomList = map[int64]*BattleRoom{} //所有房间

////////////////////////////////////////////////////////////////////////
////创建部分
////////////////////////////////////////////////////////////////////////

func CreatePvE(p *GamePlayer, battleid int32) *BattleRoom {
	room := BattleRoom{}

	room.Status = kUsed
	room.InstId = atomic.AddInt64(&roomInstId, 1)
	room.Round = 0
	room.Winner = prpc.CT_MAX
	room.Units = make([]*GameUnit, prpc.BP_MAX)
	room.PlayerList = append(room.PlayerList, p)
	room.Type = prpc.BT_PVE
	room.Point = 1
	room.BattleID = battleid

	logs.Info("CreatePvE", &room)
	BattleRoomList[room.InstId] = &room

	room.Monster = CreateMonster(battleid, room.InstId)
	p.SetProprty(&room, prpc.CT_RED)

	room.Units[prpc.BP_BLUE_5] = room.Monster.MainUnit
	room.Monster.MainUnit.Position = prpc.BP_BLUE_5

	room.Units[prpc.BP_RED_5] = p.MyUnit
	p.MyUnit.Position = prpc.BP_RED_5

	room.BattleStart()
	//go room.BattleUpdate()

	return &room
}
func CreatePvR(p *GamePlayer, battleid int32) *BattleRoom {
	room := BattleRoom{}

	room.Status = kUsed
	room.InstId = atomic.AddInt64(&roomInstId, 1)
	room.Round = 0
	room.Winner = prpc.CT_MAX
	room.Units = make([]*GameUnit, prpc.BP_MAX)
	room.PlayerList = append(room.PlayerList, p)
	room.Type = prpc.BT_PVR
	room.Point = 1
	room.BattleID = battleid

	logs.Info("CreatePvE", &room)
	BattleRoomList[room.InstId] = &room

	room.Monster = CreateMonster(battleid, room.InstId)
	p.SetProprty(&room, prpc.CT_RED)

	room.Units[prpc.BP_BLUE_5] = room.Monster.MainUnit
	room.Monster.MainUnit.Position = prpc.BP_BLUE_5

	room.Units[prpc.BP_RED_5] = p.MyUnit
	p.MyUnit.Position = prpc.BP_RED_5

	room.BattleStart()
	//go room.BattleUpdate()

	return &room
}

func CreateMonster(battleid int32, roomid int64) *Monster {
	t := GetBattleRecordById(battleid)

	m := Monster{}

	m.MainUnit = CreateUnitFromTable(t.MainId)
	logs.Info("adasdas", battleid, t.MainId, m.MainUnit)
	m.MainUnit.ResetBattle(prpc.CT_BLUE, true, roomid)
	//m.MainUnit.IsMain = true
	//m.MainUnit.Camp = prpc.CT_BLUE
	//m.MainUnit.BattleId = roomid

	for _, uid := range t.SmallId {
		t1 := CreateUnitFromTable(uid)
		t1.ResetBattle(prpc.CT_BLUE, false, roomid)
		//t1.Camp = prpc.CT_BLUE
		//t1.IsMain = false
		//t1.BattleId = roomid
		m.BattleUnitList = append(m.BattleUnitList, t1)
	}

	m.BattleCamp = prpc.CT_BLUE

	return &m
}

func CreatePvP(p0 *GamePlayer, p1 *GamePlayer) *BattleRoom {

	if p0 == p1 {
		return nil
	}

	room := BattleRoom{}
	room.Status = kUsed
	room.InstId = atomic.AddInt64(&roomInstId, 1)
	room.Round = 0
	room.Winner = prpc.CT_MAX
	room.Units = make([]*GameUnit, prpc.BP_MAX)
	room.PlayerList = append(room.PlayerList, p0, p1)
	room.Type = prpc.BT_PVP
	room.Point = 1
	room.BattleID = 0
	room.TargetCOM = prpc.COM_BattleActionTarget{}

	//p0.BattleId = room.InstId
	//p1.BattleId = room.InstId
	//
	//p0.BattleCamp = prpc.CT_RED
	//p1.BattleCamp = prpc.CT_BLUE

	BattleRoomList[room.InstId] = &room
	logs.Info("CreatePvP", &room)

	p0.SetProprty(&room, prpc.CT_RED)
	p1.SetProprty(&room, prpc.CT_BLUE)

	room.Units[prpc.BP_BLUE_5] = p1.MyUnit
	p1.MyUnit.Position = prpc.BP_BLUE_5

	room.Units[prpc.BP_RED_5] = p0.MyUnit
	p0.MyUnit.Position = prpc.BP_RED_5

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

func PopBattle(battleId int64) {
	logs.Debug("PopBattle 1 ", BattleRoomList)
	_, ok := BattleRoomList[battleId]
	if !ok {
		return
	}

	delete(BattleRoomList, battleId)
	logs.Debug("PopBattle 2 ", BattleRoomList)

}

func (this *BattleRoom) BattleStart() {

	ul := []prpc.COM_BattleUnit{}

	unitllist := this.SortUnits()

	for _, unit := range unitllist {
		if unit == nil {
			continue
		}
		logs.Debug("卡牌敏捷: 1 ", unit.GetCProperty(prpc.CPT_AGILE))
		ul = append(ul, unit.GetBattleUnitCOM())
	}

	for _, p := range this.PlayerList {

		if p == nil || p.session == nil {
			continue
		}
		targetList := this.findCardsByTarget(p.BattleCamp)
		logs.Debug("JoinBattleOk, p.id", p.MyUnit.InstId, " and batlecamp is ", int32(p.BattleCamp), "targetList is ", targetList)
		p.session.JoinBattleOk(int32(p.BattleCamp), this.BattleID, targetList, ul)
	}
}

func (this *BattleRoom) BattleStrongOver() {
	this.BattleRoomOver(prpc.CT_MAX)
	this.PlayerList = []*GamePlayer{}

	this.Status = kIdle

}

func (this *BattleRoom) PlayerLeft(player *GamePlayer) {
	del_index := -1
	for index, p := range this.PlayerList {
		if p.MyUnit.InstId == player.MyUnit.InstId {
			del_index = index
		}
	}
	if del_index != -1 {
		this.PlayerList = append(this.PlayerList[:del_index], this.PlayerList[del_index+1:]...)
	}

	if len(this.PlayerList) == 0 {
		this.BattleRoomOver(prpc.CT_MAX)
		this.Status = kIdle
	}

}

func (this *BattleRoom) SendReportFirst() {

}

func (this *BattleRoom) findCardsByTarget(camp int) []int32 {
	li := []int32{}
	for _, p := range this.PlayerList {
		if p == nil {
			continue
		}

		if p.BattleCamp == camp {
			continue
		}

		var group *prpc.COM_UnitGroup

		if this.Type == prpc.BT_PVP {
			group = p.GetUnitGroupById(p.BattleUnitGroup)
		} else {
			group = p.GetUnitGroupById(p.BattleUnitGroup)
		}
		if group == nil {
			continue
		}

		for _, instid := range group.UnitList {
			unit := p.GetUnit(instid)
			if unit == nil {
				continue
			}
			li = append(li, unit.UnitId)
		}
	}

	return li
}

//臨時用
func (this *BattleRoom) BattleUpdate() {
	start := time.Now().Unix()
	now_start := time.Now().Unix()
	//Round_start := time.Now().Unix()
	checkindex := 0

	defer func() {

		if r := recover(); r != nil {
			logs.Error("battle panic %s",fmt.Sprint(r))
		}

	}()

	for this.Status == kUsed {

		now := time.Now().Unix()

		if this.Type == prpc.BT_PVP {
			if now-kTimeSleep < now_start { //每隔5S檢測一次
				continue
			}
		}

		if now-kTimeMax >= start { //超時直接結束 並且沒有勝負方
			this.Status = kIdle
			this.BattleRoomOver(prpc.CT_MAX)
			continue

		}
		//if now - Round_start >= 30 {			// 20S不行动就自动帮你打了
		//	for _, p := range this.PlayerList {
		//
		//		//if p == nil || p.session == nil {
		//		//	continue
		//		//}
		//
		//		if p == nil {
		//			continue
		//		}
		//		p.IsActive = true
		//	}
		//	Round_start = time.Now().Unix()
		//}

		logs.Info("BattleUpdate, roomId is ", this.InstId, "index is", checkindex)
		this.Update()
		now_start = time.Now().Unix()
		checkindex += 1
		if this.Type == prpc.BT_PVE || this.Type == prpc.BT_PVR { //pve只执行一次就跳出
			break
		}
	}
}

////////////////////////////////////////////////////////////////////////
////销毁部分
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) BattleRoomOver(camp int) {
	for _, p := range this.PlayerList {

		if p == nil {
			continue
		}

		var win int32

		if p.BattleCamp == camp {
			win = 1
		} else {
			win = 0
		}

		result := prpc.COM_BattleResult{}

		result.Win = win
		result.BattleRound = this.Round
		result.KillMonsters = p.KillUnits
		result.MySelfDeathNum = p.MyDeathNum

		if this.Type == prpc.BT_PVE {
			dropId := p.CalcSmallChapterStar(result)
			if dropId != 0 {
				drop := GetDropById(dropId)
				if drop == nil {
					logs.Info("PVE Can Not Find Drop By DropId=", dropId)
					return
				}
				if drop.Exp != 0 {
					p.AddExp(drop.Exp)
					result.Exp = drop.Exp
				}
				if drop.Money != 0 {
					p.AddCopper(drop.Money)
					result.Money = drop.Money
				}
				if len(drop.Items) != 0 {
					for _, item := range drop.Items {
						p.AddBagItemByItemId(item.ItemId, item.ItemNum)
						logs.Info("PVE GiveDrop AddItem ItemId=", item.ItemId, "ItemNum=", item.ItemNum)
						itemInst := prpc.COM_ItemInst{}
						itemInst.ItemId = item.ItemId
						itemInst.Stack = item.ItemNum
						result.BattleItems = append(result.BattleItems, itemInst)
					}
				}
				if drop.Hero != 0 {
					if p.HasUnitByTableId(drop.Hero) {
						//有这个卡就不给了
						logs.Info("PlayerName=", p.MyUnit.InstName, "GiveDrop AddUnit Have not to UnitId=", drop.Hero)
					} else {
						unit := p.NewGameUnit(drop.Hero)
						if unit != nil {
							logs.Debug("PlayerName=", p.MyUnit.InstName, "GiveDrop AddUnit OK UnitId=", drop.Hero)
							temp := unit.GetUnitCOM()
							if p.session != nil {
								p.session.AddNewUnit(temp)
							}
						}
					}
				}
			}
			logs.Info("Battle Over PVE DropId=", dropId)
		}
		if this.Type == prpc.BT_PVP {
			for _, once := range this.PlayerList {
				if once == nil {
					continue
				}
				if once.MyUnit.InstId == p.MyUnit.InstId {
					continue
				}
				if p.BattleCamp == once.BattleCamp {
					continue
				}
				dropId := CaleTianTiVal(p, once, camp)

				if dropId != 0 {
					drop := GetDropById(dropId)
					if drop == nil {
						logs.Info("PVP Can Not Find Drop By DropId=", dropId)
						return
					}
					if drop.Exp != 0 {
						p.AddExp(drop.Exp)
						result.Exp = drop.Exp
					}
					if drop.Money != 0 {
						p.AddCopper(drop.Money)
						result.Money = drop.Money
					}
					if len(drop.Items) != 0 {
						for _, item := range drop.Items {
							p.AddBagItemByItemId(item.ItemId, item.ItemNum)
							logs.Info("PVP GiveDrop AddItem ItemId=", item.ItemId, "ItemNum=", item.ItemNum)
							itemInst := prpc.COM_ItemInst{}
							itemInst.ItemId = item.ItemId
							itemInst.Stack = item.ItemNum
							result.BattleItems = append(result.BattleItems, itemInst)
						}
					}
				}
			}
		}
		if this.Type == prpc.BT_PVR {
			dropId := CaleTiantiPVR(p, camp)
			if dropId != 0 {
				drop := GetDropById(dropId)
				if drop == nil {
					logs.Info("PVP Can Not Find Drop By DropId=", dropId)
					return
				}
				if drop.Exp != 0 {
					p.AddExp(drop.Exp)
					result.Exp = drop.Exp
				}
				if drop.Money != 0 {
					p.AddCopper(drop.Money)
					result.Money = drop.Money
				}
				if len(drop.Items) != 0 {
					for _, item := range drop.Items {
						p.AddBagItemByItemId(item.ItemId, item.ItemNum)
						logs.Info("PVP GiveDrop AddItem ItemId=", item.ItemId, "ItemNum=", item.ItemNum)
						itemInst := prpc.COM_ItemInst{}
						itemInst.ItemId = item.ItemId
						itemInst.Stack = item.ItemNum
						result.BattleItems = append(result.BattleItems, itemInst)
					}
				}
			}
		}
		p.BattleId = 0
		p.ClearAllBuff()

		if p.session != nil {
			p.session.BattleExit(result)
		}
		logs.Info("BattleRoomOver, result is ", result, "player is ", p.MyUnit.InstId, "p.battlecampis ", p.BattleCamp, "wincampis ", camp, "winis", win)
		p.BattleCamp = prpc.CT_MAX
	}

	logs.Debug("BattleRoomOver, winner is ", camp)
	PopBattle(this.InstId)
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

	if this.Type != prpc.BT_PVP {
		this.MonsterMove()
		//this.PlayerMove()
	}

	//this.CheckPlayerIsMove()

	//顺序排序

	logs.Info("站前回合为 ", this.Round)

	this.ReportOne = prpc.COM_BattleReport{}
	this.Dead = []*GameUnit{}
	this.ReportOne.BattleID = this.BattleID

	unitllist := this.SortUnits()

	for _, unit := range unitllist {
		if unit == nil {
			continue
		}
		//logs.Info("卡牌敏捷: 1 ", unit.GetCProperty(prpc.CPT_AGILE))
		this.ReportOne.UnitList = append(this.ReportOne.UnitList, unit.GetBattleUnitCOM())
	}

	//if this.Round > 0 {
	for _, u := range unitllist {
		if u == nil {
			continue
		}
		//logs.Info("卡牌敏捷 2 : ", u.GetCProperty(prpc.CPT_AGILE))
		//logs.Info("report.UnitList, append", u)
		//this.ReportOne.UnitList = append(this.ReportOne.UnitList, u.GetBattleUnitCOM())

		if u.IsDead() { // 非主角死亡跳過
			continue
		}

		//logs.Info("行动的卡牌信息: ", u.InstName)

		this.AcctionList = prpc.COM_BattleAction{}
		//this.TargetOn()
		this.AcctionList.InstId = u.InstId

		u.CheckBuff(this.Round)
		u.CheckDebuff(this.Round)

		del_buf := u.CheckAllBuff(this.Round)
		this.UpdateBuffState(del_buf)

		if u.IsJump() {
			continue
		}

		u.CastSkill(this)

		//logs.Info("行动的卡牌信息 11 : ", u.InstName, u.ChoiceSKill)
		u.ChoiceSKill = 0

		//this.TargetOver()

		this.ReportOne.ActionList = append(this.ReportOne.ActionList, this.AcctionList)

		//logs.Info("BattleAcction", u.InstId, "acction", this.AcctionList)
		//logs.Info("BattleAcction", u.InstId, "ReportOne", this.ReportOne)
		if this.calcWinner() == true {
			for _, a := range this.AcctionList.TargetList {
				unit := this.SelectOneUnit(a.InstId)
				if unit == nil {
					continue
				}
				//logs.Info("append", unit)
				//this.ReportOne.UnitList = append(this.ReportOne.UnitList, unit.GetBattleUnitCOM())
			}
			//logs.Info("this.Winner", this.Winner)

			this.Round += 1
			this.SendReport(this.ReportOne)
			this.BattleRoomOver(this.Winner)
			this.Status = kIdle
			break

		}
	}
	//}
	logs.Debug("Battle report battleid is ", this.ReportOne.BattleID)
	logs.Debug("Battle report unitlist is ", this.ReportOne.UnitList)
	logs.Debug("Battle report acctionlist is ", this.ReportOne.ActionList)

	this.showReport()

	this.SetBattleUnits()

	logs.Debug("Battle status ", this.Status)

	for _, p := range this.PlayerList { //戰鬥結束之後要重置屬性
		p.IsActive = false
	}

	if this.Status == kUsed {
		this.Round += 1
		this.Point += 1
		this.SendReport(this.ReportOne)
	}

	logs.Debug("站后回合为 ", this.Round)

	this.ReportAll = append(this.ReportAll, this.ReportOne)

}

func (this *BattleRoom) SortUnits() []*GameUnit {
	ul := []*GameUnit{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		ul = append(ul, u)
	}

	sort.Sort(UnitList(ul))

	return ul
}

func (this *BattleRoom) calcWinner() bool {
	return this.Winner != prpc.CT_MAX
}

func (this *BattleRoom) CheckPlayerIsMove() {
	logs.Info("CheckPlayerIsMove 1", this.Units)
	for _, p := range this.PlayerList {
		if p == nil || p.session == nil {
			continue
		}
		p.IsActive = true
		if p.MyUnit.Position != prpc.BP_MAX {
			continue
		}
		pos := this.positionMiddle(p.BattleCamp)
		p.MyUnit.Position = int32(pos)
		this.Units[pos] = p.MyUnit
	}
	logs.Info("CheckPlayerIsMove 1", this.Units)
}

func (this *BattleRoom) SetBattleUnits() {
	for _, unit := range this.Dead {
		this.Units[unit.Position] = nil
	}
}

func (this *BattleRoom) showReport() {
	logs.Debug("第", this.Round+1, "回合")
	for idx, re := range this.ReportOne.ActionList {
		logs.Debug("第", idx+1, "条动作单元")
		logs.Debug("行动的卡牌ID ", re.InstId)
		logs.Debug("身上的buff变更 ", re.BuffList)

		//logs.DebugInfo("\tbuff", re.BuffList)
		//logs.DebugInfo("\tBUFF变更", re.BuffList)
		//logs.DebugInfo("\t本卡是否因为buff死亡", re.BuffList)

		logs.Debug("使用的技能 ", re.SkillId)
		logs.Debug("技能自带的buff ", re.SkillBuff)
		logs.Debug("技能释放的目标信息为")
		logs.Debug("目标链表", re.TargetList)
		for idx1, l := range re.TargetList {
			logs.Debug("\t第", idx1+1, "个目标")
			logs.Debug("\t目标实例ID为", l.InstId)
			logs.Debug("\t目标受击类型", l.ActionType)
			logs.Debug("\t目标伤害", l.ActionParam)
			logs.Debug("\t目标额外信息", l.ActionParamExt)
			logs.Debug("\t目标是否死亡", l.Dead)
			logs.Debug("\t目标中的buff", l.BuffAdd)

		}

	}

}

func (this *BattleRoom) UpdateBuffState(bufflist []int32) {
	logs.Info("UpdateBuffState", bufflist)
	for _, buffid := range bufflist {
		buffCOM := prpc.COM_BattleBuffAction{}
		buffCOM.BuffChange.BuffId = buffid
		buffCOM.BuffChange.Change = 0

		buffCOM.BuffData = 0
		buffCOM.Dead = false

		this.AcctionList.BuffList = append(this.AcctionList.BuffList, buffCOM)
	}

	return
}

func (this *BattleRoom) SendReport(report prpc.COM_BattleReport) {
	b, _ := json.Marshal(report)
	logs.Info(string(b))
	for _, p := range this.PlayerList {
		if p == nil || p.session == nil {
			continue
		}
		p.BattlePoint = this.Point

		p.session.BattleReport(report)
	}
}

////////////////////////////////////////////////////////////////////////
////數據處理
////////////////////////////////////////////////////////////////////////

//取得全部目标
func (this *BattleRoom) SelectAllTarget(camp int) []int64 {
	targets := []int64{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.Camp == camp {
			continue
		}
		targets = append(targets, u.InstId)
	}

	return targets
}

//前排
func (this *BattleRoom) SelectFrontTarget(camp int) []int64 {
	targets := []int64{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp == camp {
			continue
		}
		if !u.isFront() {
			continue
		}
		targets = append(targets, u.InstId)
	}

	if len(targets) == 0 {
		return this.SelectBackTarget(camp)
	}

	return targets
}

//后排
func (this *BattleRoom) SelectBackTarget(camp int) []int64 {
	targets := []int64{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp == camp {
			continue
		}
		if !u.isBack() {
			continue
		}
		targets = append(targets, u.InstId)
	}

	if len(targets) == 0 {
		return this.SelectFrontTarget(camp)
	}

	return targets
}

func GetCampPos(camp int) []int {
	if camp == prpc.CT_RED {
		return []int{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_3, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_6}
	} else if camp == prpc.CT_BLUE {
		return []int{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_3, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
	}
	return []int{}
}

//一竖排敌人,传递第一个选中的敌人,就是先调用getTarget
func (this *BattleRoom) SelectLineTraget(unitid int64) []int64 {
	unit := this.SelectOneUnit(unitid)

	targetList := []int64{unitid}
	var targetPos int32

	pos_list := GetCampPos(unit.Camp)
	if unit.Position > int32(pos_list[2]) {
		targetPos = unit.Position - 3
	} else {
		targetPos = unit.Position + 3
	}

	targetunit := this.Units[targetPos]
	if targetunit == nil {
		return targetList
	}

	if targetunit.IsDead() {
		return targetList
	}

	targetList = append(targetList, targetunit.InstId)

	return targetList
}

//周围敌人,传递第一个选中的敌人,就是先调用getTarget
func (this *BattleRoom) SelectAroundTraget(unitid int64) []int64 {
	unit := this.SelectOneUnit(unitid)

	targetList := []int64{}

	pos_list := GetCampPos(unit.Camp)
	pos_front := unit.Position - 3
	pos_back := unit.Position + 3
	pos_left := unit.Position - 1
	pos_right := unit.Position + 1

	for _, pos := range pos_list {
		if this.Units[pos] == nil {
			continue
		}
		if this.Units[pos].IsDead() {
			continue
		}

		if pos == int(pos_front) {
			targetList = append(targetList, this.Units[pos].InstId)
			continue
		} else if pos == int(pos_back) {
			targetList = append(targetList, this.Units[pos].InstId)
			continue
		} else if pos == int(pos_left) {
			targetList = append(targetList, this.Units[pos].InstId)
			continue
		} else if pos == int(pos_right) {
			targetList = append(targetList, this.Units[pos].InstId)
			continue
		}
	}

	return targetList
}

//取得全部目标
func (this *BattleRoom) SelectMoreTarget(instid int64, num int) []int64 {
	logs.Info("targets start = ", num)
	unit := this.SelectOneUnit(instid)

	targets := []int64{}

	_bool, ok := unit.Special[prpc.BF_FRIENDLOCK]

	if ok && len(_bool) > 0 {

		friend := this.SelectOneFriend(unit.InstId)
		if friend != -1 {
			return []int64{friend}
		}
	}
	idx := 0
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp == unit.Camp {
			continue
		}
		if num > 0 && idx >= num {
			break
		}
		idx += 1
		targets = append(targets, u.InstId)
	}

	logs.Info("targets length1 = ", len(targets))

	if num > 0 && int(num) < int(len(targets)) {
		rand.Seed(time.Now().UnixNano())
		l := make([]int64, len(targets))
		tmp := map[int64]int{}
		var uid int64 = 0
		for len(tmp) < num {
			logs.Info("len(tmp)", len(tmp), num)
			_, ok := tmp[uid]
			for !ok {
				//这里从targets里面随机选择
				idx := rand.Intn(num - 1)
				l = append(l, targets[idx])
				tmp[targets[idx]] = 1
			}
		}
	}
	logs.Info("targets length2 = ", len(targets))
	return targets
}

//取得全部目标
func (this *BattleRoom) SelectOneTarget(instid int64) int64 {
	rand.Seed(time.Now().UnixNano())
	unit := this.SelectOneUnit(instid)
	u_list := []int64{}

	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp == unit.Camp {
			continue
		}

		u_list = append(u_list, u.InstId)
	}
	if len(u_list) == 1 {
		return u_list[0]
	}

	if len(u_list) == 0 {
		if this.Round == 0 {
			for _, p := range this.PlayerList {
				if p.BattleCamp != unit.Camp {
					return p.MyUnit.InstId
				}
			}
		}
		return -1
	}

	index := len(u_list)

	logs.Info("目标索引", index)
	idx := rand.Intn(index)

	return u_list[idx]
}

//取得全部目标
func (this *BattleRoom) SelectOneTargetNoMain(instid int64) int64 {
	rand.Seed(time.Now().UnixNano())
	unit := this.SelectOneUnit(instid)
	u_list := []int64{}

	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp == unit.Camp {
			continue
		}
		if u.IsMain {
			continue
		}

		u_list = append(u_list, u.InstId)
	}

	if len(u_list) == 0 {
		return -1
	}

	if len(u_list) == 1 {
		return u_list[0]
	}

	index := len(u_list)
	idx := rand.Intn(index)

	return u_list[idx]
}

func (this *BattleRoom) SelectOneFriend(instid int64) int64 {
	rand.Seed(time.Now().UnixNano())
	unit := this.SelectOneUnit(instid)
	u_list := []int64{}

	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp != unit.Camp {
			continue
		}
		if u.InstId == instid {
			continue
		}

		u_list = append(u_list, u.InstId)
	}

	logs.Info("友方目标", u_list)
	if len(u_list) == 1 {
		return u_list[0]
	}

	if len(u_list) == 0 {
		if this.Round == 0 {
			for _, p := range this.PlayerList {
				if p.BattleCamp == unit.Camp {
					return p.MyUnit.InstId
				}
			}
		}
		return -1
	}

	index := len(u_list) - 1

	idx := rand.Intn(index)
	//idx := rand.Intn(5)

	logs.Info("一个友方目标", u_list[idx], "index", idx)

	return u_list[idx]
}

//写死
func GetNearPos(pos int32) []int32 {
	if pos < prpc.BP_RED_1 || pos >= prpc.BP_MAX {
		return []int32{}
	}
	switch int(pos) {
	case prpc.BP_RED_1:
		return []int32{prpc.BP_BLUE_3, prpc.BP_BLUE_2, prpc.BP_BLUE_6, prpc.BP_BLUE_5, prpc.BP_BLUE_1, prpc.BP_BLUE_4}
	case prpc.BP_RED_2:
		return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_1, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_6}
	case prpc.BP_RED_3:
		return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_3, prpc.BP_BLUE_6}
	case prpc.BP_RED_4:
		return []int32{prpc.BP_BLUE_3, prpc.BP_BLUE_2, prpc.BP_BLUE_6, prpc.BP_BLUE_5, prpc.BP_BLUE_1, prpc.BP_BLUE_4}
	case prpc.BP_RED_5:
		return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_1, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_6}
	case prpc.BP_RED_6:
		return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_3, prpc.BP_BLUE_6}
	case prpc.BP_BLUE_1:
		return []int32{prpc.BP_RED_3, prpc.BP_RED_2, prpc.BP_RED_6, prpc.BP_RED_5, prpc.BP_RED_1, prpc.BP_RED_4}
	case prpc.BP_BLUE_2:
		return []int32{prpc.BP_RED_2, prpc.BP_RED_1, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_6}
	case prpc.BP_BLUE_3:
		return []int32{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_3, prpc.BP_RED_6}
	case prpc.BP_BLUE_4:
		return []int32{prpc.BP_RED_3, prpc.BP_RED_2, prpc.BP_RED_6, prpc.BP_RED_5, prpc.BP_RED_1, prpc.BP_RED_4}
	case prpc.BP_BLUE_5:
		return []int32{prpc.BP_RED_2, prpc.BP_RED_1, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_6}
	case prpc.BP_BLUE_6:
		return []int32{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_3, prpc.BP_RED_6}
	default:
		return []int32{}
	}
	return []int32{}
}
func GetNearFriend(pos int32) []int32 {
	if pos < prpc.BP_RED_1 || pos >= prpc.BP_MAX {
		return []int32{}
	}
	switch int(pos) {
	case prpc.BP_RED_1:
		return []int32{prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_6}
	case prpc.BP_RED_2:
		return []int32{prpc.BP_RED_1, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_6}
	case prpc.BP_RED_3:
		return []int32{prpc.BP_RED_2, prpc.BP_RED_6, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_1}
	case prpc.BP_RED_4:
		return []int32{prpc.BP_RED_1, prpc.BP_RED_5, prpc.BP_RED_2, prpc.BP_RED_3, prpc.BP_RED_6}
	case prpc.BP_RED_5:
		return []int32{prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_6, prpc.BP_RED_3, prpc.BP_RED_1}
	case prpc.BP_RED_6:
		return []int32{prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_1}

	case prpc.BP_BLUE_1:
		return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
	case prpc.BP_BLUE_2:
		return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_6}
	case prpc.BP_BLUE_3:
		return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_6, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_1}
	case prpc.BP_BLUE_4:
		return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_5, prpc.BP_BLUE_2, prpc.BP_BLUE_3, prpc.BP_BLUE_6}
	case prpc.BP_BLUE_5:
		return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_6, prpc.BP_BLUE_3, prpc.BP_BLUE_1}
	case prpc.BP_BLUE_6:
		return []int32{prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_1}

	default:
		return []int32{}
	}
	return []int32{}
}

func GetAllMyPos(camp int) []int32 {
	if camp == prpc.CT_RED {
		return []int32{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_3, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_6}
	} else if camp == prpc.CT_BLUE {
		return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_3, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
	}
	return []int32{}
}

func GetAllEnemyPos(camp int) []int32 {
	if camp == prpc.CT_BLUE {
		return []int32{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_3, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_6}
	} else if camp == prpc.CT_RED {
		return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_3, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
	}
	return []int32{}
}


func GetEnemyCamp(camp int) int {
	if camp == prpc.CT_BLUE {
		return prpc.CT_RED
	} else if camp == prpc.CT_RED {
		return prpc.CT_BLUE
	}
	return prpc.CT_MAX
}

func (this *BattleRoom) SelectRandomTarget(instid int64, targetnum int32) []int64 {
	unit := this.SelectOneUnit(instid)
	u_list := []int64{}
	r := []int64{}

	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp == unit.Camp {
			continue
		}

		u_list = append(u_list, u.InstId)
	}
	if len(u_list) == 1 {
		for i := 0; i < int(targetnum); i++ {
			r = append(r, u_list[0])
		}
		return r
	}

	if len(u_list) == 0 {
		if this.Round == 0 {
			for _, p := range this.PlayerList {
				if p.BattleCamp != unit.Camp {
					for i := 0; i < int(targetnum); i++ {
						r = append(r, p.MyUnit.InstId)
					}
					return r
				}
			}
		}
		return r
	}

	index := len(u_list)

	logs.Info("目标索引", index)
	for i := 0; i < int(targetnum); i++ {
		idx := rand.Intn(index)
		r = append(r, u_list[idx])
	}

	return r
}

func (this *BattleRoom) SelectNearTarget(instid int64) int64 {
	unit := this.SelectOneUnit(instid)

	near_pos := GetNearPos(unit.Position)

	if near_pos == nil || len(near_pos) == 0 {
		return this.SelectOneTarget(unit.InstId)
	}

	buff_lis, ok := unit.Special[prpc.BF_FRIENDLOCK] //检测有无选择锁定随机己方的目标buff
	logs.Info("目标 111", unit.InstId, "己方随机目标buff ,", unit.Special[prpc.BF_FRIENDLOCK])
	if ok {
		if len(buff_lis) > 0 {
			logs.Info("目标 222", unit.InstId, "己方随机目标buff ,", unit.Special[prpc.BF_FRIENDLOCK])
			return this.SelectOneFriend(unit.InstId)
		}
	}
	for _, pos := range near_pos {
		if this.Units[pos] == nil {
			continue
		}
		if this.Units[pos].IsDead() {
			continue
		}

		return this.Units[pos].InstId
	}

	return this.SelectOneTarget(unit.InstId)
}

func (this *BattleRoom) SelectNearFriend(instid int64) int64 {
	unit := this.SelectOneUnit(instid)

	near_pos := GetNearFriend(unit.Position)

	if near_pos == nil || len(near_pos) == 0 {
		return this.SelectOneFriend(unit.InstId)
	}
	for _, pos := range near_pos {
		if this.Units[pos] == nil {
			continue
		}
		if this.Units[pos].IsDead() {
			continue
		}
		return this.Units[pos].InstId
	}

	return this.SelectOneFriend(unit.InstId)
}

//取得全部友方目标
func (this *BattleRoom) SelectMoreFriend(instid int64, num int) []int64 {
	unit := this.SelectOneUnit(instid)

	friends := []int64{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.IsDead() {
			continue
		}
		if u.Camp != unit.Camp {
			continue
		}
		friends = append(friends, u.InstId)
	}

	//if num > 0 && int(num) < int(len(targets)){
	//	rand.Seed(time.Now().UnixNano())
	//	l := make([]int64, len(targets))
	//	tmp := map[int64]int{}
	//	var uid int64 = 0
	//	for len(tmp) <= num {
	//		_, ok := tmp[uid]
	//		for !ok {
	//			//这里从targets里面随机选择
	//			idx := rand.Intn(num - 1)
	//			l = append(l, targets[idx])
	//			tmp[targets[idx]] = 1
	//		}
	//	}
	//}
	logs.Info("friends length = ", len(friends))
	return friends
}

func (this *BattleRoom) SelectThrowCard(instid int64) (int64, int32, int32) {
	unit := this.SelectOneUnit(instid)

	canThrow := []int64{}
	units := []*GameUnit{}

	if this.Type == prpc.BT_PVP {
		for _, p := range this.PlayerList {
			if p.MyUnit.InstId == unit.Owner.MyUnit.InstId {
				continue
			}
			for _, g := range p.UnitGroup {
				if g.GroupId != p.BattleUnitGroup {
					continue
				}
				for _, ud := range g.UnitList {
					u := p.GetUnit(ud)
					if u.IsDead() {
						continue
					}
					if u.OutBattle {
						continue
					}
					canThrow = append(canThrow, ud)
					units = append(units, u)
				}
			}
		}
	} else {
		for _, m := range this.Monster.BattleUnitList {
			if m.IsDead() {
				continue
			}
			if m.OutBattle {
				continue
			}
			canThrow = append(canThrow, m.InstId)
			units = append(units, m)
		}
	}

	logs.Info("can throw all cards", canThrow)

	if len(canThrow) == 0 {
		return 0, 0, 0
	}

	var del_card int64
	var idx int

	if len(canThrow) == 1 {
		idx = 0
	} else {
		index := len(canThrow) - 1

		idx = rand.Intn(index)
	}

	del_card = canThrow[idx]
	del_unit := units[idx]

	logs.Info("throw one card", del_card)

	return del_card, del_unit.UnitId, del_unit.IProperties[prpc.IPT_PROMOTE]
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

	for _, p := range this.PlayerList {
		if instid == p.MyUnit.InstId {
			return p.MyUnit
		}
		for _, u := range p.UnitList {
			if u == nil {
				continue
			}
			if u.InstId == instid {
				return u
			}
		}
	}

	return nil
}

func (this *BattleRoom) selectMainUnit(instid int64, MyCamp bool) int64 {
	unit := this.SelectOneUnit(instid)

	if unit == nil {
		return 0
	}

	if this.Type == prpc.BT_PVP {
		for _, p := range this.PlayerList {
			if MyCamp && p.MyUnit.Camp == unit.Camp {
				return p.MyUnit.InstId
			}

			if !MyCamp && p.MyUnit.Camp != unit.Camp {
				return p.MyUnit.InstId
			}
		}
	} else {
		if unit.Camp == this.PlayerList[0].MyUnit.Camp {
			if MyCamp {
				return this.PlayerList[0].MyUnit.InstId
			} else {
				return this.Monster.MainUnit.InstId
			}
		} else {
			if MyCamp {
				return this.Monster.MainUnit.InstId
			} else {
				return this.PlayerList[0].MyUnit.InstId
			}
		}
	}

	logs.Info("selectMainUnit end ", 0)

	return 0
}

////////////////////////////////////////////////////////////////////////
////player测试
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) PlayerMove() {
	logs.Info("PlayerMove 1", this.Units)

	p := this.PlayerList[0]
	if this.Round == 0 {
		pos := this.positionMiddle(p.BattleCamp)
		this.Units[pos] = p.MyUnit
		p.MyUnit.Position = int32(pos)
	} else {
		//pos := this.monsterPos(p.BattleCamp)
		//this.Units[pos] = this.Monster.BattleUnitList[0]
		//this.Monster.BattleUnitList[0].Position = int32(pos)
		//this.Monster.BattleUnitList = this.Monster.BattleUnitList[1:]
	}

	logs.Info("PlayerMove 2", this.Units)

	return
}

////////////////////////////////////////////////////////////////////////
////PVE行为
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) MonsterMove() {
	logs.Info("MonsterMove 1", this.Units)
	if this.Round == 0 {
		pos := this.positionMiddle(this.Monster.BattleCamp)
		this.Units[pos] = this.Monster.MainUnit
		this.Monster.MainUnit.Position = int32(pos)
	} else {
		pos := this.monsterPos(this.Monster.BattleCamp)
		if pos == prpc.BP_MAX {
			return
		}
		if len(this.Monster.BattleUnitList) > 0 {
			for index, m := range this.Monster.BattleUnitList {
				if m.OutBattle {
					logs.Info("outbattle", m.InstId)
					continue
				}
				logs.Info("outbattle1111111 ", m.InstId)
				this.Units[pos] = m
				m.Position = int32(pos)
				this.Monster.BattleUnitList = this.Monster.BattleUnitList[index+1:]
				break
			}
		}
	}

	this.Monster.MainUnit.ChoiceSKill = this.Monster.MainUnit.SelectSkill(this.Round).SkillID

	logs.Info("MonsterMove 2", this.Units)
	logs.Info("MonsterMove 2.1 ", this.Monster.MainUnit.ChoiceSKill)

	return
}

func (this *BattleRoom) positionMiddle(camp int) int {
	if camp == prpc.CT_RED {
		return prpc.BP_RED_5
	} else {
		return prpc.BP_BLUE_5
	}
}

func (this *BattleRoom) monsterPos(camp int) int {
	if camp == prpc.CT_RED {
		for i := prpc.BP_RED_1; i <= prpc.BP_RED_6; i++ {
			if this.Units[i] != nil {
				continue
			}
			return i
		}
	} else {
		for i := prpc.BP_BLUE_1; i <= prpc.BP_BLUE_6; i++ {
			if this.Units[i] != nil {
				continue
			}
			return i
		}
	}

	return prpc.BP_MAX
}

////////////////////////////////////////////////////////////////////////
////属性操控
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) GetUnitProperty(instid int64, property string) int {
	p_d := prpc.ToId_CPropertyType(property)

	unit := this.SelectOneUnit(instid)

	pro := unit.CProperties[p_d]

	return int(pro)
}
func (this *BattleRoom) ChangeCptProperty(instid int64, data int32, property string) { //CPT
	logs.Info("增加攻击力,目标为 ", instid, data)
	p_d := prpc.ToId_CPropertyType(property)

	unit := this.SelectOneUnit(instid)
	if unit == nil {
		for _, p := range this.PlayerList {
			if instid == p.MyUnit.InstId {
				unit = p.MyUnit
				break
			}
			for _, u := range p.UnitList {
				if u.InstId == instid {
					unit = u
					break
				}
			}
		}
	}
	logs.Info("属性修改前", unit.CProperties[p_d], data)

	unit.CProperties[p_d] = unit.CProperties[p_d] + float32(data)
	logs.Info("属性修改后", unit.CProperties[p_d], data)

	return
}

func (this *BattleRoom) ChangeIptProperty(instid int64, data int32, property string) { //IPT
	logs.Info("增加攻击力,目标为 ", instid, data)
	p_d := prpc.ToId_IPropertyType(property)

	unit := this.SelectOneUnit(instid)
	if unit == nil {
		for _, p := range this.PlayerList {
			if instid == p.MyUnit.InstId {
				unit = p.MyUnit
				break
			}
			for _, u := range p.UnitList {
				if u.InstId == instid {
					unit = u
					break
				}
			}
		}
	}
	logs.Info("属性修改前", unit.CProperties[p_d], data)

	unit.CProperties[p_d] = unit.CProperties[p_d] + float32(data)
	logs.Info("属性修改后", unit.CProperties[p_d], data)

	return
}

func (this *BattleRoom) MintsHp(casterid int64, target int64, damage int32, crit int32) {

	unit := this.SelectOneUnit(target)

	if unit.IsDead() {
		return
	}

	for _, debuff := range unit.Debuff {
		logs.Info("debuff", debuff)
	}

	for _, debuff := range unit.Buff {
		logs.Info("debuff", debuff)
	}

	isover := unit.CheckSpec("BF_UNDAMAGE", this.Round)
	logs.Info("BF_UNDAMAGE out ", isover, unit.InstName, unit.Special[prpc.BF_UNDAMAGE])
	if isover {
		logs.Info("BF_UNDAMAGE", damage, unit.InstName)
		damage = 0
	}

	//////////////////////////////////////////////////////////////////////////
	//if unit.VirtualHp >= damage {			//计算护盾减伤之后的伤害
	//	damage = 0
	//	unit.VirtualHp = unit.VirtualHp - damage
	//} else {
	//	damage = damage - unit.VirtualHp
	//}
	//////////////////////////////////////////////////////////////////////////

	//////////////////////////////////////////////////////////////////////////
	//per := unit.ClacSheldPer(this.Round)		//百分比减伤 结果是最终减伤多少
	//if per > 0 {
	//	persent := 1.0 - per					//实际受到的伤害百分比
	//	damage = int32(float32(damage) * persent)
	//}
	//////////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////////

	caster := this.SelectOneUnit(casterid)

	exercise := GetRaceRecordById(caster.Race)

	if exercise != nil {
		if unit.Race == exercise.Exercise {
			logs.Info("unit.Race == exercise.Exercise 1", damage)
			damage_q := damage * exercise.Quotiety / 100
			damage += damage_q
			logs.Info("unit.Race == exercise.Exercise 2", damage)
		}
	}

	if float32(damage) >= unit.CProperties[prpc.CPT_CHP] { //检测免死
		bf, ok := unit.Special[prpc.BF_UNDEAD]
		true_list := []int32{}
		if ok {
			if len(bf) > 0 {
				for _, bid := range bf {
					buff := unit.SelectBuff(bid)
					if buff == nil {
						continue
					}
					if buff.IsOver(this.Round) {
						continue
					}
					true_list = append(true_list, bid)
				}
			}
		}

		if len(true_list) > 0 {
			logs.Info("免死触发")
			damage = 0
			if len(true_list) == 1 { // 只有一个就删除掉这个效果
				delete(unit.Special, prpc.BF_UNDEAD)
			} else {
				unit.Special[prpc.BF_UNDEAD] = true_list[:1]
			}
			buff := unit.SelectBuff(true_list[0])
			buff.Over = true
		}
	}
	////////////////////////////////////////////////////////////////////////

	unit.CProperties[prpc.CPT_CHP] = unit.CProperties[prpc.CPT_CHP] - float32(damage)

	if this.NewAction == false {
		this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
		this.TargetCOM = prpc.COM_BattleActionTarget{}
		this.TargetCOM.TransPostion = prpc.BP_MAX
	}

	if crit == 0 {
		crit = prpc.BE_MAX
	} else {
		crit = prpc.BE_Crit
	}

	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = -damage
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(int(crit))
	this.TargetCOM.Dead = unit.IsDead()
	this.TargetCOM.BuffAdd = []prpc.COM_BattleBuff{}
	this.NewAction = false

	//logs.Info("MintsHp", target, damage, t)

	logs.Info("攻擊  catserid ", casterid, this.AcctionList.TargetList)
	logs.Info("MintsHp 1  ", this.TargetCOM)

	if unit.IsDead() {
		unit.OutBattle = true
		unit.PopAllBuffByDead(this)
		this.isDeadOwner(casterid, target)
		this.Dead = append(this.Dead, unit)
		if unit.Owner != nil {
			unit.Owner.MyDeathNum += 1
		}
		if caster.Owner != nil {
			caster.Owner.KillUnits = append(caster.Owner.KillUnits, unit.UnitId)
		}
	}

}

func (this *BattleRoom) ThrowCard(target int64, throwcard int64, entity int32, level int32) {

	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = -0
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(prpc.BE_MAX)
	this.TargetCOM.Dead = false
	this.TargetCOM.BuffAdd = []prpc.COM_BattleBuff{}
	this.TargetCOM.ThrowCard = prpc.COM_ThrowCard{InstId: throwcard, EntityId: entity, Level: level}
	this.NewAction = false

}

func (this *BattleRoom) Throw(main int64, throwcard int64) {

	if this.Type == prpc.BT_PVP {
		for _, p := range this.PlayerList {
			if p.MyUnit.InstId != main {
				continue
			}
			unit := p.GetUnit(throwcard)
			unit.OutBattle = true
		}
	} else {
		for _, m := range this.Monster.BattleUnitList {
			if m.InstId != throwcard {
				continue
			}
			m.OutBattle = true
		}
	}

}

func (this *BattleRoom) AddHp(target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead() {
		return
	}

	if this.NewAction == false {
		this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
		this.TargetCOM = prpc.COM_BattleActionTarget{}
		this.TargetCOM.TransPostion = prpc.BP_MAX
	}

	if crit == 0 {
		crit = prpc.BE_MAX
	} else {
		crit = prpc.BE_Crit
	}

	unit.CProperties[prpc.CPT_CHP] = unit.CProperties[prpc.CPT_CHP] + float32(damage)

	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = damage
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(int(crit))
	this.TargetCOM.Dead = unit.IsDead()
	this.TargetCOM.BuffAdd = []prpc.COM_BattleBuff{}
	this.NewAction = false

	logs.Info("加血  catserid ", target, this.AcctionList.TargetList)
	logs.Info("AddHp 1  ", this.TargetCOM)

	//this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

//降低法术

func (this *BattleRoom) ReduceSpell(target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead() {
		return
	}

	unit.CProperties[prpc.CPT_MAGIC_ATK] = unit.CProperties[prpc.CPT_MAGIC_ATK] - float32(damage)

	logs.Info("实例ID为", target, "的卡牌在第", this.Round+1, "法术降到", unit.CProperties[prpc.CPT_MAGIC_ATK], "降低了", damage)

	this.TargetCOM = prpc.COM_BattleActionTarget{}
	this.TargetCOM.TransPostion = prpc.BP_MAX
	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = damage
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(int(crit))

	this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

//增加法术

func (this *BattleRoom) IncreaseSpell(target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead() {
		return
	}

	unit.CProperties[prpc.CPT_MAGIC_ATK] = unit.CProperties[prpc.CPT_MAGIC_ATK] + float32(damage)

	logs.Info("实例ID为", target, "的卡牌在第", this.Round+1, "法术增加到", unit.CProperties[prpc.CPT_MAGIC_ATK], "增加了", damage)

	this.TargetCOM = prpc.COM_BattleActionTarget{}
	this.TargetCOM.TransPostion = prpc.BP_MAX
	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = damage
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(int(crit))

	this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

func (this *BattleRoom) AddSkillBuff(casterid int64, target int64, buffid int32, data int32) {
	// 上buff

	buffCOM := prpc.COM_BattleBuff{}
	buffCOM.BuffId = buffid
	buffCOM.Change = 1

	unit := this.SelectOneUnit(target)

	buff := NewBuff(unit, casterid, buffid, data, this.Round)
	buff.AddProperty()

	unit.Allbuff = append(unit.Allbuff, buff)

	this.AcctionList.SkillBuff = append(this.AcctionList.SkillBuff, buffCOM)

}
func (this *BattleRoom) AddBuff(casterid int64, target int64, buffid int32, data int32) {
	// 上buff

	buffCOM := prpc.COM_BattleBuff{}
	buffCOM.BuffId = buffid
	buffCOM.Change = 1

	unit := this.SelectOneUnit(target)

	buff := NewBuff(unit, casterid, buffid, data, this.Round)
	unit.Allbuff = append(unit.Allbuff, buff)

	buff.AddProperty()

	logs.Debug("bufflen front", this.TargetCOM)
	this.TargetCOM.BuffAdd = append(this.TargetCOM.BuffAdd, buffCOM)
	logs.Debug("实例ID为", target, "的卡牌在第", this.Round+1, "回合获得了id为", buff.InstId, "的buff, buff表中的ID为", buffid, "目前该卡牌一共有", len(unit.Allbuff), "个buff, ", buff.Round)
	logs.Debug("bufflen back", this.TargetCOM)

}

func (this *BattleRoom) DeleteBuff(target int64, buffinstid int32, data int32) {
	// 去buff

	buffCOM := prpc.COM_BattleBuff{}

	unit := this.SelectOneUnit(target)

	buff := unit.SelectBuff(buffinstid)
	buff.DeleteProperty()

	buffCOM.BuffId = buff.BuffId
	buffCOM.Change = 0

	logs.Info("bufflen front", this.TargetCOM)
	this.TargetCOM.BuffAdd = append(this.TargetCOM.BuffAdd, buffCOM)
	logs.Info("实例ID为", target, "的卡牌在第", this.Round+1, "回合失去了id为", buff.InstId, "的buff, buff表中的ID为", buff.BuffId, "目前该卡牌一共有", len(unit.Allbuff), "个buff, ", buff.Round)
	logs.Info("bufflen back", this.TargetCOM)

}

func (this *BattleRoom) BuffMintsHp(casterid int64, target int64, buffid int32, data int32, over bool) {
	logs.Info("BuffMintsHp", " buff 给id为", target, "的卡牌造成了", data, "点伤害, over", over)
	unit := this.SelectOneUnit(target)

	if unit.CheckSpec("BF_UNDAMAGE", this.Round) {
		data = 0
	}

	unit.CProperties[prpc.CPT_CHP] = unit.CProperties[prpc.CPT_CHP] - float32(data)

	buffCOM := prpc.COM_BattleBuffAction{}
	buffCOM.BuffChange.BuffId = buffid
	buffCOM.BuffChange.Change = 2
	if over {
		buffCOM.BuffChange.Change = 0
	}

	buffCOM.BuffData = -data
	buffCOM.Dead = unit.IsDead()

	this.AcctionList.BuffList = append(this.AcctionList.BuffList, buffCOM)

	if unit.IsDead() {
		unit.OutBattle = true
		unit.PopAllBuffByDead(this)
		this.isDeadOwner(casterid, target)
		this.Dead = append(this.Dead, unit)
		if unit.Owner != nil {
			unit.Owner.MyDeathNum += 1
		}
		caster := this.SelectOneUnit(casterid)
		if caster.Owner != nil {
			caster.Owner.KillUnits = append(caster.Owner.KillUnits, unit.UnitId)
		}
	}
}

func (this *BattleRoom) BuffAddHp(target int64, buffid int32, data int32, over bool) {
	logs.Info("BuffMintsHp", " buff 给id为", target, "的卡牌增加了", data, "点血量, over", over)
	unit := this.SelectOneUnit(target)

	unit.CProperties[prpc.CPT_CHP] = unit.CProperties[prpc.CPT_CHP] + float32(data)

	buffCOM := prpc.COM_BattleBuffAction{}
	buffCOM.BuffChange.BuffId = buffid
	buffCOM.BuffChange.Change = 2
	if over {
		buffCOM.BuffChange.Change = 0
	}

	buffCOM.BuffData = data
	buffCOM.Dead = unit.IsDead()

	this.AcctionList.BuffList = append(this.AcctionList.BuffList, buffCOM)

}

func (this *BattleRoom) InToBattleOnFighting(PlayerInstId int64) {
	p := FindPlayerByInstId(PlayerInstId)
	group := p.GetUnitGroupById(p.BattleUnitGroup)
	space := []*GameUnit{}
	//logs.Debug("group",group)
	for _, uid := range group.UnitList {
		u := p.GetUnit(uid)
		if u.Position != 12 {
			//logs.Debug(uid, " 11111 ", u.InstName, u.Position)
			continue
		}

		space = append(space, u)
	}
	//logs.Debug("space"," ",len(space))
	if len(space) == 0{
		return
	}

	poss := []int32{}
	mypos := GetAllMyPos(p.BattleCamp)
	//logs.Debug("mypos"," ",len(mypos))
	for _, pos := range mypos {
		//logs.Debug("循环")
		if this.Units[pos] != nil && !this.Units[pos].IsDead(){
			//logs.Debug("1111111")
			continue
		}
		//logs.Debug("9999999")
		poss = append(poss, pos)
	}
	//logs.Debug("poss",len(poss))
	if len(poss) == 0{
		return
	}

	this.Units[poss[0]] = space[0]
	this.Units[poss[0]].Position = poss[0]

	info := prpc.COM_ChangeUnit{}
	info.Status = true
	info.Unit = space[0].GetBattleUnitCOM()

	//logs.Debug("人物链表",this.AcctionList.UnitList)

	this.AcctionList.UnitList = append(this.AcctionList.UnitList, info)

}

func (this *BattleRoom) TransPosOnFighting(instid int64) {
	unit := this.SelectOneUnit(instid)

	if unit.IsMain {
		return
	}

	if unit == nil {
		return
	}

	space_pos := []int32{}

	for _, pos := range GetAllEnemyPos(unit.Camp) {
		if this.Units[pos] != nil && !this.Units[pos].IsDead(){
			continue
		}
		space_pos = append(space_pos, pos)
	}

	if len(space_pos) == 0 {
		return
	}

	transCamp := GetEnemyCamp(unit.Camp)

	if transCamp == prpc.CT_MAX {
		return
	}

	unit.Camp = transCamp
	this.Units[unit.Position] = nil

	this.Units[space_pos[0]] = unit
	unit.Position = space_pos[0]

	this.TargetCOM.TransPostion = space_pos[0]

}

func (this *BattleRoom) TargetOn() {
	this.TargetCOM = prpc.COM_BattleActionTarget{}
	this.TargetCOM.TransPostion = prpc.BP_MAX
	this.NewAction = true
}

func (this *BattleRoom) TargetOver() {
	this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

func (this *BattleRoom) isDeadOwner(casterid int64, target int64) {
	unit := this.SelectOneUnit(target)

	if !unit.IsMain {
		return
	}

	if !unit.IsDead() {
		return
	}

	caster := this.SelectOneUnit(casterid)

	if casterid == target {
		if caster.Camp == prpc.CT_RED {
			this.Winner = prpc.CT_BLUE
		} else if caster.Camp == prpc.CT_BLUE {
			this.Winner = prpc.CT_RED
		} else {
			this.Winner = prpc.CT_MAX
		}
	} else {
		this.Winner = caster.Camp
	}

}

////////////////////////////////////////////////////////////////////////
////独立接口
////////////////////////////////////////////////////////////////////////

func IsCrit(skillid int32) int {
	skill := GetSkillRecordById(skillid)

	if skill == nil {
		return 0
	}

	per := rand.Intn(100)
	//logs.Info("IsCrit", skill.Crit)

	if per <= int(skill.Crit) {
		return 1
	}

	return 0
}

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SetupPosition(p *GamePlayer, posList []prpc.COM_BattlePosition, skillid int32) {

	logs.Info("SetupPosition.start", posList, p.BattleCamp)
	//if this.Round == 0 { //第一回合 必须设置主角卡
	//	for _, pos := range posList {
	//		//logs.Info("SetupPosition, set ", pos.InstId, p.MyUnit.InstId)
	//		if pos.InstId == p.MyUnit.InstId {
	//			goto setup_check_success
	//		}
	//		logs.Info("SetupPosition.error no main")
	//		return //没有主角卡
	//	}
	//}
	//setup_check_success:
	var needPoint int32
	for i := 0; i < len(posList)-1; i++ {
		for j := i + 1; j < len(posList); j++ {
			if posList[i].InstId == posList[j].InstId {
				logs.Info("SetupPosition.error card same")
				return //有重复卡牌设置
			}
			if posList[i].Position == posList[j].Position {
				logs.Info("SetupPosition.error pos same")
				return //有重复位置设置
			}
			unit := p.GetUnit(posList[i].InstId)

			if unit.IsDead() {
				logs.Info("SetupPosition.error card is dead")
				return
			}
			needPoint += unit.Cost
		}
	}

	needPoint += 0 // 主角技能应该有消耗

	if needPoint > p.BattlePoint {
		logs.Info("SetupPosition.error point less", needPoint, p.BattlePoint)
		return //能量点不足
	}

	for _, pos := range posList {
		for _, u := range this.Units {
			if u == nil {
				continue
			}
			if u.InstId == pos.InstId {
				logs.Info("SetupPosition.error pos same 2")
				return //已经上场
			}
		}

		if this.Units[pos.Position] != nil {
			logs.Info("SetupPosition.error pos same 3")
			return //这个位置上有人
		}
	}

	//处理数据
	for _, pos := range posList {
		this.Units[pos.Position] = p.GetUnit(pos.InstId)
		this.Units[pos.Position].Position = pos.Position
	}
	p.IsActive = true
	p.BattlePoint -= needPoint
	p.MyUnit.ChoiceSKill = skillid

	if this.Type == prpc.BT_PVE || this.Type == prpc.BT_PVR {
		this.BattleUpdate()
	}
	//logs.Info("SetupPosition", this.Units, p.BattleCamp, p.IsActive)
}

////////////////////////////////////////////////////////////////////////
////属性检测
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) HasBuff(target int64) bool {
	unit := this.SelectOneUnit(target)

	for _, buff := range unit.Allbuff {
		if buff.BuffType == kTypeBuff {
			return true
		}
	}
	return false
}

func (this *BattleRoom) HasDebuff(target int64) bool {
	unit := this.SelectOneUnit(target)

	for _, buff := range unit.Allbuff {
		if buff.BuffType == kTypeDebuff {
			return true
		}
	}
	return false
}
