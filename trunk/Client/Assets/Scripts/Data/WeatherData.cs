using UnityEngine;
using System;
using System.Collections.Generic;

public class WeatherData {

    public int _Id;
    public string _Asset;

    public static Dictionary<int, WeatherData> metaData;

    static public void ParseData(string content, string fileName)
    {
        metaData = new Dictionary<int, WeatherData> ();

        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("WeatherData 解析错误");
            return;
        }

        int recordCounter = parser.GetRecordCounter();
        WeatherData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            data = new WeatherData ();
            data._Id = parser.GetInt (i, "Hour");
            data._Asset = parser.GetInt (i, "Asset");

            if(metaData.ContainsKey(data._Id))
            {
                Debug.LogError("WeatherData ID重复");
                return;
            }
            metaData[data._Id] = data;
        }
        parser.Dispose ();
        parser = null;
    }

    public static string GetWeather(int hour)
    {
        if (!metaData.ContainsKey(hour))
            return null;

        return metaData[hour]._Asset;
    }
}
