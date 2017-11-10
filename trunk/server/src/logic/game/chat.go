package game

import (
	"container/list"
	"jimny/logs"
	"logic/prpc"
)

type AudioInfo struct {
	AudioId int64
	Audio   []uint8
}

var (
	AudioGuid  int64 = 0
	AudioCatch []AudioInfo
	Record_MAX int  = 2000
	AudioList       = list.New()
	CK_World   int8 = 1
	CK_GM      int8 = 2
	CK_System  int8 = 0
	CK_Friend  int8 = 3
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

func BroadFriendChat(info prpc.COM_Chat) {

	friend := FindPlayerByInstId(info.PlayerInstId)

	if friend == nil {
		return
	}

	if friend.session == nil {
		return
	}

	friend.session.ReceiveChat(info)

}

//func (player *GamePlayer) RequestAudio(guid int64) {
//	dummy := []uint8{}
//	audio := FindAudioData(guid)
//	logs.Info("Player[", player.MyUnit.InstName, "]", "RequestAudio", audio)
//	if player.session == nil {
//		return
//	}
//	if audio == nil {
//		player.session.RequestAudioOk(-1, dummy)
//	} else {
//		player.session.RequestAudioOk(audio.AudioId, audio.Audio)
//	}
//
//}

func (player *GamePlayer) SendChat(info prpc.COM_Chat) {
	if info.Type == CK_GM {
		///啦啦啦啦啦啦
	} else {
		if info.Type == CK_System {

		} else if info.Type == CK_World {
			BroadcastChat(info)
		} else if info.Type == CK_Friend {
			BroadFriendChat(info)
		}
		logs.Info("Player[", player.MyUnit.InstName, "]", "SendChat", info)
	}
}

//func PushAudioInfo(audio []uint8) int64 {
//	af := AudioInfo{}
//	af.AudioId = atomic.AddInt64(&AudioGuid, 1)
//	af.Audio = audio
//
//	AudioList.PushBack(af)
//	if AudioList.Len() > Record_MAX {
//		AudioList.Remove(AudioList.Front())
//	}
//	return af.AudioId
//}

//func FindAudioData(guid int64) *AudioInfo {
//	af := AudioInfo{}
//	for e := AudioList.Front(); e != nil; e = e.Next() {
//		if e.Value.(AudioInfo).AudioId == guid {
//			af = e.Value.(AudioInfo)
//			return &af
//		}
//	}
//	return nil
//}

func (player *GamePlayer) TestChat() {
	//for i := 0; i < 2022; i++ {
	//	af := prpc.COM_Chat{}
	//	u8Array := []uint8{11, 22, 33, 44, 55, 66}
	//	af.Type = CK_World
	//	af.Audio = u8Array
	//	af.PlayerInstId = player.MyUnit.InstId
	//	af.PlayerName = player.MyUnit.InstName
	//	af.HeadIcon = "1111111"
	//	af.Content = "test test test test test[123123123@$%^^&&*]"
	//	player.SendChat(af)
	//}
	//player.RequestAudio(2022)
	//for e := AudioList.Front(); e != nil; e = e.Next() {
	//	af := AudioInfo{}
	//	af = e.Value.(AudioInfo)
	//	logs.Info("123123123123", af)
	//}
}
