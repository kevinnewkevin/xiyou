using UnityEngine;
using System.Collections.Generic;

public class GamePlayer {

    static public long _InstID;

    static public string _Name;

    static public bool _IsAuto;

    static public COM_Unit _Data;

    static public Dictionary<string, List<COM_Unit>> _Cards = new Dictionary<string, List<COM_Unit>>();

    static public List<List<long>> _CardGroup = new List<List<long>>();
    static public List<string> _CardGroupName = new List<string>();

	static public int _TianTiVal;
	
    static public int _CrtBattleGroupIdx;

    static public List<string> _IPropDirty = new List<string>();
    static public List<string> _CPropDirty = new List<string>();
	static public bool showNewCard;
	static public COM_Unit newCard;
    static public void Init(COM_Player player)
    {
        Clear();
        string name = "";
		for (int i = 0; i < player.UnitGroup.Length; ++i) 
		{
			_CardGroup.Add (new List<long> ());
		}
        for(int i=0; i < _CardGroup.Count; ++i)
        {
            if (player.UnitGroup [i].UnitList != null)
            {
//                for(int j=0; j < player.UnitGroup [i].UnitList.Length; ++j)
                    _CardGroup [player.UnitGroup [i].GroupId - 1].AddRange(player.UnitGroup [i].UnitList);
            }

            name = PlayerPrefs.GetString("XYSK_XIYOU_ACCOUNT_PLUGINID" + i);
            _CardGroupName.Add(name);
        }

        _InstID = player.InstId;
        _Name = player.Name;
        _Data = player.Unit;
		_TianTiVal = player.TianTiVal;
        JieHunSystem.instance.InitChapterData(player.Chapters);
        EntityData eData = null;
        for(int i=0; i < player.Employees.Length; ++i)
        {
            eData = EntityData.GetData(player.Employees[i].UnitId);
            string costTypeKey = eData._Cost + "_" + eData._Type;
            if(!_Cards.ContainsKey(costTypeKey))
                _Cards.Add(costTypeKey, new List<COM_Unit>());
            _Cards [costTypeKey].Add(player.Employees [i]);

            if(!_Cards.ContainsKey("0_0"))
                _Cards.Add("0_0", new List<COM_Unit>());
            _Cards ["0_0"].Add(player.Employees [i]);

            if(!_Cards.ContainsKey("0_" + eData._Type))
                _Cards.Add("0_" + eData._Type, new List<COM_Unit>());
            _Cards ["0_" + eData._Type].Add(player.Employees [i]);

            if(!_Cards.ContainsKey(eData._Cost + "_0"))
                _Cards.Add(eData._Cost + "_0", new List<COM_Unit>());
            _Cards [eData._Cost + "_0"].Add(player.Employees [i]);
        
        }

        if (player.SkillBase != null)
        {
            for(int i=0; i < player.SkillBase.Length; ++i)
            {
                RoleSkillData.SetData(player.SkillBase[i].SkillIdx, player.SkillBase[i].SkillId);
            }
        }

        GuideSystem.SetProgress(player.Guide);

        UIManager.SetDirty("zhujiemian");
    }

    static public void AddCard(COM_Unit card)
    {
        EntityData eData = EntityData.GetData(card.UnitId);
        string costTypeKey = eData._Cost + "_" + eData._Type;
        if(!_Cards.ContainsKey(costTypeKey))
            _Cards.Add(costTypeKey, new List<COM_Unit>());
        _Cards [costTypeKey].Add(card);

        if(!_Cards.ContainsKey("0_" + eData._Type))
            _Cards.Add("0_" + eData._Type, new List<COM_Unit>());
        _Cards ["0_" + eData._Type].Add(card);

        if(!_Cards.ContainsKey(eData._Cost + "_0"))
            _Cards.Add(eData._Cost + "_0", new List<COM_Unit>());
        _Cards [eData._Cost + "_0"].Add(card);

        if(!_Cards.ContainsKey("0_0"))
            _Cards.Add("0_0", new List<COM_Unit>());
        _Cards ["0_0"].Add(card);
    }

    static public void UpdateEquipedSkill(int idx, int skillid)
    {
        if (_Data == null)
            return;

        if (_Data.Skills == null)
            return;

        for(int i=0; i < _Data.Skills.Length; ++i)
        {
            if (_Data.Skills [i].SkillId == skillid)
            {
                _Data.Skills [i].SkillId = 0;
                UIManager.SetDirty("jineng");
            }
        }

        for(int i=0; i < _Data.Skills.Length; ++i)
        {
            if (_Data.Skills [i].Pos == idx)
            {
                _Data.Skills [i].SkillId = skillid;
                UIManager.SetDirty("jineng");
            }
        }
    }

    //通过InstID获取卡牌
    static public COM_Unit GetCardByInstID(long instid)
    {
        if (_InstID == instid)
            return _Data;
        
        for(int i=0; i < _Cards ["0_0"].Count; ++i)
        {
            if (_Cards ["0_0"] [i].InstId == instid)
                return _Cards ["0_0"] [i];
        }
        return null;
    }


	static public COM_Unit GetCardByEntityID(int unitId)
	{
        for(int i=0; i < _Cards ["0_0"].Count; ++i)
		{
            if (_Cards ["0_0"] [i].UnitId== unitId)
                return _Cards ["0_0"] [i];
		}
		return null;
	}


    static public void UpdateUnitIProperty(long instId, int type, int vaule)
    {
        for (int i = 0; i < _Cards ["0_0"].Count; ++i)
        {
            if (_Cards ["0_0"][i].InstId == instId)
            {
                _Cards ["0_0"][i].IProperties[type] = vaule;
            }
        }

        if (_InstID == instId)
        {
			if (type == 11) 
			{
				if (vaule > _Data.IProperties [type]) 
				{
					Proxy4Lua.PopMsg("获得魂币" + (vaule-_Data.IProperties [type]));	
				}
			}
            _Data.IProperties [type] = vaule;

        }

        for(int i=0; i < _IPropDirty.Count; ++i)
        {
            UIManager.SetDirty(_IPropDirty[i]);
        }
    }

    static public void UpdateUnitCProperty(long instId, int type, float vaule)
    {
        for (int i = 0; i < _Cards ["0_0"].Count; ++i)
        {
            if (_Cards ["0_0"][i].InstId == instId)
            {
                _Cards ["0_0"][i].CProperties[type] = vaule;
            }
        }

        if (_InstID == instId)
        {
            _Data.CProperties [type] = vaule;
        }

        for(int i=0; i < _CPropDirty.Count; ++i)
        {
            UIManager.SetDirty(_CPropDirty[i]);
        }
    }

    //通过索引获取卡组
    static public List<long> GetGroupCards(int idx)
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
        
        for(int i=0; i < _Cards ["0_0"].Count; ++i)
        {
            if (_Cards ["0_0"] [i].InstId == instid)
                return true;
        }

        return false;
    }

    static public void ChangeGroupName(int groupidx, string name)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return;
        
        _CardGroupName[groupidx] = name;
        PlayerPrefs.SetString("XYSK_XIYOU_ACCOUNT_PLUGINID" + groupidx, name);
    }

    static public string GetGroupName(int groupidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroupName.Count)
            return "";

        return _CardGroupName [groupidx];
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
            if (_CardGroup [groupidx] [j] == instid)
                return true;
        }
        return false;
    }

    //卡组最大数
    static public bool IsGroupMax(int groupidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return false;

        return _CardGroup [groupidx].Count == 10;
    }

    //通过索引获得卡牌形象
    static public DisplayData GetDisplayDataByIndex(int fee, int type, int idx)
    {
        if (!_Cards.ContainsKey(fee + "_" + type))
            return null;

        if (idx < 0 || idx >= _Cards[fee + "_" + type].Count)
            return null;

        EntityData edata = EntityData.GetData(_Cards[fee + "_" + type][idx].UnitId);
        if (edata == null)
            return null;

        DisplayData ddata = DisplayData.GetData(edata._DisplayId);
        return ddata;
    }

    static public DisplayData GetDisplayDataByInstID(long instid)
    {
        for(int i=0; i < _Cards ["0_0"].Count; ++i)
        {
            if (_Cards ["0_0"] [i].InstId == instid)
            {
                EntityData edata = EntityData.GetData(_Cards ["0_0"] [i].UnitId);
                if (edata == null)
                    return null;

                DisplayData ddata = DisplayData.GetData(edata._DisplayId);
                return ddata;
            }
        }
        return null;
    }

    static public EntityData GetEntityDataByIndex(int fee, int type, int idx)
    {
        if (!_Cards.ContainsKey(fee + "_" + type))
            return null;

        if (idx < 0 || idx >= _Cards[fee + "_" + type].Count)
            return null;

        EntityData edata = EntityData.GetData(_Cards[fee + "_" + type][idx].UnitId);
        return edata;
    }

    static public EntityData GetEntityDataByInstID(long instid)
    {
        if (instid == _InstID)
            return GetMyEntityData();
        
        for (int i = 0; i < _Cards["0_0"].Count; ++i)
        {
            if (_Cards["0_0"][i].InstId == instid)
            {
                EntityData edata = EntityData.GetData(_Cards["0_0"][i].UnitId);
                return edata;
            }
        }
        return null;
    }

    //通过索引获得卡组中卡牌形象
    static public DisplayData GetDisplayDataByIndexFromGroup(int groupidx, int cardidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return null;

        if (cardidx < 0 || cardidx >= _CardGroup[groupidx].Count)
            return null;

        COM_Unit unit = GetCardByInstID(_CardGroup[groupidx][cardidx]);
        if (unit == null)
            return null;

        EntityData edata = EntityData.GetData(unit.UnitId);
        if (edata == null)
            return null;

        DisplayData ddata = DisplayData.GetData(edata._DisplayId);
        return ddata;
    }

    //通过索引获得卡组中卡牌形象
    static public EntityData GetEntityDataByIndexFromGroup(int groupidx, int cardidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return null;

        if (cardidx < 0 || cardidx >= _CardGroup[groupidx].Count)
            return null;

        COM_Unit unit = GetCardByInstID(_CardGroup[groupidx][cardidx]);
        if (unit == null)
            return null;

        EntityData edata = EntityData.GetData(unit.UnitId);
        return edata;
    }

    static public DisplayData GetMyDisplayData()
    {
        EntityData eData = EntityData.GetData(_Data.UnitId);
        if (eData == null)
            return null;

        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        return dData;
    }

    static public EntityData GetMyEntityData()
    {
        EntityData eData = EntityData.GetData(_Data.UnitId);
        if (eData == null)
            return null;

        return eData;
    }

    static public int[] GetMyActiveSkill()
    {
        List<int> skills = new List<int>();
        RoleSkillData rsData;
        for (int i = 0; i < _Data.Skills.Length; ++i)
        {
            rsData = RoleSkillData.GetDataBySkillID(_Data.Skills[i].SkillId);
            if (rsData._Type == 1)
                skills.Add(rsData._SkillId);
        }
        return skills.ToArray();
    }

    //通过索引获得卡牌InstID
    static public long GetInstID(int fee, int type, int idx)
    {
        if (!_Cards.ContainsKey(fee + "_" + type))
            return 0;
        
        if (idx < 0 || idx >= _Cards [fee + "_" + type].Count)
            return 0;

        return _Cards [fee + "_" + type][idx].InstId;
    }

    //通过索引获得卡组中卡牌InstID
    static public long GetInstIDFromGroup(int groupidx, int cardidx)
    {
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return 0;

        if (cardidx < 0 || cardidx >= _CardGroup[groupidx].Count)
            return 0;

        return _CardGroup[groupidx][cardidx];
    }

    static public void PutInCard(long instid, int groupidx)
    {
        COM_Unit card =  GetCardByInstID(instid);
        if (card == null)
            return;
        
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return;

        _CardGroup [groupidx].Add(instid);

        NetWoking.S.SetBattleUnit(instid, groupidx + 1, true);
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
            if (_CardGroup [groupidx] [i] == instid)
            {
                _CardGroup [groupidx].RemoveAt(i);
                break;
            }
        }

        NetWoking.S.SetBattleUnit(instid, groupidx + 1, false);
    }

    static public void DeleteGroup(int groupidx)
	{
        if (groupidx < 0 || groupidx >= _CardGroup.Count)
            return;

        _CardGroup [groupidx].Clear();
        NetWoking.S.DelUnitGroup(groupidx + 1);
    }

    static public List<long> GetBattleCardsCopy()
    {
        return new List<long>(_CardGroup [_CrtBattleGroupIdx]);
    }

    static public List<COM_Unit> CardsByFeeAndType(int fee, int type)
    {
        if (!_Cards.ContainsKey(fee + "_" + type))
            return null;
        return _Cards[fee + "_" + type];
    }

	static public int GetTianTiLevel()
	{
				if (_TianTiVal < 400)
						return 1;
				else if (_TianTiVal > 400 && _TianTiVal <= 700)
						return 2;
				else if (_TianTiVal > 700 && _TianTiVal <= 1000)
						return 3;
				else if (_TianTiVal > 1000 && _TianTiVal <= 1400)
						return 4;
				else if (_TianTiVal > 1400 && _TianTiVal <= 1700)
						return 5;
				else if (_TianTiVal > 1700 && _TianTiVal <= 2000)
						return 6;
				else if (_TianTiVal > 2000 && _TianTiVal <= 2500)
						return 7;
				else if (_TianTiVal > 2500 && _TianTiVal <= 3000)
						return 8;
				else if (_TianTiVal > 3000 && _TianTiVal <= 4000)
						return 9;
				else if (_TianTiVal > 4000)
						return 10;
				else
					return 1;
			
	}

    static public void Clear()
    {
        foreach(List<COM_Unit> list in _Cards.Values)
        {
            list.Clear();
        }
        _Cards.Clear();
        _CardGroup.Clear();
    }
}
