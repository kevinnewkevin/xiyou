package game

import (
	"fmt"
	"logic/prpc"
	"sync"
	"sync/atomic"
	"time"
	"math/rand"
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

type Monster struct {
	sync.Locker
	MainUnit         	*GameUnit   //自己的卡片
	BattleUnitList      []*GameUnit //拥有的卡片

	//战斗相关辅助信息
	BattleCamp 			int   //阵营 //prpc.CompType
}

type BattleRoom struct {
	sync.Mutex
	Type 	   int32		 //战斗类型 1是pvp 2是pve
	InstId     int64         //房间ID
	Status     int32         //战斗房间状态
	Round      int32         //回合计数
	Units      []*GameUnit   //当前战斗中牌 数组索引跟下面玩家对应
	PlayerList []*GamePlayer //房间中玩家信息
	Monster    *Monster
	Turn       int32
	Winner     int //获胜者
	ReportAll  []prpc.COM_BattleReport		//zhanbao all
	ReportOne  prpc.COM_BattleReport		//zhanbao
	AcctionList  prpc.COM_BattleAction		//zhanbao
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
	room.Type = 2

	room.Monster = CreateMonster(battleid)

	p.SetProprty(room.InstId, prpc.CT_RED)

	fmt.Println("CreatePvE", &room)
	BattleRoomList[room.InstId] = &room

	room.BattleStart()
	go room.BattleUpdate()

	return &room
}

func CreateMonster(battleid int32) *Monster{
	t := GetBattleRecordById(battleid)

	m := Monster{}

	m.MainUnit = CreateUnitFromTable(t.MainId)
	m.MainUnit.IsMain = true
	m.MainUnit.Camp = prpc.CT_BLUE
	fmt.Println("CreateMonster t1.SmallID", t.SmallId)

	for _, uid := range t.SmallId {
		fmt.Println("CreateMonster for", uid)
		t1 := CreateUnitFromTable(uid)
		t1.Camp = prpc.CT_BLUE
		t1.IsMain = false
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
	room.Type = 1

	//p0.BattleId = room.InstId
	//p1.BattleId = room.InstId
	//
	//p0.BattleCamp = prpc.CT_RED
	//p1.BattleCamp = prpc.CT_BLUE

	p0.SetProprty(room.InstId, prpc.CT_RED)
	p1.SetProprty(room.InstId, prpc.CT_BLUE)

	BattleRoomList[room.InstId] = &room
	fmt.Println("CreatePvP", &room)

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
	for _, p := range this.PlayerList {
		if !p.IsActive {
			return
		}
	}

	if this.Type == 2 {
		this.MonsterMove()
	}

	//顺序排序

	this.ReportOne = prpc.COM_BattleReport{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}

		//fmt.Println("report.UnitList, append", u)
		this.ReportOne.UnitList = append(this.ReportOne.UnitList, u.GetBattleUnitCOM())

		if u.IsDead() { // 非主角死亡跳過
			continue
		}

		this.AcctionList = prpc.COM_BattleAction{}

		u.CheckBuff(this.Round)
		u.CheckDebuff(this.Round)

		u.CastSkill2(this)

		this.ReportOne.ActionList = append(this.ReportOne.ActionList, this.AcctionList)

		//fmt.Println("BattleAcction", u.InstId, "acction", this.AcctionList)
		//fmt.Println("BattleAcction", u.InstId, "ReportOne", this.ReportOne)
		if this.calcWinner() == true {
			for _, a := range this.AcctionList.TargetList {
				unit := this.SelectOneUnit(a.InstId)
				if unit == nil {
					continue
				}
				//fmt.Println("append", unit)
				this.ReportOne.UnitList = append(this.ReportOne.UnitList, unit.GetBattleUnitCOM())
			}
			//fmt.Println("this.Winner", this.Winner)

			this.Round += 1
			this.SendReport(this.ReportOne)
			this.BattleRoomOver(this.Winner)
			this.Status = kIdle
			break

		}
	}
	fmt.Println("Battle report unitlist is ", this.ReportOne.UnitList)
	fmt.Println("Battle report acctionlist is ", this.ReportOne.ActionList)
	fmt.Println("Battle status ", this.Status)

	for _, p := range this.PlayerList { //戰鬥結束之後要重置屬性
		p.IsActive = false
	}

	if this.Status == kUsed {
		this.Round += 1
		this.SendReport(this.ReportOne)
	}

	this.ReportAll = append(this.ReportAll, this.ReportOne)

}

func (this *BattleRoom) calcWinner() bool {
	return this.Winner != prpc.CT_MAX
}

func (this *BattleRoom) SendReport(report prpc.COM_BattleReport) {
	fmt.Println("SendReport", "111111111111111111111111")
	for _, p := range this.PlayerList {
		p.session.BattleReport(report)
	}
}

////////////////////////////////////////////////////////////////////////
////數據處理
////////////////////////////////////////////////////////////////////////

//取得全部目标
func (this *BattleRoom) SelectAllTarget(camp int) []*GameUnit {
	targets := []*GameUnit{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.Camp == camp {
			continue
		}
		targets = append(targets, u)
	}

	return targets
}

//前排
func (this *BattleRoom) SelectFrontTarget(camp int) []*GameUnit {
	targets := []*GameUnit{}
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
		targets = append(targets, u)
	}

	return targets
}

//后排
func (this *BattleRoom) SelectBackTarget(camp int) []*GameUnit {
	targets := []*GameUnit{}
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
		targets = append(targets, u)
	}

	return targets
}

//取得全部目标
func (this *BattleRoom) SelectMoreTarget(instid int64, num int) []int64 {
	unit := this.SelectOneUnit(instid)

	targets := []int64{}
	for _, u := range this.Units {
		if u == nil {
			continue
		}
		if u.Camp == unit.Camp {
			continue
		}
		targets = append(targets, u.InstId)
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
	fmt.Println("targets length = ", len(targets))
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
////PVE行为
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) MonsterMove() {
	if this.Round == 0 {
		pos := this.monsterMiddle(this.Monster.BattleCamp)
		this.Units[pos] = this.Monster.MainUnit
		this.Monster.MainUnit.Position = int32(pos)
	} else {
		pos := this.monsterPos(this.Monster.BattleCamp)
		this.Units[pos] = this.Monster.BattleUnitList[0]
		this.Monster.BattleUnitList[0].Position = int32(pos)
		this.Monster.BattleUnitList = this.Monster.BattleUnitList[1:]
	}

	return
}

func (this *BattleRoom)monsterMiddle(camp int) int {
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

func (this *BattleRoom) MintsHp (casterid int64, target int64, damage int32, crit int32) {

	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	for _, debuff := range unit.Debuff {
		fmt.Println("debuff", debuff)
	}

	for _, debuff := range unit.Buff {
		fmt.Println("debuff", debuff)
	}

	unit.CProperties[prpc.CPT_HP] = unit.CProperties[prpc.CPT_HP] - float32(damage)

	t := prpc.COM_BattleActionTarget{}
	t.InstId = target
	t.ActionType = 1
	t.ActionParam = damage
	t.ActionParamExt = crit

	//fmt.Println("MintsHp", target, damage, t)

	this.AcctionList.TargetList = append(this.AcctionList.TargetList, t)

	//fmt.Println("MintsHp2 ", this.AcctionList.TargetList)

	this.isDeadOwner(casterid, target)

}
func (this *BattleRoom) AddHp (target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	unit.CProperties[prpc.CPT_HP] = unit.CProperties[prpc.CPT_HP] - float32(damage)

	t := prpc.COM_BattleActionTarget{}
	t.InstId = target
	t.ActionType = 1
	t.ActionParam = damage
	t.ActionParamExt = crit

	this.AcctionList.TargetList = append(this.AcctionList.TargetList, t)
}

func (this *BattleRoom) AddBuff(target int64, buffid int32, data int32) {
// 上buff

	unit := this.SelectOneUnit(target)

	buff := NewBuff(unit, buffid, data, this.Round)

	unit.Buff = append(unit.Buff, buff)

}

func (this *BattleRoom) AddDebuff(target int64, buffid int32, data int32) {
// 上buff

	unit := this.SelectOneUnit(target)

	buff := NewBuff(unit, buffid, data, this.Round)

	unit.Debuff = append(unit.Debuff, buff)

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

	per := rand.Intn(100)
	//fmt.Println("IsCrit", skill.Crit)

	if per <= int(skill.Crit) {
		return 1
	}

	return 0
}

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) SetupPosition(p *GamePlayer, posList []prpc.COM_BattlePosition) {

	//Println("SetupPosition.start")
	if this.Round == 0 { //第一回合 必须设置主角卡
		for _, pos := range posList {
			//fmt.Println("SetupPosition, set ", pos.InstId, p.MyUnit.InstId)
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

	//fmt.Println("SetupPosition", this.Units, p.BattleCamp, p.IsActive)
}
