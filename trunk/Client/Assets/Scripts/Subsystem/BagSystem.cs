using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BagSystem {

    public static Dictionary<int, List<COM_ItemInst>> _ItemsByType;

	public static void Init(COM_ItemInst[] items)
    {
		_ItemsByType = new Dictionary<int, List<COM_ItemInst>>();
		List<COM_ItemInst> itemList = new List<COM_ItemInst> ();

        if ( items != null )
        {
            for (int i = 0; i < items.Length; i++)
            {
                itemList.Add(items[i]);
            }
        }
        
		_ItemsByType.Add(0, itemList);

       // _ItemsByType [0].Sort(Sort);
        ItemData iData = null;
		
	/*	for(int i=0; i < items.Length; ++i)
        {
			iData = ItemData.GetData(items[i].ItemId);
            if (iData == null)
                continue;
            if (!_ItemsByType.ContainsKey(iData._Type))
					_ItemsByType.Add(iData._Type, new List<COM_ItemInst>());
			_ItemsByType [iData._Type].Add(items[i]);
        }
        */
        
        UIManager.SetDirty("bagui");
    }

	static int Sort(COM_ItemInst i1, COM_ItemInst i2)
    {
        if (i1.ItemId > i2.ItemId)
            return -1;
        if (i1.ItemId < i2.ItemId)
            return 1;
        else
            return 0;
    }

	public static void AddItem(COM_ItemInst inst)
    {
        _ItemsByType [0].Add(inst);
			ItemData iData = ItemData.GetData(inst.ItemId);
        if (iData != null)
        {
            //if (!_ItemsByType.ContainsKey(iData._Type))
			//	_ItemsByType.Add(iData._Type, new List<COM_ItemInst>());
            //_ItemsByType [iData._Type].Add(inst);
        }
        UIManager.SetDirty("bagui");
    }

    public static void DelItem(long instid)
    {
        int type = 0;
        for(int i=0; i < _ItemsByType [0].Count; ++i)
        {
			if (_ItemsByType [0] [i].InstId == instid)
            {
				ItemData iData = ItemData.GetData(_ItemsByType [0] [i].ItemId);
                type = iData._Type;
                _ItemsByType [0].RemoveAt(i);
                break;
            }
        }
        for(int i=0; i < _ItemsByType[type].Count; ++ i)
        {
			if (_ItemsByType [type] [i].InstId == instid)
            {
                _ItemsByType [type].RemoveAt(i);
            }
        }
        UIManager.SetDirty("bagui");
    }

	public static void UpdateItem(COM_ItemInst inst)
    {
        int type = 0;
        for(int i=0; i < _ItemsByType [0].Count; ++i)
        {
			if (_ItemsByType [0] [i].InstId == inst.InstId)
            {
				ItemData iData = ItemData.GetData(_ItemsByType [0] [i].ItemId);
                type = iData._Type;
                _ItemsByType [0] [i] = inst;
                break;
            }
        }
        for(int i=0; i < _ItemsByType[type].Count; ++i)
        {
						if (_ItemsByType [type] [i].InstId == inst.InstId)
                _ItemsByType [type] [i] = inst;
        }
        UIManager.SetDirty("bagui");
    }

	public static  COM_ItemInst GetItemInstByIndex(int idx, int crtTab = 0)
    {
        if (!_ItemsByType.ContainsKey(crtTab))
            return null;
        
        if (idx < 0 || idx >= _ItemsByType[crtTab].Count)
            return null;

        return _ItemsByType[crtTab][idx];
    }

	public static  int GetItemCount(int crtTab)
    {
        if (!_ItemsByType.ContainsKey(crtTab))
            return 0;
		return _ItemsByType [crtTab].Count;
    }

    public static int GetItemMaxNum(int itemId)
    {
        int num = 0;
        for (int i = 0; i < _ItemsByType[0].Count; ++i)
        {
            if (_ItemsByType[0][i].ItemId == itemId)
            {
                num += _ItemsByType[0][i].Stack_;
            }
        }
        return num;
    }

    static public void Fini()
    {
        if(_ItemsByType != null)
            _ItemsByType.Clear();
    }
}
