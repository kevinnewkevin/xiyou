
#ifndef __UNIT_TABLE_H__
#define __UNIT_TABLE_H__

#include "logic.h"

struct UnitData{
	int32_t Id;
	int32_t DispId;
	std::array < int32_t, IPropertyType::IPT_MAX > IProp;
	std::array < float, CPropertyType::CPT_MAX > CProp;
	std::array < int32_t, SKILL_MAX > Skills;
	
	static bool Load(const char* fileName);
	static const UnitData* GetUnitById(int32_t id);
	static std::vector<UnitData> records_;
};

#endif