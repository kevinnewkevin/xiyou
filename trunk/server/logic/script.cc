#include "config.h"
#include "script.h"

std::unordered_map<lua_State*, boost::shared_ptr<Script> > Script::holder_;
boost::shared_ptr<Script> Script::GetHandler(lua_State* state){
	BOOST_ASSERT(state);
	if (holder_.find(state) == holder_.end())
		return nullptr;
	return holder_[state];
}

static int __panic(lua_State *L)
{
	(void)L;  /* to avoid warnings */
	BOOST_LOG_TRIVIAL(info) << "PANIC: unprotected error in call to Lua API " << lua_tostring(L, -1);
	return 0;
}

Script::StackChecker::StackChecker(lua_State *s) :state_(s){
	BOOST_ASSERT(state_);
	top_ = lua_gettop(state_);
}
Script::StackChecker::~StackChecker(){
	BOOST_ASSERT(top_ == lua_gettop(state_));
}


void Script::Init(){
	BOOST_LOG_TRIVIAL(trace) << "Init script system";
	BOOST_ASSERT_MSG(state_ = luaL_newstate(), "Init script system failed");
	lua_atpanic(state_, &__panic);

	luaopen_base(state_);
	luaopen_math(state_);
	luaopen_os(state_);
	luaopen_table(state_);
	luaopen_debug(state_);

#include "script_apireg.h"
#include "script_api.h"

	holder_[state_] = shared_from_this();
}

void Script::Fini(){
	BOOST_ASSERT_MSG(state_, "Fini script system failed");
	lua_close(state_);
	holder_.erase(holder_.find(state_));
	state_ = nullptr;
}

void Script::SetTable(int32_t index){
	BOOST_ASSERT(state_);
	lua_settable(state_, index);
}

void Script::PushStack(int32_t val){
	BOOST_ASSERT(state_);
	lua_pushinteger(state_, val);
}
void Script::PushStack(float val){
	BOOST_ASSERT(state_);
	lua_pushnumber(state_, val);
}
void Script::PushStack(bool val){
	BOOST_ASSERT(state_);
	lua_pushboolean(state_, val);
}
void Script::PushStack(const std::string& val){
	BOOST_ASSERT(state_);
	lua_pushstring(state_, val.c_str());
}

void Script::SetGlobal(const std::string& varName, int32_t i){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	PushStack(i);
	SetTable(LUA_GLOBALSINDEX);
}
void Script::SetGlobal(const std::string& varName, float f){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	PushStack(f);
	SetTable(LUA_GLOBALSINDEX);
}
void Script::SetGlobal(const std::string& varName, bool b){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	PushStack(b);
	SetTable(LUA_GLOBALSINDEX);
}
void Script::SetGlobal(const std::string& varName, const std::string& s){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	PushStack(s);
	SetTable(LUA_GLOBALSINDEX);
}

void Script::DelGlobal(const std::string& varName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	lua_pushnil(state_);
	lua_settable(state_, LUA_GLOBALSINDEX);
}

void Script::LoadFile(const std::string& fileName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!fileName.empty() && fileName != "");
	CHK_LUA_ERROR(luaL_loadfile(state_, fileName.c_str()));
	CHK_LUA_ERROR(lua_pcall(state_, 0, 0, 0));
}
void Script::LoadString(const std::string& chunk){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!chunk.empty() && chunk != "");
	CHK_LUA_ERROR(luaL_loadbuffer(state_, chunk.c_str(), chunk.length(), 0));
	CHK_LUA_ERROR(lua_pcall(state_, 0, 0, 0));
}
void Script::OpenAPI(lua_CFunction func, const std::string &apiName, const std::string &libName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(func);
	BOOST_ASSERT(!apiName.empty() && apiName != "");
	//BOOST_ASSERT(!libName.empty() && libName != "");
	luaL_reg lib[2];
	lib[0].name = apiName.c_str();
	lib[0].func = func;
	lib[1].name = nullptr;
	lib[1].func = nullptr;
	luaL_openlib(state_, libName.c_str(), lib, 0);
	lua_pop(state_, 1);
}