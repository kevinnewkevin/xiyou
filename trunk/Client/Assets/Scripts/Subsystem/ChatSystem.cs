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
