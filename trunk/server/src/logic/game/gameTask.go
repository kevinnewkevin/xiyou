package game

import (
	"github.com/astaxie/beego/toolbox"
)

func PassZeroHourTask() error {
	OnlinePlayerPassZeroHour()
	return nil
}

func BlackMarkteRefreshTask() error {
	CheckMyBlackMarkte()
	return nil
}

func PlayerSave() error {
	Save()
	return nil
}

func RefreshTopList() error {
	RefreshAllTopList()
	RefreshFriendTopList()
	return nil
}

func InitGameTask()  {
	passzerohourTimer := GetGlobalString("C_PassZeroHour")
	passzerohourTask := toolbox.NewTask("Passzerohour",passzerohourTimer,PassZeroHourTask)
	toolbox.AddTask("Passzerohour",passzerohourTask)

	blackMarketTimer := GetGlobalString("C_BlackMarkteRefresh")
	blackMarketTask := toolbox.NewTask("BlackMarkteRefresh",blackMarketTimer,BlackMarkteRefreshTask)
	toolbox.AddTask("BlackMarkteRefresh",blackMarketTask)

	playerSaveTimer := GetGlobalString("C_SaveDataToDB")
	playerSaveTask := toolbox.NewTask("PlayerSaveToDB",playerSaveTimer,PlayerSave)
	toolbox.AddTask("PlayerSaveToDB",playerSaveTask)

	TopListTimer := GetGlobalString("C_TopListRefresh")
	RefreshTopList := toolbox.NewTask("RefreshTopList", TopListTimer, RefreshTopList)
	toolbox.AddTask("RefreshTopList",RefreshTopList)


}