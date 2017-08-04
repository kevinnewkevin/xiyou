#include "logic.h"
#include "player.h"
#include "entity.h"
#include "battle.h"
#include "session.h"
#include "unit_table.h"

std::vector<boost::shared_ptr<GamePlayer> > GamePlayer::playerList_;
std::unordered_map<int64_t, boost::shared_ptr<GamePlayer> > GamePlayer::playerIdTable_;
std::unordered_map<std::string, boost::shared_ptr<GamePlayer> > GamePlayer::playerNameTable_;

boost::shared_ptr<GamePlayer> GamePlayer::CreatePlayer(boost::shared_ptr<class Session> session, const std::string &name, int32_t tmpId){
	boost::shared_ptr<GamePlayer> player = boost::make_shared<GamePlayer>();
	player->Init(tmpId, name);
	playerList_.push_back(player);
	playerIdTable_[player->GetPlayerId()] = player;
	playerNameTable_[name] = player;
	return player;
}
boost::shared_ptr<GamePlayer> GamePlayer::FindPlayer(int64_t playerId){
	if (playerIdTable_.find(playerId) == playerIdTable_.end())
		return nullptr;
	return playerIdTable_[playerId];
}
boost::shared_ptr<GamePlayer> GamePlayer::FindPlayer(const std::string &playerName){
	if (playerNameTable_.find(playerName) == playerNameTable_.end())
		return nullptr;
	return playerNameTable_[playerName];
}

//////////////////////////////////////////////////////////////////////////
void GamePlayer::Init(int32_t tmpId, const std::string &name){
	BOOST_ASSERT_MSG(UnitData::GetUnitById(tmpId) != nullptr, "Can not find tmpId");
	playerName_ = name;
	myUnit_ = boost::make_shared<GameUnit>();
	myUnit_->Init(tmpId);
	///
	boost::shared_ptr<GameUnit> unit = boost::make_shared<GameUnit>();
	unit->Init(2);
	unitList_.push_back(unit);
	unit = boost::make_shared<GameUnit>();
	unit->Init(3);
	unitList_.push_back(unit);
}

void GamePlayer::Init(const COM_Player &inst){
	
}

void GamePlayer::Save(COM_Player &inst){
	inst.InstId = myUnit_->GetInstId();
	inst.Name = playerName_;
	myUnit_->Save(inst.Unit);
	for (auto var : unitList_){
		COM_Unit unit;
		var->Save(unit);
		inst.Employees.push_back(unit);
	}
}

int64_t GamePlayer::GetPlayerId(){
	return myUnit_->GetInstId();
}

boost::shared_ptr<GameUnit> GamePlayer::GetUnit(int64_t instId){
	if (myUnit_->GetInstId() == instId)
		return myUnit_;
	else {
		for (size_t i = 0; i < unitList_.size(); ++i){
			if (unitList_[i]->GetInstId() == instId){
				return unitList_[i];
			}
		}
	}
	return nullptr;
}

boost::shared_ptr<GameUnit> GamePlayer::GetBattleUnit(int64_t instId){
	if (myUnit_->GetInstId() == instId)
		return myUnit_;
	for (size_t i = 0; i < battleUnitList_.size(); ++i){
		if (battleUnitList_[i] == instId){
			return GetUnit(instId);
		}
	}
	return nullptr;
}

void GamePlayer::SetBattleUnit(int64_t instId){
	if (GetUnit(instId) == nullptr)
		return;
	if (GetBattleUnit(instId) != nullptr)
		return;
	if (battleUnitList_.size() > GAME_PLAYER_BATTLE_UNIT_MAX)
		return;
	battleUnitList_.push_back(instId);
}

void GamePlayer::SetBattle(boost::shared_ptr<class Battle> battle){
	battleInst_ = battle;
}

boost::shared_ptr<Battle> GamePlayer::GetBattlle(){
	return battleInst_;
}

void GamePlayer::JoinBattle(){
	if (GetBattlle() != nullptr)
		return;

	static std::vector<boost::shared_ptr<GamePlayer> > tmp;

	tmp.push_back(shared_from_this());

	if (tmp.size() >= 2){
		tmp.clear();
	}

}

void GamePlayer::SetupBattle(std::vector<COM_BattlePosition> &positionList){
	if (GetBattlle() == nullptr)
		return;

	GetBattlle()->SetupPosition(shared_from_this(), positionList);
}

void GamePlayer::SendBattleReport(COM_BattleReport &report){
	session_->BattleReport(report);
}