package game

import (
	"fmt"
	"sync/atomic"
	"logic/prpc"
)

const(
	kIdle = 0 // 无效状态
	kUsed = 1 // 使用状态

	kTurn = 1		//回合數
	kMaxUnit = 6	//雙方最多上陣卡牌
	kMaxMove = 2	//行动结束
)
var roomInstId int64 = 1
type BattleRoom struct{
	InstId      int64			//房间ID
	Status 		int32 			//战斗房间状态
	PlayerList []*BattlePlayer 	//房间中玩家信息
	Target		map[int64]*BattlePlayer 	//目標信息
	Bout 		int32			//回合
	TurnMove	int32			//行動次數
	NowPlayer	*BattlePlayer
	Turn		int32
	TurnPlayer	map[int64]int32	//是否移动
}

type BattlePlayer struct {
	Player *GamePlayer			//玩家信息
	MaxPoint int32				//最大能量點
	CurPoint int32				//當前能量點
	Camp	 int				//阵营
	BattlePosition map[int64]*Position //戰鬥位置信息
}

type Position struct{
	InstId int64  	//0
	Position int32  //1
}

var BattleRoomList = map[int64]*BattleRoom{}			//所有房间


////////////////////////////////////////////////////////////////////////
////创建部分
////////////////////////////////////////////////////////////////////////

func CreateBattlePlayer(player *GamePlayer) *BattlePlayer {
	p := BattlePlayer{}

	p.Player = player
	p.MaxPoint = 1
	p.CurPoint = 1
	p.BattlePosition = map[int64]*Position{}

	if player.session == nil {
		return nil
	}
	player.session.JoinBattleOk()
	fmt.Println(p.Player.session, "JoinBattleOk")

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
	Player1.Camp = 0
	player2.Camp = 1

	room.PlayerList = append(room.PlayerList, Player1)
	room.PlayerList = append(room.PlayerList, player2)
	war1.BattleRoom = room.InstId
	war2.BattleRoom = room.InstId

	room.Target = map[int64]*BattlePlayer{}
	room.Target[war1.MyUnit.InstId] = player2
	room.Target[war2.MyUnit.InstId] = Player1

	room.Turn = 0
	room.TurnPlayer = map[int64]int32{}

	room.NowPlayer = Player1

	BattleRoomList[room.InstId] = &room
	fmt.Println("CreateBattleRoom", &room)

	return &room
}

////////////////////////////////////////////////////////////////////////
////销毁部分
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) BattleRoomOver() {
	this.Status = kIdle
}

func (this *BattleRoom) DestoryBattleRoom() {
	this = nil
}

////////////////////////////////////////////////////////////////////////
////获取数据
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) GetPlayer(instId int64) *BattlePlayer {
	for _, v := range this.PlayerList {
		if v.Player.MyUnit.InstId == instId {
			return v
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////
////回合操作
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) TurnOver(Player *GamePlayer) {
	_, ok := this.TurnPlayer[Player.MyUnit.InstId]

	if !ok{
		this.Turn += 1
		this.TurnPlayer[Player.MyUnit.InstId] = 1
	}

	if this.Turn == 2 {
		//回合结束 并且重置Turn,向两方发送战报信息
		this.BattleAcction(this.NowPlayer, this.Target)
		this.Turn = 0
	}
	//this.BattleAcction(this.NowPlayer, this.Target)
}

func (this *BattleRoom) TurnGround(thisPlayer *GamePlayer) {
	// 操作切換
	this.TurnMove += 1
	if this.TurnMove % 2 == 0 {
		this.Bout += 1
	}
	this.NowPlayer = this.Target[thisPlayer.MyUnit.InstId]
}


func (this *BattleRoom) SendReport(report prpc.COM_BattleReport)  {
	for _, player := range this.PlayerList{
		player.Player.session.BattleReport(report)
	}
}

////////////////////////////////////////////////////////////////////////
////战斗过程
////////////////////////////////////////////////////////////////////////

func (this *BattleRoom) BattleAcction(Player *BattlePlayer, AllPlayer map[int64]*BattlePlayer) {
	fmt.Println("BattleAcction start", Player.Player.MyUnit.InstId)
	report := prpc.COM_BattleReport{}
	fmt.Println("BattleAcction report start",report.UnitList, "ssssss", report.ActionList)

	for _, bp := range AllPlayer {				//遍历所有参与战斗的
		for _, unit := range bp.BattlePosition {		//遍历本人所有战斗单位
			instid := unit.InstId
			//u := bp.Player.GetBattleUnit(instid)

			var u *GameUnit
			u = bp.Player.GetUnit(instid)

			if u == nil {
				continue
			}

			skill := u.Skill[0]

			targetPlayer, _ := this.Target[bp.Player.MyUnit.InstId]
			targetList := []*GameUnit{}

			for _, pos := range targetPlayer.BattlePosition {
				e_u := targetPlayer.Player.GetBattleUnit(pos.InstId)
				if e_u == nil{
					continue
				}
				targetList = append(targetList, e_u)
			}

			if len(targetList) == 0{
				targetList = append(targetList, targetPlayer.Player.MyUnit)
			}		//测试用代码

			unit_con := prpc.COM_BattleUnit{}

			unit_con.InstId = unit.InstId
			unit_con.Position = unit.Position
			unit_con.Camp = bp.Camp
			unit_con.UnitId = u.UnitId
			unit_con.Name = u.InstName
			unit_con.HP = int32(u.CProperties[prpc.CPT_HP])

			report.UnitList = append(report.UnitList, unit_con)

			fmt.Println("UnitId", unit_con.UnitId, "Instid ", unit_con.InstId)

			al, dl := skill.Action(u, targetList, this.Bout)

			if len(dl) > 0{
				//这里需要在战斗单位中删除某些角色
				for _, iid := range dl {
					delete(bp.BattlePosition, iid)
				}
				if len(bp.BattlePosition) == 0{
				}
			}

			fmt.Println("BattleAcction", bp.Player.MyUnit.InstId, "target", targetPlayer.Player.MyUnit.InstId, "acction", al)
			report.ActionList = append(report.ActionList, al)
			fmt.Println("Battle report", report)
		}
	}

	fmt.Println("Battle end, unitlist is ", report.UnitList, "		action is ", report.ActionList)
	this.SendReport(report)

}

