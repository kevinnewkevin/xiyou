package game

import (
	"logic/conf"
)

type (
	TianTiData struct {
		Id         int32
		ScoreFloor int32
		ScoreUpper int32
		LadderDrop int32 //赛季奖励
		WinDrop    int32 //单次输赢奖励
		LoseDop    int32
	}
)

var (
	TianTiTableData = map[int32]*TianTiData{}
)

func LoadTianTiTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		c := TianTiData{}

		id := csv.GetInt32(r, "ID")
		if TianTiTableData[id] != nil {
			return nil
		}
		c.Id = id
		c.WinDrop = csv.GetInt32(r, "WinDrop")
		c.LoseDop = csv.GetInt32(r, "LoseDop")
		c.LadderDrop = csv.GetInt32(r, "LadderDrop")
		c.ScoreFloor = csv.GetInt32(r, "ScoreL")
		c.ScoreUpper = csv.GetInt32(r, "ScoreH")

		TianTiTableData[id] = &c
	}
	return nil
}

func GetTianTiIdByVal(val int32) int32 {
	var id int32 = 0
	for _, t := range TianTiTableData {
		if t.ScoreFloor <= val && t.ScoreUpper >= val {
			id = t.Id
		}
	}
	return id
}

func GetTianTiTableDataById(id int32) *TianTiData {
	return TianTiTableData[id]
}
