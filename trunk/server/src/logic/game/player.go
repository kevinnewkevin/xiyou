package game

import (
	"errors"
	"fmt"
	"logic/prpc"
	"sync"
	"strings"
	"strconv"
)

const (
	unitGroupMax		= 5			//卡组上限
	onceUnitGroupMax 	= 10		//每组卡片上限
	bagMaxGrid			= 200
)

type GamePlayer struct {
	sync.Locker
	session        			*Session    //链接
	MyUnit         			*GameUnit   //自己的卡片
	UnitList       			[]*GameUnit //拥有的卡片
	BattleUnitList 			[]int64     //默认出战卡片
	DefaultUnitGroup		int			//默认战斗卡片组
	BattleUnitGroup			int32		//战斗卡片组
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

	//Bag
	BagItems				[]*prpc.COM_ItemInst
}

var (
	PlayerStore		[]*GamePlayer
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

func SetDefaultUnits(cards string) {
	s1 := strings.Split(cards, ",")
	for _, c := range s1{
		e_id, _ := strconv.Atoi(c)
		DefaultUnits = append(DefaultUnits, int32(e_id))
	}
}

func (this *GamePlayer) SetSession(session *Session) {
	this.session = session
}

func CreatePlayer(tid int32, name string) *GamePlayer {
	p := GamePlayer{}
	p.MyUnit = p.NewGameUnit(tid)
	p.MyUnit.InstName = name
	//来两个默认的小兵
	//p.UnitList = append(p.UnitList, p.NewGameUnit(4))
	//p.UnitList = append(p.UnitList, p.NewGameUnit(5))
	//p.UnitList = append(p.UnitList, p.NewGameUnit(6))
	//p.UnitList = append(p.UnitList, p.NewGameUnit(7))
	//p.UnitList = append(p.UnitList, p.NewGameUnit(8))
	//p.UnitList = append(p.UnitList, p.NewGameUnit(9))
	//p.UnitList = append(p.UnitList, p.NewGameUnit(10))
	for _, e_id := range DefaultUnits {
		p.UnitList = append(p.UnitList, p.NewGameUnit(e_id))
	}
	p.DefaultUnitGroup = 1
	p.TianTiVal	= 0
	PlayerStore = append(PlayerStore, &p)
	p.InitUnitGroup()

	for _,u := range p.UnitList{
		fmt.Println("Myself Unit InstId",u.InstId,"InstName",u.InstName)
	}
	
	return &p

}

func (this *GamePlayer) NewGameUnit(tid int32) *GameUnit {
	unit := CreateUnitFromTable(tid)
	unit.Owner = this
	chapterids := GetUnitChapterById(tid)
	for i:=0;i<len(chapterids) ;i++  {
		OpenChapter(this,chapterids[i])
	}
	return unit
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

	//
	this.SyncBag()

	return p
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//技能相关
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) StudySkill(UnitID int64, skillpos int32, skillid int32) error {
	if skillpos >= 2 {
		fmt.Println("技能位置錯誤")
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
	fmt.Println("JoinBattle", battlePlayerList)
}

func (this *GamePlayer) JoinBattlePvE(bigGuanqia int32, SmallGuanqia int32) {
	//this.Lock()
	//defer this.Unlock()

	CreatePvE(battlePlayerList[0], 1)

	fmt.Println("JoinBattlePvE", battlePlayerList)
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
	fmt.Println("SetBattleUnit ", this.BattleUnitList)
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
	fmt.Println("SetBattleUnitGroup InstId",instId,"GroupID",groupId,"IsBattle",isBattle)
	if isBattle {
		addError := this.AddUnitToGroup(instId,groupId)
		if addError != 0{
			fmt.Println("AddUnitToGroup Error",addError)
		}
	}else {
		addError := this.RemoveUnitToGroup(instId,groupId)
		if addError != 0{
			fmt.Println("RemoveUnitToGroup Error",addError)
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

func (this *GamePlayer) SetupBattle(pos []prpc.COM_BattlePosition) error { //卡牌上阵	每次回合之前
	//this.Lock()
	//defer this.Unlock()
	//fmt.Println("SetupBattle", pos)
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
	fmt.Println("SetupBattle 1 ", battleRoom.Units)
	battleRoom.SetupPosition(this, pos)
	fmt.Println("SetupBattle 2 ", battleRoom.Units)

	this.session.SetupBattleOK()

	//battleRoom.Update()

	return nil
}

func (this *GamePlayer) SetProprty(battleid int64, camp int) {
	this.BattleId = battleid
	this.BattleCamp = camp
	this.BattlePoint = 1
	this.MyUnit.ResetBattle(camp, true, battleid)

	for _, u := range this.UnitList {
		u.ResetBattle(camp, false, battleid)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer)SyncBag()  {
//////给客户端测试的东西
	for i:=0;i<len(ItemTableData) ;i++  {
		if i==1 {
			this.AddBagItemByItemId(int32(i+1),1050)
		}else {
			this.AddBagItemByItemId(int32(i+1),10)
		}
	}
	for _,item := range this.BagItems{
		if item.ItemId== 2 {
			this.DelItemByTableId(2,1000)
		}

		fmt.Println("ItemInst  ItemInstId=",item.InstId,"ItemId=",item.ItemId,"itemStack=",item.Stack_,"Bag len",len(this.BagItems))
	}
/////
	items := []prpc.COM_ItemInst{}

	for _,itemInst := range this.BagItems {
		items = append(items,*itemInst)
	}

	for _,item := range items{
		fmt.Println("To Client Item TableId=",item.ItemId,"Stack=",item.Stack_,"InstId=",item.InstId)
	}

	if this.session != nil {
		this.session.InitBagItems(items)
	}
}

func (this *GamePlayer)AddBagItemByItemId(itemId int32,itemCount int32)  {
	itemData := GetItemTableDataById(itemId)
	if itemData==nil {
		fmt.Println("ItemTable Not Find This ItemId=",itemId)
		return
	}
	for _,itemInst := range this.BagItems{
		if itemInst.ItemId == itemId {
			itemInst.Stack_ += itemCount
			if itemInst.Stack_ > itemData.MaxCount {
				itemCount = itemInst.Stack_ - itemData.MaxCount
				itemInst.Stack_ = itemData.MaxCount
			}
			//updata bag itemInst
			if this.session != nil {
				this.session.UpdateBagItem(*itemInst)
			}
			break;
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
		fmt.Println("Not Find Item In The Bag",instid)
		return
	}
	for i:=0;i<len(this.BagItems) ;i++  {
		if this.BagItems[i] == nil{
			continue
		}
		if this.BagItems[i].InstId == instid {
			if this.BagItems[i].Stack_ > stack {
				this.BagItems[i].Stack_ -= stack
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

	itemData := GetItemTableDataById(itemInst.ItemId)
	if itemData==nil {
		return
	}

	v := []interface{}{0}
	r := []interface{}{0}


	_L.CallFuncEx(itemData.GloAction, v, &r)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestPlayer() {
	P1 := CreatePlayer(1, "testPlayer")
	fmt.Println(len(ItemTableData))
	for i:=0;i<len(ItemTableData) ;i++  {
		if i==1 {
			P1.AddBagItemByItemId(int32(i+1),1050)
		}else {
			P1.AddBagItemByItemId(int32(i+1),10)
		}
	}
	for _,item := range P1.BagItems{
		if item.ItemId== 2 {
			P1.DelItemByTableId(2,1000)
		}
		
		fmt.Println("ItemInst  ItemInstId=",item.InstId,"ItemId=",item.ItemId,"itemStack=",item.Stack_,"Bag len",len(P1.BagItems))
	}
	//CreatePvE(P, 1)
}


func (this *GamePlayer) ClearAllBuff ()  {
	fmt.Println("ClearAllBuff")
	this.MyUnit.ClearAllbuff()

	for _, unit := range this.UnitList {
		unit.ClearAllbuff()
	}
}
