using UnityEngine;
using System;
using System.Collections.Generic;

public class HeroStroyData 
{
	public int Id_;
	public int Type_;
	public int Level_;
    public int[] Rewards_;
    public int[] Star_;
	public string Name_;
	public string Icon_;
	public string Desc_;
	public int EntityID_;
	static public List<HeroStroyData> easyList = new List<HeroStroyData>(); 
	static public Dictionary<int, HeroStroyData> metaData;
	
	static public void ParseData(string content, string fileName)
	{
		metaData = new Dictionary<int, HeroStroyData> ();
		
		CSVParser parser = new CSVParser ();
		if(!parser.Parse (content))
		{
			Debug.LogError("HeroStroyData 解析错误");       
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

            string[] starStr = parser.GetString(i, "Star").Split(new char[] { ';' }, StringSplitOptions.RemoveEmptyEntries);
            data.Star_ = new int[starStr.Length];
            for (int j = 0; j < starStr.Length; ++j)
            {
                data.Star_[j] = int.Parse(starStr[j]);
            }

            string[] rewardStr = parser.GetString(i, "Reward").Split(new char[] { ';' }, StringSplitOptions.RemoveEmptyEntries);
            data.Rewards_ = new int[starStr.Length];
            for (int r = 0; r < rewardStr.Length; ++r)
            {
                data.Rewards_[r] = int.Parse(rewardStr[r]);
            }

			data.Name_ = parser.GetString (i, "Name");
			data.Icon_ = parser.GetString (i, "Icon");
  			data.Desc_ = parser.GetString (i, "Desc"); 
			data.EntityID_ = parser.GetInt(i, "EntityID");
			if(data.Type_ == 1)
			{
				easyList.Add (data); 
			}
			if(metaData.ContainsKey(data.Id_))
			{
				Debug.LogError("HeroStroyData ID重复");
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

	static public int GetEasyListNum()
	{
		return easyList.Count;
	}

}
