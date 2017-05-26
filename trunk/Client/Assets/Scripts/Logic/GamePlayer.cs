using protocol;
using System.Collections.Generic;

public class GamePlayer {

    static public ulong _InstID;

    static public string _Name;

    static public COM_EntityInstance _Data;

    static public List<COM_EntityInstance> _Cards;

    static public void Init(protocol.COM_PlayerInstance player)
    {
        _InstID = player.InstanceId;
        _Name = player.PlayerName;
        _Data = player.PlayerEntity;
        _Cards = player.Employees;
    }

    static public COM_EntityInstance GetCardByInstID(ulong instid)
    {
        for(int i=0; i < _Cards.Count; ++i)
        {
            if (_Cards [i].InstanceId == instid)
                return _Cards [i];
        }
        return null;
    }
}
