package game

import (
	"fmt"
	"sync/atomic"
)

const(
	kIdle = 0 // 无效状态
	kUsed = 1 // 使用状态

	kTurn = 1		//回合數
	kMaxUnit = 6	//雙方最多上陣卡牌
)
var roomInstId int64 = 1
type BattleRoom struct{
	InstId      int64			//房间ID
	Status 		int 			//战斗房间状态
	PlayerList []*BattlePlayer 	//房间中玩家信息
	Target		map[int64]*BattlePlayer 	//目標信息
	Bout 		int32			//回合
	TurnMove	int32			//行動次數
	NextPlayer	*BattlePlayer
}

type BattlePlayer struct {
	Player *GamePlayer			//玩家信息
	MaxPoint int				//最大能量點
	CurPoint int				//當前能量點
	BattlePosition []*Position 	//戰鬥位置信息
}

type Position struct{
	InstId int64  	//0
	Position int32  //1
}

var BattleRoomList = map[int64]*BattleRoom{}			//所有房间

func CreateBattlePlayer(player *GamePlayer) *BattlePlayer {
	p := BattlePlayer{}

	p.Player = player
	p.MaxPoint = 1
	p.CurPoint = 1

	player.session.JoinBattleOk()

	return &p
}

func CreateBattleRoom(war1 *GamePlayer, war2 *GamePlayer) *BattleRoom {
	room := BattleRoom{}
	room.Status = kUsed
	room.InstId = atomic.AddInt64(&roomInstId, 1)
	room.Bout = kTurn
	room.TurnMove = 0

	Player1 := CreateBattlePlayer(war1)
	player2 := CreateBattlePlayer(war2)

	room.PlayerList = append(room.PlayerList, Player1)
	room.PlayerList = append(room.PlayerList, player2)
	war1.BattleRoom = room.InstId
	war2.BattleRoom = room.InstId

	room.Target = map[int64]*BattlePlayer{}
	room.Target[war1.MyUnit.InstId] = player2
	fmt.Println("Player1 target", player2)
	room.Target[war2.MyUnit.InstId] = Player1
	fmt.Println("Player2 target", Player1)

	room.NextPlayer = player2

	BattleRoomList[room.InstId] = &room
	fmt.Println("CreateBattleRoom", &room)

	return &room
}

func (this *BattleRoom) BattleRoomOver() {
	this.Status = kIdle
}
func (this *BattleRoom) GetPlayer(instId int64) *BattlePlayer {
	for _, v := range this.PlayerList {
		if v.Player.MyUnit.InstId == instId {
			return v
		}
	}
	return nil
}

func (this *BattleRoom) DestoryBattleRoom() {
	this = nil
}
func (this *BattleRoom) CheckPlayerMove(Player *GamePlayer) bool {

	if this.NextPlayer.Player.MyUnit.InstId == Player.MyUnit.InstId {
		return false
	}

	return true
}

func (this *BattleRoom) TurnGround(thisPlayer *GamePlayer) {
	// 操作切換
	this.TurnMove += 1
	if this.TurnMove % 2 == 0 {
		this.Bout += 1
	}
	this.NextPlayer = this.Target[thisPlayer.MyUnit.InstId]
}
