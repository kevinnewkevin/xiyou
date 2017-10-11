﻿using UnityEngine;
using System;
using System.Collections.Generic;

public class EntityData {

    public int _UnitId;
    public int _DisplayId;
    public int _Cost;
    public string _Name;
    public int[] _Skills;
    public string _Desc;

    static Dictionary<int, EntityData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, EntityData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("EntityData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        EntityData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new EntityData ();
            data._UnitId = parser.GetInt (i, "UnitId");
            data._DisplayId = parser.GetInt(i, "DisplayId");
            data._Cost = parser.GetInt(i, "Cost");
            data._Name = parser.GetString(i, "Name");
            data._Desc = parser.GetString(i, "Desc");
            data._Skills = new int[4];
            for(int j=0; j < 4; ++j)
            {
                data._Skills[j] = parser.GetInt(i, "Skill" + (j + 1).ToString());
            }

            if(metaData.ContainsKey(data._UnitId))
            {
                Debug.LogError("EntityData ID重复");
                return;
            }
            metaData[data._UnitId] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public EntityData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
