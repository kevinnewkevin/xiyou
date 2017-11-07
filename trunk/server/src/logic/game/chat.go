package game

import (
	"logic/prpc"
	"sync/atomic"
)

type AudioInfo struct {
 	AudioId 	int64
	Audio		[]uint8
}

var (
	AudioGuid	int64 = 0
	AudioCatch	[]AudioInfo
	Record_MAX	int	= 2000
	CK_World  int8	= 1
	CK_GM	  int8	= 2
	CK_System int8	= 0
)

func BroadcastChat(info prpc.COM_Chat)  {
	for i:=0;i<len(PlayerStore) ;i++  {
		if PlayerStore[i]== nil {
			continue
		}
		if PlayerStore[i].session == nil {
			continue
		}
		PlayerStore[i].session.ReceiveChat(info)
	}
}

func (player *GamePlayer)RequestAudio(guid int64)  {
	dummy := []uint8{}
	audio := FindAudioData(guid)
	if player.session == nil {
		return
	}
	if audio==nil {
		player.session.RequestAudioOk(-1,dummy)
	}else {
		player.session.RequestAudioOk(audio.AudioId,audio.Audio)
	}

}

func (player *GamePlayer)SendChat(info prpc.COM_Chat)  {
	if info.Type == CK_GM {
		///啦啦啦啦啦啦
	}else {
		if info.Type == CK_System {

		}else if info.Type == CK_World {
			if len(info.Audio) != 0 {
				info.AudioId = PushAudioInfo(info.Audio)
				//for i:=0;i<len(info.Audio) ;i++  {
				//	info.Audio = append(info.Audio[:i],info.Audio[i+1:]...)
				//}
				info.Audio = nil
			}
		}
		BroadcastChat(info)
	}
}

func PushAudioInfo(audio []uint8) int64 {
	af := AudioInfo{}
	af.AudioId 	= atomic.AddInt64(&AudioGuid,1)
	af.Audio 	= audio
	AudioCatch 	= append(AudioCatch,af)
	if len(AudioCatch) > Record_MAX {
		i := 0
		AudioCatch = append(AudioCatch[:i],AudioCatch[i+1:]...)
	}

	return af.AudioId
}

func FindAudioData(guid int64) *AudioInfo {
	for i:=0;i < len(AudioCatch) ;i++  {
		if AudioCatch[i].AudioId == guid {
			return &AudioCatch[i]
		}
	}
	return nil
}

func (player *GamePlayer)TestChat()  {
	for i:=0;i<2010 ;i++  {
		af := prpc.COM_Chat{}
		u8Array := []uint8{11, 22, 33, 44, 55, 66}
		af.Type 		= CK_World

		af.Audio		= u8Array

		af.PlayerInstId = player.MyUnit.InstId
		af.PlayerName	= player.MyUnit.InstName
		af.HeadIcon		= "1111111"
		af.Content		= "test test test test test[123123123@$%^^&&*]"
		player.SendChat(af)
	}
	player.RequestAudio(2000)
}