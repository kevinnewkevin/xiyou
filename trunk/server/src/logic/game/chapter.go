package game

import (
	"jimny/logs"
	"logic/prpc"
)

func OpenChapter(player *GamePlayer, cid int32) {
	if player == nil {
		return
	}
	if !IsOpenChapter(player, cid) {
		chapterData := prpc.COM_Chapter{}
		chapterData.ChapterId = cid

		chapter := GetChapterById(cid)
		if chapter == nil {
			return
		}

		for i := 0; i < len(chapter.SmallChapter); i++ {
			small := prpc.COM_SmallChapter{}
			small.SmallChapterId = chapter.SmallChapter[i]
			chapterData.SmallChapters = append(chapterData.SmallChapters, small)
		}
		player.Chapters = append(player.Chapters, &chapterData)
		if player.session == nil {
			return
		}

		player.session.OpenChapter(chapterData)
	}
}

func IsOpenChapter(player *GamePlayer, cid int32) bool {
	if player == nil {
		return false
	}
	for _, chapter := range player.Chapters {
		if cid == chapter.ChapterId {
			return true
		}
	}
	return false
}

func (player *GamePlayer) GetMyChapterDataById(chapterId int32) *prpc.COM_Chapter {
	if player == nil {
		return nil
	}
	for i := 0; i < len(player.Chapters); i++ {
		if player.Chapters[i] == nil {
			logs.Info("GetMyChapterDataById player.Chapters[i] == nil")
			continue
		}
		if player.Chapters[i].ChapterId == chapterId {
			return player.Chapters[i]
		}
	}

	return nil
}

func (player *GamePlayer) SycnMyChapterDataById(chapterId int32) {
	if player == nil {
		return
	}
	onceChapter := player.GetMyChapterDataById(chapterId)
	if onceChapter == nil {
		return
	}
	if onceChapter.ChapterId == 0 {
		return
	}
	if player.session == nil {
		return
	}

	for _, temp := range onceChapter.SmallChapters {
		logs.Info("SyncChapter ChapterId=", onceChapter.ChapterId, "SmallChapterId=", temp.SmallChapterId, "Star1=", temp.Star1, "Star2=", temp.Star2, "Star3=", temp.Star3)
	}

	player.session.SycnChapterData(*onceChapter)
}

func (player *GamePlayer) GetMySmallChapterDataById(smallid int32) *prpc.COM_SmallChapter {
	if player == nil {
		return nil
	}
	for _, chapter := range player.Chapters {
		for i := 0; i < len(chapter.SmallChapters); i++ {
			if chapter.SmallChapters[i].SmallChapterId == smallid {
				return &chapter.SmallChapters[i]
			}
		}
	}
	return nil
}

func (player *GamePlayer) AttackChapter(smallchapterid int32) {
	if player == nil {
		return
	}

	small := player.GetMySmallChapterDataById(smallchapterid)
	if small == nil {
		logs.Info("PlayerName", player.MyUnit.InstName, "AttackChapter ErrorNo=", "2")
		return
	}

	smallData := GetSmallChapterById(small.SmallChapterId)
	if smallData == nil {
		logs.Info("PlayerName", player.MyUnit.InstName, "AttackChapter ErrorNo=", "3")
		return
	}

	chapterData := GetChapterById(smallData.SmallChapterType)
	myUnitLevel := player.MyUnit.GetIProperty(prpc.IPT_LEVEL)
	if chapterData == nil {
		logs.Info("PlayerName", player.MyUnit.InstName, "AttackChapter ErrorNo=", "4")
		return
	}

	logs.Debug("AttackChapter smallchapterid=", smallchapterid, " smallData.SmallChapterType=", smallData.SmallChapterType, "chapterData.ChapterType", chapterData.ChapterType)

	myEnergy := player.MyUnit.GetIProperty(prpc.IPT_ENERGY)

	if smallData.EnergyExpend > myEnergy {
		logs.Info("PlayerName", player.MyUnit.InstName, "AttackChapter ErrorNo=", "1")
		return
	}

	if chapterData.ChapterType == 1 {
		for _, id := range smallData.UnLockId {
			//
			checksmall := player.GetMySmallChapterDataById(id)
			if checksmall == nil {
				logs.Debug("AttackChapter General checksmall==nil smallchapterid=", id)
				return
			}
			if !checksmall.Star1 {
				logs.Info("AttackChapter checksmall.Star1==false")
				return
			}
		}
	} else {
		if chapterData.ChapterLevel > myUnitLevel {
			logs.Info("PlayerName", player.MyUnit.InstName, "AttackChapter ErrorNo=", "6")
			return
		}
		for _, id := range smallData.UnLockId {
			//
			checksmall := player.GetMySmallChapterDataById(id)
			if checksmall == nil {
				logs.Debug("AttackChapter Difficulty checksmall==nil smallchapterid=", id)
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
		logs.Info("PlayerName", player.MyUnit.InstName, "AttackChapter ErrorNo=", "7")
		return
	}

	isok := CreatePvE(player, smallData.BattleID)
	if isok == nil {
		return
	}

	player.SetMyEnergy(smallData.EnergyExpend, false)

	player.ChapterID = smallchapterid
}

func (player *GamePlayer) CalcSmallChapterStar(battledata prpc.COM_BattleResult) int32 {
	if player == nil {
		return 0
	}
	if player.ChapterID == 0 {
		logs.Info("CalcSmallChapterStar player.ChapterID == 0")
		return 0
	}
	if battledata.Win == 0 {
		return 0
	}
	small := player.GetMySmallChapterDataById(player.ChapterID)
	if small == nil {
		logs.Info("CalcSmallChapterStar small == nil")
		return 0
	}
	smallData := GetSmallChapterById(player.ChapterID)
	if smallData == nil {
		logs.Info("CalcSmallChapterStar smallData == nil")
		return 0
	}

	if !small.Star1 {
		for i := 0; i < len(battledata.KillMonsters); i++ {
			logs.Info("battledata.KillMonsters monsterId=", battledata.KillMonsters[i], "Star1TargetId=", smallData.SmallChapterCase1)
			if battledata.KillMonsters[i] == smallData.SmallChapterCase1 {
				small.Star1 = true
				logs.Debug("SmallChapter=", small.SmallChapterId, "Star1 Succeed")
			}
		}
	}

	if !small.Star2 {
		logs.Info("CheckChapterStar3 TableCase=", smallData.SmallChapterCase2, "BattleVal=", battledata.BattleRound)
		if smallData.SmallChapterCase2 >= battledata.BattleRound {
			small.Star2 = true
			logs.Debug("SmallChapter=", small.SmallChapterId, "Star2 Succeed")
		}
	}

	if !small.Star3 {
		logs.Info("CheckChapterStar3 TableCase=", smallData.SmallChapterCase3, "BattleVal=", battledata.MySelfDeathNum)
		if smallData.SmallChapterCase3 >= battledata.MySelfDeathNum {
			small.Star3 = true
			logs.Debug("SmallChapter=", small.SmallChapterId, "Star3 Succeed")
		}
	}
	logs.Debug("CalcSmallChapterStar DropId = ", smallData.DropID)
	player.SycnMyChapterDataById(smallData.SmallChapterType)
	return smallData.DropID
}

func (player *GamePlayer) GetChapterStarReward(chapterId int32, star int32) {
	myChapter := player.GetMyChapterDataById(chapterId)
	if myChapter == nil {
		return
	}
	chapterData := GetChapterById(chapterId)
	if chapterData == nil {
		return
	}

	var myStar int32 = player.GetChapterStarById(chapterId)
	logs.Debug("chapter ", chapterId, " Star ", myStar, "StarReward", myChapter.StarReward)

	if myStar < star {
		logs.Debug("Lacking Star", chapterId, star)
		return
	}

	var index int = -1

	for i := 0; i < len(chapterData.ChapterStar); i++ {
		if chapterData.ChapterStar[i] == star {
			index = i
		}
	}
	if index == -1 {
		logs.Debug("chapter Without this reward star=", star)
		return
	}
	for _, s := range myChapter.StarReward {
		if s == star {
			logs.Info("chapter reward have been received star", star)
			return
		}
	}
	if index >= len(chapterData.ChapterReward) {
		return
	}
	player.GiveDrop(chapterData.ChapterReward[index])
	myChapter.StarReward = append(myChapter.StarReward, star)
	if player.session != nil {
		player.session.RequestChapterStarRewardOK()
	}
}

func (player *GamePlayer) GetChapterStarById(chapterId int32) int32 {
	myChapter := player.GetMyChapterDataById(chapterId)
	if myChapter == nil {
		return 0
	}
	var allStar int32 = 0

	for _, small := range myChapter.SmallChapters {
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
