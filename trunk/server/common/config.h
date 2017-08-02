#ifndef ___CONFIG_H__
#define ___CONFIG_H__

#if defined(_WIN32) || defined(_WIN64) || defined(MSVC)
#	ifndef _CRT_SECURE_NO_WARNINGS
#		define _CRT_SECURE_NO_WARNINGS
#	endif
#pragma warning(error : 4258)
#pragma warning(error : 4715)
#endif // (_WIN32) || defined(_WIN64) || defined(MSVC)


#include <stdlib.h>
#include <stdint.h>
#include <float.h>
#include <stdarg.h>
#include <memory>
#include <array>
#include <vector>
#include <string>
#include <set>
#include <map>
#include <unordered_map>
#include <queue>
#include <stack>
#include <algorithm>
#include <iostream>
#include <sstream>
#include <thread>
#include <fstream>

#include <boost/asio.hpp>
#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <boost/enable_shared_from_this.hpp>  



#endif //end ifndef __RPC_CONFIG_H__