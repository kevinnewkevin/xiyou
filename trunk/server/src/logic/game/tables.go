package game

import (
	"logic/prpc"
	"suzuki/conf"
	"strconv"
)

type (
	UnitRecord struct {
		Id     int32
		DispId int32
		IProp  []int32
		CProp  []float32
		Skills []int32
	}
	SkillRecord struct {
		SkillID int32
		Crit int32
		Damage int32
		BuffList []int32
		CoolDown int32
		TargetNum int
		TargetCamp int
	}
)

var (
	unitTable =  map[int32]*UnitRecord{}
	skillTable = map[int32]*SkillRecord{}
)

func LoadUnitTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		u := UnitRecord{}
		u.Id = int32(csv.GetInt(r, "EntityId"))
		u.DispId = int32(csv.GetInt(r, "DisplayId"))
		u.IProp = make([]int32,prpc.IPT_MAX)
		u.CProp = make([]float32,prpc.CPT_MAX)
		u.IProp[prpc.IPT_PHYLE] = csv.GetInt32(r, prpc.K_IPT_PHYLE)
		u.IProp[prpc.IPT_LEVEL] = csv.GetInt32(r, prpc.K_IPT_LEVEL)
		u.IProp[prpc.IPT_COPPER] = csv.GetInt32(r, prpc.K_IPT_COPPER)
		u.IProp[prpc.IPT_SILVER] = csv.GetInt32(r, prpc.K_IPT_SILVER)
		u.IProp[prpc.IPT_GOLD] = csv.GetInt32(r, prpc.K_IPT_GOLD)

		u.CProp[prpc.CPT_HP] = csv.GetFloat32(r, prpc.K_CPT_HP)
		u.CProp[prpc.CPT_ATK] = csv.GetFloat32(r, prpc.K_CPT_ATK)
		u.CProp[prpc.CPT_DEF] = csv.GetFloat32(r, prpc.K_CPT_DEF)
		u.CProp[prpc.CPT_MAGIC_ATK] = csv.GetFloat32(r, prpc.K_CPT_MAGIC_ATK)
		u.CProp[prpc.CPT_MAGIC_DEF] = csv.GetFloat32(r, prpc.K_CPT_MAGIC_DEF)
		u.CProp[prpc.CPT_AGILE] = csv.GetFloat32(r, prpc.K_CPT_AGILE)
		u.CProp[prpc.CPT_KILL] = csv.GetFloat32(r, prpc.K_CPT_KILL)
		u.CProp[prpc.CPT_CRIT] = csv.GetFloat32(r, prpc.K_CPT_CRIT)
		u.CProp[prpc.CPT_COUNTER_ATTACK] = csv.GetFloat32(r, prpc.K_CPT_COUNTER_ATTACK)
		u.CProp[prpc.CPT_SPUTTERING] = csv.GetFloat32(r, prpc.K_CPT_SPUTTERING)
		u.CProp[prpc.CPT_DOUBLE_HIT] = csv.GetFloat32(r, prpc.K_CPT_DOUBLE_HIT)
		u.CProp[prpc.CPT_RECOVERY] = csv.GetFloat32(r, prpc.K_CPT_RECOVERY)
		u.CProp[prpc.CPT_REFLEX] = csv.GetFloat32(r, prpc.K_CPT_REFLEX)
		u.CProp[prpc.CPT_SUCK_BLOOD] = csv.GetFloat32(r, prpc.K_CPT_SUCK_BLOOD)
		u.CProp[prpc.CPT_INCANTER] = csv.GetFloat32(r, prpc.K_CPT_INCANTER)
		u.CProp[prpc.CPT_RESISTANCE] = csv.GetFloat32(r, prpc.K_CPT_RESISTANCE)

		u.Skills = make([]int32,4)
		u.Skills[0] = csv.GetInt32(r, "Skill1")
		u.Skills[1] = csv.GetInt32(r, "Skill2")
		u.Skills[2] = csv.GetInt32(r, "Skill3")
		u.Skills[3] = csv.GetInt32(r, "Skill4")

		unitTable[u.Id] = &u
	}
	return nil
}

func GetUnitRecordById(id int32) *UnitRecord {
	return unitTable[id]
}


func LoadSkillTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		s := SkillRecord{}
		s.SkillID = int32(csv.GetInt(r, "SkilId"))
		s.Crit = int32(csv.GetInt(r, "Crit"))
		s.Damage = int32(csv.GetInt(r, "Damage"))
		string1 := csv.GetStrings(r, "Bufflist")
		for i := 0; i < len(string1); i++ {
			buffid, _ := strconv.Atoi(string1[i])
			s.BuffList = append(s.BuffList, int32(buffid))
		}
		s.CoolDown = int32(csv.GetInt(r, "CoolDown"))
		s.TargetNum = csv.GetInt(r, "TargetNum")
		s.TargetCamp = csv.GetInt(r, "TargetCamp")
		skillTable[s.SkillID] = &s
	}
	return nil
}

func GetSkillRecordById(id int32) *SkillRecord {
	return skillTable[id]
}