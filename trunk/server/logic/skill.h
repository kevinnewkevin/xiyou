#ifndef __SKILL_H__
#define __SKILL_H__
#include "logic.h"

class SkillAction{

};

class SkillCondition{

};

class Skill{
public:
	void Init(int32_t tmpId);
private:
	int32_t	skillId_;
	int32_t crit_;
	int32_t damage_;
	boost::shared_ptr<SkillAction> action_;
	boost::shared_ptr<SkillCondition> condition_;
};

#endif