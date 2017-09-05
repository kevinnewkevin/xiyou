using UnityEngine;
using System;
using System.Collections.Generic;

public class AnimationData {

    public int _Id;
    public string _ClipName;
    public int[] _Actors;
    public int[] _Talks;

    static Dictionary<int, AnimationData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, AnimationData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("AnimationData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        AnimationData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new AnimationData ();
            data._Id = parser.GetInt (i, "id");
            data._ClipName = parser.GetString(i, "clipName");
            string[] actorstr = parser.GetString(i, "actors").Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._Actors = new int[actorstr.Length];
            for(int j=0; j < actorstr.Length; ++j)
            {
                data._Actors[j] = int.Parse(actorstr[j]);
            }
            string[] talkstr = parser.GetString(i, "talks").Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._Talks = new int[talkstr.Length];
            for(int j=0; j < talkstr.Length; ++j)
            {
                data._Talks[j] = int.Parse(talkstr[j]);
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

    static public AnimationData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
