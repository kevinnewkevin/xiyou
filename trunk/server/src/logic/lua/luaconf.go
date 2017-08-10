package lua

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

	adapter luaAPI
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
