using UnityEngine;
using System;
using System.Collections.Generic;

public class TalkData {

    public int _Id;
    public string _Name;
    public string _Content;
    public int _Side;
    public int _DisplayId;
    public float _Scale;

    static Dictionary<int, TalkData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, TalkData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("TalkData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        TalkData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new TalkData ();
            data._Id = parser.GetInt (i, "id");
            data._Name = parser.GetString(i, "name");
            data._Content = parser.GetString(i, "content");
            data._DisplayId = parser.GetInt(i, "DisplayId");
            data._Side = parser.GetInt(i, "Side");
            data._Scale = parser.GetFloat(i, "Scale");

            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("TalkData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public TalkData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
