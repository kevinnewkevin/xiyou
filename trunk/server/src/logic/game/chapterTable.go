package game

import (
	"suzuki/conf"
	"strconv"
	"strings"
)

type (
	Chapter struct {
		ChapterId			int32
		ChapterType			int32
		ChapterLevel   		int32
		ChapterStar			[]int32		//星级和奖励对应
		ChapterReward		[]int32
		SmallChapter		[]int32
	}

	SmallChapter struct {
		SmallChapterId		int32
		SmallChapterType	int32
		SmallChapterCase1	int32		//小章节星级达成条件  case1:击杀某个怪物   case2:多少回合内  case3:阵亡数量
		SmallChapterCase2	int32
		SmallChapterCase3	int32
		BattleID			int32
		DropID				int32
	}
)

var (
	ChapterTable  		= map[int32]*Chapter{}
	SmallChapterTable  	= map[int32]*SmallChapter{}
)

func LoadStoryChapterTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		c := Chapter{}

		c.ChapterId 		= int32(csv.GetInt(r,"ID"))
		c.ChapterLevel		= int32(csv.GetInt(r,"Type"))
		c.ChapterLevel		= int32(csv.GetInt(r,"Level"))

		strTmp := strings.Split(csv.GetString(r,"Star"),";")
		for i:=0;i<len(strTmp);i++{
			num,_ := strconv.Atoi(strTmp[i])
			c.ChapterStar = append(c.ChapterStar,int32(num))
		}

		strTmp1 := strings.Split(csv.GetString(r,"Reward"),";")
		for i:=0;i<len(strTmp1);i++{
			id,_ := strconv.Atoi(strTmp1[i])
			c.ChapterReward = append(c.ChapterReward,int32(id))
		}

		strTmp2 := strings.Split(csv.GetString(r,"CheckpointID"),";")
		for i:=0;i<len(strTmp2);i++{
			id,_ := strconv.Atoi(strTmp2[i])
			c.SmallChapter = append(c.SmallChapter,int32(id))
		}

		ChapterTable[c.ChapterId] = &c
	}
	return nil
}

func GetChapterById(id int32) *Chapter {
	return ChapterTable[id]
}

//--------------------------------------------------------------------------------------------------------------------------------------------

func LoadSmallChapterTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		c := SmallChapter{}
		c.SmallChapterId 		= int32(csv.GetInt(r,"ID"))
		c.SmallChapterType		= int32(csv.GetInt(r,"HeroID"))
		c.SmallChapterCase1		= int32(csv.GetInt(r,"Star1Need"))
		c.SmallChapterCase2		= int32(csv.GetInt(r,"Star2Need"))
		c.SmallChapterCase3		= int32(csv.GetInt(r,"Star3Need"))
		c.BattleID				= int32(csv.GetInt(r,"BattleID"))
		c.DropID				= int32(csv.GetInt(r,"DropID"))
		SmallChapterTable[c.SmallChapterId] = &c
	}
	return nil
}

func GetSmallChapterById(id int32) *SmallChapter{
	return SmallChapterTable[id]
}