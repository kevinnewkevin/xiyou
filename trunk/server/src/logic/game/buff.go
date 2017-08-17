package game

import (
	"sync/atomic"
	"fmt"
)

type Buff struct {
	Owner       *GameUnit	//挂在谁身上
	BuffId		int32		//基础id
	CasterId	int64		//基础id
	InstId		int32
	Round		int32		//哪个回合上的
	BuffUntil	int32		//持续多久
	BuffType	int32		//buff类型 增益还是减益
	BuffKind	int32		//buff种类 有行动还是没行动 有行动就是类似每回合恢复血量或者每回合掉血 没行动就是增加个盾之类的
	Data		int32 		//数值 加血 掉血 护盾 可以为0
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

	return &NewBuff
}


func (this *Buff) Update(round int32) bool {
	fmt.Println("buff Update, buff id is ", this.BuffId, this.BuffKind == kKindNow, this.BuffKind == kKindUntil)

	if this.Round == round {
		return false
	}

	if this.BuffKind == kKindNow {		//沒有行為的不需要進行結算
		return false
	} else if this.BuffKind == kKindUntil {
		if this.IsOver(round) {			//buff結束需要刪除
			return true
		}
		//v := []interface{}{int(this.Owner.Owner.BattleId),int(this.BuffId), int(this.InstId)}
		//r := []interface{}{0}
		//
		//
		//_L.CallFuncEx("", v, &r)
		fmt.Println("11111")
		testBattleBuff(this)

		return false
	} else {
		return false
	}
}

func (this *Buff) IsOver(round int32) bool {
	return this.Round + this.BuffUntil <= round
}

func testBattleBuff(buff *Buff) {
	fmt.Println("testBattleBuff 1, buffid:", buff.BuffId)
	fmt.Println(buff.Owner.BattleId)
	battle := FindBattle(buff.Owner.BattleId)

	if battle == nil {
		return
	}
	fmt.Println("testBattleBuff 2")

	battle.BuffMintsHp(buff.CasterId, buff.Owner.InstId, buff.Data)
}


