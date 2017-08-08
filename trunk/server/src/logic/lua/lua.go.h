#ifndef __LUA_GO_H__
#define __LUA_GO_H__

#include <stdlib.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

/// ADD FOR GOLUA
void luaL_loadapi(lua_State* L, void* api, const char* funcName, const char* libName );
void luaL_openpanic(lua_State* L, void* pnc);
const char* luaL_tostring(lua_State* L, int i);
#endif
