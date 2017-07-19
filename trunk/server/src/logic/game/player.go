package game

import (
	"logic/prpc"
	"fmt"
)

type GamePlayer struct{
	MyUnit 		*GameUnit
	UnitList	[]*GameUnit
}

func CreatePlayer(tid int32, name string) *GamePlayer{
	p := GamePlayer{}
	p.MyUnit = CreateUnitFromTable(tid)
	p.MyUnit.InstName = name

	//来两个默认的小兵
	p.UnitList = append(p.UnitList,CreateUnitFromTable(2))
	p.UnitList = append(p.UnitList,CreateUnitFromTable(3))
	fmt.Println("CreatePlayer")
	fmt.Println("CreatePlayer", p.MyUnit)

	return  &p
}

func(this* GamePlayer)GetPlayerCOM() prpc.COM_Player{
	p := prpc.COM_Player{}
	p.InstId = this.MyUnit.InstId
	p.Name = this.MyUnit.InstName
	p.Unit = this.MyUnit.GetUnitCOM()
	for _, u := range this.UnitList{
		p.Employees = append(p.Employees,u.GetUnitCOM())
	}
	return p
}