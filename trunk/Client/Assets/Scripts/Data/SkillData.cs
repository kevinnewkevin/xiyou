using UnityEngine;
using System;
using System.Collections.Generic;

public class SkillData {

    public int _Id;
    public bool _IsMelee;
    public string _CastAnim;
    public string _CastEffect;
    public string _BeattackAnim;
    public string _BeattackEffect;
    public float _BeattackTime;
    public float _EmitNumTime;
    public float _TotalTime;

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
            data._CastAnim = parser.GetString (i, "CastAnim");
            data._CastEffect = parser.GetString (i, "CastEffect");
            data._BeattackAnim = parser.GetString (i, "BeattackAnim");
            data._BeattackEffect = parser.GetString (i, "BeattackEffect");
            data._BeattackTime = parser.GetFloat (i, "BeattackTime");
            data._EmitNumTime = parser.GetFloat (i, "EmitNumTime");
            data._TotalTime = parser.GetFloat(i, "TotalTime");

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
