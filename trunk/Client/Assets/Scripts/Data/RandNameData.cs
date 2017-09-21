using UnityEngine;
using System;
using System.Collections.Generic;

public class RandNameData {

    static List<string> _LastName;
    static List<string> _Name;

    static public void ParseData(string content, string fileName)
    {
        CSVParser parser = new CSVParser ();
        if(!parser.Parse (content))
        {
            Debug.LogError("RandNameData 解析错误");
            return;
        }

        _LastName = new List<string>();
        _Name = new List<string>();
        int recordCounter = parser.GetRecordCounter();
        RandNameData data = null;
        for(int i=0; i < recordCounter; ++i)
        {
            _LastName.Add(parser.GetString(i, "lastname"));
            _Name.Add(parser.GetString(i, "name"));
        }
        parser.Dispose ();
        parser = null;
    }

    static public string Rand()
    {
        string lastname = _LastName[UnityEngine.Random.Range(0, _LastName.Count)];
        string name = _Name[UnityEngine.Random.Range(0, _Name.Count)];
        return lastname + name;
    }
}
