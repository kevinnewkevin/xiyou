package lua

/*
#include "lua.go.h"
extern int __panic(lua_State*);
extern int __err(lua_State*);
extern int __test(lua_State*);
*/
import "C"

//export __panic
func __panic(L LuaState) C.int {
	return 0
}

//export __err
func __err(L LuaState) C.int {
	idx := 1
	s := ToString(L, idx)
	panic(s)
	//	return 0
}

//export __test
func __test(L LuaState) C.int {
	idx := 1
	v1 := ToInteger(L, idx)
	idx++
	v2 := ToInteger(L, idx)

	idx = 0
	PushInteger(L, v1+v2)
	idx++
	return C.int(idx)
}

func RegistSystemAPI(L LuaState) {
	LoadApi(L, C.__err, "err", "sys")
	LoadApi(L, C.__test, "test", "sys")
}
