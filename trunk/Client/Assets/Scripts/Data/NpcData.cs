using UnityEngine;
using System;
using System.Collections.Generic;

public class NpcData {

	public int _Id;
    public int _Display;
    public string _OpenUI;
    public int[] _QuestID;
    public Vector3 _Position;

    static Dictionary<int, NpcData> metaData;
	
    static public void ParseData(string content, string fileName)
	{
        metaData = new Dictionary<int, NpcData> ();

		CSVParser parser = new CSVParser ();
		if(!parser.Parse (content))
		{
            Debug.LogError("NpcData 解析错误");
			return;
		}

		int recordCounter = parser.GetRecordCounter();
        NpcData data = null;
		for(int i=0; i < recordCounter; ++i)
		{
            data = new NpcData ();
            data._Id = parser.GetInt (i, "ID");
            data._Display = parser.GetInt (i, "DisPlayID");
            data._OpenUI = parser.GetString (i, "UIPath");
            string questes = parser.GetString (i, "QuestID");
            string[] qusetesArr = questes.Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            data._QuestID = new int[qusetesArr.Length];
            for (int j = 0; j < qusetesArr.Length; ++j)
            {
                data._QuestID[j] = int.Parse(qusetesArr[j]);
            }
            string[] posStr = parser.GetString (i, "Position").Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
            if (posStr.Length != 3)
                Debug.Log("Error Position in table: Npc.csv");
            else
                data._Position = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));

			if(metaData.ContainsKey(data._Id))
			{
                Debug.LogError("NpcData ID重复");
				return;
			}
			metaData[data._Id] = data;
		}
		parser.Dispose ();
		parser = null;
	}

    static public NpcData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
