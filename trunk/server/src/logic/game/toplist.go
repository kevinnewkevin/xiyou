package game

import (
	"logic/prpc"
	"strconv"

)

var TrueTopList []prpc.COM_TopUnit
var TMPTopList []prpc.COM_TopUnit

const (
	show_num = 5		//每页显示五个人
	Testpaiming = "测试用"
	num = 200
	tian = 10000
)


type TopList []prpc.COM_TopUnit

func (a TopList) Len() int {    // 重写 Len() 方法
	return len(a)
}
func (a TopList) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a TopList) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j].TianTi < a[i].TianTi
}

func InitTopList(){
	if len(TrueTopList) == 0{
		for i := 0; i < num; i++ {
			p := prpc.COM_TopUnit{}
			p.Name = Testpaiming + strconv.Itoa(i)
			if i / 2 == 0 {
				p.DisplayID = 1
			} else {
				p.DisplayID = 2
			}
			p.Level = num - int32(i)
			p.TianTi = tian - int32(i) * 2

			TrueTopList = append(TrueTopList, p)
		}
	}
	TMPTopList = TrueTopList

}


func (this *GamePlayer) AllTopByPage()  {

	this.session.RecvTopList(TrueTopList, this.TianTiRank)

	return
}


func (this *GamePlayer) FriendTopByPage(page int32) {


}


func (this *GamePlayer) FindMyTianTiRank() int32 {
	for i, t := range TrueTopList {
		if t.Name == this.MyUnit.InstName {
			return int32(i)
		}
	}

	return -1
}

func (this *GamePlayer) UpdateTianTiVal() {		//只更新不操作

	if this.TianTiRank == -1 {		// 无排名
		top := prpc.COM_TopUnit{}
		top.TianTi = this.TianTiVal
		top.Level = this.MyUnit.Level
		top.Name = this.MyUnit.InstName
		top.DisplayID = this.MyUnit.UnitId

		TMPTopList = append(TMPTopList, top)
	} else {
		my_top := TMPTopList[this.TianTiRank]

		my_top.TianTi = this.TianTiVal

		TMPTopList[this.TianTiRank] = my_top
	}

	//sort.Sort(TopList(TrueTopList))		// 重新排名

	//this.TianTiRank = this.FindMyTianTiRank()
}