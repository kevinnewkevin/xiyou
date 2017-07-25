package game

import (
	"logic/prpc"
	"fmt"
	"errors"
)

type GamePlayer struct {
	session        *Session    //链接
	MyUnit         *GameUnit   //自己的卡片
	UnitList       []*GameUnit //拥有的卡片
	BattleUnitList []int64     //默认出战卡片
	BattleRoom     int64	   //所在房间编号
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
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (this *GamePlayer) GetUnit(instId int64) *GameUnit {
	if this.MyUnit.InstId == instId{
		return  this.MyUnit
	}

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
//技能相关
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) StudySkill(UnitID int64, skillpos int32, skillid int32) error {
	if skillpos >= 2{
		fmt.Println("技能位置錯誤")
		return errors.New("技能位置錯誤")
	}
	unit := this.GetUnit(UnitID)
	skill := InitSkillFromTable(skillid)

	unit.Skill[skillpos] = skill

	return nil
}

func (this *GamePlayer) UseSkill(attacker int64, defender int64, skillid int32) {
	attack := this.GetBattleUnit(attacker)		//攻擊卡牌
	skill, ok := attack.Skill[skillid]

	if !ok {
		fmt.Println("這個卡牌沒有這個技能")
	}

	if !skill.Condition() {
		fmt.Println("技能不能使用")
	}

	battleRoom, ok := BattleRoomList[this.BattleRoom]
	if !ok {
		fmt.Println("不在房間中")
	}

	targetPlayer, ok := battleRoom.Target[this.MyUnit.InstId]
	if !ok {
		fmt.Println("目標卡牌的主人不在房間中")
	}


	skill.Action(attack, skill.StandbySkill(defender, targetPlayer.Player), battleRoom.Bout)
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//战斗相关 设置卡牌
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) SetBattleUnit(instId int64) {		//往战斗池里设置出战卡牌
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

func (this *GamePlayer) SetupBattle(pos []prpc.COM_BattlePosition) error {		//卡牌上阵
	poss := map[int64]*Position{}

	for _, po := range pos {
		//if this.GetBattleUnit(int64(po.InstId)) == nil {
		//	continue
		//}
		p := Position{}
		p.Position = po.Position
		p.InstId = po.InstId
		poss[po.InstId] = &p
	}

	battleplayer := this.GetBattlePlayer(this.BattleRoom)

	battleplayer.BattlePosition = poss

	this.TurnOver(this.BattleRoom)
	fmt.Println("SetupBattle end ", &battleplayer.BattlePosition )

	return nil
}

func (this *GamePlayer) SetBattleUnitOK(instId int64) error{			//返回卡牌上阵
	return nil
}

func (this *GamePlayer) GetBattlePlayer (battleroom int64) *BattlePlayer {		//获取玩家的战斗属性
	battleRoom, _ := BattleRoomList[this.BattleRoom]

	player := battleRoom.GetPlayer(this.MyUnit.InstId)
	return player
}


func (this *GamePlayer) TurnOver (battleroom int64)  {			//本回合结束
	battleRoom, _ := BattleRoomList[this.BattleRoom]

	battleRoom.TurnOver(this)

	return
}
