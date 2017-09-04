using UnityEngine;
using System;
using System.Collections.Generic;

public class ItemData {
    
    public int _Id;
    public int _Type;
    public string _Icon;

    static Dictionary<int, ItemData> metaData;

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
            data._Type = parser.GetInt (i, "ItemType");
            data._Icon = parser.GetString (i, "ItemIcon");

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

    static public ItemData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
