using UnityEngine;
using System;
using System.Collections.Generic;

public class BuffData {

    public int _Id;
    public string _Icon;
    public string _Anim;
    public string _Effect;
    public float _AnimTime;
    public float _EffectTime;
    public float _EmitTime;

    static Dictionary<int, BuffData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, BuffData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("BuffData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        BuffData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new BuffData ();
            data._Id = parser.GetInt (i, "BuffId");
            data._Icon = parser.GetString(i, "Icon");

            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("BuffData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public BuffData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
