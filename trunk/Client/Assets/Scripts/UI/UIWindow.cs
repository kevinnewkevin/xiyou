using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;

public class UIWindow {

    string _UiName;
    LuaState _Lua;
    LuaFunction _InitFunc;
    LuaFunction _UpdateFunc;
    LuaFunction _TickFunc;
    object _ParamValue;

    public UIWindow(string uiName, object parVal = null)
    {
        Init(uiName, parVal);
    }

    public string UIName
    {
        get{ return _UiName; }
    }

    void Init(string uiName, object paramVal)
    {
        _UiName = uiName;
        _ParamValue = paramVal;
        string resName = UIManager.GetUIResName(uiName);
        Define.LaunchUIBundle(resName);
        _Lua = UIManager._Lua;
        _Lua.DoFile(_UiName + ".lua");
        Debug.Log(" UI Name : " + uiName);

        _InitFunc = _Lua.GetFunction(_UiName + ".OnEntry");
        if (_InitFunc == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no Entry function.");
        }
        _UpdateFunc = _Lua.GetFunction(_UiName + ".OnUpdate");
        if (_UpdateFunc == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no Update function.");
        }
        _TickFunc = _Lua.GetFunction(_UiName + ".OnTick");
        if (_TickFunc == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no Tick function.");
        }
        _InitFunc.Call(_ParamValue);
        _Lua.CheckTop();
    }

    public FairyGUI.Window GetWindow()
    {
        LuaFunction getWindow = _Lua.GetFunction(_UiName + ".GetWindow");
        if (getWindow == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no getWindow function.");
        }
        return (FairyGUI.Window)getWindow.Call(0)[0];
    }

    //每帧更新
    public void Update()
    {
        if (_Lua == null)
            return;

        if (_UpdateFunc == null)
            return;

        _UpdateFunc.Call();
        _Lua.CheckTop();
    }

    //每秒更新
    public void Tick()
    {
        if (_Lua == null)
            return;
        
        if (_TickFunc == null)
            return;
        
        _TickFunc.Call();
        _Lua.CheckTop();
    }

    #region 外部调用请使用UIManager接口
    public bool IsShow
    {
        get
        {
            if (_Lua == null)
                return false;
            
            LuaFunction func = _Lua.GetFunction(_UiName + ".isShow");
            if (func == null)
            {
                Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no isShow function.");
                return false;
            }
            return (bool)func.Call(0)[0];
        }
    }

    public void Dispose()
    {
        if (_Lua == null)
            return;
        if (_InitFunc != null)
            _InitFunc.Dispose();
        if (_UpdateFunc != null)
            _UpdateFunc.Dispose();
        if (_TickFunc != null)
            _TickFunc.Dispose();
        LuaTable table = _Lua.GetTable(_UiName);
        if(table != null)
            table.Dispose();

        LuaFunction func = _Lua.GetFunction(_UiName + ".OnDispose");
        if (func == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no OnDispose function.");
            return;
        }
        func.Call();
    }

    public void Hide()
    {
        LuaFunction func = _Lua.GetFunction(_UiName + ".OnHide");
        if (func == null)
        {
            Debug.LogWarning(" UI lua Script Named: " + UIName + ".lua has no OnHide function.");
            return;
        }
        func.Call();
    }
    #endregion
}
