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

    static Actor[] _SelfActorInScene = new Actor[6/*BP_Max*/];
    static Actor[] _OppoActorInScene = new Actor[6/*BP_Max*/];
    static Transform[] _SelfPosInScene = new Transform[6/*BP_Max*/];
    static Transform[] _OppoPosInScene = new Transform[6/*BP_Max*/];
    static GameObject _SceneConfig;

    static public BattleState _CurrentState = BattleState.BS_Max;
    static public BattleResult _Result = BattleResult.BR_None;
    static public protocol.COM_BattleReport _BattleReport;

    static bool _IsStagePointInitSuc;
    static public bool _OperationFinish;
    static public bool _ReportIsPlaying;    //仅代表每一个战报单元的状态

    static public ulong _SelectedHandCardInstID;
    static public List<protocol.COM_EntityInstance> _HandCards = new List<protocol.COM_EntityInstance>();

    static public List<protocol.COM_BattlePosition> _OperatList = new List<protocol.COM_BattlePosition>();

    static public void Update()
    {
        switch(_CurrentState)
        {
            case BattleState.BS_Init:
                if (LoadAssets() && PlaceActor())
                {
                    _CurrentState = BattleState.BS_Oper;
                }
                break;
            case BattleState.BS_Oper:
                if (_OperationFinish)
                {
                    _OperatList.Clear();
                    _CurrentState = BattleState.BS_Play;
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
    static public void Init()
    {
        _CurrentState = BattleState.BS_Init;
        _OperatList.Clear();
        _HandCards.Clear();

        _HandCards.Add(GamePlayer._Data);
        RandHandCards(2);
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
                    if (toIdx < 6)
                    {
                        _OppoPosInScene [toIdx] = point;
                        _OppoPosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                        _OppoPosInScene [toIdx].gameObject.SetActive(false);
                    }
                    else
                    {
                        _SelfPosInScene [toIdx - 6] = point;
                        _SelfPosInScene [toIdx - 6].GetComponent<PointHandle>().Init(toIdx - 6);
                        _SelfPosInScene [toIdx - 6].gameObject.SetActive(false);
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

        if (_BattleReport.BattleUnit.Count > 0)
        {
            DisplayData display;
            for(int i=0; i < _BattleReport.BattleUnit.Count; ++i)
            {
                display = DisplayData.GetData(_BattleReport.BattleUnit[i].EntityId);
                AddActor(AssetLoader.LoadAsset(display._AssetPath), _BattleReport.BattleUnit[i].PositionType, _BattleReport.BattleUnit[i].InstanceId);
            }
            _BattleReport.BattleUnit.Clear();
            return;
        }

        if (_ReportIsPlaying)
            return;

        // cast skill
        Actor actor = GetActor(_BattleReport.BattleAction[0].InstanceId);
        List<Actor> targets = new List<Actor>();
        Actor target;
        for(int i=0; i < _BattleReport.BattleAction[0].BattleTarget.Count; ++i)
        {
            target = GetActor(_BattleReport.BattleAction[0].BattleTarget[i].InstanceId);
            targets.Add(target);
        }
        Skill skill = new Skill(_BattleReport.BattleAction[0].SkillId, actor, targets.ToArray());
        skill.Cast();

        _BattleReport.BattleAction.RemoveAt(0);

        _ReportIsPlaying = true;

        //if final report play to end;
        if(_BattleReport.BattleAction.Count == 0)
            Judgement();
    }

    //场上添加一个角色
    static void AddActor(GameObject go, int pos, ulong instid)
    {
        Actor actor = GetActor(instid);
        if (actor != null)
            return;
        
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
    static public Actor GetActor(ulong instid)
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
        NetWoking.S.BattleSetup(Battle._OperatList);
        Battle._OperatList.Clear();
        _CurrentState = BattleState.BS_Play;
    }

    static public void Judgement()
    {
        if (_Result == BattleResult.BR_None)
            _CurrentState = BattleState.BS_Oper;
        else
            _CurrentState = BattleState.BS_Result;
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
            if (_OperatList [i].InstanceId == _SelectedHandCardInstID)
            {
                _OperatList [i].PosotionId = (sbyte)pos;
                contains = true;
                break;
            }
        }

        if(!contains)
        {
            protocol.COM_BattlePosition bp = new protocol.COM_BattlePosition();
            bp.InstanceId = _SelectedHandCardInstID;
            bp.PosotionId = (sbyte)pos;
            _OperatList.Add(bp);

            protocol.COM_EntityInstance entity = GamePlayer.GetCardByInstID(_SelectedHandCardInstID);
            int displayId = 1;

            DisplayData displayData = DisplayData.GetData(displayId);
            AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, _SelectedHandCardInstID);
        }

        SwitchPoint(false);
    }

    //销毁场景 角色 UI
    static public void Fini()
    {
        _CurrentState = BattleState.BS_Max;
        _Result = BattleResult.BR_None;
        _ReportIsPlaying = false;
        UnLoadAssets();
    }
}
