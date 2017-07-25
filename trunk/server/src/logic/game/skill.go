package game

import (
	"logic/prpc"
	"fmt"
)

type Skill struct {
	SkillID int32		//技能ID
	Crit int32			//暴击几率
	Damage int32		//伤害数值
	BuffList []int32	//buff列表
	CoolDown int32		//冷卻時間
	UseTime int32		//使用時間
	TargetNum int32		//目標個數
	TargetCamp int32	//目標陣營 我方 敵方
}

const(
	kCampMine = 1
	kCampOther = 2
)	// 目標類型 1我方 2敵方

func(this *Skill)Condition()bool{
	return this.checkUse()
} //能不能使用

func(this *Skill)Action(caster *GameUnit, targetList []*GameUnit, bout int32) (prpc.COM_BattleAction, []int64) {
	actionList := []prpc.COM_BattleActionTarget{}
	allDeat := []int64{}
	for i:=0; i<len(targetList); i++ {
		fmt.Println(i, "Action", targetList[i], "		")
		finl := int32(targetList[i].CProperties[prpc.CPT_HP]) - this.Damage
		if finl <= 0 {
			finl = 0
			allDeat = append(allDeat, targetList[i].InstId)
		}
		targetList[i].CProperties[prpc.CPT_HP] = float32(finl)
		t := prpc.COM_BattleActionTarget{}
		t.InstId = targetList[i].InstId
		t.ActionType = 1
		t.ActionParam = this.Damage
		t.ActionParamExt = finl
		actionList = append(actionList, t)
	}
	this.UseTime = bout

	action := prpc.COM_BattleAction{}
	action.InstId = caster.InstId
	action.SkillId = this.SkillID
	action.TargetList = actionList


	return action, allDeat

}//使用技能

func (this *Skill) StandbySkill(targetid int64, targetPlayer *GamePlayer) []*GameUnit {

	l := []*GameUnit{targetPlayer.GetBattleUnit(targetid)}
	index := 1
	for int32(index) < this.TargetNum {
		if index >= len(targetPlayer.BattleUnitList){
			break
		}
		uid := targetPlayer.BattleUnitList[index]			//按順序選擇
		l = append(l, targetPlayer.GetBattleUnit(uid))
		index ++
	}

	return l
}//技能準備 選定目標

func (this *Skill) checkUse() bool {
	return true
}//技能是否可用

func (this *Skill) RefreshBattle() {
	this.UseTime = 0
}//戰鬥開始時刷新技能


func InitSkillFromTable(SkillId int32) *Skill {
	t := GetSkillRecordById(SkillId)
	if t == nil {
		return nil
	}

	s := Skill{}

	s.SkillID = t.SkillID
	s.Crit = t.Crit
	s.Damage = t.Damage
	s.BuffList = t.BuffList
	s.CoolDown = t.CoolDown

	return &s
}