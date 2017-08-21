
using System.Collections.Generic;

public class GamePlayer {

    static public long _InstID;

    static public string _Name;

    static public bool _IsAuto;

    static public COM_Unit _Data;

    static public Dictionary<int, List<COM_Unit>> _CardsByFee = new Dictionary<int, List<COM_Unit>>();

    static public Dictionary<string, List<COM_Unit>> _CardsByType = new Dictionary<string, List<COM_Unit>>();

    static public List<List<COM_Unit>> _CardGroup = new List<List<COM_Unit>>();

    static public int _CrtBattleGroupIdx;

    static public void Init(COM_Player player)
    {
        Clear();
        for(int i=0; i < 5; ++i)
        {
            _CardGroup.Add(new List<COM_Unit>());
        }
        _InstID = player.InstId;
        _Name = player.Name;
        _Data = player.Unit;
        EntityData eData = null;
        for(int i=0; i < player.Employees.Length; ++i)
        {
            eData = EntityData.GetData(player.Employees[i].UnitId);
            if(!_CardsByFee.ContainsKey(eData._Cost))
                _CardsByFee.Add(eData._Cost, new List<COM_Unit>());
            _CardsByFee [eData._Cost].Add(player.Employees [i]);

            if(!_CardsByType.ContainsKey(eData.ToString()))
                _CardsByType.Add(eData.ToString(), new List<COM_Unit>());
            _CardsByType [eData.ToString()].Add(player.Employees [i]);

            if(!_CardsByFee.ContainsKey(0))
                _CardsByFee.Add(0, new List<COM_Unit>());
            _CardsByFee [0].Add(player.Employees [i]);
        }

        UIManager.SetDirty("zhujiemian");
    }

    //通过InstID获取卡牌
    static public COM_Unit GetCardByInstID(long instid)
    {
        if (_InstID == instid)
            return _Data;
        
        for(int i=0; i < _CardsByFee [0].Count; ++i)
        {
            if (_CardsByFee [0] [i].InstId == instid)
                return _CardsByFee [0] [i];
        }
        return null;
    }

    //通过索引获取卡组
    static public List<COM_Unit> GetGroupCards(int idx)
    {
        if (idx < 0 || idx >= _CardGroup.Count)
            return null;

        return _CardGroup[idx];
    }

    //是我的卡牌
    static public bool IsMy(long instid)
    {
        if (_InstID == instid)
            return true;
        
        for(int i=0; i < _CardsByFee [0].Count; ++i)
        {
            if (_CardsByFee [0] [i].InstId == instid)
                return true;
        }

        return false;
    }

    //是我自己
    static public bool IsMe(long instid)
    {
        return _InstID == instid;
    }

    //在我的卡组里
    static public bool IsInGroup(long instid, int groupidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return false;
        
        for(int j=0; j < _CardGroup[groupidx].Count; ++j)
        {
            if (_CardGroup [groupidx] [j].InstId == instid)
                return true;
        }
        return false;
    }

    //通过索引获得卡牌形象
    static public string GetResPathInMyCards(int idx)
    {
        if (idx < 0 || idx >= _CardsByFee [0].Count)
            return "";

        EntityData edata = EntityData.GetData(_CardsByFee [0][idx].UnitId);
        if (edata == null)
            return "";

        DisplayData ddata = DisplayData.GetData(edata._DisplayId);
        if (ddata == null)
            return "";
        
        return ddata._AssetPath;
    }

    static public string GetResPath(long instid)
    {
        for(int i=0; i < _CardsByFee [0].Count; ++i)
        {
            if (_CardsByFee [0] [i].InstId == instid)
            {
                EntityData edata = EntityData.GetData(_CardsByFee [0] [i].UnitId);
                if (edata == null)
                    return "";

                DisplayData ddata = DisplayData.GetData(edata._DisplayId);
                if (ddata == null)
                    return "";

                return ddata._AssetPath;
            }
        }
        return "";
    }

    //通过索引获得卡组中卡牌形象
    static public string GetResPathInMyGroup(int groupidx, int cardidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return "";

        if (cardidx < 0 || cardidx >= _CardGroup[groupidx].Count)
            return "";

        EntityData edata = EntityData.GetData(_CardGroup[groupidx][cardidx].UnitId);
        if (edata == null)
            return "";

        DisplayData ddata = DisplayData.GetData(edata._DisplayId);
        if (ddata == null)
            return "";

        return ddata._AssetPath;
    }

    static public string GetMyHeadIcon()
    {
        EntityData eData = EntityData.GetData(_Data.UnitId);
        if (eData == null)
            return string.Empty;

        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        if (dData == null)
            return string.Empty;

        return dData._HeadIcon;
    }

    //通过索引获得卡牌UnitID
    static public int GetUnitIDInMyCards(int idx)
    {
        if (idx < 0 || idx >= _CardsByFee [0].Count)
            return 0;
        
        return _CardsByFee [0][idx].UnitId;
    }

    //通过索引获得卡牌InstID
    static public long GetInstIDInMyCards(int fee, int idx)
    {
        if (!_CardsByFee.ContainsKey(fee))
            return 0;
        
        if (idx < 0 || idx >= _CardsByFee [fee].Count)
            return 0;

        return _CardsByFee [fee][idx].InstId;
    }

    //通过索引获得卡牌费用
    static public int GetFeeInMyCards(int fee, int idx)
    {
        if (!_CardsByFee.ContainsKey(fee))
            return -1;

        if (idx < 0 || idx >= _CardsByFee [fee].Count)
            return -1;

        EntityData eData = EntityData.GetData(_CardsByFee [fee][idx].UnitId);
        if (eData == null)
            return -1;
        
        return eData._Cost;
    }

    //通过索引获得卡组中卡牌InstID
    static public long GetInstIDInMyGroup(int groupidx, int cardidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return 0;

        if (cardidx < 0 || cardidx >= _CardGroup[groupidx].Count)
            return 0;

        return _CardGroup[groupidx][cardidx].InstId;
    }

    static public void PutInCard(long instid, int groupidx)
    {
        COM_Unit card =  GetCardByInstID(instid);
        if (card == null)
            return;
        
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return;

        _CardGroup [groupidx].Add(card);

        NetWoking.S.AddBattleUnit(instid, groupidx);
    }

    static public void TakeOffCard(long instid, int groupidx)
    {
        COM_Unit card =  GetCardByInstID(instid);
        if (card == null)
            return;
        
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return;

        for(int i=0; i < _CardGroup [groupidx].Count; ++i)
        {
            if (_CardGroup [groupidx] [i].InstId == instid)
            {
                _CardGroup [groupidx].RemoveAt(i);
                break;
            }
        }

        NetWoking.S.PopBattleUnit(instid, groupidx);
    }

    static public void DeleteGroup(int groupidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return;

        _CardGroup [groupidx].Clear();
    }

    static public List<COM_Unit> GetBattleCardsCopy()
    {
        return new List<COM_Unit>(_CardGroup [_CrtBattleGroupIdx]);
    }

    static public List<COM_Unit> CardsByFee(int fee)
    {
        if (!_CardsByFee.ContainsKey(fee))
            return null;
        return _CardsByFee[fee];
    }

    static public void Clear()
    {
        foreach(List<COM_Unit> list in _CardsByFee.Values)
        {
            list.Clear();
        }
        foreach(List<COM_Unit> list in _CardsByType.Values)
        {
            list.Clear();
        }
        _CardsByFee.Clear();
        _CardsByType.Clear();
        _CardGroup.Clear();
    }
}
