#ifndef __SKILL_H__
#define __SKILL_H__
#include "logic.h"

class Selecter{

};

class Action{

};

class Condition{

};


class Skill{
public:
	void Init(int32_t tmpId);

	void Active(boost::shared_ptr<class GameUnit> caster, boost::shared_ptr<class Battle> context);

private:
	int32_t	skillId_;
	int32_t crit_;
	int32_t damage_;
	boost::shared_ptr<Condition>		 condition_;
	boost::shared_ptr<Selecter>			 selecter_;
	boost::shared_ptr<Action>			 action_;
};

#endif