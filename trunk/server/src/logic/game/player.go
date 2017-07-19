package game

import (
	"logic/prpc"
)

type GamePlayer struct {
	session        *Session    //链接
	MyUnit         *GameUnit   //自己的卡片
	UnitList       []*GameUnit //拥有的卡片
	BattleUnitList []int64     //默认出战卡片
}

func (this *GamePlayer) SetSession(session *Session) {
	this.session = session
}

func CreatePlayer(tid int32, name string) *GamePlayer {
	p := GamePlayer{}
	p.MyUnit = CreateUnitFromTable(tid)
	p.MyUnit.InstName = name

	//来两个默认的小兵
	p.UnitList = append(p.UnitList, CreateUnitFromTable(2))
	p.UnitList = append(p.UnitList, CreateUnitFromTable(3))

	return &p
}

func (this *GamePlayer) GetPlayerCOM() prpc.COM_Player {
	p := prpc.COM_Player{}
	p.InstId = this.MyUnit.InstId
	p.Name = this.MyUnit.InstName
	p.Unit = this.MyUnit.GetUnitCOM()
	for _, u := range this.UnitList {
		p.Employees = append(p.Employees, u.GetUnitCOM())
	}
	return p
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//角色数据接口
func (this *GamePlayer) GetUnit(instId int64) *GameUnit {
	for _, v := range this.UnitList {
		if v.InstId == instId {
			return v
		}
	}
	return nil
}
func (this *GamePlayer) GetBattleUnit(instId int64) *GameUnit {
	for _, v := range this.BattleUnitList {
		if v == instId {
			return this.GetUnit(instId)
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//战斗相关
func (this *GamePlayer) SetBattleUnit(instId int64) {
	if this.GetUnit(instId) == nil {
		return //没有设置你妹
	}
	if this.GetBattleUnit(instId) != nil {
		return //在出战设置你妹
	}
	this.BattleUnitList = append(this.BattleUnitList, instId)
}

func (this *GamePlayer) JoinBattle() {

}

func (this *GamePlayer) SetupBattle() {

}
