#ifndef __ENTITY_H__
#define __ENTITY_H__

#include "logic.h"

class GameUnit {
public:
	void Init(int32_t tmpId);
	void Init(const COM_Unit &inst);
	void Save(COM_Unit &inst);
	void Save(COM_BattleUnit &inst);
public:
	int32_t GetUnitId(){
		return unitId_;
	}
	int64_t GetInstId(){
		return instId_;
	}
	boost::shared_ptr<class GamePlayer> GetOwner(){
		return owner_;
	}
	int32_t GetIProperty(IPropertyType::Type type){
		if (type <= IPropertyType::IPT_MIN || type >= IPropertyType::IPT_MAX){
			return 0;
		}
		return ipropertyList_[type];
	}
	float GetCProperty(CPropertyType::Type type){
		if (type <= CPropertyType::CPT_MIN || type >= CPropertyType::CPT_MAX){
			return 0;
		}
		return cpropertyList_[type];
	}
	BattlePosotion::Type GetBattlePosition(){
		return battlePosition_;
	}
	void SetBattlePosition(BattlePosotion::Type bp){
		battlePosition_ = bp;
	}
	bool IsDeadth(){
		return GetCProperty(CPropertyType::CPT_HP) <= 0.F;
	}
	
	void CastSkill(boost::shared_ptr<class Battle> battle);
private:
	int32_t													unitId_;
	int64_t													instId_;
	std::array<int32_t, IPropertyType::IPT_MAX>				ipropertyList_;
	std::array<float, CPropertyType::CPT_MAX>				cpropertyList_;
	std::array< boost::shared_ptr<class Skill>, SKILL_MAX>	skills_;
	boost::shared_ptr<class GamePlayer>						owner_; //ÓµÓÐÕß
	
	BattlePosotion::Type									battlePosition_;
};

#endif