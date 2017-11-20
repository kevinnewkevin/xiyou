using System.Collections.Generic;

public class ChatSystem {

    // type -1 为全部信息 前端自定义
    static public Dictionary<int, List<COM_Chat>> _AllMsg;

    static public List<COM_Assistant> _Assistant;

    static public void Init()
    {
        _Assistant = new List<COM_Assistant>();
        _AllMsg = new Dictionary<int, List<COM_Chat>>();
        _AllMsg.Add(-1, new List<COM_Chat>());
        EmojiParser.inst.RegistEmojiTags();
    }

    static public void AddMsg(COM_Chat chat)
    {
        if (!_AllMsg.ContainsKey(chat.Type))
            _AllMsg.Add(chat.Type, new List<COM_Chat>());

        _AllMsg [chat.Type].Add(chat);
        _AllMsg [-1].Add(chat);

        UIManager.SetDirty("zhujiemian_liaotian");
        UIManager.SetDirty("liaotian");
        UIManager.SetDirty("squad");
    }

    static public void UpdateAss(COM_Assistant ass)
    {
        bool updated = false;
        for(int i=0; i < _Assistant.Count; ++i)
        {
            if (_Assistant [i].Id == ass.Id)
            {
                _Assistant [i] = ass;
                updated = true;
            }
        }
        if (!updated)
        {
            _Assistant.Add(ass);
            if(ass.PlayerName.Equals(GamePlayer._Name))
                GamePlayer.AddCoolDown("AssistantCoolDown", TimerManager.GetTimeStamp());
            else
                LuaManager.Call("global.lua", "ReceivedAssistant", ass);
        }

        UIManager.SetDirty("liaotian");
        UIManager.SetDirty("squad");
    }

    static public COM_Assistant GetAss(int assid)
    {
        for(int i=0; i < _Assistant.Count; ++i)
        {
            if (_Assistant [i].Id == assid)
            {
                return _Assistant [i];
            }
        }
        return null;
    }

    static public void SetRecord(string audioid)
    {
        for(int i=0; i < _AllMsg[-1].Count; ++i)
        {
            if (_AllMsg [-1] [i].AudioId.Equals(audioid))
            {
                _AllMsg [-1] [i].AudioOld = true;
                UIManager.SetDirty("zhujiemian_liaotian");
                UIManager.SetDirty("liaotian");
                UIManager.SetDirty("squad");
                break;
            }
        }
    }

    static public COM_Chat GetRecord(string audioid)
    {
        for(int i=0; i < _AllMsg[-1].Count; ++i)
        {
            if (_AllMsg [-1] [i].AudioId.Equals(audioid))
            {
                return _AllMsg [-1] [i];
            }
        }
        return null;
    }

    static public List<COM_Chat> MsgByType(int type)
    {
        if (!_AllMsg.ContainsKey(type))
            return new List<COM_Chat>();

        return _AllMsg [type];
    }

    static public List<COM_Chat> LastestMsgByType(int type, int num)
    {
        List<COM_Chat> tList = new List<COM_Chat>();
        if (!_AllMsg.ContainsKey(type))
            return tList;
        
        for(int i = _AllMsg[type].Count - 1; i >= 0; --i)
        {
            tList.Add(_AllMsg[type][i]);
            if (tList.Count >= num)
                break;
        }
        tList.Reverse();
        return tList;
    }

    static public void ClearMsgByType(int type)
    {
        if (_AllMsg.ContainsKey(type))
            _AllMsg[type].Clear();
    }

    static public void Clear()
    {
        
    }
}
