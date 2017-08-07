
#include "lua.h"
#include "lualib.h"
#include "lauxlib.h"
#include "lua.go.h"

void luaL_loadapi(lua_State* L, void* api, const char* funcName, const char* libName ){
	enum {
		MaxLib = 2
	};
	luaL_Reg lib[MaxLib] = {{funcName,(lua_CFunction)api},{NULL,NULL}};
	luaL_openlib(L,libName,lib,0);
	lua_pop(L,1);
}

void luaL_openpanic(lua_State* L, void* pnc){
	lua_atpanic(L ,(lua_CFunction)pnc);
}

const char* luaL_tostring(lua_State* L, int i){
    return lua_tostring(L,i);
}