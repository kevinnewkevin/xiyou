package game

import (
	"jimny/logs"
	"logic/prpc"
)

type AudioInfo struct {
	AudioId int64
	Audio   []uint8
}

const (
	CK_System  		int8 = 0
	CK_World   		int8 = 1
	CK_GM      		int8 = 2
	CK_Friend  		int8 = 3
	CK_Guild   		int8 = 4
)

func BroadcastChat(info prpc.COM_Chat) {
	for i := 0; i < len(PlayerStore); i++ {
		if PlayerStore[i] == nil {
			continue
		}

		if PlayerStore[i].session == nil {
			continue
		}
		PlayerStore[i].session.ReceiveChat(info)
	}
}

func BroadFriendChat(info prpc.COM_Chat, sendid int64) {

	friend := FindPlayerByInstId(info.PlayerInstId)

	if friend == nil {
		return
	}

	if friend.session == nil {
		return
	}

	info.PlayerInstId = sendid
	logs.Info("BroadFriendChat ", info)
	friend.session.ReceiveChat(info)

}


func (player *GamePlayer) SendChat(info prpc.COM_Chat) {
	if info.Type == CK_GM {
		///啦啦啦啦啦啦
	} else {
		if info.Type == CK_System {

		} else if info.Type == CK_World {
			BroadcastChat(info)
		} else if info.Type == CK_Friend {
			BroadFriendChat(info, player.MyUnit.InstId)
		} else if info.Type == CK_Guild{
			if player.GuildId == 0 {
				return
			}
			pGuild := FindGuildById(player.GuildId)
			if pGuild==nil {
				return
			}
			pGuild.GuildChat(info)
		}
		logs.Info("Player[", player.MyUnit.InstName, "]", "SendChat", info)
	}
}

func (player *GamePlayer) TestChat() {

}
