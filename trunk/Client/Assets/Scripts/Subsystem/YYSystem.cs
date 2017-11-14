using System;
using UnityEngine;
using YunvaIM;

public class YYSystem {

    private const uint sUserId = 1002318;

    static string _LastFilePath;
    static int _LastRecordLength;
    static int _LastChatType;

    static public void Init()
    {
        #if UNITY_EDITOR 
        return;
        #endif
        EventListenerManager.AddListener(ProtocolEnum.IM_RECORD_VOLUME_NOTIFY, ImRecordVolume);//录音音量大小回调监听
        int init = YunVaImSDK.instance.YunVa_Init(0, sUserId, Application.persistentDataPath, false, false);
        if (init == 0)
        {
            Debug.Log("初始化成功...");
        }
        else
        {
            Debug.Log("初始化失败...");
        }
    }

    static public void ImRecordVolume(object data)
    {
        ImRecordVolumeNotify RecordVolumeNotify = data as ImRecordVolumeNotify;
    }

    static public void Login()
    {
        #if UNITY_EDITOR 
        return;
        #endif
        string ttFormat = "{{\"nickname\":\"{0}\",\"uid\":\"{1}\"}}";
        string tt = string.Format(ttFormat, sUserId, sUserId);
        string[] wildcard = new string[2];
        wildcard[0] = "0x001";
        wildcard[1] = "0x002";
        YunVaImSDK.instance.YunVaOnLogin(tt, "1", wildcard, 0, (data) => 
        {
            if (data.result == 0)
            {
                Debug.Log(string.Format("登录成功，昵称:{0},用户ID:{1}", data.nickName, data.userId));
                YunVaImSDK.instance.RecordSetInfoReq(true);//开启录音的音量大小回调
            }
            else
            {
                Debug.Log(string.Format("登录失败，错误消息：{0}", data.msg));
            }
        });
    }

    static public void StartRecord()
    {
        #if UNITY_EDITOR
        return;
        #endif

        string filePath = string.Format("{0}/{1}.amr", Application.persistentDataPath, DateTime.Now.ToFileTime());
        YunVaImSDK.instance.RecordStartRequest(filePath,2);
    }

    static public void StopRecord(bool cancel, int chatType)
    {
        #if UNITY_EDITOR 
        return;
        #endif

        if (cancel)
        {
            YunVaImSDK.instance.RecordStopRequest((data1) => {
                
            });
            return;
        }
        _LastChatType = chatType;

        YunVaImSDK.instance.RecordStopRequest((data1) => {
//            recordPath = data1.strfilepath;
//            Debug.Log("停止录音返回:" + recordPath);
            _LastRecordLength = (int)(data1.time / 1000f);
            _LastFilePath = data1.strfilepath;
        }, 
            (data2) => {
            Debug.Log("上传返回:" + data2.fileurl);
            COM_Chat chat = new COM_Chat();
            chat.AudioId = System.Guid.NewGuid().ToString();
            chat.AudioUrl = data2.fileurl;
            chat.AudioPath = _LastFilePath;
            chat.Type = _LastChatType;
            chat.AudioLen = _LastRecordLength;
            chat.PlayerInstId = GamePlayer._InstID;
            chat.PlayerName = GamePlayer._Name;
            chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
            chat.Level = GamePlayer._Data.IProperties[9].ToString();
            NetWoking.S.SendChat(chat);
        }, 
            (data3) => {
            Debug.Log("识别返回:" + data3.text);
        });
    }

    static public void PlayRecord(string path, string url)
    {
        #if UNITY_EDITOR 
        return;
        #endif

        YunVaImSDK.instance.RecordStartPlayRequest(path, url, "", (data2) =>
        {
            if (data2.result == 0)
            {
                Debug.Log("播放成功");  
            }
            else
            {
                Debug.Log("播放失败");
            }
        });
    }

    static public void StopPlayRecord()
    {
        YunVaImSDK.instance.RecordStopPlayRequest();
    }

    static public void Logout()
    {
        YunVaImSDK.instance.YunVaLogOut();
    }
}
