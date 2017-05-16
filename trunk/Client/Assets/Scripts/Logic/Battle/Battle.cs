using UnityEngine;
using System.Collections.Generic;

public class Battle {

    public enum BattleState
    {
        BS_Init,
        BS_Oper,
        BS_Play,
        BS_Result,
        BS_Max
    }

    public enum BattleResult
    {
        BR_Win,
        BR_Lose,
        BR_None
    }

    static Actor[] _ActorInScene = new Actor[12/*BP_Max*/];

    static public BattleState _CurrentState = BattleState.BS_Max;
    static public BattleResult _Result = BattleResult.BR_None;

    static public bool _OperationFinish;

    static public void Update()
    {
        switch(_CurrentState)
        {
            case BattleState.BS_Init:
                if (LoadAssets() && PlaceActor())
                    _CurrentState = BattleState.BS_Oper;
                break;
            case BattleState.BS_Oper:
                if (_OperationFinish)
                    _CurrentState = BattleState.BS_Play;
                break;
            case BattleState.BS_Play:
                Play();
                break;
            case BattleState.BS_Result:
                break;
            default:
                break;
        }
    }

    //初始化战斗
    static public void Init()
    {
        _CurrentState = BattleState.BS_Init;
    }

    static bool LoadAssets()
    {
        return false;
    }

    static bool PlaceActor()
    {
        return false;
    }

    static void UnLoadAssets()
    {
        for (int i = 0; i < _ActorInScene.Length; ++i)
        {
            if (_ActorInScene[i] != null)
                _ActorInScene[i].Fini();
        }
    }

    //播放一回合战报 处理快照
    static void Play()
    {
        //if battleReport is null return;

        //if is playing return;

        //if final report play to end;
        Judgement();
    }

    //场上添加一个角色
    static void AddActor(GameObject go/*, BattlePosition pos*/, int pos)
    {
        DelActor(pos);
        _ActorInScene[pos] = new Actor(go);
    }

    //场上删除一个角色
    static void DelActor(/*, BattlePosition pos*/ int pos)
    {
        if (_ActorInScene[pos] != null)
            _ActorInScene[pos].Fini();
        _ActorInScene[pos] = null;
    }

    //场上找到一个角色
    static public Actor GetActor(/*, BattlePosition pos*/ int pos)
    {
        return _ActorInScene[pos];
    }

    static public void Judgement()
    {
        if (_Result == BattleResult.BR_None)
            _CurrentState = BattleState.BS_Oper;
        else
            _CurrentState = BattleState.BS_Result;
    }

    //销毁场景 角色 UI
    static public void Fini()
    {
        _CurrentState = BattleState.BS_Max;
        _Result = BattleResult.BR_None;
        UnLoadAssets();
    }
}
