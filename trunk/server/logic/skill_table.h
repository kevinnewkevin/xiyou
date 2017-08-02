#ifndef __SKILL_TABLE_H__
#define __SKILL_TABLE_H__

struct SkillData{
	int32_t SkillId;
	int32_t Crit;
	int32_t Damage;
	int32_t Cooldown;
	std::vector<int32_t> BufferList;

	static bool Load(const char* fileName);
	static const SkillData* GetSkillById(int32_t id);

	static std::vector<SkillData> records_;
};



#endif