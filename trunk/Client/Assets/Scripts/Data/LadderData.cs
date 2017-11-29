using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class LadderData 
{
	public int _Id;
	public int _ScoreL;
	public int _ScoreH;
	public string _Name;
	public int _LadderDrop;
	public int _WinDrop;
	public int _LoseDop;
	public static Dictionary<int, LadderData> metaData;

	static public void ParseData(string content, string fileName)
	{
		metaData = new Dictionary<int, LadderData> ();
		CSVParser parser = new CSVParser ();
		if(!parser.Parse (content))
		{
			Debug.LogError("LadderData 解析错误");
			return;
		}
		int recordCounter = parser.GetRecordCounter();
		LadderData data = null;
		for(int i=0; i < recordCounter; ++i)
		{
			data = new LadderData ();
			data._Id = parser.GetInt (i, "ID");
			data._ScoreL = parser.GetInt (i, "ScoreL");
			data._ScoreH = parser.GetInt (i, "ScoreH");
			data._Name = parser.GetString (i, "Name");
			data._LadderDrop = parser.GetInt (i, "LadderDrop");
			data._WinDrop = parser.GetInt (i, "WinDrop");
			data._LoseDop = parser.GetInt(i, "LoseDop");

			if(metaData.ContainsKey(data._Id))
			{
				Debug.LogError("LadderData ID重复");
				return;
			}
			metaData[data._Id] = data;
		}
		parser.Dispose ();
		parser = null;
	}

	public static LadderData GetData(int id)
	{
		if (!metaData.ContainsKey(id))
			return null;

		return metaData[id];
	}
}

