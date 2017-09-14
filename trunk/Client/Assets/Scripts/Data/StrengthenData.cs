using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class StrengthenData 
{
	public int _Id;
	public int _Level;
	public int _Hp;
	public int _Atk;
	public int _Def;
	public int _MagicAtk;
	public int _MagicDef;
	public int _Agile;
	public int _ItemId;
	public int _ItemNum;

	static Dictionary<int, List< StrengthenData>> metaData;
	static public void ParseData(string content, string fileName)
	{
        metaData = new Dictionary<int, List<StrengthenData>>();
		CSVParser parser = new CSVParser ();
		if(!parser.Parse (content))
		{
            Debug.LogError("StrengthenData 解析错误");
				return;
		}
		int recordCounter = parser.GetRecordCounter();
		StrengthenData data = null;
		for(int i=0; i < recordCounter; ++i)
		{
			data = new StrengthenData ();
			data._Id = parser.GetInt (i, "ID");
			data._Level = parser.GetInt (i, "Level");
            data._Hp = parser.GetInt(i, "Hp");
            data._Atk = parser.GetInt(i, "ATK");
            data._Def = parser.GetInt(i, "DEF");
            data._MagicAtk = parser.GetInt(i, "MAGIC_ATK");
            data._MagicDef = parser.GetInt(i, "MAGIC_DEF");
            data._Agile = parser.GetInt(i, "AGILE");
            data._ItemId = parser.GetInt(i, "ItemID");
            data._ItemNum = parser.GetInt(i, "ItemNum");

            if (metaData.ContainsKey(data._Id))
            {
                metaData[data._Id].Add(data);
            }
            else
            {
                metaData[data._Id] = new List<StrengthenData>();
                metaData[data._Id].Add(data);
            }
		}
		parser.Dispose ();
		parser = null;
	}

    static public StrengthenData GetData(int id,int level)
    {
        if (!metaData.ContainsKey(id))
            return null;
        return metaData[id][level-1];
    }
	
}

