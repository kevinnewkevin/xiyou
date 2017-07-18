package game

import (
	"logic/application"
	"sync/atomic"
	"logic/prpc"
)

var genInstId = 1

type GameUnit struct {
	InstId      int64
	InstName	string
	IProperties []int32
	CProperties []float32
}

func CreateUnitFromTable(id int32) *GameUnit {
	t := application.GetUnitRecordById(id)
	if t == nil {
		return nil
	}
	u := GameUnit{}
	u.InstId = atomic.AddInt64(&genInstId, 1)
	copy(u.IProperties, t.IProp)
	copy(u.CProperties, t.CProp)
	return &u
}

func(this* GameUnit)GetUnitCOM()prpc.COM_Unit{
	u := prpc.COM_Unit{}
	u.InstId = this.InstId
	copy(u.IProperties, this.IProperties)
	copy(u.CProperties, this.CProperties)
	return &u
}
