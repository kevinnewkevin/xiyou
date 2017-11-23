using UnityEngine;
using System;
using System.Collections.Generic;

public class DrawData {

    public int _ID;
    public int[] _HeroStoryIDs;
    public int _ItemId;
    public int _ItemNum;
    public static Dictionary<int, DrawData > metaData;
    public static List<int> _IDList;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, DrawData > ();
        _IDList = new List<int>();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {  
            Debug.LogError("DrawData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter(); 
        DrawData  data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new DrawData  ();
            data._ID = parser.GetInt (i, "ID");
            data._ItemId = parser.GetInt (i, "ItemID");
            data._ItemNum = parser.GetInt (i, "ItemNumber");
            string hsidstr = parser.GetString (i, "HeroStroyID");
            if(string.IsNullOrEmpty(hsidstr))
                Debug.LogError("DrawData 存在废库 ID : " + data._ID);
            string[] hsidstrArr = hsidstr.Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._HeroStoryIDs = new int[hsidstrArr.Length];
            for(int j=0; j < hsidstrArr.Length; ++j)
            {
                data._HeroStoryIDs[j] = int.Parse(hsidstrArr[j]);
            }
            if(metaData.ContainsKey(data._ID))
            {
                Debug.LogError("DrawData ID重复");
                return;
            }
            metaData[data._ID] = data;
            _IDList.Add(data._ID);
        }
        _IDList.Sort();
        parser.Dispose ();
        parser = null;
    }

    public static DrawData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }

    public static DrawData GetNextGarageData(List<int> ids)
    {
        if (ids == null || ids.Count == 0)
            return null;

        DrawData dd = null;
        for(int i=0; i < _IDList.Count; ++i)
        {
            dd = GetData(_IDList[i]);
            for(int j=0; j < dd._HeroStoryIDs.Length; ++j)
            {
                if (!ids.Contains(dd._HeroStoryIDs [j]))
                    return dd;
            }
        }
        return null;
    }
}
