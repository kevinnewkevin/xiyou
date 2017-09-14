package game

import (
	"logic/prpc"
	"strconv"
	"suzuki/conf"
	"strings"
	"fmt"
)

type (
	UnitRecord struct {
		Id     int32
		DispId int32
		Cost   int32
		IProp  []int32
		CProp  []float32
		Skills []int32
		chapters []int32
		BaseName string
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
		BuffId		int32
		Until		int32
		Type 		int32
		Kind 		int32
		Times 		int32
		AddLua 		string
		UpdateLua 	string
		PopLua 		string
	}
	BattleRecord struct {
		BattleId	int32
		MainId		int32
		SmallId		[]int32
	}
	PromoteInfo struct {
		Level		int32
		Hp			int32
		ATK			float32
		DEF			float32
		MATK		float32
		MDEF		float32
		AGILE		float32
		ItemId		int32
		ItemNum		int32
	}
)

var (
	UnitTable  = map[int32]*UnitRecord{}
	SkillTable = map[int32]*SkillRecord{}
	SkillLuaTable = map[int32]*SkillLuaRecord{}
	BuffTable = map[int32]*BuffRecord{}
	BattleTable = map[int32]*BattleRecord{}
	ExpTable = map[int32]int32{}
	PromoteTable = map[int32][]*PromoteInfo{}
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
		u.Cost = int32(csv.GetInt(r, "Cost"))
		u.IProp = make([]int32, prpc.IPT_MAX)
		u.CProp = make([]float32, prpc.CPT_MAX)
		u.BaseName = csv.GetString(r, "Name")
		u.IProp[prpc.IPT_HP] = csv.GetInt32(r, prpc.K_IPT_HP)
		u.IProp[prpc.IPT_PHYLE] = csv.GetInt32(r, prpc.K_IPT_PHYLE)
		u.IProp[prpc.IPT_LEVEL] = csv.GetInt32(r, prpc.K_IPT_LEVEL)
		u.IProp[prpc.IPT_COPPER] = csv.GetInt32(r, prpc.K_IPT_COPPER)
		u.IProp[prpc.IPT_SILVER] = csv.GetInt32(r, prpc.K_IPT_SILVER)
		u.IProp[prpc.IPT_GOLD] = csv.GetInt32(r, prpc.K_IPT_GOLD)
		u.IProp[prpc.IPT_PROMOTE] = csv.GetInt32(r, prpc.K_IPT_PROMOTE)

		u.CProp[prpc.CPT_HP] = csv.GetFloat32(r, prpc.K_CPT_HP)
		u.CProp[prpc.CPT_CHP] = csv.GetFloat32(r, prpc.K_CPT_HP)
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
		fmt.Println("1111")
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		b := BuffRecord{}
		b.BuffId = int32(csv.GetInt(r, "BuffId"))
		b.Until = int32(csv.GetInt(r, "Until"))
		b.Times = int32(csv.GetInt(r, "Times"))
		b.Type = int32(csv.GetInt(r, "Type"))
		b.Kind = int32(csv.GetInt(r, "Kind"))
		b.AddLua = csv.GetString(r, "AddScript")
		b.UpdateLua = csv.GetString(r, "UpdateScript")
		b.PopLua = csv.GetString(r, "PopScript")
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

func LoadPromoteTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {

		e_id := csv.GetInt32(r, "ID")

		p_info := PromoteInfo{}
		p_info.Level = csv.GetInt32(r, "Level")
		p_info.Hp = csv.GetInt32(r, "Hp")
		p_info.ATK = csv.GetFloat32(r, "ATK")
		p_info.DEF = csv.GetFloat32(r, "DEF")
		p_info.MATK = csv.GetFloat32(r, "MAGIC_ATK")
		p_info.MDEF = csv.GetFloat32(r, "MAGIC_DEF")
		p_info.AGILE = csv.GetFloat32(r, "AGILE")
		p_info.ItemId = csv.GetInt32(r, "ItemID")
		p_info.ItemNum = csv.GetInt32(r, "ItemNum")


		_, ok := PromoteTable[e_id]
		if !ok {
			PromoteTable[e_id] = []*PromoteInfo{}
		}

		PromoteTable[e_id] = append(PromoteTable[e_id], &p_info)


	}
	fmt.Println("PromoteTable", PromoteTable)
	return nil
}

func GetPromoteRecordById(id int32) []*PromoteInfo {
	return PromoteTable[id]
}
func LoadExpTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {

		Level := csv.GetInt32(r, "Lv")
		Exp := csv.GetInt32(r, "Exp")

		ExpTable[Level] = Exp
	}
	return nil
}

func GetExpRecordById(level int32) int32 {
	return ExpTable[level]
}


