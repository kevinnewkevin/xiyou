package application

import (
	"fmt"
	"logic/game"

	"logic/socket"
	"net"
	"suzuki/logs"
	"time"
)

type App struct {
	l net.Listener
}

func (this *App) Run() {
	var (
		err        error
		conn       net.Conn
		endRunning = make(chan bool, 1)
	)
	logs.Init()

	err = game.LoadUnitTable("../../../config/tables/entity.csv")
	if err != nil {
		fmt.Println("LoadUnitTable", err.Error())
		return
	}

	err = game.LoadSkillTable("../../../config/tables/skill.csv")
	if err != nil {
		fmt.Println("LoadSkillTable", err.Error())
		return
	}

	err = game.LoadBuffTable("../../../config/tables/buff.csv")
	if err != nil {
		fmt.Println("LoadBuffTable", err.Error())
		return
	}

	err = game.LoadBattleTable("../../../config/tables/Battle.csv")
	if err != nil {
		fmt.Println("LoadBattleTable", err.Error())
		return
	}

	err = game.LoadStoryChapterTable("../../../config/tables/HeroStroy.csv")
	if err != nil {
		fmt.Println("LoadStoryTable", err.Error())
		return
	}

	err = game.LoadSmallChapterTable("../../../config/tables/Checkpoint.csv")
	if err != nil {
		fmt.Println("LoadSmallChapterTable", err.Error())
		return
	}

	err = game.LoadItemTable("../../../config/tables/Item.csv")
	if err != nil {
		fmt.Println("LoadItemTable", err.Error())
		return
	}

	err = game.LoadDropTable("../../../config/tables/Drop.csv")
	if err != nil {
		fmt.Println("LoadDropTable", err.Error())
		return
	}

	err = game.LoadPromoteTable("../../../config/tables/Strengthen.csv")
	if err != nil {
		fmt.Println("LoadPromoteTable", err.Error())
		return
	}
	err = game.LoadExpTable("../../../config/tables/Exp.csv")
	if err != nil {
		fmt.Println("LoadExpTable", err.Error())
		return
	}

	err = game.LoadTianTiTable("../../../config/tables/Ladder.csv")
	if err != nil {
		fmt.Println("LoadTianTiTable", err.Error())
		return
	}

	err = game.LoadRoleSkillTable("../../../config/tables/RoleSkill.csv")
	if err != nil {
		fmt.Println("LoadRoleSkillTable", err.Error())
		return
	}

	err = game.LoadRoleSkillUpdateTable("../../../config/tables/RoleSkillUpdate.csv")
	if err != nil {
		fmt.Println("LoadRoleSkillUpdateTable", err.Error())
		return
	}

	err = game.LoadShopTable("../../../config/tables/ShopData.csv")
	if err != nil {
		fmt.Println("LoadShopTable", err.Error())
		return
	}

	err = game.LoadCardPondTable("../../../config/tables/Cardclose.csv")
	if err != nil {
		fmt.Println("LoadCardPondTable", err.Error())
		return
	}

	game.InitLua("../../../config/scripts/")

	//game.InitGlobalLuaState()
	game.InitTianTi()
	//game.TestPlayer()
	this.l, err = net.Listen("tcp", "0.0.0.0:10999")
	if err != nil {
		fmt.Println(err.Error())
		return
	}


	go func() {
		for {
			conn, err = this.l.Accept()
			if err != nil {
				fmt.Println(err.Error())
				endRunning <- true
			}
			fmt.Println("Has one connect ")
			conn.SetDeadline(time.NewTimer(time.Second))
			peer := socket.NewPeer(conn)
			client := game.NewClient(peer)
			//
			go client.Update()
		}
	}()

	<-endRunning

	logs.Fini()
}

func NewApp() *App {
	return &App{}
}
