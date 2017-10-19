package game

import (
	"github.com/astaxie/beego/toolbox"
)

func PassZeroHourTask() error {
	OnlinePlayerPassZeroHour()
	return nil
}

func InitGameTask()  {
	passzerohourTimer := GetGlobalString("C_PassZeroHour")
	passzerohourTask := toolbox.NewTask("huodong",passzerohourTimer,PassZeroHourTask)
	toolbox.AddTask("huodong",passzerohourTask)
}