#include "logic.h"
#include "player.h"
#include "entity.h"
#include "battle.h"

std::vector<boost::shared_ptr<Battle> > Battle::battleList_;
std::unordered_map<int32_t, boost::shared_ptr<Battle> > Battle::battleIdTable_;

void Battle::CreateBattle(boost::shared_ptr<GamePlayer> pL, boost::shared_ptr<GamePlayer> pR){
	if (pL->GetBattlle() != nullptr)
		return;
	if (pR->GetBattlle() != nullptr)
		return;
	static int32_t idGen = 1;
	boost::shared_ptr<Battle> battle = boost::make_shared<Battle>();
	battleList_.push_back(battle);
	battleIdTable_[idGen] = battle;
	battle->battleId_ = idGen++;
	battle->playerList_.push_back(pL);
	battle->playerList_.push_back(pR);

	pL->SetBattle(battle);
	pR->SetBattle(battle);
}

boost::shared_ptr<Battle> Battle::FindBattle(int32_t battleId){
	if (battleIdTable_.find(battleId) == battleIdTable_.end())
		return nullptr;
	return battleIdTable_[battleId];
}

void Battle::UpdateBattleList(){

	for (auto itr = battleList_.begin(); itr != battleList_.end(); ++itr){
		(*itr)->Update();
	}
}

void Battle::SetupPosition(boost::shared_ptr<class GamePlayer> player, std::vector<COM_BattlePosition> &posList){
	if (round_ == 0){
		for (auto var : posList){
			if (player->GetMyUnit()->GetInstId() == var.InstId){
				goto setup_check_success;
			}
		}
		return;// 没有主角卡
	}
setup_check_success:
	for( size_t i = 0; i < posList.size() - 1; ++i) {
		for(size_t j = i + 1; j < posList.size(); ++j){
			
			if(posList[i].InstId == posList[j].InstId){
				return; //有重复卡牌设置
			}
			
			if(posList[i].Position == posList[j].Position){
				return; //有重复位置设置
			}
		}
	}

	for( auto pos : posList){
		for ( auto unit :unitList_){
			if(unit == nullptr){
				continue;
			}
			if(unit->GetInstId() == pos.InstId){
				return; //已经上场
			}
		}

		if(unitList_[pos.Position] != nullptr){
			return; //这个位置上有人
		}
	}

	for(auto pos : posList){
		unitList_[pos.Position] = player->GetBattleUnit(pos.InstId);
		if (unitList_[pos.Position] != nullptr)
			unitList_[pos.Position]->SetBattlePosition((BattlePosotion::Type)pos.Position);
	}
	player->SetBattleActive(true);
}

void Battle::Update(){

	for (auto var : playerList_){
		if (!var->IsBattleActive()){
			return;
		}
	}

	report_ = COM_BattleReport();

	for (auto var : unitList_){
		if (var == nullptr)
			continue;
		if (var->IsDeadth())
			continue;

		COM_BattleUnit battleUnit;
		var->Save(battleUnit);
		report_.UnitList.push_back(battleUnit);
		
		var->CastSkill(shared_from_this());
	}

	for (auto var : playerList_){
		var->SetBattleActive(false);
		var->SendBattleReport(report_);
	}
	
	if (calcWinner()){
		release();
	}
}

bool Battle::calcWinner(){
	return false;
}

void Battle::release(){

}