package game

import (
	"fmt"
	"jimny/logs"

	"github.com/astaxie/beego/toolbox"
)

func PassZeroHourTask() error {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("main panic %s", fmt.Sprint(r))
		}

	}()
	OnlinePlayerPassZeroHour()
	return nil
}

func BlackMarkteRefreshTask() error {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("main panic %s", fmt.Sprint(r))
		}

	}()
	CheckMyBlackMarkte()
	return nil
}

func PlayerSave() error {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("main panic %s", fmt.Sprint(r))
		}

	}()
	Save()
	return nil
}

func RefreshTopList() error {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("main panic %s", fmt.Sprint(r))
		}

	}()
	//logs.Debug("RefreshTopList")
	RefreshAllTopList()
	RefreshFriendTopList()
	return nil
}

func InitGameTask() {
	passzerohourTimer := GetGlobalString("C_PassZeroHour")
	passzerohourTask := toolbox.NewTask("Passzerohour", passzerohourTimer, PassZeroHourTask)
	toolbox.AddTask("Passzerohour", passzerohourTask)

	blackMarketTimer := GetGlobalString("C_BlackMarkteRefresh")
	blackMarketTask := toolbox.NewTask("BlackMarkteRefresh", blackMarketTimer, BlackMarkteRefreshTask)
	toolbox.AddTask("BlackMarkteRefresh", blackMarketTask)

	playerSaveTimer := GetGlobalString("C_SaveDataToDB")
	playerSaveTask := toolbox.NewTask("PlayerSaveToDB", playerSaveTimer, PlayerSave)
	toolbox.AddTask("PlayerSaveToDB", playerSaveTask)

	TopListTimer := GetGlobalString("C_TopListRefresh")
	RefreshTopList := toolbox.NewTask("RefreshTopList", TopListTimer, RefreshTopList)
	toolbox.AddTask("RefreshTopList", RefreshTopList)

}
