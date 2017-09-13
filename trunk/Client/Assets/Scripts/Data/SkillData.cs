using UnityEngine;
using System;
using System.Collections.Generic;

public class SkillData {

    public enum MotionType
    {
        MT_None,
        MT_Self,
        MT_Target,
        MT_Fly,
        MT_Sec,
    }

    public int _Id;
    public bool _IsMelee;
    public string _Name;
    public string _CastAnim;
    public string _CastEffect;
    public string _SkillEffect;
    public string _AttackAnim;
    public string _BeattackEffect;
    public float _CastTime;
    public float[] _BeattackTime;
    public float[] _EmitNumTime;
    public float _TotalTime;
    public MotionType _Motion;
    public bool _Single;
    public string _Camera;
    public string _Icon;

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
            data._Name = parser.GetString (i, "Name");
            data._IsMelee = parser.GetBool (i, "IsMelee");
            data._CastAnim = parser.GetString (i, "CastAnim");
            data._CastEffect = parser.GetString (i, "CastEffect");
            data._BeattackEffect = parser.GetString (i, "BeattackEffect");
            data._SkillEffect = parser.GetString (i, "SkillEffect");
            data._AttackAnim = parser.GetString(i, "AttackAnim");
            data._CastTime = parser.GetFloat(i, "CastTime");
            string beatkTimes = parser.GetString (i, "BeattackTime");
            string[] beatkTimeAry = beatkTimes.Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            int j;
            data._BeattackTime = new float[beatkTimeAry.Length];
            for(j=0; j < beatkTimeAry.Length; ++j)
            {
                data._BeattackTime[j] = float.Parse(beatkTimeAry[j]);
            }
            string emitTimes = parser.GetString (i, "EmitNumTime");
            string[] emitTimeAry = beatkTimes.Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._EmitNumTime = new float[emitTimeAry.Length];
            for(j=0; j < emitTimeAry.Length; ++j)
            {
                data._EmitNumTime[j] = float.Parse(emitTimeAry[j]);
            }
            data._TotalTime = parser.GetFloat(i, "TotalTime");
            data._Motion = (MotionType)Enum.Parse(typeof(MotionType), parser.GetString(i, "MotionType"));
            data._Single = parser.GetInt(i, "SingleSkill") == 1;
            data._Icon = parser.GetString(i, "ICON");
            data._Camera = parser.GetString(i, "Camera");

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
