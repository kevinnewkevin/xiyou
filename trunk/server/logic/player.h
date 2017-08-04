#ifndef __PLAYER_H__
#define __PLAYER_H__
#include "logic.h"

class GamePlayer 
	:public boost::enable_shared_from_this<GamePlayer>{
	static std::vector<boost::shared_ptr<GamePlayer> > playerList_;
	static std::unordered_map<int64_t, boost::shared_ptr<GamePlayer> > playerIdTable_;
	static std::unordered_map<std::string, boost::shared_ptr<GamePlayer> > playerNameTable_;
public:
	static boost::shared_ptr<GamePlayer> CreatePlayer(boost::shared_ptr<class Session> session, const std::string &name, int32_t tmpId);
	static boost::shared_ptr<GamePlayer> FindPlayer(int64_t playerId);
	static boost::shared_ptr<GamePlayer> FindPlayer(const std::string &playerName);
	
public:
	void Init(int32_t tmpId, const std::string &name);
	void Init(const COM_Player &inst);
	void Save(COM_Player &inst);
	int64_t GetPlayerId();
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
	std::string											playerName_;
	boost::shared_ptr<class Session>					session_; //链接
	boost::shared_ptr<class GameUnit>					myUnit_;
	std::vector<boost::shared_ptr<class GameUnit> >		unitList_;
	std::vector<int64_t>								battleUnitList_;

	//战斗信息
	bool battleActive_;	
	boost::shared_ptr<class Battle> battleInst_;

};

#endif