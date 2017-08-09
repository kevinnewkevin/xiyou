﻿//this source code was auto-generated by tolua#, do not modify it
using System;
using LuaInterface;

public class BattleWrap
{
	public static void Register(LuaState L)
	{
		L.BeginClass(typeof(Battle), typeof(System.Object));
		L.RegFunction("Update", Update);
		L.RegFunction("Init", Init);
		L.RegFunction("RandHandCards", RandHandCards);
		L.RegFunction("GetActor", GetActor);
		L.RegFunction("BattleSetup", BattleSetup);
		L.RegFunction("Judgement", Judgement);
		L.RegFunction("SwitchPoint", SwitchPoint);
		L.RegFunction("OperateSetActor", OperateSetActor);
		L.RegFunction("IsSelfCard", IsSelfCard);
		L.RegFunction("Fini", Fini);
		L.RegFunction("New", _CreateBattle);
		L.RegFunction("__tostring", ToLua.op_ToString);
		L.RegVar("_CurrentState", get__CurrentState, set__CurrentState);
		L.RegVar("_Result", get__Result, set__Result);
		L.RegVar("_BattleReport", get__BattleReport, set__BattleReport);
		L.RegVar("_ReportAction", get__ReportAction, set__ReportAction);
		L.RegVar("_OperationFinish", get__OperationFinish, set__OperationFinish);
		L.RegVar("_ReportIsPlaying", get__ReportIsPlaying, set__ReportIsPlaying);
		L.RegVar("_SelectedHandCardInstID", get__SelectedHandCardInstID, set__SelectedHandCardInstID);
		L.RegVar("_HandCards", get__HandCards, set__HandCards);
		L.RegVar("_Turn", get__Turn, set__Turn);
		L.RegVar("_Side", get__Side, set__Side);
		L.RegVar("_OperatList", get__OperatList, set__OperatList);
		L.RegVar("BattleReport", null, set_BattleReport);
		L.RegVar("_LeftCardNum", get__LeftCardNum, null);
		L.RegVar("SetResult", null, set_SetResult);
		L.RegVar("CurrentState", get_CurrentState, set_CurrentState);
		L.RegVar("IsWin", get_IsWin, null);
		L.EndClass();
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int _CreateBattle(IntPtr L)
	{
		try
		{
			int count = LuaDLL.lua_gettop(L);

			if (count == 0)
			{
				Battle obj = new Battle();
				ToLua.PushObject(L, obj);
				return 1;
			}
			else
			{
				return LuaDLL.luaL_throw(L, "invalid arguments to ctor method: Battle.New");
			}
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int Update(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 0);
			Battle.Update();
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int Init(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 1);
			Battle.Init(arg0);
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int RandHandCards(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 1);
			Battle.RandHandCards(arg0);
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int GetActor(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			long arg0 = LuaDLL.tolua_checkint64(L, 1);
			Actor o = Battle.GetActor(arg0);
			ToLua.PushObject(L, o);
			return 1;
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
			Battle.BattleSetup();
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int Judgement(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 0);
			Battle.Judgement();
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int SwitchPoint(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			bool arg0 = LuaDLL.luaL_checkboolean(L, 1);
			Battle.SwitchPoint(arg0);
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int OperateSetActor(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 1);
			Battle.OperateSetActor(arg0);
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int IsSelfCard(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 1);
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 1);
			bool o = Battle.IsSelfCard(arg0);
			LuaDLL.lua_pushboolean(L, o);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int Fini(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 0);
			Battle.Fini();
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__CurrentState(IntPtr L)
	{
		try
		{
			ToLua.Push(L, Battle._CurrentState);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__Result(IntPtr L)
	{
		try
		{
			ToLua.Push(L, Battle._Result);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__BattleReport(IntPtr L)
	{
		try
		{
			ToLua.PushObject(L, Battle._BattleReport);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__ReportAction(IntPtr L)
	{
		try
		{
			ToLua.PushObject(L, Battle._ReportAction);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__OperationFinish(IntPtr L)
	{
		try
		{
			LuaDLL.lua_pushboolean(L, Battle._OperationFinish);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__ReportIsPlaying(IntPtr L)
	{
		try
		{
			LuaDLL.lua_pushboolean(L, Battle._ReportIsPlaying);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__SelectedHandCardInstID(IntPtr L)
	{
		try
		{
			LuaDLL.tolua_pushint64(L, Battle._SelectedHandCardInstID);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__HandCards(IntPtr L)
	{
		try
		{
			ToLua.PushObject(L, Battle._HandCards);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__Turn(IntPtr L)
	{
		try
		{
			LuaDLL.lua_pushinteger(L, Battle._Turn);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__Side(IntPtr L)
	{
		try
		{
			LuaDLL.lua_pushinteger(L, Battle._Side);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__OperatList(IntPtr L)
	{
		try
		{
			ToLua.PushObject(L, Battle._OperatList);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get__LeftCardNum(IntPtr L)
	{
		try
		{
			LuaDLL.lua_pushinteger(L, Battle._LeftCardNum);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get_CurrentState(IntPtr L)
	{
		try
		{
			ToLua.Push(L, Battle.CurrentState);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get_IsWin(IntPtr L)
	{
		try
		{
			LuaDLL.lua_pushboolean(L, Battle.IsWin);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__CurrentState(IntPtr L)
	{
		try
		{
			Battle.BattleState arg0 = (Battle.BattleState)ToLua.CheckObject(L, 2, typeof(Battle.BattleState));
			Battle._CurrentState = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__Result(IntPtr L)
	{
		try
		{
			Battle.BattleResult arg0 = (Battle.BattleResult)ToLua.CheckObject(L, 2, typeof(Battle.BattleResult));
			Battle._Result = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__BattleReport(IntPtr L)
	{
		try
		{
			COM_BattleReport arg0 = (COM_BattleReport)ToLua.CheckObject(L, 2, typeof(COM_BattleReport));
			Battle._BattleReport = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__ReportAction(IntPtr L)
	{
		try
		{
			System.Collections.Generic.List<COM_BattleAction> arg0 = (System.Collections.Generic.List<COM_BattleAction>)ToLua.CheckObject(L, 2, typeof(System.Collections.Generic.List<COM_BattleAction>));
			Battle._ReportAction = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__OperationFinish(IntPtr L)
	{
		try
		{
			bool arg0 = LuaDLL.luaL_checkboolean(L, 2);
			Battle._OperationFinish = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__ReportIsPlaying(IntPtr L)
	{
		try
		{
			bool arg0 = LuaDLL.luaL_checkboolean(L, 2);
			Battle._ReportIsPlaying = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__SelectedHandCardInstID(IntPtr L)
	{
		try
		{
			long arg0 = LuaDLL.tolua_checkint64(L, 2);
			Battle._SelectedHandCardInstID = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__HandCards(IntPtr L)
	{
		try
		{
			System.Collections.Generic.List<COM_Unit> arg0 = (System.Collections.Generic.List<COM_Unit>)ToLua.CheckObject(L, 2, typeof(System.Collections.Generic.List<COM_Unit>));
			Battle._HandCards = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__Turn(IntPtr L)
	{
		try
		{
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 2);
			Battle._Turn = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__Side(IntPtr L)
	{
		try
		{
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 2);
			Battle._Side = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set__OperatList(IntPtr L)
	{
		try
		{
			System.Collections.Generic.List<COM_BattlePosition> arg0 = (System.Collections.Generic.List<COM_BattlePosition>)ToLua.CheckObject(L, 2, typeof(System.Collections.Generic.List<COM_BattlePosition>));
			Battle._OperatList = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set_BattleReport(IntPtr L)
	{
		try
		{
			COM_BattleReport arg0 = (COM_BattleReport)ToLua.CheckObject(L, 2, typeof(COM_BattleReport));
			Battle.BattleReport = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set_SetResult(IntPtr L)
	{
		try
		{
			Battle.BattleResult arg0 = (Battle.BattleResult)ToLua.CheckObject(L, 2, typeof(Battle.BattleResult));
			Battle.SetResult = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int set_CurrentState(IntPtr L)
	{
		try
		{
			Battle.BattleState arg0 = (Battle.BattleState)ToLua.CheckObject(L, 2, typeof(Battle.BattleState));
			Battle.CurrentState = arg0;
			return 0;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}
}

