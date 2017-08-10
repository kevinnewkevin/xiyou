package lua

//#include <stdlib.h>
import "C"
import (
	"syscall"
	"unsafe"
)

const (
	LUA_VERSION     = "Lua 5.1"
	LUA_RELEASE     = "Lua 5.1.4"
	LUA_VERSION_NUM = 501

	LUA_COPYRIGHT = "Copyright (C) 1994-2008 Lua.org, PUC-Rio"
	LUA_AUTHORS   = "R. Ierusalimschy, L. H. de Figueiredo & W. Celes"
	/* mark for precompiled code (`<esc>Lua') */
	LUA_SIGNATURE = "\033Lua"

	/* option for multiple returns in `lua_pcall' and `lua_call' */
	LUA_MULTRET = (-1)

	LUA_REGISTRYINDEX = (-10000)
	LUA_ENVIRONINDEX  = (-10001)
	LUA_GLOBALSINDEX  = (-10002)

	LUA_GCSTOP       = 0
	LUA_GCRESTART    = 1
	LUA_GCCOLLECT    = 2
	LUA_GCCOUNT      = 3
	LUA_GCCOUNTB     = 4
	LUA_GCSTEP       = 5
	LUA_GCSETPAUSE   = 6
	LUA_GCSETSTEPMUL = 7

	LUA_HOOKCALL    = 0
	LUA_HOOKRET     = 1
	LUA_HOOKLINE    = 2
	LUA_HOOKCOUNT   = 3
	LUA_HOOKTAILRET = 4

	LUA_MASKCALL  = (1 << LUA_HOOKCALL)
	LUA_MASKRET   = (1 << LUA_HOOKRET)
	LUA_MASKLINE  = (1 << LUA_HOOKLINE)
	LUA_MASKCOUNT = (1 << LUA_HOOKCOUNT)
)

var (
	luaApiNameList []string = []string{
		//lua.h
		// state manipulation
		"lua_newstate",  //LUA_API lua_State *(lua_newstate) (lua_Alloc f, void *ud);
		"lua_close",     //LUA_API void       (lua_close) (lua_State *L);
		"lua_newthread", //LUA_API lua_State *(lua_newthread) (lua_State *L);
		"lua_atpanic",   //LUA_API lua_CFunction (lua_atpanic) (lua_State *L, lua_CFunction panicf);
		// basic stack manipulation
		"lua_gettop",     //LUA_API int   (lua_gettop) (lua_State *L);
		"lua_settop",     //LUA_API void  (lua_settop) (lua_State *L, int idx);
		"lua_pushvalue",  //LUA_API void  (lua_pushvalue) (lua_State *L, int idx);
		"lua_remove",     //LUA_API void  (lua_remove) (lua_State *L, int idx);
		"lua_insert",     //LUA_API void  (lua_insert) (lua_State *L, int idx);
		"lua_replace",    //LUA_API void  (lua_replace) (lua_State *L, int idx);
		"lua_checkstack", //LUA_API int   (lua_checkstack) (lua_State *L, int sz);
		"lua_xmove",      //LUA_API void  (lua_xmove) (lua_State *from, lua_State *to, int n);
		//access functions (stack -> C)
		"lua_isnumber",    //LUA_API int             (lua_isnumber) (lua_State *L, int idx);
		"lua_isstring",    //LUA_API int             (lua_isstring) (lua_State *L, int idx);
		"lua_iscfunction", //LUA_API int             (lua_iscfunction) (lua_State *L, int idx);
		"lua_isuserdata",  //LUA_API int             (lua_isuserdata) (lua_State *L, int idx);
		"lua_type",        //LUA_API int             (lua_type) (lua_State *L, int idx);
		"lua_typename",    //LUA_API const char     *(lua_typename) (lua_State *L, int tp);
		"lua_equal",       //LUA_API int            (lua_equal) (lua_State *L, int idx1, int idx2);
		"lua_rawequal",    //LUA_API int            (lua_rawequal) (lua_State *L, int idx1, int idx2);
		"lua_lessthan",    //LUA_API int            (lua_lessthan) (lua_State *L, int idx1, int idx2);
		"lua_tonumber",    //LUA_API lua_Number      (lua_tonumber) (lua_State *L, int idx);
		"lua_tointeger",   //LUA_API lua_Integer     (lua_tointeger) (lua_State *L, int idx);
		"lua_toboolean",   //LUA_API int             (lua_toboolean) (lua_State *L, int idx);
		"lua_tolstring",   //LUA_API const char     *(lua_tolstring) (lua_State *L, int idx, size_t *len);
		"lua_objlen",      //LUA_API size_t          (lua_objlen) (lua_State *L, int idx);
		"lua_tocfunction", //LUA_API lua_CFunction   (lua_tocfunction) (lua_State *L, int idx);
		"lua_touserdata",  //LUA_API void	       *(lua_touserdata) (lua_State *L, int idx);
		"lua_tothread",    //LUA_API lua_State      *(lua_tothread) (lua_State *L, int idx);
		"lua_topointer",   //LUA_API const void     *(lua_topointer) (lua_State *L, int idx);
		//push functions (C -> stack)
		"lua_pushnil",           //LUA_API void  (lua_pushnil) (lua_State *L);
		"lua_pushnumber",        //LUA_API void  (lua_pushnumber) (lua_State *L, lua_Number n);
		"lua_pushinteger",       //LUA_API void  (lua_pushinteger) (lua_State *L, lua_Integer n);
		"lua_pushlstring",       //LUA_API void  (lua_pushlstring) (lua_State *L, const char *s, size_t l);
		"lua_pushstring",        //LUA_API void  (lua_pushstring) (lua_State *L, const char *s);
		"lua_pushvfstring",      //LUA_API const char *(lua_pushvfstring) (lua_State *L, const char *fmt, va_list argp);
		"lua_pushfstring",       //LUA_API const char *(lua_pushfstring) (lua_State *L, const char *fmt, ...);
		"lua_pushcclosure",      //LUA_API void  (lua_pushcclosure) (lua_State *L, lua_CFunction fn, int n);
		"lua_pushboolean",       //LUA_API void  (lua_pushboolean) (lua_State *L, int b);
		"lua_pushlightuserdata", //LUA_API void  (lua_pushlightuserdata) (lua_State *L, void *p);
		"lua_pushthread",        //LUA_API int   (lua_pushthread) (lua_State *L);
		//get functions (Lua -> stack)
		"lua_gettable",     //LUA_API void  (lua_gettable) (lua_State *L, int idx);
		"lua_getfield",     //LUA_API void  (lua_getfield) (lua_State *L, int idx, const char *k);
		"lua_rawget",       //LUA_API void  (lua_rawget) (lua_State *L, int idx);
		"lua_rawgeti",      //LUA_API void  (lua_rawgeti) (lua_State *L, int idx, int n);
		"lua_createtable",  //LUA_API void  (lua_createtable) (lua_State *L, int narr, int nrec);
		"lua_newuserdata",  //LUA_API void *(lua_newuserdata) (lua_State *L, size_t sz);
		"lua_getmetatable", //LUA_API int   (lua_getmetatable) (lua_State *L, int objindex);
		"lua_getfenv",      //LUA_API void  (lua_getfenv) (lua_State *L, int idx);
		//set functions (stack -> Lua)
		"lua_settable",     //LUA_API void  (lua_settable) (lua_State *L, int idx);
		"lua_setfield",     //LUA_API void  (lua_setfield) (lua_State *L, int idx, const char *k);
		"lua_rawset",       //LUA_API void  (lua_rawset) (lua_State *L, int idx);
		"lua_rawseti",      //LUA_API void  (lua_rawseti) (lua_State *L, int idx, int n);
		"lua_setmetatable", //LUA_API int   (lua_setmetatable) (lua_State *L, int objindex);
		"lua_setfenv",      //LUA_API int   (lua_setfenv) (lua_State *L, int idx);
		//`load' and `call' functions (load and run Lua code)
		"lua_call",   //LUA_API void  (lua_call) (lua_State *L, int nargs, int nresults);
		"lua_pcall",  //LUA_API int   (lua_pcall) (lua_State *L, int nargs, int nresults, int errfunc);
		"lua_cpcall", //LUA_API int   (lua_cpcall) (lua_State *L, lua_CFunction func, void *ud);
		"lua_load",   //LUA_API int   (lua_load) (lua_State *L, lua_Reader reader, void *dt, const char *chunkname);
		"lua_dump",   //LUA_API int (lua_dump) (lua_State *L, lua_Writer writer, void *data);
		//coroutine functions
		"lua_yield",  //LUA_API int  (lua_yield) (lua_State *L, int nresults);
		"lua_resume", //LUA_API int  (lua_resume) (lua_State *L, int narg);
		"lua_status", //LUA_API int  (lua_status) (lua_State *L);
		//garbage-collection function and options
		"lua_gc", //LUA_API int (lua_gc) (lua_State *L, int what, int data);
		//miscellaneous functions
		"lua_error",     //LUA_API int   (lua_error) (lua_State *L);
		"lua_next",      //LUA_API int   (lua_next) (lua_State *L, int idx);
		"lua_concat",    //LUA_API void  (lua_concat) (lua_State *L, int n);
		"lua_getallocf", //LUA_API lua_Alloc (lua_getallocf) (lua_State *L, void **ud);
		"lua_setallocf", //LUA_API void lua_setallocf (lua_State *L, lua_Alloc f, void *ud);
		/* Functions to be called by the debuger in specific events
		 * typedef void (*lua_Hook) (lua_State *L, lua_Debug *ar);
		 */
		"lua_getstack",     //LUA_API int lua_getstack (lua_State *L, int level, lua_Debug *ar);
		"lua_getinfo",      //LUA_API int lua_getinfo (lua_State *L, const char *what, lua_Debug *ar);
		"lua_getlocal",     //LUA_API const char *lua_getlocal (lua_State *L, const lua_Debug *ar, int n);
		"lua_setlocal",     //LUA_API const char *lua_setlocal (lua_State *L, const lua_Debug *ar, int n);
		"lua_getupvalue",   //LUA_API const char *lua_getupvalue (lua_State *L, int funcindex, int n);
		"lua_setupvalue",   //LUA_API const char *lua_setupvalue (lua_State *L, int funcindex, int n);
		"lua_sethook",      //LUA_API int lua_sethook (lua_State *L, lua_Hook func, int mask, int count);
		"lua_gethook",      //LUA_API lua_Hook lua_gethook (lua_State *L);
		"lua_gethookmask",  //LUA_API int lua_gethookmask (lua_State *L);
		"lua_gethookcount", //LUA_API int lua_gethookcount (lua_State *L);
		/* From Lua 5.2. */
		"lua_upvalueid",   //LUA_API void *lua_upvalueid (lua_State *L, int idx, int n);
		"lua_upvaluejoin", //LUA_API void lua_upvaluejoin (lua_State *L, int idx1, int n1, int idx2, int n2);
		"lua_loadx",       //LUA_API int lua_loadx (lua_State *L, lua_Reader reader, void *dt, const char *chunkname, const char *mode);
		//lauxlib.h
		"luaL_openlib",      //LUALIB_API void (luaL_openlib) (lua_State *L, const char *libname, const luaL_Reg *l, int nup);
		"luaL_register",     //LUALIB_API void (luaL_register) (lua_State *L, const char *libname,const luaL_Reg *l);
		"luaL_register2",    //LUALIB_API void (luaL_register2) (lua_State *L, const char *libname, const char * funcname, lua_CFunction func);
		"luaL_getmetafield", //LUALIB_API int (luaL_getmetafield) (lua_State *L, int obj, const char *e);
		"luaL_callmeta",     //LUALIB_API int (luaL_callmeta) (lua_State *L, int obj, const char *e);
		"luaL_typerror",     //LUALIB_API int (luaL_typerror) (lua_State *L, int narg, const char *tname);
		"luaL_argerror",     //LUALIB_API int (luaL_argerror) (lua_State *L, int numarg, const char *extramsg);
		"luaL_checklstring", //LUALIB_API const char *(luaL_checklstring) (lua_State *L, int numArg,size_t *l);
		"luaL_optlstring",   //LUALIB_API const char *(luaL_optlstring) (lua_State *L, int numArg,const char *def, size_t *l);
		"luaL_checknumber",  //LUALIB_API lua_Number (luaL_checknumber) (lua_State *L, int numArg);
		"luaL_optnumber",    //LUALIB_API lua_Number (luaL_optnumber) (lua_State *L, int nArg, lua_Number def);
		"luaL_checkinteger", //LUALIB_API lua_Integer (luaL_checkinteger) (lua_State *L, int numArg);
		"luaL_optinteger",   //LUALIB_API lua_Integer (luaL_optinteger) (lua_State *L, int nArg,lua_Integer def);
		"luaL_checkstack",   //LUALIB_API void (luaL_checkstack) (lua_State *L, int sz, const char *msg);
		"luaL_checktype",    //LUALIB_API void (luaL_checktype) (lua_State *L, int narg, int t);
		"luaL_checkany",     //LUALIB_API void (luaL_checkany) (lua_State *L, int narg);
		"luaL_newmetatable", //LUALIB_API int   (luaL_newmetatable) (lua_State *L, const char *tname);
		"luaL_checkudata",   //LUALIB_API void *(luaL_checkudata) (lua_State *L, int ud, const char *tname);
		"luaL_where",        //LUALIB_API void (luaL_where) (lua_State *L, int lvl);
		"luaL_error",        //LUALIB_API int (luaL_error) (lua_State *L, const char *fmt, ...);
		"luaL_checkoption",  //LUALIB_API int (luaL_checkoption) (lua_State *L, int narg, const char *def, const char *const lst[]);
		"luaL_ref",          //LUALIB_API int (luaL_ref) (lua_State *L, int t);
		"luaL_unref",        //LUALIB_API void (luaL_unref) (lua_State *L, int t, int ref);
		"luaL_loadfile",     //LUALIB_API int (luaL_loadfile) (lua_State *L, const char *filename);
		"luaL_loadbuffer",   //LUALIB_API int (luaL_loadbuffer) (lua_State *L, const char *buff, size_t sz,const char *name);
		"luaL_loadstring",   //LUALIB_API int (luaL_loadstring) (lua_State *L, const char *s);
		"luaL_newstate",     //LUALIB_API lua_State *(luaL_newstate) (void);
		"luaL_gsub",         //LUALIB_API const char *(luaL_gsub) (lua_State *L, const char *s, const char *p,const char *r);
		"luaL_findtable",    //LUALIB_API const char *(luaL_findtable) (lua_State *L, int idx,const char *fname, int szhint);
		/* From Lua 5.2. */
		"luaL_fileresult",  //LUALIB_API int luaL_fileresult(lua_State *L, int stat, const char *fname);
		"luaL_execresult",  //LUALIB_API int luaL_execresult(lua_State *L, int stat);
		"luaL_loadfilex",   //LUALIB_API int (luaL_loadfilex) (lua_State *L, const char *filename, const char *mode);
		"luaL_loadbufferx", //LUALIB_API int (luaL_loadbufferx) (lua_State *L, const char *buff, size_t sz,const char *name, const char *mode);
		"luaL_traceback",   //LUALIB_API void luaL_traceback (lua_State *L, lua_State *L1, const char *msg,int level);
		//lualib.h
		"luaopen_base",
		"luaopen_math",
		"luaopen_string",
		"luaopen_table",
		"luaopen_io",
		"luaopen_os",
		"luaopen_package",
		"luaopen_debug",
		"luaopen_bit",
		"luaopen_jit",
		"luaopen_ffi",
	}

	luaapi luaAPI
)

type luaAPI interface {
	init()
	fini()
	last_error() error
	// state manipulation
	//lua_newstate (f uintptr, ud uintptr) uintptr
	lua_close(L uintptr)
	//lua_newthread(L uintptr) uintptr
	lua_atpanic(L uintptr, panicf uintptr) uintptr
	// basic stack manipulation
	lua_gettop(L uintptr) int
	lua_settop(L uintptr, idx int)
	lua_pushvalue(L uintptr, idx int)
	lua_remove(L uintptr, idx int)
	lua_insert(L uintptr, idx int)
	lua_replace(L uintptr, idx int)
	lua_checkstack(L uintptr, sz int) int
	//lua_xmove(from uintptr, to uintptr, n uintptr)
	//access functions stack -> C)
	lua_isnumber(L uintptr, idx int) int
	lua_isstring(L uintptr, idx int) int
	lua_iscfunction(L uintptr, idx int) int
	lua_isuserdata(L uintptr, idx int) int
	lua_type(L uintptr, idx int) int
	lua_typename(L uintptr, tp int) string
	lua_equal(L uintptr, idx1 int, idx2 int) int
	lua_rawequal(L uintptr, idx1 int, idx2 int) int
	lua_lessthan(L uintptr, idx1 int, idx2 int) int
	lua_tonumber(L uintptr, idx int) float64
	lua_tointeger(L uintptr, idx int) int
	lua_toboolean(L uintptr, idx int) int
	lua_tolstring(L uintptr, idx int, len int) string
	lua_objlen(L uintptr, idx int) int
	//lua_tocfunction(L uintptr, idx int) uintptr
	//lua_touserdata(L uintptr, idx int) uintptr
	//lua_tothread(L uintptr, idx int) uintptr
	//lua_topointer(L uintptr, idx int) uintptr
	//(push functions C -> stack)
	lua_pushnil(L uintptr)
	lua_pushnumber(L uintptr, n float64)
	lua_pushinteger(L uintptr, n int)
	//lua_pushlstring(L uintptr, s string, l int)
	lua_pushstring(L uintptr, s string)
	//lua_pushvfstring(L uintptr, const char *fmt, va_list argp) string
	//lua_pushfstring(L uintptr, const char *fmt, ...) string
	//lua_pushcclosure(L uintptr, lua_CFunction fn, int n)
	lua_pushboolean(L uintptr, b int)
	//lua_pushlightuserdata(L uintptr, p uintptr) int
	//lua_pushthread(L uintptr)
	//(get functions Lua -> stack)
	lua_gettable(L uintptr, idx int)
	lua_getfield(L uintptr, idx int, k string)
	lua_rawget(L uintptr, idx int)
	lua_rawgeti(L uintptr, idx int, n int)
	lua_createtable(L uintptr, narr int, nrec int)
	//lua_newuserdata(L uintptr, sz int) uintptr
	lua_getmetatable(L uintptr, objindex int) int
	lua_getfenv(L uintptr, idx int)
	//set functions stack -> Lua)
	lua_settable(L uintptr, idx int)
	lua_setfield(L uintptr, idx int, k string)
	lua_rawset(L uintptr, idx int)
	lua_rawseti(L uintptr, idx int, n int)
	lua_setmetatable(L uintptr, objindex int) int
	lua_setfenv(L uintptr, idx int) int
	//`load' and `call' functions load and run Lua code)
	lua_call(L uintptr, nargs int, nresults int)
	lua_pcall(L uintptr, nargs int, nresults int, errfunc int) int
	//lua_cpcall(L uintptr, lua_CFunction func, void *ud)                        int
	//lua_load(L uintptr, lua_Reader reader, void *dt, const char *chunkname)    int
	//lua_dump(L uintptr, lua_Writer writer, void *data)                           int
	//coroutine functions
	//lua_yield(L uintptr, nresults int) int
	//lua_resume(L uintptr, narg int) int
	//lua_status(L uintptr) int
	//garbage-collection function and options
	lua_gc(L uintptr, what int, data int) int
	//miscellaneous functions
	lua_error(L uintptr) int
	lua_next(L uintptr, idx int) int
	lua_concat(L uintptr, n int)
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
	lua_upvalueid(L uintptr, idx int, n int) uintptr
	//lua_upvaluejoin (L uintptr, idx int1,  n1 int, idx int2, int n2)
	//lua_loadx (L uintptr, lua_Reader reader, void *dt, const char *chunkname, const char *mode) int
	//lauxlib.h
	luaL_openlib(L uintptr, libname string, l uintptr, nup int)
	luaL_register(L uintptr, libname string, l uintptr)
	luaL_register2(L uintptr, libname string, funcname string, l uintptr)
	luaL_getmetafield(L uintptr, obj int, e string) int
	luaL_callmeta(L uintptr, obj int, e string) int
	luaL_typerror(L uintptr, narg int, tname string) int
	luaL_argerror(L uintptr, numarg int, extramsg string) int
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
	luaL_loadfile(L uintptr, filename string) int
	luaL_loadbuffer(L uintptr, buff string, name string) int
	luaL_loadstring(L uintptr, s string) int
	luaL_newstate() uintptr
	//luaL_gsub(L uintptr, s string, p string, r string) string
	//luaL_findtable(L uintptr, idx int, fname string, szhint int) string
	/* From Lua 5.2. */
	//luaL_fileresult(L uintptr, stat int, fname string) int
	//luaL_execresult(L uintptr, stat int) int
	//luaL_loadfilex(L uintptr, filename string, mode string) int
	//luaL_loadbufferx(L uintptr, buff string, name string, mode string) int
	//luaL_traceback(L uintptr, L1 uintptr, msg string, level int)

	//lualib.h
	luaopen_base(L uintptr) int
	luaopen_math(L uintptr) int
	luaopen_string(L uintptr) int
	luaopen_table(L uintptr) int
	luaopen_io(L uintptr) int
	luaopen_os(L uintptr) int
	luaopen_package(L uintptr) int
	luaopen_debug(L uintptr) int
	luaopen_bit(L uintptr) int
	luaopen_jit(L uintptr) int
	luaopen_ffi(L uintptr) int

	luaL_openlibs(L uintptr)
}

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
//func (this *luaDLL)  lua_touserdata (L uintptr, idx int)                  void	       *
//func (this *luaDLL)  lua_tothread (L uintptr, idx int)                    lua_State   *
//func (this *luaDLL)  lua_topointer(L uintptr, idx int)                   const void     *

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
//func(this*luaDLL)lua_gethookmask (L uintptr);
//func(this*luaDLL)lua_gethookcount (L uintptr);
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

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func __init() {
	if luaapi == nil {
		luaapi = &luaDLL{}
		luaapi.init()
	}
}

func __fini() {
	if luaapi != nil {
		luaapi.fini()
		luaapi = nil
	}
}

func luaapi_lasterror() error {
	return luaapi.last_error()
}

func lua_close(L uintptr) {
	luaapi.lua_close(L)
}

func lua_atpanic(L uintptr, panicf uintptr) uintptr {
	return luaapi.lua_atpanic(L, panicf)
}

func lua_gettop(L uintptr) int {
	return luaapi.lua_gettop(L)
}
func lua_settop(L uintptr, idx int) {
	luaapi.lua_settop(L, idx)
}
func lua_pushvalue(L uintptr, idx int) {
	luaapi.lua_pushvalue(L, idx)
}
func lua_remove(L uintptr, idx int) {
	luaapi.lua_remove(L, idx)
}
func lua_insert(L uintptr, idx int) {
	luaapi.lua_insert(L, idx)
}
func lua_replace(L uintptr, idx int) {
	luaapi.lua_replace(L, idx)
}
func lua_checkstack(L uintptr, sz int) int {
	return luaapi.lua_checkstack(L, sz)
}

//lua_xmove(from uintptr, to uintptr, n uintptr)
//access functions stack -> C)
func lua_isnumber(L uintptr, idx int) int {
	return luaapi.lua_isnumber(L, idx)
}
func lua_isstring(L uintptr, idx int) int {
	return luaapi.lua_isstring(L, idx)
}

//func lua_iscfunction(L uintptr, idx int) int
func lua_isuserdata(L uintptr, idx int) int {
	return luaapi.lua_isuserdata(L, idx)
}
func lua_type(L uintptr, idx int) int {
	return luaapi.lua_type(L, idx)
}
func lua_typename(L uintptr, tp int) string {
	return luaapi.lua_typename(L, tp)
}
func lua_equal(L uintptr, idx1 int, idx2 int) int {
	return luaapi.lua_equal(L, idx1, idx2)
}
func lua_rawequal(L uintptr, idx1 int, idx2 int) int {
	return luaapi.lua_rawequal(L, idx1, idx2)
}
func lua_lessthan(L uintptr, idx1 int, idx2 int) int {
	return luaapi.lua_lessthan(L, idx1, idx2)
}
func lua_tonumber(L uintptr, idx int) float64 {
	return luaapi.lua_tonumber(L, idx)
}
func lua_tointeger(L uintptr, idx int) int {
	return luaapi.lua_tointeger(L, idx)
}
func lua_toboolean(L uintptr, idx int) int {
	return luaapi.lua_toboolean(L, idx)
}
func lua_tolstring(L uintptr, idx int, len int) string {
	return luaapi.lua_tolstring(L, idx, len)
}
func lua_objlen(L uintptr, idx int) int {
	return luaapi.lua_objlen(L, idx)
}

//lua_tocfunction(L uintptr, idx int) uintptr
//lua_touserdata(L uintptr, idx int) uintptr
//lua_tothread(L uintptr, idx int) uintptr
//lua_topointer(L uintptr, idx int) uintptr
//(push functions C -> stack)
func lua_pushnil(L uintptr) {
	luaapi.lua_pushnil(L)
}
func lua_pushnumber(L uintptr, n float64) {
	luaapi.lua_pushnumber(L, n)
}
func lua_pushinteger(L uintptr, n int) {
	luaapi.lua_pushinteger(L, n)
}

//lua_pushlstring(L uintptr, s string, l int)
func lua_pushstring(L uintptr, s string) {
	luaapi.lua_pushstring(L, s)
}

//lua_pushvfstring(L uintptr, const char *fmt, va_list argp) string
//lua_pushfstring(L uintptr, const char *fmt, ...) string
//lua_pushcclosure(L uintptr, lua_CFunction fn, int n)
func lua_pushboolean(L uintptr, b int) {
	luaapi.lua_pushboolean(L, b)
}

//lua_pushlightuserdata(L uintptr, p uintptr) int
//lua_pushthread(L uintptr)
//(get functions Lua -> stack)
func lua_gettable(L uintptr, idx int) {
	luaapi.lua_gettable(L, idx)
}
func lua_getfield(L uintptr, idx int, k string) {
	luaapi.lua_getfield(L, idx, k)
}
func lua_rawget(L uintptr, idx int) {
	luaapi.lua_rawget(L, idx)
}
func lua_rawgeti(L uintptr, idx int, n int) {
	luaapi.lua_rawgeti(L, idx, n)
}
func lua_createtable(L uintptr, narr int, nrec int) {
	luaapi.lua_createtable(L, narr, nrec)
}

//lua_newuserdata(L uintptr, sz int) uintptr
func lua_getmetatable(L uintptr, objindex int) int {
	return luaapi.lua_getmetatable(L, objindex)
}
func lua_getfenv(L uintptr, idx int) {
	luaapi.lua_getfenv(L, idx)
}

//set functions stack -> Lua)
func lua_settable(L uintptr, idx int) {
	luaapi.lua_settable(L, idx)
}
func lua_setfield(L uintptr, idx int, k string) {
	luaapi.lua_setfield(L, idx, k)
}
func lua_rawset(L uintptr, idx int) {
	luaapi.lua_rawset(L, idx)
}
func lua_rawseti(L uintptr, idx int, n int) {
	luaapi.lua_rawseti(L, idx, n)
}
func lua_setmetatable(L uintptr, objindex int) int {
	return luaapi.lua_setmetatable(L, objindex)
}
func lua_setfenv(L uintptr, idx int) int {
	return luaapi.lua_setfenv(L, idx)
}

//`load' and `call' functions load and run Lua code)
func lua_call(L uintptr, nargs int, nresults int) {
	luaapi.lua_call(L, nargs, nresults)
}
func lua_pcall(L uintptr, nargs int, nresults int, errfunc int) int {
	return luaapi.lua_pcall(L, nargs, nresults, errfunc)
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
	return luaapi.lua_gc(L, what, data)
}

//miscellaneous functions
func lua_error(L uintptr) int {
	return luaapi.lua_error(L)
}
func lua_next(L uintptr, idx int) int {
	return luaapi.lua_next(L, idx)
}
func lua_concat(L uintptr, n int) {
	luaapi.lua_concat(L, n)
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
	return luaapi.lua_upvalueid(L, idx, n)
}

//lua_upvaluejoin (L uintptr, idx int1,  n1 int, idx int2, int n2)
//lua_loadx (L uintptr, lua_Reader reader, void *dt, const char *chunkname, const char *mode) int
//lauxlib.h
func luaL_openlib(L uintptr, libname string, l uintptr, nup int) {
	luaapi.luaL_openlib(L, libname, l, nup)
}
func luaL_register(L uintptr, libname string, l uintptr) {
	luaapi.luaL_register(L, libname, l)
}
func luaL_register2(L uintptr, libname string, funcname string, l uintptr) {
	luaapi.luaL_register2(L, libname, funcname, l)
}
func luaL_getmetafield(L uintptr, obj int, e string) int {
	return luaapi.luaL_getmetafield(L, obj, e)
}
func luaL_callmeta(L uintptr, obj int, e string) int {
	return luaapi.luaL_callmeta(L, obj, e)
}
func luaL_typerror(L uintptr, narg int, tname string) int {
	return luaapi.luaL_typerror(L, narg, tname)
}
func luaL_argerror(L uintptr, numarg int, extramsg string) int {
	return luaapi.luaL_argerror(L, numarg, extramsg)
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
	return luaapi.luaL_loadfile(L, filename)
}
func luaL_loadbuffer(L uintptr, buff string, name string) int {
	return luaapi.luaL_loadbuffer(L, buff, name)
}
func luaL_loadstring(L uintptr, s string) int {
	return luaapi.luaL_loadstring(L, s)
}
func luaL_newstate() uintptr {
	return luaapi.luaL_newstate()
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
	return luaapi.luaopen_base(L)
}
func luaopen_math(L uintptr) int {
	return luaapi.luaopen_math(L)
}
func luaopen_string(L uintptr) int {
	return luaapi.luaopen_string(L)
}
func luaopen_table(L uintptr) int {
	return luaapi.luaopen_table(L)
}
func luaopen_io(L uintptr) int {
	return luaapi.luaopen_io(L)
}
func luaopen_os(L uintptr) int {
	return luaapi.luaopen_os(L)
}
func luaopen_package(L uintptr) int {
	return luaapi.luaopen_package(L)
}
func luaopen_debug(L uintptr) int {
	return luaapi.luaopen_debug(L)
}
func luaopen_bit(L uintptr) int {
	return luaapi.luaopen_bit(L)
}
func luaopen_jit(L uintptr) int {
	return luaapi.luaopen_jit(L)
}
func luaopen_ffi(L uintptr) int {
	return luaapi.luaopen_ffi(L)
}
func luaL_openlibs(L uintptr) {
	luaapi.luaL_openlibs(L)
}
