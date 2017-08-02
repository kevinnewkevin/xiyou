#include "logic.h"
#include "unit_table.h"
std::vector<UnitData> UnitData::records_;
bool UnitData::Load(const char* fileName){
	CSV csv;
	if (!csv.LoadFile(fileName)){
		return false;
	}

	for (size_t r = 0; r < csv.GetRowMax(); ++r){
		UnitData u;
		u.Id = csv.Get<int32_t>(r, "UnitId");
		u.DispId = csv.Get<int32_t>(r, "DisplayId");
		
		u.IProp[IPropertyType::IPT_PHYLE] = csv.Get<int32_t>(r, IPropertyType::ToName(IPropertyType::IPT_PHYLE));
		u.IProp[IPropertyType::IPT_LEVEL] = csv.Get<int32_t>(r, IPropertyType::ToName(IPropertyType::IPT_LEVEL));
		u.IProp[IPropertyType::IPT_COPPER] = csv.Get<int32_t>(r, IPropertyType::ToName(IPropertyType::IPT_COPPER));
		u.IProp[IPropertyType::IPT_SILVER] = csv.Get<int32_t>(r, IPropertyType::ToName(IPropertyType::IPT_SILVER));
		u.IProp[IPropertyType::IPT_GOLD] = csv.Get<int32_t>(r, IPropertyType::ToName(IPropertyType::IPT_GOLD));

		u.CProp[CPropertyType::CPT_HP] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_HP));
		u.CProp[CPropertyType::CPT_ATK] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_ATK));
		u.CProp[CPropertyType::CPT_DEF] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_DEF));
		u.CProp[CPropertyType::CPT_MAGIC_ATK] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_MAGIC_ATK));
		u.CProp[CPropertyType::CPT_MAGIC_DEF] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_MAGIC_DEF));
		u.CProp[CPropertyType::CPT_AGILE] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_AGILE));
		u.CProp[CPropertyType::CPT_KILL] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_KILL));
		u.CProp[CPropertyType::CPT_CRIT] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_CRIT));
		u.CProp[CPropertyType::CPT_COUNTER_ATTACK] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_COUNTER_ATTACK));
		u.CProp[CPropertyType::CPT_SPUTTERING] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_SPUTTERING));
		u.CProp[CPropertyType::CPT_DOUBLE_HIT] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_DOUBLE_HIT));
		u.CProp[CPropertyType::CPT_RECOVERY] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_RECOVERY));
		u.CProp[CPropertyType::CPT_REFLEX] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_REFLEX));
		u.CProp[CPropertyType::CPT_SUCK_BLOOD] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_SUCK_BLOOD));
		u.CProp[CPropertyType::CPT_INCANTER] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_INCANTER));
		u.CProp[CPropertyType::CPT_RESISTANCE] = csv.Get<float>(r, CPropertyType::ToName(CPropertyType::CPT_RESISTANCE));

		u.Skills[0] = csv.Get<int32_t>(r, "Skill1");
		u.Skills[1] = csv.Get<int32_t>(r, "Skill2");
		u.Skills[2] = csv.Get<int32_t>(r, "Skill3");
		u.Skills[3] = csv.Get<int32_t>(r, "Skill4");
		
		records_.push_back(u);
	}
	return true;
}

const UnitData* UnitData::GetUnitById(int32_t id){
	for (size_t i = 0; i < records_.size(); ++i){
		if (records_[i].Id == id){
			return &records_[i];
		}
	}
	return nullptr;
}