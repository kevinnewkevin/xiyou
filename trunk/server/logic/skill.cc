#include "logic.h"
#include "skill.h"
#include "skill_table.h"

void Skill::Init(int32_t tmpId){
	BOOST_ASSERT_MSG(SkillData::GetSkillById(tmpId) != nullptr, "can not find skill data");
	const SkillData* skData = SkillData::GetSkillById(tmpId);
	skillId_ = skData->SkillId;
	crit_ = skData->Crit;
	damage_ = skData->Damage;
}

void Skill::Active(boost::shared_ptr<class GameUnit> caster, boost::shared_ptr<class Battle> context){
	
}