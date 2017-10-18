package application

import (
	"fmt"
	"logic/game"

	"logic/socket"
	"net"
	"logic/std"
)

type App struct {
	l *net.TCPListener
}

func (this *App) Run() {
	var (
		err        error
		conn       *net.TCPConn
		endRunning = make(chan bool, 1)
	)

	err = game.LoadUnitTable("../../../config/tables/entity.csv")
	if err != nil {
		std.LogDebug("LoadUnitTable %s ", err.Error())
		return
	}

	err = game.LoadSkillTable("../../../config/tables/skill.csv")
	if err != nil {
		std.LogDebug("LoadSkillTable %s ", err.Error())
		return
	}

	err = game.LoadBuffTable("../../../config/tables/buff.csv")
	if err != nil {
		std.LogDebug("LoadBuffTable %s ", err.Error())
		return
	}

	err = game.LoadBattleTable("../../../config/tables/Battle.csv")
	if err != nil {
		std.LogDebug("LoadBattleTable %s ", err.Error())
		return
	}

	err = game.LoadStoryChapterTable("../../../config/tables/HeroStroy.csv")
	if err != nil {
		std.LogDebug("LoadStoryTable %s ", err.Error())
		return
	}

	err = game.LoadSmallChapterTable("../../../config/tables/Checkpoint.csv")
	if err != nil {
		std.LogDebug("LoadSmallChapterTable %s ", err.Error())
		return
	}

	err = game.LoadItemTable("../../../config/tables/Item.csv")
	if err != nil {
		std.LogDebug("LoadItemTable", err.Error())
		return
	}

	err = game.LoadDropTable("../../../config/tables/Drop.csv")
	if err != nil {
		std.LogDebug("LoadDropTable %s ", err.Error())
		return
	}

	err = game.LoadPromoteTable("../../../config/tables/Strengthen.csv")
	if err != nil {
		std.LogDebug("LoadPromoteTable %s ", err.Error())
		return
	}
	err = game.LoadExpTable("../../../config/tables/Exp.csv")
	if err != nil {
		std.LogDebug("LoadExpTable %s ", err.Error())
		return
	}

	err = game.LoadTianTiTable("../../../config/tables/Ladder.csv")
	if err != nil {
		std.LogDebug("LoadTianTiTable %s ", err.Error())
		return
	}

	err = game.LoadRoleSkillTable("../../../config/tables/RoleSkill.csv")
	if err != nil {
		std.LogDebug("LoadRoleSkillTable %s ", err.Error())
		return
	}

	err = game.LoadRoleSkillUpdateTable("../../../config/tables/RoleSkillUpdate.csv")
	if err != nil {
		std.LogDebug("LoadRoleSkillUpdateTable %s ", err.Error())
		return
	}

	err = game.LoadShopTable("../../../config/tables/ShopData.csv")
	if err != nil {
		std.LogDebug("LoadShopTable %s ", err.Error())
		return
	}

	err = game.LoadCardPondTable("../../../config/tables/Cardclose.csv")
	if err != nil {
		std.LogDebug("LoadCardPondTable %s ", err.Error())
		return
	}

	game.InitLua("../../../config/scripts/")

	//game.InitGlobalLuaState()
	game.InitTianTi()
	//game.TestPlayer()
	this.l, err = net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("0.0.0.0"),10999,"ipv4"})
	if err != nil {
		std.LogFatal(err.Error())
		return
	}


	go func() {
		for {
			defer func() {

				if r := recover(); r != nil {
					std.LogError("main panic %s",fmt.Sprint(r))
				}

			}()
			conn, err = this.l.AcceptTCP()
			if err != nil {
				fmt.Println(err.Error())
				endRunning <- true
			}
			std.LogDebug("Has one connect ")

			peer := socket.NewPeer(conn)
			client := game.NewClient(peer)
			//


			go client.Update()
		}
	}()

	<-endRunning


}

func NewApp() *App {
	return &App{}
}
