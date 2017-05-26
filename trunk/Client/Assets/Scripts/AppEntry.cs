using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;

public class AppEntry : MonoBehaviour {

    public Transform _UiRoot;

    public Transform[] _GlobalObjs;

	// Use this for initialization
	void Start () {
        DontDestroyOnLoad(this);
        for(int i=0; i < _GlobalObjs.Length; ++i)
        {
            DontDestroyOnLoad(_GlobalObjs[i]);
        }

        DataLoader.Init();
        UIManager.Init(_UiRoot);

        //init network
        if (NetWoking.Open("10.10.10.2", 9900))
        {
            UIManager.Show("LoginPanel");
            DataLoader.BeginLoad();
        }
	}
	
	// Update is called once per frame
	void Update () {
        NetWoking.SetupNetFPS();
        TimerManager.Update();
        DataLoader.Update();
        UIManager.Update();
	}
}
