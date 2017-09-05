using UnityEngine;
using System;
using System.Collections.Generic;

public class BattleData {

    public int _Id;
    public int[] _Animations;

    static Dictionary<int, BattleData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, BattleData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("BattleData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        BattleData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new BattleData ();
            data._Id = parser.GetInt (i, "ID");
            string[] animstr = parser.GetString (i, "Animation").Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._Animations = new int[animstr.Length];
            for(int j=0; j < animstr.Length; ++j)
            {
                data._Animations [j] = int.Parse(animstr[j]);
            }
            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("BattleData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public BattleData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
