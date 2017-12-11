using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using LuaInterface;
using FairyGUI;

public class AppEntry : MonoBehaviour {

    string context;
    string logUrl = "";

    long _PauseTimeStamp;

	// Use this for initialization
	void Start () {
        DontDestroyOnLoad(this);
        Screen.sleepTimeout = SleepTimeout.NeverSleep;
        SceneManager.sceneLoaded += OnSceneLoaded;
        GRoot.inst.SetContentScaleFactor(1920, 1080, UIContentScaler.ScreenMatchMode.MatchHeight);

        Define.LaunchUIBundle("jiazai");
        UIConfig.defaultFont = "方正楷体_GBK";
        UIConfig.windowModalWaiting = "ui://jiazai/zairu_com";
        UIConfig.globalModalWaiting = "ui://jiazai/jiazai_com";
        //UIConfig.buttonSound = AssetLoader.LoadAudio("Audio/effect");
        Application.logMessageReceived += (condition, stackTrace, type) => {
            if(type == LogType.Log || string.IsNullOrEmpty(logUrl))
                return;
            
            context = condition + "\n" + stackTrace + "\n" + type;
            WWWForm form = new WWWForm();
            form.AddField("log", context);
            WWW www = new WWW(logUrl, form);
        };
        AudioSystem.Init();
        DataLoader.Init();
        UIManager.Init();
        LuaManager.Init();
        GuideSystem.Init();
        Define.Init();
        CameraEffect.Init();
        BagSystem.Init();
        ChatSystem.Init();
        RankSystem.Init();
        YYSystem.Init();
        YYSystem.Login();
//        WeatherSystem.Init();
        UIManager.Show("yemiantishi");
        UIManager.Show("denglu");

        logUrl = Define.GetStr("LogUrl");
        if (string.IsNullOrEmpty(logUrl))
            logUrl = "http://106.75.78.151:8080/log";

//        UIPackage.AddPackage("UI/ModalWaiting");
//        UIConfig.globalModalWaiting = "ui://ModalWaiting/GlobalModalWaiting";
//        UIConfig.windowModalWaiting = "ui://ModalWaiting/WindowModalWaiting";
        //UIObjectFactory.SetPackageItemExtension("ui://ModalWaiting/GlobalModalWaiting", typeof(GlobalWaiting));
        //GRoot.inst.ShowModalWait();
//        UIManager.GetWindow("denglu").ShowModalWait();
        //UIObjectFactory.SetLoaderExtension(typeof(MyGLoader));
	}

	// Update is called once per frame
	void Update () {
        NetWoking.SetupNetFPS();
        TimerManager.Update();
        DataLoader.Update();
        UIManager.Update();
        World.Update();
        Battle.Update();
        GuideSystem.Update();
        SceneLoader.Update();
        CameraEffect.Update();
        ExceptionHandle.Update();
//        WeatherSystem.Update();
        if (StageCamera.main != null)
        {
            if (StageCamera.main.transform.position.y != 995)
            {
                StageCamera.main.transform.position = new Vector3(StageCamera.main.transform.position.x, 995, StageCamera.main.transform.position.z);
            }
        }
        if(Stage.inst.y != -1000)
            Stage.inst.SetXY(0f, -1000f);
	}

    void OnApplicationPause(bool pause)
    {
        if (pause)
        {
            _PauseTimeStamp = TimerManager.GetTimeStamp();
        }
        else
        {
            long timegap = TimerManager.GetTimeStamp() - _PauseTimeStamp;
            TimerManager.SetTickerGap(timegap);
            if (Battle.InBattle)
                NetWoking.S.QueryBattleRound(Battle._Turn);
        }
    }

    public static void OnSceneLoaded(Scene arg0, LoadSceneMode arg1)
    {
        GuideSystem.ClearGuide();
        switch(arg0.name)
        {
            case Define.SCENE_LOGIN:
                UIManager.Show("denglu");
                break;
            case Define.SCENE_CREATE:
                UIManager.Show("xuanren");
                break;
            case Define.SCENE_MAIN:
                UIManager.Show("zhujiemian");
                break;
            default:
                break;
        }
        if (Proxy4Lua._SceneUI.ContainsKey(arg0.name))
        {
            for(int i=0; i < Proxy4Lua._SceneUI[arg0.name].Count; ++i)
            {
                UIManager.Show(Proxy4Lua._SceneUI[arg0.name][i]);
            }
        }
        Proxy4Lua.ClearHoldUI(arg0.name);
    }
//
//    void OnGUI()
//    {
//        if (GUILayout.Button("PlayEffect"))
//        {
//            COM_Chat cha = new COM_Chat();
//            cha.PlayerInstId = 111;
//            cha.Level = "3";
//            cha.PlayerName = "哈哈哈";
//            cha.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
//            cha.Content = "温柔攻克欧派让客户";
//            cha.Type = 1;
//            Proxy4Lua.SendChat(cha);
//        }
//
//        if (GUILayout.Button("PlayBackground"))
//        {
//            COM_Chat cha = new COM_Chat();
//            cha.Content = "温泛塞封柔攻克欧派让客户";
//            cha.Type = 0;
//            Proxy4Lua.SendChat(cha);
//        }
//
//        if (GUILayout.Button("PlayBackground2"))
//            AudioSystem.PlayBackground("Audio/background2");
//    }
        }
