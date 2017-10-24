using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using LuaInterface;
using FairyGUI;

public class AppEntry : MonoBehaviour {

    string context;
	// Use this for initialization
	void Start () {
        DontDestroyOnLoad(this);
        SceneManager.sceneLoaded += OnSceneLoaded;
        GRoot.inst.SetContentScaleFactor(1920, 1080, UIContentScaler.ScreenMatchMode.MatchHeight);
        UIConfig.defaultFont = "方正楷体_GBK";
        //UIConfig.buttonSound = AssetLoader.LoadAudio("Audio/effect");
        Application.logMessageReceived += (condition, stackTrace, type) => {
            context += condition + "\n" + stackTrace + "\n" + type;
        };
        AudioSystem.Init();
        DataLoader.Init();
        UIManager.Init();
        Define.Init();
        CameraEffect.Init();
        BagSystem.Init();
//        WeatherSystem.Init();
        UIManager.Show("yemiantishi");
        UIManager.Show("denglu");

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

    public static void OnSceneLoaded(Scene arg0, LoadSceneMode arg1)
    {
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
    }
//
//    void OnGUI()
//    {
//        if (GUILayout.Button("PlayEffect"))
//            Proxy4Lua.ReturnToLogin();
//
//        if (GUILayout.Button("PlayBackground"))
//            AudioSystem.PlayBackground("Audio/background");
//
//        if (GUILayout.Button("PlayBackground2"))
//            AudioSystem.PlayBackground("Audio/background2");
//    }
}
