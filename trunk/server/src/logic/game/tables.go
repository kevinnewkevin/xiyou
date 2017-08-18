package game

import (
	"logic/prpc"
	"strconv"
	"suzuki/conf"
	"strings"
)

type (
	UnitRecord struct {
		Id     int32
		DispId int32
		IProp  []int32
		CProp  []float32
		Skills []int32
		chapters []int32
	}
	SkillRecord struct {
		SkillID    int32
		Crit       int32
		Damage     int32
		BuffList   []int32
		CoolDown   int32
		TargetNum  int
		TargetCamp int
		LuaScprit  string
	}
	SkillLuaRecord struct {
		SkillID    int32
		LuaScprit  string
	}
	BuffRecord struct {
		BuffId	int32
		Until	int32
		Type 	int32
		Kind 	int32
	}
	BattleRecord struct {
		BattleId	int32
		MainId		int32
		SmallId		[]int32
	}
)

var (
	UnitTable  = map[int32]*UnitRecord{}
	SkillTable = map[int32]*SkillRecord{}
	SkillLuaTable = map[int32]*SkillLuaRecord{}
	BuffTable = map[int32]*BuffRecord{}
	BattleTable = map[int32]*BattleRecord{}
)

func LoadUnitTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		u := UnitRecord{}
		u.Id = int32(csv.GetInt(r, "UnitId"))
		u.DispId = int32(csv.GetInt(r, "DisplayId"))
		u.IProp = make([]int32, prpc.IPT_MAX)
		u.CProp = make([]float32, prpc.CPT_MAX)
		u.IProp[prpc.IPT_HP] = csv.GetInt32(r, prpc.K_IPT_HP)
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

		u.Skills = make([]int32, 4)
		u.Skills[0] = csv.GetInt32(r, "Skill1")
		u.Skills[1] = csv.GetInt32(r, "Skill2")
		u.Skills[2] = csv.GetInt32(r, "Skill3")
		u.Skills[3] = csv.GetInt32(r, "Skill4")

		strTmp := strings.Split(csv.GetString(r,"ChapterID"),";")
		for i:=0;i<len(strTmp);i++{
			id,_ := strconv.Atoi(strTmp[i])
			u.chapters = append(u.chapters,int32(id))
		}

		UnitTable[u.Id] = &u
	}
	return nil
}

func GetUnitRecordById(id int32) *UnitRecord {
	return UnitTable[id]
}

func GetUnitChapterById(id int32) []int32 {
	if UnitTable[id] == nil {
		return nil
	}
	return UnitTable[id].chapters
}

func LoadSkillTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		s := SkillRecord{}
		s.SkillID = int32(csv.GetInt(r, "SkillId"))
		s.Crit = int32(csv.GetInt(r, "Crit"))
		s.Damage = int32(csv.GetInt(r, "Damage"))
		string1 := csv.GetStrings(r, "Bufflist")
		for i := 0; i < len(string1); i++ {
			buffid, _ := strconv.Atoi(string1[i])
			s.BuffList = append(s.BuffList, int32(buffid))
		}
		s.TargetNum = csv.GetInt(r, "TargetNum")
		s.TargetCamp = csv.GetInt(r, "TargetCamp")
		s.LuaScprit = csv.GetString(r, "LuaName")

		SkillTable[s.SkillID] = &s
	}
	return nil
}

func GetSkillRecordById(id int32) *SkillRecord {
	return SkillTable[id]
}

func LoadSkillLuaTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		s := SkillLuaRecord{}
		s.SkillID = int32(csv.GetInt(r, "SkillId"))
		s.LuaScprit = csv.GetString(r, "TargetCamp")
		SkillLuaTable[s.SkillID] = &s
	}
	return nil
}

func GetSkillLuaRecordById(id int32) *SkillLuaRecord {
	return SkillLuaTable[id]
}

func LoadBuffTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		b := BuffRecord{}
		b.BuffId = int32(csv.GetInt(r, "BuffId"))
		b.Until = int32(csv.GetInt(r, "Until"))
		b.Type = int32(csv.GetInt(r, "Type"))
		b.Kind = int32(csv.GetInt(r, "Kind"))
		BuffTable[b.BuffId] = &b
	}
	return nil
}

func GetBuffRecordById(id int32) *BuffRecord {
	return BuffTable[id]
}

func LoadBattleTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		b := BattleRecord{}
		b.BattleId = int32(csv.GetInt(r, "ID"))
		b.MainId = int32(csv.GetInt(r, "MainID"))
		string1 := csv.GetStrings(r, "SmallID")
		for i := 0; i < len(string1); i++ {
			buffid, _ := strconv.Atoi(string1[i])
			b.SmallId = append(b.SmallId, int32(buffid))
		}
		BattleTable[b.BattleId] = &b
	}
	return nil
}

func GetBattleRecordById(id int32) *BattleRecord {
	return BattleTable[id]
}
