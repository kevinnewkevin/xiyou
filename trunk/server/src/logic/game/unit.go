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
}

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
	return this.GetCProperty(prpc.CPT_HP) <= 0
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
	this.Buff = []*Buff{}
	this.Debuff = []*Buff{}
	this.Allbuff = []*Buff{}
	this.Camp = camp
	this.IsMain = ismain
	this.BattleId = battleid
}

func (this *GameUnit)CheckBuff (round int32){
	//检测那些有行为的buff 比如定时增加血量的那种

}

func (this *GameUnit)CheckDebuff (round int32){
	//检测那些有行为的debuff 比如定时损血

}

func (this *GameUnit)CheckAllBuff (round int32){
	fmt.Println("checkallBuff round is ", round)			//檢測buff效果
	needDelete := map[*Buff]int{}
	this.DelBuff = []*Buff{}

	for _, buff := range this.Allbuff{
		if this.IsDead() {		//buff執行中玩家卡牌可能死掉
			break
		}
		if buff.Update(round) {
			needDelete[buff] = 1
			this.DelBuff = append(this.DelBuff, buff)		//這個是給戰鬥房間用的 用來寫入戰報
		}
	}

	this.deletBuff(needDelete)
}

func (this *GameUnit) deletBuff (need map[*Buff]int){
	newList := []*Buff{}
	for _, buff := range this.Allbuff {
		_, ok := need[buff]
		if ok {
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

func (this *GameUnit) PopAllBuffByDebuff() {
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