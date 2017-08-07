package lua

/*
#include "lua.go.h"
*/
import "C"
import "unsafe"

const (
	LUA_VSERSION = C.LUA_VERSION
)

const (
	LUA_REGISTRYINDEX = C.LUA_REGISTRYINDEX
	LUA_ENVIRONINDEX  = C.LUA_ENVIRONINDEX
	LUA_GLOBALSINDEX  = C.LUA_GLOBALSINDEX
)

type LuaState *C.lua_State

func Open() LuaState {
	return C.luaL_newstate()
}
func Close(L LuaState) { C.lua_close(L) }

func GetTop(L LuaState) int             { return int(C.lua_gettop(L)) }
func SetTop(L LuaState, idx int)        { C.lua_settop(L, C.int(idx)) }
func PushValue(L LuaState, idx int)     { C.lua_pushvalue(L, C.int(idx)) }
func Remove(L LuaState, idx int)        { C.lua_remove(L, C.int(idx)) }
func Insert(L LuaState, idx int)        { C.lua_insert(L, C.int(idx)) }
func Replace(L LuaState, idx int)       { C.lua_replace(L, C.int(idx)) }
func CheckStack(L LuaState, sz int) int { return int(C.lua_checkstack(L, C.int(sz))) }

func XMove(L1 LuaState, L2 LuaState, n int) { C.lua_xmove(L1, L2, C.int(n)) }

func IsNumber(L LuaState, idx int) bool    { return C.lua_isnumber(L, C.int(idx)) != 0 }
func IsString(L LuaState, idx int) bool    { return C.lua_isstring(L, C.int(idx)) != 0 }
func IsCFunction(L LuaState, idx int) bool { return C.lua_iscfunction(L, C.int(idx)) != 0 }
func IsUserdata(L LuaState, idx int) bool  { return C.lua_isuserdata(L, C.int(idx)) != 0 }
func Type(L LuaState, idx int) int         { return int(C.lua_type(L, C.int(idx))) }
func TypeName(L LuaState, idx int) string  { return C.GoString(C.lua_typename(L, C.int(idx))) }

func ToNumber(L LuaState, idx int) float64 { return float64(C.lua_tonumber(L, C.int(idx))) }
func ToInteger(L LuaState, idx int) int    { return int(C.lua_tointeger(L, C.int(idx))) }
func ToBoolean(L LuaState, idx int) bool   { return C.lua_toboolean(L, C.int(idx)) != 0 }
func ToString(L LuaState, idx int) string {
	return C.GoString(C.lua_tolstring(L, C.int(idx), (*C.size_t)(unsafe.Pointer(nil))))
}

func PushNil(L LuaState)               { C.lua_pushnil(L) }
func PushNumber(L LuaState, n float64) { C.lua_pushnumber(L, C.lua_Number(n)) }
func PushInteger(L LuaState, n int)    { C.lua_pushinteger(L, C.lua_Integer(n)) }
func PushString(L LuaState, s string)  { 
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_pushstring(L,cstr)
}

func PushBoolean(L LuaState, b bool) {
	if b {
		C.lua_pushboolean(L, 1)
	} else {
		C.lua_pushboolean(L, 0)
	}
}
func PushLightUserdata(L LuaState, p unsafe.Pointer) { C.lua_pushlightuserdata(L, p) }

func GetTable(L LuaState, idx int)               { C.lua_gettable(L, C.int(idx)) }
func GetField(L LuaState, idx int, k string)     { 
	cstr := C.CString(k)	
	defer C.free(unsafe.Pointer(cstr))
	C.lua_getfield(L, C.int(idx), cstr)
}
func RawGet(L LuaState, idx int)                 { C.lua_rawget(L, C.int(idx)) }
func RawGetI(L LuaState, idx int, n int)         { C.lua_rawgeti(L, C.int(idx), C.int(n)) }
func CreateTable(L LuaState, nArr int, nRec int) { C.lua_createtable(L, C.int(nArr), C.int(nRec)) }
func GetMetaTable(L LuaState, objIndex int) int  { return int(C.lua_getmetatable(L, C.int(objIndex))) }
func GetEnvF(L LuaState, idx int)                { C.lua_getfenv(L, C.int(idx)) }

func SetTable(L LuaState, idx int)           { C.lua_settable(L, C.int(idx)) }
func SetField(L LuaState, idx int, k string) { 
	cstr := C.CString(k)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_setfield(L, C.int(idx), cstr)
}
func RawSet(L LuaState, idx int)             { C.lua_rawset(L, C.int(idx)) }
func RawSetI(L LuaState, idx int, n int)     { C.lua_rawseti(L, C.int(idx), C.int(n)) }
func SetMetaTable(L LuaState, objIndex int)  { C.lua_setmetatable(L, C.int(objIndex)) }
func SetEnvF(L LuaState, idx int)            { C.lua_setfenv(L, C.int(idx)) }

func Call(L LuaState, nArgs int, nResults int) { C.lua_call(L, C.int(nArgs), C.int(nResults)) }
func CallP(L LuaState, nArgs int, nResults int, errFunc int) int {
	return int(C.lua_pcall(L, C.int(nArgs), C.int(nResults), C.int(errFunc)))
}

// lualib.h
func OpenBase(L LuaState) {
	C.luaopen_base(L)
}
func OpenTable(L LuaState) {
	C.luaopen_table(L)
}
func OpenIO(L LuaState) {
	C.luaopen_io(L)
}
func OpenOS(L LuaState) {
	C.luaopen_os(L)
}
func OpenString(L LuaState) {
	C.luaopen_string(L)
}
func OpenMath(L LuaState) {
	C.luaopen_math(L)
}
func OpenDebug(L LuaState) {
	C.luaopen_debug(L)
}
func OpenPackage(L LuaState) {
	C.luaopen_package(L)
}
func OpenLibs(L LuaState) {
	C.luaopen_base(L)
	C.luaopen_table(L)
	C.luaopen_io(L)
	C.luaopen_os(L)
	C.luaopen_string(L)
	C.luaopen_math(L)
	C.luaopen_debug(L)
	C.luaopen_package(L)
}

func LoadFile(L LuaState, fileName string) int {
	fileNameC := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	{
		ret := int(C.luaL_loadfile(L, fileNameC))

		if ret != 0 {
			errorDescC := C.luaL_tostring(L, C.int(-1));
			//defer C.free(errorDescC)
			panic(errorDescC)
			return ret
		}
	}
	{
		ret := int(C.lua_pcall(L, 0, 0, 0 ))
		if ret != 0{
			errorDescC := C.luaL_tostring(L,-1);
			//defer C.free(errorDescC)
			panic(errorDescC)
			return ret
		}
	}
	return 0
}

func LoadString(L LuaState, s string) int {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.luaL_loadstring(L, cstr))
}

func LoadApi(L LuaState, api unsafe.Pointer, funcName string, libName string) {
	cstr := C.CString(funcName)
	defer C.free(unsafe.Pointer(cstr))
	C.luaL_loadapi(L, api, C.CString(funcName), C.CString(libName))
}

func OpenPanic(L LuaState, pnc unsafe.Pointer) {
	C.luaL_openpanic(L, pnc)
}
