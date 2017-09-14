using UnityEngine;
using System;
using System.Collections.Generic;

public class ExpData {

    public int _Level;
    public int _Exp;

    static Dictionary<int, ExpData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, ExpData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("ExpData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        ExpData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new ExpData ();
            data._Level = parser.GetInt(i, "Lv");
            data._Exp = parser.GetInt(i, "Exp");

            if(metaData.ContainsKey(data._Level))
            {
                Debug.LogError("ExpData Lv重复");
                return;
            }
            metaData[data._Level] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public int NeedExp(int level)
    {
        if (!metaData.ContainsKey(level))
            return null;

        return metaData[level]._Exp;
    }
}
