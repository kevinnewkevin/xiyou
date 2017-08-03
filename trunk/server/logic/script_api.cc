#include "config.h"
#include "script.h"
#include "script_apidef.h"

SCRIPT_API(sys, log){
	boost::shared_ptr<Script> handle = Script::GetHandler(L);
	BOOST_ASSERT(handle != nullptr);
	BOOST_LOG_TRIVIAL(info) << "LUA >> " << handle->GetStack<std::string>(1);
	return 0;
}

SCRIPT_API(sys, loadfile){
	boost::shared_ptr<Script> handle = Script::GetHandler(L);
	BOOST_ASSERT(handle != nullptr);
	std::string fileName = handle->GetStack<std::string>(1);
	handle->LoadFile(fileName);
	return 0;
}