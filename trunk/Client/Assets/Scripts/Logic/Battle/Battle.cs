using UnityEngine;
using System.Collections.Generic;
using FairyGUI;

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
        BR_Lose,
        BR_Win,
        BR_None
    }

    static Actor[] _ActorInScene = new Actor[12/*BP_Max*/];
    static Transform[] _PosInScene = new Transform[12/*BP_Max*/];
    static GameObject _SceneConfig;

    static public BattleState _CurrentState = BattleState.BS_Max;
    static public BattleResult _Result = BattleResult.BR_None;
    static public COM_BattleReport _BattleReport;
    static public List<COM_BattleAction> _ReportAction;

    static bool _IsStagePointInitSuc;
    static public bool _OperationFinish;
    static public bool _ReportIsPlaying;    //仅代表每一个战报单元的状态
    static public bool _ActorLaunched;

    static BuffChecker _CrtBuffChecker;
    static Skill _CrtSkill;
    static Actor _CrtActor;
    static List<Actor> _CrtTargets;

    static public long _SelectedHandCardInstID;
    static public List<COM_Unit> _HandCards = new List<COM_Unit>();

    static public int _Turn;

    static public int _Fee;

    static public int _MaxFee;

    static public int _Side;

    static public List<COM_Unit> _MyGroupCards;

    static public COM_BattleReport BattleReport
    {
        set
        {
            _BattleReport = value;
            _ReportAction = new List<COM_BattleAction>(_BattleReport.ActionList);
        }
    }

    static public int _LeftCardNum
    {
        get
        {
            return _HandCards.Count;
        }
    }

    static public BattleResult SetResult
    {
        set
        {
            _Result = value;
            if (_ReportAction == null || _ReportAction.Count == 0)
                Judgement();
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
                if (GamePlayer._IsAuto)
                {
                    CurrentState = BattleState.BS_Play;
                    PutCardInBattle();
                }
                else
                {
                    if (_OperationFinish)
                    {
                        _OperatList.Clear();
                        CurrentState = BattleState.BS_Play;
                    }
                }
                break;
            case BattleState.BS_Play:
                Play();
                break;
            case BattleState.BS_Result:
                End();
                break;
            default:
                break;
        }

        for(int i=0; i < _ActorInScene.Length; ++i)
        {
            if (_ActorInScene[i] != null)
                _ActorInScene [i].Update();
        }
    }

    //初始化战斗
    static public void Init(int side)
    {
        _SceneConfig = null;
        _IsStagePointInitSuc = false;
        _ReportIsPlaying = false;
        _Turn = 1;
        _Side = side;
        CurrentState = BattleState.BS_Init;
        _OperatList.Clear();
        _HandCards.Clear();

        _HandCards.Add(GamePlayer._Data);
        _MyGroupCards = GamePlayer.GetBattleCardsCopy();

        _MaxFee = Define.GetInt("MaxFee");

        UIManager.SetDirty("BattlePanel");
    }

    static public void RandHandCards(int count)
    {
        if (_MyGroupCards.Count <= 0)
            return;
        
        if (count <= 0)
            return;

        int idx = Random.Range(0, _MyGroupCards.Count);
        if (idx < 0 || idx >= _MyGroupCards.Count)
            return;
        
        _HandCards.Add(_MyGroupCards[idx]);
        _MyGroupCards.RemoveAt(idx);
        RandHandCards(--count);
    }

    static public void PutCardInBattle()
    {
        EntityData entity = null;
        for(int i=0; i < _HandCards.Count; ++i)
        {
            entity = EntityData.GetData(_HandCards [i].UnitId);
            if (entity != null && entity._Cost <= _Fee)
            {
                _SelectedHandCardInstID = _HandCards [i].InstId;
                OperateSetActor(FindEmptyPos());
            }
        }
        BattleSetup();
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
                        _PosInScene [toIdx] = point;
                        _PosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                        _PosInScene [toIdx].gameObject.SetActive(false);
                    }
                    else
                    {
                        _PosInScene [toIdx] = point;
                        _PosInScene [toIdx].GetComponent<PointHandle>().Init(ConvertedPos(toIdx));
                        _PosInScene [toIdx].gameObject.SetActive(false);
                    }
                }
                _IsStagePointInitSuc = true;
            }
        }

        // 加载角色资源
        //TODO

        return _IsStagePointInitSuc;
    }

    static int ConvertedPos(int pos)
    {
        return (pos + 6) % 12;
    }

    static bool PlaceActor()
    {
        return true;
    }

    static void UnLoadAssets()
    {
        for (int i = 0; i < _PosInScene.Length; ++i)
        {
            if (_ActorInScene[i] != null)
                _ActorInScene[i].Fini();
        }
        _PosInScene = new Transform[12];
        _ActorInScene = new Actor[12];
    }

    //播放一回合战报 处理快照
    static void Play()
    {
        if (_ReportAction == null || _ReportAction.Count == 0)
            return;

        // 处理每回合新上场角色
        if (_BattleReport.UnitList != null && _BattleReport.UnitList.Length > 0)
        {
            EntityData entity;
            DisplayData display;
            Actor actor;
            for (int i = 0; i < _BattleReport.UnitList.Length; ++i)
            {
                actor = GetActor(_BattleReport.UnitList[i].InstId);
                if (actor != null)
                {
                    actor.SetValue(_BattleReport.UnitList[i].CHP, _BattleReport.UnitList[i].HP);
                    continue;
                }
                
                entity = EntityData.GetData(_BattleReport.UnitList[i].UnitId);
                display = DisplayData.GetData(entity._DisplayId);
                AddActor(AssetLoader.LoadAsset(display._AssetPath), _BattleReport.UnitList[i].Position, _BattleReport.UnitList[i].InstId, _BattleReport.UnitList[i].CHP, _BattleReport.UnitList[i].HP);
            }
            _BattleReport.UnitList = null;
            return;
        }


        if (!_ActorLaunched)
        {
            // 获取该次行动的施法者 目标
            _CrtActor = GetActor(_ReportAction [0].InstId);
            _CrtTargets = new List<Actor>();
            Actor target;
            for (int i = 0; i < _ReportAction [0].TargetList.Length; ++i)
            {
                target = GetActor(_ReportAction [0].TargetList [i].InstId);
                _CrtTargets.Add(target);
            }
            _ActorLaunched = true;
        }

        if (_CrtBuffChecker == null)
        {
            // 处理buff结算 和自身buff的增删
            _CrtBuffChecker = new BuffChecker(_CrtActor, _ReportAction [0].BuffList);
            _CrtBuffChecker.Check();
            return;
        }
        else if (!_CrtBuffChecker._IsChecked)
        {
            return;
        }

        // cast skill
        if (_CrtSkill == null)
        {
            _CrtSkill = new Skill(_ReportAction [0].SkillId, _CrtActor, _CrtTargets.ToArray(), _ReportAction [0].TargetList, _ReportAction[0].SkillBuff);
            _CrtSkill.Cast();
            return;
        }
        else if(!_CrtSkill._IsCasted)
        {
            return;
        }

        _ReportAction.RemoveAt(0);
        CheckEnd();
    }

    static public void CheckEnd()
    {
        if (_ReportAction == null || _ReportAction.Count == 0)
            Judgement();

        _ActorLaunched = false;
        _CrtBuffChecker = null;
        _CrtSkill = null;
    }

    static void End()
    {
        if (_BattleReport != null)
        {
            CurrentState = BattleState.BS_Play;
            return;
        }

        Debug.Log(" You" + (_Result == BattleResult.BR_Win? " Win ": " Lose"));
        UIManager.Show("jiesuanjiemian");
        Fini();
    }

    //场上添加一个角色
    static void AddActor(GameObject go, int pos, long instid, int crtHp, int maxHp)
    {
        int tpos = _Side == 0? pos: ConvertedPos(pos);
        Actor actor = GetActor(instid);
        if (actor != null)
        {
            actor.MoveTo(_PosInScene [tpos].position, null);
            return;
        }
        _ActorInScene[tpos] = new Actor(go, _PosInScene[tpos], instid, pos, crtHp, maxHp);
        _ActorInScene[tpos].Play(Define.ANIMATION_PLAYER_ACTION_SHOW);
        _ActorInScene [tpos].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);

        UIManager.SetDirty("BattlePanel");
    }

    //场上删除一个角色
    static void DelActor(int pos)
    {
        int tpos = _Side == 0? pos: ConvertedPos(pos);
        if (_ActorInScene [tpos] != null)
            _ActorInScene [tpos].Fini();
        _ActorInScene [tpos] = null;
    }

    //场上找到一个角色
    static public Actor GetActor(long instid)
    {
        for (int i = 0; i < _ActorInScene.Length; ++i)
        {
            if (_ActorInScene [i] == null)
                continue;

            if (_ActorInScene [i].InstID == instid)
                return _ActorInScene [i];
        }
        return null;
    }

    static public int FindEmptyPos()
    {
        int emptyPos = -1;
        for(int i = 0; i < _PosInScene.Length; ++i)
        {
            emptyPos = i;
            for (int j = 0; j < _ActorInScene.Length; ++j)
            {
                if (_ActorInScene [j] == null)
                    continue;

                if (_ActorInScene [j]._RealPosInScene == i)
                {
                    emptyPos = -1;
                    break;
                }
            }
            if (emptyPos != -1)
                return (_Side == 0? emptyPos: ConvertedPos(emptyPos));
        }
        return (_Side == 0? emptyPos: ConvertedPos(emptyPos));
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

        // 从第二回合开始 每回合结束加 1 费
        _Turn++;
        if (_Turn > 1)
            AddFee(1);

        if (_Turn == 2)
            RandHandCards(3);
        if(_Turn > 2)
            RandHandCards(1);
    }

    static public void SwitchPoint(bool on)
    {
        for(int i=0; i < 6; ++i)
        {
            if (_PosInScene [i] != null)
            {
                _PosInScene [i].gameObject.SetActive(on);
            }
        }
    }

    static public void OperateSetActor(int pos)
    {
        if (pos == -1)
            return;
        
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
            EntityData eData = EntityData.GetData(entity.UnitId);
            DisplayData displayData = DisplayData.GetData(eData._DisplayId);
            AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, _SelectedHandCardInstID, 100, 100);
            RemoveHandCard(_SelectedHandCardInstID);
            CostFee(eData._Cost);
        }

        SwitchPoint(false);
    }

    static public EntityData GetHandCard(int idx)
    {
        if (idx < 0 || idx >= _HandCards.Count)
            return null;

        EntityData eData = EntityData.GetData(_HandCards[idx].UnitId);
        
        return eData;
    }

    static public DisplayData GetHandCardDisplay(int idx)
    {
        if (idx < 0 || idx >= _HandCards.Count)
            return null;

        EntityData eData = GetHandCard(idx);
        if (eData == null)
            return null;

        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        return dData;
    }

    static public void RemoveHandCard(long instid)
    {
        for(int i=0; i < _HandCards.Count; ++i)
        {
            if (_HandCards [i].InstId == instid)
            {
                _HandCards.RemoveAt(i);
                break;
            }
        }
    }

    static public bool IsSelfCard(int cardIdx)
    {
        if (cardIdx < 0 || cardIdx >= _HandCards.Count)
            return false;
        return GamePlayer.IsMe(_HandCards[cardIdx].InstId);
    }

    static public int CardsInGroupCount
    {
        get
        {
            if (_MyGroupCards == null)
                return 0;
            
            return _MyGroupCards.Count;
        }
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

    static public bool IsWin
    {
        get
        {
            return _Result == BattleResult.BR_Win;
        }
    }

    static public void AddFee(int count)
    {
        _Fee += count;
        if (_Fee > _MaxFee)
            _Fee = _MaxFee;

        UIManager.SetDirty("BattlePanel");
    }

    static public bool CostFee(int count)
    {
        if (count > _MaxFee)
            return false;
        
        if (_Fee < count)
            return false;
        
        _Fee -= count;
        UIManager.SetDirty("BattlePanel");
        return true;
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
