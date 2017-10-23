package game

import (
	"logic/prpc"
	"sync"
	"sync/atomic"
	"time"
	"math/rand"
	"sort"
	"encoding/json"
	"logic/std"
	//"fmt"

)

type UnitList []*GameUnit

func (a UnitList) Len() int {    // 重写 Len() 方法
	return len(a)
}
func (a UnitList) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a UnitList) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j].CProperties[prpc.CPT_AGILE] < a[i].CProperties[prpc.CPT_AGILE]
}

const (
	kIdle = 0 			// 无效状态
	kUsed = 1 			// 使用状态

	kTurn    = 1 		//回合數
	kMaxUnit = 6 		//雙方最多上陣卡牌
	kMaxMove = 2 		//行动结束

	kTimeSleep = 3   	//檢測間隔
	kTimeMax   = 600	//戰鬥持續時間
)

var roomInstId int64 = 1

type Monster struct {
	sync.Locker
	MainUnit         	*GameUnit   //自己的卡片
	BattleUnitList      []*GameUnit //拥有的卡片

	//战斗相关辅助信息
	BattleCamp 			int   //阵营 //prpc.CompType
}

type BattleRoom struct {
	sync.Mutex
	Type 	   	int32		 				//战斗类型 1是pvp 2是pve
	InstId     	int64         				//房间ID
	BattleID    int32         				//戰鬥ID
	Status     	int32         				//战斗房间状态
	Round      	int32         				//回合计数
	Point      	int32         				//本场战斗的能量点
	Units      	[]*GameUnit   				//当前战斗中牌 数组索引跟下面玩家对应
	Dead      	[]*GameUnit   				//本回合死亡的人数
	PlayerList 	[]*GamePlayer 				//房间中玩家信息
	Monster    	*Monster
	Turn       	int32
	Winner     	int 						//获胜者
	ReportAll  	[]prpc.COM_BattleReport		//整场战斗的所有战报
	ReportOne  	prpc.COM_BattleReport		//每回合的战报
	AcctionList prpc.COM_BattleAction		//行动单元
	TargetCOM	prpc.COM_BattleActionTarget	//行动单元中的每个子元素
	NewAction	bool						//是否行动过
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

	std.LogInfo("CreatePvE", &room)
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

func CreateMonster(battleid int32, roomid int64) *Monster{
	t := GetBattleRecordById(battleid)

	m := Monster{}

	m.MainUnit = CreateUnitFromTable(t.MainId)
	std.LogInfo("adasdas", battleid, t.MainId, m.MainUnit)
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
	std.LogInfo("CreatePvP", &room)

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

func (this *BattleRoom) BattleStart() {

	ul := []prpc.COM_BattleUnit{}

	unitllist := this.SortUnits()

	for _, unit := range unitllist {
		if unit == nil {
			continue
		}
		std.LogInfo("卡牌敏捷: 1 ", unit.GetCProperty(prpc.CPT_AGILE))
		ul = append(ul, unit.GetBattleUnitCOM())
	}

	for _, p := range this.PlayerList {

		if p == nil || p.session == nil {
			continue
		}
		targetList := this.findCardsByTarget(p.BattleCamp)
		std.LogInfo("JoinBattleOk, p.id", p.MyUnit.InstId, " and batlecamp is ", int32(p.BattleCamp), "targetList is ", targetList)
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
	for index, p := range this.PlayerList{
		if p.MyUnit.InstId == player.MyUnit.InstId {
			del_index = index
		}
	}
	if del_index != -1 {
		this.PlayerList = append(this.PlayerList[:del_index], this.PlayerList[del_index+1:]...)
	}

	if len(this.PlayerList) == 0{
		this.BattleRoomOver(prpc.CT_MAX)
		this.Status = kIdle
	}

}

func (this *BattleRoom) SendReportFirst() {

}

func (this *BattleRoom) findCardsByTarget(camp int) []int32 {
	li := []int32{}
	for _, p := range this.PlayerList {
		if p == nil{
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

	//defer func() {
	//
	//	if r := recover(); r != nil {
	//		std.LogError("main panic %s",fmt.Sprint(r))
	//	}
	//
	//}()

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

		std.LogInfo("BattleUpdate, roomId is ", this.InstId, "index is", checkindex)
		this.Update()
		now_start = time.Now().Unix()
		checkindex += 1
		if this.Type == prpc.BT_PVE{	//pve只执行一次就跳出
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

		result.Win	= win
		result.BattleRound = this.Round
		result.KillMonsters = p.KillUnits
		result.MySelfDeathNum = p.MyDeathNum

		if this.Type == prpc.BT_PVE {
			dropId := p.CalcSmallChapterStar(result)
			if dropId != 0 {
				drop := GetDropById(dropId)
				if drop==nil {
					std.LogInfo("PVE Can Not Find Drop By DropId=",dropId)
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
					for _,item := range drop.Items{
						p.AddBagItemByItemId(item.ItemId,item.ItemNum)
						std.LogInfo("PVE GiveDrop AddItem ItemId=",item.ItemId,"ItemNum=",item.ItemNum)
						itemInst := prpc.COM_ItemInst{}
						itemInst.ItemId = item.ItemId
						itemInst.Stack = item.ItemNum
						result.BattleItems = append(result.BattleItems,itemInst)
					}
				}
				if drop.Hero != 0 {
					if p.HasUnitByTableId(drop.Hero) {
						//有这个卡就不给了
						std.LogInfo("PlayerName=",p.MyUnit.InstName,"GiveDrop AddUnit Have not to UnitId=",drop.Hero)
					}else {
						unit := p.NewGameUnit(drop.Hero)
						if unit!=nil {
							std.LogInfo("PlayerName=",p.MyUnit.InstName,"GiveDrop AddUnit OK UnitId=",drop.Hero)
							temp := unit.GetUnitCOM()
							if p.session != nil {
								p.session.AddNewUnit(temp)
							}
						}
					}
				}
			}
			std.LogInfo("Battle Over PVE DropId=",dropId)
		}
		if this.Type == prpc.BT_PVP {
			for _,once := range this.PlayerList{
				if once == nil {
					continue
				}
				if once.MyUnit.InstId == p.MyUnit.InstId {
					continue
				}
				if p.BattleCamp == once.BattleCamp {
					continue
				}
				dropId := CaleTianTiVal(p,once,camp)

				if dropId != 0 {
					drop := GetDropById(dropId)
					if drop==nil {
						std.LogInfo("PVP Can Not Find Drop By DropId=",dropId)
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
						for _,item := range drop.Items{
							p.AddBagItemByItemId(item.ItemId,item.ItemNum)
							std.LogInfo("PVP GiveDrop AddItem ItemId=",item.ItemId,"ItemNum=",item.ItemNum)
							itemInst := prpc.COM_ItemInst{}
							itemInst.ItemId = item.ItemId
							itemInst.Stack = item.ItemNum
							result.BattleItems = append(result.BattleItems,itemInst)
						}
					}
				}
			}
		}
		p.BattleId = 0
		p.ClearAllBuff()

		if p.session != nil {
			p.session.BattleExit(result)
		}
		std.LogInfo("BattleRoomOver, result is ", result, "player is ", p.MyUnit.InstId, "p.battlecampis ", p.BattleCamp, "wincampis ", camp, "winis", win)
		p.BattleCamp = prpc.CT_MAX
	}

	std.LogInfo("BattleRoomOver, winner is ", camp)
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

	if this.Type == 2 {
		this.MonsterMove()
		//this.PlayerMove()
	}

	//this.CheckPlayerIsMove()

	//顺序排序

	std.LogInfo("站前回合为 ", this.Round)

	this.ReportOne = prpc.COM_BattleReport{}
	this.Dead = []*GameUnit{}
	this.ReportOne.BattleID = this.BattleID

	unitllist := this.SortUnits()

	for _, unit := range unitllist {
		if unit == nil {
			continue
		}
		std.LogInfo("卡牌敏捷: 1 ", unit.GetCProperty(prpc.CPT_AGILE))
		this.ReportOne.UnitList = append(this.ReportOne.UnitList, unit.GetBattleUnitCOM())
	}

	//if this.Round > 0 {
		for _, u := range unitllist {
			if u == nil {
				continue
			}
			std.LogInfo("卡牌敏捷 2 : ", u.GetCProperty(prpc.CPT_AGILE))
			//std.LogInfo("report.UnitList, append", u)
			//this.ReportOne.UnitList = append(this.ReportOne.UnitList, u.GetBattleUnitCOM())

			if u.IsDead() { // 非主角死亡跳過
				continue
			}

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
			u.ChoiceSKill = 0

			//this.TargetOver()

			this.ReportOne.ActionList = append(this.ReportOne.ActionList, this.AcctionList)

			//std.LogInfo("BattleAcction", u.InstId, "acction", this.AcctionList)
			//std.LogInfo("BattleAcction", u.InstId, "ReportOne", this.ReportOne)
			if this.calcWinner() == true {
				for _, a := range this.AcctionList.TargetList {
					unit := this.SelectOneUnit(a.InstId)
					if unit == nil {
						continue
					}
					//std.LogInfo("append", unit)
					//this.ReportOne.UnitList = append(this.ReportOne.UnitList, unit.GetBattleUnitCOM())
				}
				//std.LogInfo("this.Winner", this.Winner)

				this.Round += 1
				this.SendReport(this.ReportOne)
				this.BattleRoomOver(this.Winner)
				this.Status = kIdle
				break

			}
		}
	//}
	std.LogInfo("Battle report battleid is ", this.ReportOne.BattleID)
	std.LogInfo("Battle report unitlist is ", this.ReportOne.UnitList)
	std.LogInfo("Battle report acctionlist is ", this.ReportOne.ActionList)

	this.showReport()

	this.SetBattleUnits()

	std.LogInfo("Battle status ", this.Status)

	for _, p := range this.PlayerList { //戰鬥結束之後要重置屬性
		p.IsActive = false
	}

	if this.Status == kUsed {
		this.Round += 1
		this.Point += 1
		this.SendReport(this.ReportOne)
	}

	std.LogInfo("站后回合为 ", this.Round)

	this.ReportAll = append(this.ReportAll, this.ReportOne)

}

func (this *BattleRoom) SortUnits() []*GameUnit {
	ul := []*GameUnit{}
	for _, u := range this.Units{
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
	std.LogInfo("CheckPlayerIsMove 1", this.Units)
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
	std.LogInfo("CheckPlayerIsMove 1", this.Units)
}

func (this *BattleRoom) SetBattleUnits() {
	for _, unit := range this.Dead {
		this.Units[unit.Position] = nil
	}
}

func (this *BattleRoom) showReport()  {
	std.Log("第", this.Round + 1,"回合")
	for idx, re := range this.ReportOne.ActionList{
		std.Log("第", idx + 1, "条动作单元")
		std.Log("行动的卡牌ID ", re.InstId)
		std.Log("身上的buff变更 ", re.BuffList)

		//std.LogInfo("\tbuff", re.BuffList)
		//std.LogInfo("\tBUFF变更", re.BuffList)
		//std.LogInfo("\t本卡是否因为buff死亡", re.BuffList)

		std.Log("使用的技能 ", re.SkillId)
		std.Log("技能自带的buff ", re.SkillBuff)
		std.Log("技能释放的目标信息为")
		for idx1, l := range re.TargetList {
			std.Log("\t第", idx1 + 1, "个目标")
			std.Log("\t目标实例ID为", l.InstId)
			std.Log("\t目标受击类型", l.ActionType)
			std.Log("\t目标伤害", l.ActionParam)
			std.Log("\t目标额外信息", l.ActionParamExt)
			std.Log("\t目标是否死亡", l.Dead)
			std.Log("\t目标中的buff", l.BuffAdd)
			std.Log("\n")
		}
		std.Log("\n")
	}
}

func (this *BattleRoom) UpdateBuffState(bufflist []int32) {
	std.LogInfo("UpdateBuffState", bufflist)
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
	b,_ := json.Marshal(report)
	std.LogInfo(string(b))
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
		if u.IsDead(){
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
		if u.IsDead(){
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
		}else if pos == int(pos_left) {
			targetList = append(targetList, this.Units[pos].InstId)
			continue
		}else if pos == int(pos_right) {
			targetList = append(targetList, this.Units[pos].InstId)
			continue
		}
	}

	return targetList
}

//取得全部目标
func (this *BattleRoom) SelectMoreTarget(instid int64, num int) []int64 {
	std.LogInfo("targets start = ", num)
	unit := this.SelectOneUnit(instid)

	targets := []int64{}
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

	std.LogInfo("targets length1 = ", len(targets))

	if num > 0 && int(num) < int(len(targets)){
		rand.Seed(time.Now().UnixNano())
		l := make([]int64, len(targets))
		tmp := map[int64]int{}
		var uid int64 = 0
		for len(tmp) < num {
			std.LogInfo("len(tmp)", len(tmp), num)
			_, ok := tmp[uid]
			for !ok {
				//这里从targets里面随机选择
				idx := rand.Intn(num - 1)
				l = append(l, targets[idx])
				tmp[targets[idx]] = 1
			}
		}
	}
	std.LogInfo("targets length2 = ", len(targets))
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

	if len(u_list) == 0{
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


	std.LogInfo("目标索引",index)
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

		u_list = append(u_list, u.InstId)
	}

	std.LogInfo("友方目标",u_list)
	if len(u_list) == 1 {
		return u_list[0]
	}

	if len(u_list) == 0{
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

	std.LogInfo("一个友方目标", u_list[idx], "index", idx)

	return u_list[idx]
}

//写死
func GetNearPos(pos int32) []int32 {
	if pos < prpc.BP_RED_1 || pos >= prpc.BP_MAX{
		return []int32{}
	}
	switch int(pos) {
		case prpc.BP_RED_1 :
			return []int32{prpc.BP_BLUE_3, prpc.BP_BLUE_2, prpc.BP_BLUE_6, prpc.BP_BLUE_5, prpc.BP_BLUE_1, prpc.BP_BLUE_4}
		case prpc.BP_RED_2 :
			return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_1, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_6}
		case prpc.BP_RED_3 :
			return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_3, prpc.BP_BLUE_6}
		case prpc.BP_RED_4 :
			return []int32{prpc.BP_BLUE_3, prpc.BP_BLUE_2, prpc.BP_BLUE_6, prpc.BP_BLUE_5, prpc.BP_BLUE_1, prpc.BP_BLUE_4}
		case prpc.BP_RED_5 :
			return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_1, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_6}
		case prpc.BP_RED_6 :
			return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_3, prpc.BP_BLUE_6}
		case prpc.BP_BLUE_1 :
			return []int32{prpc.BP_RED_3, prpc.BP_RED_2, prpc.BP_RED_6, prpc.BP_RED_5, prpc.BP_RED_1, prpc.BP_RED_4}
		case prpc.BP_BLUE_2 :
			return []int32{prpc.BP_RED_2, prpc.BP_RED_1, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_6}
		case prpc.BP_BLUE_3 :
			return []int32{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_3, prpc.BP_RED_6}
		case prpc.BP_BLUE_4 :
			return []int32{prpc.BP_RED_3, prpc.BP_RED_2, prpc.BP_RED_6, prpc.BP_RED_5, prpc.BP_RED_1, prpc.BP_RED_4}
		case prpc.BP_BLUE_5 :
			return []int32{prpc.BP_RED_2, prpc.BP_RED_1, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_6}
		case prpc.BP_BLUE_6 :
			return []int32{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_3, prpc.BP_RED_6}
		default: return []int32{}
	}
	return []int32{}
}
func GetNearFriend(pos int32) []int32 {
	if pos < prpc.BP_RED_1 || pos >= prpc.BP_MAX{
		return []int32{}
	}
	switch int(pos) {
		case prpc.BP_RED_1 :
			return []int32{prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_6}
		case prpc.BP_RED_2 :
			return []int32{prpc.BP_RED_1, prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_6}
		case prpc.BP_RED_3 :
			return []int32{prpc.BP_RED_2, prpc.BP_RED_6, prpc.BP_RED_5, prpc.BP_RED_4, prpc.BP_RED_1}
		case prpc.BP_RED_4 :
			return []int32{prpc.BP_RED_1, prpc.BP_RED_5, prpc.BP_RED_2, prpc.BP_RED_3, prpc.BP_RED_6}
		case prpc.BP_RED_5:
			return []int32{prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_6, prpc.BP_RED_3, prpc.BP_RED_1}
		case prpc.BP_RED_6 :
			return []int32{prpc.BP_RED_3, prpc.BP_RED_5, prpc.BP_RED_2, prpc.BP_RED_4, prpc.BP_RED_1}

		case prpc.BP_BLUE_1 :
			return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
		case prpc.BP_BLUE_2 :
			return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_6}
		case prpc.BP_BLUE_3 :
			return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_6, prpc.BP_BLUE_5, prpc.BP_BLUE_4, prpc.BP_BLUE_1}
		case prpc.BP_BLUE_4 :
			return []int32{prpc.BP_BLUE_1, prpc.BP_BLUE_5, prpc.BP_BLUE_2, prpc.BP_BLUE_3, prpc.BP_BLUE_6}
		case prpc.BP_BLUE_5:
			return []int32{prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_6, prpc.BP_BLUE_3, prpc.BP_BLUE_1}
		case prpc.BP_BLUE_6 :
			return []int32{prpc.BP_BLUE_3, prpc.BP_BLUE_5, prpc.BP_BLUE_2, prpc.BP_BLUE_4, prpc.BP_BLUE_1}

		default: return []int32{}
	}
	return []int32{}
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

	if len(u_list) == 0{
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

	std.LogInfo("目标索引",index)
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

	buff_lis,ok:= unit.Special[prpc.BF_FRIENDLOCK]  //检测有无选择锁定随机己方的目标buff
	std.LogInfo("目标 111",unit.InstId,"己方随机目标buff ,", unit.Special[prpc.BF_FRIENDLOCK])
	if ok {
		if len(buff_lis) > 0 {
			std.LogInfo("目标 222",unit.InstId,"己方随机目标buff ,", unit.Special[prpc.BF_FRIENDLOCK])
			return this.SelectOneFriend(unit.InstId)
		}
	}
	for _, pos := range near_pos{
		if this.Units[pos] == nil{
			continue
		}
		if this.Units[pos].IsDead() {
			continue
		}
		buff_list,ok:= this.Units[pos].Special[prpc.BF_FRIENDLOCK]

		if ok {
			if len(buff_list) > 0 {
				std.LogInfo("目标  333",this.Units[pos].InstId,"己方随机目标buff ,", this.Units[pos].Special[prpc.BF_FRIENDLOCK])
				return this.SelectOneFriend(this.Units[pos].InstId)
			}
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
	for _, pos := range near_pos{
		if this.Units[pos] == nil{
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
	std.LogInfo("friends length = ", len(friends))
	return friends
}

func (this *BattleRoom) SelectThrowCard(instid int64)  (int64, int32, int32) {
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
				for _, ud := range g.UnitList{
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

	std.LogInfo("can throw all cards", canThrow)

	if len(canThrow) == 0{
		return 0, 0, 0
	}

	var del_card int64
	var idx int

	if len(canThrow) == 1{
		idx = 0
	} else {
		index := len(canThrow) - 1

		idx = rand.Intn(index)
	}

	del_card = canThrow[idx]
	del_unit := units[idx]

	std.LogInfo("throw one card", del_card)

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
		if instid == p.MyUnit.InstId{
			return p.MyUnit
		}
		for _, u := range p.UnitList {
			if u == nil{
				continue
			}
			if u.InstId == instid{
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

	std.LogInfo("selectMainUnit end ", 0)

	return 0
}

////////////////////////////////////////////////////////////////////////
////player测试
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) PlayerMove() {
	std.LogInfo("PlayerMove 1", this.Units)
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

	std.LogInfo("PlayerMove 2", this.Units)

	return
}

////////////////////////////////////////////////////////////////////////
////PVE行为
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) MonsterMove() {
	std.LogInfo("MonsterMove 1", this.Units)
	if this.Round == 0 {
		pos := this.positionMiddle(this.Monster.BattleCamp)
		this.Units[pos] = this.Monster.MainUnit
		this.Monster.MainUnit.Position = int32(pos)
	} else {
		pos := this.monsterPos(this.Monster.BattleCamp)
		if pos == prpc.BP_MAX {
			return
		}
		if len(this.Monster.BattleUnitList) == 0 {
			return
		}
		for index, m:= range this.Monster.BattleUnitList{
			if m.OutBattle{
				std.LogInfo("outbattle", m.InstId)
				continue
			}
			std.LogInfo("outbattle1111111 ", m.InstId)
			this.Units[pos] = m
			m.Position = int32(pos)
			this.Monster.BattleUnitList = this.Monster.BattleUnitList[index + 1:]
			break
		}
	}

	this.Monster.MainUnit.ChoiceSKill = this.Monster.MainUnit.SelectSkill(this.Round).SkillID

	std.LogInfo("MonsterMove 2", this.Units)

	return
}

func (this *BattleRoom)positionMiddle(camp int) int {
	if camp == prpc.CT_RED {
		return prpc.BP_RED_5
	} else {
		return prpc.BP_BLUE_5
	}
}

func (this *BattleRoom)monsterPos(camp int) int {
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
func (this *BattleRoom) ChangeCptProperty(instid int64, data int32, property string) {//CPT
	std.LogInfo("增加攻击力,目标为 ", instid, data)
	p_d := prpc.ToId_CPropertyType(property)

	unit := this.SelectOneUnit(instid)
	if unit == nil {
		for _, p := range this.PlayerList {
			if instid == p.MyUnit.InstId {
				unit = p.MyUnit
				break
			}
			for _, u := range p.UnitList {
				if u.InstId == instid{
					unit = u
					break
				}
			}
		}
	}
	std.LogInfo("属性修改前", unit.CProperties[p_d], data)

	unit.CProperties[p_d] = unit.CProperties[p_d] + float32(data)
	std.LogInfo("属性修改后", unit.CProperties[p_d], data)

	return
}

func (this *BattleRoom) ChangeIptProperty(instid int64, data int32, property string) { //IPT
	std.LogInfo("增加攻击力,目标为 ", instid, data)
	p_d := prpc.ToId_IPropertyType(property)

	unit := this.SelectOneUnit(instid)
	if unit == nil {
		for _, p := range this.PlayerList {
			if instid == p.MyUnit.InstId {
				unit = p.MyUnit
				break
			}
			for _, u := range p.UnitList {
				if u.InstId == instid{
					unit = u
					break
				}
			}
		}
	}
	std.LogInfo("属性修改前", unit.CProperties[p_d], data)

	unit.CProperties[p_d] = unit.CProperties[p_d] + float32(data)
	std.LogInfo("属性修改后", unit.CProperties[p_d], data)

	return
}

func (this *BattleRoom) MintsHp (casterid int64, target int64, damage int32, crit int32) {

	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	for _, debuff := range unit.Debuff {
		std.LogInfo("debuff", debuff)
	}

	for _, debuff := range unit.Buff {
		std.LogInfo("debuff", debuff)
	}

	_bf, _ok := unit.Special[prpc.BF_UNDAMAGE] //检测免伤

	if _ok {
		if len(_bf) > 0 {
			damage = 0
		}
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
	if float32(damage) >= unit.CProperties[prpc.CPT_CHP] {			//检测免死
		bf, ok := unit.Special[prpc.BF_UNDEAD]
		true_list := []int32{}
		if ok {
			if len(bf) > 0 {
				for  _, bid := range bf {
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
			std.LogInfo("免死触发")
			damage = 0
			if len(true_list) == 1 {			// 只有一个就删除掉这个效果
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

	//std.LogInfo("MintsHp", target, damage, t)

	std.LogInfo("攻擊  catserid ", casterid, this.AcctionList.TargetList)
	std.LogInfo("MintsHp 1  ", this.TargetCOM)

	if unit.IsDead() {
		unit.OutBattle = true
		this.isDeadOwner(casterid, target)
		this.Dead = append(this.Dead, unit)
		if unit.Owner != nil{
			unit.Owner.MyDeathNum += 1
		}
		caster := this.SelectOneUnit(casterid)
		if caster.Owner != nil {
			caster.Owner.KillUnits = append(caster.Owner.KillUnits, unit.UnitId)
		}
	}

}

func (this *BattleRoom) ThrowCard (target int64, throwcard int64, entity int32, level int32) {

	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = -0
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(prpc.BE_MAX)
	this.TargetCOM.Dead = false
	this.TargetCOM.BuffAdd = []prpc.COM_BattleBuff{}
	this.TargetCOM.ThrowCard = prpc.COM_ThrowCard{InstId:throwcard, EntityId:entity, Level:level}
	this.NewAction = false

}

func (this *BattleRoom) Throw (main int64, throwcard int64) {

	if this.Type == prpc.BT_PVP {
		for _, p :=range this.PlayerList {
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

func (this *BattleRoom) AddHp (target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	if this.NewAction == false {
		this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
		this.TargetCOM = prpc.COM_BattleActionTarget{}
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

	std.LogInfo("加血  catserid ", target, this.AcctionList.TargetList)
	std.LogInfo("AddHp 1  ", this.TargetCOM)

	//this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

//降低法术

func (this *BattleRoom) ReduceSpell (target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	unit.CProperties[prpc.CPT_MAGIC_ATK] = unit.CProperties[prpc.CPT_MAGIC_ATK] - float32(damage)

	std.LogInfo("实例ID为", target, "的卡牌在第", this.Round + 1, "法术降到",unit.CProperties[prpc.CPT_MAGIC_ATK],"降低了",damage )

	this.TargetCOM = prpc.COM_BattleActionTarget{}
	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = damage
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(int(crit))

	this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

//增加法术

func (this *BattleRoom) IncreaseSpell (target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	unit.CProperties[prpc.CPT_MAGIC_ATK] = unit.CProperties[prpc.CPT_MAGIC_ATK] + float32(damage)

	std.LogInfo("实例ID为", target, "的卡牌在第", this.Round + 1, "法术增加到",unit.CProperties[prpc.CPT_MAGIC_ATK],"增加了",damage )

	this.TargetCOM = prpc.COM_BattleActionTarget{}
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

	std.LogInfo("bufflen front", this.TargetCOM)
	this.TargetCOM.BuffAdd = append(this.TargetCOM.BuffAdd, buffCOM)
	std.LogInfo("实例ID为", target, "的卡牌在第", this.Round + 1, "回合获得了id为", buff.InstId, "的buff, buff表中的ID为", buffid, "目前该卡牌一共有", len(unit.Allbuff), "个buff, ", buff.Round)
	std.LogInfo("bufflen back", this.TargetCOM)

}

func (this *BattleRoom) DeleteBuff(target int64, buffinstid int32, data int32) {
// 去buff

	buffCOM := prpc.COM_BattleBuff{}

	unit := this.SelectOneUnit(target)

	buff := unit.SelectBuff(buffinstid)
	buff.DeleteProperty()

	buffCOM.BuffId = buff.BuffId
	buffCOM.Change = 0

	std.LogInfo("bufflen front", this.TargetCOM)
	this.TargetCOM.BuffAdd = append(this.TargetCOM.BuffAdd, buffCOM)
	std.LogInfo("实例ID为", target, "的卡牌在第", this.Round + 1, "回合失去了id为", buff.InstId, "的buff, buff表中的ID为", buff.BuffId, "目前该卡牌一共有", len(unit.Allbuff), "个buff, ", buff.Round)
	std.LogInfo("bufflen back", this.TargetCOM)

}

func (this *BattleRoom) BuffMintsHp(casterid int64, target int64, buffid int32, data int32, over bool) {
	std.LogInfo("BuffMintsHp", " buff 给id为", target, "的卡牌造成了", data, "点伤害, over", over)
	unit := this.SelectOneUnit(target)

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
		this.isDeadOwner(casterid, target)
		this.Dead = append(this.Dead, unit)
		if unit.Owner != nil{
			unit.Owner.MyDeathNum += 1
		}
		caster := this.SelectOneUnit(casterid)
		if caster.Owner != nil {
			caster.Owner.KillUnits = append(caster.Owner.KillUnits, unit.UnitId)
		}
	}
}

func (this *BattleRoom) BuffAddHp(target int64, buffid int32, data int32, over bool) {
	std.LogInfo("BuffMintsHp", " buff 给id为", target, "的卡牌增加了", data, "点血量, over", over)
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

func (this *BattleRoom) TargetOn() {
	this.TargetCOM = prpc.COM_BattleActionTarget{}
	this.NewAction = true
}


func (this *BattleRoom) TargetOver() {
	this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

func (this *BattleRoom) isDeadOwner (casterid int64, target int64) {
	unit := this.SelectOneUnit(target)

	if !unit.IsMain{
		return
	}

	if !unit.IsDead() {
		return
	}

	caster := this.SelectOneUnit(casterid)
	this.Winner = caster.Camp
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
	//std.LogInfo("IsCrit", skill.Crit)

	if per <= int(skill.Crit) {
		return 1
	}

	return 0
}

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SetupPosition(p *GamePlayer, posList []prpc.COM_BattlePosition, skillid int32) {

	std.LogInfo("SetupPosition.start", posList, p.BattleCamp)
	//if this.Round == 0 { //第一回合 必须设置主角卡
	//	for _, pos := range posList {
	//		//std.LogInfo("SetupPosition, set ", pos.InstId, p.MyUnit.InstId)
	//		if pos.InstId == p.MyUnit.InstId {
	//			goto setup_check_success
	//		}
	//		std.LogInfo("SetupPosition.error no main")
	//		return //没有主角卡
	//	}
	//}
//setup_check_success:
	var needPoint int32
	for i := 0; i < len(posList)-1; i++ {
		for j := i + 1; j < len(posList); j++ {
			if posList[i].InstId == posList[j].InstId {
				std.LogInfo("SetupPosition.error card same")
				return //有重复卡牌设置
			}
			if posList[i].Position == posList[j].Position {
				std.LogInfo("SetupPosition.error pos same")
				return //有重复位置设置
			}
			unit := p.GetUnit(posList[i].InstId)

			if unit.IsDead() {
				std.LogInfo("SetupPosition.error card is dead")
				return
			}
			needPoint += unit.Cost
		}
	}

	needPoint += 0 // 主角技能应该有消耗

	if needPoint > p.BattlePoint {
		std.LogInfo("SetupPosition.error point less", needPoint, p.BattlePoint)
		return //能量点不足
	}

	for _, pos := range posList {
		for _, u := range this.Units {
			if u == nil {
				continue
			}
			if u.InstId == pos.InstId {
				std.LogInfo("SetupPosition.error pos same 2")
				return //已经上场
			}
		}

		if this.Units[pos.Position] != nil {
			std.LogInfo("SetupPosition.error pos same 3")
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

	if this.Type == prpc.BT_PVE {
		this.BattleUpdate()
	}
	//std.LogInfo("SetupPosition", this.Units, p.BattleCamp, p.IsActive)
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