#include "logic.h"
#include "skill_table.h"

std::vector<SkillData> SkillData::records_;

bool SkillData::Load(const char* fileName){
	CSV csv;
	if (!csv.LoadFile(fileName)){
		return false;
	}

	for (size_t r = 0; r < csv.GetRowMax(); ++r){
		SkillData s;
		s.SkillId = csv.Get<int32_t>(r, "SkillId");
		s.Crit = csv.Get<int32_t>(r, "Crit");
		s.Damage= csv.Get<int32_t>(r, "Damage");

		records_.push_back(s);
	}
	return true;
}

const SkillData* SkillData::GetSkillById(int32_t id){
	for (size_t i = 0; i < records_.size(); ++i)
	{
		if (records_[i].SkillId == id)
			return &records_[i];
	}
	return nullptr;
}