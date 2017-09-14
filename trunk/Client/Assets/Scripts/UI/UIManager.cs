using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;
using System;
using FairyGUI;

public class UIManager {

    static float _Timer;

    static public LuaState _Lua;

    static public void Init()
    {
        _Lua = new LuaState();
        LuaBinder.Bind(_Lua);
        _Lua.Start();
    }

    static Dictionary<string, UIWindow> _Windows = new Dictionary<string, UIWindow>();
    static Dictionary<string, bool> _DirtyPool = new Dictionary<string, bool>();
    static List<string> _WantClearDirty = new List<string>();

    static public void Show(string uiName)
    {
        if (IsShow(uiName))
            return;

        if (!_Windows.ContainsKey(uiName))
            _Windows.Add(uiName, new UIWindow(uiName, null));
        else
            _Windows[uiName].GetWindow().Show();
        
        if (!_DirtyPool.ContainsKey(uiName))
            _DirtyPool.Add(uiName, true);
        else
            _DirtyPool [uiName] = true;
    }

    static public void Show(string uiName, object parVal)
    {
        if (IsShow(uiName))
            return;

        if (!_Windows.ContainsKey(uiName))
            _Windows.Add(uiName, new UIWindow(uiName, parVal));
        else
            _Windows[uiName].GetWindow().Show();

        if (!_DirtyPool.ContainsKey(uiName))
            _DirtyPool.Add(uiName, true);
        else
            _DirtyPool [uiName] = true;
    }

    static public bool IsShow(string uiName)
    {
        bool isShow = _Windows.ContainsKey(uiName) && _Windows [uiName].IsShow;
        return isShow;
    }

    static public void Hide(string uiName)
    {
        if (!IsShow(uiName))
            return;
        
        _Windows [uiName].Hide();

        //AssetLoader.UnloadAsset(PathDefine.UI_ASSET_PATH + uiName);

        if (_DirtyPool.ContainsKey(uiName))
            _DirtyPool.Remove(uiName);
    }

    static public void HideAll()
    {
        foreach(UIWindow window in _Windows.Values)
        {
            window.Hide();
//            AssetLoader.UnloadAsset(PathDefine.UI_ASSET_PATH + window.UIName);

            if (_DirtyPool.ContainsKey(window.UIName))
                _DirtyPool.Remove(window.UIName);
        }
    }

    static public void Dispose(string uiName)
    {
        if (!IsShow(uiName))
            return;

        _Windows [uiName].GetWindow().Dispose();
        _Windows [uiName].Dispose();
        _Windows.Remove(uiName);

        AssetLoader.UnloadAsset(PathDefine.UI_ASSET_PATH + uiName);
        Define.UnloadUIBundle(uiName);

        if (_DirtyPool.ContainsKey(uiName))
            _DirtyPool.Remove(uiName);
    }

    static public void DisposeAll()
    {
        foreach(UIWindow window in _Windows.Values)
        {
            window.GetWindow().Dispose();
            window.Dispose();
            AssetLoader.UnloadAsset(PathDefine.UI_ASSET_PATH + window.UIName);
            Define.UnloadUIBundle(window.UIName);

            if (_DirtyPool.ContainsKey(window.UIName))
                _DirtyPool.Remove(window.UIName);
        }
        _Windows.Clear();
    }

    static public void Update()
    {
        _Timer += Time.deltaTime;

        foreach(UIWindow window in _Windows.Values)
        {
            window.Update();
            if (_Timer >= 1f)
                window.Tick();
        }

        if (_Timer >= 1f)
            _Timer = 0f;

        //重置界面Dirty属性
        for(int i=0; i < _WantClearDirty.Count; ++i)
        {
            if (!_DirtyPool.ContainsKey(_WantClearDirty[i]))
                continue;

            _DirtyPool [_WantClearDirty[i]] = false;
        }
        _WantClearDirty.Clear();
    }

    static public void SetDirty(string uiName)
    {
        if (!_DirtyPool.ContainsKey(uiName))
            return;

        _DirtyPool [uiName] = true;
    }

    static public bool IsDirty(string uiName)
    {
        if (!_DirtyPool.ContainsKey(uiName))
            return false;

        return _DirtyPool [uiName];
    }

    static public void ClearDirty(string uiName)
    {
        if (_WantClearDirty.Contains(uiName))
            return;
        
        _WantClearDirty.Add(uiName);
    }

    static public Window ShowMessageBox()
    {
        if (!_Windows.ContainsKey("tanchukuang"))
            _Windows.Add("tanchukuang", new UIWindow("tanchukuang"));
        else
            _Windows["tanchukuang"].GetWindow().Show();

        if (!_DirtyPool.ContainsKey("tanchukuang"))
            _DirtyPool.Add("tanchukuang", false);

        return _Windows["tanchukuang"].GetWindow();
    }

    static public void HideMessageBox()
    {
        Hide("tanchukuang");
    }

    static public Window GetWindow(string uiName)
    {
        if (!_Windows.ContainsKey(uiName))
            return null;

        return _Windows [uiName].GetWindow();
    }

    static public void RegIDirty(string uiName)
    {
        if (GamePlayer._IPropDirty.Contains(uiName))
            GamePlayer._IPropDirty.Add(uiName);
    }

    static public void RegCDirty(string uiName)
    {
        if (GamePlayer._CPropDirty.Contains(uiName))
            GamePlayer._CPropDirty.Add(uiName);
    }
}
