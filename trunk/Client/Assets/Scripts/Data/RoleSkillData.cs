using UnityEngine;
using System;
using System.Collections.Generic;

public class RoleSkillData {

    public int _ID;
    public int _OpenLv;
    public int _SkillId;
    public int _Type;

    static public List<RoleSkillData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new List<RoleSkillData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("RoleSkillData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        RoleSkillData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new RoleSkillData ();
            data._ID = parser.GetInt(i, "ID");
            data._OpenLv = parser.GetInt(i, "OpenLv");
            data._SkillId = parser.GetInt(i, "SkillID");
            data._Type = parser.GetInt(i, "Type");
            metaData.Add(data);
        }
        parser.Dispose ();
        parser = null;
    }

    static public RoleSkillData GetDataBySkillID(int skillid)
    {
        for(int i=0; i < metaData.Count; ++i)
        {
            if (metaData [i]._SkillId == skillid)
                return metaData [i];
        }
        return null;
    }

    static public void SetData(int id, int skillid)
    {
        if (skillid == 0)
            return;
        
        for(int i=0; i < metaData.Count; ++i)
        {
            if (metaData [i]._ID == id)
                metaData [i]._SkillId = skillid;
        }
    }

    static public RoleSkillData GetData(int id)
    {
        for(int i=0; i < metaData.Count; ++i)
        {
            if (metaData [i]._ID == id)
                return metaData [i];
        }
        return null;
    }
}
