package lua

import (
	"fmt"
	"unsafe"
	"github.com/yuin/gopher-lua"
)



type LuaState struct {
	//luaState uintptr
	L *lua.LState
}

func (this *LuaState) Open() {
	this.L = lua.NewState()
}
func (this *LuaState) Close() {
	this.L.Close()
}

func (this *LuaState) Pop(n int)             { this.SetTop(-n - 1) }
func (this *LuaState) GetTop() int           { return lua_gettop(this.luaState) }
func (this *LuaState) SetTop(idx int)        { lua_settop(this.luaState, idx) }
func (this *LuaState) PushValue(idx int)     { lua_pushvalue(this.luaState, idx) }
func (this *LuaState) Remove(idx int)        { lua_remove(this.luaState, idx) }
func (this *LuaState) Insert(idx int)        { lua_insert(this.luaState, idx) }
func (this *LuaState) Replace(idx int)       { lua_replace(this.luaState, idx) }
func (this *LuaState) CheckStack(sz int) int { return lua_checkstack(this.luaState, sz) }

//func (this *LuaState) XMove(L1 *LuaState, L2 *LuaState, n int) { C.lua_xmove(L1.luaState, L2.luaState, C.int(n)) }

func (this *LuaState) IsNumber(idx int) bool { return lua_isnumber(this.luaState, idx) != 0 }
func (this *LuaState) IsString(idx int) bool { return lua_isstring(this.luaState, idx) != 0 }

//func (this *LuaState) IsCFunction( idx int) bool { return lua_iscfunction(this.luaState, C.int(idx)) != 0 }
func (this *LuaState) IsUserdata(idx int) bool { return lua_isuserdata(this.luaState, idx) != 0 }
func (this *LuaState) Type(idx int) int        { return lua_type(this.luaState, idx) }
func (this *LuaState) TypeName(idx int) string { return lua_typename(this.luaState, idx) }

func (this *LuaState) ToNumber(idx int) float64 { return lua_tonumber(this.luaState, idx) }
func (this *LuaState) ToInteger(idx int) int    { return int(lua_tointeger(this.luaState, idx)) }
func (this *LuaState) ToLong(idx int) int64    { return int64(lua_tointeger(this.luaState, idx)) }
func (this *LuaState) ToBoolean(idx int) bool   { return lua_toboolean(this.luaState, idx) != 0 }
func (this *LuaState) ToString(idx int) string {
	return lua_tolstring(this.luaState, idx, nil)
}

func (this *LuaState) PushNil()             { lua_pushnil(this.luaState) }
func (this *LuaState) PushNumber(n float64) { lua_pushnumber(this.luaState, n) }
func (this *LuaState) PushInteger(n int)    { lua_pushinteger(this.luaState, uintptr(n)) }
func (this *LuaState) PushLong(n int64)    { lua_pushLong(this.luaState, n) }
func (this *LuaState) PushString(s string) {
	lua_pushstring(this.luaState, s)
}
func (this *LuaState) PushBoolean(b bool) {
	if b {
		lua_pushboolean(this.luaState, 1)
	} else {
		lua_pushboolean(this.luaState, 0)
	}
}
func (this *LuaState) Push(v interface{}){
	switch v.(type) {
	case int:
		this.PushInteger(v.(int))
		break
	case int64:
		this.PushLong(v.(int64))
		break
	case float64:
		this.PushNumber(v.(float64))
		break
	case string:
		this.PushString(v.(string))
		break
	case bool:
		this.PushBoolean(v.(bool))
	default:
		panic("cant not use lua params")
		break
	}
}
func (this* LuaState) PushArray(arr []interface{}){
	this.NewTable()
	for k, v :=range arr{
		this.PushInteger(k)
		this.Push(v)
		this.SetTable(-3)
	}
}

func (this *LuaState) To(v *interface{}, idx int){
	switch (*v).(type) {
	case int:
		(*v) = this.ToInteger(idx)
		break
	case int64:
		(*v)= this.ToLong(idx)
		break
	case float64:
		(*v) = this.ToNumber(idx)
		break
	case string:
		(*v) = this.ToString(idx)
		break
	case bool:
		(*v) = this.ToBoolean(idx)
	default:
		panic("cant not use lua params rsuly")
		break
	}
}

func (this* LuaState) ToArray(arr *[]interface{}){

}

//func(this *LuaState)  PushLightUserdata( p unsafe.Pointer) { lua_pushlightuserdata(this.luaState, p) }

func (this *LuaState) GetTable(idx int) { lua_gettable(this.luaState, idx) }
func (this *LuaState) GetField(idx int, k string) {
	lua_getfield(this.luaState, idx, k)
}
func (this *LuaState) RawGet(idx int)                 { lua_rawget(this.luaState, idx) }
func (this *LuaState) RawGetI(idx int, n int)         { lua_rawgeti(this.luaState, idx, n) }
func (this *LuaState) CreateTable(nArr int, nRec int) { lua_createtable(this.luaState, nArr, nRec) }
func (this *LuaState) NewTable()                      { lua_createtable(this.luaState, 0, 0) }
func (this *LuaState) GetMetaTable(objIndex int) int {
	return lua_getmetatable(this.luaState, objIndex)
}
func (this *LuaState) GetEnvF(idx int) { lua_getfenv(this.luaState, idx) }

func (this *LuaState) SetTable(idx int) { lua_settable(this.luaState, idx) }
func (this *LuaState) SetField(idx int, k string) {

	lua_setfield(this.luaState, idx, k)
}
func (this *LuaState) RawSet(idx int)            { lua_rawset(this.luaState, idx) }
func (this *LuaState) RawSetI(idx int, n int)    { lua_rawseti(this.luaState, idx, n) }
func (this *LuaState) SetMetaTable(objIndex int) { lua_setmetatable(this.luaState, objIndex) }
func (this *LuaState) SetEnvF(idx int)           { lua_setfenv(this.luaState, idx) }

func (this *LuaState) Call(nArgs int, nResults int) { lua_call(this.luaState, nArgs, nResults) }
func (this *LuaState) PCall(nArgs int, nResults int, errFunc int) int {
	return lua_pcall(this.luaState, nArgs, nResults, errFunc)
}
func (this *LuaState) CallFunc(funcName string, nArgs int, nResults int) {
	this.PushString(funcName)
	this.GetTable(LUA_GLOBALSINDEX)
	this.Call(nArgs, nResults)
}

func (this *LuaState) CallFuncEx(funcName string, args []interface{}, results *[]interface{}) {
	lastTop := this.GetTop()
	this.PushString(funcName)
	this.GetTable(LUA_GLOBALSINDEX)

	for _, arg := range args {
		switch arg.(type) {
		case int:
			this.PushInteger(arg.(int))
			break
		case int64:
			this.PushLong(arg.(int64))
			break
		case float64:
			this.PushNumber(arg.(float64))
			break
		case string:
			this.PushString(arg.(string))
			break
		case bool:
			this.PushBoolean(arg.(bool))
		default:
			panic("cant not use lua params")
			break
		}
	}

	fmt.Println("CallFuncEx", len(*results), funcName)

	this.Call(len(args), len(*results))

	resLen := len(*results)
	for i := 0; i < resLen; i++ {
		idx := i-resLen
		switch (*results)[i].(type) {
		case int:
			(*results)[i] = this.ToInteger(idx)
			break
		case int64:
			(*results)[i] = this.ToLong(idx)
			break
		case float64:
			(*results)[i] = this.ToNumber(idx)
			break
		case string:
			(*results)[i] = this.ToString(idx)
			break
		case bool:
			(*results)[i] = this.ToBoolean(idx)
		default:
			panic("cant not use lua params rsuly")
			break
		}
	}

	this.Pop(len(*results))

	currTop := this.GetTop()

	if lastTop != currTop {
		panic("Lua stack failed")
	}
}

// lualib.h
func (this *LuaState) OpenBase() {
	luaopen_base(this.luaState)
}
func (this *LuaState) OpenTable() {
	luaopen_table(this.luaState)
}
func (this *LuaState) OpenIO() {
	luaopen_io(this.luaState)
}
func (this *LuaState) OpenOS() {
	luaopen_os(this.luaState)
}
func (this *LuaState) OpenString() {
	luaopen_string(this.luaState)
}
func (this *LuaState) OpenMath() {
	luaopen_math(this.luaState)
}
func (this *LuaState) OpenDebug() {
	luaopen_debug(this.luaState)
}
func (this *LuaState) OpenPackage() {
	luaopen_package(this.luaState)
}
func (this *LuaState) OpenLibs() {
	luaopen_base(this.luaState)
	luaopen_table(this.luaState)
	luaopen_string(this.luaState)
	luaopen_math(this.luaState)
}

func (this *LuaState) LoadFile(fileName string) int {

	r := luaL_loadfile(this.luaState, fileName)

	if r != 0 {
		fmt.Println(lua_tolstring(this.luaState, -1, nil))
	}

	this.Call(0,0)

	return 0
}

func (this *LuaState) LoadString(s string) int {
	return luaL_loadstring(this.luaState, s)
}

func (this *LuaState) LoadApi(api unsafe.Pointer, funcName string, libName string) {
	luaL_register2(this.luaState, libName, funcName, uintptr(api))
}

func (this *LuaState) OpenPanic(pnc unsafe.Pointer) {
	lua_atpanic(this.luaState, uintptr(pnc))
}
