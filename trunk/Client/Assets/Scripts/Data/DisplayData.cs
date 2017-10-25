using UnityEngine;
using System;
using System.Collections.Generic;

public class DisplayData {

    public int _Id;
    public string _AssetPath;
    public string _Quality;
	public string _Race;
    public string _AssetPathDetail;
    public string _HeadIcon;
    public string _CardIcon;
    public float _Distance;
    public float _HeadBarHeight;
    public float _BattleSkillScale;
    public float _BattleSkillHeight;


    static Dictionary<int, DisplayData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, DisplayData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("DisplayData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        DisplayData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new DisplayData ();
            data._Id = parser.GetInt (i, "ID");
            data._AssetPath = parser.GetString(i, "AssetsPath");
            data._Quality = parser.GetString(i, "Quality");
			data._Race = parser.GetString(i, "Race");
            data._AssetPathDetail = parser.GetString(i, "AssetsPath2");
            data._HeadIcon = parser.GetString(i, "HeadIcon");
            data._CardIcon = parser.GetString(i, "CardIcon");
            data._Distance = parser.GetFloat(i, "Distance");
            data._HeadBarHeight = parser.GetFloat(i, "HeadBarHeight");
            data._BattleSkillScale = parser.GetFloat(i, "ScaleInBattleSkill");
            data._BattleSkillHeight = parser.GetFloat(i, "HeightInBattleSkill");
					
            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("DisplayData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    static public DisplayData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
