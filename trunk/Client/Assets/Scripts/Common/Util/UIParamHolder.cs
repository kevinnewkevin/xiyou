using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class UIParamHolder {

    static Dictionary<string, object> _UiparamDic = new Dictionary<string, object>();

    static public void Set(string key, object val)
    {
        if (_UiparamDic.ContainsKey(key))
            _UiparamDic [key] = val;
        else
            _UiparamDic.Add(key, val);
    }

    static public object Get(string key)
    {
        if (_UiparamDic.ContainsKey(key))
            return _UiparamDic [key];
        else
            return null;
    }
}
