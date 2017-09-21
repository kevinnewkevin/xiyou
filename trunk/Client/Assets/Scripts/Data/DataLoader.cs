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

    public static bool _IsLoading;

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
        #if EDITOR_MODE
            _ParseTables.Enqueue(new TableInfo(tableName, callback));
        #else
            int idx = tableName.LastIndexOf("/");
            if(idx != -1)
                tableName = tableName.Substring(idx + 1);
            _ParseTables.Enqueue(new TableInfo(tableName, callback));
        #endif
    }

    static public void Init()
    {
        RegistTables("DisPlay", DisplayData.ParseData);
        RegistTables("Npc", NpcData.ParseData);
        RegistTables("Scene", SceneData.ParseData);
        RegistTables("skill", SkillData.ParseData);
        RegistTables("entity", EntityData.ParseData);
		RegistTables("HeroStroy", HeroStroyData.ParseData);
		RegistTables("Checkpoint", CheckpointData.ParseData);
		RegistTables("buff", BuffData.ParseData);
        RegistTables("Item", ItemData.ParseData);
        RegistTables("animation", AnimationData.ParseData);
        RegistTables("animation_actor", ActorData.ParseData);
        RegistTables("Battle", BattleData.ParseData);
        RegistTables("Talk", TalkData.ParseData);
        RegistTables("Strengthen", StrengthenData.ParseData);
        RegistTables("Exp", ExpData.ParseData);
        RegistTables("Weather", WeatherData.ParseData);
        RegistTables("Drop", DropData.ParseData);
        RegistTables("RoleSkill", RoleSkillData.ParseData);
        RegistTables("RoleSkillUpdate", RoleSkillUpdateData.ParseData);
    }

    static public void BeginLoad()
    {
        _IsLoading = true;
        UIManager.SetDirty("denglu");
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
                        string filename = _LoadingTable._Name;
                        int idx = filename.LastIndexOf("/");
                        if(idx != -1)
                            filename = filename.Substring(idx + 1);
                        _LoadingTable.Excute(_Request.assetBundle.LoadAsset(filename).ToString());
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
        UIManager.SetDirty("denglu");
    }
}
