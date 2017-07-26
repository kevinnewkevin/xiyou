using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class DataLoader {

#if EDITOR_MODE
    static ResourceRequest _Request;
#else
    static AssetBundleCreateRequest _Request;
#endif
    static TableInfo _LoadingTable;

    static bool _IsLoading;

    class TableInfo
    {
        public string _Name;
        public ParseHandler _Callback;
        public TableInfo(string name, ParseHandler callback)
        {
            _Name = name;
            _Callback = callback;
        }

        public void Excute(string content)
        {
            if (_Callback != null)
                _Callback(content, _Name);
            _Callback = null;
        }
    }

    public delegate void ParseHandler(string content, string fileName);
    static Queue<TableInfo> _ParseTables = new Queue<TableInfo>();

    static public void RegistTables(string tableName, ParseHandler callback)
    {
        _ParseTables.Enqueue(new TableInfo(tableName, callback));
    }

    static public void Init()
    {
        RegistTables("DisPlay", DisplayData.ParseData);
        RegistTables("Npc", NpcData.ParseData);
        RegistTables("Scene", SceneData.ParseData);
        RegistTables("skill", SkillData.ParseData);
        RegistTables("tables/entity", EntityData.ParseData);
    }

    static public void BeginLoad()
    {
        _IsLoading = true;
    }

    static public void Update()
    {
        if (_IsLoading == false)
            return;
        
        if (_Request != null)
        {
            if (_Request.isDone)
            {
                if (_LoadingTable != null)
                {
                    #if EDITOR_MODE
                        _LoadingTable.Excute(_Request.asset.ToString());
                    #else
                        _LoadingTable.Excute(_Request.assetBundle.LoadAsset(_LoadingTable._Name).ToString());
                    #endif
                }
                _Request = null;
            }
        }
        else
        {
            if (_ParseTables.Count > 0)
            {
                _LoadingTable = _ParseTables.Dequeue();
                _Request = AssetLoader.LoadAssetAsync(PathDefine.TABLE_ASSET_PATH + _LoadingTable._Name);
            }
            else
                Dispose();
        }
    }

    static void Dispose()
    {
        _Request = null;
        _LoadingTable = null;
        _IsLoading = false;
    }
}
