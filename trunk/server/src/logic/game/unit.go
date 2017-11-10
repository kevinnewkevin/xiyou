package game

import (
	"errors"
	"jimny/logs"
	"logic/prpc"
	"sync"
)


type GameUnit struct {
	sync.Mutex
	Owner       *GamePlayer //所有者
	IsMain      bool
	Camp        int
	UnitId      int32
	InstId      int64
	InstName    string
	DisPlay     int32
	Level       int32
	IProperties []int32
	CProperties []float32
	Cost        int32
	Race        int32
	Skill       map[int32]*Skill

	//战斗的实际信息
	ChoiceSKill int32
	Position    int32   //prpc.BattlePosition
	Buff        []*Buff //增益状态
	Debuff      []*Buff //负面状态
	Allbuff     []*Buff //全体buff
	DelBuff     []*Buff //需要刪除的buff
	BattleId    int64   //zhandou id
	MoveStage   int32   //行动信息
	//战斗 buff需要的数据
	VirtualHp int32             //护盾数值
	Special   map[int32][]int32 //特殊属性效果
	OutBattle bool              //脱离战斗
}

//如果是创建怪物卡牌的话 player = 你来
func CreateUnitFromTable(id int32) *GameUnit {
	t := GetUnitRecordById(id)
	if t == nil {
		return nil
	}
	u := GameUnit{}
	u.UnitId = t.Id
	u.IProperties = append(u.IProperties, t.IProp...)
	u.CProperties = append(u.CProperties, t.CProp...)
	u.DisPlay = t.DispId
	u.Skill = map[int32]*Skill{}
	u.Level = u.IProperties[prpc.IPT_PROMOTE]
	u.InstName = t.BaseName
	u.Cost = t.Cost
	u.Race = t.Race
	u.IsMain = false
	for i := 0; i < len(t.Skills); i++ {
		skill := InitSkillFromTable(t.Skills[i])
		if skill == nil {
			u.Skill[int32(i)] = nil
			continue
		}
		u.Skill[int32(i)] = skill
	}
	return &u
}

func (this *GameUnit) GetBattleCamp() int {

	return this.Camp
}

func (this *GameUnit) GetCProperty(id int32) float32 {
	if id <= prpc.CPT_MIN || id >= prpc.CPT_MAX {
		return 0
	}
	return this.CProperties[id]
}

func (this *GameUnit) GetIProperty(id int32) int32 {
	if id <= prpc.IPT_MIN || id >= prpc.IPT_MAX {
		return 0
	}
	return this.IProperties[id]
}

func (this *GameUnit) AddSpec(spec string, buffinstid int32) {
	spe := prpc.ToId_BuffSpecial(spec)
	bufflist, ok := this.Special[int32(spe)]
	if !ok {
		this.Special[int32(spe)] = []int32{buffinstid}
	} else {
		this.Special[int32(spe)] = append(bufflist, buffinstid)
	}
	logs.Debug("AddSpec %d %d", spe, this.Special[int32(spe)])
	return
}

func (this *GameUnit) PopSpec(spec string, buffinstid int32) {
	spe := prpc.ToId_BuffSpecial(spec)
	logs.Debug("PopSpec %d %d %d ", buffinstid, spe, this.Special)

	bufflist, ok := this.Special[int32(spe)]
	if ok {
		if len(bufflist) > 0 {
			delete(this.Special, int32(spe))
		} else {
			tmpList := []int32{}
			for _, buff := range bufflist {
				if buff == buffinstid {
					continue
				}
				tmpList = append(tmpList, buff)
			}
			this.Special[int32(spe)] = tmpList
		}
	}

	return
}

func (this *GameUnit) GetSpecial(spec string) []int32 { //获取对应sepce枚举对应的buffid 可能为空
	tmp := []int32{}
	spe := prpc.ToId_BuffSpecial(spec)
	bufflist, ok := this.Special[int32(spe)]
	if !ok {
		return tmp
	} else {
		for _, v := range bufflist {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

func (this *GameUnit) GetOneSpecial(spec string, round int32) int32 { //获取对应sepce枚举对应的实力id 可能为空
	var tmp int32
	spe := prpc.ToId_BuffSpecial(spec)
	bufflist, ok := this.Special[int32(spe)]
	if !ok {
		return tmp
	} else if len(bufflist) == 0 {
		return tmp
	}

	for _, buff_id := range bufflist {
		buff := this.SelectBuff(buff_id)
		if buff.IsOver(round) {
			continue
		}
		tmp = buff_id
		break
	}

	return tmp
}

func (this *GameUnit) CheckSpec(spec string, round int32) bool { //unit.checkspec(是否有免死)
	spe := prpc.ToId_BuffSpecial(spec)
	bufflist, ok := this.Special[int32(spe)]

	if !ok {
		return false
	}

	if len(bufflist) == 0 {
		return false
	}

	for _, bfid := range bufflist {
		buff := this.SelectBuff(bfid)
		if buff == nil {
			continue
		}

		if buff.IsOver(round) {
			continue
		}
		return true
	}

	return false
}

func (this *GameUnit) ClacSheldPer(round int32) float32 { //计算百分比减伤 所有buff的百分比减伤加起来 有个最大值
	maxPer := 75

	bl, ok := this.Special[prpc.BF_SHELD]

	if !ok || len(bl) == 0 {
		return 0
	}

	sheld := 0

	for _, instid := range bl {
		buff := this.SelectBuff(instid)
		if buff == nil || buff.IsOver(round) {
			continue
		}
		sheld += int(buff.Data)
	}

	if sheld > maxPer {
		sheld = maxPer
	}

	per := float32(sheld) / 100.0

	return per
}

func (this *GameUnit) ClacStrongPer(round int32) float32 { //计算百分比增加输出伤 所有buff的百分比减伤加起来 有个最大值
	maxPer := 75

	bl, ok := this.Special[prpc.BF_STRONG]

	if !ok || len(bl) == 0 {
		return 0
	}

	sheld := 0

	for _, instid := range bl {
		buff := this.SelectBuff(instid)
		if buff == nil || buff.IsOver(round) {
			continue
		}
		sheld += int(buff.Data)
	}

	if sheld > maxPer {
		sheld = maxPer
	}

	per := float32(sheld) / 100.0

	return per
}

func (this *GameUnit) ClacWeakPer(round int32) float32 { //计算百分比增加承受伤 所有buff的百分比减伤加起来 有个最大值
	maxPer := 75

	bl, ok := this.Special[prpc.BF_WEAK]

	if !ok || len(bl) == 0 {
		return 0
	}

	sheld := 0

	for _, instid := range bl {
		buff := this.SelectBuff(instid)
		if buff == nil || buff.IsOver(round) {
			continue
		}
		sheld += int(buff.Data)
	}

	if sheld > maxPer {
		sheld = maxPer
	}

	per := float32(sheld) / 100.0

	return per
}

func (this *GameUnit) SetUnitCOM(u *prpc.COM_Unit) {
	this.UnitId = u.UnitId
	this.InstId = u.InstId
	this.Level = u.Level
	this.Race = u.Race
	this.IProperties = u.IProperties
	this.CProperties = u.CProperties

	this.Skill = map[int32]*Skill{}
	for _, sk := range u.Skills {
		this.Skill[sk.Pos] = InitSkillFromTable(sk.SkillId)
	}
}

func (this *GameUnit) GetUnitCOM() prpc.COM_Unit {
	u := prpc.COM_Unit{}
	u.UnitId = this.UnitId
	u.InstId = this.InstId
	u.Level = this.Level
	u.Race = this.Race
	u.IProperties = append(u.IProperties, this.IProperties...)
	u.CProperties = append(u.CProperties, this.CProperties...)

	for idx, skill := range this.Skill {
		unit_skill := prpc.COM_UnitSkill{}
		unit_skill.Pos = idx
		if skill == nil {
			unit_skill.SkillId = 0
		} else {
			unit_skill.SkillId = skill.SkillID
		}

		u.Skills = append(u.Skills, unit_skill)
	}
	return u
}

func (this *GameUnit) GetBattleUnitCOM() prpc.COM_BattleUnit {
	u := prpc.COM_BattleUnit{}
	u.Position = this.Position
	u.InstId = this.InstId
	u.UnitId = this.UnitId
	u.HP = int32(this.GetCProperty(prpc.CPT_HP))
	u.CHP = int32(this.GetCProperty(prpc.CPT_CHP))
	u.Position = this.Position
	u.Name = this.InstName
	u.Level = this.IProperties[prpc.IPT_PROMOTE]

	return u
}

func (this *GameUnit) SelectSkill(round int32) *Skill {
	var idx int32
	if this.MoveStage > 2 {
		idx = this.MoveStage % 3
	} else {
		idx = this.MoveStage
	}
	this.MoveStage += 1

	if this.Skill[idx+1] == nil {
		for _, skill := range this.Skill {
			if skill == nil {
				continue
			}
			return skill
		}
	}

	return this.Skill[idx+1]
}

func (this *GameUnit) CastSkill(battle *BattleRoom) bool {
	if this.IsDead() {
		return false
	}

	var skill *Skill
	if this.IsMain {
		if this.ChoiceSKill == 0 {
			return false
		} else {
			for _, sk := range this.Skill {
				if sk == nil {
					continue
				}
				if sk.SkillID == this.ChoiceSKill {
					skill = sk
					break
				}
			}
			if skill == nil {
				return false
			}
		}
	} else {
		skill = this.SelectSkill(battle.Round)
	}

	logs.Debug("CastSkill ", skill)

	//logs.Info("CastSkill skill id is ", skill.SkillID)

	battle.AcctionList.SkillId = skill.SkillID

	skill.ActionBylua(battle.InstId, this.InstId)

	//logs.Info("CastSkill, AcctionList ", battle.AcctionList)

	return false
}

func (this *GameUnit) CastPassiveSkill(battle *BattleRoom) bool {
	skill := this.Skill[0]

	logs.Debug("CastPassiveSkill ", skill)
	if skill == nil {
		return false
	}

	battle.AcctionList.SkillId = skill.SkillID

	skill.ActionBylua(battle.InstId, this.InstId)

	return false
}

func (this *GameUnit) GetFrontList() []int {
	if this.GetBattleCamp() == prpc.CT_RED {
		return []int{prpc.BP_RED_1, prpc.BP_RED_2, prpc.BP_RED_3}
	} else {
		return []int{prpc.BP_BLUE_1, prpc.BP_BLUE_2, prpc.BP_BLUE_3}
	}
}

func (this *GameUnit) GetBackList() []int {
	if this.GetBattleCamp() == prpc.CT_RED {
		return []int{prpc.BP_RED_4, prpc.BP_RED_5, prpc.BP_RED_6}
	} else {
		return []int{prpc.BP_BLUE_4, prpc.BP_BLUE_5, prpc.BP_BLUE_6}
	}
}

func (this *GameUnit) IsDead() bool {
	return this.GetCProperty(prpc.CPT_CHP) <= 0
}

func (this *GameUnit) IsJump() bool {
	buff_lis, ok := this.Special[prpc.BF_JUMP]
	if !ok {
		return false
	}
	if len(buff_lis) == 0 {
		return false
	}
	return true
}

func (this *GameUnit) isFront() bool {
	li := this.GetFrontList()

	for _, v := range li {
		if v == int(this.Position) {
			return true
		}
	}
	return false
}

func (this *GameUnit) isBack() bool {
	li := this.GetBackList()

	for _, v := range li {
		if v == int(this.Position) {
			return true
		}
	}
	return false
}

func (this *GameUnit) ClearAllbuff() {

	//logs.Debug("ClearAllbuff, unitid: ", this.InstId)
	for _, buff := range this.Allbuff {
		buff.DeleteProperty()
	}

	return
}

func (this *GameUnit) ResetBattle(camp int, ismain bool, battleid int64) {
	this.CProperties[prpc.CPT_HP] = float32(this.IProperties[prpc.IPT_HP])
	this.CProperties[prpc.CPT_CHP] = float32(this.IProperties[prpc.IPT_HP])
	this.Buff = []*Buff{}
	this.Debuff = []*Buff{}
	this.Allbuff = []*Buff{}
	this.Camp = camp
	this.IsMain = ismain
	this.BattleId = battleid
	this.VirtualHp = 0
	this.Position = prpc.BP_MAX
	this.Special = map[int32][]int32{}
	this.MoveStage = 0
	this.OutBattle = false
}

func (this *GameUnit) CheckBuff(round int32) {
	//检测那些有行为的buff 比如定时增加血量的那种

}

func (this *GameUnit) CheckDebuff(round int32) {
	//检测那些有行为的debuff 比如定时损血

}
func (this *GameUnit) MustUpdateBuff(spe string, round int32) {
	special := prpc.ToId_BuffSpecial(spe)
	bufflist, _ := this.Special[int32(special)]

	for _, buffid := range bufflist {
		buff := this.SelectBuff(buffid)
		if buff.IsOver(round) {
			continue
		}
		buff.MustUpdate()
	}

}

func (this *GameUnit) SelectBuff(instid int32) *Buff {
	for _, buff := range this.Allbuff {
		if buff.InstId == instid {
			return buff
		}
	}

	return nil
}

func (this *GameUnit) CheckAllBuff(round int32) []int32 {
	logs.Debug(string(this.InstId), "checkallBuff round is ", round) //檢測buff效果
	needDelete := map[*Buff]int{}
	this.DelBuff = []*Buff{}

	for _, buff := range this.Allbuff {
		if this.IsDead() { //buff執行中玩家卡牌可能死掉
			break
		}
		if buff.Update(round) {
			logs.Info("CheckAllBuff one", buff.InstId, buff.Round)
			needDelete[buff] = 1
			this.DelBuff = append(this.DelBuff, buff) //這個是給戰鬥房間用的 用來寫入戰報
		}
	}

	need := this.deletBuff(needDelete)

	logs.Debug(string(this.InstId), "checkallBuff over 1", len(needDelete)) //檢測buff效果
	logs.Debug(string(this.InstId), "checkallBuff over 2", need)            //檢測buff效果

	return need
}

func (this *GameUnit) deletBuff(need map[*Buff]int) []int32 {
	newList := []*Buff{}
	delete_id := []int32{}
	for _, buff := range this.Allbuff {
		_, ok := need[buff]
		if ok {
			buff.DeleteProperty()
			if buff.BuffKind == kKindNow {
				delete_id = append(delete_id, buff.BuffId)
			}
			continue
		}
		newList = append(newList, buff)
	}

	logs.Debug("deletBuff", need)
	this.Allbuff = newList
	return delete_id
}

func erase(arr []interface{}, idx int) []interface{} {
	return append(arr[:idx], arr[idx+1:]...)
}

func (this *GameUnit) PopAllBuffByDebuff() int {
	//删除卡牌身上所有的debuff
	tmp := map[*Buff]int{}

	if len(this.Allbuff) == 0 || this.Allbuff == nil {
		return 0
	}

	logs.Debug("allbuff 1", this.Allbuff)
	for _, buff := range this.Allbuff {
		logs.Info("this buff", buff)
		if buff == nil {
			continue
		}
		if buff.BuffType == kTypeBuff {
			continue
		}
		tmp[buff] = 1
	}

	newBufflist := []*Buff{}

	for _, v := range this.Allbuff {
		_, ok := tmp[v]
		if ok {
			v.DeleteProperty()
			continue
		}
		newBufflist = append(newBufflist, v)
	}

	logs.Debug("PopAllBuffByDebuff")
	this.Allbuff = newBufflist
	logs.Debug("allbuff 2", this.Allbuff)
	logs.Debug(string(len(tmp)), tmp)
	return len(tmp)
}

func (this *GameUnit) PopAllBuffByBuff() int {
	//删除卡牌身上的buff
	tmp := map[*Buff]int{}

	for _, buff := range this.Allbuff {
		if buff.BuffType == kTypeDebuff {
			continue
		}
		tmp[buff] = 1
	}

	newBufflist := []*Buff{}

	for _, v := range this.Allbuff {
		_, ok := tmp[v]
		if ok {
			v.DeleteProperty()
			continue
		}

		newBufflist = append(newBufflist, v)
	}

	logs.Debug("PopAllBuffByBuff")
	this.Allbuff = newBufflist
	logs.Debug("allbuff 3", this.Allbuff)
	logs.Debug(string(len(tmp)), tmp)
	return len(tmp)
}

func (this *GameUnit) ClearBuffByDead(ownerid int64, buffids []int32) {

}

func (this *GameUnit) PopAllBuffByDead(battle *BattleRoom) {
	if this.UnitId != 19 {
		return
	}
	for _, unit := range battle.Units {
		if unit == nil {
			continue
		}

		if unit.IsDead() {
			continue
		}

		if unit.Camp != this.Camp {
			continue
		}

		for _, buff := range unit.Allbuff {
			if buff.CasterId == this.InstId {
				buff.Over = true
			}
		}
	}

}

func (this *GameUnit) ChangeBuffTimes(round int32) {

	for _, buff := range this.Allbuff {
		if buff == nil {
			continue
		}
		if buff.IsOver(round) {
			continue
		}
		if buff.Times == 0 {
			continue
		}
		buff.ChangeTimes()
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *GameUnit) UpdateIProperty(iType int32, value int32) error {

	if iType <= prpc.IPT_MIN || iType >= prpc.IPT_MAX {
		return errors.New("error iType")
	}

	logs.Debug("UpdateIProperty, itype", iType, "front pro ", this.IProperties[iType])
	this.IProperties[iType] += value
	logs.Debug("UpdateIProperty, itype", iType, "after pro ", this.IProperties[iType])

	if this.Owner.session == nil {
		return nil
	}
	this.Owner.session.UpdateUnitIProperty(this.InstId, iType, this.IProperties[iType])

	return nil
}

func (this *GameUnit) UpdateCProperty(cType int32, value float32) error {

	if cType <= prpc.CPT_MIN || cType >= prpc.CPT_MAX {
		return errors.New("error cType")
	}

	logs.Debug("UpdateCProperty, cType", cType, "front pro ", this.CProperties[cType])
	this.CProperties[cType] += value
	logs.Debug("UpdateCProperty, cType", cType, "after pro ", this.CProperties[cType])

	if this.Owner.session == nil {
		return nil
	}
	this.Owner.session.UpdateUnitCProperty(this.InstId, cType, this.CProperties[cType])

	return nil
}
func (this *GameUnit) SetIProperty(iType int32, value int32) error {

	if iType <= prpc.IPT_MIN || iType >= prpc.IPT_MAX {
		return errors.New("error iType")
	}

	if value < 0 {
		value = 0
	}

	logs.Debug("SetIProperty, itype", iType, "front pro ", this.IProperties[iType])
	this.IProperties[iType] = value
	logs.Debug("SetIProperty, itype", iType, "after pro ", this.IProperties[iType])

	if this.Owner.session == nil {
		return nil
	}

	this.Owner.session.UpdateUnitIProperty(this.InstId, iType, value)

	return nil
}

func (this *GameUnit) SetCProperty(cType int32, value float32) error {

	if cType <= prpc.CPT_MIN || cType >= prpc.CPT_MAX {
		return errors.New("error cType")
	}

	logs.Debug("SetCProperty, cType", cType, "front pro ", this.CProperties[cType])
	this.CProperties[cType] = value
	logs.Debug("SetCProperty, cType", cType, "after pro ", this.CProperties[cType])

	this.Owner.session.UpdateUnitCProperty(this.InstId, cType, value)

	return nil
}

func (this *GameUnit) CheckExp(exp int32) int32 {
	logs.Debug("CheckExp in", exp)
	if this.Owner == nil {
		return 0
	}

	if this.Owner.MyUnit.InstId != this.InstId {
		return 0
	}

	exp_info := GetExpRecordById(this.IProperties[prpc.IPT_PROMOTE])
	if exp_info == 0 {
		return 0
	}
	if exp_info > exp {
		return exp
	}

	for {
		if exp_info > exp {
			break
		}
		promote := GetPromoteRecordById(this.UnitId)
		logs.Debug("this.IProperties[prpc.IPT_PROMOTE]", this.IProperties[prpc.IPT_PROMOTE])
		logs.Debug("this.Promote", promote[this.IProperties[prpc.IPT_PROMOTE]-1])
		this.Promote(promote[this.IProperties[prpc.IPT_PROMOTE]-1])

		exp -= exp_info
		exp_info = GetExpRecordById(this.IProperties[prpc.IPT_PROMOTE])
		this.Level += 1
		this.Owner.CheckSkillBase()
	}

	logs.Debug("CheckExp out final", exp)

	return exp
}

func (this *GameUnit) Promote(info *PromoteInfo) error {

	//this.Level = info.Level
	//this.IProperties[prpc.IPT_HP] += info.Hp
	//this.CProperties[prpc.CPT_HP] += float32(info.Hp)
	//this.CProperties[prpc.CPT_ATK] += info.ATK
	//this.CProperties[prpc.CPT_DEF] += info.DEF
	//this.CProperties[prpc.CPT_MAGIC_ATK] += info.MATK
	//this.CProperties[prpc.CPT_MAGIC_DEF] += info.MDEF
	//this.CProperties[prpc.CPT_AGILE] += info.AGILE

	this.SetIProperty(prpc.IPT_PROMOTE, info.Level)
	this.UpdateIProperty(prpc.IPT_HP, info.Hp)
	this.UpdateCProperty(prpc.CPT_HP, float32(info.Hp))
	this.UpdateCProperty(prpc.CPT_ATK, info.ATK)
	this.UpdateCProperty(prpc.CPT_DEF, info.DEF)
	this.UpdateCProperty(prpc.CPT_MAGIC_ATK, info.MATK)
	this.UpdateCProperty(prpc.CPT_MAGIC_DEF, info.MDEF)
	this.UpdateCProperty(prpc.CPT_AGILE, info.AGILE)

	return nil
}
