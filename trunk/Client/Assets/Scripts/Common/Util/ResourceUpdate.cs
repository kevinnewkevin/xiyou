using System.IO;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ResourceUpdate : MonoBehaviour {

    public string _CenterServer;

    List<string> toDownload = new List<string>();

    int _CrtDownload = 0;

    public bool _UpdateFinish = false;

    public bool _IsDoing = false;

    string _RemoteMd5Content = "";

    public void Init()
    {
        #if EDITOR_MODE
            _UpdateFinish = true;
            StartCoroutine(SkipUpdate());
        #else
            _IsDoing = true;
            _UpdateFinish = false;
            _CenterServer = Define.GetStr("CenterServer");
            CheckNew();
        #endif
    }

    IEnumerator SkipUpdate()
    {
        yield return true;
        UIManager.SetDirty("denglu");
    }

    public void CheckNew()
    {
        string localMd5 = Application.persistentDataPath + "/md5.txt";
        if (!File.Exists(localMd5))
        {
            WWW www = new WWW(Application.streamingAssetsPath + "/md5.txt");
            while(true)
            {
                if (www.isDone)
                {
                    if (!string.IsNullOrEmpty(www.error))
                    {
                        //LuaManager.CallGlobal("ErrorMessage", www.error);
                        break;
                    }
                    else
                    {
                        File.WriteAllText(localMd5, www.text);
                        UnityEngine.Debug.Log("md5 writed");
                        break;
                    }
                }
            }
        }

        StartCoroutine(CheckResource(_CenterServer));
    }

    IEnumerator CheckResource(string url)
    {
        WWW www = new WWW(url);
        yield return true;

        if (www.isDone)
        {
            if (!string.IsNullOrEmpty(www.error))
            {
                //LuaManager.CallGlobal("ErrorMessage", www.error);
                //请求cdn md5时 有个开关 代表是否跳过热更（cdn服务器有问题时的解决方案）
                _UpdateFinish = true;
                _IsDoing = false;
                UIManager.SetDirty("denglu");
            }
            else
            {
                string localMd5 = Application.persistentDataPath + "/md5.txt";
                string localMd5Content = File.ReadAllText(localMd5);
                _RemoteMd5Content = www.text;

                List<string> lmd5 = new List<string>(localMd5Content.Split(new char[]{'\n', ':'}, System.StringSplitOptions.RemoveEmptyEntries));
                List<string> rmd5 = new List<string>(_RemoteMd5Content.Split(new char[]{'\n', ':'}, System.StringSplitOptions.RemoveEmptyEntries));

                for(int i=0; i < lmd5.Count; ++i)
                {
                    if (rmd5.Contains(lmd5 [i]))
                    {
                        if (!lmd5 [i + 1].Equals(rmd5 [i + 1]))
                            toDownload.Add(rmd5[i] + ":" + rmd5[i + 1]);
                    }
                    else
                        toDownload.Add(rmd5[i] + ":" + rmd5[i + 1]);
                }
                StartUpdate();
            }
        }
    }

    void StartUpdate()
    {
        if (_CrtDownload >= toDownload.Count)
        {
            File.WriteAllText(Application.persistentDataPath + "/md5.txt", _RemoteMd5Content);
            _UpdateFinish = true;
            _IsDoing = false;
            UIManager.SetDirty("denglu");
            return;
        }

        string[] path = toDownload [_CrtDownload].Split(new char[]{':'}, System.StringSplitOptions.RemoveEmptyEntries);
        DownLoad(_CenterServer + path[0]);
    }

    IEnumerator DownLoad(string url)
    {
        WWW www = new WWW(url);
        yield return true;

        if (www.isDone)
        {
            if (!string.IsNullOrEmpty(www.error))
            {
                LuaManager.CallGlobal("ErrorMessage", www.error);
            }
            else
            {
                File.WriteAllBytes(url.Replace(_CenterServer, Application.persistentDataPath), www.bytes);
                _CrtDownload ++;
                StartUpdate();
            }
        }
    }
}
