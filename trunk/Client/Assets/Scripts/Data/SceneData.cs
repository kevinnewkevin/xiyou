using UnityEngine;
using System;
using System.Collections.Generic;

public class SceneData {

	public int _Id;
    public int[] _NpcId;

    static Dictionary<int, SceneData> metaData;
	
    static public void ParseData(string content, string fileName)
	{
        metaData = new Dictionary<int, SceneData> ();

		CSVParser parser = new CSVParser ();
		if(!parser.Parse (content))
		{
            Debug.LogError("SceneData 解析错误");
			return;
		}

		int recordCounter = parser.GetRecordCounter();
        SceneData data = null;
		for(int i=0; i < recordCounter; ++i)
		{
            data = new SceneData ();
            data._Id = parser.GetInt (i, "ID");
            string[] npcidStr = parser.GetString (i, "NPC").Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._NpcId = new int[npcidStr.Length];
            for(int j=0; j < npcidStr.Length; ++j)
            {
                data._NpcId [j] = int.Parse(npcidStr [j]);
            }
			if(metaData.ContainsKey(data._Id))
			{
                Debug.LogError("SceneData ID重复");
				return;
			}
			metaData[data._Id] = data;
		}
		parser.Dispose ();
		parser = null;
	}

    static public SceneData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
