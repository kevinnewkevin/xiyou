#include "logic.h"
#include "entity.h"
#include "skill.h"
#include "player.h"
#include "session.h"
#include "unit_table.h"

void GameUnit::Init(int32_t tmpId){

	static int64_t instIdGen = 0;

	const UnitData* t = UnitData::GetUnitById(tmpId);
	if (t == nullptr){
		BOOST_LOG_TRIVIAL(info) << "aa";
	}
	unitId_ = t->Id;
	instId_ = ++instIdGen;
	std::copy(t->IProp.begin(), t->IProp.end(), ipropertyList_.begin());
	std::copy(t->CProp.begin(), t->CProp.end(), cpropertyList_.begin());
	
	for (size_t i = 0; i < t->Skills.size(); ++i){
		if (t->Skills[i]){
			skills_[i] = boost::make_shared<Skill>();
			skills_[i]->Init(t->Skills[i]);
		}
	}

}

void GameUnit::Init(const COM_Unit &inst){

}

void GameUnit::Save(COM_Unit &inst){
	inst.UnitId = unitId_;
	inst.InstId = instId_;
	inst.IProperties.resize(IPropertyType::IPT_MAX);
	inst.CProperties.resize(CPropertyType::CPT_MAX);
	std::copy(ipropertyList_.begin(), ipropertyList_.end(), inst.IProperties.begin());
	std::copy(cpropertyList_.begin(), cpropertyList_.end(), inst.CProperties.begin());
}

void GameUnit::Save(COM_BattleUnit &inst){
	inst.UnitId = unitId_;
	inst.InstId = instId_;
	inst.HP = static_cast<int32_t>(GetCProperty(CPropertyType::CPT_HP));
	inst.Position = battlePosition_;
}

void GameUnit::CastSkill(boost::shared_ptr<class Battle> battle){

}