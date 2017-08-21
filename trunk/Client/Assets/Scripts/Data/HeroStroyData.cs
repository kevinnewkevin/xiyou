using UnityEngine;
using System;
using System.Collections.Generic;

public class HeroStroyData 
{
	public int Id_;
	public int Type_;
	public int Level_;
	public int Star_;
	public string Name_;
	public string Icon_;
	public string Desc_;
	
	static public Dictionary<int, HeroStroyData> metaData;
	
	static public void ParseData(string content, string fileName)
	{
		metaData = new Dictionary<int, HeroStroyData> ();
		
		CSVParser parser = new CSVParser ();
		if(!parser.Parse (content))
		{
			Debug.LogError("EntityData 解析错误");       
			return;
		} 
		
		int recordCounter = parser.GetRecordCounter();
		HeroStroyData data = null;
		for(int i=0; i < recordCounter; ++i)
		{
			data = new HeroStroyData ();
			data.Id_ = parser.GetInt (i, "ID");
			data.Type_ = parser.GetInt (i, "Type");
			data.Level_ = parser.GetInt (i, "Level");
			data.Star_ = parser.GetInt (i, "Star");
			data.Name_ = parser.GetString (i, "Name");
			data.Icon_ = parser.GetString (i, "Icon");
  			data.Desc_ = parser.GetString (i, "Desc"); 
			
			if(metaData.ContainsKey(data.Id_))
			{
				Debug.LogError("EntityData ID重复");
				return;
			}
			metaData[data.Id_] = data;
		}
		parser.Dispose ();
		parser = null;
	}
	
	static public HeroStroyData GetData(int id)
	{
		if (!metaData.ContainsKey(id))
			return null;
		
				return metaData[id];
	}
}
