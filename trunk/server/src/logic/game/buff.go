package game

type Buff struct {
	Owner       *GameUnit	//挂在谁身上
	BuffId		int32
	Round		int32		//哪个回合上的
	BuffUntil	int32		//持续多久
	BuffType	int32		//buff类型 增益还是减益
	BuffKind	int32		//buff种类 有行动还是没行动 有行动就是类似每回合恢复血量或者每回合掉血 没行动就是增加个盾之类的
	Data		int32 		//数值 加血 掉血 护盾 可以为0
}

func (this *Buff) Update() {

}

func (this *Buff) IsOver(round int32) bool {
	return this.Round + this.BuffUntil >= round
}


