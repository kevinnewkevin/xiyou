package game

import (
	"logic/prpc"
	"strconv"
	"sort"
)

var TestTop []prpc.COM_TopUnit

const (
	show_num = 5		//每页显示五个人
	Testpaiming = "测试用"
	num = 1000
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
	if len(TestTop) == 0{
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

			TestTop = append(TestTop, p)
		}
	}

}


func (this *GamePlayer) AllTopByPage(page int32)  {
	start := show_num * page - show_num
	end := show_num * page - 1

	show_p := []prpc.COM_TopUnit{}
	for i := start; i <= end; i++ {
		show_p = append(show_p, TestTop[i])
	}

	this.UpdateTianTiVal()

	this.session.RecvTopList(show_p, this.TianTiRank)

	return
}


func (this *GamePlayer) FriendTopByPage(page int32) {


}


func (this *GamePlayer) FindMyTianTiRank() int32 {
	for i, t := range TestTop {
		if t.Name == this.MyUnit.InstName {
			return int32(i)
		}
	}

	return -1
}

func (this *GamePlayer) UpdateTianTiVal() {

	if this.TianTiRank == -1 {		// 无排名
		top := prpc.COM_TopUnit{}
		top.TianTi = this.TianTiVal
		top.Level = this.MyUnit.Level
		top.Name = this.MyUnit.InstName
		top.DisplayID = this.MyUnit.UnitId

		TestTop = append(TestTop, top)
	} else {
		my_top := TestTop[this.TianTiRank]

		my_top.TianTi = this.TianTiVal

		TestTop[this.TianTiRank] = my_top
	}

	sort.Sort(TopList(TestTop))		// 重新排名

	if len(TestTop) > 1000 {
		TestTop = TestTop[:1000]
	}

	this.TianTiRank = this.FindMyTianTiRank()
}