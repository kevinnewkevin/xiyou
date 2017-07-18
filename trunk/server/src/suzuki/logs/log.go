package logs

import (
	"fmt"
	"time"
)

const (
	ERROR   = 1 << 0
	WARNING = 1 << 1
	INFO    = 1 << 2
	DEBUG   = 1 << 3
	ALL     = ERROR | WARNING | INFO | DEBUG
)

type (
	core struct {
		bTimestamp bool
		priority   int
		ostream    chan (string)
	}
)

func (this *core) setPriority(mask int) {
	this.priority = mask
}

func (this *core) setCallBack(f func(ch chan (string))) {
	this.ostream <- "" //关闭原来的
	go f(this.ostream)
}

func (this *core) getTimestamp() string {
	if this.bTimestamp {
		return time.Now().Format("[2006-01-02 15:04:05.000000]")
	}
	return ""
}

func (this *core) logError(s string) {
	if this.priority&ERROR == 0 {
		return
	}
	this.ostream <- fmt.Sprintf("[ERRO]%s%s", this.getTimestamp(), s)
}

func (this *core) logWarning(s string) {
	if this.priority&WARNING == 0 {
		return
	}
	this.ostream <- fmt.Sprintf("[WARN]%s%s", this.getTimestamp(), s)
}

func (this *core) logInfo(s string) {
	if this.priority&INFO == 0 {
		return
	}
	this.ostream <- fmt.Sprintf("[INFO]%s%s", this.getTimestamp(), s)
}

func (this *core) logDebug(s string) {
	if this.priority&DEBUG == 0 {
		return
	}
	this.ostream <- fmt.Sprintf("[DEBG]%s%s", this.getTimestamp(), s)

}

var _core *core

func Fini() {
	if _core != nil {
		_core.ostream <- "" //关闭
	}
	_core = nil
}

func Init() {
	Fini()
	_core = &core{
		true,
		ALL,
		make(chan string, 2048),
	}
	go func(ch chan (string)) {
		for {
			select {
			case s := <-ch:
				if s == "" {
					return
				}
				fmt.Println(s)
				break
			}
		}
	}(_core.ostream)
}

func Error(f string, a ...interface{}) {
	if _core == nil {
		return
	}
	_core.logError(fmt.Sprintf(f, a...))
}

func Warning(f string, a ...interface{}) {
	if _core == nil {
		return
	}
	_core.logWarning(fmt.Sprintf(f, a...))
}

func Info(f string, a ...interface{}) {
	if _core == nil {
		return
	}
	_core.logInfo(fmt.Sprintf(f, a...))
}

func Debug(f string, a ...interface{}) {
	if _core == nil {
		return
	}
	_core.logDebug(fmt.Sprintf(f, a...))
}

func SetCallBack(f func(ch chan (string))) {
	if _core == nil {
		return
	}
	_core.setCallBack(f)
}
