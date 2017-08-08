package game

/*
#include "lua.go.h"
extern int __loadfile(lua_State*);
extern int __GetPlayerUnit(lua_State*);
extern int __GetUnitProperty(lua_State*);
extern int __GetEnemy(lua_State*);
*/
import "C"
import (
	"fmt"
	"logic/game"
)

var _GL LuaState
var _GLFileRoot string

func InitGlobalLuaState(r string) {
	_GLFileRoot = r
	_GL = Open()
	RegistSystemAPI(_GL)
	RegistGameAPI(_GL)

	LoadFile(_GL, _GLFileRoot+"main.lua")
}

func GetGlobalLuaState() LuaState {
	return _GL
}

//export __loadfile
func __loadfile(L LuaState) C.int {
	fileName := ToString(L, 1)

	LoadFile(L, _GLFileRoot+fileName)

	return C.int(0)
}

//export __GetPlayerUnit
func __GetPlayerUnit(L LuaState) C.int {
	idx := 1
	v1 := ToInteger(L, idx)
	idx++
	v2 := ToInteger(L, idx)

	idx = 0
	PushInteger(L, v1+v2)
	idx++
	fmt.Println("GetPlayerUnit", v1, v2)
	return C.int(idx)
}

//export __GetUnitProperty
func __GetUnitProperty(L LuaState) C.int {
	idx := 1
	v1 := ToInteger(L, idx)
	var end int
	if v1 == 10 {
		end = 1000000
	} else {
		end = 999999
	}

	idx = 0
	PushInteger(L, end)
	idx++
	fmt.Println("GetUnitProperty", v1)
	return C.int(idx)
}

//export __GetEnemy
func __GetEnemy(L LuaState) C.int {
	idx := 1
	battleId := ToInteger(L, idx)
	idx++
	camp := ToInteger(L, idx)
	room := game.FindBattle(int64(battleId))
	if room == nil {
		return idx
	}

	targets := []int{}

	for _, u := range room.Units {
		if u.Owner.camp == camp {
			continue
		}
		targets = append(targets, u.InstId)
	}

	idx = 0
	PushInteger(L, targets)
	idx++
	fmt.Println("GetEnemy", v1)
	return C.int(idx)
}

func RegistGameAPI(L LuaState) {
	LoadApi(L, C.__GetPlayerUnit, "GetPlayerUnit", "Player")
	LoadApi(L, C.__GetUnitProperty, "GetUnitProperty", "Player")
	LoadApi(L, C.__GetEnemy, "GetEnemy", "Player")
}
