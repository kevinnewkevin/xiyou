using UnityEngine;
using System.Collections.Generic;
using FairyGUI;

public class Battle {

    public enum BattleState
    {
        BS_Init,
        BS_Opra,
        BS_Eff,
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

    static public List<ReportBase> _ReportTips = new List<ReportBase>();

    static public int _Turn;

    static public int _Fee;

    static public int _MaxFee;

    static public int _MaxTimeLeft;

    static public int _Side;

    static public int _BattleId;

    static public bool _IsRecord;

    static public Vector3 _Center;

    static public Transform _CenterTrans;

    static public List<long> _MyGroupCards;

    static public BattleCamera _BattleCamera;

    static public BattleSceneTouch _BattleTouch;

    static public int[] _OpponentCards;

    static public int _CasterDisplayID;

    static public string _CasterSkillName;

    static public int _SelectReportIdx; //当前选择的战报索引

    static int _SelectSkillID;      //主角选择的技能id

    static COM_BattleUnit[] _OriginUnits;

    static int _BattleStateAfterInit;

    static public COM_Unit _ThrowCardInst;

    static public int SelectSkillID
    {
        set
        {
            _SelectSkillID = value;
            UIManager.SetDirty("BattlePanel");
            PushReportTip(_SelectSkillID);
        }
        get
        {
            return _SelectSkillID;
        }
    }

    static public COM_BattleReport BattleReport
    {
        set
        {
            _BattleReport = value;
            _ReportAction = new List<COM_BattleAction>(_BattleReport.ActionList);
            _DealUnitList = false;
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
                    // battle has anim but is not record
                    if (!_IsRecord && bData != null && bData._Animations != null && bData._Animations.Length > 0 && _BattleStateAfterInit == 0)
                    {
                        GRoot.inst.modalLayer.visible = false;
                        op.Begin(bData._Animations);
                        op.Play();
                        CurrentState = BattleState.BS_Opra;
                    }
                    else
                    {
                        LoadOrigin();
                        UIManager.Show("BattlePanel");
                    }
                }
                break;
            case BattleState.BS_Opra:
                break;
            case BattleState.BS_Oper:
                if (_IsRecord)
                {
                    BattleRecordSystem.LaunchOperate();
                    CurrentState = BattleState.BS_Play;
                    return;
                }

                if (GamePlayer._IsAuto)
                {
                    if (_Turn == 1)
                        PutMainInBattle();
                    else
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
    static public void Init(int side, int battleid = 0, int[] opponentCards = null, COM_BattleUnit[] units = null, bool isRecord = false, int turn = 1, int timeleft = 0, int state = 0)//0正常 1操作 2播放
    {
        _OpponentCards = opponentCards;
        _SceneConfig = null;
        _IsStagePointInitSuc = false;
        _ReportIsPlaying = false;
        _Turn = turn;
        _Side = side;
        _BattleId = battleid;
        CurrentState = BattleState.BS_Init;
        _OperatList.Clear();
        _HandCards.Clear();
        _OriginUnits = units;
        _IsRecord = isRecord;
        _BattleStateAfterInit = state;

        //_HandCards.Add(GamePlayer._Data);
        _MyGroupCards = GamePlayer.GetBattleCardsCopy();

        _MaxFee = Define.GetInt("MaxFee");
        _MaxTimeLeft = timeleft == 0? Define.GetInt("BattleMaxTime"): timeleft;
        AddFee(_Turn);
        RandHandCards(3);
    }

    static public void ResetTimeLeft()
    {
        _MaxTimeLeft = Define.GetInt("BattleMaxTime");
    }

    static public void FadedCallback()
    {
        if(CurrentState == BattleState.BS_Eff)
            return;

        if (_BattleStateAfterInit == 0)
        {
            UIManager.GetUI("BattlePanel").Call("ShowBattleStart");
            new Timer().Start(1f, delegate
            {
                ShowTurn();
            });
            CurrentState = BattleState.BS_Eff;
        }
        else
        {
            CurrentState = _BattleStateAfterInit == 1 ? BattleState.BS_Oper : BattleState.BS_Play;
            _BattleStateAfterInit = 0;
        }
    }

    static void ShowTurn()
    {
        UIManager.GetUI("BattlePanel").Call("ShowTurn");
        new Timer().Start(1.5f, delegate {
            if(InBattle)
            {
                CurrentState = BattleState.BS_Oper;
                NetWoking.S.ChangeBattleState(1);
            }
        });

        if (_ActorInScene != null)
        {
            for(int i=0; i < _ActorInScene.Length; ++i)
            {
                if (_ActorInScene [i] != null && _ActorInScene [i]._Headbar != null)
                    _ActorInScene [i]._Headbar._IsDirty = true;
            }
        }
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

    static public void PutMainInBattle()
    {
        EntityData entity = null;
        for(int i=0; i < _HandCards.Count; ++i)
        {
            entity = EntityData.GetData(_HandCards [i].UnitId);
            if (entity != null && entity._Cost <= _Fee)
            {
                num++;
                _SelectedHandCardInstID = _HandCards [i].InstId;
                OperateSetActor(4);
            }
        }
        BattleSetup();
    }

    static public OpraSystem op;
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
                    {
                        _Center = point.transform.position;
                        _CenterTrans = point.transform;
                    }
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

        if(_BattleTouch == null)
            _BattleTouch = GameObject.FindObjectOfType<BattleSceneTouch>();

        if(op == null)
            op = GameObject.Find("OpraSystem").GetComponent<OpraSystem>();

        if (!_IsStagePointInitSuc)
            return false;

        return _IsStagePointInitSuc && _BattleCamera != null && op != null && _BattleTouch != null;
    }

    static public void LoadOrigin()
    {
        // 加载角色资源
        if (_OriginUnits != null && _OriginUnits.Length > 0)
        {
            EntityData entity;
            DisplayData display;
            Actor actor;
            int localPos;
            for (int i = 0; i < _OriginUnits.Length; ++i)
            {
                localPos = (_Side == 0? _OriginUnits [i].Position : ConvertedPos(_OriginUnits [i].Position));
                actor = GetActorByPos(localPos);
                if (actor != null)
                {
                    actor.SetValue(_OriginUnits[i].CHP, _OriginUnits[i].HP);
                    actor.InstID = _OriginUnits [i].InstId;
                    continue;
                }

                entity = EntityData.GetData(_OriginUnits[i].UnitId);
                display = DisplayData.GetData(entity._DisplayId);
                actor = AddActor(AssetLoader.LoadAsset(display._AssetPath), localPos, _OriginUnits[i].InstId, _OriginUnits[i].CHP, _OriginUnits[i].HP, entity._UnitId, _OriginUnits[i].Level);
            }
            _OriginUnits = null;
        }
    }

    static int ConvertedPos(int pos)
    {
        return (pos + 6) % 12;
    }

    static bool PlaceActor()
    {
        bool ok = false;
        if (StageCamera.main != null)
        {
            StageCamera.main.transform.position = new Vector3(StageCamera.main.transform.position.x, Stage.inst.y - 5f, StageCamera.main.transform.position.z);
            ok = true;
        }
        return ok;
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

    static bool _DealUnitList;
    //播放一回合战报 处理快照
    static void Play()
    {
        if (_ReportAction == null || _ReportAction.Count == 0)
            return;

        if (_ShowTimeDoing)
            return;

        // 处理每回合新上场角色
        if (!_DealUnitList && _BattleReport.UnitList != null && _BattleReport.UnitList.Length > 0)
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
//                    _BattleReport.UnitList [i] = null;
                    continue;
                }
                
                entity = EntityData.GetData(_BattleReport.UnitList[i].UnitId);
                display = DisplayData.GetData(entity._DisplayId);
                actor = AddActor(AssetLoader.LoadAsset(display._AssetPath), localPos, _BattleReport.UnitList[i].InstId, _BattleReport.UnitList[i].CHP, _BattleReport.UnitList[i].HP, entity._UnitId, _BattleReport.UnitList[i].Level);
                float clipLen = actor.ClipLength(Define.ANIMATION_PLAYER_ACTION_SHOW);
                if (_LongestShowTime < clipLen)
                    _LongestShowTime = clipLen;
            }
            PushReportTip(_BattleReport.UnitList);
//            _BattleReport.UnitList = null;
            _DealUnitList = true;
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
            if (_ReportAction [0].TargetList != null)
            {
                for (int i = 0; i < _ReportAction [0].TargetList.Length; ++i)
                {
                    target = GetActor(_ReportAction [0].TargetList [i].InstId);
                    if (target == null)
                    {
                        Debug.LogError("施法者： " + _CrtActor._Name + " 释放技能： " + _ReportAction [0].SkillId + " 目标ID： " + _ReportAction [0].TargetList [i].InstId + " 为空!");
                    }
                    _CrtTargets.Add(target);
                }
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
            _CrtSkill = new Skill(_ReportAction [0].SkillId, _CrtActor, _CrtTargets.ToArray(), _ReportAction [0].TargetList, _ReportAction[0].SkillBuff, _ReportAction[0].UnitList);
            _CrtSkill.Cast();
            PushReportTip(_ReportAction[0]);
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

        //check dead anim
        for(int i=0; i < _ActorInScene.Length; ++i)
        {
            if (_ActorInScene [i] == null)
                continue;
            
            if (_ActorInScene [i].IsPlay(Define.ANIMATION_PLAYER_ACTION_DEAD))
                return;
        }
        
        if (_BattleReport != null)
        {
            CurrentState = BattleState.BS_Play;
            return;
        }

        UIManager.Show("jiesuanjiemian");
        _IsEnding = true;
    }

    //场上添加一个角色
    static public Actor AddActor(GameObject go, int pos, long instid, int crtHp, int maxHp, int entityid, int strLv)
    {
        Actor actor = GetActor(instid);
        if (actor != null)
        {
            if(actor._RealPosInScene != pos)
                actor.MoveTo(_PosInScene [pos].position, null);
            GameObject.Destroy(go);
            return actor;
        }
        _ActorInScene[pos] = new Actor(go, _PosInScene[pos], instid, pos, crtHp, maxHp, entityid, strLv);
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

    static public void SwapActor(int frompos, int topos)
    {
        if (frompos < 0 || frompos >= _ActorInScene.Length)
            return;

        if (topos < 0 || topos >= _ActorInScene.Length)
            return;

        _ActorInScene[topos] = _ActorInScene[frompos];
        _ActorInScene [frompos] = null;
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

                if (_ActorInScene [j]._RealPosInScene == i && _ActorInScene [j].InstID != 0)
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
        NetWoking.S.SetupBattle(Battle._OperatList.ToArray(), SelectSkillID);
        _SelectSkillID = 0;
        Battle._OperatList.Clear();
        CurrentState = BattleState.BS_Play;
    }

    static public void Judgement()
    {
        if (_Result == null || (BattleResult)_Result.Win == BattleResult.BR_None)
        {
            //每回合结束加 1 费
            _Turn++;
            AddFee(_Turn);
            if(_Turn > 2)
                RandHandCards(1);
            CurrentState = BattleState.BS_Eff;
            ShowTurn();
        }
        else
            CurrentState = BattleState.BS_Result;

        _BattleReport = null;
   }

    static public void SwitchPoint(bool on)
    {
        Actor actor = null;
        int localPos = 0;
        for(int i=0; i < 6; ++i)
        {
            localPos = (_Side == 0? i : ConvertedPos(i));
            actor = GetActorByPos(localPos);
            if (_PosInScene [i] != null || (actor != null && actor.IsPlay(Define.ANIMATION_PLAYER_ACTION_DEAD)))
            {
                if(IsEmptyPos(i) && on)
                    _PosInScene [i].gameObject.SetActive(on);

                if(!on)
                    _PosInScene [i].gameObject.SetActive(on);
            }
        }

        if (on == false)
        {
            UIWindow win = UIManager.GetUI("BattlePanel");
            if(win != null)
                win.Call("NormalCard");
        }

        GuideSystem.SpecialEvt("battle_switchpoint", on);
    }

    static public Transform GetPoint(int idx)
    {
        return _PosInScene [idx].transform;
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
            PushReportTip(bp);

            COM_Unit entity = GamePlayer.GetCardByInstID(_SelectedHandCardInstID);
            EntityData eData = EntityData.GetData(entity.UnitId);
            DisplayData displayData = DisplayData.GetData(eData._DisplayId);
            AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, _SelectedHandCardInstID, 100, 100, eData._UnitId, entity.IProperties[9]);
            RemoveHandCard(_SelectedHandCardInstID);
            _SelectedHandCardInstID = 0;
            CostFee(eData._Cost);
            AudioSystem.PlayVoice(eData._Voice);
        }

        SwitchPoint(false);
        GuideSystem.SpecialEvt("battle_cardonbattle");
    }

    static public void SimSetActor(int pos)
    {
        if (pos < 0 || pos > 5)
            return;

        COM_Unit entity = GamePlayer.GetCardByInstID(_SelectedHandCardInstID);
        EntityData eData = EntityData.GetData(entity.UnitId);
        DisplayData displayData = DisplayData.GetData(eData._DisplayId);
        AddActor(AssetLoader.LoadAsset(displayData._AssetPath), pos, 0, 100, 100, eData._UnitId, entity.IProperties[9]);
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
            return 0;

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

    static public void RemoveHandCard(long instid, bool anim = false)
    {
        for(int i=0; i < _HandCards.Count; ++i)
        {
            if (_HandCards [i].InstId == instid)
            {
                _HandCards.RemoveAt(i);
                if (anim)
                    _ThrowCardInst = GamePlayer.GetCardByInstID(instid);
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
            if (_Result != null && _CurrentState == BattleState.BS_Result)
                return;
            
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

    //Push 自身出牌
    static void PushReportTip(COM_BattlePosition operate)
    {
        ReportBase rb = new ReportBase();
        rb._RBType = ReportBase.RBType.RBT_SelfAppear;
        EntityData eData = GamePlayer.GetEntityDataByInstID(operate.InstId);
        if (eData != null)
            rb._CasterEntityID = eData._UnitId;
        rb._Self = true;
        _ReportTips.Add(rb);
        while (_ReportTips.Count > 15)
        {
            _ReportTips.RemoveAt(0);
        }
        UIManager.SetDirty("BattlePanel");
    }

    //Push 自身技能
    static void PushReportTip(int operateSkillId)
    {
        if (operateSkillId == 0)
            return;
        
        ReportBase rb = new ReportBase();
        rb._RBType = ReportBase.RBType.RBT_SelfSkill;
        rb._SkillID = operateSkillId;
        rb._Self = true;
        _ReportTips.Add(rb);
        while (_ReportTips.Count > 15)
        {
            _ReportTips.RemoveAt(0);
        }
        UIManager.SetDirty("BattlePanel");
    }

    //Push 战报上场
    static void PushReportTip(COM_BattleUnit[] appearActors)
    {
        ReportBase rb = null;
        for(int i=0; i < appearActors.Length; ++i)
        {
            if (appearActors [i] == null)
                continue;
            
            rb = new ReportBase();
            rb._RBType = ReportBase.RBType.RBT_AllAppear;
            rb._CasterEntityID = appearActors [i].UnitId;
            Actor actor = GetActor(appearActors [i].InstId);
            if (actor != null)
            {
                rb._Self = actor._RealPosInScene < 6;
            }
            _ReportTips.Add(rb);
        }
        while (_ReportTips.Count > 15)
        {
            _ReportTips.RemoveAt(0);
        }
        UIManager.SetDirty("BattlePanel");
    }

    //Push 战报技能
    //Push 战报buff(回合开始自然增删忽略， 行动后自身buff处理，目标buff处理)
    static void PushReportTip(COM_BattleAction report)
    {
        if (report.SkillId == 0)
            return;
        
        ReportBase rb = new ReportBase();
        rb._RBType = ReportBase.RBType.RBT_AllSkill;
        Actor actor = GetActor(report.InstId);
        if (actor != null)
        {
            rb._CasterEntityID = actor._EntityID;
            rb._Self = actor._RealPosInScene < 6;
        }
        
        rb._SkillID = report.SkillId;
        if (report.SkillBuff != null)
        {
            rb._Buffs = new COM_BattleBuff[report.SkillBuff.Length];
            for(int i=0; i < rb._Buffs.Length; ++i)
            {
                rb._Buffs [i] = new COM_BattleBuff();
                rb._Buffs [i].BuffId = report.SkillBuff [i].BuffId;
                rb._Buffs [i].Change = report.SkillBuff [i].Change;
            }
        }
        if (report.TargetList != null)
        {
            rb._Targets = new COM_BattleActionTarget[report.TargetList.Length];
            rb._TargetEntityID = new int[report.TargetList.Length];
            rb._TargetSelf = new bool[report.TargetList.Length];
            for(int i=0; i < rb._Targets.Length; ++i)
            {
                rb._Targets [i] = new COM_BattleActionTarget();
                actor = GetActor(report.TargetList [i].InstId);
                rb._TargetEntityID [i] = actor._EntityID;
                rb._TargetSelf[i] = actor._RealPosInScene < 6;
                rb._Targets [i].ActionType = report.TargetList [i].ActionType;
                rb._Targets [i].ActionParam = report.TargetList [i].ActionParam;
                rb._Targets [i].ActionParamExt = report.TargetList [i].ActionParamExt;
                rb._Targets [i].Dead = report.TargetList [i].Dead;
                rb._Targets [i].ThrowCard = new COM_ThrowCard();
                rb._Targets [i].ThrowCard.InstId = report.TargetList [i].ThrowCard.InstId;
                rb._Targets [i].ThrowCard.EntityId = report.TargetList [i].ThrowCard.EntityId;
                rb._Targets [i].ThrowCard.Level = report.TargetList [i].ThrowCard.Level;

                if (report.TargetList [i].BuffAdd != null)
                {
                    rb._Targets [i].BuffAdd = new COM_BattleBuff[report.TargetList[i].BuffAdd.Length];
                    for(int j=0; j < rb._Targets [i].BuffAdd.Length; ++j)
                    {
                        rb._Targets [i].BuffAdd[j] = new COM_BattleBuff();
                        rb._Targets [i].BuffAdd[j].BuffId = report.TargetList[i].BuffAdd [j].BuffId;
                        rb._Targets [i].BuffAdd[j].Change = report.TargetList[i].BuffAdd [j].Change;
                    }
                }
            }
        }
        _ReportTips.Add(rb);
        while (_ReportTips.Count > Define._MaxReportTips)
        {
            _ReportTips.RemoveAt(0);
        }
        UIManager.SetDirty("BattlePanel");
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
            if (_Result == null)
                return 0;

            if (_Result.BattleItems == null)
                return 0;
            
            return _Result.BattleItems.Length;
        }
    }

    static public COM_ItemInst DropItem(int idx)
    {
        if (_Result == null)
            return null;

        if (_Result.BattleItems == null)
            return null;

        if(idx < 0 || idx >= _Result.BattleItems.Length)
            return null;
        
        return _Result.BattleItems[idx];
    }

    static public bool InBattle
    {
        get{ return CurrentState != BattleState.BS_Max; } 
    }

    //销毁场景 角色 UI
    static public void Fini()
    {
        UnLoadAssets();
        _Result = null;
        _IsRecord = false;
        _ReportIsPlaying = false;
        _ShowTimeDoing = false;
        _IsEnding = false;
        _ActorLaunched = false;
        _ShowTimeDoing = false;
        _LongestShowTime = 0f;
        _Fee = 0;
        _Turn = 0;
        _BattleId = 0;
        _BattleCamera.Reset();
        _SelectSkillID = 0;
        _OriginUnits = null;
        _ReportTips.Clear();
        _CrtSkill = null;
        _CrtBuffChecker = null;
        CurrentState = BattleState.BS_Max;
    }
}

public class ReportBase
{
    public enum RBType
    {
        RBT_SelfSkill,
        RBT_SelfAppear,
        RBT_AllAppear,
        RBT_AllSkill,
        RBT_AllBuff,
    }
    public bool _Self;
    public RBType _RBType;
    public int _SkillID;
    public int _CasterEntityID;
    public COM_BattleBuff[] _Buffs;
    public COM_BattleActionTarget[] _Targets;
    public int[] _TargetEntityID;
    public bool[] _TargetSelf;
}