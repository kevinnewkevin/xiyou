package game

import (
	"jimny/logs"
	"logic/prpc"
	"strconv"
)

////////////////////////////////////////////////////
////好友
////////////////////////////////////////////////////

func (this *GamePlayer) SerchFriendByName(name string) {
	t := FindPlayerByInstName(name)

	if t == nil {
		logs.Debug("不存在")
		return
	}

	if t.session == nil {
		logs.Debug("不在线")
		return
	}

	for _, b := range t.Enemys {
		if b.Name == name {
			logs.Debug("你在对方黑名单中,不能查看")
			return
		}
	}

	info := prpc.COM_Friend{}
	info.InstId = t.MyUnit.InstId
	info.Name = t.MyUnit.InstName
	info.Level = t.MyUnit.Level
	info.DisplayID = t.MyUnit.UnitId

	t.session.SerchFriendInfo(info)		//向对方发送好友信息

}

func (this *GamePlayer) SerchFriendRandom() {
	var step int32 = 10
	all := []prpc.COM_Friend{}

	low := this.MyUnit.Level - step
	if low <= 1 {
		low = 1
	}

	high := this.MyUnit.Level - step

	for _, p := range PlayerStore {
		if low <= p.MyUnit.Level && p.MyUnit.Level <= high {
			info := prpc.COM_Friend{}
			info.InstId = p.MyUnit.InstId
			info.Name = p.MyUnit.InstName
			info.Level = p.MyUnit.Level
			info.DisplayID = p.MyUnit.UnitId

			all = append(all, info)
		}
	}

	if len(all) == 0 {
		for i := 1; i <= 6; i++ {
			info := prpc.COM_Friend{}
			info.InstId = int64(i)
			info.Name = strconv.Itoa(i)
			info.Level = this.MyUnit.Level
			info.DisplayID = this.MyUnit.UnitId

			all = append(all, info)
		}
	}

	if len(all) > 6 {
		this.session.FriendInfo(all[:6])
	} else {
		this.session.FriendInfo(all)
	}

}

func (this *GamePlayer) ApplicationFriend(name string) {
	t := FindPlayerByInstName(name)

	if t == nil {
		logs.Debug("不存在")
		return
	}

	if t.session == nil {
		logs.Debug("不在线")
		return
	}

	for _, b := range t.Enemys {
		if b.Name == name {
			logs.Debug("你在对方黑名单中,不能添加好友")
			return
		}
	}

	t.session.ApplyFriend(this.MyUnit.InstName) //向对方发送好友信息
}

func (this *GamePlayer) ProcessingFriend(name string) {
	t := FindPlayerByInstName(name)

	if t == nil {
		logs.Debug("不存在")
		return
	}

	if t.session == nil {
		logs.Debug("不在线")
		return
	}

	for _, b := range t.Enemys {
		if b.Name == name {
			logs.Debug("你在对方黑名单中,不能添加好友")
			return
		}
	}

	//把自己加入到对方好友名单
	info := prpc.COM_Friend{}
	info.InstId = this.MyUnit.InstId
	info.Name = this.MyUnit.InstName
	info.Level = this.MyUnit.Level
	info.DisplayID = this.MyUnit.UnitId
	t.Friends = append(t.Friends, &info)
	t.session.RecvFriend(info)

	//把别人加入到自己的好友名单
	t_info := prpc.COM_Friend{}
	info.InstId = t.MyUnit.InstId
	t_info.Name = t.MyUnit.InstName
	t_info.Level = t.MyUnit.Level
	t_info.DisplayID = t.MyUnit.UnitId
	this.Friends = append(this.Friends, &t_info)
	this.session.RecvFriend(t_info)

}

func (this *GamePlayer) DeleteFriend(instid int64) {

	this.delFriend(instid)

	friend := FindPlayerByInstId(instid)

	if friend == nil {
		return
	}

	friend.delFriend(this.MyUnit.InstId)
}

func (this *GamePlayer) delFriend(instid int64) {
	idx := -1
	for i, f := range this.Friends {
		if f.InstId == instid {
			idx = i
		}
	}
	this.Friends = append(this.Friends[:idx], this.Friends[idx+1:]...)

	if this.session == nil {
		return
	}

	//给player发送删除好友的信息
	this.session.DelFriend(instid)

}

func (this *GamePlayer) CheckMe() {
	for _, f := range this.Friends {
		friend := FindPlayerByInstId(f.InstId)
		if friend == nil {
			continue
		}
		friend.delFriend(this.MyUnit.InstId)
	}

}

////////////////////////////////////////////////////
////黑名单
////////////////////////////////////////////////////

func (this *GamePlayer) AddBlackList() {

}

func (this *GamePlayer) DeleteBlackList() {

}

////////////////////////////////////////////////////
////信息
////////////////////////////////////////////////////

func (this *GamePlayer) SendMessage() {

}

////////////////////////////////////////////////////
///测试
////////////////////////////////////////////////////

func (this *GamePlayer) InitTestFriend() {
	var uid int64 = 1000000000
	for i := 1; i <= 10; i++ {
		info := prpc.COM_Friend{}
		info.InstId = uid - int64(i)
		info.Name = "测试Friend" + strconv.Itoa(i)
		info.Level = int32(i) + 1
		info.DisplayID = this.MyUnit.UnitId

		this.Friends = append(this.Friends, &info)
	}

}
