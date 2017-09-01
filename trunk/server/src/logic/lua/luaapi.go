package lua

/*
#include <stddef.h>
#include <stdlib.h>
#cgo windows LDFLAGS: -L.  lua51.dll
#cgo linux LDFLAGS: -L. liblua51
typedef int (*lua_CFunction) (void *L);
//
// functions that read/write blocks when loading/dumping Lua chunks
//
typedef const char * (*lua_Reader) (void *L, void *ud, size_t *sz);

typedef int (*lua_Writer) (void *L, const void* p, size_t sz, void* ud);


//
// prototype for memory-allocation functions
//
typedef void * (*lua_Alloc) (void *ud, void *ptr, size_t osize, size_t nsize);

// type of numbers in Lua
typedef double lua_Number;


// type for integer functions
typedef ptrdiff_t lua_Integer;


extern void *(lua_newstate) (lua_Alloc f, void *ud);
extern void       (lua_close) (void *L);
extern void *(lua_newthread) (void *L);

extern lua_CFunction (lua_atpanic) (void *L, lua_CFunction panicf);


//
// basic stack manipulation
//
extern int   (lua_gettop) (void *L);
extern void  (lua_settop) (void *L, int idx);
extern void  (lua_pushvalue) (void *L, int idx);
extern void  (lua_remove) (void *L, int idx);
extern void  (lua_insert) (void *L, int idx);
extern void  (lua_replace) (void *L, int idx);
extern int   (lua_checkstack) (void *L, int sz);

extern void  (lua_xmove) (void *from, void *to, int n);


//
// access functions (stack -> C)
//

extern int             (lua_isnumber) (void *L, int idx);
extern int             (lua_isstring) (void *L, int idx);
extern int             (lua_iscfunction) (void *L, int idx);
extern int             (lua_isuserdata) (void *L, int idx);
extern int             (lua_type) (void *L, int idx);
extern const char     *(lua_typename) (void *L, int tp);

extern int            (lua_equal) (void *L, int idx1, int idx2);
extern int            (lua_rawequal) (void *L, int idx1, int idx2);
extern int            (lua_lessthan) (void *L, int idx1, int idx2);

extern lua_Number      (lua_tonumber) (void *L, int idx);
extern lua_Integer     (lua_tointeger) (void *L, int idx);
extern int             (lua_toboolean) (void *L, int idx);
extern const char     *(lua_tolstring) (void *L, int idx, void *len);
extern size_t          (lua_objlen) (void *L, int idx);
extern lua_CFunction   (lua_tocfunction) (void *L, int idx);
extern void	       *(lua_touserdata) (void *L, int idx);
extern void      *(lua_tothread) (void *L, int idx);
extern const void     *(lua_topointer) (void *L, int idx);


//
// push functions (C -> stack)
//
extern void  (lua_pushnil) (void *L);
extern void  (lua_pushnumber) (void *L, lua_Number n);
extern void  (lua_pushinteger) (void *L, lua_Integer n);
extern void  (lua_pushlstring) (void *L, const char *s, size_t l);
extern void  (lua_pushstring) (void *L, const char *s);
//extern const char *(lua_pushvfstring) (void *L, const char *fmt,va_list argp);
extern const char *(lua_pushfstring) (void *L, const char *fmt, ...);
extern void  (lua_pushcclosure) (void *L, lua_CFunction fn, int n);
extern void  (lua_pushboolean) (void *L, int b);
extern void  (lua_pushlightuserdata) (void *L, void *p);
extern int   (lua_pushthread) (void *L);


//
// get functions (Lua -> stack)
//
extern void  (lua_gettable) (void *L, int idx);
extern void  (lua_getfield) (void *L, int idx, const char *k);
extern void  (lua_rawget) (void *L, int idx);
extern void  (lua_rawgeti) (void *L, int idx, int n);
extern void  (lua_createtable) (void *L, int narr, int nrec);
extern void *(lua_newuserdata) (void *L, size_t sz);
extern int   (lua_getmetatable) (void *L, int objindex);
extern void  (lua_getfenv) (void *L, int idx);


//
// set functions (stack -> Lua)
//
extern void  (lua_settable) (void *L, int idx);
extern void  (lua_setfield) (void *L, int idx, const char *k);
extern void  (lua_rawset) (void *L, int idx);
extern void  (lua_rawseti) (void *L, int idx, int n);
extern int   (lua_setmetatable) (void *L, int objindex);
extern int   (lua_setfenv) (void *L, int idx);


//
// `load' and `call' functions (load and run Lua code)
//
extern void  (lua_call) (void *L, int nargs, int nresults);
extern int   (lua_pcall) (void *L, int nargs, int nresults, int errfunc);
extern int   (lua_cpcall) (void *L, lua_CFunction func, void *ud);
extern int   (lua_load) (void *L, lua_Reader reader, void *dt,const char *chunkname);

extern int (lua_dump) (void *L, lua_Writer writer, void *data);


//
// coroutine functions
//
extern int  (lua_yield) (void *L, int nresults);
extern int  (lua_resume) (void *L, int narg);
extern int  (lua_status) (void *L);

//
// garbage-collection function and options
//

extern int (lua_gc) (void *L, int what, int data);


//
// miscellaneous functions
//

extern int   (lua_error) (void *L);

extern int   (lua_next) (void *L, int idx);

extern void  (lua_concat) (void *L, int n);

extern lua_Alloc (lua_getallocf) (void *L, void **ud);
extern void lua_setallocf (void *L, lua_Alloc f, void *ud);



// hack
extern void lua_setlevel	(void *from, void *to);


//
// {======================================================================
// Debug API
// =======================================================================
//

typedef struct lua_Debug lua_Debug;  // activation record


// Functions to be called by the debuger in specific events
typedef void (*lua_Hook) (void *L, lua_Debug *ar);


extern int lua_getstack (void *L, int level, lua_Debug *ar);
extern int lua_getinfo (void *L, const char *what, lua_Debug *ar);
extern const char *lua_getlocal (void *L, const lua_Debug *ar, int n);
extern const char *lua_setlocal (void *L, const lua_Debug *ar, int n);
extern const char *lua_getupvalue (void *L, int funcindex, int n);
extern const char *lua_setupvalue (void *L, int funcindex, int n);
extern int lua_sethook (void *L, lua_Hook func, int mask, int count);
extern lua_Hook lua_gethook (void *L);
extern int lua_gethookmask (void *L);
extern int lua_gethookcount (void *L);

// From Lua 5.2.
extern void *lua_upvalueid (void *L, int idx, int n);
extern void lua_upvaluejoin (void *L, int idx1, int n1, int idx2, int n2);
extern int lua_loadx (void *L, lua_Reader reader, void *dt,
const char *chunkname, const char *mode);


////////////////////////////////////////////////////////////////////////////////////////////////////////////

typedef struct luaL_Reg luaL_Reg;
typedef struct luaL_Buffer luaL_Buffer;


extern void (luaL_openlib) (void *L, const char *libname, const luaL_Reg *l, int nup);
extern void (luaL_register) (void *L, const char *libname,const luaL_Reg *l);
extern void (luaL_register2) (void *L, const char *libname, const char * funcname, lua_CFunction func);

extern int (luaL_getmetafield) (void *L, int obj, const char *e);
extern int (luaL_callmeta) (void *L, int obj, const char *e);
extern int (luaL_typerror) (void *L, int narg, const char *tname);
extern int (luaL_argerror) (void *L, int numarg, const char *extramsg);
extern const char *(luaL_checklstring) (void *L, int numArg, size_t *l);
extern const char *(luaL_optlstring) (void *L, int numArg,const char *def, size_t *l);
extern lua_Number (luaL_checknumber) (void *L, int numArg);
extern lua_Number (luaL_optnumber) (void *L, int nArg, lua_Number def);

extern lua_Integer (luaL_checkinteger) (void *L, int numArg);
extern lua_Integer (luaL_optinteger) (void *L, int nArg, lua_Integer def);

extern void (luaL_checkstack) (void *L, int sz, const char *msg);
extern void (luaL_checktype) (void *L, int narg, int t);
extern void (luaL_checkany) (void *L, int narg);

extern int   (luaL_newmetatable) (void *L, const char *tname);
extern void *(luaL_checkudata) (void *L, int ud, const char *tname);

extern void (luaL_where) (void *L, int lvl);
extern int (luaL_error) (void *L, const char *fmt, ...);

extern int (luaL_checkoption) (void *L, int narg, const char *def,const char *const lst[]);

extern int (luaL_ref) (void *L, int t);
extern void (luaL_unref) (void *L, int t, int ref);

extern int (luaL_loadfile) (void *L, const char *filename);
extern int (luaL_loadbuffer) (void *L, const char *buff, size_t sz,const char *name);
extern int (luaL_loadstring) (void *L, const char *s);

extern void *(luaL_newstate) (void);


extern const char *(luaL_gsub) (void *L, const char *s, const char *p,const char *r);

extern const char *(luaL_findtable) (void *L, int idx,const char *fname, int szhint);

// From Lua 5.2.
extern int luaL_fileresult(void *L, int stat, const char *fname);
extern int luaL_execresult(void *L, int stat);
extern int (luaL_loadfilex) (void *L, const char *filename,const char *mode);
extern int (luaL_loadbufferx) (void *L, const char *buff, size_t sz,const char *name, const char *mode);
extern void luaL_traceback (void *L, void *L1, const char *msg,int level);



//
// {======================================================
// Generic Buffer manipulation
// =======================================================
//


extern void (luaL_buffinit) (void *L, luaL_Buffer *B);
extern char *(luaL_prepbuffer) (luaL_Buffer *B);
extern void (luaL_addlstring) (luaL_Buffer *B, const char *s, size_t l);
extern void (luaL_addstring) (luaL_Buffer *B, const char *s);
extern void (luaL_addvalue) (luaL_Buffer *B);
extern void (luaL_pushresult) (luaL_Buffer *B);


extern int luaopen_base(void *L);
extern int luaopen_math(void *L);
extern int luaopen_string(void *L);
extern int luaopen_table(void *L);
extern int luaopen_io(void *L);
extern int luaopen_os(void *L);
extern int luaopen_package(void *L);
extern int luaopen_debug(void *L);
extern int luaopen_bit(void *L);
extern int luaopen_jit(void *L);
extern int luaopen_ffi(void *L);

*/
import "C"
import (
	"unsafe"
)

func lua_close(L uintptr) {
	C.lua_close(unsafe.Pointer(L))
}

// basic stack manipulation
func lua_gettop(L uintptr) int {
	return int(C.lua_gettop(unsafe.Pointer(L)))
}
func lua_settop(L uintptr, idx int) {
	C.lua_settop(unsafe.Pointer(L), C.int(idx))

}
func lua_pushvalue(L uintptr, idx int) {
	C.lua_pushvalue(unsafe.Pointer(L), C.int(idx))
}
func lua_remove(L uintptr, idx int) {
	C.lua_remove(unsafe.Pointer(L), C.int(idx))
}
func lua_insert(L uintptr, idx int) {
	C.lua_insert(unsafe.Pointer(L), C.int(idx))
}
func lua_replace(L uintptr, idx int) {
	C.lua_replace(unsafe.Pointer(L), C.int(idx))
}
func lua_checkstack(L uintptr, sz int) int {
	return int(C.lua_checkstack(unsafe.Pointer(L), C.int(sz)))
}

//access functions (stack -> C)
func lua_isnumber(L uintptr, idx int) int {
	return int(C.lua_isnumber(unsafe.Pointer(L), C.int(idx)))
}
func lua_isstring(L uintptr, idx int) int {
	return int(C.lua_isstring(unsafe.Pointer(L), C.int(idx)))
}
func lua_iscfunction(L uintptr, idx int) int {
	return int(C.lua_iscfunction(unsafe.Pointer(L), C.int(idx)))
}

func lua_isuserdata(L uintptr, idx int) int {
	return int(C.lua_isuserdata(unsafe.Pointer(L), C.int(idx)))
}
func lua_type(L uintptr, idx int) int {
	return int(C.lua_type(unsafe.Pointer(L), C.int(idx)))
}
func lua_typename(L uintptr, tp int) string {
	return C.GoString(C.lua_typename(unsafe.Pointer(L), C.int(tp)))
}
func lua_equal(L uintptr, idx1 int, idx2 int) int {
	return int(C.lua_equal(unsafe.Pointer(L), C.int(idx1), C.int(idx2)))
}
func lua_rawequal(L uintptr, idx1 int, idx2 int) int {
	return int(C.lua_rawequal(unsafe.Pointer(L), C.int(idx1), C.int(idx2)))
}
func lua_lessthan(L uintptr, idx1 int, idx2 int) int {
	return int(C.lua_lessthan(unsafe.Pointer(L), C.int(idx1), C.int(idx2)))
}
func lua_tonumber(L uintptr, idx int) float64 {
	return float64(C.lua_tonumber(unsafe.Pointer(L), C.int(idx)))
}
func lua_tointeger(L uintptr, idx int) uintptr {
	return uintptr(C.lua_tointeger(unsafe.Pointer(L), C.int(idx)))
}
func lua_toboolean(L uintptr, idx int) int {
	return int(C.lua_toboolean(unsafe.Pointer(L), C.int(idx)))
}

func lua_tolstring(L uintptr, idx int, len *int) string {
	return C.GoString(C.lua_tolstring(unsafe.Pointer(L), C.int(idx), unsafe.Pointer(len)))
}
func lua_objlen(L uintptr, idx int) int {
	return int(C.lua_objlen(unsafe.Pointer(L), C.int(idx)))
}

func lua_tocfunction(L uintptr, idx int) unsafe.Pointer {
	return unsafe.Pointer(C.lua_tocfunction(unsafe.Pointer(L), C.int(idx)))
}
func lua_touserdata(L uintptr, idx int) uintptr {
	return uintptr(C.lua_touserdata(unsafe.Pointer(L), C.int(idx)))
}
func lua_tothread(L uintptr, idx int) uintptr {
	return uintptr(C.lua_tothread(unsafe.Pointer(L), C.int(idx)))
}
func lua_topointer(L uintptr, idx int) uintptr {
	return uintptr(C.lua_topointer(unsafe.Pointer(L), C.int(idx)))
}

//push functions (C -> stack)
func lua_pushnil(L uintptr) {
	C.lua_pushnil(unsafe.Pointer(L))

}
func lua_pushnumber(L uintptr, n float64) {
	C.lua_pushnumber(unsafe.Pointer(L), C.lua_Number(n))

}
func lua_pushinteger(L uintptr, n uintptr) {
	C.lua_pushinteger(unsafe.Pointer(L), C.lua_Integer(n))
}

//func (lua_pushlstring) (L uintptr, const char *s, size_t l);
func lua_pushstring(L uintptr, s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.lua_pushstring(unsafe.Pointer(L), p)
}

//func(this*luaDLL) (lua_pushvfstring) (L uintptr, const char *fmt, va_list argp) string
//func(this*luaDLL) (lua_pushfstring) (L uintptr, const char *fmt, ...) string
//func(this*luaDLL) (lua_pushcclosure) (L uintptr, lua_CFunction fn, int n)
func lua_pushboolean(L uintptr, b int) {
	C.lua_pushinteger(unsafe.Pointer(L), C.lua_Integer(b))
}

//func(lua_pushlightuserdata) (L uintptr, void *p);
//func (lua_pushthread) (L uintptr);
//get functions (Lua -> stack)
func lua_gettable(L uintptr, idx int) {
	C.lua_gettable(unsafe.Pointer(L), C.int(idx))
}
func lua_getfield(L uintptr, idx int, k string) {
	p := C.CString(k)
	defer C.free(unsafe.Pointer(p))
	C.lua_getfield(unsafe.Pointer(L), C.int(idx), p)

}
func lua_rawget(L uintptr, idx int) {
	C.lua_rawget(unsafe.Pointer(L), C.int(idx))
}
func lua_rawgeti(L uintptr, idx int, n int) {
	C.lua_rawgeti(unsafe.Pointer(L), C.int(idx), C.int(n))
}
func lua_createtable(L uintptr, narr int, nrec int) {
	C.lua_createtable(unsafe.Pointer(L), C.int(narr), C.int(nrec))
}

//func(this*luaDLL)lua_newuserdata (L,  sz int) uintptr;
func lua_getmetatable(L uintptr, objindex int) int {
	return int(C.lua_getmetatable(unsafe.Pointer(L), C.int(objindex)))
}
func lua_getfenv(L uintptr, idx int) {
	C.lua_getfenv(unsafe.Pointer(L), C.int(idx))
}

//set functions (stack -> Lua)
func lua_settable(L uintptr, idx int) {
	C.lua_settable(unsafe.Pointer(L), C.int(idx))
}
func lua_setfield(L uintptr, idx int, k string) {
	p := C.CString(k)
	defer C.free(unsafe.Pointer(p))
	C.lua_setfield(unsafe.Pointer(L), C.int(idx), p)
}
func lua_rawset(L uintptr, idx int) {
	C.lua_rawset(unsafe.Pointer(L), C.int(idx))
}
func lua_rawseti(L uintptr, idx int, n int) {
	C.lua_rawseti(unsafe.Pointer(L), C.int(idx), C.int(n))
}
func lua_setmetatable(L uintptr, objindex int) int {
	return int(C.lua_setmetatable(unsafe.Pointer(L), C.int(objindex)))
}
func lua_setfenv(L uintptr, idx int) int {
	return int(C.lua_setfenv(unsafe.Pointer(L), C.int(idx)))
}

//`load' and `call' functions (load and run Lua code)
func lua_call(L uintptr, nargs int, nresults int) {
	C.lua_call(unsafe.Pointer(L), C.int(nargs), C.int(nresults))
}
func lua_pcall(L uintptr, nargs int, nresults int, errfunc int) int {
	return int(C.lua_pcall(unsafe.Pointer(L), C.int(nargs), C.int(nresults), C.int(errfunc)))
}

//func lua_cpcall (L uintptr,  fun uintptr,  ud uintptr){
//	C.lua_cpcall(unsafe.Pointer(L),unsafe.Pointer(fun),unsafe.Pointer(ud))
//}
//func(this*luaDLL)lua_load (L uintptr, lua_Reader reader, void *dt, const char *chunkname);
//func(this*luaDLL)lua_dump (L uintptr, lua_Writer writer, void *data);
//coroutine functions
//LUA_API int  (lua_yield) (lua_State *L, int nresults);
//LUA_API int  (lua_resume) (lua_State *L, int narg);
//LUA_API int  (lua_status) (lua_State *L);
//garbage-collection function and options
func lua_gc(L uintptr, what int, data int) int {
	return int(C.lua_gc(unsafe.Pointer(L), C.int(what), C.int(data)))
}

//miscellaneous functions
func lua_error(L uintptr) int {
	return int(C.lua_error(unsafe.Pointer(L)))
}
func lua_next(L uintptr, idx int) int {
	return int(C.lua_next(unsafe.Pointer(L), C.int(idx)))
}
func lua_concat(L uintptr, n int) {
	C.lua_concat(unsafe.Pointer(L), C.int(n))
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
//func  lua_upvalueid(L uintptr, idx int, n int) uintptr {
//	return uintptr(C.lua_call(unsafe.Pointer(L),C.int(idx),C.int(n)))
//}

//LUA_API void lua_upvaluejoin (lua_State *L, idx int1, int n1, idx int2, int n2);
//LUA_API int lua_loadx (lua_State *L, lua_Reader reader, void *dt, const char *chunkname, const char *mode);

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>lauxlib<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func luaL_getmetafield(L uintptr, obj int, e string) int {
	p := C.CString(e)
	defer C.free(unsafe.Pointer(p))
	return int(C.luaL_getmetafield(unsafe.Pointer(L), C.int(obj), p))
}
func luaL_callmeta(L uintptr, obj int, e string) int {
	p := C.CString(e)
	defer C.free(unsafe.Pointer(p))
	return int(C.luaL_callmeta(unsafe.Pointer(L), C.int(obj), p))
}
func luaL_typerror(L uintptr, narg int, tname string) int {
	p := C.CString(tname)
	defer C.free(unsafe.Pointer(p))
	return int(C.luaL_typerror(unsafe.Pointer(L), C.int(narg), p))
}
func luaL_argerror(L uintptr, numarg int, extramsg string) int {
	p := C.CString(extramsg)
	defer C.free(unsafe.Pointer(p))
	return int(C.luaL_argerror(unsafe.Pointer(L), C.int(numarg), p))
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
func luaL_loadfile(L uintptr, filename string) int {
	p := C.CString(filename)
	defer C.free(unsafe.Pointer(p))
	return int(C.luaL_loadfile(unsafe.Pointer(L), p))
}

//func  luaL_loadbuffer(L uintptr, buff string, name string) int {
//	p := C.CString(buff)
//	p1 := C.CString(name)
//	defer C.free(unsafe.Pointer(p))
//	defer C.free(unsafe.Pointer(p1))
//	return int(C.luaL_loadbuffer(unsafe.Pointer(L),p,p1))
//}

func luaL_loadstring(L uintptr, s string) int {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	return int(C.luaL_loadstring(unsafe.Pointer(L), p))
}

func luaL_newstate() uintptr {
	return uintptr(C.luaL_newstate())
}

//func(this*luaDLL) const char *(luaL_gsub) (L uintptr, const char *s, const char *p,const char *r);
//func(this*luaDLL) const char *(luaL_findtable) (L uintptr, idx int,const char *fname, int szhint);
///* From Lua 5.2. */
//func(this*luaDLL) int luaL_fileresult(L uintptr, int stat, const char *fname);
//func(this*luaDLL) int luaL_execresult(L uintptr, int stat);
//func(this*luaDLL) int (luaL_loadfilex) (L uintptr, const char *filename, const char *mode);
//func(this*luaDLL) int (luaL_loadbufferx) (L uintptr, const char *buff, size_t sz,const char *name, const char *mode);
//func(this*luaDLL) void luaL_traceback (L uintptr, lua_State *L1, const char *msg,int level);

func luaL_register2 (L uintptr, libname, funname string, fun uintptr){
	libnameC :=C.CString(libname)
	funnameC := C.CString(funname)
	C.luaL_register2(unsafe.Pointer(L),libnameC,funnameC,C.lua_CFunction(unsafe.Pointer(fun)))
	C.free(unsafe.Pointer(libnameC))
	C.free(unsafe.Pointer(funnameC))
}

func lua_atpanic (L , panicf uintptr) uintptr{
	return uintptr(unsafe.Pointer(C.lua_atpanic(unsafe.Pointer(L),C.lua_CFunction(unsafe.Pointer(panicf)))))
}

func luaopen_base(L uintptr) int {
	return int(C.luaopen_base(unsafe.Pointer(L)))
}
func luaopen_math(L uintptr) int {
	return int(C.luaopen_math(unsafe.Pointer(L)))
}
func luaopen_string(L uintptr) int {
	return int(C.luaopen_string(unsafe.Pointer(L)))
}
func luaopen_table(L uintptr) int {
	return int(C.luaopen_table(unsafe.Pointer(L)))
}
func luaopen_io(L uintptr) int {
	return int(C.luaopen_io(unsafe.Pointer(L)))
}
func luaopen_os(L uintptr) int {
	return int(C.luaopen_os(unsafe.Pointer(L)))
}
func luaopen_package(L uintptr) int {
	return int(C.luaopen_package(unsafe.Pointer(L)))
}
func luaopen_debug(L uintptr) int {
	return int(C.luaopen_debug(unsafe.Pointer(L)))
}
func luaopen_bit(L uintptr) int {
	return int(C.luaopen_bit(unsafe.Pointer(L)))
}
func luaopen_jit(L uintptr) int {
	return int(C.luaopen_jit(unsafe.Pointer(L)))
}
func luaopen_ffi(L uintptr) int {
	return int(C.luaopen_ffi(unsafe.Pointer(L)))
}
