using UnityEngine;
using System;
using System.Collections.Generic;

public class ItemData {
    
    public int _Id;
    public string _Type;
    public string _Icon;
	public string _IconBack;
	public string _Name;
	public string _Desc;
	public string _EntityID;
	public int _Price;
    public static Dictionary<int, ItemData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, ItemData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("ItemData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        ItemData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new ItemData ();
            data._Id = parser.GetInt (i, "ItemID");
			data._Type = parser.GetString (i, "ItemType");
            data._Icon = parser.GetString (i, "ItemIcon");
			data._IconBack = parser.GetString (i, "Quality");
			data._Name = parser.GetString (i, "Name");
			data._Desc = parser.GetString (i, "Desc");
			data._EntityID = parser.GetString (i, "entityID");
			data._Price = parser.GetInt(i, "Price");
            
			if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("ItemData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

	public static ItemData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
