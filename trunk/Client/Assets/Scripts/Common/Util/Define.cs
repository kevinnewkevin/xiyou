using UnityEngine;

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

    public static void LaunchUIBundle(string uiName)
    {
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
}
