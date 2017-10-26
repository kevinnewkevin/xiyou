using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.IO;

public class AssetLoader {

    static AssetBundleManifest _Manifest;
    //读取总Common表
    static public void InitCommonList()
    {
        string tpath = Application.persistentDataPath +  "/" + Define.PackageVersion + "/" + Define.PackageVersion;
        SmartPath(ref tpath);
        AssetBundle bundle = AssetBundle.LoadFromFile(tpath);
        _Manifest = (AssetBundleManifest)bundle.LoadAsset("AssetBundleManifest");
        bundle.Unload(false);
    }

    static public void SmartPath(ref string path)
    {
        if (!File.Exists(path))
            path = path.Replace(Application.persistentDataPath, Application.streamingAssetsPath);
    }

	static public GameObject LoadAsset(string path)
    {
        if (string.IsNullOrEmpty(path))
            return null;
        path = path.ToLower();
#if EDITOR_MODE
        Object obj = Resources.Load(path);
        if (obj == null)
        {
            Debug.LogError("资源: " + path + " 未找到!");
            LuaManager.Call("global.lua", "ErrorMessage", "资源: " + path + " 未找到!");
            return null;
        }
        return GameObject.Instantiate(obj) as GameObject;
#else
        if(_Manifest == null)
        InitCommonList();

        try
        {
            string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
            string assetPath;
            for(int i=0; i < dep.Length; ++i)
            {
                assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + dep[i];
                SmartPath(ref assetPath);
                if(!AssetCounter.Excist(assetPath))
                    AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
                else
                    AssetCounter.GetBundle(assetPath);
            }
            assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
            SmartPath(ref assetPath);
            AssetBundle ab = null;
            if(!AssetCounter.Excist(assetPath))
            {
                ab = AssetBundle.LoadFromFile(assetPath);
                AssetCounter.AddRef(assetPath, ab);
            }
            else
                ab = AssetCounter.GetBundle(assetPath);

            string assetName = path.Substring(path.LastIndexOf("/") + 1);
            Object o = ab.LoadAsset(assetName);
            return GameObject.Instantiate(o) as GameObject;
        }
        catch(System.Exception e)
        {
            Debug.LogWarning("AssetPath: " + path + " is not excist!");
            return null;
        }

#endif
    }

#if !EDITOR_MODE
    static public AssetBundle LoadAssetBundle(string path)
    {
        if(_Manifest == null)
            InitCommonList();
        try
        {
            string assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
            SmartPath(ref assetPath);
            AssetBundle ab = AssetBundle.LoadFromFile(assetPath);
            return ab;
        }
        catch(System.Exception e)
        {
            Debug.LogWarning("AssetPath: " + path + " is not excist!");
            return null;
        }
    }
#endif

    static public void LaunchBundle(int entityid)
    {
        EntityData eData = EntityData.GetData(entityid);
        DisplayData dData = null;
        SkillData sData = null;
        if(eData != null)
        {
            dData = DisplayData.GetData(eData._DisplayId);
            if(dData != null)
                AssetLoader.LaunchBundle(dData._AssetPath);
            for(int j=0; j < eData._Skills.Length; ++j)
            {
                sData = SkillData.GetData(eData._Skills[j]);
                if(sData != null)
                {
                    AssetLoader.LaunchBundle(sData._CastEffect);
                    AssetLoader.LaunchBundle(sData._BeattackEffect);
                    AssetLoader.LaunchBundle(sData._SkillEffect);
                }
            }
        }
    }

    static public void DisposeBundle(int entityid)
    {
        EntityData eData = EntityData.GetData(entityid);
        DisplayData dData = null;
        SkillData sData = null;
        if(eData != null)
        {
            dData = DisplayData.GetData(eData._DisplayId);
            if(dData != null)
                AssetLoader.UnloadAsset(dData._AssetPath, true);
            for(int j=0; j < eData._Skills.Length; ++j)
            {
                sData = SkillData.GetData(eData._Skills[j]);
                if(sData != null)
                {
                    AssetLoader.UnloadAsset(sData._CastEffect, true);
                    AssetLoader.UnloadAsset(sData._BeattackEffect, true);
                    AssetLoader.UnloadAsset(sData._SkillEffect, true);
                }
            }
        }
    }

    static public void LaunchBundle(string path)
    {
        if (string.IsNullOrEmpty(path))
            return;
        path = path.ToLower();
        #if !EDITOR_MODE
        if(_Manifest == null)
            InitCommonList();

        try
        {
            string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
            string assetPath;
            for(int i=0; i < dep.Length; ++i)
            {
                assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + dep[i];
                SmartPath(ref assetPath);
                if(!AssetCounter.Excist(assetPath))
                    AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
                else
                    AssetCounter.GetBundle(assetPath);
            }
            assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
            SmartPath(ref assetPath);
            if(!AssetCounter.Excist(assetPath))
                AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
            else
                AssetCounter.GetBundle(assetPath);
        }
        catch(System.Exception e)
        {
            Debug.LogWarning("AssetPath: " + path + " is not excist!");
        }

        #endif
    }

    static public AnimationClip LoadClip(string path)
    {
        if (string.IsNullOrEmpty(path))
            return null;
        path = path.ToLower();
        #if EDITOR_MODE
        AnimationClip clip = Resources.Load<AnimationClip>(path);
        if (clip == null)
            return null;
        return clip;
        #else
        if(_Manifest == null)
        InitCommonList();

        try
        {
            string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
            string assetPath;
            for(int i=0; i < dep.Length; ++i)
            {
                assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + dep[i];
                SmartPath(ref assetPath);
                if(!AssetCounter.Excist(assetPath))
                    AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
                else
                    AssetCounter.GetBundle(assetPath);
            }
            assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
            SmartPath(ref assetPath);
            AssetBundle ab = null;
            if(!AssetCounter.Excist(assetPath))
            {
                ab = AssetBundle.LoadFromFile(assetPath);
                AssetCounter.AddRef(assetPath, ab);
            }
            else
                ab = AssetCounter.GetBundle(assetPath);

            string assetName = path.Substring(path.LastIndexOf("/") + 1);
            AnimationClip clip = ab.LoadAsset<AnimationClip>(assetName);
            return clip;
        }
        catch(System.Exception e)
        {
            Debug.LogWarning("AssetPath: " + path + " is not excist!");
            return null;
        }
        #endif
    }

    static public AudioClip LoadAudio(string path)
    {
        if (string.IsNullOrEmpty(path))
            return null;
        path = path.ToLower();
        #if EDITOR_MODE
        AudioClip clip = Resources.Load<AudioClip>(path);
        if (clip == null)
            return null;
        return clip;
        #else
        if(_Manifest == null)
        InitCommonList();

        try
        {
            string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
            string assetPath;
            for(int i=0; i < dep.Length; ++i)
            {
                assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + dep[i];
                SmartPath(ref assetPath);
                if(!AssetCounter.Excist(assetPath))
                    AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
                else
                    AssetCounter.GetBundle(assetPath);
            }
            assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
            SmartPath(ref assetPath);
            AssetBundle ab = null;
            if(!AssetCounter.Excist(assetPath))
            {
                ab = AssetBundle.LoadFromFile(assetPath);
                AssetCounter.AddRef(assetPath, ab);
            }
            else
                ab = AssetCounter.GetBundle(assetPath);

            string assetName = path.Substring(path.LastIndexOf("/") + 1);
                AudioClip clip = ab.LoadAsset<AudioClip>(assetName);
            return clip;
        }
        catch(System.Exception e)
        {
            Debug.LogWarning("AssetPath: " + path + " is not excist!");
            return null;
        }
        #endif
    }

#if EDITOR_MODE
    static public ResourceRequest LoadAssetAsync(string path)
    {
        path = path.ToLower();
        return Resources.LoadAsync(path);
    }
#else
    static public AssetBundleCreateRequest LoadAssetAsync(string path)
    {
        if (string.IsNullOrEmpty(path))
            return null;
        
        path = path.ToLower();
        if(_Manifest == null)
            InitCommonList();

        string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
        string assetPath;
        for(int i=0; i < dep.Length; ++i)
        {
            assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + dep[i];
            SmartPath(ref assetPath);
            AssetCounter.AddRef(assetPath, AssetBundle.LoadFromFile(assetPath));
        }
        assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
        SmartPath(ref assetPath);
        return AssetBundle.LoadFromFileAsync(assetPath);
    }
#endif

    static void AssetLoadCallback(AssetBundleRequest request)
    {
        
    }

    static public void UnloadAsset(string path, bool destroyObj = false)
    {
        if (string.IsNullOrEmpty(path))
            return;
        
        path = path.ToLower();
#if EDITOR_MODE

        Resources.UnloadUnusedAssets();
#else
        if(_Manifest == null)
        InitCommonList();

        string[] dep = _Manifest.GetAllDependencies(path + Define.ASSET_EXT);
        string assetPath;

        int refCount = 1;
        if(destroyObj)
            refCount = AssetCounter.GetRef(path);
        for(int j = 0; j < refCount; ++j)
        {
            for(int i=0; i < dep.Length; ++i)
            {
                assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + dep[i];
                SmartPath(ref assetPath);
                AssetCounter.DelRef(assetPath);
            }
        }

        assetPath = Application.persistentDataPath + "/" + Define.PackageVersion + "/" + path + Define.ASSET_EXT;
        SmartPath(ref assetPath);
        if (destroyObj)
            AssetCounter.Dispose(assetPath);
        else
            AssetCounter.DelRef(assetPath);
#endif
    }
}
