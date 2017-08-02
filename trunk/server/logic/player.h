#ifndef __PLAYER_H__
#define __PLAYER_H__
#include "logic.h"
class GamePlayer 
	:public boost::enable_shared_from_this<GamePlayer>{
public:
	void Init(int32_t tmpId, const std::string &name);
	void Init(const COM_Player &inst);
	void Save(COM_Player &inst);

	boost::shared_ptr<class GameUnit> GetMyUnit(){
		return myUnit_;
	}
	boost::shared_ptr<class GameUnit> GetUnit(int64_t instId);
	boost::shared_ptr<class GameUnit> GetBattleUnit(int64_t instId);
	void SetBattleUnit(int64_t instId);

	void SetBattle(boost::shared_ptr<class Battle>);
	boost::shared_ptr<class Battle> GetBattlle();
	void JoinBattle();
	void SetupBattle(std::vector<COM_BattlePosition> &positionList);
	void SetBattleActive(bool isActive){
		battleActive_ = true;
	}
	bool IsBattleActive(){
		return battleActive_;
	}
	void SendBattleReport(COM_BattleReport &report);
private:
	std::string											name_;
	boost::shared_ptr<class Session>					session_; //链接
	boost::shared_ptr<class GameUnit>					myUnit_;
	std::vector<boost::shared_ptr<class GameUnit> >		unitList_;
	std::vector<int64_t>								battleUnitList_;

	//战斗信息
	bool battleActive_;	
	boost::shared_ptr<class Battle> battleInst_;

};

#endif