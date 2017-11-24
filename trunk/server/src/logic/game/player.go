package game

import (
	"errors"
	"jimny/logs"
	"logic/prpc"
	"strconv"
	"strings"
	"time"
)

const (
	unitGroupMax     = 5  //卡组上限
	onceUnitGroupMax = 10 //每组卡片上限
	bagMaxGrid       = 200
)

type GamePlayer struct {
	session        *Session //链接
	PlayerId       int64    //角色ID
	Username       string   //账户名
	LoginTime      int64
	LogoutTime     int64
	MyUnit         *GameUnit   //自己的卡片
	UnitList       []*GameUnit //拥有的卡片
	BattleUnitList []int64     //默认出战卡片
	//DefaultUnitGroup		int			//默认战斗卡片组
	BattleUnitGroup int32 //战斗卡片组
	UnitGroup       []*prpc.COM_UnitGroup
	//战斗相关辅助信息
	BattleId    int64   //所在房间编号
	BattleCamp  int     //阵营 //prpc.CompType
	IsActive    bool    //是否激活
	KillUnits   []int32 //杀掉的怪物
	MyDeathNum  int32   //战斗中自身死亡数量
	BattlePoint int32   //战斗點數

	//story chapter
	ChapterPondId int32
	ChapterID int32 //正在进行的关卡
	Chapters  []*prpc.COM_Chapter

	TianTiVal        int32
	TianTiRank       int32
	FriendTianTiRank int32
	EnergyTimer      float64
	//Bag
	BagItems []*prpc.COM_ItemInst
	GenItemMaxGuid		int64
	//经验
	Exp int32
	//主角可学习技能
	SkillBase map[int32]int32
	//黑市
	BlackMarketData *prpc.COM_BlackMarket

	//状态标记 时间戳
	LockTime int64

	//好友排行榜
	FriendTop []prpc.COM_TopUnit

	//新手引導步驟
	Guide uint64

	//好友
	Friends []*prpc.COM_Friend
	Enemys  []*prpc.COM_Friend

	GuildId			int32
	AssistantId		int32
	AssistantCreateTime		int64
}

var (
	PlayerStore       		= []*GamePlayer{}
	InstIdPlayerStore		= map[int64]*GamePlayer{}
	InstNamePlayerStore     = map[string]*GamePlayer{}
	DefaultUnits      []int32
	DefaultBattleUnit int32
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//角色创建
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func FindPlayerByInstId(instid int64) *GamePlayer {
	for _, p := range PlayerStore {
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
	for _, p := range PlayerStore {
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
	for _, p := range PlayerStore {
		if p == nil {
			continue
		}
		if p.MyUnit.InstName == instName {
			return p
		}
	}

	return nil
}

func RemovePlayerByInstName(instName string) {
	for i := 0; i < len(PlayerStore); i++ {
		if PlayerStore[i] == nil {
			continue
		}
		if PlayerStore[i].MyUnit.InstName == instName {
			PlayerStore = append(PlayerStore[:i], PlayerStore[i+1:]...)
		}
	}
	logs.Info("RemovePlayerByInstName Name= %s PlayerStore len= %d", instName, len(PlayerStore))
}

func RemovePlayer(player *GamePlayer) {
	if player == nil {
		return
	}
	RemovePlayerByInstName(player.MyUnit.InstName)
	delete(InstNamePlayerStore, player.MyUnit.InstName)
	delete(InstIdPlayerStore, player.MyUnit.InstId)
}

func SetDefaultUnits(cards string) {
	s1 := strings.Split(cards, ",")
	for _, c := range s1 {
		e_id, _ := strconv.Atoi(c)
		DefaultUnits = append(DefaultUnits, int32(e_id))
	}
}

func SetDefaultBattleUnit(card int32) {
	DefaultBattleUnit = card

	logs.Info("SetDefaultBattleUnit, ", DefaultBattleUnit)
}

func PlayerTick(dt float64) {
	for _, p := range PlayerStore {
		if p == nil {
			continue
		}
		p.CaleMyEnergy(dt)
	}
}

func OnlinePlayerPassZeroHour() {
	for _, p := range PlayerStore {
		if p == nil {
			continue
		}
		p.PassZeroHour()
	}
}

func (this *GamePlayer) SetSession(session *Session) {
	this.session = session
}

func CreatePlayer(tid int32, name string, username string) *GamePlayer {

	p := GamePlayer{}
	p.Username = username
	p.PlayerId = GenPlayerInstId()
	p.MyUnit = p.NewGameUnit(tid)

	p.MyUnit.InstName = name
	p.MyUnit.IsMain = true
	p.Exp = 0
	p.Guide = 0

	p.InitUnitGroup()
	logs.Debug("createplayer ", DefaultUnits)
	for _, e_id := range DefaultUnits {
		u := p.NewGameUnit(e_id)
		if e_id == DefaultBattleUnit {
			p.SetBattleUnitGroup(u.InstId, 1, true)
		}
	}
	//p.DefaultUnitGroup = 1
	p.TianTiVal = 0
	p.TianTiRank = -1
	p.FriendTianTiRank = -1
	p.GenItemMaxGuid = 1
	PlayerStore = append(PlayerStore, &p)
	p.Friends = []*prpc.COM_Friend{}
	p.Enemys = []*prpc.COM_Friend{}
	p.BattleUnitGroup = 1

	for _, u := range p.UnitList {
		logs.Info("Myself Unit InstId %d InstName %s", u.InstId, u.InstName)
	}
	p.SkillBase = map[int32]int32{}

	for ID, info := range RoleSkillTable {
		if p.MyUnit.Level >= info.OpenLv {
			p.SkillBase[ID] = info.SKillID
		} else {
			p.SkillBase[ID] = 0
		}
	}

	p.ChapterPondId = 1

	chapters := GetGlobalString("C_DefaultChatper");

	s1 := strings.Split(chapters, ",")
	for _, c := range s1 {
		e_id, _ := strconv.Atoi(c)
		p.OpenChapter(int32(e_id))
	}

	p.FriendTop = []prpc.COM_TopUnit{}

	return &p

}

func (this *GamePlayer) NewGameUnit(tid int32) *GameUnit {
	unit := CreateUnitFromTable(tid)
	if unit == nil {
		return nil
	}
	unit.Owner = this

	if tid == 1 || tid == 2 { //主角卡不要放在牌库中
		return unit
	}

	this.UnitList = append(this.UnitList, unit)

	InsertUnit(this.PlayerId,unit.GetUnitCOM())

	return unit
}

func (this *GamePlayer) SetPlayerCOM(p *prpc.COM_Player) {
	this.MyUnit = &GameUnit{}
	this.MyUnit.Owner = this
	this.MyUnit.SetUnitCOM(&p.Unit)
	this.MyUnit.InstName = p.Name

	for _, u := range p.Employees {
		unit := GameUnit{}
		unit.SetUnitCOM(&u)
		unit.Owner = this
		this.UnitList = append(this.UnitList, &unit)
	}
	for i, _ := range p.Chapters {
		this.Chapters = append(this.Chapters, &p.Chapters[i])
	}

	for i, _ := range p.UnitGroup {
		this.UnitGroup = append(this.UnitGroup, &p.UnitGroup[i])
	}
	this.TianTiVal = p.TianTiVal
	this.TianTiRank = p.TianTiRank
	this.FriendTianTiRank = p.FriendTianTiRank
	this.Guide = p.Guide
	this.BattleUnitGroup = p.BattleUnitGroup

	this.AssistantCreateTime = p.AssistantCreateTime
	this.SkillBase = map[int32]int32{}
	for _, skb := range p.SkillBase {
		this.SkillBase[skb.SkillId] = skb.SkillId
	}

	for i := range p.Friends {
		this.Friends = append(this.Friends, &p.Friends[i])
	}

	for i := range p.Enemys {
		this.Enemys = append(this.Enemys, &p.Enemys[i])
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
	for _, ug := range this.UnitGroup {
		p.UnitGroup = append(p.UnitGroup, *ug)
	}

	for _, fr := range this.Friends {
		p.Friends = append(p.Friends, *fr)
	}

	for _, en := range this.Enemys {
		p.Enemys = append(p.Enemys, *en)
	}
	p.TianTiVal = this.TianTiVal
	p.TianTiRank = this.TianTiRank
	p.FriendTianTiRank = this.FriendTianTiRank
	p.Guide = this.Guide
	p.BattleUnitGroup = this.BattleUnitGroup

	p.AssistantCreateTime = this.AssistantCreateTime
	for index, skillid := range this.SkillBase {
		skillbase := prpc.COM_SkillBase{}
		skillbase.SkillIdx = index
		skillbase.SkillId = skillid

		p.SkillBase = append(p.SkillBase, skillbase)
	}

	return p
}

func (this *GamePlayer) SetPlayerSGE(p prpc.SGE_DBPlayer) {

	this.SetPlayerCOM(&p.COM_Player)
	this.PlayerId 	= p.PlayerId
	this.Username 	= p.Username
	this.LoginTime 	= p.LoginTime
	this.LogoutTime = p.LogoutTime
	this.GuildId	= p.GuildId
	this.AssistantId	= p.AssistantId
	this.GenItemMaxGuid = p.GenItemMaxGuid
	this.ChapterPondId = p.ChapterPondId
	for i := range p.BagItemList {
		this.BagItems = append(this.BagItems, &p.BagItemList[i])
	}

	this.BlackMarketData = &p.BlackMarketData
}

func (this *GamePlayer) GetPlayerSGE() prpc.SGE_DBPlayer {
	items := []prpc.COM_ItemInst{}
	for _, a := range this.BagItems {
		items = append(items, *a)
	}

	data := prpc.SGE_DBPlayer{}
	data.COM_Player = this.GetPlayerCOM()
	data.PlayerId = this.PlayerId
	data.Username = this.Username
	data.LoginTime = this.LoginTime
	data.LogoutTime = this.LogoutTime
	data.GuildId	= this.GuildId
	data.AssistantId = this.AssistantId
	data.GenItemMaxGuid = this.GenItemMaxGuid
	data.ChapterPondId = this.ChapterPondId
	data.BagItemList = items

	if this.BlackMarketData != nil {
		data.BlackMarketData = *this.BlackMarketData
	}

	return data
}



func (this *GamePlayer) IsBattle() bool {
	if this.BattleId == 0 {
		return false
	}
	return true
}

func (this *GamePlayer) PassZeroHour() {
	if this.BlackMarketData != nil {
		this.BlackMarketData.RefreshNum = int32(GetGlobalInt("C_BlackMarkteRefreshNum"))
		this.session.SycnBlackMarkte(*this.BlackMarketData)
	}
}

func (this *GamePlayer) PlayerLogin() {
	this.LoginTime = time.Now().Unix()

	loginDT := time.Unix(this.LoginTime, 0)
	logoutDT := time.Unix(this.LogoutTime, 0)

	if loginDT.Day() != logoutDT.Day() || loginDT.Month() != logoutDT.Month() || loginDT.Year() != logoutDT.Year() {
		this.PassZeroHour()
	}

	this.SyncBag()
	if this.BlackMarketData != nil {
		this.session.SycnBlackMarkte(*this.BlackMarketData)
	} else {
		this.InitMyBlackMarket()
	}

	this.RefreshFriendTopList()

	pGuild := FindGuildByPlayerId(this.MyUnit.InstId)
	if pGuild != nil{
		pGuild.GuildMemberOnLine(this)
	}
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

func (this *GamePlayer) HasUnitByTableId(tableId int32) bool {
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
		logs.Debug("技能位置錯誤")
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
	logs.Info("JoinBattle %s", battlePlayerList)
}

func (this *GamePlayer) JoinBattlePvE(bigGuanqia int32, SmallGuanqia int32) {
	//this.Lock()
	//defer this.Unlock()

	CreatePvE(battlePlayerList[0], 1)

	logs.Info("JoinBattlePvE %s", battlePlayerList)
}

func (this *GamePlayer) LeftBattle_strong() {
	battle := FindBattle(this.BattleId)
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
	logs.Info("SetBattleUnit  %s", this.BattleUnitList)
}

func (this *GamePlayer) InitUnitGroup() {
	for i := 0; i < unitGroupMax; i++ {
		unitgroup := prpc.COM_UnitGroup{}
		unitgroup.GroupId = int32(i + 1)
		this.UnitGroup = append(this.UnitGroup, &unitgroup)
	}
}

func (this *GamePlayer) SetBattleUnitGroup(instId int64, groupId int32, isBattle bool) {
	if this == nil {
		return
	}
	if this.GetUnit(instId) == nil {
		return
	}
	logs.Info("SetBattleUnitGroup InstId", instId, "GroupID", groupId, "IsBattle", isBattle)
	if isBattle {
		addError := this.AddUnitToGroup(instId, groupId)
		if addError != 0 {
			logs.Info("AddUnitToGroup Error", addError)
		}
	} else {
		addError := this.RemoveUnitToGroup(instId, groupId)
		if addError != 0 {
			logs.Info("RemoveUnitToGroup Error", addError)
		}
	}
}

func (this *GamePlayer) GetUnitGroupById(groupId int32) *prpc.COM_UnitGroup {
	for _, g := range this.UnitGroup {
		if g == nil {
			continue
		}
		if g.GroupId == groupId {
			return g
		}
	}
	return nil
}

func (this *GamePlayer) AddUnitToGroup(instId int64, groupId int32) int {
	group := this.GetUnitGroupById(groupId)
	if group == nil {
		return 1
	}

	if len(group.UnitList) >= onceUnitGroupMax {
		return 2
	}

	card := this.GetUnit(instId)
	if card == nil {
		return 3
	}

	for _, unit := range group.UnitList {
		tmp := this.GetUnit(unit)
		if tmp == nil {
			continue
		}
		if card.UnitId == tmp.UnitId {
			return 4
		}
	}

	group.UnitList = append(group.UnitList, card.InstId)

	return 0
}

func (this *GamePlayer) RemoveUnitToGroup(instId int64, groupId int32) int {
	group := this.GetUnitGroupById(groupId)
	if group == nil {
		return 1
	}

	card := this.GetUnit(instId)
	if card == nil {
		return 2
	}

	index := 100

	for i := 0; i < len(group.UnitList); i++ {
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

func (this *GamePlayer) DeleteUnitGroup(groupId int32) {
	for i := 0; i < len(this.UnitGroup); i++ {
		if this.UnitGroup[i].GroupId == groupId {
			unitgroup := prpc.COM_UnitGroup{}
			this.UnitGroup[i] = &unitgroup
		}
	}
}

func (this *GamePlayer) SetupBattle(pos []prpc.COM_BattlePosition, skillid int32) error { //卡牌上阵	每次回合之前
	//this.Lock()
	//defer this.Unlock()
	//logs.Info("SetupBattle", pos)

	logs.Info("SetupBattle")
	if this.IsActive {
		logs.Info("this.IsActive")
		return nil
	}

	for _, p := range pos {
		//if this.GetBattleUnit(int64(p.InstId)) == nil {
		//	return nil //错误消息
		//}
		if p.Position >= prpc.BP_MAX || p.Position < prpc.BP_RED_1 {
			logs.Info(" p.Position >= prpc.BP_MAX || p.Position < prpc.BP_RED_1", p.Position)
			return nil //错误消息 //检测缺失 阵营与位置关系
		}
	}

	battleRoom := FindBattle(this.BattleId)

	if battleRoom == nil {
		//错误消息
		logs.Info("battleRoom")
		return nil
	}
	logs.Debug("SetupBattle 1 ", battleRoom.Units)
	battleRoom.SetupPosition(this, pos, skillid)
	logs.Debug("SetupBattle 2 ", battleRoom.Units)

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

func (this *GamePlayer) SyncBag() {
	items := []prpc.COM_ItemInst{}

	for _, itemInst := range this.BagItems {
		items = append(items, *itemInst)
	}

	for _, item := range items {
		logs.Info("To Client Item TableId=", item.ItemId, " Stack=", item.Stack, " InstId=", item.InstId)
	}

	if this.session != nil {
		this.session.InitBagItems(items)
	}
}

func (this *GamePlayer) AddBagItemByItemId(itemId int32, itemCount int32) {
	itemData := GetItemTableDataById(itemId)
	if itemData == nil {
		logs.Info("ItemTable Not Find This ItemId=", itemId)
		return
	}
	for _, itemInst := range this.BagItems {
		if itemInst.ItemId == itemId {
			itemInst.Stack += itemCount
			if itemInst.Stack > itemData.MaxCount {
				itemCount = itemInst.Stack - itemData.MaxCount
				itemInst.Stack = itemData.MaxCount
			} else {
				itemCount = 0
			}

			//updata bag itemInst
			if this.session != nil {
				this.session.UpdateBagItem(*itemInst)
			}
			if itemCount == 0 {
				break
			}
		}
	}
	if itemCount > 0 {
		newItems := this.GenItemInst(itemId, itemCount)
		if len(newItems) == 0 {
			return
		}
		if len(this.BagItems)+len(newItems) >= bagMaxGrid {
			//newItems To Mall
			return
		}

		for _, newItem := range newItems {
			if newItem == nil {
				continue
			}
			this.BagItems = append(this.BagItems, newItem)
			//add newItem
			if this.session != nil {
				this.session.AddBagItem(*newItem)
			}
		}
	}
}

func (this *GamePlayer) DelItemByInstId(instid int64, stack int32) {
	itemInst := this.GetBagItemByInstId(instid)
	if itemInst == nil {
		logs.Info("Not Find Item In The Bag", instid)
		return
	}
	for i := 0; i < len(this.BagItems); i++ {
		if this.BagItems[i] == nil {
			continue
		}
		if this.BagItems[i].InstId == instid {
			if this.BagItems[i].Stack > stack {
				this.BagItems[i].Stack -= stack
				//updata item
				if this.session != nil {
					this.session.UpdateBagItem(*itemInst)
				}
			} else {
				this.BagItems = append(this.BagItems[:i], this.BagItems[i+1:]...)
				//del item
				if this.session != nil {
					this.session.DeleteItemOK(instid)
				}
			}
		}
	}
}

func (this *GamePlayer) DelItemByTableId(tableId int32, delNum int32) {
	items := this.GetBagItemByTableId(tableId)
	if len(items) == 0 {
		logs.Info("Can Not Find Item In Bag By TableId =", tableId)
		return
	}
	for _, item := range items {
		if item.Stack > delNum {
			item.Stack -= delNum
			if this.session != nil {
				this.session.UpdateBagItem(*item)
			}
		} else {
			delNum -= item.Stack
			this.DelItemByInstId(item.InstId, item.Stack)
		}
	}
}

func (this *GamePlayer) GetBagItemByInstId(instId int64) *prpc.COM_ItemInst {
	for _, itemInst := range this.BagItems {
		if itemInst == nil {
			continue
		}
		if itemInst.InstId == instId {
			return itemInst
		}
	}
	return nil
}

func (this *GamePlayer) GetBagItemByTableId(itemid int32) []*prpc.COM_ItemInst {
	items := []*prpc.COM_ItemInst{}
	for _, itemInst := range this.BagItems {
		if itemInst == nil {
			continue
		}
		if itemInst.ItemId == itemid {
			items = append(items, itemInst)
		}
	}
	return items
}

func (this *GamePlayer)GetBagItemNumByTableId(itemId int32) int32 {
	var itemNum int32 = 0
	for _, itemInst := range this.BagItems {
		if itemInst == nil {
			continue
		}
		if itemInst.ItemId == itemId {
			itemNum += itemInst.Stack
		}
	}
	return itemNum
}

func (this *GamePlayer)GetBagItemNumByInstId(instId int64) int32 {
	var itemNum int32 = 0
	for _, itemInst := range this.BagItems {
		if itemInst == nil {
			continue
		}
		if itemInst.InstId == instId {
			itemNum += itemInst.Stack
		}
	}
	return itemNum
}

func (this *GamePlayer) UseItem(instId int64, useNum int32) {
	itemInst := this.GetBagItemByInstId(instId)
	if itemInst == nil {
		return
	}

	if itemInst.Stack < useNum {
		return
	}

	itemData := GetItemTableDataById(itemInst.ItemId)
	if itemData == nil {
		return
	}

	v := []interface{}{int64(this.MyUnit.InstId), int64(itemData.ItemId)}
	r := []interface{}{""}

	usestack := 0

	for i := 0; i < int(useNum); i++ {
		CallLuaFunc(itemData.GloAction, v, &r)
		errorNo := r[0]
		if errorNo != "" {
			errorId := prpc.ToId_ErrorNo(errorNo.(string))
			if this.session != nil {
				this.session.ErrorMessage(errorId)
			}
			logs.Info("useItem errorId=", errorId, "itemId=", itemData.ItemId)
			break
		}
		usestack++
	}
	if usestack != 0 {
		this.DelItemByInstId(instId, int32(usestack))
	}
	logs.Info("123123123123", r, len(r), r[0], usestack)
}

func (this *GamePlayer) GiveDrop(dropId int32) {
	drop := GetDropById(dropId)
	if drop == nil {
		logs.Info("Can Not Find Drop By DropId=", dropId)
		return
	}
	if drop.Exp != 0 {
		this.AddExp(drop.Exp)
	}
	if drop.Money != 0 {
		this.AddCopper(drop.Money)
	}
	if len(drop.Items) != 0 {
		for _, item := range drop.Items {
			this.AddBagItemByItemId(item.ItemId, item.ItemNum)
			logs.Info("PlayerName=", this.MyUnit.InstName, "GiveDrop AddItem ItemId=", item.ItemId, "ItemNum=", item.ItemNum)
		}
	}
	if drop.Hero != 0 {
		if this.HasUnitByTableId(drop.Hero) {
			//有这个卡就不给了
			logs.Info("PlayerName=", this.MyUnit.InstName, "GiveDrop AddUnit Have not to UnitId=", drop.Hero)
		} else {
			unit := this.NewGameUnit(drop.Hero)
			if unit != nil {
				logs.Info("PlayerName=", this.MyUnit.InstName, "GiveDrop AddUnit OK UnitId=", drop.Hero)
				temp := unit.GetUnitCOM()
				if this.session != nil {
					this.session.AddNewUnit(temp)
				}
			}
		}
	}
}

func (this *GamePlayer) AddExp(val int32) {
	curExp := this.MyUnit.GetIProperty(prpc.IPT_EXPERIENCE)
	curExp += val
	if curExp < 0 {
		curExp = 0
	}
	//在这里加上对于经验值和等级的判断

	curExp = this.MyUnit.CheckExp(curExp)

	this.MyUnit.SetIProperty(prpc.IPT_EXPERIENCE, curExp)

	logs.Info("append EXP", val, "all EXP", curExp)
}

func (this *GamePlayer) AddCopper(val int32) {
	oldCopper := this.MyUnit.GetIProperty(prpc.IPT_COPPER)
	curCopper := oldCopper + val
	if curCopper < 0 {
		curCopper = 0
	}
	if curCopper > CopperMax {
		curCopper = CopperMax
	}
	this.MyUnit.SetIProperty(prpc.IPT_COPPER, curCopper)

	logs.Debug("Player[", this.MyUnit.InstName, "]", "Old Copper=", oldCopper, "curCopper=", curCopper)
}

func (this *GamePlayer) AddGold(val int32) {
	oldGold := this.MyUnit.GetIProperty(prpc.IPT_GOLD)
	curGold := oldGold + val
	if curGold < 0 {
		curGold = 0
	}
	if curGold > CopperMax {
		curGold = CopperMax
	}
	this.MyUnit.SetIProperty(prpc.IPT_GOLD, curGold)

	logs.Info("Player[", this.MyUnit.InstName, "]", "Old MyGold=", oldGold, "curGold=", curGold)
}

func (this *GamePlayer) AddSoulCur(val int32) {
	oldSoul := this.MyUnit.GetIProperty(prpc.IPT_SOULCUR)
	curSoul := oldSoul + val
	if curSoul < 0 {
		curSoul = 0
	}
	if curSoul > CopperMax {
		curSoul = CopperMax
	}
	this.MyUnit.SetIProperty(prpc.IPT_SOULCUR, curSoul)

	logs.Info("Player[", this.MyUnit.InstName, "]", "Old MySoulCur=", oldSoul, "curSoulCur=", curSoul)
}

func (this *GamePlayer) ClearAllBuff() {
	logs.Info("ClearAllBuff")
	this.MyUnit.ClearAllbuff()

	for _, unit := range this.UnitList {
		unit.ClearAllbuff()
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////升级 强化....
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) PromoteUnit(unitid int64) {
	logs.Info("PromoteUnit", unitid)

	if this.IsLock() {
		logs.Info("toooooo fast")
		return
	}

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

	if items == nil || len(items) <= 0 {
		logs.Info("cant find item, this id is ", level_info.ItemId)
		return
	}

	var num int32
	for _, item := range items {
		num += item.Stack
	}

	if num < level_info.ItemNum {
		logs.Info("数量不足, this id is ", level_info.ItemId, "need itemnum is ", level_info.ItemNum, "i have num is ", num)
		return
	}

	this.DelItemByTableId(level_info.ItemId, level_info.ItemNum)

	unit.Promote(level_info)

	logs.Info("PromoteUnitOK", unitid)
	this.session.PromoteUnitOK()
}

func (this *GamePlayer) MyUnitLevelUp() {

	LevelUp_info := GetPromoteRecordById(1)
	logs.Info("MyUnitLevelUp 1", LevelUp_info)

	if LevelUp_info == nil {
		return
	}

	if this.MyUnit.Level >= int32(len(LevelUp_info)) {
		return
	}

	level_info := LevelUp_info[this.MyUnit.Level]

	this.MyUnit.Promote(level_info)
	this.MyUnit.Level = level_info.Level
	logs.Info("MyUnitLevelUp 2", level_info)

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

func (this *GamePlayer) CaleMyEnergy(dt float64) {
	myEnergy := this.MyUnit.GetIProperty(prpc.IPT_ENERGY)
	if myEnergy >= 1000 {
		return
	}
	this.EnergyTimer += dt
	if this.EnergyTimer >= 300 {
		this.SetMyEnergy(1, true)
		this.EnergyTimer = 0
	}
}

func (this *GamePlayer) SetMyEnergy(val int32, isAdd bool) {
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
	logs.Debug("SetMyEnergy Val=", val, "CurmyEnergy=", myEnergy)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////技能学习
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) EquipSkill(skillinfo prpc.COM_LearnSkill) {

	logs.Info("EquipSkill", skillinfo)

	skill := GetRoleSkillRecordById(skillinfo.SkillID)
	if skill == nil {
		return //错误的skill
	}

	if this.MyUnit.Level < skill.OpenLv {
		return
	}

	real_skillid := this.SkillBase[skillinfo.SkillID]
	learnSkill := InitSkillFromTable(real_skillid)
	if learnSkill == nil {
		learnSkill = InitSkillFromTable(skill.SKillID)
	}
	if learnSkill == nil {
		return
	}

	switch skillinfo.Position {
	case 0:
		if skill.Type != 0 {
			logs.Info("EquipSkill skill.Type wrong", 0)
			return
		}
	case 1, 2:
		if skill.Type != 1 {
			logs.Info("EquipSkill skill.Type wrong", 1, 2)
			return
		}
	case 3:
		if skill.Type != 2 {
			logs.Info("EquipSkill skill.Type wrong", 3)
			return
		}
	}

	this.MyUnit.Skill[skillinfo.Position] = learnSkill

	var idx int32 = -1
	for index, skill_ := range this.MyUnit.Skill {
		logs.Info("skill", skill_, &skill_)
		if skill_ == nil {
			continue
		}
		if skill_.SkillID == learnSkill.SkillID {
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

	logs.Info("skillall", this.MyUnit.Skill)

	this.session.EquipSkillOK(skillinfo.Position, learnSkill.SkillID)
	logs.Info("EquipSkillOK", skillinfo.Position, learnSkill.SkillID)

	//如果是被动技能 需要修改buff

	return
}

func (this *GamePlayer) SkillUpdate(skillindex int32, skillId int32) {
	logs.Info("SkillUpdate", skillindex, skillId)
	if this.SkillBase[skillindex] == 0 {
		return
	}

	updateInfo := GetRoleSkillUpdateRecordById(skillId)

	if updateInfo == nil {
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
	for _, item := range items {
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
	logs.Info("SkillUpdateOK", skillindex, updateInfo.NextID, skillpos)
	logs.Info("SkillUpdateOK 1", this.SkillBase)

}

func (this *GamePlayer) SkillUpdate_equip(position int32, skillId int32) {
	updateInfo := GetRoleSkillUpdateRecordById(skillId)

	if updateInfo == nil {
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

	if change != 0 {
		this.SkillBase[change] = updateInfo.NextID
	}

	//这里需要buff更新一下

}

func (this *GamePlayer) CheckSkillBase() {
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

func (this *GamePlayer) BuyShopItem(shopId int32) {
	shopData := GetShopDataById(shopId)
	if shopData == nil {
		return
	}

	if shopData.CurrenciesKind == prpc.IPT_GOLD {

		myGold := this.MyUnit.GetIProperty(prpc.IPT_GOLD)
		if myGold < shopData.Price {
			return
		}
		this.AddGold(-shopData.Price)
		logs.Info("Player[", this.MyUnit.InstName, "]", "BuyShopItem ShopId=", shopId, "Shoping Spend=", shopData.Price)
	}

	if shopData.CurrenciesKind == prpc.IPT_SOULCUR {
		mySoul := this.MyUnit.GetIProperty(prpc.IPT_SOULCUR)
		if mySoul < shopData.Price {
			return
		}
		this.AddSoulCur(-shopData.Price)
		logs.Info("Player[", this.MyUnit.InstName, "]", "BuyShopItem ShopId=", shopId, "Shoping Spend=", shopData.Price)
	}

	if shopData.ShopType == prpc.SHT_BuyCopper {
		if shopData.Copper > 0 {
			this.AddCopper(shopData.Copper)
			if shopData.CardPondId != 0 {
				this.OpenTreasureBox(shopData.CardPondId)
			}
		} else {
			logs.Info("Player[", this.MyUnit.InstName, "]", "shopData.Copper A wrong number", "BuyShopItem ShopId=", shopId)
		}
	}

	if shopData.ShopType == prpc.SHT_BlackMarket {
		if this.IsBuyBlackMarketItem(shopData.ShopId) {
			this.AddBagItemByItemId(shopData.ShopItemId, shopData.ShopItemNum)
			itemInsts := []prpc.COM_ItemInst{}
			item := prpc.COM_ItemInst{}
			item.ItemId = shopData.ShopItemId
			item.Stack = shopData.ShopItemNum
			itemInsts = append(itemInsts, item)
			if this.session != nil {
				this.session.BuyShopItemOK(itemInsts)
			}
		} else {
			logs.Info("Player[", this.MyUnit.InstName, "]", "IsBuyBlackMarketItem==false", "BuyShopItem ShopId=", shopId)
		}
	}
}

func (this *GamePlayer) OpenTreasureBox(pondId int32) bool {
	data := GetCardPondTableDataById(pondId)
	if data == nil {
		logs.Info("Player[", this.MyUnit.InstName, "]", "OpenTreasureBox GetCardPondTableDataById Can Not Find pondId=", pondId)
		return false
	}

	var items []int32

	greenItems := data.GetGreenCardItems()
	buleItems := data.GetBlueCardItems()
	purplenItems := data.GetPurpleCardItems()
	orangeItems := data.GetOrangeCardItems()

	for _, itemId := range greenItems {
		items = append(items, itemId)
		//logs.Info("Player[",this.MyUnit.InstName,"]","OpenTreasureBox GreenItem itemId",itemId)
	}
	for _, itemId := range buleItems {
		items = append(items, itemId)
		//logs.Info("Player[",this.MyUnit.InstName,"]","OpenTreasureBox BuleItem itemId",itemId)
	}
	for _, itemId := range purplenItems {
		items = append(items, itemId)
		//logs.Info("Player[",this.MyUnit.InstName,"]","OpenTreasureBox PurlenItem itemId",itemId)
	}
	for _, itemId := range orangeItems {
		items = append(items, itemId)
		//logs.Info("Player[",this.MyUnit.InstName,"]","OpenTreasureBox OrangeItem itemId",itemId)
	}

	itemInsts := []prpc.COM_ItemInst{}
	for _, itemId := range items {
		var isHave bool = false
		for i := 0; i < len(itemInsts); i++ {
			if itemInsts[i].ItemId == itemId {
				itemInsts[i].Stack++
				isHave = true
				break
			}
		}
		if !isHave {
			item := prpc.COM_ItemInst{}
			item.ItemId = itemId
			item.Stack = 1
			itemInsts = append(itemInsts, item)
		}
	}

	for _, item := range itemInsts {
		this.AddBagItemByItemId(item.ItemId, item.Stack)
		logs.Info("Player[", this.MyUnit.InstName, "]", "OpenTreasureBox AddItem ID=", item.ItemId, "Num=", item.Stack)
	}

	if this.session != nil {
		this.session.BuyShopItemOK(itemInsts)
	}

	return true
}

func (this *GamePlayer) Logout() {

	logs.Debug("Logout", "PlayerName=", this.MyUnit.InstName)
	logs.Debug("Logout", "PlayerName=", this.Friends)
	logs.Debug("Logout", "PlayerName=", this.Enemys)

	this.LogoutTime = time.Now().Unix()

	if this.GuildId != 0 {
		pGuild := FindGuildById(this.GuildId)
		if pGuild!=nil {
			pGuild.GuildMemberOffOnLine(this.MyUnit.InstId)
		}
	}

	//清理战斗信息
	this.LeftBattle_strong()

	this.GamePlayerSave()

	RemovePlayer(this)
}

func (this *GamePlayer) GamePlayerSave() {
	p := this.GetPlayerSGE()

	logs.Debug("GamePlayerSave", p.Enemys)
	UpdatePlayer(p)
}

func Save() {
	for _, p := range PlayerStore {
		if p == nil {
			continue
		}
		p.GamePlayerSave()
	}
}

func (this *GamePlayer) CardDebrisResolve(itemInstId int64, num int32) {
	item := this.GetBagItemByInstId(itemInstId)
	if item == nil {
		return
	}

	if item.Stack < num {
		return
	}

	itemData := GetItemTableDataById(item.ItemId)
	if itemData == nil {
		return
	}
	if itemData.ItemMainType == prpc.IMT_Debris {
		this.AddSoulCur(itemData.SoulVal * num)
	}
	this.DelItemByInstId(itemInstId, num)
}

func (this *GamePlayer) InitMyBlackMarket() {
	tempData := prpc.COM_BlackMarket{}
	tempData.RefreshTime = time.Now().Unix()
	tempData.RefreshNum = int32(GetGlobalInt("C_BlackMarkteRefreshNum"))

	items := GetBlackMarketShopItems()
	for i := 0; i < len(items); i++ {
		item := prpc.COM_MlackMarketItemData{}
		item.ItemId = items[i]
		item.IsBuy = true
		tempData.ShopItems = append(tempData.ShopItems, item)
	}

	this.BlackMarketData = &tempData

	logs.Info("Player[", this.MyUnit.InstName, "]", "InitMyBlackMarket ", this.BlackMarketData.ShopItems, this.BlackMarketData.RefreshTime, this.BlackMarketData.RefreshNum, len(this.BlackMarketData.ShopItems))

	if this.session != nil {
		logs.Info("Player[", this.MyUnit.InstName, "]", "InitMyBlackMarket ", this.BlackMarketData.ShopItems, this.BlackMarketData.RefreshTime, this.BlackMarketData.RefreshNum)
		this.session.SycnBlackMarkte(*this.BlackMarketData)
	}
}

func (this *GamePlayer) RefreshMyBlackMarket(isActive bool) {
	if this.BlackMarketData == nil {
		return
	}
	if isActive {
		need := GetGlobalInt("C_BlackMarkteRefreshSpeed")
		if this.MyUnit.GetIProperty(prpc.IPT_SOULCUR) < int32(need) {
			return
		}
		this.AddSoulCur(-int32(need))
		this.BlackMarketData.RefreshNum--
	} else {
		this.BlackMarketData.RefreshTime = time.Now().Unix()
	}

	tempData := prpc.COM_BlackMarket{}
	tempData.RefreshNum = this.BlackMarketData.RefreshNum
	tempData.RefreshTime = this.BlackMarketData.RefreshTime
	items := GetBlackMarketShopItems()
	for i := 0; i < len(items); i++ {
		item := prpc.COM_MlackMarketItemData{}
		item.ItemId = items[i]
		item.IsBuy = true
		tempData.ShopItems = append(tempData.ShopItems, item)
	}
	this.BlackMarketData = &tempData

	logs.Info("Player[", this.MyUnit.InstName, "]", "RefreshMyBlackMarket ", this.BlackMarketData.ShopItems, this.BlackMarketData.RefreshTime, this.BlackMarketData.RefreshNum, len(this.BlackMarketData.ShopItems))
	if this.session != nil {
		this.session.SycnBlackMarkte(*this.BlackMarketData)
	}
}

func CheckMyBlackMarkte() {
	for _, p := range PlayerStore {
		if p == nil {
			continue
		}
		p.RefreshMyBlackMarket(false)
	}
}

func (this *GamePlayer) IsBuyBlackMarketItem(shopId int32) bool {
	if this.BlackMarketData == nil {
		return false
	}
	for i := 0; i < len(this.BlackMarketData.ShopItems); i++ {
		if this.BlackMarketData.ShopItems[i].ItemId == shopId && this.BlackMarketData.ShopItems[i].IsBuy {
			this.BlackMarketData.ShopItems[i].IsBuy = false
			if this.session != nil {
				this.session.SycnBlackMarkte(*this.BlackMarketData)
			}
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//查询
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) QueryPlayerInfo(Instid int64) {
	info := prpc.COM_PlayerInfo{}
	info.IsOnline = false
	p := FindPlayerByInstId(Instid)
	if p == nil {
		var p *prpc.SGE_DBPlayer
		if p = <- QueryPlayerById(Instid); p!=nil {		//数据库里有这个人
			player := &GamePlayer{}
			player.SetPlayerSGE(*p)
			info.Name = player.MyUnit.InstName
			info.Level = player.MyUnit.Level
			info.UnitID = player.MyUnit.UnitId
			info.TiatiVal = player.TianTiVal
			info.TiatiRank = player.TianTiRank

			if player.GuildId == 0 {
				info.ClanName = "无帮派"
			} else {
				guild := FindGuildById(player.GuildId)
				if guild == nil  {
					info.ClanName = "无帮派"
				} else {
					info.ClanName = guild.GuildData.GuildName
				}
			}

			group := player.GetUnitGroupById(player.BattleUnitGroup)
			if group != nil {
				for _, unitid := range group.UnitList {
					u := player.GetUnit(unitid)
					info.UnitLIst = append(info.UnitLIst, u.GetUnitCOM())
				}
			}

			logs.Debug("QueryPlayerInfo offline ", info)
		} else {
			return
		}
	} else {
		info.IsOnline = true
		info.Name = p.MyUnit.InstName
		info.Level = p.MyUnit.Level
		info.UnitID = p.MyUnit.UnitId
		info.TiatiVal = p.TianTiVal
		info.TiatiRank = p.TianTiRank

		if p.GuildId == 0 {
			info.ClanName = "无帮派"
		} else {
			guild := FindGuildById(p.GuildId)
			if guild == nil  {
				info.ClanName = "无帮派"
			} else {
				info.ClanName = guild.GuildData.GuildName
			}
		}

		group := p.GetUnitGroupById(p.BattleUnitGroup)
		if group != nil {
			for _, unitid := range group.UnitList {
				u := p.GetUnit(unitid)
				info.UnitLIst = append(info.UnitLIst, u.GetUnitCOM())
			}
		}
		logs.Debug("QueryPlayerInfo online ", info)
	}

	logs.Debug("QueryPlayerInfo end", info)

	this.session.QueryPlayerInfoOK(info)

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//新手引導
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) NewPlayerGuide(Step uint64) {
	logs.Debug("NewPlayerGuide ", Step)
	this.Guide = Step

	return
}

func (this *GamePlayer) NewPlayerGuide_SetOver() {
	this.Guide = 9999999

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) IsLock() bool {
	//if this.LockTime == 0 {
	//	this.LockTime = time.Now().Unix()
	//	return false
	//}
	//
	//Now := time.Now().Unix()
	//
	//if Now - this.LockTime < 1 {
	//	return true
	//}
	//
	//this.LockTime = Now

	return false
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestPlayer() {
	player := CreatePlayer(1,"test123","123")
	if player==nil {
		return
	}
	logs.Info("Chapters",player.Chapters)

	for i:=0;i<10 ;i++  {
		player.RandChapterGo()
	}

	logs.Info("Chapters new",player.Chapters)
}


func (this *GamePlayer) TestItem() {
	for i := 1; i < 9; i++ { //测试用
		this.AddBagItemByItemId(int32(i), 10)
	}
	this.AddBagItemByItemId(5000, 2000)
	this.AddCopper(10000000)
	this.AddGold(10000)
	this.GiveDrop(1000)

	for _, u := range this.UnitList {
		logs.Info("Myself Unit InstId", u.InstId, "InstName", u.InstName)
	}
}

func (this *GamePlayer) TestTop() {
	this.UpdateTianTiVal()
	this.TianTiVal = 10000
	this.UpdateTianTiVal()
	RefreshAllTopList()
}
