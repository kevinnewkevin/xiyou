using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;
using FairyGUI;

public class AppEntry : MonoBehaviour {

    public Transform[] _GlobalObjs;

	// Use this for initialization
	void Start () {
        DontDestroyOnLoad(this);
        for(int i=0; i < _GlobalObjs.Length; ++i)
        {
            DontDestroyOnLoad(_GlobalObjs[i]);
        }
        CopyPastePatch.Apply();
        GRoot.inst.SetContentScaleFactor(1920, 1080);

        DataLoader.Init();
        UIManager.Init();
        DataLoader.BeginLoad();
//
        //init network
        if (NetWoking.Open("10.10.10.188", 10999))
        {
            UIManager.Show("denglu");
            DataLoader.BeginLoad();
        }
        //UIObjectFactory.SetLoaderExtension(typeof(MyGLoader));
	}
	
	// Update is called once per frame
	void Update () {
        NetWoking.SetupNetFPS();
        TimerManager.Update();
        DataLoader.Update();
        UIManager.Update();
        Battle.Update();
        SceneLoader.Update();
	}
}
