using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;

public class UIWindow {

    string _UiName;
    GameObject _Ui;
    LuaState _Lua;
    LuaFunction _Func;

    public UIWindow(string uiName)
    {
        Init(uiName);
    }

    public string UIName
    {
        get{ return _UiName; }
    }

    public void Init(string uiName)
    {
        _UiName = uiName;
        _Ui = AssetLoader.LoadAsset(PathDefine.UI_ASSET_PATH + uiName);
        if (_Ui == null)
        {
            Debug.LogWarning("No Asset has been loaded : " + _UiName);
        }
        else
        {
            _Ui.transform.SetParent(UIManager.UIRoot);
            _Ui.transform.localPosition = Vector3.zero;
        }
        Debug.Log(" UI Name : " + uiName);
        _Lua = UIManager._Lua;
        _Lua.DoFile(_UiName + ".lua");
    }

    //初始数据赋值
    public void Start()
    {
        if (_Ui == null)
            return;

        if (_Ui.activeSelf == false)
            return;
        
        if (_Lua == null)
            return;
        
        _Func = _Lua.GetFunction(_UiName + "_start");
        if (_Func == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no start function.");
            return;
        }
        _Func.Call();
        _Lua.CheckTop();
        _Func.Dispose();
    }

    //每帧更新
    public void Update()
    {
        if (_Ui == null)
            return;

        if (_Ui.activeSelf == false)
            return;

        if (_Lua == null)
            return;
        
        _Func = _Lua.GetFunction(_UiName + "_update");
        if (_Func == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no update function.");
            return;
        }
        _Func.Call();
        _Lua.CheckTop();
        _Func.Dispose();
    }

    //每秒更新
    public void Tick()
    {
        if (_Ui == null)
            return;

        if (_Ui.activeSelf == false)
            return;

        if (_Lua == null)
            return;
        
        _Func = _Lua.GetFunction(_UiName + "_tick");
        if (_Func == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no tick function.");
            return;
        }
        _Func.Call();
        _Lua.CheckTop();
        _Func.Dispose();
    }

    #region 外部调用请使用UIManager接口
    public bool IsShow
    {
        get
        {
            if (_Ui == null)
                return false;
            return _Ui.activeSelf;
        }
    }

    public void Show()
    {
        if (_Ui == null)
            return;

        if (_Ui.activeSelf)
            return;

        _Ui.SetActive(true);
    }

    public void Hide()
    {
        if (_Ui == null)
            return;

        if (!_Ui.activeSelf)
            return;
        
        _Ui.SetActive(false);
    }

    public void Dispose()
    {
        if (_Ui == null)
            return;

        if (!_Ui.activeSelf)
            return;

        GameObject.Destroy(_Ui);
    }
    #endregion
}
