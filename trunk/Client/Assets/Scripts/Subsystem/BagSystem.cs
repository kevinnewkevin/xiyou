using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BagSystem {

    public static Dictionary<int, List<COM_ItemInst>> _ItemsByType;
	public static Dictionary<int, List<COM_ItemInst>> _DebrisItemsByType;

	public static void Init(COM_ItemInst[] items = null)
    {
		_ItemsByType = new Dictionary<int, List<COM_ItemInst>>();
		_DebrisItemsByType = new Dictionary<int, List<COM_ItemInst>>();
		List<COM_ItemInst> itemList = new List<COM_ItemInst> ();
		List<COM_ItemInst> debrisItemList = new List<COM_ItemInst> ();	
        if ( items != null )
        {
            for (int i = 0; i < items.Length; i++)
            {
                itemList.Add(items[i]);
            }
        }

        if (!_ItemsByType.ContainsKey(0))
            _ItemsByType.Add(0, itemList);
        else
            _ItemsByType [0] = itemList;

		if (!_DebrisItemsByType.ContainsKey (0))
				_DebrisItemsByType.Add (0, debrisItemList);
       // _ItemsByType [0].Sort(Sort);
        ItemData iData = null;
		if ( items != null )
		{
			for (int k = 0; k < items.Length; ++k) 
			{
				iData = ItemData.GetData (items [k].ItemId);
				if (iData == null)
						continue;
				if (iData._Type == "IMT_Debris") 
				{
					_DebrisItemsByType [0].Add (items [k]);
				}
			}
		}
        
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
					
			if (iData._Type == "IMT_Debris")
				_DebrisItemsByType [0].Add(inst);
					
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
              //  type = iData._Type;
                _ItemsByType [0].RemoveAt(i);
                break;
            }
        }
		for(int i=0; i < _DebrisItemsByType[0].Count; ++ i)
        {
			if (_DebrisItemsByType [0] [i].InstId == instid)
            {
				_DebrisItemsByType [0].RemoveAt(i);
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
               // type = iData._Type;
                _ItemsByType [0] [i] = inst;
                break;
            }
        }
		for(int i=0; i < _DebrisItemsByType[0].Count; ++i)
        {
			if (_DebrisItemsByType [0] [i].InstId == inst.InstId)
				_DebrisItemsByType [0] [i] = inst;
        }
        UIManager.SetDirty("bagui");
    }

	public static  COM_ItemInst GetItemInstByIndex(int idx, int crtTab = 0)
    {
		if (crtTab == 1) 
		{
			if (idx < 0 || idx >= _DebrisItemsByType[0].Count)
					return null;
			return _DebrisItemsByType[0][idx];
		}
        if (!_ItemsByType.ContainsKey(crtTab))
            return null;
        
        if (idx < 0 || idx >= _ItemsByType[crtTab].Count)
            return null;

        return _ItemsByType[crtTab][idx];
    }

	public static  int GetItemCount(int crtTab)
{
		if (crtTab == 1) 
		{
			return _DebrisItemsByType [0].Count;
		}
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
                num += _ItemsByType[0][i].Stack;
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
