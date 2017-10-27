using UnityEngine;
using System.Collections.Generic;

public class Define {

    public const string PackageVersion = "1_0_0";

    public const string ASSET_EXT = ".ab";
    public const string TXT_EXT = ".bytes";

    public const int MALE_ID = 1;
    public const int FEMALE_ID = 2;

    public const string SCENE_LOGIN = "Login";
    public const string SCENE_MAIN = "main";
    public const string SCENE_CREATE = "createRole";
    public const string SCENE_BATTLE = "haidizhandou";
    public static string[] _BattleScenes = null;

    public const string ANIMATION_PLAYER_ACTION_RUN = "run";
    public const string ANIMATION_PLAYER_ACTION_WALK = "walk";
    public const string ANIMATION_PLAYER_ACTION_IDLE = "stand";
    public const string ANIMATION_PLAYER_ACTION_SHOW = "ruchang";
    public const string ANIMATION_PLAYER_ACTION_BEATTACK = "beattack";
    public const string ANIMATION_PLAYER_ACTION_DEAD = "dead";

    static Dictionary<string, object> globalValues;

    public static void LaunchUIBundle(string uiName)
    {
        uiName = uiName.ToLower();
        #if EDITOR_MODE
            FairyGUI.UIPackage.AddPackage(PathDefine.UI_ASSET_PATH + uiName);
        #else
            string path = string.Format("{0}{1}_desc", PathDefine.UI_ASSET_PATH, uiName);
            AssetBundle descBundle = AssetLoader.LoadAssetBundle(path);
            path = string.Format("{0}{1}_res", PathDefine.UI_ASSET_PATH, uiName);
            AssetBundle resBundle = AssetLoader.LoadAssetBundle(path);
            FairyGUI.UIPackage.AddPackage(descBundle, resBundle);
        #endif
    }

    public static void UnloadUIBundle(string uiName)
    {
        uiName = uiName.ToLower();
        #if EDITOR_MODE
        FairyGUI.UIPackage.RemovePackage(uiName, true);
        #endif
    }

    public static void Init()
    {
        globalValues = new Dictionary<string, object>();
        LuaManager.Call("global.lua", "RegGlobalValue");

        _BattleScenes = GetStr("BattleScenePool").Split(new char[]{','}, System.StringSplitOptions.RemoveEmptyEntries);
    }

    public static int GetInt(string key)
    {
        if (!globalValues.ContainsKey(key))
            return 0;

        return System.Convert.ToInt32(globalValues [key]);
    }

    public static float GetFloat(string key)
    {
        if (!globalValues.ContainsKey(key))
            return 0.0f;

        return float.Parse(globalValues [key].ToString());
    }

    public static string GetStr(string key)
    {
        if (!globalValues.ContainsKey(key))
            return string.Empty;

        return string.Format("{0}", globalValues [key]);
    }

    public static void Set(string key, object val)
    {
        if (!globalValues.ContainsKey(key))
        {
            globalValues.Add(key, val);
            return;
        }

        globalValues [key] = val;
    }

    public static string RandomBattleScene
    {
        get
        {
            if (_BattleScenes == null)
                return SCENE_BATTLE;
            return _BattleScenes [Random.Range(0, _BattleScenes.Length)];
        }
    }
}
