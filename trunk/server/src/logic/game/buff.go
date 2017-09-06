package game

import (
	"sync/atomic"
	"fmt"
)

type Buff struct {
	Owner       *GameUnit		//挂在谁身上
	BuffId		int32			//基础id
	CasterId	int64			//buff释放者id
	InstId		int32			//buff实例ID
	Round		int32			//哪个回合上的
	BuffUntil	int32			//持续多久
	BuffType	int32			//buff类型 增益还是减益
	BuffKind	int32			//buff种类 有行动还是没行动 有行动就是类似每回合恢复血量或者每回合掉血 没行动就是增加个盾之类的
	Data		int32 			//数值 加血 掉血 护盾 可以为0
	Times		int32 			//buff生效的次数
	Over 		bool			//是否中断
	DataMap 	map[int]int32	//buff数值key是buff的idx,value是数值
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
	NewBuff.Times = b.Times

	return &NewBuff
}

func (this *Buff) AddProperty() {

	v := []interface{}{int(this.Owner.BattleId), int(this.Owner.InstId), int(this.InstId), int(this.Data)}
	r := []interface{}{0}
	buff_t := GetBuffRecordById(this.BuffId)
	fmt.Println("AddProperty", int(this.Owner.BattleId), this.Data, "buffID是", buff_t.BuffId)

	_L.CallFuncEx(buff_t.AddLua, v, &r)
}

func (this *Buff) DeleteProperty() {

	fmt.Println("DeleteProperty", this.Data, this.InstId)
	v := []interface{}{int(this.Owner.BattleId), int(this.Owner.InstId), int(this.InstId), int(this.Data)}
	r := []interface{}{0}
	this.Over = true

	buff_t := GetBuffRecordById(this.BuffId)

	fmt.Println(buff_t.PopLua, this.Owner.InstId, this.InstId)
	_L.CallFuncEx(buff_t.PopLua, v, &r)
}


func (this *Buff) Update(round int32) bool {
	fmt.Println("buff每回合更新 实例ID为:", this.InstId, "round is ", round, "myRound is ", this.Round)

	needDel := false

	if this.Round == round && this.BuffKind == kKindUntil{
		fmt.Println("本回合上的有结算的buff本回合不生效", this.Round, round, needDel)
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

		buff_t := GetBuffRecordById(this.BuffId)

		fmt.Println(buff_t.UpdateLua, int(this.Owner.BattleId), int(this.InstId), "是否需要删除", needDel, "unitID为:", this.Owner.InstId)
		_L.CallFuncEx(buff_t.UpdateLua, v, &r)
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

func (this *Buff) ChangeTimes() int {
	bf_table := GetBuffRecordById(this.BuffId)
	if bf_table == nil {
		return 1
	}

	if bf_table.Times == 0 {
		return 1
	}

	this.Times -= 1

	if this.Times <= 0 {
		this.Over = true
	}

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


