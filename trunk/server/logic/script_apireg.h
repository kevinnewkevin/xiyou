
#define SCRIPT_API( LIB, NAME ) do{int SCRIPT_API_##LIB##_##NAME( lua_State *L ); OpenAPI(SCRIPT_API_##LIB##_##NAME, #NAME, #LIB );}while(0)