#ifndef __BATTLE_H__
#define __BATTLE_H__
#include "logic.h"

class Battle 
	:public boost::enable_shared_from_this<Battle>{
	static std::vector<boost::shared_ptr<Battle> > battles_;
public:
	static void CreateBattle(boost::shared_ptr<class GamePlayer> pL, boost::shared_ptr<class GamePlayer> pR);
	static void UpdateBattleList();
public:
	int32_t GetBattleId();
	void SetupPosition(boost::shared_ptr<class GamePlayer> player, std::vector<COM_BattlePosition> &posList);
	void Update();
private:
	void release();
	bool calcWinner();
private:
	int32_t round_;
	COM_BattleReport report_;
	std::vector< boost::shared_ptr<class GameUnit> > unitList_;
	std::vector< boost::shared_ptr<class GamePlayer> > playerList_;
};

#endif