using System;
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

    public static int _MaxReportTips;

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
            if(descBundle == null || resBundle == null)
                return;
        
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
        _MaxReportTips = GetInt("MaxReportTips");
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
            return _BattleScenes [UnityEngine.Random.Range(0, _BattleScenes.Length)];
        }
    }

    public static void RegNoPopUI(string uiName)
    {
        if (!UIManager._NoPopIgnoreUI.Contains(uiName))
            UIManager._NoPopIgnoreUI.Add(uiName);
    }

    public static void CheckFileExcsit(string path)
    {
				return;
        #if EDITOR_MODE
        if (string.IsNullOrEmpty(path))
            return;
        
        UnityEngine.Object obj = Resources.Load<UnityEngine.Object>(path);
       // if(obj == null)
         //  Debug.LogError(path + " not Excsit");
        #endif
    }

    //加到类的定义部分
    private static string[] cstr={"零","一","二","三","四", "五", "六","七","八","九"};
    private  static string[] wstr={"","","十","百","千","万","十","百","千","亿","十","百","千"};
    //数字必须在12位整数以内的字符串
    //调用方式如：Label1.Text=ConvertInt("数字字符串");
    public static string ConvertInt(string str)
    {
        int len=str.Length;
        int i;
        string tmpstr,rstr;
        rstr="";
        for(i=1;i<=len;i++)
        {
            tmpstr=str.Substring(len-i,1);
            rstr=string.Concat(cstr[Int32.Parse(tmpstr)]+wstr[i],rstr);
        }
        rstr=rstr.Replace("十零","十");
        rstr=rstr.Replace("零十","零");
        rstr=rstr.Replace("零百","零");
        rstr=rstr.Replace("零千","零");
        rstr=rstr.Replace("零万","万");
        for(i=1;i<=6;i++)
            rstr=rstr.Replace("零零","零");
        rstr=rstr.Replace("零万","零");
        rstr=rstr.Replace("零亿","億");
        rstr=rstr.Replace("零零","零"); 
        return rstr;
    }
}