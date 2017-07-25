
using System.Collections.Generic;

public class GamePlayer {

    static public long _InstID;

    static public string _Name;

    static public bool _IsAuto;

    static public COM_Unit _Data;

    static public List<COM_Unit> _Cards = new List<COM_Unit>();

    static public void Init(COM_Player player)
    {
        _InstID = player.InstId;
        _Name = player.Name;
        _Data = player.Unit;
        _Cards.AddRange(player.Employees);
        UIManager.SetDirty("zhujiemian");
    }

    static public COM_Unit GetCardByInstID(long instid)
    {
        if (_InstID == instid)
            return _Data;
        
        for(int i=0; i < _Cards.Count; ++i)
        {
            if (_Cards [i].InstId == instid)
                return _Cards [i];
        }
        return null;
    }

    static public bool IsMy(long instid)
    {
        if (_InstID == instid)
            return true;
        
        for(int i=0; i < _Cards.Count; ++i)
        {
            if (_Cards [i].InstId == instid)
                return true;
        }

        return false;
    }

    static public bool IsMe(long instid)
    {
        return _InstID == instid;
    }
}
