package game

import (
	"sync/atomic"
	"logic/prpc"
	"sync"
)

var genInstId int64 = 1

type GameUnit struct {
	sync.Mutex
	Owner		*GamePlayer //所有者
	UnitId	   	int32
	InstId      	int64
	InstName 	string
	DisPlay		int32
	IProperties 	[]int32
	CProperties 	[]float32
	Skill 		map[int32]*Skill

	//战斗的实际信息
	Position	int32 //prpc.BattlePosition
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
	for i := 0; i <len(t.Skills); i++ {
		if t.Skills[i] == 0{
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

func (this *GameUnit)GetBattleCamp()int{
	if this.Owner != nil {
		return  this.Owner.BattleCamp
	}
	return  prpc.CT_MAX
}

func (this* GameUnit) GetCProperty(id int32)float32{
	if id <= prpc.CPT_MIN || id >= prpc.CPT_MAX{
		return 0
	}
	return this.CProperties[id]
}

func(this* GameUnit)GetUnitCOM()prpc.COM_Unit{
	u := prpc.COM_Unit{}
	u.UnitId = this.UnitId
	u.InstId = this.InstId
	u.IProperties = append(u.IProperties, this.IProperties...)
	u.CProperties = append(u.CProperties, this.CProperties...)
	return u
}

func(this *GameUnit)GetBattleUnitCOM()prpc.COM_BattleUnit{
	u := prpc.COM_BattleUnit{}
	u.Position = this.Position
	u.InstId = this.InstId
	u.UnitId = this.UnitId
	u.HP = int32(this.GetCProperty(prpc.CPT_HP))
	u.Position = this.Position
	u.Name = this.InstName
	return  u
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

func(this* GameUnit)CastSkill(battle *BattleRoom) (prpc.COM_BattleAction, bool) {
	skill := this.SelectSkill(battle.Round)

	tagetList := battle.SelectAllTarget(this.Owner.BattleCamp)

	acc, dead := skill.Action(this, tagetList, battle.Round)

	return acc, dead
}

func (this *GameUnit) IsDead() bool {
	return this.GetCProperty(prpc.CPT_HP) <= 0
}
