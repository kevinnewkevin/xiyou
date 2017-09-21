using UnityEngine;
using System;
using System.Collections.Generic;

public class DropData
{
    public int Id_;
    public int exp_;
    public int money_;
    public int item1_;
    public int itemNum1_;
    static public Dictionary<int, DropData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, DropData>();

        CSVParser parser = new CSVParser();
        if (!parser.Parse(content))
        {
            Debug.LogError("DropData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        DropData data = null;
        for (int i = 0; i < recordCounter; ++i)
        {
            data = new DropData();
            data.Id_ = parser.GetInt(i, "DropID");
            data.exp_ = parser.GetInt(i, "exp");
            data.money_ = parser.GetInt(i, "money");
            data.item1_ = parser.GetInt(i, "item-1");
            data.itemNum1_ = parser.GetInt(i, "item-num-1");


            if (metaData.ContainsKey(data.Id_))
            {
                Debug.LogError("DropData ID重复");
                return;
            }
            metaData[data.Id_] = data;
        }
        parser.Dispose();
        parser = null;
    }

    static public DropData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
