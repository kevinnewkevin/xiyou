package game

import (
	"sync/atomic"
	"logic/prpc"
	"fmt"
)

var genInstId int64 = 1

type GameUnit struct {
	UnitId	    int32
	InstId      int64
	InstName 	string
	DisPlay		int32
	IProperties []int32
	CProperties []float32
	Skill 		map[int32]*Skill
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
	fmt.Println(&u)
	return &u
}

func(this* GameUnit)GetUnitCOM()prpc.COM_Unit{
	u := prpc.COM_Unit{}
	u.UnitId = this.UnitId
	u.InstId = this.InstId
	u.IProperties = append(u.IProperties, this.IProperties...)
	u.CProperties = append(u.CProperties, this.CProperties...)
	return u
}
