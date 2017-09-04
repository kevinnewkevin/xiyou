using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class COM_Item
{
    public int tableid;
    public int instid;
    public int stack;
}

public class BagSystem {

    public static Dictionary<int, List<COM_Item>> _ItemsByType;

    public static void Init(COM_Item[] items)
    {
        _ItemsByType = new Dictionary<int, List<COM_Item>>();
        _ItemsByType.Add(0, new List<COM_Item>(items));
        _ItemsByType [0].Sort(Sort);
        ItemData iData = null;
        for(int i=0; i < _ItemsByType [0].Count; ++i)
        {
            iData = ItemData.GetData(_ItemsByType [0][i].tableid);
            if (iData == null)
                continue;
            
            if (!_ItemsByType.ContainsKey(iData._Type))
                _ItemsByType.Add(iData._Type, new List<COM_Item>());
            _ItemsByType [iData._Type].Add(_ItemsByType [0][i]);
        }
        UIManager.SetDirty("bagui");
    }

    static int Sort(COM_Item i1, COM_Item i2)
    {
        if (i1.tableid > i2.tableid)
            return -1;
        if (i1.tableid < i2.tableid)
            return 1;
        else
            return 0;
    }

    public static void AddItem(COM_Item inst)
    {
        _ItemsByType [0].Add(inst);
        ItemData iData = ItemData.GetData(inst.tableid);
        if (iData != null)
        {
            if (!_ItemsByType.ContainsKey(iData._Type))
                _ItemsByType.Add(iData._Type, new List<COM_Item>());
            _ItemsByType [iData._Type].Add(inst);
        }
        UIManager.SetDirty("bagui");
    }

    public static void DelItem(int instid)
    {
        int type = 0;
        for(int i=0; i < _ItemsByType [0].Count; ++i)
        {
            if (_ItemsByType [0] [i].instid == instid)
            {
                ItemData iData = ItemData.GetData(_ItemsByType [0] [i].tableid);
                type = iData._Type;
                _ItemsByType [0].RemoveAt(i);
                break;
            }
        }
        for(int i=0; i < _ItemsByType[type].Count; ++ i)
        {
            if (_ItemsByType [type] [i].instid == instid)
            {
                _ItemsByType [type].RemoveAt(i);
            }
        }
        UIManager.SetDirty("bagui");
    }

    public static void UpdateItem(COM_Item inst)
    {
        int type = 0;
        for(int i=0; i < _ItemsByType [0].Count; ++i)
        {
            if (_ItemsByType [0] [i].instid == inst.instid)
            {
                ItemData iData = ItemData.GetData(_ItemsByType [0] [i].tableid);
                type = iData._Type;
                _ItemsByType [0] [i] = inst;
                break;
            }
        }
        for(int i=0; i < _ItemsByType[type].Count; ++i)
        {
            if (_ItemsByType [type] [i].instid == inst.instid)
                _ItemsByType [type] [i] = inst;
        }
        UIManager.SetDirty("bagui");
    }

    static COM_Item GetItemInstByIndex(int idx, int crtTab = 0)
    {
        if (!_ItemsByType.ContainsKey(crtTab))
            return null;
        
        if (idx < 0 || idx >= _ItemsByType[crtTab].Count)
            return null;

        return _ItemsByType[crtTab][idx];
    }

    static int GetCount(int crtTab)
    {
        if (!_ItemsByType.ContainsKey(crtTab))
            return 0;

        return _ItemsByType [crtTab].Count;
    }

    static public void Fini()
    {
        if(_ItemsByType != null)
            _ItemsByType.Clear();
    }
}
