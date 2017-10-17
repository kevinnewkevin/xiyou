package game

import (
	"errors"
	"logic/prpc"
	"strings"
	"strconv"
	"logic/std"
)

const (
	unitGroupMax		= 5			//卡组上限
	onceUnitGroupMax 	= 10			//每组卡片上限
	bagMaxGrid			= 200
)

type GamePlayer struct {
	session        			*Session    		//链接
	PlayerId 			int64			//角色ID
	Username 			string 			//账户名
	MyUnit         			*GameUnit   		//自己的卡片
	UnitList       			[]*GameUnit 		//拥有的卡片
	BattleUnitList 			[]int64     		//默认出战卡片
	//DefaultUnitGroup		int			//默认战斗卡片组
	BattleUnitGroup			int32			//战斗卡片组
	UnitGroup				[]*prpc.COM_UnitGroup
	//战斗相关辅助信息
	BattleId   				int64 		//所在房间编号
	BattleCamp 				int   		//阵营 //prpc.CompType
	IsActive   				bool  		//是否激活
	KillUnits 	 			[]int32 	//杀掉的怪物
	MyDeathNum				int32		//战斗中自身死亡数量
	BattlePoint				int32		//战斗點數

	//story chapter
	ChapterID				int32		//正在进行的关卡
	Chapters				[]*prpc.COM_Chapter

	TianTiVal				int32
	EnergyTimer				float64
	//Bag
	BagItems				[]*prpc.COM_ItemInst

	//经验
	Exp 					int32
	//主角可学习技能
	SkillBase				map[int32]int32
}

var (
	PlayerStore	[]*GamePlayer = []*GamePlayer{}
	DefaultUnits	[]int32
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//角色创建
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func FindPlayerByInstId(instid int64) *GamePlayer{
	for _, p :=range PlayerStore{
		if p == nil {
			continue
		}
		if p.MyUnit.InstId == instid {
			return p
		}
	}
	return nil
}

func FindPlayerByUsername(username string) *GamePlayer {
	for _,p:=range PlayerStore{
		if p == nil {
			continue
		}
		if p.Username == username {
			return p
		}
	}

	return nil
}

func FindPlayerByInstName(instName string) *GamePlayer {
	for _,p:=range PlayerStore{
		if p == nil {
			continue
		}
		if p.MyUnit.InstName == instName {
			return p
		}
	}

	return nil
}

func RemovePlayerByInstName(instName string){
	for i:=0;i<len(PlayerStore) ;i++  {
		if PlayerStore[i] == nil {
			continue
		}
		if PlayerStore[i].MyUnit.InstName == instName {
			PlayerStore = append(PlayerStore[:i],PlayerStore[i+1:]...)
		}
	}
	std.LogInfo("RemovePlayerByInstName Name=",instName,"PlayerStore len=",len(PlayerStore))
}

func RemovePlayer(player *GamePlayer)  {
	if player==nil {
		return
	}
	RemovePlayerByInstName(player.MyUnit.InstName)
}

func SetDefaultUnits(cards string) {
	s1 := strings.Split(cards, ",")
	for _, c := range s1{
		e_id, _ := strconv.Atoi(c)
		DefaultUnits = append(DefaultUnits, int32(e_id))
	}
}

func PlayerTick(dt float64)  {
	for _, p :=range PlayerStore{
		if p == nil {
			continue
		}
		p.CaleMyEnergy(dt)
	}
}

func (this *GamePlayer) SetSession(session *Session) {
	this.session = session
}


func CreatePlayer(tid int32, name string) *GamePlayer {
	p := GamePlayer{}
	p.MyUnit = p.NewGameUnit(tid)
	p.MyUnit.InstName = name
	p.MyUnit.IsMain = true
	p.Exp = 0

	for _, e_id := range DefaultUnits {
		p.NewGameUnit(e_id)
	}
	//p.DefaultUnitGroup = 1
	p.TianTiVal	= 0
	PlayerStore = append(PlayerStore, &p)
	p.InitUnitGroup()

	for _,u := range p.UnitList{
		std.LogInfo("Myself Unit InstId",u.InstId,"InstName",u.InstName)
	}
	p.SkillBase = map[int32]int32{}

	for ID, info := range RoleSkillTable {
		if p.MyUnit.Level >= info.OpenLv {
			p.SkillBase[ID] = info.SKillID
		} else {
			p.SkillBase[ID] = 0
		}
	}

	return &p

}

func (this *GamePlayer) NewGameUnit(tid int32) *GameUnit {
	unit := CreateUnitFromTable(tid)
	if unit==nil {
		return nil
	}
	unit.Owner = this
	chapterids := GetUnitChapterById(tid)
	for i:=0;i<len(chapterids) ;i++  {
		OpenChapter(this,chapterids[i])
	}
	this.UnitList = append(this.UnitList, unit)
	return unit
}

func (this *GamePlayer) SetPlayerCOM(p *prpc.COM_Player){
	this.MyUnit = &GameUnit{}
	this.MyUnit.Owner = this
	this.MyUnit.SetUnitCOM(&p.Unit)

	for _, u := range p.Employees {
		unit := GameUnit{}
		unit.SetUnitCOM(&u)
		unit.Owner = this
		this.UnitList = append(this.UnitList, &unit)
	}
	for i, _ := range p.Chapters {
		this.Chapters = append(this.Chapters,&p.Chapters[i])
	}

	for i ,_ := range p.UnitGroup{
		this.UnitGroup = append(this.UnitGroup,&p.UnitGroup[i])
	}
	this.TianTiVal = p.TianTiVal

	this.SkillBase = map[int32]int32{}
	for _, skb:= range p.SkillBase {
		this.SkillBase[skb.SkillId] = skb.SkillId
	}
}

func (this *GamePlayer) GetPlayerCOM() prpc.COM_Player {
	//this.Lock()
	//defer this.Unlock()
	p := prpc.COM_Player{}
	p.InstId = this.MyUnit.InstId
	p.Name = this.MyUnit.InstName
	p.Unit = this.MyUnit.GetUnitCOM()
	for _, u := range this.UnitList {
		p.Employees = append(p.Employees, u.GetUnitCOM())
	}
	for _, c := range this.Chapters {
		p.Chapters = append(p.Chapters, *c)
	}
	for _,ug := range this.UnitGroup{
		p.UnitGroup = append(p.UnitGroup,*ug)
	}
	p.TianTiVal = this.TianTiVal

	for index, skillid := range this.SkillBase {
		skillbase := prpc.COM_SkillBase{}
		skillbase.SkillIdx = index
		skillbase.SkillId = skillid

		p.SkillBase = append(p.SkillBase, skillbase)
	}

	//


	return p
}

func( this* GamePlayer) SetPlayerSGE(p prpc.SGE_DBPlayer){
	this.SetPlayerCOM(&p.COM_Player)
	this.PlayerId = p.PlayerId
	this.Username = p.Username
	for _, a := range p.BagItemList{
		this.BagItems = append(this.BagItems,&a)
	}
}

func (this* GamePlayer) GetPlayerSGE() prpc.SGE_DBPlayer{
	items := []prpc.COM_ItemInst{}
	for _, a := range this.BagItems{
		items = append(items,*a)
	}
	return prpc.SGE_DBPlayer{COM_Player:this.GetPlayerCOM(),PlayerId:this.PlayerId,Username:this.Username ,BagItemList:items}
}

func (this *GamePlayer)IsBattle() bool {
	if this.BattleId == 0 {
		return false
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//角色数据接口
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (this *GamePlayer) GetUnit(instId int64) *GameUnit {
	//this.Lock()
	//defer this.Unlock()
	if this.MyUnit.InstId == instId {
		return this.MyUnit
	}

	for _, v := range this.UnitList {
		if v.InstId == instId {
			return v
		}
	}
	return nil
}
func (this *GamePlayer) GetBattleUnit(instId int64) *GameUnit {
	//this.Lock()
	//defer this.Unlock()
	for _, v := range this.BattleUnitList {
		if v == instId {
			return this.GetUnit(instId)
		}
	}
	return nil
}

func (this *GamePlayer) HasUnitByTableId(tableId int32) bool{
	for _, v := range this.UnitList {
		if v.UnitId == tableId {
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//技能相关
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) StudySkill(UnitID int64, skillpos int32, skillid int32) error {
	if skillpos >= 2 {
		std.LogInfo("技能位置錯誤")
		return errors.New("技能位置錯誤")
	}
	unit := this.GetUnit(UnitID)
	skill := InitSkillFromTable(skillid)

	unit.Skill[skillpos] = skill

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//战斗相关 设置卡牌
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//dont care mutli thread
var battlePlayerList = []*GamePlayer{}

func (this *GamePlayer) JoinBattle() {
	//this.Lock()
	//defer this.Unlock()

	for _, v := range battlePlayerList {
		if v == this {
			return
		}
	}

	battlePlayerList = append(battlePlayerList, this)

	if len(battlePlayerList) == 2 {
		//把他俩都拉到战斗力去			这里还要加一个判断,不能重复加入战斗
		CreatePvP(battlePlayerList[0], battlePlayerList[1])
		//CreatePvE(battlePlayerList[0], 1)

		battlePlayerList = battlePlayerList[:0]
	}
	std.LogInfo("JoinBattle", battlePlayerList)
}

func (this *GamePlayer) JoinBattlePvE(bigGuanqia int32, SmallGuanqia int32) {
	//this.Lock()
	//defer this.Unlock()

	CreatePvE(battlePlayerList[0], 1)

	std.LogInfo("JoinBattlePvE", battlePlayerList)
}


func (this *GamePlayer) LeftBattle_strong() {
	battle:= FindBattle(this.BattleId)
	if battle != nil {
		battle.PlayerLeft(this)
	}
	this.BattleId = 0
	this.BattleCamp = prpc.CT_MAX
	this.ClearAllBuff()
}

func (this *GamePlayer) SetBattleUnit(instId int64) { //往战斗池里设置出战卡牌  战斗开始之前
	//this.Lock()
	//defer this.Unlock()
	if instId == 0 {
		return // 0是錯誤的
	}

	if this.GetUnit(instId) == nil {
		return //没有设置你妹
	}
	if this.GetBattleUnit(instId) != nil {
		return //在出战设置你妹
	}
	this.BattleUnitList = append(this.BattleUnitList, instId)
	std.LogInfo("SetBattleUnit ", this.BattleUnitList)
}

func (this *GamePlayer)InitUnitGroup()  {
	for i:=0;i<unitGroupMax ;i++  {
		unitgroup := prpc.COM_UnitGroup{}
		unitgroup.GroupId = int32(i+1)
		this.UnitGroup = append(this.UnitGroup,&unitgroup)
	}
}

func (this *GamePlayer)SetBattleUnitGroup(instId int64, groupId int32, isBattle bool)  {
	if this == nil {
		return
	}
	if this.GetUnit(instId) == nil {
		return
	}
	std.LogInfo("SetBattleUnitGroup InstId",instId,"GroupID",groupId,"IsBattle",isBattle)
	if isBattle {
		addError := this.AddUnitToGroup(instId,groupId)
		if addError != 0{
			std.LogInfo("AddUnitToGroup Error",addError)
		}
	}else {
		addError := this.RemoveUnitToGroup(instId,groupId)
		if addError != 0{
			std.LogInfo("RemoveUnitToGroup Error",addError)
		}
	}
}

func (this *GamePlayer)GetUnitGroupById(groupId int32) *prpc.COM_UnitGroup {
	for _,g:=range this.UnitGroup{
		if g==nil {
			continue
		}
		if g.GroupId == groupId {
			return g;
		}
	}
	return nil
}

func (this *GamePlayer)AddUnitToGroup(instId int64,groupId int32) int {
	group := this.GetUnitGroupById(groupId)
	if group==nil {
		return 1
	}

	if len(group.UnitList) >= onceUnitGroupMax {
		return 2
	}

	card := this.GetUnit(instId)
	if card==nil {
		return 3
	}

	for _,unit := range group.UnitList{
		tmp := this.GetUnit(unit)
		if tmp==nil {
			continue
		}
		if card.UnitId == tmp.UnitId {
			return 4
		}
	}

	group.UnitList = append(group.UnitList,card.InstId)

	return 0
}

func (this *GamePlayer)RemoveUnitToGroup(instId int64,groupId int32) int {
	group := this.GetUnitGroupById(groupId)
	if group==nil {
		return 1
	}

	card := this.GetUnit(instId)
	if card==nil {
		return 2
	}

	index := 100

	for i:=0;i<len(group.UnitList);i++  {
		if group.UnitList[i] == instId {
			index = i
		}
	}

	if index > onceUnitGroupMax {
		return 3
	}

	group.UnitList = append(group.UnitList[:index], group.UnitList[index+1:]...)

	return 0
}

func (this *GamePlayer)DeleteUnitGroup(groupId int32)  {
	for i:=0;i<len(this.UnitGroup) ;i++  {
		if this.UnitGroup[i].GroupId == groupId {
			unitgroup := prpc.COM_UnitGroup{}
			this.UnitGroup[i] = &unitgroup
		}
	}
}

func (this *GamePlayer) SetupBattle(pos []prpc.COM_BattlePosition , skillid int32) error { //卡牌上阵	每次回合之前
	//this.Lock()
	//defer this.Unlock()
	//std.LogInfo("SetupBattle", pos)
	for _, p := range pos {
		//if this.GetBattleUnit(int64(p.InstId)) == nil {
		//	return nil //错误消息
		//}
		if p.Position >= prpc.BP_MAX || p.Position < prpc.BP_RED_1 {
			return nil //错误消息 //检测缺失 阵营与位置关系
		}
	}

	battleRoom := FindBattle(this.BattleId)

	if battleRoom == nil {
		//错误消息
		return nil
	}
	std.LogInfo("SetupBattle 1 ", battleRoom.Units)
	battleRoom.SetupPosition(this, pos, skillid)
	std.LogInfo("SetupBattle 2 ", battleRoom.Units)

	this.session.SetupBattleOK()

	//battleRoom.Update()

	return nil
}

func (this *GamePlayer) SetProprty(battle *BattleRoom, camp int) {
	this.BattleId = battle.InstId
	this.BattleCamp = camp
	this.BattlePoint = 1
	this.MyUnit.ResetBattle(camp, true, battle.InstId)
	this.MyUnit.CastPassiveSkill(battle)

	for _, u := range this.UnitList {
		if u.IsMain {
			continue
		}
		u.ResetBattle(camp, false, battle.InstId)
		u.CastPassiveSkill(battle)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer)SyncBag()  {
	items := []prpc.COM_ItemInst{}

	for _,itemInst := range this.BagItems {
		items = append(items,*itemInst)
	}

	for _,item := range items{
		std.LogInfo("To Client Item TableId=",item.ItemId,"Stack=",item.Stack,"InstId=",item.InstId)
	}

	if this.session != nil {
		this.session.InitBagItems(items)
	}
}

func (this *GamePlayer)AddBagItemByItemId(itemId int32,itemCount int32)  {
	itemData := GetItemTableDataById(itemId)
	if itemData==nil {
		std.LogInfo("ItemTable Not Find This ItemId=",itemId)
		return
	}
	for _,itemInst := range this.BagItems{
		if itemInst.ItemId == itemId {
			itemInst.Stack += itemCount
			if itemInst.Stack > itemData.MaxCount {
				itemCount = itemInst.Stack - itemData.MaxCount
				itemInst.Stack = itemData.MaxCount
			}else {
				itemCount = 0
			}

			//updata bag itemInst
			if this.session != nil {
				this.session.UpdateBagItem(*itemInst)
			}
			if itemCount==0 {
				break;
			}
		}
	}
	if itemCount > 0 {
		newItems := GenItemInst(itemId,itemCount)
		if len(newItems) == 0 {
			return
		}
		if len(this.BagItems) + len(newItems) >= bagMaxGrid {
			//newItems To Mall
			return
		}

		for _,newItem := range newItems{
			if newItem == nil {
				continue
			}
			this.BagItems = append(this.BagItems,newItem)
			//add newItem
			if this.session != nil {
				this.session.AddBagItem(*newItem)
			}
		}
	}
}

func (this *GamePlayer)DelItemByInstId(instid int64,stack int32)  {
	itemInst := this.GetBagItemByInstId(instid)
	if itemInst == nil {
		std.LogInfo("Not Find Item In The Bag",instid)
		return
	}
	for i:=0;i<len(this.BagItems) ;i++  {
		if this.BagItems[i] == nil{
			continue
		}
		if this.BagItems[i].InstId == instid {
			if this.BagItems[i].Stack > stack {
				this.BagItems[i].Stack -= stack
				//updata item
				if this.session != nil {
					this.session.UpdateBagItem(*itemInst)
				}
			}else {
				this.BagItems = append(this.BagItems[:i], this.BagItems[i+1:]...)
				//del item
				if this.session != nil {
					this.session.DeleteItemOK(instid)
				}
			}
		}
	}
}

func (this *GamePlayer)DelItemByTableId(tableId int32,delNum int32)  {
	items := this.GetBagItemByTableId(tableId)
	if len(items) == 0 {
		std.LogInfo("Can Not Find Item In Bag By TableId =",tableId)
		return
	}
	for _,item := range items{
		if item.Stack > delNum {
			item.Stack -= delNum
			if this.session != nil {
				this.session.UpdateBagItem(*item)
			}
		}else {
			delNum -= item.Stack
			this.DelItemByInstId(item.InstId,item.Stack)
		}
	}
}

func (this *GamePlayer)GetBagItemByInstId(instId int64) *prpc.COM_ItemInst {
	for _,itemInst := range this.BagItems{
		if itemInst == nil {
			continue
		}
		if itemInst.InstId == instId {
			return itemInst
		}
	}
	return nil
}

func (this *GamePlayer)GetBagItemByTableId(itemid int32) []*prpc.COM_ItemInst {
	items := []*prpc.COM_ItemInst{}
	for _,itemInst := range this.BagItems{
		if itemInst == nil {
			continue
		}
		if itemInst.ItemId == itemid {
			items = append(items,itemInst)
		}
	}
	return items
}

func (this *GamePlayer)UseItem(instId int64,useNum int32)  {
	itemInst := this.GetBagItemByInstId(instId)
	if itemInst==nil {
		return
	}

	if itemInst.Stack < useNum {
		return
	}

	itemData := GetItemTableDataById(itemInst.ItemId)
	if itemData==nil {
		return
	}

	v := []interface{}{int64(this.MyUnit.InstId),int64(itemData.ItemId)}
	r := []interface{}{""}

	usestack := 0

	for i := 0; i < int(useNum) ; i++ {
		_L.CallFuncEx(itemData.GloAction, v, &r)
		errorNo := r[0]
		if errorNo != "" {
			errorId := prpc.ToId_ErrorNo(errorNo.(string))
			if this.session != nil {
				this.session.ErrorMessage(errorId)
			}
			std.LogInfo("useItem errorId=",errorId,"itemId=",itemData.ItemId)
			break
		}
		usestack++
	}
	if usestack != 0 {
		this.DelItemByInstId(instId,int32(usestack))
	}
	std.LogInfo("123123123123",r,len(r),r[0],usestack)
}

func (this *GamePlayer)GiveDrop(dropId int32)  {
	drop := GetDropById(dropId)
	if drop==nil {
		std.LogInfo("Can Not Find Drop By DropId=",dropId)
		return
	}
	if drop.Exp != 0 {
		this.AddExp(drop.Exp)
	}
	if drop.Money != 0 {
		this.AddCopper(drop.Money)
	}
	if len(drop.Items) != 0 {
		for _,item := range drop.Items{
			this.AddBagItemByItemId(item.ItemId,item.ItemNum)
			std.LogInfo("PlayerName=",this.MyUnit.InstName,"GiveDrop AddItem ItemId=",item.ItemId,"ItemNum=",item.ItemNum)
		}
	}
	if drop.Hero != 0 {
		if this.HasUnitByTableId(drop.Hero) {
			//有这个卡就不给了
			std.LogInfo("PlayerName=",this.MyUnit.InstName,"GiveDrop AddUnit Have not to UnitId=",drop.Hero)
		}else {
			unit := this.NewGameUnit(drop.Hero)
			if unit!=nil {
				std.LogInfo("PlayerName=",this.MyUnit.InstName,"GiveDrop AddUnit OK UnitId=",drop.Hero)
				temp := unit.GetUnitCOM()
				if this.session != nil {
					this.session.AddNewUnit(temp)
				}
			}
		}
	}
}

func (this *GamePlayer)AddExp(val int32)  {
	curExp := this.MyUnit.GetIProperty(prpc.IPT_EXPERIENCE)
	curExp += val
	if curExp<0 {
		curExp = 0
	}
	//在这里加上对于经验值和等级的判断

	curExp = this.MyUnit.CheckExp(curExp)

	this.MyUnit.SetIProperty(prpc.IPT_EXPERIENCE,curExp)

	std.LogInfo("append EXP",val,"all EXP",curExp)
}

func (this *GamePlayer)AddCopper(val int32)  {
	oldCopper := this.MyUnit.GetIProperty(prpc.IPT_COPPER)
	curCopper := oldCopper + val
	if curCopper < 0  {
		curCopper=0
	}
	if curCopper>CopperMax {
		curCopper=CopperMax
	}
	this.MyUnit.SetIProperty(prpc.IPT_COPPER,curCopper)

	std.LogInfo("Player[",this.MyUnit.InstName,"]","Old Copper=",oldCopper,"curCopper=",curCopper)
}

func (this *GamePlayer)AddGold(val int32)  {
	oldGold := this.MyUnit.GetIProperty(prpc.IPT_GOLD)
	curGold := oldGold + val
	if curGold < 0  {
		curGold=0
	}
	if curGold>CopperMax {
		curGold=CopperMax
	}
	this.MyUnit.SetIProperty(prpc.IPT_GOLD,curGold)

	std.LogInfo("Player[",this.MyUnit.InstName,"]","Old MyGold=",oldGold,"curGold=",curGold)
}

func (this *GamePlayer) ClearAllBuff ()  {
	std.LogInfo("ClearAllBuff")
	this.MyUnit.ClearAllbuff()

	for _, unit := range this.UnitList {
		unit.ClearAllbuff()
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////升级 强化....
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) PromoteUnit (unitid int64)  {
	std.LogInfo("PromoteUnit", unitid)
	unit := this.GetUnit(unitid)
	if unit == nil {
		return
	}
	promote_info := GetPromoteRecordById(unit.UnitId)

	if promote_info == nil {
		return
	}

	if unit.IProperties[prpc.IPT_PROMOTE] >= int32(len(promote_info)) {
		return
	}

	level_info := promote_info[unit.IProperties[prpc.IPT_PROMOTE]]

	items := this.GetBagItemByTableId(level_info.ItemId)

	if items == nil || len(items) <= 0{
		std.LogInfo("cant find item, this id is ", level_info.ItemId)
		return
	}

	var num int32
	for _, item := range items {
		num += item.Stack
	}

	if num < level_info.ItemNum {
		std.LogInfo("数量不足, this id is ", level_info.ItemId, "need itemnum is ", level_info.ItemNum, "i have num is ", num)
		return
	}

	this.DelItemByTableId(level_info.ItemId, level_info.ItemNum)

	unit.Promote(level_info)

	this.session.PromoteUnitOK()
}


func (this *GamePlayer) MyUnitLevelUp()  {

	LevelUp_info := GetPromoteRecordById(1)
	std.LogInfo("MyUnitLevelUp 1", LevelUp_info)

	if LevelUp_info == nil {
		return
	}

	if this.MyUnit.Level >= int32(len(LevelUp_info)) {
		return
	}

	level_info := LevelUp_info[this.MyUnit.Level]

	this.MyUnit.Promote(level_info)
	this.MyUnit.Level = level_info.Level
	std.LogInfo("MyUnitLevelUp 2", level_info)

	for ID, info := range RoleSkillTable {
		if this.MyUnit.Level < info.OpenLv {
			continue
		}
		if this.SkillBase[ID] != 0 {
			continue
		}

		this.SkillBase[ID] = info.SKillID
	}

	//this.session.PromoteUnitOK()
}

func (this *GamePlayer)CaleMyEnergy(dt float64)  {
	myEnergy := this.MyUnit.GetIProperty(prpc.IPT_ENERGY)
	if myEnergy >= 1000 {
		return
	}
	this.EnergyTimer += dt
	if this.EnergyTimer >= 300 {
		this.SetMyEnergy(1,true)
		this.EnergyTimer = 0
	}
}

func (this *GamePlayer)SetMyEnergy(val int32,isAdd bool) {
	myEnergy := this.MyUnit.GetIProperty(prpc.IPT_ENERGY)

	if isAdd {
		if myEnergy >= 1000 {
			return
		}
		
		myEnergy += val

	} else {
		myEnergy -= val
	}

	this.MyUnit.SetIProperty(prpc.IPT_ENERGY, myEnergy)
	std.LogInfo("SetMyEnergy Val=",val,"CurmyEnergy=",myEnergy)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////技能学习
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) EquipSkill(skillinfo prpc.COM_LearnSkill) {

	std.LogInfo("EquipSkill", skillinfo)

	skill := GetRoleSkillRecordById(skillinfo.SkillID)
	if skill == nil {
		return 		//错误的skill
	}

	if this.MyUnit.Level < skill.OpenLv {
		return
	}

	real_skillid := this.SkillBase[skillinfo.SkillID]
	learnSkill := InitSkillFromTable(real_skillid)
	if learnSkill == nil{
		learnSkill = InitSkillFromTable(skill.SKillID)
	}
	if learnSkill == nil {
		return
	}

	switch skillinfo.Position {
		case 0:
			if skill.Type != 0{
				std.LogInfo("EquipSkill skill.Type wrong", 0)
				return
			}
		case 1, 2:
			if skill.Type != 1{
				std.LogInfo("EquipSkill skill.Type wrong", 1, 2)
				return
			}
		case 3:
			if skill.Type != 2 {
				std.LogInfo("EquipSkill skill.Type wrong", 3)
				return
			}
	}

	this.MyUnit.Skill[skillinfo.Position] = learnSkill

	var idx int32 = -1
	for index, skill_ := range this.MyUnit.Skill {
		std.LogInfo("skill", skill_, &skill_)
		if skill_ == nil {
			continue
		}
		if skill_.SkillID == learnSkill.SkillID{
			if index == skillinfo.Position {
				continue
			}
			idx = index
			break
		}
	}

	if idx != -1 {
		this.MyUnit.Skill[idx] = nil
	}

	std.LogInfo("skillall", this.MyUnit.Skill)

	this.session.EquipSkillOK(skillinfo.Position, learnSkill.SkillID)
	std.LogInfo("EquipSkillOK", skillinfo.Position, learnSkill.SkillID)

	//如果是被动技能 需要修改buff

	return
}

func (this * GamePlayer)SkillUpdate(skillindex int32, skillId int32) {
	std.LogInfo("SkillUpdate", skillindex, skillId)
	if this.SkillBase[skillindex] == 0 {
		return
	}

	updateInfo := GetRoleSkillUpdateRecordById(skillId)

	if updateInfo == nil{
		return
	}

	//检测道具是否拥有
	curCopper := this.MyUnit.GetIProperty(prpc.IPT_COPPER)

	if curCopper < updateInfo.NeedMoney {
		return
	}

	curCopper -= updateInfo.NeedMoney

	items := this.GetBagItemByTableId(updateInfo.NeedItem)

	if len(items) <= 0 {
		return
	}

	var curnum int32
	for _, item := range items{
		curnum += item.Stack
	}

	if curnum < updateInfo.NeedNum {
		return
	}

	new_skill := InitSkillFromTable(updateInfo.NextID)

	if new_skill == nil {
		return
	}

	this.MyUnit.SetIProperty(prpc.IPT_COPPER, curCopper)
	this.DelItemByTableId(updateInfo.NeedItem, updateInfo.NeedNum)

	var skillpos int32 = -1
	for idx, skill := range this.MyUnit.Skill {
		if skill == nil {
			continue
		}
		if skill.SkillID == skillId {
			skillpos = idx
		}
	}

	if skillpos != -1 {
		this.MyUnit.Skill[skillpos] = new_skill
	}
	this.SkillBase[skillindex] = updateInfo.NextID

	this.session.SkillUpdateOK(skillindex, updateInfo.NextID, skillpos)
	std.LogInfo("SkillUpdateOK", skillindex, updateInfo.NextID, skillpos)
	std.LogInfo("SkillUpdateOK 1", this.SkillBase)

}

func (this * GamePlayer)SkillUpdate_equip(position int32, skillId int32) {
	updateInfo := GetRoleSkillUpdateRecordById(skillId)

	if updateInfo == nil{
		return
	}

	//检测道具是否拥有

	new_skill := InitSkillFromTable(updateInfo.NextID)

	if new_skill == nil {
		return
	}

	var change int32
	for index, skill_id := range this.SkillBase {
		if skill_id == skillId {
			change = index
		}
	}

	this.MyUnit.Skill[position] = new_skill


	if change != 0{
		this.SkillBase[change] = updateInfo.NextID
	}

	//这里需要buff更新一下

}

func (this * GamePlayer)CheckSkillBase() {
	for ID, info := range RoleSkillTable {
		if this.MyUnit.Level < info.OpenLv {
			continue
		}
		if this.SkillBase[ID] != 0 {
			continue
		}

		this.SkillBase[ID] = info.SKillID
	}

}

func (this *GamePlayer)BuyShopItem(shopId int32)  {
	shopData := GetShopDataById(shopId)
	if shopData == nil {
		return
	}

	if shopData.CurrenciesKind == prpc.IPT_GOLD{

		myGold := this.MyUnit.GetIProperty(prpc.IPT_GOLD)
		if myGold < shopData.Price {
			return
		}
		this.AddGold(-shopData.Price)
		std.LogInfo("Player[",this.MyUnit.InstName,"]","BuyShopItem ShopId=",shopId,"Shoping Spend=",shopData.Price)
	}

	if shopData.ShopType == prpc.SHT_BuyCopper {
		if shopData.Copper > 0 {
			this.AddCopper(shopData.Copper)
			if shopData.CardPondId != 0 {
				this.OpenTreasureBox(shopData.CardPondId)
			}
		}else {
			std.LogInfo("Player[",this.MyUnit.InstName,"]","shopData.Copper A wrong number","BuyShopItem ShopId=",shopId)
		}
	}
}

func (this *GamePlayer)OpenTreasureBox(pondId int32) bool {
	data := GetCardPondTableDataById(pondId)
	if data==nil {
		std.LogInfo("Player[",this.MyUnit.InstName,"]","OpenTreasureBox GetCardPondTableDataById Can Not Find pondId=",pondId)
		return false
	}

	var items []int32

	greenItems 		:= data.GetGreenCardItems()
	buleItems  		:= data.GetBlueCardItems()
	purplenItems 	:= data.GetPurpleCardItems()
	orangeItems 	:= data.GetOrangeCardItems()

	for _,itemId := range greenItems{
		items = append(items,itemId)
		//std.LogInfo("Player[",this.MyUnit.InstName,"]","OpenTreasureBox GreenItem itemId",itemId)
	}
	for _,itemId := range buleItems{
		items = append(items,itemId)
		//std.LogInfo("Player[",this.MyUnit.InstName,"]","OpenTreasureBox BuleItem itemId",itemId)
	}
	for _,itemId := range purplenItems{
		items = append(items,itemId)
		//std.LogInfo("Player[",this.MyUnit.InstName,"]","OpenTreasureBox PurlenItem itemId",itemId)
	}
	for _,itemId := range orangeItems{
		items = append(items,itemId)
		//std.LogInfo("Player[",this.MyUnit.InstName,"]","OpenTreasureBox OrangeItem itemId",itemId)
	}

	itemInsts := []prpc.COM_ItemInst{}
	for _,itemId := range items {
		var isHave bool = false
		for i:=0;i<len(itemInsts) ;i++  {
			if itemInsts[i].ItemId == itemId {
				itemInsts[i].Stack++
				isHave = true
				break
			}
		}
		if !isHave {
			item := prpc.COM_ItemInst{}
			item.ItemId = itemId
			item.Stack	= 1
			itemInsts = append(itemInsts,item)
		}
	}

	for _,item := range itemInsts{
		this.AddBagItemByItemId(item.ItemId,item.Stack)
		std.LogInfo("Player[",this.MyUnit.InstName,"]","OpenTreasureBox AddItem ID=",item.ItemId,"Num=",item.Stack)
	}

	if this.session != nil {
		this.session.BuyShopItemOK(itemInsts)
	}
	
	return true
}

func (this *GamePlayer) Logout(){

	std.LogInfo("Logout","PlayerName=",this.MyUnit.InstName)
	//清理战斗信息
	this.LeftBattle_strong()
	
	//

	this.Save()

	//
	RemovePlayer(this)
}

func (this* GamePlayer)Save(){
	std.LogInfo("SAVE ")
	UpdatePlayer(this.GetPlayerSGE())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestPlayer() {
	P1 := CreatePlayer(1, "testPlayer")
	P1.BuyShopItem(1000)
	P1.BuyShopItem(1001)
	P1.BuyShopItem(1002)
	for i:=0;i<len(P1.BagItems) ;i++  {
		std.LogInfo("BagItems ItemId=",P1.BagItems[i].ItemId,"ItemNum=",P1.BagItems[i].Stack)
	}
	P1.TestItem()
}

func (this *GamePlayer)TestItem()  {
	//for i := 1; i < 9 ; i++ {	//测试用
	//	this.AddBagItemByItemId(int32(i), 10)
	//}
	//this.AddBagItemByItemId(5000,2000)
	this.AddCopper(10000000)
	this.AddGold(10000)
	this.GiveDrop(1000)

	for _,u := range this.UnitList{
		std.LogInfo("Myself Unit InstId",u.InstId,"InstName",u.InstName)
	}
}