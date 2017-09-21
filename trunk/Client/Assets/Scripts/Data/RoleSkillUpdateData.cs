using UnityEngine;
using System;
using System.Collections.Generic;

public class RoleSkillUpdateData {

    public int _SkillId;
    public int _NeedItem;
    public int _NeedNum;
    public int _NeedGold;
    public int _NextId;

    static Dictionary<int, RoleSkillUpdateData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, RoleSkillUpdateData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("RoleSkillUpdateData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        RoleSkillUpdateData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new RoleSkillUpdateData ();
            data._SkillId = parser.GetInt(i, "SkillID");
            data._NeedItem = parser.GetInt(i, "NeedItem");
            data._NeedNum = parser.GetInt(i, "NeedNum");
            data._NeedGold = parser.GetInt(i, "NeedMoney");
            data._NextId = parser.GetInt(i, "NextID");

            if(metaData.ContainsKey(data._SkillId))
            {
                Debug.LogError("RoleSkillUpdateData Lv重复");
                return;
            }
            metaData[data._SkillId] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public RoleSkillUpdateData GetData(int skillid)
    {
        if (!metaData.ContainsKey(skillid))
            return null;

        return metaData [skillid];
    }
}
