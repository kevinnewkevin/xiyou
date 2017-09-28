using UnityEngine;
using System;
using System.Collections.Generic;

public class ShopData {
    
    public int _ShopId;
	public string _Name;
	public string _ShopType;
	public int _CardId;
	public int _Num;
	public string _PayType;
	public int _Price;
	public int _copper;
	public static Dictionary<int, ShopData > metaData;

    static public void ParseData(string content, string fileName)
    {
		metaData = new Dictionary<int, ShopData > ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {  
			Debug.LogError("ShopData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter(); 
		ShopData  data = null;
        for(int i=0; i < recordCounter; ++i)
        {
			data = new ShopData  ();
			data._ShopId = parser.GetInt (i, "ID");
			data._ShopType = parser.GetString (i, "ShopType");
			data._Name = parser.GetString (i, "Name");
			data._Num = parser.GetInt (i, "Num");
			data._CardId = parser.GetInt (i, "CardcloseID");
			data._PayType = parser.GetString (i, "ShopPayType");
			data._Price = parser.GetInt(i, "Price");
			data._copper = parser.GetInt(i, "COPPER");
			if(metaData.ContainsKey(data._ShopId))
            {
				Debug.LogError("ShopData ID重复");
                return;
            }
			metaData[data._ShopId] = data;
        }
        parser.Dispose ();
        parser = null;
    }

	public static ShopData GetData(int id)
    {
        if (!metaData.ContainsKey(id))
            return null;

        return metaData[id];
    }
}
