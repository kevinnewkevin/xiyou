#ifndef __LOGIC_H__
#define __LOGIC_H__

#include "config.h"
#include "prpc.h"
#include "prpc.gen.h"
#include "csv.h"
//#define BOOST_USE_WINAPI_VERSION WINVER
#include <boost/log/core.hpp>
#include <boost/log/trivial.hpp>
#include <boost/log/expressions.hpp>
enum {
	SKILL_MAX = 4, //最大技能数量
	BATTLE_UNIT_MAX = 6, //战斗中最大牌数
	GAME_PLAYER_BATTLE_UNIT_MAX = 10
};


#endif