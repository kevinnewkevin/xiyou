using UnityEngine;
using System.Collections;
using System.Collections.Generic;
public class CardcloseData
{
		public int _ID;
		public int _GreenitemID;
		public int _Greennum;
		public int _BlueitemID;
		public int _Bluenum;
		public int _PurpleitemID;
		public int _Purplenum;
		public int _OrangetemID;
		public int _Orangenum;
		public static Dictionary<int, CardcloseData > metaData;

		static public void ParseData(string content, string fileName)
		{
				metaData = new Dictionary<int, CardcloseData > ();

				CSVParser parser = new CSVParser ();
				if(!parser.Parse (content))
				{  
						Debug.LogError("CardcloseData 解析错误");
						return;
				}

				int recordCounter = parser.GetRecordCounter(); 
				CardcloseData  data = null;
				for(int i=0; i < recordCounter; ++i)
				{
						data = new CardcloseData  ();
						data._ID = parser.GetInt (i, "ID");
						data._GreenitemID= parser.GetInt (i, "GreenitemID");
						data._Greennum = parser.GetInt (i, "Greennum");
						data._BlueitemID = parser.GetInt (i, "BlueitemID");
						data._Bluenum = parser.GetInt (i, "BlueitemID");
						data._PurpleitemID = parser.GetInt(i, "PurpleitemID");
						data._Purplenum = parser.GetInt(i, "Purplenum");
						data._OrangetemID = parser.GetInt(i, "OrangeitemID");
						data._Orangenum = parser.GetInt(i, "Orangenum");
						if(metaData.ContainsKey(data._ID))
						{
								Debug.LogError("CardcloseData ID重复");
								return;
						}
						metaData[data._ID] = data;
				}
				parser.Dispose ();
				parser = null;
		}

		public static CardcloseData GetData(int id)
		{
				if (!metaData.ContainsKey(id))
						return null;

				return metaData[id];
		}	
}

