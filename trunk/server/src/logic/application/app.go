package application

import (
	"fmt"
	"logic/game"

	"logic/socket"
	"net"
	"logic/log"
	"github.com/astaxie/beego/toolbox"
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
		log.Debug("LoadUnitTable %s ", err.Error())
		return
	}

	err = game.LoadSkillTable("../../../config/tables/skill.csv")
	if err != nil {
		log.Debug("LoadSkillTable %s ", err.Error())
		return
	}

	err = game.LoadBuffTable("../../../config/tables/buff.csv")
	if err != nil {
		log.Debug("LoadBuffTable %s ", err.Error())
		return
	}

	err = game.LoadBattleTable("../../../config/tables/Battle.csv")
	if err != nil {
		log.Debug("LoadBattleTable %s ", err.Error())
		return
	}

	err = game.LoadStoryChapterTable("../../../config/tables/HeroStroy.csv")
	if err != nil {
		log.Debug("LoadStoryTable %s ", err.Error())
		return
	}

	err = game.LoadSmallChapterTable("../../../config/tables/Checkpoint.csv")
	if err != nil {
		log.Debug("LoadSmallChapterTable %s ", err.Error())
		return
	}

	err = game.LoadItemTable("../../../config/tables/Item.csv")
	if err != nil {
		log.Debug("LoadItemTable", err.Error())
		return
	}

	err = game.LoadDropTable("../../../config/tables/Drop.csv")
	if err != nil {
		log.Debug("LoadDropTable %s ", err.Error())
		return
	}

	err = game.LoadPromoteTable("../../../config/tables/Strengthen.csv")
	if err != nil {
		log.Debug("LoadPromoteTable %s ", err.Error())
		return
	}
	err = game.LoadExpTable("../../../config/tables/Exp.csv")
	if err != nil {
		log.Debug("LoadExpTable %s ", err.Error())
		return
	}

	err = game.LoadTianTiTable("../../../config/tables/Ladder.csv")
	if err != nil {
		log.Debug("LoadTianTiTable %s ", err.Error())
		return
	}

	err = game.LoadRoleSkillTable("../../../config/tables/RoleSkill.csv")
	if err != nil {
		log.Debug("LoadRoleSkillTable %s ", err.Error())
		return
	}

	err = game.LoadRoleSkillUpdateTable("../../../config/tables/RoleSkillUpdate.csv")
	if err != nil {
		log.Debug("LoadRoleSkillUpdateTable %s ", err.Error())
		return
	}

	err = game.LoadShopTable("../../../config/tables/ShopData.csv")
	if err != nil {
		log.Debug("LoadShopTable %s ", err.Error())
		return
	}

	err = game.LoadCardPondTable("../../../config/tables/Cardclose.csv")
	if err != nil {
		log.Debug("LoadCardPondTable %s ", err.Error())
		return
	}
	err = game.LoadRaceTable("../../../config/tables/Race.csv")
	if err != nil {
		log.Debug("LoadRaceTable %s ", err.Error())
		return
	}

	err = game.LoadRobotTable("../../../config/tables/AI.csv")
	if err != nil {
		log.Debug("LoadRobotTable %s ", err.Error())
		return
	}

	game.InitLua("../../../config/scripts/")

	//game.InitGlobalLuaState()
	game.InitTianTi()
	game.InitGameTask()
	toolbox.StartTask()
	defer toolbox.StopTask()
	//game.TestPlayer()
	this.l, err = net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("0.0.0.0"),10999,"ipv4"})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	go func() {
		for {

			defer func() {

				if r := recover(); r != nil {
					log.Error("main panic %s",fmt.Sprint(r))
				}

			}()
			conn, err = this.l.AcceptTCP()
			if err != nil {
				log.Debug(err.Error())
				endRunning <- true
			}
			log.Info("Client connected %s ",conn.RemoteAddr().String())



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
