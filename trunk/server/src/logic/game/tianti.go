package game

import (
	"time"
	"fmt"
	"logic/prpc"
)

const (
	timer		= 100			//计时器更新间隔
)

type (
	OncePlayer struct {
		PlayerInstId	int64
		TianTiVal		int32
		MatchingTime	float64
	}
)

var TianTiStore		[]*OncePlayer

func InitTianTi()  {
	go func() {
		t1 := time.NewTimer(time.Millisecond * timer)

		for {
			select {
			case <-t1.C:
				//fmt.Println("50ms timer")
				Tick(timer)
				t1.Reset(time.Millisecond * timer)
			}
		}
	}()
}

func Tick(dt float64)  {
	var TempSecond	float64
	TempSecond = dt/1000
	for i:=0;i<len(TianTiStore) ;i++  {
		if TianTiStore[i]==nil {
			continue
		}
		CheckMatching(TianTiStore[i],TempSecond)
	}

	PlayerTick(TempSecond)
}

func CheckMatching(oncePlayer *OncePlayer, dt float64)  {
	oncePlayer.MatchingTime += dt
	for _,t := range TianTiStore{
		if t.PlayerInstId == oncePlayer.PlayerInstId {
			continue
		}
		if oncePlayer.MatchingTime > 300{
			player := FindPlayerByInstId(oncePlayer.PlayerInstId)
			StopMatching(player)

			if player.session != nil {
				player.session.ErrorMessage(prpc.EN_TIANTI_MATCHING_TIMEOUT)
			}
			
			fmt.Println("Matching TimeOut ----> RemoveMatching",oncePlayer.PlayerInstId)
			return
		}

		tempV := (int32(oncePlayer.MatchingTime/30) +1)*50

		if oncePlayer.TianTiVal >= (t.TianTiVal - tempV) && oncePlayer.TianTiVal <= (t.TianTiVal + tempV) {
			//fmt.Println("Matching InstId=",oncePlayer.PlayerInstId,"MyTiantiVal",oncePlayer.TianTiVal,"tempV=",tempV,"[",(t.TianTiVal - tempV),(t.TianTiVal + tempV),"]","MatchingTime",
			//	oncePlayer.MatchingTime)
			myself := FindPlayerByInstId(oncePlayer.PlayerInstId)
			rival  := FindPlayerByInstId(t.PlayerInstId)
			RemoveMatching(oncePlayer.PlayerInstId)
			RemoveMatching(t.PlayerInstId)
			if CreatePvP(myself,rival) != nil {
				fmt.Println("Matching Succeed")
			}else {
				fmt.Println("Tianti CreatePvP Loser",oncePlayer.PlayerInstId,t.PlayerInstId)
			}

		}
	}
}

func StartMatching(player *GamePlayer,groupId int32)  {
	if player==nil {
		return
	}

	if player.GetUnitGroupById(groupId) == nil {
		fmt.Println("Can Not Find UnitGroup GroupId=",groupId)
		return
	}

	player.BattleUnitGroup = groupId

	tmp := OncePlayer{}
	tmp.PlayerInstId = player.MyUnit.InstId
	tmp.TianTiVal	 = player.TianTiVal
	TianTiStore = append(TianTiStore,&tmp)

	fmt.Println("StartMatching OK InstId=",tmp.PlayerInstId,"TianTiVal=",tmp.TianTiVal)
}

func StopMatching(player *GamePlayer)  {
	if player==nil {
		return
	}
	if RemoveMatching(player.MyUnit.InstId) {
		player.BattleUnitGroup = 0
	}
}

func RemoveMatching(instId int64) bool {
	for i:=0;i<len(TianTiStore) ;i++  {
		if instId == TianTiStore[i].PlayerInstId {
			TianTiStore = append(TianTiStore[:i], TianTiStore[i+1:]...)
			fmt.Println("RemoveMatching...",instId)
			return true
		}
	}
	return false
}

func CaleTianTiVal(player1 *GamePlayer,player2 *GamePlayer,winCamp int) int32 {
	if player1 == nil || player2 == nil {
		return 0
	}
	coef := int32((player1.TianTiVal - player2.TianTiVal)/5)

	if player1.BattleCamp == winCamp {
		player1.TianTiVal += (30-coef*2)
	}else {
		if player1.TianTiVal > 400 && player1.TianTiVal <= 1000 {
			player1.TianTiVal = player1.TianTiVal - (15-coef)
		}else if player1.TianTiVal > 1000 {
			player1.TianTiVal = player1.TianTiVal - (30-coef*2)
		}
	}
	if player1.session != nil {
		player1.session.UpdateTiantiVal(player1.TianTiVal)
	}

	tableId := GetTianTiIdByVal(player1.TianTiVal)
	ttData := GetTianTiTableDataById(tableId)
	if ttData == nil {
		fmt.Println("Can Not Find TianTiTableData By TableId=",tableId)
		return 0
	}

	var dropId int32 = 0;

	if player1.BattleCamp == winCamp {
		dropId = ttData.WinDrop
		fmt.Println("Tianti Battle Over CaleVal Winer Player[",player1.MyUnit.InstId,"]","TianTiVal[",player1.TianTiVal,"]","DropId=",ttData.WinDrop)
	}else {
		dropId = ttData.LoseDop
		fmt.Println("Tianti Battle Over CaleVal Loser Player[",player1.MyUnit.InstId,"]","TianTiVal[",player1.TianTiVal,"]","DropId=",ttData.LoseDop)
	}

	return dropId
}