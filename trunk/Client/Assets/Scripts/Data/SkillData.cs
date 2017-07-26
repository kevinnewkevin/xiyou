using UnityEngine;
using System;
using System.Collections.Generic;

public class SkillData {

    public int _Id;
    public bool _IsMelee;

    static Dictionary<int, SkillData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, SkillData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("SkillData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        SkillData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new SkillData ();
            data._Id = parser.GetInt (i, "SkillId");
            data._IsMelee = parser.GetBool (i, "IsMelee");

            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("SkillData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public SkillData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
