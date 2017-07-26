﻿using UnityEngine;
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

    static Actor[] _SelfActorInScene = new Actor[12/*BP_Max*/];
    static Actor[] _OppoActorInScene = new Actor[12/*BP_Max*/];
    static Transform[] _SelfPosInScene = new Transform[12/*BP_Max*/];
    static Transform[] _OppoPosInScene = new Transform[12/*BP_Max*/];
    static GameObject _SceneConfig;

    static public BattleState _CurrentState = BattleState.BS_Max;
    static public BattleResult _Result = BattleResult.BR_None;
    static public COM_BattleReport _BattleReport;

    static bool _IsStagePointInitSuc;
    static public bool _OperationFinish;
    static public bool _ReportIsPlaying;    //仅代表每一个战报单元的状态

    static public long _SelectedHandCardInstID;
    static public List<COM_Unit> _HandCards = new List<COM_Unit>();

    static public int _Turn;

    static public int _Side;

    static public int _LeftCardNum
    {
        get
        {
            return _HandCards.Count;
        }
    }

    static public List<COM_BattlePosition> _OperatList = new List<COM_BattlePosition>();

    static public void Update()
    {
        switch(_CurrentState)
        {
            case BattleState.BS_Init:
                if (LoadAssets() && PlaceActor())
                {
                    CurrentState = BattleState.BS_Oper;
                }
                break;
            case BattleState.BS_Oper:
                if (_OperationFinish)
                {
                    _OperatList.Clear();
                    CurrentState = BattleState.BS_Play;
                }
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
    static public void Init(int side)
    {
        _IsStagePointInitSuc = false;
        _ReportIsPlaying = false;
        _Turn = 1;
        _Side = side;
        CurrentState = BattleState.BS_Init;
        _OperatList.Clear();
        _HandCards.Clear();

        _HandCards.Add(GamePlayer._Data);
        RandHandCards(2);

        UIManager.SetDirty("BattlePanel");
    }

    static public void RandHandCards(int count)
    {
        while(count > 0)
        {
            _HandCards.Add(GamePlayer._Cards[Random.Range(0, GamePlayer._Cards.Count)]);
            count--;
        }
    }

    static bool LoadAssets()
    {
        // 加载场景站位点信息
        if (_IsStagePointInitSuc == false)
        {
            _SceneConfig = GameObject.Find("SceneConfig");
            if (_SceneConfig != null)
            {
                Transform point;
                for(int i=0; i < _SceneConfig.transform.childCount; ++i)
                {
                    point = _SceneConfig.transform.GetChild(i);
                    int toIdx = int.Parse(point.name) - 1;
                    if (_Side == 0)
                    {
                        if (toIdx < 6)
                        {
                            _SelfPosInScene [toIdx] = point;
                            _SelfPosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                            _SelfPosInScene [toIdx].gameObject.SetActive(false);
                        }
                        else
                        {
                            _OppoPosInScene [toIdx] = point;
                            _OppoPosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                            _OppoPosInScene [toIdx].gameObject.SetActive(false);
                        }
                    }
                    else
                    {
                        if (toIdx < 6)
                        {
                            _OppoPosInScene [toIdx] = point;
                            _OppoPosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                            _OppoPosInScene [toIdx].gameObject.SetActive(false);
                        }
                        else
                        {
                            _SelfPosInScene [toIdx] = point;
                            _SelfPosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                            _SelfPosInScene [toIdx].gameObject.SetActive(false);
                        }
                    }
                }
                _IsStagePointInitSuc = true;
            }
        }

        // 加载角色资源
        //TODO

        return _IsStagePointInitSuc;
    }

    static bool PlaceActor()
    {
        return true;
    }

    static void UnLoadAssets()
    {
        for (int i = 0; i < _SelfActorInScene.Length; ++i)
        {
            if (_SelfActorInScene[i] != null)
                _SelfActorInScene[i].Fini();
        }

        for (int i = 0; i < _OppoActorInScene.Length; ++i)
        {
            if (_OppoActorInScene[i] != null)
                _OppoActorInScene[i].Fini();
        }
    }

    //播放一回合战报 处理快照
    static void Play()
    {
        if (_BattleReport == null)
            return;

        if (_BattleReport.UnitList != null && _BattleReport.UnitList.Length > 0)
        {
            EntityData entity;
            DisplayData display;
            for (int i = 0; i < _BattleReport.UnitList.Length; ++i)
            {
                entity = EntityData.GetData(_BattleReport.UnitList[i].UnitId);
                display = DisplayData.GetData(entity._DisplayId);
                AddActor(AssetLoader.LoadAsset(display._AssetPath), _BattleReport.UnitList[i].Position, _BattleReport.UnitList[i].InstId);
            }
            _BattleReport.UnitList = null;
            return;
        }

        if (_ReportIsPlaying)
            return;

        // cast skill
        Actor actor = GetActor(_BattleReport.ActionList[0].InstId);
        List<Actor> targets = new List<Actor>();
        Actor target;
        for (int i = 0; i < _BattleReport.ActionList[0].TargetList.Length; ++i)
        {
            target = GetActor(_BattleReport.ActionList[0].TargetList[i].InstId);
            targets.Add(target);
        }
        Skill skill = new Skill(_BattleReport.ActionList[0].SkillId, actor, targets.ToArray());
        skill.Cast();

       System.Array.Copy( _BattleReport.ActionList,1, _BattleReport.ActionList, 0, _BattleReport.ActionList.Length - 1);

        _ReportIsPlaying = true;

        //if final report play to end;
        if (_BattleReport.ActionList.Length == 0)
            Judgement();
    }

    //场上添加一个角色
    static void AddActor(GameObject go, int pos, long instid)
    {
        Actor actor = GetActor(instid);
        if (actor != null)
        {
            if(GamePlayer.IsMy(instid))
                actor.MoveTo(_SelfPosInScene[pos].position, null);
            else
                actor.MoveTo(_OppoPosInScene[pos].position, null);
            return;
        }
        
        if(GamePlayer.IsMy(instid))
            _SelfActorInScene[pos] = new Actor(go, _SelfPosInScene[pos].position, instid);
        else
            _OppoActorInScene[pos] = new Actor(go, _OppoPosInScene[pos].position, instid);
    }

    //场上删除一个角色
    static void DelActor(int pos, bool self)
    {
        if (self)
        {
            if (_SelfActorInScene [pos] != null)
                _SelfActorInScene [pos].Fini();
            _SelfActorInScene [pos] = null;
        }
        else
        {
            if (_OppoActorInScene [pos] != null)
                _OppoActorInScene [pos].Fini();
            _OppoActorInScene [pos] = null;
        }
    }

    //场上找到一个角色
    static public Actor GetActor(long instid)
    {
        if (GamePlayer.IsMy(instid))
        {
            for (int i = 0; i < _SelfActorInScene.Length; ++i)
            {
                if (_SelfActorInScene [i] == null)
                    continue;

                if (_SelfActorInScene [i].InstID == instid)
                    return _SelfActorInScene [i];
            }
        }
        else
        {
            for (int i = 0; i < _OppoActorInScene.Length; ++i)
            {
                if (_OppoActorInScene [i] == null)
                    continue;

                if (_OppoActorInScene [i].InstID == instid)
                    return _OppoActorInScene [i];
            }
        }

        return null;
    }

    static public void BattleSetup()
    {
        NetWoking.S.SetupBattle(Battle._OperatList.ToArray());
        Battle._OperatList.Clear();
        CurrentState = BattleState.BS_Play;
    }

    static public void Judgement()
    {
        if (_Result == BattleResult.BR_None)
            CurrentState = BattleState.BS_Oper;
        else
            CurrentState = BattleState.BS_Result;

        _BattleReport = null;
    }

    static public void SwitchPoint(bool on)
    {
        for(int i=0; i < _SelfPosInScene.Length; ++i)
        {
            if (_SelfPosInScene [i] != null)
            {
                _SelfPosInScene [i].gameObject.SetActive(on);
            }
        }
//        for(int i=0; i < _OppoPosInScene.Length; ++i)
//        {
//            if (_OppoPosInScene [i] != null)
//            {
//                _OppoPosInScene [i].gameObject.SetActive(on);
//            }
//        }
    }

    static public void OperateSetActor(int pos)
    {
        bool contains = false;
        for(int i=0; i < _OperatList.Count; ++i)
        {
            if (_OperatList [i].InstId == _SelectedHandCardInstID)
            {
                _OperatList [i].Position = pos;
                contains = true;
                break;
            }
        }

        if(!contains)
        {
            COM_BattlePosition bp = new COM_BattlePosition();
            bp.InstId = _SelectedHandCardInstID;
            bp.Position = pos;
            _OperatList.Add(bp);

            COM_Unit entity = GamePlayer.GetCardByInstID(_SelectedHandCardInstID);
            int displayId = 1;

            DisplayData displayData = DisplayData.GetData(displayId);
            AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, _SelectedHandCardInstID);
        }

        SwitchPoint(false);
    }

    static public bool IsSelfCard(int cardIdx)
    {
        if (cardIdx < 0 || cardIdx >= _HandCards.Count)
            return false;
        return GamePlayer.IsMe(_HandCards[cardIdx].InstId);
    }

    static public BattleState CurrentState
    {
        set
        {
            _CurrentState = value;
            UIManager.SetDirty("BattlePanel");
        }
        get { return _CurrentState; }
    }


    //销毁场景 角色 UI
    static public void Fini()
    {
        CurrentState = BattleState.BS_Max;
        _Result = BattleResult.BR_None;
        _ReportIsPlaying = false;
        UnLoadAssets();
    }
}
