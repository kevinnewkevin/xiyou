package game

import (
	"fmt"
	"logic/prpc"
	"sync"
	"sync/atomic"
	"time"
	"math/rand"
	"sort"
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

	kTimeSleep = 5   	//檢測間隔
	kTimeMax   = 50000	//戰鬥持續時間
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
	room.Point = 1
	room.BattleID = battleid

	room.Monster = CreateMonster(battleid, room.InstId)

	p.SetProprty(room.InstId, prpc.CT_RED)

	fmt.Println("CreatePvE", &room)
	BattleRoomList[room.InstId] = &room

	room.BattleStart()
	go room.BattleUpdate()

	return &room
}

func CreateMonster(battleid int32, roomid int64) *Monster{
	t := GetBattleRecordById(battleid)

	m := Monster{}

	m.MainUnit = CreateUnitFromTable(t.MainId)
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
	room.Type = 1
	room.Point = 1
	room.BattleID = 0

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
		p.session.JoinBattleOk(int32(p.BattleCamp), this.BattleID)
	}
}

//臨時用
func (this *BattleRoom) BattleUpdate() {
	start := time.Now().Unix()
	now_start := time.Now().Unix()
	//Round_start := time.Now().Unix()
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
		round := this.Round
		deathnum := p.MyDeathNum
		killmonster := p.KillUnits

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
		result.BattleRound = round
		result.KillMonsters = killmonster
		result.MySelfDeathNum = deathnum

		if this.Type == prpc.BT_PVE {
			p.CalcSmallChapterStar(result)
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
				CaleTianTiVal(p,once,camp)
			}
		}
		p.BattleId = 0
		p.BattleCamp = prpc.CT_MAX

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
		//this.PlayerMove()
	}

	//this.CheckPlayerIsMove()

	//顺序排序

	fmt.Println("站前回合为 ", this.Round)

	this.ReportOne = prpc.COM_BattleReport{}
	this.Dead = []*GameUnit{}
	this.ReportOne.BattleID = this.BattleID

	unitllist := this.SortUnits()

	for _, unit := range unitllist {
		if unit == nil {
			continue
		}
		this.ReportOne.UnitList = append(this.ReportOne.UnitList, unit.GetBattleUnitCOM())
	}

	for _, u := range unitllist {
		if u == nil {
			continue
		}

		//fmt.Println("report.UnitList, append", u)
		//this.ReportOne.UnitList = append(this.ReportOne.UnitList, u.GetBattleUnitCOM())

		if u.IsDead() { // 非主角死亡跳過
			continue
		}

		this.AcctionList = prpc.COM_BattleAction{}
		//this.TargetOn()

		u.CheckBuff(this.Round)
		u.CheckDebuff(this.Round)

		del_buf := u.CheckAllBuff(this.Round)
		this.UpdateBuffState(del_buf)

		if u.IsDead() { // 非主角死亡跳過
			continue
		}

		u.CastSkill2(this)

		//this.TargetOver()

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
	fmt.Println("Battle report battleid is ", this.ReportOne.BattleID)
	fmt.Println("Battle report unitlist is ", this.ReportOne.UnitList)
	fmt.Println("Battle report acctionlist is ", this.ReportOne.ActionList)

	this.showReport()

	this.SetBattleUnits()

	fmt.Println("Battle status ", this.Status)

	for _, p := range this.PlayerList { //戰鬥結束之後要重置屬性
		p.IsActive = false
	}

	if this.Status == kUsed {
		this.Round += 1
		this.Point += 1
		this.SendReport(this.ReportOne)
	}

	fmt.Println("站后回合为 ", this.Round)

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
	fmt.Println("CheckPlayerIsMove 1", this.Units)
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
	fmt.Println("CheckPlayerIsMove 1", this.Units)
}

func (this *BattleRoom) SetBattleUnits() {
	for _, unit := range this.Dead {
		this.Units[unit.Position] = nil
	}
}

func (this *BattleRoom) showReport()  {
	fmt.Println("第", this.Round + 1,"回合")
	for idx, re := range this.ReportOne.ActionList{
		fmt.Println("第", idx + 1, "条动作单元")
		fmt.Println("行动的卡牌ID ", re.InstId)
		fmt.Println("身上的buff变更 ", re.BuffList)

		//fmt.Println("\tbuff", re.BuffList)
		//fmt.Println("\tBUFF变更", re.BuffList)
		//fmt.Println("\t本卡是否因为buff死亡", re.BuffList)

		fmt.Println("使用的技能 ", re.SkillId)
		fmt.Println("技能自带的buff ", re.SkillBuff)
		fmt.Println("技能释放的目标信息为")
		for idx1, l := range re.TargetList {
			fmt.Println("\t第", idx1 + 1, "个目标")
			fmt.Println("\t目标实例ID为", l.InstId)
			fmt.Println("\t目标受击类型", l.ActionType)
			fmt.Println("\t目标伤害", l.ActionParam)
			fmt.Println("\t目标额外信息", l.ActionParamExt)
			fmt.Println("\t目标是否死亡", l.Dead)
			fmt.Println("\t目标中的buff", l.BuffAdd)
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}

func (this *BattleRoom) UpdateBuffState(bufflist []int32) {
	fmt.Println("UpdateBuffState", bufflist)
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
	fmt.Println("SendReport", "111111111111111111111111")
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
	fmt.Println("targets start = ", num)
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

	fmt.Println("targets length1 = ", len(targets))

	//if num > 0 && int(num) < int(len(targets)){
	//	rand.Seed(time.Now().UnixNano())
	//	l := make([]int64, len(targets))
	//	tmp := map[int64]int{}
	//	var uid int64 = 0
	//	for len(tmp) < num {
	//		_, ok := tmp[uid]
	//		for !ok {
	//			//这里从targets里面随机选择
	//			idx := rand.Intn(num - 1)
	//			l = append(l, targets[idx])
	//			tmp[targets[idx]] = 1
	//		}
	//	}
	//}
	//fmt.Println("targets length2 = ", len(targets))
	return targets
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
	fmt.Println("friends length = ", len(friends))
	return friends
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
////player测试
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) PlayerMove() {
	fmt.Println("PlayerMove 1", this.Units)
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

	fmt.Println("PlayerMove 2", this.Units)

	return
}

////////////////////////////////////////////////////////////////////////
////PVE行为
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) MonsterMove() {
	fmt.Println("MonsterMove 1", this.Units)
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
		this.Units[pos] = this.Monster.BattleUnitList[0]
		this.Monster.BattleUnitList[0].Position = int32(pos)
		this.Monster.BattleUnitList = this.Monster.BattleUnitList[1:]
	}

	fmt.Println("MonsterMove 2", this.Units)

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
func (this *BattleRoom) ChangeUnitProperty(instid int64, data int32, property string) {
	fmt.Println("增加攻击力,目标为 ", instid, data)
	p_d := prpc.ToId_CPropertyType(property)

	unit := this.SelectOneUnit(instid)
	fmt.Println("属性修改前", unit.CProperties[p_d], data)

	unit.CProperties[p_d] = unit.CProperties[p_d] + float32(data)
	fmt.Println("属性修改后", unit.CProperties[p_d], data)

	return
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

	////////////////////////////////////////////////////////////////////////
	if unit.VirtualHp >= damage {			//计算护盾减伤之后的伤害
		damage = 0
		unit.VirtualHp = unit.VirtualHp - damage
	} else {
		damage = damage - unit.VirtualHp
	}
	////////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////////
	per := unit.ClacSheldPer(this.Round)		//百分比减伤 结果是最终减伤多少
	if per > 0 {
		persent := 1.0 - per					//实际受到的伤害百分比
		damage = int32(float32(damage) * persent)
	}
	////////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////////
	if float32(damage) >= unit.CProperties[prpc.CPT_CHP] {			//检测免死
		bf, ok := unit.Special[prpc.BF_UNDEAD]
		if ok && len(bf) > 0 {
			fmt.Println("免死触发")
			damage = 0
			if len(bf) == 1 {			// 只有一个就删除掉这个效果
				delete(unit.Special, prpc.BF_UNDEAD)
			} else {
				unit.Special[prpc.BF_UNDEAD] = bf[:1]
			}
			buff := unit.SelectBuff(bf[0])
			buff.Over = true
		}
	}
	////////////////////////////////////////////////////////////////////////

	unit.CProperties[prpc.CPT_CHP] = unit.CProperties[prpc.CPT_CHP] - float32(damage)

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

	//fmt.Println("MintsHp", target, damage, t)

	fmt.Println("攻擊  catserid ", casterid, this.AcctionList.TargetList)
	fmt.Println("MintsHp 1  ", this.TargetCOM)

	if unit.IsDead() {
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
func (this *BattleRoom) AddHp (target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	unit.CProperties[prpc.CPT_CHP] = unit.CProperties[prpc.CPT_CHP] + float32(damage)

	this.TargetCOM = prpc.COM_BattleActionTarget{}
	this.TargetCOM.InstId = target
	this.TargetCOM.ActionType = 1
	this.TargetCOM.ActionParam = damage
	this.TargetCOM.ActionParamExt = prpc.ToName_BattleExt(int(crit))

	//this.AcctionList.TargetList = append(this.AcctionList.TargetList, this.TargetCOM)
}

//降低法术

func (this *BattleRoom) ReduceSpell (target int64, damage int32, crit int32) {
	unit := this.SelectOneUnit(target)

	if unit.IsDead(){
		return
	}

	unit.CProperties[prpc.CPT_MAGIC_ATK] = unit.CProperties[prpc.CPT_MAGIC_ATK] - float32(damage)

	fmt.Println("实例ID为", target, "的卡牌在第", this.Round + 1, "法术降到",unit.CProperties[prpc.CPT_MAGIC_ATK],"降低了",damage )

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

	fmt.Println("实例ID为", target, "的卡牌在第", this.Round + 1, "法术增加到",unit.CProperties[prpc.CPT_MAGIC_ATK],"增加了",damage )

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

	fmt.Println("bufflen front", this.TargetCOM)
	this.TargetCOM.BuffAdd = append(this.TargetCOM.BuffAdd, buffCOM)
	fmt.Println("实例ID为", target, "的卡牌在第", this.Round + 1, "回合获得了id为", buff.InstId, "的buff, buff表中的ID为", buffid, "目前该卡牌一共有", len(unit.Allbuff), "个buff, ", buff.Round)
	fmt.Println("bufflen back", this.TargetCOM)

}

func (this *BattleRoom) DeleteBuff(target int64, buffinstid int32, data int32) {
// 去buff

	buffCOM := prpc.COM_BattleBuff{}

	unit := this.SelectOneUnit(target)

	buff := unit.SelectBuff(buffinstid)
	buff.DeleteProperty()

	buffCOM.BuffId = buff.BuffId
	buffCOM.Change = 0

	fmt.Println("bufflen front", this.TargetCOM)
	this.TargetCOM.BuffAdd = append(this.TargetCOM.BuffAdd, buffCOM)
	fmt.Println("实例ID为", target, "的卡牌在第", this.Round + 1, "回合失去了id为", buff.InstId, "的buff, buff表中的ID为", buff.BuffId, "目前该卡牌一共有", len(unit.Allbuff), "个buff, ", buff.Round)
	fmt.Println("bufflen back", this.TargetCOM)

}

func (this *BattleRoom) BuffMintsHp(casterid int64, target int64, buffid int32, data int32, over bool) {
	fmt.Println("BuffMintsHp", " buff 给id为", target, "的卡牌造成了", data, "点伤害, over", over)
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
	fmt.Println("BuffMintsHp", " buff 给id为", target, "的卡牌增加了", data, "点血量, over", over)
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

	fmt.Println("SetupPosition.start", posList)
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
	var needPoint int32
	for i := 0; i < len(posList)-1; i++ {
		for j := i + 1; j < len(posList); j++ {
			if posList[i].InstId == posList[j].InstId {
				return //有重复卡牌设置
			}
			if posList[i].Position == posList[j].Position {
				return //有重复位置设置
			}
			unit := p.GetUnit(posList[i].InstId)
			needPoint += unit.Cost
		}
	}

	if needPoint > p.BattlePoint {
		return //能量点不足
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
	p.BattlePoint -= needPoint

	//fmt.Println("SetupPosition", this.Units, p.BattleCamp, p.IsActive)
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