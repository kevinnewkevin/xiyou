package lua

/*
#include "lua.go.h"
*/
import "C"
import (
	"unsafe"
	"fmt"
)

const (
	LUA_VSERSION = C.LUA_VERSION
)

const (
	LUA_REGISTRYINDEX = C.LUA_REGISTRYINDEX
	LUA_ENVIRONINDEX  = C.LUA_ENVIRONINDEX
	LUA_GLOBALSINDEX  = C.LUA_GLOBALSINDEX
)

var _CLua2GLua map[unsafe.Pointer]*LuaState = map[unsafe.Pointer]*LuaState{}

func GetLuaState(clua unsafe.Pointer) *LuaState{
	return _CLua2GLua[clua]
}

type LuaState struct {
	luaState *C.lua_State
}

func (this *LuaState) Open()  {
	this.luaState = C.luaL_newstate()
	_CLua2GLua[unsafe.Pointer(this.luaState)] = this
}
func (this *LuaState) Close() {
	_CLua2GLua[unsafe.Pointer(this.luaState)] = nil
	C.lua_close(this.luaState)
}

func (this *LuaState) Pop(n int)             { this.SetTop(-n-1) }
func (this *LuaState) GetTop() int             { return int(C.lua_gettop(this.luaState)) }
func (this *LuaState) SetTop( idx int)        { C.lua_settop(this.luaState, C.int(idx)) }
func (this *LuaState) PushValue(idx int)     { C.lua_pushvalue(this.luaState, C.int(idx)) }
func (this *LuaState) Remove(idx int)        { C.lua_remove(this.luaState, C.int(idx)) }
func (this *LuaState) Insert( idx int)        { C.lua_insert(this.luaState, C.int(idx)) }
func (this *LuaState) Replace( idx int)       { C.lua_replace(this.luaState, C.int(idx)) }
func (this *LuaState) CheckStack( sz int) int { return int(C.lua_checkstack(this.luaState, C.int(sz))) }

//func (this *LuaState) XMove(L1 *LuaState, L2 *LuaState, n int) { C.lua_xmove(L1.luaState, L2.luaState, C.int(n)) }

func (this *LuaState) IsNumber(idx int) bool    { return C.lua_isnumber(this.luaState, C.int(idx)) != 0 }
func (this *LuaState) IsString(idx int) bool    { return C.lua_isstring(this.luaState, C.int(idx)) != 0 }
func (this *LuaState) IsCFunction( idx int) bool { return C.lua_iscfunction(this.luaState, C.int(idx)) != 0 }
func (this *LuaState) IsUserdata( idx int) bool  { return C.lua_isuserdata(this.luaState, C.int(idx)) != 0 }
func (this *LuaState) Type( idx int) int         { return int(C.lua_type(this.luaState, C.int(idx))) }
func (this *LuaState) TypeName( idx int) string  { return C.GoString(C.lua_typename(this.luaState, C.int(idx))) }

func (this *LuaState) ToNumber( idx int) float64 { return float64(C.lua_tonumber(this.luaState, C.int(idx))) }
func(this *LuaState)  ToInteger( idx int) int    { return int(C.lua_tointeger(this.luaState, C.int(idx))) }
func (this *LuaState) ToBoolean( idx int) bool   { return C.lua_toboolean(this.luaState, C.int(idx)) != 0 }
func (this *LuaState) ToString( idx int) string {
	return C.GoString(C.lua_tolstring(this.luaState, C.int(idx), (*C.size_t)(unsafe.Pointer(nil))))
}

func (this *LuaState) PushNil()               { C.lua_pushnil(this.luaState) }
func(this *LuaState)  PushNumber( n float64) { C.lua_pushnumber(this.luaState, C.lua_Number(n)) }
func (this *LuaState) PushInteger( n int)    { C.lua_pushinteger(this.luaState, C.lua_Integer(n)) }
func (this *LuaState) PushString( s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_pushstring(this.luaState, cstr)
}

func (this *LuaState) PushBoolean( b bool) {
	if b {
		C.lua_pushboolean(this.luaState, 1)
	} else {
		C.lua_pushboolean(this.luaState, 0)
	}
}
func(this *LuaState)  PushLightUserdata( p unsafe.Pointer) { C.lua_pushlightuserdata(this.luaState, p) }

func (this *LuaState) GetTable( idx int) { C.lua_gettable(this.luaState, C.int(idx)) }
func(this *LuaState)  GetField( idx int, k string) {
	cstr := C.CString(k)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_getfield(this.luaState, C.int(idx), cstr)
}
func(this *LuaState)  RawGet( idx int)                 { C.lua_rawget(this.luaState, C.int(idx)) }
func (this *LuaState) RawGetI( idx int, n int)         { C.lua_rawgeti(this.luaState, C.int(idx), C.int(n)) }
func (this *LuaState) CreateTable( nArr int, nRec int) { C.lua_createtable(this.luaState, C.int(nArr), C.int(nRec)) }
func (this *LuaState) NewTable() 					   { C.lua_newtable (this.luaState) }
func (this *LuaState) GetMetaTable( objIndex int) int  { return int(C.lua_getmetatable(this.luaState, C.int(objIndex))) }
func (this *LuaState) GetEnvF( idx int)                { C.lua_getfenv(this.luaState, C.int(idx)) }

func (this *LuaState) SetTable( idx int) { C.lua_settable(this.luaState, C.int(idx)) }
func (this *LuaState) SetField( idx int, k string) {
	cstr := C.CString(k)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_setfield(this.luaState, C.int(idx), cstr)
}
func (this *LuaState) RawSet( idx int)            { C.lua_rawset(this.luaState, C.int(idx)) }
func (this *LuaState) RawSetI( idx int, n int)    { C.lua_rawseti(this.luaState, C.int(idx), C.int(n)) }
func (this *LuaState) SetMetaTable( objIndex int) { C.lua_setmetatable(this.luaState, C.int(objIndex)) }
func (this *LuaState) SetEnvF( idx int)           { C.lua_setfenv(this.luaState, C.int(idx)) }

func (this *LuaState) Call( nArgs int, nResults int) { C.lua_call(this.luaState, C.int(nArgs), C.int(nResults)) }
func (this *LuaState) PCall( nArgs int, nResults int, errFunc int) int {
	return int(C.lua_pcall(this.luaState, C.int(nArgs), C.int(nResults), C.int(errFunc)))
}
func (this *LuaState) CallFunc( funcName string, nArgs int, nResults int) {
	this.PushString( funcName)
	this.GetTable( LUA_GLOBALSINDEX)
	this.Call( nArgs, nResults)
}

func (this *LuaState) CallFuncEx( funcName string, args []interface{}, results *[]interface{}) {
	lastTop := this.GetTop()
	this.PushString( funcName)
	this.GetTable( LUA_GLOBALSINDEX)

	for _, arg := range args {
		switch arg.(type) {
		case int:
			this.PushInteger( arg.(int))
			break
		case int64:
			this.PushInteger( arg.(int))
			break
		case float64:
			this.PushNumber(arg.(float64))
			break
		case string:
			this.PushString( arg.(string))
			break
		default:
			panic("cant not use lua params")
			break
		}
	}

	fmt.Println("CallFuncEx", len(*results))

	this.Call(len(args), len(*results))

	for i := 0; i < len(*results); i++ {
		(*results)[i] = this.ToString( i+1)
	}

	this.Pop(len(*results))

	currTop := this.GetTop()

	if lastTop != currTop {
		panic("Lua stack failed")
	}
}

// lualib.h
func (this *LuaState) OpenBase() {
	C.luaopen_base(this.luaState)
}
func (this *LuaState) OpenTable() {
	C.luaopen_table(this.luaState)
}
func (this *LuaState) OpenIO() {
	C.luaopen_io(this.luaState)
}
func (this *LuaState) OpenOS() {
	C.luaopen_os(this.luaState)
}
func (this *LuaState) OpenString() {
	C.luaopen_string(this.luaState)
}
func (this *LuaState) OpenMath() {
	C.luaopen_math(this.luaState)
}
func (this *LuaState) OpenDebug() {
	C.luaopen_debug(this.luaState)
}
func (this *LuaState) OpenPackage() {
	C.luaopen_package(this.luaState)
}
func (this *LuaState) OpenLibs(){
	C.luaopen_base(this.luaState)
	C.luaopen_table(this.luaState)
	//C.luaopen_io(this.luaState)
	//C.luaopen_os(this.luaState)
	C.luaopen_string(this.luaState)
	C.luaopen_math(this.luaState)
	//C.luaopen_debug(L)
	//C.luaopen_package(L)
}

func (this *LuaState) LoadFile( fileName string) int {
	fileNameC := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	{
		ret := int(C.luaL_loadfile(this.luaState, fileNameC))

		if ret != 0 {
			errorDescC := this.ToString(-1)
			//defer C.free(errorDescC)
			panic(errorDescC)
			return ret
		}
	}
	{
		ret := this.PCall(0,0,0)
		if ret != 0 {
			errorDescC := this.ToString(-1)
			//defer C.free(errorDescC)
			panic(errorDescC)
			return ret
		}
	}
	return 0
}

func (this *LuaState) LoadString( s string) int {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.luaL_loadstring(this.luaState, cstr))
}

func (this *LuaState)LoadApi( api unsafe.Pointer, funcName string, libName string) {
	cstr := C.CString(funcName)
	defer C.free(unsafe.Pointer(cstr))
	C.luaL_loadapi(this.luaState, api, C.CString(funcName), C.CString(libName))
}

func (this *LuaState) OpenPanic( pnc unsafe.Pointer) {
	C.luaL_openpanic(this.luaState, pnc)
}
