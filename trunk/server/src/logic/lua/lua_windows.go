package lua

//#include <stdlib.h>
import "C"
import (
	"syscall"
	"unsafe"
)

type luaDLL struct {
	object    *syscall.DLL
	lasterror error
	apimap    map[string]*syscall.Proc
}

func (this *luaDLL) init() {
	this.object, this.lasterror = syscall.LoadDLL("lua51.dll")
	if this.lasterror != nil {
		panic(this.lasterror)
	}
	this.apimap = map[string]*syscall.Proc{}
	for _, name := range luaApiNameList {
		this.apimap[name], this.lasterror = this.object.FindProc(name)
		if this.lasterror != nil {
			panic(this.lasterror)
		}
	}

}

func (this *luaDLL) fini() {
	this.object.Release()
}

func (this *luaDLL) last_error() error {
	return this.lasterror
}

// state manipulation
//lua_newstate (f uintptr, ud uintptr) uintptr
func (this *luaDLL) lua_close(L uintptr) {
	_, _, err := this.apimap["lua_close"].Call(L)
	this.lasterror = err
}

//lua_newthread(L uintptr) uintptr
func (this *luaDLL) lua_atpanic(L uintptr, panicf uintptr) uintptr {
	r1, _, err := this.apimap["lua_atpanic"].Call(L, panicf)
	this.lasterror = err
	return r1
}

// basic stack manipulation
func (this *luaDLL) lua_gettop(L uintptr) int {
	r1, _, err := this.apimap["lua_gettop"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_settop(L uintptr, idx int) {
	_, _, err := this.apimap["lua_settop"].Call(L, uintptr(idx))
	this.lasterror = err

}
func (this *luaDLL) lua_pushvalue(L uintptr, idx int) {
	_, _, err := this.apimap["lua_pushvalue"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_remove(L uintptr, idx int) {
	_, _, err := this.apimap["lua_remove"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_insert(L uintptr, idx int) {
	_, _, err := this.apimap["lua_insert"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_replace(L uintptr, idx int) {
	_, _, err := this.apimap["lua_replace"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_checkstack(L uintptr, sz int) int {
	r1, _, err := this.apimap["lua_checkstack"].Call(L, uintptr(sz))
	this.lasterror = err
	return int(r1)
}

//access functions (stack -> C)
func (this *luaDLL) lua_isnumber(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_isnumber"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_isstring(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_isstring"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_iscfunction(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_iscfunction"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}

func (this *luaDLL) lua_isuserdata(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_isuserdata"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_type(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_type"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_typename(L uintptr, tp int) string {
	r1, _, err := this.apimap["lua_typename"].Call(L, uintptr(tp))
	this.lasterror = err
	return C.GoString((*C.char)(unsafe.Pointer(r1)))
}
func (this *luaDLL) lua_equal(L uintptr, idx1 int, idx2 int) int {
	r1, _, err := this.apimap["lua_equal"].Call(L, uintptr(idx1), uintptr(idx2))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_rawequal(L uintptr, idx1 int, idx2 int) int {
	r1, _, err := this.apimap["lua_rawequal"].Call(L, uintptr(idx1), uintptr(idx2))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_lessthan(L uintptr, idx1 int, idx2 int) int {
	r1, _, err := this.apimap["lua_lessthan"].Call(L, uintptr(idx1), uintptr(idx2))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_tonumber(L uintptr, idx int) float64 {
	r1, _, err := this.apimap["lua_tonumber"].Call(L, uintptr(idx))
	this.lasterror = err
	return float64(r1)
}
func (this *luaDLL) lua_tointeger(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_tointeger"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_toboolean(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_toboolean"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}

func (this *luaDLL) lua_tolstring(L uintptr, idx int, len int) string {
	r1, _, err := this.apimap["lua_tolstring"].Call(L, uintptr(idx), uintptr(len))
	this.lasterror = err
	return C.GoString((*C.char)(unsafe.Pointer(r1)))
}
func (this *luaDLL) lua_objlen(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_objlen"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}

//func (this *luaDLL)  lua_tocfunction (L uintptr, idx int)                 lua_CFunction
//func (this *luaDLL)  lua_touserdata (L uintptr, idx int)                  void*
//func (this *luaDLL)  lua_tothread (L uintptr, idx int)                    lua_State*
//func (this *luaDLL)  lua_topointer(L uintptr, idx int)                   const void*

//push functions (C -> stack)
func (this *luaDLL) lua_pushnil(L uintptr) {
	_, _, err := this.apimap["lua_pushnil"].Call(L)
	this.lasterror = err

}
func (this *luaDLL) lua_pushnumber(L uintptr, n float64) {
	_, _, err := this.apimap["lua_pushnumber"].Call(L, *(*uintptr)(unsafe.Pointer(&n)))
	this.lasterror = err

}
func (this *luaDLL) lua_pushinteger(L uintptr, n int) {
	_, _, err := this.apimap["lua_pushinteger"].Call(L, uintptr(n))
	this.lasterror = err
}

//func(this*luaDLL) (lua_pushlstring) (L uintptr, const char *s, size_t l);
func (this *luaDLL) lua_pushstring(L uintptr, s string) {
	ps := unsafe.Pointer(C.CString(s))
	defer C.free(ps)
	_, _, err := this.apimap["lua_pushstring"].Call(L, uintptr(ps))
	this.lasterror = err
}

//func(this*luaDLL) (lua_pushvfstring) (L uintptr, const char *fmt, va_list argp) string
//func(this*luaDLL) (lua_pushfstring) (L uintptr, const char *fmt, ...) string
//func(this*luaDLL) (lua_pushcclosure) (L uintptr, lua_CFunction fn, int n)
func (this *luaDLL) lua_pushboolean(L uintptr, b int) {
	_, _, err := this.apimap["lua_pushboolean"].Call(L, uintptr(b))
	this.lasterror = err
}

//func(this*luaDLL) (lua_pushlightuserdata) (L uintptr, void *p);
//func(this*luaDLL) (lua_pushthread) (L uintptr);
//get functions (Lua -> stack)
func (this *luaDLL) lua_gettable(L uintptr, idx int) {
	_, _, err := this.apimap["lua_gettable"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_getfield(L uintptr, idx int, k string) {
	pk := unsafe.Pointer(C.CString(k))
	//defer C.free(unsafe.Pointer(pk))
	_, _, err := this.apimap["lua_getfield"].Call(L, uintptr(idx), uintptr(pk))
	this.lasterror = err

}
func (this *luaDLL) lua_rawget(L uintptr, idx int) {
	_, _, err := this.apimap["lua_rawget"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_rawgeti(L uintptr, idx int, n int) {
	_, _, err := this.apimap["lua_rawgeti"].Call(L, uintptr(idx), uintptr(n))
	this.lasterror = err
}
func (this *luaDLL) lua_createtable(L uintptr, narr int, nrec int) {
	_, _, err := this.apimap["lua_createtable"].Call(L, uintptr(narr), uintptr(nrec))
	this.lasterror = err
}

//func(this*luaDLL)lua_newuserdata (L,  sz int) uintptr;
func (this *luaDLL) lua_getmetatable(L uintptr, objindex int) int {
	r1, _, err := this.apimap["lua_getmetatable"].Call(L, uintptr(objindex))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_getfenv(L uintptr, idx int) {
	_, _, err := this.apimap["lua_getfenv"].Call(L, uintptr(idx))
	this.lasterror = err
}

//set functions (stack -> Lua)
func (this *luaDLL) lua_settable(L uintptr, idx int) {
	_, _, err := this.apimap["lua_settable"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_setfield(L uintptr, idx int, k string) {
	pk := unsafe.Pointer(C.CString(k))
	_, _, err := this.apimap["lua_setfield"].Call(L, uintptr(idx), uintptr(pk))
	this.lasterror = err
}
func (this *luaDLL) lua_rawset(L uintptr, idx int) {
	_, _, err := this.apimap["lua_rawset"].Call(L, uintptr(idx))
	this.lasterror = err
}
func (this *luaDLL) lua_rawseti(L uintptr, idx int, n int) {
	_, _, err := this.apimap["lua_rawseti"].Call(L, uintptr(idx), uintptr(n))
	this.lasterror = err
}
func (this *luaDLL) lua_setmetatable(L uintptr, objindex int) int {
	r1, _, err := this.apimap["lua_setmetatable"].Call(L, uintptr(objindex))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_setfenv(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_setfenv"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}

//`load' and `call' functions (load and run Lua code)
func (this *luaDLL) lua_call(L uintptr, nargs int, nresults int) {
	_, _, err := this.apimap["lua_call"].Call(L, uintptr(nargs), uintptr(nresults))
	this.lasterror = err
}
func (this *luaDLL) lua_pcall(L uintptr, nargs int, nresults int, errfunc int) int {
	r1, _, err := this.apimap["lua_pcall"].Call(L, uintptr(nargs), uintptr(nresults), uintptr(errfunc))
	this.lasterror = err
	return int(r1)
}

//func(this*luaDLL)lua_cpcall (L uintptr, lua_CFunction func, void *ud);
//func(this*luaDLL)lua_load (L uintptr, lua_Reader reader, void *dt, const char *chunkname);
//func(this*luaDLL)lua_dump (L uintptr, lua_Writer writer, void *data);
//coroutine functions
//LUA_API int  (lua_yield) (lua_State *L, int nresults);
//LUA_API int  (lua_resume) (lua_State *L, int narg);
//LUA_API int  (lua_status) (lua_State *L);
//garbage-collection function and options
func (this *luaDLL) lua_gc(L uintptr, what int, data int) int {
	r1, _, err := this.apimap["lua_gc"].Call(L, uintptr(what), uintptr(data))
	this.lasterror = err
	return int(r1)
}

//miscellaneous functions
func (this *luaDLL) lua_error(L uintptr) int {
	r1, _, err := this.apimap["lua_error"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_next(L uintptr, idx int) int {
	r1, _, err := this.apimap["lua_next"].Call(L, uintptr(idx))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) lua_concat(L uintptr, n int) {
	_, _, err := this.apimap["lua_concat"].Call(L, uintptr(n))
	this.lasterror = err
}

//func(this*luaDLL) lua_getallocf (L, void **ud);
//func(this*luaDLL)lua_setallocf (L, lua_Alloc f, void *ud);
/* Functions to be called by the debuger in specific events
 * typedef void (*lua_Hook) (lua_State *L, lua_Debug *ar);
 */
//func(this*luaDLL)lua_getstack (L uintptr,  level int, lua_Debug *ar);
//func(this*luaDLL)lua_getinfo (L uintptr, const char *what, lua_Debug *ar);
//func(this*luaDLL)lua_getlocal (L uintptr, const lua_Debug *ar, int n);
//func(this*luaDLL)lua_setlocal (L uintptr, const lua_Debug *ar, int n);
//func(this*luaDLL)lua_getupvalue (L uintptr, int funcindex, int n);
//func(this*luaDLL)lua_setupvalue (L uintptr, int funcindex, int n);
//func(this*luaDLL)lua_sethook (L, lua_Hook func, int mask, int count);
//func(this*luaDLL)Hook lua_gethook (lua_State *L);
//func(this*luaDLL)lua_gethookmask (L uintpttr);
//func(this*luaDLL)lua_gethookcount (L uinptr);
/* From Lua 5.2. */
func (this *luaDLL) lua_upvalueid(L uintptr, idx int, n int) uintptr {
	r1, _, err := this.apimap["lua_upvalueid"].Call(L, uintptr(idx), uintptr(n))
	this.lasterror = err
	return r1
}

//LUA_API void lua_upvaluejoin (lua_State *L, idx int1, int n1, idx int2, int n2);
//LUA_API int lua_loadx (lua_State *L, lua_Reader reader, void *dt, const char *chunkname, const char *mode);

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>lauxlib<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func (this *luaDLL) luaL_openlib(L uintptr, libname string, l uintptr, nup int) {
	plibname := unsafe.Pointer(C.CString(libname))
	_, _, err := this.apimap["luaL_openlib"].Call(L, uintptr(plibname), uintptr(l), uintptr(nup))
	this.lasterror = err
}
func (this *luaDLL) luaL_register(L uintptr, libname string, l uintptr) {
	plibname := unsafe.Pointer(C.CString(libname))
	_, _, err := this.apimap["luaL_register"].Call(L, uintptr(plibname), uintptr(l))
	this.lasterror = err
}
func (this *luaDLL) luaL_register2(L uintptr, libname string, funcname string, l uintptr) {
	plibname := unsafe.Pointer(C.CString(libname))
	pfuncname := unsafe.Pointer(C.CString(funcname))
	defer func() {
		C.free(plibname)
		C.free(pfuncname)
	}()
	_, _, err := this.apimap["luaL_register2"].Call(L, uintptr(plibname), uintptr(pfuncname), uintptr(l))
	this.lasterror = err
}
func (this *luaDLL) luaL_getmetafield(L uintptr, obj int, e string) int {
	pe := unsafe.Pointer(C.CString(e))
	r1, _, err := this.apimap["luaL_getmetafield"].Call(L, uintptr(obj), uintptr(pe))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaL_callmeta(L uintptr, obj int, e string) int {
	pe := unsafe.Pointer(C.CString(e))
	r1, _, err := this.apimap["luaL_callmeta"].Call(L, uintptr(obj), uintptr(pe))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaL_typerror(L uintptr, narg int, tname string) int {
	ptname := unsafe.Pointer(C.CString(tname))
	r1, _, err := this.apimap["lua_upvalueid"].Call(L, uintptr(narg), uintptr(ptname))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaL_argerror(L uintptr, numarg int, extramsg string) int {
	pextramsg := unsafe.Pointer(C.CString(extramsg))
	r1, _, err := this.apimap["luaL_argerror"].Call(L, uintptr(numarg), uintptr(pextramsg))
	this.lasterror = err
	return int(r1)
}

//func(this*luaDLL) const char *(luaL_checklstring) (L uintptr, int numArg,size_t *l);
//func(this*luaDLL) const char *(luaL_optlstring) (L uintptr, int numArg,const char *def, size_t *l);
//func(this*luaDLL) lua_Number (luaL_checknumber) (L uintptr, int numArg);
//func(this*luaDLL) lua_Number (luaL_optnumber) (L uintptr, int nArg, lua_Number def);
//func(this*luaDLL) lua_Integer (luaL_checkinteger) (L uintptr, int numArg);
//func(this*luaDLL) lua_Integer (luaL_optinteger) (L uintptr, int nArg,lua_Integer def);
//func(this*luaDLL) void (luaL_checkstack) (L uintptr, int sz, const char *msg);
//func(this*luaDLL) void (luaL_checktype) (L uintptr, int narg, int t);
//func(this*luaDLL) void (luaL_checkany) (L uintptr, int narg);
//func(this*luaDLL) int   (luaL_newmetatable) (L uintptr, const char *tname);
//func(this*luaDLL) void *(luaL_checkudata) (L uintptr, int ud, const char *tname);
//func(this*luaDLL) void (luaL_where) (L uintptr, int lvl);
//func(this*luaDLL) int (luaL_error) (L uintptr, const char *fmt, ...);
//func(this*luaDLL) int (luaL_checkoption) (L uintptr, int narg, const char *def, const char *const lst[]);
//func(this*luaDLL) int (luaL_ref) (L uintptr, int t);
//func(this*luaDLL) void (luaL_unref) (L uintptr, int t, int ref);
func (this *luaDLL) luaL_loadfile(L uintptr, filename string) int {
	pfilename := unsafe.Pointer(C.CString(filename))
	r1, _, err := this.apimap["luaL_loadfile"].Call(L, uintptr(pfilename))
	this.lasterror = err
	return int(r1)
}

func (this *luaDLL) luaL_loadbuffer(L uintptr, buff string, name string) int {
	pbuffer := unsafe.Pointer(C.CString(buff))
	pname := unsafe.Pointer(C.CString(name))
	r1, _, err := this.apimap["luaL_loadfile"].Call(L, uintptr(pbuffer), uintptr(len(buff)), uintptr(pname))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaL_loadstring(L uintptr, s string) int {
	ps := unsafe.Pointer(C.CString(s))
	r1, _, err := this.apimap["luaL_loadfile"].Call(L, uintptr(ps))
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaL_newstate() uintptr {
	r1, _, err := this.apimap["luaL_newstate"].Call()
	this.lasterror = err
	return r1
}

//func(this*luaDLL) const char *(luaL_gsub) (L uintptr, const char *s, const char *p,const char *r);
//func(this*luaDLL) const char *(luaL_findtable) (L uintptr, idx int,const char *fname, int szhint);
///* From Lua 5.2. */
//func(this*luaDLL) int luaL_fileresult(L uintptr, int stat, const char *fname);
//func(this*luaDLL) int luaL_execresult(L uintptr, int stat);
//func(this*luaDLL) int (luaL_loadfilex) (L uintptr, const char *filename, const char *mode);
//func(this*luaDLL) int (luaL_loadbufferx) (L uintptr, const char *buff, size_t sz,const char *name, const char *mode);
//func(this*luaDLL) void luaL_traceback (L uintptr, lua_State *L1, const char *msg,int level);

func (this *luaDLL) luaopen_base(L uintptr) int {
	r1, _, err := this.apimap["luaopen_base"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_math(L uintptr) int {
	r1, _, err := this.apimap["luaopen_math"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_string(L uintptr) int {
	r1, _, err := this.apimap["luaopen_string"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_table(L uintptr) int {
	r1, _, err := this.apimap["luaopen_table"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_io(L uintptr) int {
	r1, _, err := this.apimap["luaopen_io"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_os(L uintptr) int {
	r1, _, err := this.apimap["luaopen_os"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_package(L uintptr) int {
	r1, _, err := this.apimap["luaopen_package"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_debug(L uintptr) int {
	r1, _, err := this.apimap["luaopen_debug"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_bit(L uintptr) int {
	r1, _, err := this.apimap["luaopen_bit"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_jit(L uintptr) int {
	r1, _, err := this.apimap["luaopen_jit"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaopen_ffi(L uintptr) int {
	r1, _, err := this.apimap["luaopen_ffi"].Call(L)
	this.lasterror = err
	return int(r1)
}
func (this *luaDLL) luaL_openlibs(L uintptr) {
	_, _, err := this.apimap["luaL_openlibs"].Call(L)
	this.lasterror = err
}


func __init_windows(){
	adapter = &luaDLL{}
	adapter.init()
}
