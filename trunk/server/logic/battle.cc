#include "logic.h"
#include "player.h"
#include "entity.h"
#include "battle.h"
std::vector<boost::shared_ptr<Battle> > Battle::battles_;

void Battle::CreateBattle(boost::shared_ptr<GamePlayer> pL, boost::shared_ptr<GamePlayer> pR){
	if (pL->GetBattlle() != nullptr)
		return;
	if (pR->GetBattlle() != nullptr)
		return;
	boost::shared_ptr<Battle> battle = boost::make_shared<Battle>();
	battle->playerList_.push_back(pL);
	battle->playerList_.push_back(pR);
	battles_.push_back(battle);
	pL->SetBattle(battle);
	pR->SetBattle(battle);
}

void Battle::UpdateBattleList(){
	BOOST_LOG_TRIVIAL(trace) << "AASD";
	for (auto itr = battles_.begin(); itr != battles_.end(); ++itr){
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