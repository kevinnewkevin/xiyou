using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;
using System;

public class LuaManager {

    static LuaState _GlobalLua;

    static public void Init()
    {
        _GlobalLua = new LuaState();
        LuaBinder.Bind(_GlobalLua);
        _GlobalLua.Start();
        _GlobalLua.DoFile("global.lua");
    }

    static public object[] CallGlobal(string func, params object[] p)
    {
        LuaFunction luafunc = _GlobalLua.GetFunction(func);
        object[] rp = luafunc.Call(p);
        _GlobalLua.CheckTop();
        luafunc.Dispose();
        return rp;
    }

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