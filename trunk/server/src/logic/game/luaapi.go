package game

/*
extern int __loadfile(void*);
*/
import "C"
import (
	"unsafe"
	"logic/lua"

)

var (
	_L *lua.LuaState
	_R string
)

func InitLua(r string){
	_R = r
	_L = &lua.LuaState{}
	_L.Open()

	_L.OpenSys()
	_L.LoadApi(C.__loadfile,"loadfile","sys")
	_L.LoadFile(_R + "main.lua")

}

//export __loadfile
func __loadfile(p unsafe.Pointer) C.int {

	L := lua.GetLuaState(p)
	fileName := L.ToString(-1)
	L.LoadFile(_R + fileName)
	return 0
}
