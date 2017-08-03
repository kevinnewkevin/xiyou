#ifndef SCRIPT_API
#define SCRIPT_API(LIB,NAME) int SCRIPT_API_##LIB##_##NAME( lua_State *L )
#endif