using System;
using System.IO;
using UnityEngine;
using UnityEditor;
using System.Collections;
using System.Collections.Generic;
using SevenZip.Compression.LZMA;
using SevenZip;
using System.Diagnostics;
using System.Security.Cryptography;
using UnityEditor.SceneManagement;

public class Tools {

    static void RefreshScene()
    {
        List<string> dirs = new List<string>();
        GetDirs(Application.dataPath + "/Scenes/", ref dirs);
        EditorBuildSettingsScene[] scenes = new EditorBuildSettingsScene[dirs.Count];
        for(int i=0; i < scenes.Length; ++i)
        {
            scenes [i] = new EditorBuildSettingsScene(dirs[i], true);
        }
        EditorBuildSettings.scenes = scenes;
        AssetDatabase.SaveAssets();
        EditorSceneManager.OpenScene("Assets/Scenes/Scene/Login.unity");
    }

    private static void GetDirs(string dirPath, ref List<string> dirs)
    {
        foreach (string path in Directory.GetFiles(dirPath))
        {
            if(System.IO.Path.GetExtension(path) == ".unity") 
            {
                dirs.Add(path.Substring(path.IndexOf("Assets/")));
            }
        }
        if (Directory.GetDirectories(dirPath).Length > 0)
        {
            foreach (string path in Directory.GetDirectories(dirPath))
                GetDirs(path,ref dirs);
        }
    }

    [MenuItem("Tools/每次更新后执行", false, 1)]
    static void AfterUpdate()
    {
        RefreshScene();
        ToLuaMenu.ClearLuaWraps();
        CopyTableAndScripts();
        GenRPC();
    }

    static void GenRPC()
    {
        Process p = new Process();
        p.StartInfo.FileName = Application.dataPath + "/../../tools/rpc.bat";
        p.StartInfo.Arguments = PathDefine.TABLE_ASSET_PATH.Replace("/", "\\");
        p.StartInfo.UseShellExecute = true;
        p.StartInfo.WorkingDirectory = "../tools/";
        p.Start();
    }

    [MenuItem("Tools/1.拷贝数据表和脚本", false, 2)]
    static void CopyTableAndScripts()
    {
        Process p = new Process();
        p.StartInfo.FileName = Application.dataPath + "/../../tools/copyClientFiles.bat";
        p.StartInfo.Arguments = PathDefine.TABLE_ASSET_PATH.Replace("/", "\\");
        p.StartInfo.UseShellExecute = true;
        p.StartInfo.WorkingDirectory = "../tools/";
        p.Start();
    }

    [MenuItem("Tools/2.打资源包/PC", false, 3)]
    static public void BuildAssetBundlePC()
    {
        SetScene();
        SetPlayer();
        SetEffect();
        SetUI();
        SetAnim();
        SetAudio();
        SetTable();
        SetLua();

        string resPkgPath = Application.streamingAssetsPath + "/" + Define.PackageVersion;
        if (!Directory.Exists(resPkgPath))
            Directory.CreateDirectory(resPkgPath);
        BuildPipeline.BuildAssetBundles(Application.streamingAssetsPath + "/" + Define.PackageVersion, BuildAssetBundleOptions.ChunkBasedCompression | BuildAssetBundleOptions.DeterministicAssetBundle, BuildTarget.StandaloneWindows64);
        CreateMD5File();
        AssetDatabase.Refresh();
    }

    [MenuItem("Tools/2.打资源包/Android", false, 3)]
    static public void BuildAssetBundleAndroid()
    {
        SetScene();
        SetPlayer();
        SetEffect();
        SetUI();
        SetAnim();
        SetAudio();
        SetTable();
        SetLua();

        string resPkgPath = Application.streamingAssetsPath + "/" + Define.PackageVersion;
        if (!Directory.Exists(resPkgPath))
            Directory.CreateDirectory(resPkgPath);
        BuildPipeline.BuildAssetBundles(Application.streamingAssetsPath + "/" + Define.PackageVersion, BuildAssetBundleOptions.ChunkBasedCompression | BuildAssetBundleOptions.DeterministicAssetBundle, BuildTarget.Android);
        CreateMD5File();
        AssetDatabase.Refresh();
    }

    [MenuItem("Tools/2.打资源包/iOS", false, 3)]
    static public void BuildAssetBundleiOS()
    {
        SetScene();
        SetPlayer();
        SetEffect();
        SetUI();
        SetAnim();
        SetAudio();
        SetTable();
        SetLua();

        string resPkgPath = Application.streamingAssetsPath + "/" + Define.PackageVersion;
        if (!Directory.Exists(resPkgPath))
            Directory.CreateDirectory(resPkgPath);
        BuildPipeline.BuildAssetBundles(Application.streamingAssetsPath + "/" + Define.PackageVersion, BuildAssetBundleOptions.ChunkBasedCompression | BuildAssetBundleOptions.DeterministicAssetBundle, BuildTarget.iOS);
        CreateMD5File();
        AssetDatabase.Refresh();
    }

    [MenuItem("Tools/3.删除临时资源", false, 4)]
    static void DeleteTempFiles()
    {
        EditorBuildSettings.scenes = new EditorBuildSettingsScene[1]{ new EditorBuildSettingsScene("Assets/Scenes/Scene/Login.unity", true) };
        AssetDatabase.SaveAssets();

        Process p = new Process();
        p.StartInfo.FileName = Application.dataPath + "/../../tools/removeResources.bat";
        p.StartInfo.Arguments = PathDefine.TABLE_ASSET_PATH.Replace("/", "\\");
        p.StartInfo.UseShellExecute = true;
        p.StartInfo.WorkingDirectory = "../tools/";
        p.Start();
    }

//    static void SetScene()
//    {
//        BuildPlayerOptions opti = new BuildPlayerOptions();
//        opti.locationPathName = Application.streamingAssetsPath + "/" + Define.PackageVersion + "/";
//        opti.options = BuildOptions.CompressWithLz4;
//        opti.targetGroup = BuildTargetGroup.Standalone;
//        BuildPipeline.BuildPlayer(opti);
//    }

    static void SetPlayer()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Resources/" + PathDefine.PLAYER_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);
            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            aimport.assetBundleName = PathDefine.PLAYER_ASSET_PATH + shortPath.Remove(shortPath.IndexOf(".")) + Define.ASSET_EXT;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;
                
                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(_assetPath + " depands on: " + deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }

    static void SetEffect()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Resources/" + PathDefine.EFFECT_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);
            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            aimport.assetBundleName = PathDefine.EFFECT_ASSET_PATH + shortPath.Remove(shortPath.IndexOf(".")) + Define.ASSET_EXT;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;
                
                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(_assetPath + " depands on: " + deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }

    static void SetUI()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Resources/" + PathDefine.UI_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);

            #if (UNITY_EDITOR && UNITY_IOS) || (UNITY_EDITOR && UNITY_ANDROID)
            TextureImporter timport;
            if (aimport is UnityEditor.TextureImporter)
            {
                if (!aimport.assetPath.Contains("!a."))
                {
                    #if UNITY_IOS
                        TextureImporterPlatformSettings tips = new TextureImporterPlatformSettings();
                        tips.format = TextureImporterFormat.PVRTC_RGB4;
                    #elif UNITY_ANDROID
                        TextureImporterPlatformSettings tips = new TextureImporterPlatformSettings();
                        tips.format = TextureImporterFormat.ETC_RGB4;
                    #endif

                    timport = (TextureImporter)aimport;
                    timport.textureCompression = TextureImporterCompression.Compressed;
                    timport.SaveAndReimport();
                    timport.SetPlatformTextureSettings(tips);
                }
            }
            #endif

            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            int idx = shortPath.IndexOf("@");
            if (idx != -1)
            {
                shortPath = shortPath.Remove(idx);
                shortPath += "_res";
            }
            else
            {
                shortPath = shortPath.Remove(shortPath.IndexOf("."));
                shortPath += "_desc";
            }
            string finalPath = PathDefine.UI_ASSET_PATH + shortPath + Define.ASSET_EXT;
            aimport.assetBundleName = finalPath;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;
                
                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(_assetPath + " depands on: " + deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }

    static void SetAnim()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Resources/" + PathDefine.ANIM_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);
            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            aimport.assetBundleName = PathDefine.ANIM_ASSET_PATH + shortPath.Remove(shortPath.IndexOf(".")) + Define.ASSET_EXT;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;

                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(_assetPath + " depands on: " + deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }

    static void SetAudio()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Resources/" + PathDefine.AUDIO_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);
            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            aimport.assetBundleName = PathDefine.AUDIO_ASSET_PATH + shortPath.Remove(shortPath.IndexOf(".")) + Define.ASSET_EXT;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;

                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(_assetPath + " depands on: " + deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }

    static void SetScene()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Scenes/" + PathDefine.SCENE_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);
            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            aimport.assetBundleName = PathDefine.SCENE_ASSET_PATH + shortPath.Remove(shortPath.IndexOf(".")) + Define.ASSET_EXT;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;

                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(_assetPath + " depands on: " + deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }

    static void SetTable()
    {
        CollectAllFiles(string.Format("{0}/{1}", Application.dataPath, "Resources/" + PathDefine.TABLE_ASSET_PATH));
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            string _assetPath = "Assets" + _AllFiles[i].Substring (Application.dataPath.Length);
            _assetPath = _assetPath.Replace("\\", "/");
            AssetImporter aimport = AssetImporter.GetAtPath(_assetPath);
            string shortPath = _assetPath.Substring(_assetPath.LastIndexOf("/") + 1);
            aimport.assetBundleName = PathDefine.TABLE_ASSET_PATH + shortPath.Remove(shortPath.IndexOf(".")) + Define.ASSET_EXT;
            string[] deps = AssetDatabase.GetDependencies(_assetPath);
            _assetPath = _assetPath.ToLower();
            for(int j=0; j < deps.Length; ++j)
            {
                deps [j] = deps [j].ToLower();
                if (deps [j].IndexOf(".js") != -1)
                    continue;

                if (deps [j].IndexOf(".cs") != -1)
                    continue;
                
                if (deps [j].Equals(_assetPath))
                    continue;

                if (_AllDependences.Contains(deps [j]))
                    continue;

                _AllDependences.Add(deps[j]);
                UnityEngine.Debug.Log(deps[j]);
            }
        }
        for(int i=0; i < _AllDependences.Count; ++i)
        {
            AssetImporter aimport = AssetImporter.GetAtPath(_AllDependences[i]);
            aimport.assetBundleName = PathDefine.COMMON_ASSET_PATH + AssetDatabase.AssetPathToGUID(_AllDependences[i]) + Define.ASSET_EXT;
        }
        _AllFiles.Clear();
        _AllDirectories.Clear();
        _AllDependences.Clear();
    }
    static void SetLua()
    {
        string sourcePath = string.Format("{0}/{1}/{2}", Application.dataPath, "Resources", PathDefine.LUA_ASSET_PATH);
        if (!Directory.Exists(sourcePath))
        {
            return;
        }

        string[] files = Directory.GetFiles(sourcePath, "*.lua", SearchOption.AllDirectories);
        int len = sourcePath.Length;

        if (sourcePath[len - 1] == '/' || sourcePath[len - 1] == '\\')
        {
            --len;
        }         

        for (int i = 0; i < files.Length; i++)
        {
            string str = files[i].Remove(0, len);
            string dest = LuaConst.luaResDir + "/" + str;
            //dest += ".bytes";
            string dir = Path.GetDirectoryName(dest);
            Directory.CreateDirectory(dir);
            File.Copy(files[i], dest, true);
        }
    }



    static List<string> _AllFiles = new List<string>();
    static List<string> _AllDirectories = new List<string>();
    static List<string> _AllDependences = new List<string>();

    static void CollectFiles(string path)
    {
        string[] files = Directory.GetFiles(path);
        for(int i=0; i < files.Length; ++i)
        {
            if (files [i].EndsWith(".meta"))
                continue;

            _AllFiles.Add(files[i]);
        }
    }

    static void CollectAllFiles(string path)
    {
        CollectFiles(path);
        string[] dirs = Directory.GetDirectories(path);
        if (dirs.Length != 0)
            _AllDirectories.AddRange(dirs);
        for(int i=0; i < dirs.Length; ++i)
        {
            CollectAllFiles(dirs[i]);
        }
    }

    [MenuItem("Tools/Clear")]
    static void ClearAssetBundlesName()
    {
        int length = AssetDatabase.GetAllAssetBundleNames().Length;
        UnityEngine.Debug.Log(length);
        string[] oldAssetBundleNames = new string[length];
        for(int i=0; i < length; ++i)
        {
            oldAssetBundleNames[i] = AssetDatabase.GetAllAssetBundleNames()[i];
        }
        for(int i=0; i < oldAssetBundleNames.Length; ++i)
        {
            AssetDatabase.RemoveAssetBundleName(oldAssetBundleNames[i], true);
        }
        length = AssetDatabase.GetAllAssetBundleNames().Length;
        UnityEngine.Debug.Log(length);
    }

    [MenuItem("Tools/TestLoad")]
    static void TestLoad()
    {
        AssetLoader.LoadAsset(PathDefine.PLAYER_ASSET_PATH + "longtaizi");
    }

    [MenuItem ("MyMenu/CompressFile")]
    static void CompressFile () 
    {
        //压缩文件
//        Compress7zZip(Application.streamingAssetsPath + "/1_0_0/", Application.streamingAssetsPath + "/1_0_0.7z", "");
        AssetDatabase.Refresh();

    }
    [MenuItem ("MyMenu/DecompressFile")]
    static void DecompressFile () 
    {
        //解压文件
        AssetDatabase.Refresh();
    }

    /// <summary>
    /// 执行压缩命令结果
    /// </summary>
    public enum CompressResults
    {
        Success,
        SourceObjectNotExist,
        UnKnown
    }

    /// <summary>
    /// 执行解压缩命令结果
    /// </summary>
    public enum UnCompressResults
    {
        Success,
        SourceObjectNotExist,
        PasswordError,
        UnKnown
    }

//    #region 7zZip压缩、解压方法
    /// <summary>
    /// 压缩文件 
    /// </summary>
    /// <param name="objectPathName">压缩对象（即可以是文件夹|也可以是文件）</param>
    /// <param name="objectZipPathName">保存压缩文件的路径</param>
    /// <param name="strPassword">加密码</param>
    /// 测试压缩文件夹：压缩文件（objectZipPathName）不能放在被压缩文件（objectPathName）内，否则报“文件夹被另一进程使用中”错误。
    /// <returns></returns>
//    static CompressResults Compress7zZip(String objectPathName, String objectZipPathName, String strPassword)
//    {
//        try
//        {
//            //http://sevenzipsharp.codeplex.com/releases/view/51254 下载sevenzipsharp.dll
//            //SevenZipSharp.dll、zLib1.dll、7z.dll必须同时存在，否则常报“加载7z.dll错误”
//            string libPath = Application.dataPath + "/Plugins/7z.dll";
//            SevenZip.SevenZipCompressor.SetLibraryPath(libPath);
//            SevenZip.SevenZipCompressor sevenZipCompressor = new SevenZip.SevenZipCompressor();
//            sevenZipCompressor.CompressionLevel = SevenZip.CompressionLevel.Fast;
//            sevenZipCompressor.ArchiveFormat = SevenZip.OutArchiveFormat.Zip;
//
//            //被压缩对象是否存在
//            int beforeObjectNameIndex = objectPathName.LastIndexOf('\\');
//            if(beforeObjectNameIndex == -1)
//                beforeObjectNameIndex = objectPathName.LastIndexOf('/');
//            string objectPath = objectPathName.Substring(0, beforeObjectNameIndex);
//            //System.IO.DirectoryInfo directoryInfo = new System.IO.DirectoryInfo(objectPathName);
//            if (Directory.Exists(objectPathName)/*directoryInfo.Exists*/ == false && System.IO.File.Exists(objectPathName) == false)
//            {
//                return CompressResults.SourceObjectNotExist;
//            }
//            int beforeObjectRarNameIndex = objectZipPathName.LastIndexOf('\\');
//            if(beforeObjectRarNameIndex == -1)
//                beforeObjectRarNameIndex = objectZipPathName.LastIndexOf('/');
//            int objectRarNameIndex = beforeObjectRarNameIndex + 1;
//            //string objectZipName = objectZipPathName.Substring(objectRarNameIndex);
//            string objectZipPath = objectZipPathName.Substring(0, beforeObjectRarNameIndex);
//            //目标目录、文件是否存在
//            if (System.IO.Directory.Exists(objectZipPath) == false)
//            {
//                System.IO.Directory.CreateDirectory(objectZipPath);
//            }
//            else if (System.IO.File.Exists(objectZipPathName) == true)
//            {
//                System.IO.File.Delete(objectZipPathName);
//            }
//
//            if (Directory.Exists(objectPathName))       //压缩对象是文件夹
//            {
//                if (String.IsNullOrEmpty(strPassword))
//                {
//                    sevenZipCompressor.CompressDirectory(objectPathName, objectZipPathName);
//                }
//                else
//                {
//                    sevenZipCompressor.CompressDirectory(objectPathName, objectZipPathName, strPassword);
//                }
//            }
//            else        //压缩对象是文件 无加密方式
//            {
//                sevenZipCompressor.CompressFiles(objectZipPathName, objectPathName);
//            }
//
//            return CompressResults.Success;
//        }
//        catch(Exception ex)
//        {
//            UnityEngine.Debug.Log("压缩文件失败！" + ex.Message);
//            return CompressResults.UnKnown;
//        }
//    }

    /// <summary>
    /// 解压缩文件
    /// </summary>
    /// <param name="zipFilePathName">zip文件具体路径+名</param>
    /// <param name="unCompressDir">解压路径</param>
    /// <param name="strPassword">解密码</param>
    /// <returns></returns>
//    static UnCompressResults UnCompress7zZip(String zipFilePathName, String unCompressDir, String strPassword)
//    {
//        try
//        {
//            //SevenZipSharp.dll、zLib1.dll、7z.dll必须同时存在，否则常报“加载7z.dll错误”而项目引用时，只引用SevenZipSharp.dll就可以了
//            string libPath = System.AppDomain.CurrentDomain.BaseDirectory + @"..\..\dll\7z.dll";
//            SevenZip.SevenZipCompressor.SetLibraryPath(libPath);
//
//            bool isFileExist = File.Exists(zipFilePathName);
//            if (false == isFileExist)
//            {
//                UnityEngine.Debug.Log("解压文件不存在！" + zipFilePathName);
//                return UnCompressResults.SourceObjectNotExist;
//            }
//            File.SetAttributes(zipFilePathName, FileAttributes.Normal);     //去掉只读属性
//
//            if (Directory.Exists(unCompressDir) == false)
//            {
//                Directory.CreateDirectory(unCompressDir);
//            }
//
//            SevenZip.SevenZipExtractor sevenZipExtractor;
//            if (String.IsNullOrEmpty(strPassword))
//            {
//                sevenZipExtractor = new SevenZip.SevenZipExtractor(zipFilePathName);
//            }
//            else
//            {
//                sevenZipExtractor = new SevenZip.SevenZipExtractor(zipFilePathName, strPassword);
//            }
//
//            sevenZipExtractor.ExtractArchive(unCompressDir);
//            sevenZipExtractor.Dispose();
//            return UnCompressResults.Success;
//        }
//        catch(Exception ex)
//        {
//            UnityEngine.Debug.Log("解压缩文件失败！" + ex.Message);
//            return UnCompressResults.UnKnown;
//        }
//    }
//    #endregion

    [MenuItem("Tools/生成资源MD5文件", false, 5)]
    static public void CreateMD5File()
    {
        CollectAllFiles(Application.streamingAssetsPath + "/" + Define.PackageVersion);
        string resMd5 = "";
        for(int i=0; i < _AllFiles.Count; ++i)
        {
            resMd5 += _AllFiles[i].Remove(0, _AllFiles[i].IndexOf(Define.PackageVersion) + Define.PackageVersion.Length + 1) + ":" + GetMD5WithFilePath(_AllFiles[i]);
            resMd5 += "\n";
        }
        File.WriteAllText(Application.streamingAssetsPath + "/md5.txt", resMd5);
        UnityEngine.Debug.Log("MD5 File Created.");
        AssetDatabase.Refresh();
    }

    static public string GetMD5WithFilePath(string filePath)
    {
       FileStream file = new FileStream(filePath, FileMode.Open, FileAccess.Read, FileShare.Read);
       MD5CryptoServiceProvider md5 = new MD5CryptoServiceProvider();
       byte[] hash_byte = md5.ComputeHash(file);
       string str = System.BitConverter.ToString(hash_byte);
       str = str.Replace("-", "");
       return str;
    }
}

public class FindMissingScriptsRecursively : EditorWindow   
{  
    static int go_count = 0, components_count = 0, missing_count = 0;  

    [MenuItem("Window/FindMissingScriptsRecursively")]  
    public static void ShowWindow()  
    {  
        EditorWindow.GetWindow(typeof(FindMissingScriptsRecursively));  
    }  

    public void OnGUI()  
    {  
        if (GUILayout.Button("Find Missing Scripts in selected GameObjects"))  
        {  
            FindInSelected();  
        }  
    }  
    private static void FindInSelected()  
    {  
        GameObject[] go = Selection.gameObjects;  
        go_count = 0;  
        components_count = 0;  
        missing_count = 0;  
        foreach (GameObject g in go)  
        {  
            FindInGO(g);  
        }
        UnityEngine.Debug.Log(string.Format("Searched {0} GameObjects, {1} components, found {2} missing", go_count, components_count, missing_count));  
    }  

    private static void FindInGO(GameObject g)  
    {  
        go_count++;  
        Component[] components = g.GetComponents<Component>();  
        for (int i = 0; i < components.Length; i++)  
        {  
            components_count++;  
            if (components[i] == null)  
            {  
                SerializedObject so = new SerializedObject(g);
                SerializedProperty sp = so.FindProperty("m_Component");
                sp.DeleteArrayElementAtIndex(i);
                missing_count++;  
                string s = g.name;  
                Transform t = g.transform;  
                while (t.parent != null)   
                {  
                    s = t.parent.name +"/"+s;  
                    t = t.parent;  
                }  
                UnityEngine.Debug.Log (s + " has an empty script attached in position: " + i, g);  
                so.ApplyModifiedProperties();
                EditorUtility.SetDirty(g);
            }  
        }  
        // Now recurse through each child GO (if there are any):  
        foreach (Transform childT in g.transform)  
        {  
            //Debug.Log("Searching " + childT.name  + " " );  
            FindInGO(childT.gameObject);  
        }  
    }

    [MenuItem("Tools/删除场景中的空动画控制器")]
    public static void FindEmptyAnimator()
    {
        int total = 0;
        Animator[] animators = GameObject.FindObjectsOfType<Animator>();
        for (int i = 0; i < animators.Length; ++i)
        {
            if (animators[i].runtimeAnimatorController == null)
            {
                total++;
                GameObject.DestroyImmediate(animators[i]);
            }
        }
        UnityEngine.Debug.Log("Destroy " + total + " null Animators.");
    }
}  