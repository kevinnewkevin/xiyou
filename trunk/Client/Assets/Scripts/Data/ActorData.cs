using UnityEngine;
using System;
using System.Collections.Generic;

public class ActorData {

    public int _Id;
    public string _Asset;
    public KeyValuePair<float, string>[] _Actions;

    static Dictionary<int, ActorData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, ActorData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("ActorData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        ActorData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new ActorData ();
            data._Id = parser.GetInt (i, "id");
            data._Asset = parser.GetString(i, "assets");
            string[] actionstr = parser.GetString(i, "action").Split(new char[]{'|'}, StringSplitOptions.RemoveEmptyEntries);
            string[] actionunitstr;
            data._Actions = new KeyValuePair<float, string>[actionstr.Length];
            for(int j=0; j < actionstr.Length; ++j)
            {
                actionunitstr = actionstr[j].Split(new char[]{':'}, StringSplitOptions.RemoveEmptyEntries);
                data._Actions [j] = new KeyValuePair<float, string>(float.Parse(actionunitstr[0]), actionunitstr[1]);
            }
            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("ActorData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public ActorData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
