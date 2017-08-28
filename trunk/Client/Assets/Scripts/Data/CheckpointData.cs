using UnityEngine;
using System;
using System.Collections.Generic;
public class CheckpointData 
{
	public int _HerID;
	public int _ID;
	public int _Star1Need;
	public int _Star2Need;
	public int _Star3Need;
	public int _battleID;
	public int _DropID;
	public int _Main;
	public string _Name;
	static public Dictionary<int, List<CheckpointData>> metaData;

	static public void ParseData(string content, string fileName)
	{
			metaData = new Dictionary<int, List<CheckpointData>> ();

			CSVParser parser = new CSVParser ();
			if(!parser.Parse (content))
			{
					Debug.LogError("EntityData 解析错误");       
					return;
			} 

			int recordCounter = parser.GetRecordCounter();
				CheckpointData data = null;
			for(int i=0; i < recordCounter; ++i)
			{
				data = new CheckpointData ();
				data._HerID = parser.GetInt (i, "HeroID");
				data._ID = parser.GetInt (i, "ID");
				data._Star1Need = parser.GetInt (i, "Star1Need");
				data._Star2Need = parser.GetInt (i, "Star2Need");
				data._Star3Need = parser.GetInt (i, "Star3Need");
				data._battleID = parser.GetInt (i, "BattleID"); 
				data._DropID = parser.GetInt (i, "DropID"); 
				data._Main = parser.GetInt (i, "Main"); 
				data._Name = parser.GetString(i, "Name");
				if (!metaData.ContainsKey (data._HerID)) 
				{
						metaData [data._HerID] = new List<CheckpointData> ();
				}
				metaData [data._HerID].Add (data);
			}
			parser.Dispose ();
			parser = null;
	}

		static public List<CheckpointData> GetData(int id)
	{
			if (!metaData.ContainsKey(id))
					return null;

			return metaData[id];
	}
}
