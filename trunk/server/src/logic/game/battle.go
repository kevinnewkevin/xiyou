package game

const(
	kIdle = 0 // 无效状态
	kUsed = 1 // 使用状态
)

type BattleRoom struct{
	Status 		int //战斗房间状态
	UnitList []*GameUnit //战斗单元列表
}

