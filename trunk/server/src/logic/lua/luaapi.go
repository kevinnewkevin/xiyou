package lua


func __init() {
	if adapter == nil {
		adapter = &luaDLL{}
		adapter.init()
	}
}

func __fini() {
	if adapter != nil {
		adapter.fini()
		adapter = nil
	}
}

func luaapi_lasterror() error {
	return adapter.last_error()
}

func lua_close(L uintptr) {
	adapter.lua_close(L)
}

func lua_atpanic(L uintptr, panicf uintptr) uintptr {
	return adapter.lua_atpanic(L, panicf)
}

func lua_gettop(L uintptr) int {
	return adapter.lua_gettop(L)
}
func lua_settop(L uintptr, idx int) {
	adapter.lua_settop(L, idx)
}
func lua_pushvalue(L uintptr, idx int) {
	adapter.lua_pushvalue(L, idx)
}
func lua_remove(L uintptr, idx int) {
	adapter.lua_remove(L, idx)
}
func lua_insert(L uintptr, idx int) {
	adapter.lua_insert(L, idx)
}
func lua_replace(L uintptr, idx int) {
	adapter.lua_replace(L, idx)
}
func lua_checkstack(L uintptr, sz int) int {
	return adapter.lua_checkstack(L, sz)
}

//lua_xmove(from uintptr, to uintptr, n uintptr)
//access functions stack -> C)
func lua_isnumber(L uintptr, idx int) int {
	return adapter.lua_isnumber(L, idx)
}
func lua_isstring(L uintptr, idx int) int {
	return adapter.lua_isstring(L, idx)
}

//func lua_iscfunction(L uintptr, idx int) int
func lua_isuserdata(L uintptr, idx int) int {
	return adapter.lua_isuserdata(L, idx)
}
func lua_type(L uintptr, idx int) int {
	return adapter.lua_type(L, idx)
}
func lua_typename(L uintptr, tp int) string {
	return adapter.lua_typename(L, tp)
}
func lua_equal(L uintptr, idx1 int, idx2 int) int {
	return adapter.lua_equal(L, idx1, idx2)
}
func lua_rawequal(L uintptr, idx1 int, idx2 int) int {
	return adapter.lua_rawequal(L, idx1, idx2)
}
func lua_lessthan(L uintptr, idx1 int, idx2 int) int {
	return adapter.lua_lessthan(L, idx1, idx2)
}
func lua_tonumber(L uintptr, idx int) float64 {
	return adapter.lua_tonumber(L, idx)
}
func lua_tointeger(L uintptr, idx int) int {
	return adapter.lua_tointeger(L, idx)
}
func lua_toboolean(L uintptr, idx int) int {
	return adapter.lua_toboolean(L, idx)
}
func lua_tolstring(L uintptr, idx int, len int) string {
	return adapter.lua_tolstring(L, idx, len)
}
func lua_objlen(L uintptr, idx int) int {
	return adapter.lua_objlen(L, idx)
}

//lua_tocfunction(L uintptr, idx int) uintptr
//lua_touserdata(L uintptr, idx int) uintptr
//lua_tothread(L uintptr, idx int) uintptr
//lua_topointer(L uintptr, idx int) uintptr
//(push functions C -> stack)
func lua_pushnil(L uintptr) {
	adapter.lua_pushnil(L)
}
func lua_pushnumber(L uintptr, n float64) {
	adapter.lua_pushnumber(L, n)
}
func lua_pushinteger(L uintptr, n int) {
	adapter.lua_pushinteger(L, n)
}

//lua_pushlstring(L uintptr, s string, l int)
func lua_pushstring(L uintptr, s string) {
	adapter.lua_pushstring(L, s)
}

//lua_pushvfstring(L uintptr, const char *fmt, va_list argp) string
//lua_pushfstring(L uintptr, const char *fmt, ...) string
//lua_pushcclosure(L uintptr, lua_CFunction fn, int n)
func lua_pushboolean(L uintptr, b int) {
	adapter.lua_pushboolean(L, b)
}

//lua_pushlightuserdata(L uintptr, p uintptr) int
//lua_pushthread(L uintptr)
//(get functions Lua -> stack)
func lua_gettable(L uintptr, idx int) {
	adapter.lua_gettable(L, idx)
}
func lua_getfield(L uintptr, idx int, k string) {
	adapter.lua_getfield(L, idx, k)
}
func lua_rawget(L uintptr, idx int) {
	adapter.lua_rawget(L, idx)
}
func lua_rawgeti(L uintptr, idx int, n int) {
	adapter.lua_rawgeti(L, idx, n)
}
func lua_createtable(L uintptr, narr int, nrec int) {
	adapter.lua_createtable(L, narr, nrec)
}

//lua_newuserdata(L uintptr, sz int) uintptr
func lua_getmetatable(L uintptr, objindex int) int {
	return adapter.lua_getmetatable(L, objindex)
}
func lua_getfenv(L uintptr, idx int) {
	adapter.lua_getfenv(L, idx)
}

//set functions stack -> Lua)
func lua_settable(L uintptr, idx int) {
	adapter.lua_settable(L, idx)
}
func lua_setfield(L uintptr, idx int, k string) {
	adapter.lua_setfield(L, idx, k)
}
func lua_rawset(L uintptr, idx int) {
	adapter.lua_rawset(L, idx)
}
func lua_rawseti(L uintptr, idx int, n int) {
	adapter.lua_rawseti(L, idx, n)
}
func lua_setmetatable(L uintptr, objindex int) int {
	return adapter.lua_setmetatable(L, objindex)
}
func lua_setfenv(L uintptr, idx int) int {
	return adapter.lua_setfenv(L, idx)
}

//`load' and `call' functions load and run Lua code)
func lua_call(L uintptr, nargs int, nresults int) {
	adapter.lua_call(L, nargs, nresults)
}
func lua_pcall(L uintptr, nargs int, nresults int, errfunc int) int {
	return adapter.lua_pcall(L, nargs, nresults, errfunc)
}

//lua_cpcall(L uintptr, lua_CFunction func, void *ud)                        int
//lua_load(L uintptr, lua_Reader reader, void *dt, const char *chunkname)    int
//lua_dump(L uintptr, lua_Writer writer, void *data)                           int
//coroutine functions
//lua_yield(L uintptr, nresults int) int
//lua_resume(L uintptr, narg int) int
//lua_status(L uintptr) int
//garbage-collection function and options
func lua_gc(L uintptr, what int, data int) int {
	return adapter.lua_gc(L, what, data)
}

//miscellaneous functions
func lua_error(L uintptr) int {
	return adapter.lua_error(L)
}
func lua_next(L uintptr, idx int) int {
	return adapter.lua_next(L, idx)
}
func lua_concat(L uintptr, n int) {
	adapter.lua_concat(L, n)
}

//lua_getallocf(L uintptr, ud uintptr)
//lua_setallocf(L uintptr, f uintptr, ud uintptr)
/* Functions to be called by the debuger in specific events
 * typedef void *lua_Hook(L uintptr, lua_Debug *ar)
 */
//lua_getstack (L uintptr,  level int, lua_Debug *ar)			int
//lua_getinfo (L uintptr, const char *what, lua_Debug *ar)      int
//lua_getlocal (L uintptr, const lua_Debug *ar, int n)	string
//lua_setlocal (L uintptr, const lua_Debug *ar, int n)	string
//lua_getupvalue (L uintptr, int funcindex, int n)		string
//lua_setupvalue (L uintptr, int funcindex, int n)		string
//lua_sethook (L uintptr, lua_Hook func, int mask, int count) int
//lua_gethook (L uintptr) uintptr
//lua_gethookmask (L uintptr) int
//lua_gethookcount( L uintptr) int
/* From Lua 5.2. */
func lua_upvalueid(L uintptr, idx int, n int) uintptr {
	return adapter.lua_upvalueid(L, idx, n)
}

//lua_upvaluejoin (L uintptr, idx int1,  n1 int, idx int2, int n2)
//lua_loadx (L uintptr, lua_Reader reader, void *dt, const char *chunkname, const char *mode) int
//lauxlib.h
func luaL_openlib(L uintptr, libname string, l uintptr, nup int) {
	adapter.luaL_openlib(L, libname, l, nup)
}
func luaL_register(L uintptr, libname string, l uintptr) {
	adapter.luaL_register(L, libname, l)
}
func luaL_register2(L uintptr, libname string, funcname string, l uintptr) {
	adapter.luaL_register2(L, libname, funcname, l)
}
func luaL_getmetafield(L uintptr, obj int, e string) int {
	return adapter.luaL_getmetafield(L, obj, e)
}
func luaL_callmeta(L uintptr, obj int, e string) int {
	return adapter.luaL_callmeta(L, obj, e)
}
func luaL_typerror(L uintptr, narg int, tname string) int {
	return adapter.luaL_typerror(L, narg, tname)
}
func luaL_argerror(L uintptr, numarg int, extramsg string) int {
	return adapter.luaL_argerror(L, numarg, extramsg)
}

//luaL_checklstring(L uintptr, numArg int, l int) string
//luaL_optlstring(L uintptr, numArg int, def int, l uintptr) string
//luaL_checknumber(L uintptr, numArg int) float64
//luaL_optnumber(L uintptr, nArg int, def int) float64
//luaL_checkinteger(L uintptr, numArg int) int
//luaL_optinteger(L uintptr, nArg int, def int) int
//luaL_checkstack(L uintptr, sz int, msg string)
//luaL_checktype(L uintptr, narg int, t int)
//luaL_checkany(L uintptr, narg int)
//luaL_newmetatable(L uintptr, tname string) int
//luaL_checkudata(L uintptr, ud int, tname string) uintptr
//luaL_where(L uintptr, lvl int)
//luaL_error(L uintptr, const char *fmt, ...)int
//luaL_checkoption(L uintptr,  narg int,def string, const char *const lst[])          int
//luaL_ref(L uintptr, t int) int
//luaL_unref(L uintptr, t int, ref int)
func luaL_loadfile(L uintptr, filename string) int {
	return adapter.luaL_loadfile(L, filename)
}
func luaL_loadbuffer(L uintptr, buff string, name string) int {
	return adapter.luaL_loadbuffer(L, buff, name)
}
func luaL_loadstring(L uintptr, s string) int {
	return adapter.luaL_loadstring(L, s)
}
func luaL_newstate() uintptr {
	return adapter.luaL_newstate()
}

//luaL_gsub(L uintptr, s string, p string, r string) string
//luaL_findtable(L uintptr, idx int, fname string, szhint int) string
/* From Lua 5.2. */
//luaL_fileresult(L uintptr, stat int, fname string) int
//luaL_execresult(L uintptr, stat int) int
//luaL_loadfilex(L uintptr, filename string, mode string) int
//luaL_loadbufferx(L uintptr, buff string, name string, mode string) int
//luaL_traceback(L uintptr, L1 uintptr, msg string, level int)

func luaopen_base(L uintptr) int {
	return adapter.luaopen_base(L)
}
func luaopen_math(L uintptr) int {
	return adapter.luaopen_math(L)
}
func luaopen_string(L uintptr) int {
	return adapter.luaopen_string(L)
}
func luaopen_table(L uintptr) int {
	return adapter.luaopen_table(L)
}
func luaopen_io(L uintptr) int {
	return adapter.luaopen_io(L)
}
func luaopen_os(L uintptr) int {
	return adapter.luaopen_os(L)
}
func luaopen_package(L uintptr) int {
	return adapter.luaopen_package(L)
}
func luaopen_debug(L uintptr) int {
	return adapter.luaopen_debug(L)
}
func luaopen_bit(L uintptr) int {
	return adapter.luaopen_bit(L)
}
func luaopen_jit(L uintptr) int {
	return adapter.luaopen_jit(L)
}
func luaopen_ffi(L uintptr) int {
	return adapter.luaopen_ffi(L)
}
func luaL_openlibs(L uintptr) {
	adapter.luaL_openlibs(L)
}
