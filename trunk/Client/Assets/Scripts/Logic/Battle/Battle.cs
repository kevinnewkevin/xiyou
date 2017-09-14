using UnityEngine;
using System.Collections.Generic;
using FairyGUI;

public class Battle {

    public enum BattleState
    {
        BS_Init,
        BS_Opra,
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
    static Dictionary<int, SkillData> skillAssets;

    static public BattleState _CurrentState = BattleState.BS_Max;
    static public COM_BattleResult _Result = null;
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

    static float _LongestShowTime;
    static bool _ShowTimeDoing;
    static bool _IsEnding;

    static public long _SelectedHandCardInstID;
    static public List<COM_Unit> _HandCards = new List<COM_Unit>();

    static public int _Turn;

    static public int _Fee;

    static public int _MaxFee;

    static public int _Side;

    static public int _BattleId;

    static public int _CasterDisplayID;

    static public Vector3 _Center;

    static public List<long> _MyGroupCards;

    static public BattleCamera _BattleCamera;

    static public int[] _OpponentCards;

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

    static public COM_BattleResult Result
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
                    BattleData bData = BattleData.GetData(_BattleId);
                    // battle has anim
                    if (bData != null && bData._Animations != null && bData._Animations.Length > 0)
                    {
                        GRoot.inst.modalLayer.visible = false;
                        op.Begin(bData._Animations);
                        op.Play();
                        CurrentState = BattleState.BS_Opra;
                    }
                    else
                    {
                        CurrentState = BattleState.BS_Oper;
                        UIManager.Show("BattlePanel");
                    }
                }
                break;
            case BattleState.BS_Opra:
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
    static public void Init(int side, int battleid = 0, int[] opponentCards = null)
    {
        _OpponentCards = opponentCards;
        _SceneConfig = null;
        _IsStagePointInitSuc = false;
        _ReportIsPlaying = false;
        _Turn = 1;
        _Side = side;
        _BattleId = battleid;
        CurrentState = BattleState.BS_Init;
        _OperatList.Clear();
        _HandCards.Clear();

        _HandCards.Add(GamePlayer._Data);
        _MyGroupCards = GamePlayer.GetBattleCardsCopy();

        _MaxFee = Define.GetInt("MaxFee");
    }

    static public void RandHandCards(int count)
    {
        if (_MyGroupCards.Count <= 0)
            return;

        if (_HandCards.Count >= 5)
            return;
        
        if (count <= 0)
            return;

        int idx = Random.Range(0, _MyGroupCards.Count);
        if (idx < 0 || idx >= _MyGroupCards.Count)
            return;

        COM_Unit unit = GamePlayer.GetCardByInstID(_MyGroupCards[idx]);
        _HandCards.Add(unit);
        _MyGroupCards.RemoveAt(idx);
        RandHandCards(--count);
    }
    static int num;
    static public void PutCardInBattle()
    {
        num = 0;
        EntityData entity = null;
        for(int i=0; i < _HandCards.Count; ++i)
        {
            entity = EntityData.GetData(_HandCards [i].UnitId);
            if (entity != null && entity._Cost <= _Fee)
            {
                num++;
                _SelectedHandCardInstID = _HandCards [i].InstId;
                OperateSetActor(FindEmptyPos());
            }
        }
        BattleSetup();
    }
    static OpraSystem op;
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
                    if (point.name.Equals("center"))
                        _Center = point.transform.position;
                    else
                    {
                        int toIdx = int.Parse(point.name) - 1;
                        _PosInScene [toIdx] = point;
                        _PosInScene [toIdx].GetComponent<PointHandle>().Init(toIdx);
                        _PosInScene [toIdx].gameObject.SetActive(false);
                    }
                }
                _IsStagePointInitSuc = true;
            }
        }

        if(_BattleCamera == null)
            _BattleCamera = Camera.main.gameObject.AddComponent<BattleCamera>();

        if(op == null)
            op = GameObject.Find("OpraSystem").GetComponent<OpraSystem>();

        // 加载角色资源
        //TODO

        return _IsStagePointInitSuc && _BattleCamera != null && op != null;
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
                _ActorInScene[i].Fini(true);
        }
        _PosInScene = new Transform[12];
        _ActorInScene = new Actor[12];

        DisposeAssets();
    }

    //播放一回合战报 处理快照
    static void Play()
    {
        if (_ReportAction == null || _ReportAction.Count == 0)
            return;

        if (_ShowTimeDoing)
            return;

        // 处理每回合新上场角色
        if (_BattleReport.UnitList != null && _BattleReport.UnitList.Length > 0)
        {
            _ShowTimeDoing = true;
            EntityData entity;
            DisplayData display;
            Actor actor;
            int localPos;
            for (int i = 0; i < _BattleReport.UnitList.Length; ++i)
            {
                localPos = (_Side == 0? _BattleReport.UnitList [i].Position : ConvertedPos(_BattleReport.UnitList [i].Position));
                actor = GetActorByPos(localPos);
                if (actor != null)
                {
                    actor.SetValue(_BattleReport.UnitList[i].CHP, _BattleReport.UnitList[i].HP);
                    actor.InstID = _BattleReport.UnitList [i].InstId;
                    continue;
                }
                
                entity = EntityData.GetData(_BattleReport.UnitList[i].UnitId);
                display = DisplayData.GetData(entity._DisplayId);
                actor = AddActor(AssetLoader.LoadAsset(display._AssetPath), localPos, _BattleReport.UnitList[i].InstId, _BattleReport.UnitList[i].CHP, _BattleReport.UnitList[i].HP, entity._DisplayId, 0);
                float clipLen = actor.ClipLength(Define.ANIMATION_PLAYER_ACTION_SHOW);
                if (_LongestShowTime < clipLen)
                    _LongestShowTime = clipLen;
            }
            _BattleReport.UnitList = null;
            new Timer().Start(_LongestShowTime, delegate {
                _ShowTimeDoing = false;
                _LongestShowTime = 0f;
            });
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
        if (_IsEnding)
            return;
        
        if (_BattleReport != null)
        {
            CurrentState = BattleState.BS_Play;
            return;
        }

        UIManager.Show("jiesuanjiemian");
        _IsEnding = true;
    }

    //场上添加一个角色
    static Actor AddActor(GameObject go, int pos, long instid, int crtHp, int maxHp, int displayId, int strLv)
    {
        Actor actor = GetActor(instid);
        if (actor != null)
        {
            if(actor._RealPosInScene != pos)
                actor.MoveTo(_PosInScene [pos].position, null);
            GameObject.Destroy(go);
            return actor;
        }
        _ActorInScene[pos] = new Actor(go, _PosInScene[pos], instid, pos, crtHp, maxHp, displayId, strLv);
        _ActorInScene[pos].Play(Define.ANIMATION_PLAYER_ACTION_SHOW);
        _ActorInScene [pos].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);

        UIManager.SetDirty("BattlePanel");

        return _ActorInScene[pos];
    }

    //场上删除一个角色
    static public void DelActor(int pos)
    {
        if (_ActorInScene [pos] != null)
            _ActorInScene [pos].Fini();
        _ActorInScene [pos] = null;
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

    //场上找到一个角色
    static public Actor GetActorByPos(int pos)
    {
        for (int i = 0; i < _ActorInScene.Length; ++i)
        {
            if (_ActorInScene [i] == null)
                continue;

            if (_ActorInScene [i]._RealPosInScene == pos)
                return _ActorInScene [i];
        }
        return null;
    }

    static public bool IsEmptyPos(int pos)
    {
        return _ActorInScene [pos] == null;
    }

    static public int FindEmptyPos()
    {
        int emptyPos = -1;
        for(int i = 0; i < 6; ++i)
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
                return emptyPos;//(_Side == 0 ? emptyPos : ConvertedPos(emptyPos));
            }
        return emptyPos;//(_Side == 0? emptyPos: ConvertedPos(emptyPos));
    }

    static public void BattleSetup()
    {
        NetWoking.S.SetupBattle(Battle._OperatList.ToArray());
        Battle._OperatList.Clear();
        CurrentState = BattleState.BS_Play;
    }

    static public void Judgement()
    {
        if(_Result == null || (BattleResult)_Result.Win == BattleResult.BR_None)
            CurrentState = BattleState.BS_Oper;
        else
            CurrentState = BattleState.BS_Result;

        _BattleReport = null;

        // 从第二回合开始 每回合结束加 1 费
        _Turn++;
        if (_Turn > 1)
            AddFee(_Turn);

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
                if(IsEmptyPos(i) && on)
                    _PosInScene [i].gameObject.SetActive(on);

                if(!on)
                    _PosInScene [i].gameObject.SetActive(on);
            }
        }
    }

    static public void OperateSetActor(int pos)
    {
        if (pos == -1)
            return;

        Battle.ClearSimActor();

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
            bp.Position = (_Side == 0 ? pos : ConvertedPos(pos));
            _OperatList.Add(bp);

            COM_Unit entity = GamePlayer.GetCardByInstID(_SelectedHandCardInstID);
            EntityData eData = EntityData.GetData(entity.UnitId);
            DisplayData displayData = DisplayData.GetData(eData._DisplayId);
            AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, _SelectedHandCardInstID, 100, 100, eData._DisplayId, entity.IProperties[9]);
            RemoveHandCard(_SelectedHandCardInstID);
            CostFee(eData._Cost);
        }

        SwitchPoint(false);
    }

    static public void SimSetActor(int pos)
    {
        if (pos < 0 || pos > 5)
            return;

        COM_Unit entity = GamePlayer.GetCardByInstID(_SelectedHandCardInstID);
        EntityData eData = EntityData.GetData(entity.UnitId);
        DisplayData displayData = DisplayData.GetData(eData._DisplayId);
        AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, 0, 100, 100, eData._DisplayId, entity.IProperties[9]);
    }

    static public void ClearSimActor()
    {
        for(int i=0; i < 6; ++i)
        {
            if (_ActorInScene [i] == null)
                continue;

            if (_ActorInScene [i].InstID == 0)
                DelActor(_ActorInScene[i]._RealPosInScene);
        }
    }

    static public EntityData GetHandCard(int idx)
    {
        if (idx < 0 || idx >= _HandCards.Count)
            return null;

        EntityData eData = EntityData.GetData(_HandCards[idx].UnitId);
        
        return eData;
    }

    static public int GetHandCardStrLv(int idx)
    {
        if (idx < 0 || idx >= _HandCards.Count)
            return null;

        return _HandCards[idx].IProperties[9];
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
            if (_Result == null)
                return false;
            return _Result.Win == (int)BattleResult.BR_Win;
        }
    }

    static public void AddFee(int count)
    {
        _Fee = count;
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

    static public void LaunchBundle()
    {
        EntityData eData = null;
        //AI
        if(Battle._BattleId != 0)
        {
            BattleData bData = BattleData.GetData(Battle._BattleId);
            for(int i=0; i < bData._Monsters.Length; ++i)
            {
                AssetLoader.LaunchBundle(bData._Monsters[i]);
            }
        }

        //Self
        for(int i=0; i < Battle._MyGroupCards.Count; ++i)
        {
            eData = GamePlayer.GetEntityDataByInstID(Battle._MyGroupCards[i]);
            if(eData != null)
                AssetLoader.LaunchBundle(eData._UnitId);
        }

        //opponent
        if (_OpponentCards != null)
        {
            for(int i=0; i < _OpponentCards.Length; ++i)
            {
                AssetLoader.LaunchBundle(_OpponentCards[i]);
            }
        }
    }

    static void DisposeAssets()
    {
        EntityData eData = null;
        //AI
        if(Battle._BattleId != 0)
        {
            BattleData bData = BattleData.GetData(Battle._BattleId);
            for(int i=0; i < bData._Monsters.Length; ++i)
            {
                AssetLoader.DisposeBundle(bData._Monsters[i]);
            }
        }

        //Self
        for(int i=0; i < Battle._MyGroupCards.Count; ++i)
        {
            eData = GamePlayer.GetEntityDataByInstID(Battle._MyGroupCards[i]);
            if(eData != null)
                AssetLoader.DisposeBundle(eData._UnitId);
        }

        //opponent
        if (_OpponentCards != null)
        {
            for(int i=0; i < _OpponentCards.Length; ++i)
            {
                AssetLoader.DisposeBundle(_OpponentCards[i]);
            }
        }
    }

    static public int DropItemCount
    {
        get
        {
            return 1;
            if (_Result == null)
                return 0;
            return 0;//_Result.dropitem.len
        }
    }

    static public COM_ItemInst DropItem(int idx)
    {
        COM_ItemInst isnt = new COM_ItemInst();
        isnt.ItemId = 2;
        isnt.Stack_ = 5;
        return isnt;
        if (_Result == null)
            return null;

        if(idx < 0 || idx >= 0)//_Result.dropitem.len)
            return null;
        
        return null;//_Result.dropitem[idx];
    }

    static public bool InBattle
    {
        get{ return CurrentState != BattleState.BS_Max; } 
    }

    //销毁场景 角色 UI
    static public void Fini()
    {
        UnLoadAssets();
        CurrentState = BattleState.BS_Max;
        _Result = null;
        _ReportIsPlaying = false;
        _ShowTimeDoing = false;
        _IsEnding = false;
        _LongestShowTime = 0f;
        _Fee = 0;
        _BattleId = 0;
        _BattleCamera.Reset();
    }
}
