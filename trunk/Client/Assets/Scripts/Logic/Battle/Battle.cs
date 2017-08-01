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
        BR_Lose,
        BR_Win,
        BR_None
    }

    static Actor[] _ActorInScene = new Actor[12/*BP_Max*/];
    //static Actor[] _OppoActorInScene = new Actor[12/*BP_Max*/];
    static Transform[] _PosInScene = new Transform[12/*BP_Max*/];
    //static Transform[] _OppoPosInScene = new Transform[12/*BP_Max*/];
    static GameObject _SceneConfig;

    static public BattleState _CurrentState = BattleState.BS_Max;
    static public BattleResult _Result = BattleResult.BR_None;
    static public COM_BattleReport _BattleReport;
    static public List<COM_BattleAction> _ReportAction;

    static bool _IsStagePointInitSuc;
    static public bool _OperationFinish;
    static public bool _ReportIsPlaying;    //仅代表每一个战报单元的状态

    static public long _SelectedHandCardInstID;
    static public List<COM_Unit> _HandCards = new List<COM_Unit>();

    static public int _Turn;

    static public int _Side;

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
                End();
                break;
            default:
                break;
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

        if (_BattleReport.UnitList != null && _BattleReport.UnitList.Length > 0)
        {
            EntityData entity;
            DisplayData display;
            for (int i = 0; i < _BattleReport.UnitList.Length; ++i)
            {
                if (GetActor(_BattleReport.UnitList [i].InstId) != null)
                    continue;
                
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
        Actor actor = GetActor(_ReportAction[0].InstId);
        List<Actor> targets = new List<Actor>();
        Actor target;
        for (int i = 0; i < _ReportAction[0].TargetList.Length; ++i)
        {
            target = GetActor(_ReportAction[0].TargetList[i].InstId);
            targets.Add(target);
        }
        Skill skill = new Skill(_ReportAction[0].SkillId, actor, targets.ToArray());
        skill.Cast();

        _ReportAction.RemoveAt(0);

        _ReportIsPlaying = true;

        //if final report play to end;
        if (_ReportAction.Count == 0)
            Judgement();
    }

    static void End()
    {
        if (_ReportIsPlaying)
            return;

        if (_BattleReport != null)
            return;

        Debug.Log(" You" + (_Result == BattleResult.BR_Win? " Win ": " Lose"));
        //SceneLoader.LoadScene("main");
        UIManager.Show("shengli_Component");
        Fini();
    }

    //场上添加一个角色
    static void AddActor(GameObject go, int pos, long instid)
    {
        int tpos = _Side == 0? pos: ConvertedPos(pos);
        Actor actor = GetActor(instid);
        if (actor != null)
        {
            actor.MoveTo(_PosInScene[tpos].position, null);
            return;
        }
        _ActorInScene[tpos] = new Actor(go, _PosInScene[tpos], instid);
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
