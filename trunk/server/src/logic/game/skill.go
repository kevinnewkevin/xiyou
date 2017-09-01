package game

import (
	"fmt"
	"logic/prpc"
)

type Skill struct {
	SkillID    int32   //技能ID
	Crit       int32   //暴击几率
	Damage     int32   //伤害数值
	BuffList   []int32 //buff列表
	CoolDown   int32   //冷卻時間
	UseTime    int32   //使用時間
	TargetNum  int32   //目標個數
	TargetCamp int32   //目標陣營 我方 敵方
	LuaScprit  string  //lua
}

const (
	kCampMine  = 1
	kCampOther = 2
) // 目標類型 1我方 2敵方

func (this *Skill) Condition() bool {
	return this.checkUse()
} //能不能使用

func (this *Skill) Action(caster *GameUnit, targetList []*GameUnit, bout int32) ([]prpc.COM_BattleActionTarget, bool) {
	actionList := []prpc.COM_BattleActionTarget{}
	OwnerDead := false
	for i := 0; i < len(targetList); i++ {
		fmt.Println(i, "Action", targetList[i], "		")
		finl := int32(targetList[i].CProperties[prpc.CPT_CHP]) - this.Damage
		if finl <= 0 {
			finl = 0
			if targetList[i].Owner.MyUnit.InstId == targetList[i].InstId {
				OwnerDead = true
			}
		}
		targetList[i].CProperties[prpc.CPT_CHP] = float32(finl)
		t := prpc.COM_BattleActionTarget{}
		t.InstId = targetList[i].InstId
		t.ActionType = 1
		t.ActionParam = this.Damage
		t.ActionParamExt = prpc.ToName_BattleExt(int(1))
		actionList = append(actionList, t)
	}
	this.UseTime = bout

	//this.ActionByLua(caster)

	return actionList, OwnerDead

} //使用技能

//临时函数
func SkillidToLuaName(skillid int32) string {
	SkillLua := GetSkillLuaRecordById(skillid)
	if SkillLua == nil {
		return "SK_1_Action"
	}
	return SkillLua.LuaScprit
}

func (this *Skill) ActionBylua(battleid int64, casterid int64) {
	//actionList := []prpc.COM_BattleActionTarget{}
	v := []interface{}{int(battleid), int(casterid)}
	r := []interface{}{0}


	_L.CallFuncEx(this.LuaScprit, v, &r)

	//this.ActionByLua(caster)

	return

} //使用技能

func (this *Skill) StandbySkill(targetid int64, targetPlayer *GamePlayer) []*GameUnit {

	l := []*GameUnit{targetPlayer.GetBattleUnit(targetid)}
	index := 1
	for int32(index) < this.TargetNum {
		if index >= len(targetPlayer.BattleUnitList) {
			break
		}
		uid := targetPlayer.BattleUnitList[index] //按順序選擇
		l = append(l, targetPlayer.GetBattleUnit(uid))
		index++
	}

	return l
} //技能準備 選定目標

func (this *Skill) checkUse() bool {
	return true
} //技能是否可用

func (this *Skill) RefreshBattle() {
	this.UseTime = 0
} //戰鬥開始時刷新技能

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
	s.LuaScprit = t.LuaScprit

	return &s
}

func TestActionByLua() {

	v := []interface{}{9999999999, 1}
	r := []interface{}{false}
	_L.CallFuncEx("SK_1_Action", v, &r)
	fmt.Println("TestActionByLua", r)

	return
}