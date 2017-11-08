using System.Collections.Generic;

public class ChatSystem {

    // type -1 为全部信息 前端自定义
    static public Dictionary<int, List<COM_Chat>> _AllMsg;

    static public void Init()
    {
        _AllMsg = new Dictionary<int, List<COM_Chat>>();
        _AllMsg.Add(-1, new List<COM_Chat>());
    }

    static public void AddMsg(COM_Chat chat)
    {
        if (!_AllMsg.ContainsKey(chat.Type))
            _AllMsg.Add(chat.Type, new List<COM_Chat>());

        _AllMsg [chat.Type].Add(chat);
        _AllMsg [-1].Add(chat);

        UIManager.SetDirty("zhujiemian_liaotian");
        UIManager.SetDirty("liaotian");
    }

    static public void SetRecord(long id, byte[] data)
    {
        for(int i=0; i < _AllMsg[-1].Count; ++i)
        {
            if (_AllMsg [-1] [i].AudioId != 0 && _AllMsg [-1] [i].AudioId == id)
            {
                _AllMsg [-1] [i].Audio = data;
                UIManager.SetDirty("zhujiemian_liaotian");
                UIManager.SetDirty("liaotian");
                break;
            }
        }
    }

    static public byte[] GetRecord(long id)
    {
        for(int i=0; i < _AllMsg[-1].Count; ++i)
        {
            if (_AllMsg [-1] [i].AudioId != 0 && _AllMsg [-1] [i].AudioId == id)
            {
                return _AllMsg [-1] [i].Audio;
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

    static public void Clear()
    {
        
    }
}
