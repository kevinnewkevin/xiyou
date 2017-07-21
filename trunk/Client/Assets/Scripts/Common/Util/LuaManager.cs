using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;
using System;

public class LuaManager {

    static public object[] Call(string file, string func, params object[] p)
    {
        LuaState lua = UIManager._Lua;
        lua.DoFile(file);
        LuaFunction luafunc = lua.GetFunction(func);
        object[] rp = luafunc.Call(p);
        lua.CheckTop();
        luafunc.Dispose();
        return rp;
    }

}