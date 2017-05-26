﻿//this source code was auto-generated by tolua#, do not modify it
using System;
using LuaInterface;

public class Proxy4LuaWrap
{
	public static void Register(LuaState L)
	{
		L.BeginClass(typeof(Proxy4Lua), typeof(System.Object));
		L.RegFunction("BattleSetup", BattleSetup);
		L.RegFunction("BattleJoin", BattleJoin);
		L.RegFunction("SelectCard4Ready", SelectCard4Ready);
		L.RegFunction("New", _CreateProxy4Lua);
		L.RegFunction("__tostring", ToLua.op_ToString);
		L.EndClass();
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int _CreateProxy4Lua(IntPtr L)
	{
		try
		{
			int count = LuaDLL.lua_gettop(L);

			if (count == 0)
			{
				Proxy4Lua obj = new Proxy4Lua();
				ToLua.PushObject(L, obj);
				return 1;
			}
			else
			{
				return LuaDLL.luaL_throw(L, "invalid arguments to ctor method: Proxy4Lua.New");
			}
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int BattleSetup(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 0);
			Proxy4Lua.BattleSetup();
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int BattleJoin(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 0);
			Proxy4Lua.BattleJoin();
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int SelectCard4Ready(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 1);
			Proxy4Lua.SelectCard4Ready(arg0);
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}
}

