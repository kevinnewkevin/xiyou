package game

import (
	"logic/prpc"
	"sort"
	"jimny/logs"
)

var TrueTopList []prpc.COM_TopUnit
var TMPTopList []prpc.COM_TopUnit

const (
	show_num    = 5 //每页显示五个人
	Testpaiming = "测试用"
	num = 200
	tian        = 10000
)

type TopList []prpc.COM_TopUnit

func (a TopList) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a TopList) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a TopList) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].TianTi < a[i].TianTi
}

func InitTopList() {
	TMPTopList = <- QueryAllTopList()

	sort.Sort(TopList(TMPTopList))

	if len(TMPTopList) > num {
		TrueTopList = TMPTopList[:num]
	} else {
		TrueTopList = TMPTopList
	}

}

func isSame(t1 []prpc.COM_TopUnit, t2 []prpc.COM_TopUnit) bool {

	for idx, t := range t1 {
		if t == t2[idx] {
			continue
		}
		return false
	}

	return true
}

//////////////////////////////////////////////////////////////////////////////////////
/////数据
//////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) SetTopSGE(p prpc.SGE_DBPlayer) {

	this.SetPlayerCOM(&p.COM_Player)
	this.PlayerId = p.PlayerId
	this.Username = p.Username
	this.LoginTime = p.LoginTime
	this.LogoutTime = p.LogoutTime
	for i := range p.BagItemList {
		this.BagItems = append(this.BagItems, &p.BagItemList[i])
	}

	this.BlackMarketData = &p.BlackMarketData
}

func (this *GamePlayer) GetTopSGE() prpc.SGE_DBTopUnit {

	data := prpc.SGE_DBTopUnit{}
	data.Name = this.MyUnit.InstName
	data.DisplayID = this.MyUnit.UnitId
	data.Level = this.MyUnit.Level
	data.TianTi = this.TianTiVal

	return data
}

//////////////////////////////////////////////////////////////////////////////////////
/////全体排行
//////////////////////////////////////////////////////////////////////////////////////

func RefreshAllTopList() {
	sort.Sort(TopList(TMPTopList)) // 重新排名

	tmp := TrueTopList


	if len(TMPTopList) > num {
		TrueTopList = TMPTopList[:num]
	}

	if isSame(tmp, TrueTopList) {
		return
	}
	for _, p := range PlayerStore {
		if p == nil || p.session == nil {
			continue
		}
		p.TianTiRank = p.FindMyTianTiRank()
		p.session.RecvTopList(TrueTopList, p.TianTiRank)
	}
}

func (this *GamePlayer) AllTopByPage() {
	//logs.Debug("AllTopByPage", TrueTopList)
	this.TianTiRank = this.FindMyTianTiRank()
	logs.Debug("AllTopByPage 1 ", this.TianTiRank, len(TrueTopList))
	this.session.RecvTopList(TrueTopList, this.TianTiRank)

	return
}

func (this *GamePlayer) FindMyTianTiRank() int32 {
	for i, t := range TrueTopList {
		if t.Name == this.MyUnit.InstName {
			return int32(i)
		}
	}

	return -1
}

func (this *GamePlayer) UpdateTianTiVal() { //只更新不操作

	if this.TianTiRank == -1 { // 无排名

		idx := this.FindMyTianTiRank()

		//logs.Debug("UpdateTianTiVal -100, ", idx)
		if idx == -1 {
			top := prpc.COM_TopUnit{}
			top.TianTi = this.TianTiVal
			top.Level = this.MyUnit.Level
			top.Name = this.MyUnit.InstName
			top.DisplayID = this.MyUnit.UnitId

			TMPTopList = append(TMPTopList, top)
		} else {
			TMPTopList[idx].TianTi = this.TianTiVal
		}
		//logs.Debug("UpdateTianTiVal -1, ", TMPTopList)
	} else {
		my_top := TMPTopList[this.TianTiRank]

		my_top.TianTi = this.TianTiVal

		TMPTopList[this.TianTiRank] = my_top
	}

	UpdateTopList(this.MyUnit.InstId, this.GetTopSGE())
}

//////////////////////////////////////////////////////////////////////////////////////
/////好友排行
//////////////////////////////////////////////////////////////////////////////////////

func (this *GamePlayer) FindMyFriendTianTiRank() int32 {
	for i, t := range this.FriendTop {
		if t.Name == this.MyUnit.InstName {
			return int32(i)
		}
	}

	return -1
}

func (this *GamePlayer) FriendTopByPage() {
	//initMeToFriendTopList(this) //测试用
	//logs.Debug("FriendTopByPage", TrueFriendTopList)

	this.FriendTianTiRank = this.FindMyFriendTianTiRank()
	logs.Debug("FriendTopByPage 1 ", this.FriendTianTiRank, len(this.FriendTop), "myname is :", this.MyUnit.InstName)
	this.session.RecvFriendTopList(this.FriendTop, this.FriendTianTiRank)
}

func RefreshFriendTopList() {

	//if isSame(tmp, TrueFriendTopList) {
	//	return
	//}

	for _, p := range PlayerStore {
		if p == nil || p.session == nil {
			continue
		}

		p.RefreshFriendTopList()

		p.FriendTianTiRank = p.FindMyTianTiRank()
		p.session.RecvFriendTopList(p.FriendTop, p.FriendTianTiRank)
	}
}

func (this *GamePlayer) FindFriendTianTiRank() int32 {
	for i, t := range this.FriendTop {
		if t.Name == this.MyUnit.InstName {
			return int32(i)
		}
	}

	return -1
}

func (this *GamePlayer) RefreshFriendTopList()  {		//好友之间的天梯排行榜
	this.FriendTop = []prpc.COM_TopUnit{}

	mt := prpc.COM_TopUnit{}
	mt.Name = this.MyUnit.InstName
	mt.Level = this.MyUnit.Level
	mt.TianTi = this.TianTiVal
	mt.DisplayID = this.MyUnit.UnitId

	this.FriendTop = append(this.FriendTop, mt)

	for _, f := range this.Friends {
		fr := FindPlayerByInstId(f.InstId)
		if fr == nil {
			var fd *prpc.SGE_DBPlayer
			if fd = <- QueryPlayer(f.Username); fd!=nil {
				p := &GamePlayer{}
				p.SetPlayerSGE(*fd)

				if p.TianTiVal == 0 {
					continue
				}

				t := prpc.COM_TopUnit{}
				t.Name = p.MyUnit.InstName
				t.Level = p.MyUnit.Level
				t.TianTi = p.TianTiVal
				t.DisplayID = p.MyUnit.UnitId

				this.FriendTop = append(this.FriendTop, t)
			}
		} else {
			if fr.TianTiVal == 0 {
				continue
			}

			t := prpc.COM_TopUnit{}
			t.Name = fr.MyUnit.InstName
			t.Level = fr.MyUnit.Level
			t.TianTi = fr.TianTiVal
			t.DisplayID = fr.MyUnit.UnitId

			this.FriendTop = append(this.FriendTop, t)
		}
	}

	sort.Sort(TopList(this.FriendTop)) // 重新排名

	return
}
