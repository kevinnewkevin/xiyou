using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class AssetLoader {

    static AssetBundleManifest _Manifest;
    //读取总Common表
    static public void InitCommonList()
    {
        AssetBundle bundle = AssetBundle.LoadFromFile(Application.streamingAssetsPath +  "/" + Define.PackageVersion + "/" + Define.PackageVersion);
        _Manifest = (AssetBundleManifest)bundle.LoadAsset("AssetBundleManifest");
        bundle.Unload(false);
    }

	static public GameObject LoadAsset(string path)
    {
#if EDITOR_MODE
        Object obj = Resources.Load(path);
        if (obj == null)
            return null;
        return GameObject.Instantiate(obj) as GameObject;
#else
        if(_Manifest == null)
        InitCommonList();

        string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
        string assetPath;
        for(int i=0; i < dep.Length; ++i)
        {
            assetPath = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/" + dep[i];
            AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
        }
        assetPath = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
        AssetBundle ab = AssetBundle.LoadFromFile(assetPath);
        AssetCounter.AddRef(assetPath, ab);
        string assetName = path.Substring(path.LastIndexOf("/") + 1);
        Object o = ab.LoadAsset(assetName);
        return GameObject.Instantiate(o) as GameObject;
#endif
    }

#if EDITOR_MODE
    static public ResourceRequest LoadAssetAsync(string path)
    {
        return Resources.LoadAsync(path);
    }
#else
    static public AssetBundleCreateRequest LoadAssetAsync(string path)
    {
        if(_Manifest == null)
            InitCommonList();

        string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
        string assetPath;
        for(int i=0; i < dep.Length; ++i)
        {
            assetPath = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/" + dep[i];
            AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
        }
        assetPath = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
        return AssetBundle.LoadFromFileAsync(assetPath);
    }
#endif

    static void AssetLoadCallback(AssetBundleRequest request)
    {
        
    }

    static public void UnloadAsset(string path)
    {
#if !EDITOR_MODE
        if(_Manifest == null)
            InitCommonList();

        string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
        string assetPath;
        for(int i=0; i < dep.Length; ++i)
        {
            assetPath = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/" + dep[i];
            AssetCounter.DelRef(assetPath);
        }
        assetPath = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
        AssetCounter.DelRef(assetPath);
#endif
    }
}
