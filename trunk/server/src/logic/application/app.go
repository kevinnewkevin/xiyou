package application

import (
	"fmt"
	"logic/game"

	"jimny/logs"
	"jimny/network"
	"net"
	"time"

	"github.com/astaxie/beego/toolbox"

	_ "net/http/pprof"
	"net/http"
	"runtime"
)

type App struct {
	l *net.TCPListener
}

func (this *App) Run() {
	var (
		err        error
		endRunning = make(chan bool, 1)
	)

	err = game.LoadUnitTable("../../../config/tables/entity.csv")
	if err != nil {
		logs.Debug("LoadUnitTable %s ", err.Error())
		return
	}

	err = game.LoadSkillTable("../../../config/tables/skill.csv")
	if err != nil {
		logs.Debug("LoadSkillTable %s ", err.Error())
		return
	}

	err = game.LoadBuffTable("../../../config/tables/buff.csv")
	if err != nil {
		logs.Debug("LoadBuffTable %s ", err.Error())
		return
	}

	err = game.LoadBattleTable("../../../config/tables/Battle.csv")
	if err != nil {
		logs.Debug("LoadBattleTable %s ", err.Error())
		return
	}

	err = game.LoadStoryChapterTable("../../../config/tables/HeroStroy.csv")
	if err != nil {
		logs.Debug("LoadStoryTable %s ", err.Error())
		return
	}

	err = game.LoadSmallChapterTable("../../../config/tables/Checkpoint.csv")
	if err != nil {
		logs.Debug("LoadSmallChapterTable %s ", err.Error())
		return
	}

	err = game.LoadItemTable("../../../config/tables/Item.csv")
	if err != nil {
		logs.Debug("LoadItemTable", err.Error())
		return
	}

	err = game.LoadDropTable("../../../config/tables/Drop.csv")
	if err != nil {
		logs.Debug("LoadDropTable %s ", err.Error())
		return
	}

	err = game.LoadPromoteTable("../../../config/tables/Strengthen.csv")
	if err != nil {
		logs.Debug("LoadPromoteTable %s ", err.Error())
		return
	}
	err = game.LoadExpTable("../../../config/tables/Exp.csv")
	if err != nil {
		logs.Debug("LoadExpTable %s ", err.Error())
		return
	}

	err = game.LoadTianTiTable("../../../config/tables/Ladder.csv")
	if err != nil {
		logs.Debug("LoadTianTiTable %s ", err.Error())
		return
	}

	err = game.LoadRoleSkillTable("../../../config/tables/RoleSkill.csv")
	if err != nil {
		logs.Debug("LoadRoleSkillTable %s ", err.Error())
		return
	}

	err = game.LoadRoleSkillUpdateTable("../../../config/tables/RoleSkillUpdate.csv")
	if err != nil {
		logs.Debug("LoadRoleSkillUpdateTable %s ", err.Error())
		return
	}

	err = game.LoadShopTable("../../../config/tables/ShopData.csv")
	if err != nil {
		logs.Debug("LoadShopTable %s ", err.Error())
		return
	}

	err = game.LoadCardPondTable("../../../config/tables/Cardclose.csv")
	if err != nil {
		logs.Debug("LoadCardPondTable %s ", err.Error())
		return
	}
	err = game.LoadRaceTable("../../../config/tables/Race.csv")
	if err != nil {
		logs.Debug("LoadRaceTable %s ", err.Error())
		return
	}

	err = game.LoadRobotTable("../../../config/tables/AI.csv")
	if err != nil {
		logs.Debug("LoadRobotTable %s ", err.Error())
		return
	}

	err = game.LoadGiveGiftTable("../../../config/tables/Givegift.csv")
	if err != nil {
		logs.Debug("LoadGiveGiftTable %s ", err.Error())
		return
	}

	game.InitLua("../../../config/scripts/")

	//game.InitGlobalLuaState()
	game.InitTianTi()
	game.InitGameTask()
	game.InitDB()
	game.InitTopList()
	game.InitGuild()
	toolbox.StartTask()
	defer toolbox.StopTask()
	//game.TestPlayer()

	acceptor := network.NewAcceptorC("tcp", "0.0.0.0:10999")

	go func() {
		http.ListenAndServe("0.0.0.0:16060", nil)
	}()

	go func() {

		for {
			defer func() {
				if r := recover(); r != nil {
					logs.Error("main panic %s", fmt.Sprint(r))
				}

			}()

			select {
			case conn := <-acceptor.Accept():
				game.NewClient(conn)
			default:
				game.TickClient()
				runtime.GC()
			}

			//conn, err = this.l.AcceptTCP()
			//if err != nil {
			//	logs.Debug(err.Error())
			//	endRunning <- true
			//}
			//logs.Info("Client connected %s ",conn.RemoteAddr().String())
			//
			//
			//
			//
			//

			time.Sleep(1)
			//go client.Update()
		}
	}()

	<-endRunning
}

func NewApp() *App {
	return &App{}
}
