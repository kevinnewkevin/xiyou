using UnityEngine;
using System.Collections.Generic;

public class Define {

    public const string PackageVersion = "1_0_0";

    public const string ASSET_EXT = ".ab";
    public const string TXT_EXT = ".bytes";

    public const string SCENE_MAIN = "main";
    public const string SCENE_BATTLE = "haidizhandou";

    public const string ANIMATION_PLAYER_ACTION_RUN = "run";
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
            string path = string.Format("{0}/{1}/{2}{3}{4}{5}", Application.streamingAssetsPath, Define.PackageVersion, PathDefine.UI_ASSET_PATH, uiName, "_desc", Define.ASSET_EXT);
            AssetBundle descBundle = AssetBundle.LoadFromFile(path);
            path = string.Format("{0}/{1}/{2}{3}{4}{5}", Application.streamingAssetsPath, Define.PackageVersion, PathDefine.UI_ASSET_PATH, uiName, "_res", Define.ASSET_EXT);
            AssetBundle resBundle = AssetBundle.LoadFromFile(path);
            FairyGUI.UIPackage.AddPackage(descBundle, resBundle);
        #endif
    }

    public static void Init()
    {
        globalValues = new Dictionary<string, object>();
        LuaManager.Call("global.lua", "RegGlobalValue");
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
}
