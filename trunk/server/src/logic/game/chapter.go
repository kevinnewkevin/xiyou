package game

import (
	"logic/prpc"
	"fmt"
)

func OpenChapter(player *GamePlayer,cid int32)  {
	if player==nil {
		return
	}
	if !IsOpenChapter(player,cid) {
		chapterData := prpc.COM_Chapter{}
		chapterData.ChapterId = cid

		chapter := GetChapterById(cid)
		if chapter == nil {
			return
		}

		for i:=0; i<len(chapter.SmallChapter);i++ {
			small := prpc.COM_SmallChapter{}
			small.SmallChapterId = chapter.SmallChapter[i]
			chapterData.SmallChapters = append(chapterData.SmallChapters,small)
		}
		player.Chapters = append(player.Chapters,&chapterData)
	}
}

func IsOpenChapter(player *GamePlayer, cid int32) bool {
	if player == nil {
		return false
	}
	for _, chapter := range player.Chapters {
		if cid== chapter.ChapterId{
			return true
		}
	}
	return false
}

func (player *GamePlayer)GetMyChapterDataById(chapterId int32) *prpc.COM_Chapter {
	if player==nil {
		return nil
	}
	onceChapter := prpc.COM_Chapter{}
	for _, chapter := range player.Chapters {
		if chapter.ChapterId == chapterId {
			onceChapter.ChapterId = chapter.ChapterId
			for _,smallchapter := range chapter.SmallChapters{
				onceChapter.SmallChapters = append(onceChapter.SmallChapters,smallchapter)
			}
		}
	}
	if onceChapter.ChapterId == 0 {
		return nil
	}
	return &onceChapter
}

func (player *GamePlayer)SycnMyChapterDataById(chapterId int32)  {
	if player==nil {
		return
	}
	onceChapter := player.GetMyChapterDataById(chapterId)
	if onceChapter == nil {
		return
	}
	if onceChapter.ChapterId == 0 {
		return
	}
	if player.session==nil {
		return
	}
	player.session.SycnChapterData(*onceChapter)
}

func (player *GamePlayer)GetMySmallChapterDataById(smallid int32) *prpc.COM_SmallChapter {
	if player==nil {
		return nil
	}
	for _, chapter := range player.Chapters {
		for _,smallchapter := range chapter.SmallChapters{
			if smallchapter.SmallChapterId == smallid {
				return &smallchapter
			}
		}
	}
	return nil
}

func (player *GamePlayer)AttackChapter(smallchapterid int32)  {
	if player==nil {
		return
	}

	small := player.GetMySmallChapterDataById(smallchapterid)
	if small == nil {
		return
	}
	smallData := GetSmallChapterById(small.SmallChapterId)
	if smallData == nil {
		return
	}
	chapterData := GetChapterById(smallData.SmallChapterType)
	myUnitLevel := player.MyUnit.GetIProperty(prpc.IPT_LEVEL)
	if chapterData == nil {
		return
	}

	myEnergy := player.MyUnit.GetIProperty(prpc.IPT_ENERGY)

	if smallData.EnergyExpend > myEnergy {
		return
	}

	if chapterData.ChapterType == 1 {
		for _,id :=  range smallData.UnLockId{
			//
			checksmall := player.GetMySmallChapterDataById(id)
			if checksmall == nil {
				fmt.Println("AttackChapter General checksmall==nil smallchapterid=",id)
				return
			}
			if !checksmall.Star1 {
				return
			}
		}
	}else {
		if chapterData.ChapterLevel > myUnitLevel {
			return
		}
		for _,id :=  range smallData.UnLockId{
			//
			checksmall := player.GetMySmallChapterDataById(id)
			if checksmall == nil {
				fmt.Println("AttackChapter Difficulty checksmall==nil smallchapterid=",id)
				return
			}
			if !checksmall.Star1 {
				return
			}
			if !checksmall.Star2 {
				return
			}
			if !checksmall.Star3 {
				return
			}
		}
	}

	if player.IsBattle() {
		return
	}

	isok := CreatePvE(player,smallData.BattleID)
	if isok==nil {
		return
	}
	myEnergy -= smallData.EnergyExpend
	player.MyUnit.SetIProperty(prpc.IPT_ENERGY,myEnergy)
	player.ChapterID = smallchapterid
}

func (player *GamePlayer)CalcSmallChapterStar(battledata prpc.COM_BattleResult) int32 {
	if player == nil {
		return 0
	}
	if player.ChapterID == 0 {
		fmt.Println("CalcSmallChapterStar player.ChapterID == 0")
		return 0
	}
	if battledata.Win==0 {
		return 0
	}
	small := player.GetMySmallChapterDataById(player.ChapterID)
	if small == nil {
		fmt.Println("CalcSmallChapterStar small == nil")
		return 0
	}
	smallData := GetSmallChapterById(player.ChapterID)
	if smallData == nil {
		fmt.Println("CalcSmallChapterStar smallData == nil")
		return 0
	}

	if !small.Star1 {
		for i:=0;i<len(battledata.KillMonsters) ;i++  {
			if battledata.KillMonsters[i] == smallData.SmallChapterCase1 {
				small.Star1 = true
			}
		}
	}

	if !small.Star2 {
		if smallData.SmallChapterCase2 > battledata.BattleRound {
			small.Star2 = true
		}
	}

	if !small.Star3 {
		if smallData.SmallChapterCase3 > battledata.MySelfDeathNum {
			small.Star3 = true
		}
	}
	fmt.Println("CalcSmallChapterStar DropId = ",smallData.DropID)
	player.SycnMyChapterDataById(smallData.SmallChapterType)
	return smallData.DropID
}

func (player *GamePlayer)GetChapterStarReward(chapterId int32,star int32)  {
	myChapter := player.GetMyChapterDataById(chapterId)
	if myChapter==nil {
		return
	}
	chapterData := GetChapterById(chapterId)
	if chapterData == nil {
		return
	}

	var myStar int32 = player.GetChapterStarById(chapterId)
	fmt.Println("chapter ",chapterId," Star ",myStar)

	if myStar < star {
		fmt.Println("Lacking Star",chapterId,star)
		return
	}

	var index int = -1

	for i:=0;i<len(chapterData.ChapterStar) ;i++  {
		if chapterData.ChapterStar[i] == star {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("chapter Without this reward star=",star)
		return
	}
	for _,s := range myChapter.StarReward{
		if s == star {
			fmt.Println("chapter reward have been received star",star)
			return
		}
	}
	if index >= len(chapterData.ChapterReward) {
		return
	}
	player.GiveDrop(chapterData.ChapterReward[index])
	myChapter.StarReward = append(myChapter.StarReward,star)
	if player.session != nil {
		player.session.RequestChapterStarRewardOK()
	}
}

func (player *GamePlayer)GetChapterStarById(chapterId int32) int32 {
	myChapter := player.GetMyChapterDataById(chapterId)
	if myChapter==nil {
		return 0
	}
	var allStar int32 = 0

	for _,small := range myChapter.SmallChapters{
		if small.Star1 {
			allStar += 1
		}
		if small.Star2 {
			allStar += 1
		}
		if small.Star3 {
			allStar += 1
		}
	}

	return allStar
}