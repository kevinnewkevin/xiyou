package game

import (
	"logic/prpc"
	"sync"
	"sync/atomic"
	"fmt"
)

var genInstId int64 = 1

type GameUnit struct {
	sync.Mutex
	Owner       *GamePlayer //所有者
	IsMain		bool
	Camp		int
	UnitId      int32
	InstId      int64
	InstName    string
	DisPlay     int32
	Level     	int32
	IProperties []int32
	CProperties []float32
	Skill       map[int32]*Skill

	//战斗的实际信息
	Position 	int32 //prpc.BattlePosition
	Buff 		[]*Buff //增益状态
	Debuff 		[]*Buff //负面状态
	Allbuff		[]*Buff	//全体buff
	DelBuff		[]*Buff	//需要刪除的buff
	BattleId	int64	//zhandou id
	//战斗 buff需要的数据
	VirtualHp	int32	//护盾数值
	Special 	map[int32][]int32	//特殊属性效果
}
//如果是创建怪物卡牌的话 player = 你来
func CreateUnitFromTable(id int32) *GameUnit {
	t := GetUnitRecordById(id)
	if t == nil {
		return nil
	}
	u := GameUnit{}
	u.UnitId = t.Id
	u.InstId = atomic.AddInt64(&genInstId, 1)
	u.IProperties = append(u.IProperties, t.IProp...)
	u.CProperties = append(u.CProperties, t.CProp...)
	u.DisPlay = t.DispId
	u.Skill = map[int32]*Skill{}
	u.Level = 1
	u.InstName = t.BaseName
	for i := 0; i < len(t.Skills); i++ {
		if t.Skills[i] == 0 {
			continue
		}
		skill := InitSkillFromTable(t.Skills[i])
		if skill == nil {
			continue
		}
		u.Skill[int32(i)] = skill
	}
	return &u
}

func (this *GameUnit) GetBattleCamp() int {

	return this.Camp
}

func (this *GameUnit) GetCProperty(id int32) float32 {
	if id <= prpc.CPT_MIN || id >= prpc.CPT_MAX {
		return 0
	}
	return this.CProperties[id]
}

func (this *GameUnit) GetIProperty(id int32) int32 {
	if id <= prpc.IPT_MIN || id >= prpc.IPT_MAX {
		return 0
	}
	return this.IProperties[id]
}

func (this *GameUnit) ChangeSpec(spec string, buffinstid int32) {
	spe := prpc.ToId_BuffSpecial(spec)
	bufflist, ok := this.Special[int32(spe)]
	if !ok {
		this.Special[int32(spe)] = []int32{buffinstid}
	} else {
		this.Special[int32(spe)] = append(bufflist, buffinstid)
	}
	return
}

func (this *GameUnit) CheckSpec(spec string) bool { //unit.checkspec(是否有免死)
	spe := prpc.ToId_BuffSpecial(spec)
	bufflist, ok := this.Special[int32(spe)]

	if !ok {
		return false
	}

	if len(bufflist) == 0 {
		return false
	}

	return true
}

func (this *GameUnit) ClacSheldPer(round int32) float32 {			//计算百分比减伤 所有buff的百分比减伤加起来 有个最大值
	maxPer := 75

	bl, ok := this.Special[prpc.BF_SHELD]

	if !ok || len(bl) == 0 {
		return 0
	}

	sheld := 0

	for _, instid := range bl {
		buff := this.SelectBuff(instid)
		if buff == nil || buff.IsOver(round) {
			continue
		}
		sheld += int(buff.Data)
	}

	if sheld > maxPer {
		sheld = maxPer
	}

	per := float32(sheld) / 100.0

	return per
}

func (this *GameUnit) GetUnitCOM() prpc.COM_Unit {
	u := prpc.COM_Unit{}
	u.UnitId = this.UnitId
	u.InstId = this.InstId
	u.IProperties = append(u.IProperties, this.IProperties...)
	u.CProperties = append(u.CProperties, this.CProperties...)
	return u
}

func (this *GameUnit) GetBattleUnitCOM() prpc.COM_BattleUnit {
	u := prpc.COM_BattleUnit{}
	u.Position = this.Position
	u.InstId = this.InstId
	u.UnitId = this.UnitId
	u.HP = int32(this.GetCProperty(prpc.CPT_HP))
	u.CHP = int32(this.GetCProperty(prpc.CPT_CHP))
	u.Position = this.Position
	u.Name = this.InstName
	return u
}

func (this *GameUnit) SelectSkill(round int32) *Skill {
	var idx int32
	if round > 2 {
		idx = round % 3
	} else {
		idx = round
	}

	return this.Skill[idx]
}

func (this *GameUnit) CastSkill(battle *BattleRoom) bool {
	skill := this.SelectSkill(battle.Round)

	tagetList := battle.SelectAllTarget(this.Camp)

	battle.AcctionList.InstId = this.InstId
	battle.AcctionList.SkillId = skill.SkillID

	acc, dead := skill.Action(this, tagetList, battle.Round)

	battle.AcctionList.TargetList = acc
	//fmt.Println("CastSkill, acc ", acc)
	//fmt.Println("CastSkill, AcctionList ", battle.AcctionList)

	return dead
}

func (this *GameUnit) CastSkill2(battle *BattleRoom) bool {
	if this.IsDead() {
		return false
	}

	skill := this.SelectSkill(battle.Round)

	//fmt.Println("CastSkill2 skill id is ", skill.SkillID)

	battle.AcctionList.InstId = this.InstId
	battle.AcctionList.SkillId = skill.SkillID

	skill.ActionBylua(battle.InstId, this.InstId)

	//fmt.Println("CastSkill, AcctionList ", battle.AcctionList)

	return false
}

func (this *GameUnit) GetFrontList() []int {
	if this.GetBattleCamp() == prpc.CT_RED {
		return []int{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_3}
	} else {
		return []int{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_3}
	}
}

func (this *GameUnit) GetBackList() []int {
	if this.GetBattleCamp() == prpc.CT_RED {
		return []int{prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_6}
	} else {
		return []int{prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
	}
}

func (this *GameUnit) IsDead() bool {
	return this.GetCProperty(prpc.CPT_CHP) <= 0
}

func (this *GameUnit) isFront() bool {
	li := this.GetFrontList()

	for _, v := range li {
		if v == int(this.Position) {
			return true
		}
	}
	return false
}

func (this *GameUnit) isBack() bool {
	li := this.GetBackList()

	for _, v := range li {
		if v == int(this.Position) {
			return true
		}
	}
	return false
}

func (this *GameUnit)ResetBattle(camp int, ismain bool, battleid int64) {
	this.CProperties[prpc.CPT_HP] = float32(this.IProperties[prpc.IPT_HP])
	this.CProperties[prpc.CPT_CHP] = float32(this.IProperties[prpc.IPT_HP])
	this.Buff = []*Buff{}
	this.Debuff = []*Buff{}
	this.Allbuff = []*Buff{}
	this.Camp = camp
	this.IsMain = ismain
	this.BattleId = battleid
	this.VirtualHp = 0
	this.Special = map[int32][]int32{}
}

func (this *GameUnit)CheckBuff (round int32){
	//检测那些有行为的buff 比如定时增加血量的那种

}

func (this *GameUnit)CheckDebuff (round int32){
	//检测那些有行为的debuff 比如定时损血

}

func (this *GameUnit)SelectBuff (instid int32) *Buff {
	for _, buff := range this.Allbuff {
		if buff.InstId == instid {
			return buff
		}
	}

	return nil
}

func (this *GameUnit)CheckAllBuff (round int32){
	fmt.Println(this.InstId, "checkallBuff round is ", round)			//檢測buff效果
	needDelete := map[*Buff]int{}
	this.DelBuff = []*Buff{}

	for _, buff := range this.Allbuff{
		if this.IsDead() {		//buff執行中玩家卡牌可能死掉
			break
		}
		if buff.Update(round) {
			fmt.Println("CheckAllBuff one", buff.InstId, buff.Round)
			needDelete[buff] = 1
			this.DelBuff = append(this.DelBuff, buff)		//這個是給戰鬥房間用的 用來寫入戰報
		}
	}

	fmt.Println(this.InstId, "checkallBuff over", len(needDelete))			//檢測buff效果
	this.deletBuff(needDelete)
}

func (this *GameUnit) deletBuff (need map[*Buff]int){
	newList := []*Buff{}
	for _, buff := range this.Allbuff {
		_, ok := need[buff]
		if ok {
			buff.DeleteProperty(this.BattleId, this.InstId)
			continue
		}
		newList = append(newList, buff)
	}

	fmt.Println("deletBuff", need)
	this.Allbuff = newList
}

func erase(arr []interface{} , idx int) []interface{}{
	return	append(arr[:idx], arr[idx+1:]...)
}

func (this *GameUnit) PopAllBuffByDebuff() int {
//删除卡牌身上所有的debuff
	tmp := map[*Buff]int{}

	for _, buff := range this.Allbuff {
		if buff.BuffType == kTypeBuff{
			continue
		}
		tmp[buff] = 1
	}

	newBufflist := []*Buff{}

	for _, v := range this.Allbuff {
		_, ok := tmp[v]
		if ok {
			continue
		}

		newBufflist = append(newBufflist, v)
	}

	fmt.Println("PopAllBuffByDebuff")
	this.Allbuff = newBufflist

	return len(tmp)
}

func (this *GameUnit) PopAllBuffByBuff() {
//删除卡牌身上的buff
	tmp := map[*Buff]int{}

	for _, buff := range this.Allbuff {
		if buff.BuffType == kTypeDebuff{
			continue
		}
		tmp[buff] = 1
	}

	newBufflist := []*Buff{}

	for _, v := range this.Allbuff {
		_, ok := tmp[v]
		if ok {
			continue
		}

		newBufflist = append(newBufflist, v)
	}

	fmt.Println("PopAllBuffByBuff")
	this.Allbuff = newBufflist
}