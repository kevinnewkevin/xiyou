package game

import (
	"sync/atomic"
	"fmt"
)

type Buff struct {
	Owner       *GameUnit	//挂在谁身上
	BuffId		int32		//基础id
	CasterId	int64		//buff释放者id
	InstId		int32		//buff实例ID
	Round		int32		//哪个回合上的
	BuffUntil	int32		//持续多久
	BuffType	int32		//buff类型 增益还是减益
	BuffKind	int32		//buff种类 有行动还是没行动 有行动就是类似每回合恢复血量或者每回合掉血 没行动就是增加个盾之类的
	Data		int32 		//数值 加血 掉血 护盾 可以为0
	Over 		bool		//是否中断
}

const (
	kTypeBuff = 1			//增益buff
	kTypeDebuff = 2			//减益buff

	kKindUntil = 1			//有行为的buff
	kKindNow = 2			//无行为的buff
)

var BuffInstId int32 = 1

func NewBuff(owner *GameUnit, casterid int64, buffid int32, data int32, round int32) *Buff {

	b := GetBuffRecordById(buffid)
	if b == nil {
		return nil
	}

	NewBuff := Buff{}

	NewBuff.BuffId = buffid
	NewBuff.CasterId = casterid
	NewBuff.InstId = atomic.AddInt32(&BuffInstId, 1)
	NewBuff.Owner = owner
	NewBuff.Round = round
	NewBuff.Data = data
	NewBuff.BuffUntil = b.Until
	NewBuff.BuffType = b.Type
	NewBuff.BuffKind = b.Kind
	NewBuff.Over = false

	return &NewBuff
}

func (this *Buff) AddProperty() {
	if this.BuffKind == kKindUntil {		// 只有直接增加属性的BUFF才会走到这里
		return
	}
	fmt.Println("AddProperty", int(this.Owner.BattleId), this.Data)
	v := []interface{}{int(this.Owner.BattleId), int(this.Owner.InstId), int(this.InstId)}
	r := []interface{}{0}

	fmt.Println("buff_1_add", int(this.Owner.BattleId), int(this.Data))
	_L.CallFuncEx("buff_1_add", v, &r)
}

func (this *Buff) DeleteProperty(battleud int64, unitid int64) {
	if this.BuffKind == kKindUntil {
		return
	}
	fmt.Println("DeleteProperty", int(this.Owner.BattleId), this.Data)
	v := []interface{}{int(this.Owner.BattleId), int(this.InstId), int(this.Owner.InstId)}
	r := []interface{}{0}

	fmt.Println("buff_1_delete", int(this.Owner.BattleId), int(unitid), int(this.Data))
	_L.CallFuncEx("buff_1_delete", v, &r)
}


func (this *Buff) Update(round int32) bool {
	fmt.Println("buff每回合更新 实例ID为:", this.InstId)

	needDel := false

	if this.Round == round {
		fmt.Println("本回合上的buff本回合不生效", this.Round, "   ", round)
		return needDel
	}

	if this.IsOver(round) {			//buff結束需要刪除
		fmt.Println("本buff到期 需要删除")
		needDel = true
	}

	if this.BuffKind == kKindNow {		//沒有行為的不需要進行結算
		fmt.Println("本buff不需要行为")
		return needDel
	} else if this.BuffKind == kKindUntil {
		v := []interface{}{int(this.Owner.BattleId), int(this.InstId), int(this.Owner.InstId)}
		r := []interface{}{0}

		fmt.Println("buff_1_update", int(this.Owner.BattleId), int(this.InstId), int(this.Owner.InstId))
		_L.CallFuncEx("buff_1_update", v, &r)
		//testBattleBuff(this, this.IsOver(round))

		return needDel
	} else {
		return needDel
	}
}

func (this *Buff) IsOver(round int32) bool {
	o := false

	if this.Over {
		o = true
	} else if this.Round + this.BuffUntil <= round {
		o = true
	}

	return o
}

func (this *Buff) Special(data int32) int {
	// 处理特殊属性

	return 1
}

func testBattleBuff(buff *Buff, over bool) {
	fmt.Println("testBattleBuff 1, buffid:", buff.BuffId, "over", over)
	fmt.Println(buff.Owner.BattleId)
	battle := FindBattle(buff.Owner.BattleId)

	if battle == nil {
		return
	}
	fmt.Println("testBattleBuff 2")

	var o bool

	if over {
		o = false
	} else {
		o = true
	}

	battle.BuffMintsHp(buff.CasterId, buff.Owner.InstId, buff.BuffId, buff.Data, o)
}


