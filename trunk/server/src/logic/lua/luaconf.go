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
