using UnityEngine;
using System;
using System.Collections.Generic;

public class DisplayData {

    public int _Id;
    public string _AssetPath;
    public string _HeadIcon;
    public string _CardIcon;
    public float _Distance;

    static Dictionary<int, DisplayData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, DisplayData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("DisplayData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        DisplayData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new DisplayData ();
            data._Id = parser.GetInt (i, "ID");
            data._AssetPath = parser.GetString(i, "AssetsPath");
            data._HeadIcon = parser.GetString(i, "HeadIcon");
            data._CardIcon = parser.GetString(i, "CardIcon");
            data._Distance = parser.GetFloat(i, "Distance");

            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("DisplayData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public DisplayData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
