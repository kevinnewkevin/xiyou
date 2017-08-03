#ifndef __SCRIPT_H__
#define __SCRIPT_H__
#include "config.h"

extern "C"
{
#include "lua.h"
#include "lualib.h"
#include "lauxlib.h"
}


#define CHK_LUA_STACK StackChecker __CHEKCER(state_);
#define CHK_LUA_ERROR(STMT) do\
{\
	int32_t ___CHECKER = STMT;\
	if(___CHECKER != 0){\
		BOOST_LOG_TRIVIAL(info) << lua_tostring(state_,-1);\
		lua_pop(state_,1);\
		return;\
			}\
} while (0);

class Script 
	: public boost::enable_shared_from_this<Script>{

	static std::unordered_map<lua_State*, boost::shared_ptr<Script> > holder_;
public:
	static boost::shared_ptr<Script> GetHandler(lua_State* state);
public:
	void Init();
	void Fini();
	
	void SetTable(int32_t index);

	void PushStack(int32_t val);
	void PushStack(float val);
	void PushStack(bool val);
	void PushStack(const std::string& val);

	template<class T>
	T GetStack(int32_t idx);
	
	void SetGlobal(const std::string& varName, int32_t val);
	void SetGlobal(const std::string& varName, float val);
	void SetGlobal(const std::string& varName, bool val);
	void SetGlobal(const std::string& varName, const std::string& val);

	template<class T>
	T GetGlobal(const std::string& varName);
	
	void DelGlobal(const std::string& varName);

	template<class T> void DefEnumer();

	void LoadFile(const std::string& filleName);
	void LoadString(const std::string& chunk);

	void OpenAPI(lua_CFunction func, const std::string &apiName, const std::string &libName);

	template<class T>
	bool Call(const std::string &func, std::string& err, const T& param);
	template<class T, class R>
	bool Call(const std::string &func, std::string& err, const T& param, R& result);
private:
	struct StackChecker{
		StackChecker(lua_State *s);
		~StackChecker();
	private:
		lua_State* state_;
		int32_t    top_;
	};
	lua_State* state_;

};

template<>
inline int32_t Script::GetStack<int32_t>(int32_t idx){
	BOOST_ASSERT(state_);
	lua_Number num = lua_tonumber(state_, idx);
	return static_cast<int32_t>(num);
}
template<>
inline float Script::GetStack<float>(int32_t idx){
	BOOST_ASSERT(state_);
	lua_Number num = lua_tonumber(state_, idx);
	return static_cast<float>(num);
}
template<>
inline bool Script::GetStack<bool>(int32_t idx){
	BOOST_ASSERT(state_);
	int32_t num = lua_toboolean(state_, idx);
	return !!num;
}
template<>
inline std::string Script::GetStack<std::string>(int32_t idx){
	BOOST_ASSERT(state_);
	const char* ret = lua_tostring(state_, idx);
	return ret;
}

template<>
inline int32_t Script::GetGlobal<int32_t>(const std::string& varName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	SetTable(LUA_GLOBALSINDEX);
	int32_t val = static_cast<int32_t>(lua_tonumber(state_, -1));
	lua_pop(state_, -1);
	return val;
}
template<>
inline float Script::GetGlobal(const std::string& varName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	SetTable(LUA_GLOBALSINDEX);
	float val = static_cast<float>(lua_tonumber(state_, -1));
	lua_pop(state_, -1);
	return val;
}
template<>
inline bool Script::GetGlobal(const std::string& varName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	SetTable(LUA_GLOBALSINDEX);
	bool val = !!lua_toboolean(state_, -1);
	lua_pop(state_, -1);
	return val;
}
template<>
inline std::string Script::GetGlobal(const std::string& varName){
	BOOST_ASSERT(state_);
	BOOST_ASSERT(!varName.empty() && varName != "");
	CHK_LUA_STACK;
	PushStack(varName);
	SetTable(LUA_GLOBALSINDEX);
	const char* val = lua_tostring(state_, -1);
	lua_pop(state_, -1);
	return val;
}
template<class T> 
inline void Script::DefEnumer(){
	BOOST_ASSERT(state_);
	for (auto var : T::NAMES){
		PushStack(var);
		PushStack(T::GetId(var));
		SetTable(LUA_GLOBALSINDEX);
	}
}

#endif