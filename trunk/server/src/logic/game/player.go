package game

import "logic/prpc"

type GamePlayer struct{
	MyUnit 		GameUnit
	UnitList	[]GameUnit
}

func CreatePlayer(tid int32, name string) *GamePlayer{
	p := GamePlayer{}
	p.MyUnit = CreateUnitFromTable(tid)
	p.MyUnit.InstName = name
	return  &p
}

func(this* GamePlayer)GetPlayerCOM() prpc.COM_Player{
	p := prpc.COM_Player{}
	p.InstId = this.MyUnit.InstId
	p.Name = this.MyUnit.InstName
	p.Unit = this.MyUnit.GetUnitCOM()
	return p
}