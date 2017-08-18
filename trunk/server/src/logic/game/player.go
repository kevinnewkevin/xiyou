package game

import (
	"errors"
	"fmt"
	"logic/prpc"
	"sync"
)

type GamePlayer struct {
	sync.Locker
	session        	*Session    //链接
	MyUnit         	*GameUnit   //自己的卡片
	UnitList       	[]*GameUnit //拥有的卡片
	BattleUnitList 	[]int64     //默认出战卡片

	//战斗相关辅助信息
	BattleId   		int64 		//所在房间编号
	BattleCamp 		int   		//阵营 //prpc.CompType
	IsActive   		bool  		//是否激活
	KillUnits 	 	[]int32 	//杀掉的怪物
	MyDeathNum		int32		//战斗中自身死亡数量
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//角色创建
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) SetSession(session *Session) {
	this.session = session
}

func CreatePlayer(tid int32, name string) *GamePlayer {
	p := GamePlayer{}
	p.MyUnit = CreateUnitFromTable(tid)
	p.MyUnit.InstName = name
	p.MyUnit.Owner = &p

	//来两个默认的小兵
	p.UnitList = append(p.UnitList, CreateUnitFromTable(2))
	p.UnitList = append(p.UnitList, CreateUnitFromTable(3))

	//給默認的小兵設置主人
	for _, unit := range p.UnitList {
		unit.Owner = &p
	}

	//TestActionByLua()

	return &p

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
	return p
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
	battleRoom.SetupPosition(this, pos)

	this.session.SetupBattleOK()

	//battleRoom.Update()

	return nil
}

func (this *GamePlayer) SetProprty(battleid int64, camp int) {
	this.BattleId = battleid
	this.BattleCamp = camp
	this.MyUnit.ResetBattle(camp, true, battleid)

	for _, u := range this.UnitList {
		u.ResetBattle(camp, false, battleid)
	}
}
