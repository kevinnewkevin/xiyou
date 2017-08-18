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

func (player *GamePlayer)GetMyChapterDataById(chapterId int32)  {
	if player==nil {
		return
	}
	onceChapter := prpc.COM_Chapter{}
	for _, chapter := range player.Chapters {
		if chapter.ChapterId == chapterId {
			onceChapter.ChapterId = chapter.ChapterId
			for _,smallchapter := range chapter.SmallChapters{
				onceChapter.SmallChapters = append(onceChapter.SmallChapters,smallchapter)
				fmt.Println(onceChapter)
			}
		}
	}
	if onceChapter.ChapterId == 0 {
		return
	}
	if player.session==nil {
		return
	}
	player.session.SycnChapterData(onceChapter)
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
	if chapterData.ChapterLevel > myUnitLevel {
		return
	}
	if player.IsBattle() {
		return
	}

	isok := CreatePvE(player,smallData.BattleID)
	if isok==nil {
		return
	}
	player.ChapterID = smallchapterid
}

func (player *GamePlayer)CalcSmallChapterStar(battledata prpc.COM_BattleResult)  {
	if player == nil {
		return
	}
	if player.ChapterID == 0 {
		return
	}
	if battledata.Win==0 {
		return
	}
	small := player.GetMySmallChapterDataById(player.ChapterID)
	if small == nil {
		return
	}
	smallData := GetSmallChapterById(player.ChapterID)
	if smallData == nil {
		return
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

	//player.GiveDrop(smallData.DropID)
	player.GetMyChapterDataById(smallData.SmallChapterType)
}