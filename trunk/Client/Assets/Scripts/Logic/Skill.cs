using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Skill {

    public bool _IsCasting;
    public bool _IsCasted;

    // 技能静态数据
    SkillData _SkillData;

    // 技能释放者
    Actor _Caster;

    // 技能目标，可能有多个
    Actor[] _Targets;

    // 释放者的原始位置(有位移时回归位置用)
    Vector3 _OriginPos;

    // 该单元战报(取需要的数据用)
    COM_BattleActionTarget[] _Actions;

    // 技能释放后自身增删的buff
    COM_BattleBuff[] _SkillBuff;

    // 释放特效缓存
    GameObject _CastEff;

    // 技能特效缓存
    GameObject[] _SkillEff;

    // 受击特效缓存
    GameObject[] _BeattackEff;

    public Skill(int skillId, Actor caster, Actor[] targets, COM_BattleActionTarget[] actionTargets, COM_BattleBuff[] skillBuffs)
    {
        _IsCasting = true;
        // get skilldata by id

        _SkillData = SkillData.GetData(skillId);
        if (_SkillData == null)
        {
            Clear();
            return;
        }

        // 根据技能类型初始化特效
        if (!string.IsNullOrEmpty(_SkillData._CastEffect))
        {
            _CastEff = AssetLoader.LoadAsset(_SkillData._CastEffect);
            if (_CastEff != null)
            {
                _CastEff.transform.parent = caster._ActorObj.transform;
                _CastEff.transform.localPosition = Vector3.zero;
                _CastEff.transform.localScale = Vector3.one;
                _CastEff.SetActive(false);
            }
        }

        if (_SkillData._Single)
        {
            _SkillEff = new GameObject[1];
            _SkillEff [0] = AssetLoader.LoadAsset(_SkillData._SkillEffect);
            if (_SkillEff [0] != null)
            {
                if (_SkillData._Motion == SkillData.MotionType.MT_Target)
                {
                    _SkillEff [0].transform.parent = targets[0]._ActorObj.transform;
                }
                else
                {
                    _SkillEff [0].transform.parent = caster._ActorObj.transform;
                }
                _SkillEff [0].transform.localPosition = Vector3.zero;
                _SkillEff [0].transform.localScale = Vector3.one;
                _SkillEff [0].SetActive(false);
            }
        }
        else
        {
            _SkillEff = new GameObject[targets.Length];
            if (_SkillData._Motion == SkillData.MotionType.MT_Target)
            {
                for (int i = 0; i < targets.Length; ++i)
                {
                    _SkillEff [i] = AssetLoader.LoadAsset(_SkillData._SkillEffect);
                    if (_SkillEff [i] != null)
                    {
                        _SkillEff [i].transform.parent = targets [i]._ActorObj.transform;
                        _SkillEff [i].transform.localPosition = Vector3.zero;
                        _SkillEff [i].transform.localScale = Vector3.one;
                        _SkillEff [i].SetActive(false);
                    }
                }
            }
            else if(_SkillData._Motion == SkillData.MotionType.MT_Self)
            {
                _SkillEff = new GameObject[1];
                _SkillEff[0] = AssetLoader.LoadAsset(_SkillData._SkillEffect);
                _SkillEff[0].transform.parent = caster._ActorObj.transform;
                _SkillEff[0].transform.localPosition = Vector3.zero;
                _SkillEff[0].transform.localScale = Vector3.one;
                _SkillEff[0].transform.rotation = caster._ActorObj.transform.rotation;
                _SkillEff[0].SetActive(false);
            }
            else if(_SkillData._Motion == SkillData.MotionType.MT_Fly)
            {
                for (int i = 0; i < targets.Length; ++i)
                {
                    _SkillEff [i] = AssetLoader.LoadAsset(_SkillData._SkillEffect);
                    if (_SkillEff [i] != null)
                    {
                        _SkillEff [i].transform.position = caster._ActorObj.transform.position;
                        _SkillEff [i].SetActive(false);
                    }
                }
            }
        }

        _BeattackEff = new GameObject[targets.Length];
        for (int i = 0; i < targets.Length; ++i)
        {
            _BeattackEff[i] = AssetLoader.LoadAsset(_SkillData._BeattackEffect);
            if (_BeattackEff [i] != null)
            {
                _BeattackEff[i].transform.parent = targets[i]._ActorObj.transform;
                _BeattackEff[i].transform.localPosition = Vector3.zero;
                _BeattackEff[i].transform.localScale = Vector3.one;
                _BeattackEff[i].SetActive(false);
            }
        }
        _OriginPos = caster._ActorObj.transform.position;
        _Caster = caster;
        _Targets = targets;
        _Actions = actionTargets;
        _SkillBuff = skillBuffs;
    }

    public bool Cast()
    {
        if (_Caster == null)
        {
            Clear();
            return false;
        }

        if (_Targets == null || _Targets.Length == 0)
        {
            Clear();
            return false;
        }

        // 判断技能是否是近战
        if (_SkillData._IsMelee)
        {
            // 播放释放者的释放动作
            _Caster.Play(_SkillData._CastAnim);
            // 播放释放特效
            if (_CastEff != null)
                _CastEff.SetActive(true);

            // 获取攻击动作的时长(用于播放攻击动作后的回归调用)
            float attackTime = _Caster.ClipLength(_SkillData._AttackAnim);

            // 同时根据技能表的释放时间计时
            new Timer().Start(new TimerParam(_SkillData._CastTime, delegate
            {
                // 释放者播放奔跑动作
                _Caster.Play(Define.ANIMATION_PLAYER_ACTION_RUN);

                // 释放者移动到目标前
                _Caster.MoveTo(_Targets[0].Forward, delegate
                {
                    _Caster.Stop();

                    // 释放者播放攻击动作
                    _Caster.Play(_SkillData._AttackAnim);
                    // 播放技能特效
                    for (int i = 0; i < _SkillEff.Length; ++i)
                    {
                        //skill effect
                        if (_SkillEff[i] != null)
                            _SkillEff[i].SetActive(true);
                    }
                    //1.目标播受击动作和特效的时间
                    //2.目标弹伤害数字的时间
                    //3.施法者回归时间
                    for(int i=0; i < _SkillData._BeattackTime.Length; ++i)
                    {
                        new Timer().Start(new TimerParam(_SkillData._BeattackTime[i], delegate
                        {
                            for (int j = 0; j < _Targets.Length; ++j)
                            {
                                _Targets[j].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                                _Targets[j].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);

                                //beattack effect
                                if (_BeattackEff[j] != null)
                                {
                                    _BeattackEff[j].SetActive(false);
                                    _BeattackEff[j].SetActive(true);
                                }
                            }
                        }));
                    }
                    for(int i=0; i < _SkillData._EmitNumTime.Length; ++i)
                    {
                        new Timer().Start(new TimerParam(_SkillData._EmitNumTime[i], delegate
                        {
                            for (int j = 0; j < _Targets.Length; ++j)
                            {
                                _Targets[j].UpdateValue(_Actions[j].ActionParam, -1);
                                int disValue = _Actions[j].ActionParam / _SkillData._EmitNumTime.Length;
                                if(disValue == 0)
                                {
                                    disValue = _Actions[j].ActionParam > 0? 1: -1;
                                }
                                _Targets[j].PopContent(disValue);
                            }
                        }));
                    }
                    new Timer().Start(new TimerParam(attackTime, delegate
                    {
                        HandleBuff();
                        _Caster.MoveTo(_OriginPos, delegate {
                            Clear();
                            _Caster.Stop();
                            _Caster.Reset();
                        });
                    }));
                });
            }));
        }
        else
        {
            // 播放释放动作
            _Caster.Play(_SkillData._CastAnim);
            _Caster.PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
            // 播放释放特效
            if (_CastEff != null)
                _CastEff.SetActive(true);
            //1.目标播受击动作和特效的时间
            //2.目标弹伤害数字的时间
            //3.技能总时间
            new Timer().Start(new TimerParam(_SkillData._CastTime, delegate
            {
                for (int i = 0; i < _SkillEff.Length; ++i)
                {
                    //播放技能特效
                    if (_SkillEff[i] != null)
                        _SkillEff[i].SetActive(true);
                }

                // 如果技能移动方式为fly 即有位移 则控制其移动
                if(_SkillData._Motion == SkillData.MotionType.MT_Fly)
                {
                    for (int i = 0; i < _SkillEff.Length; ++i)
                    {
                        _SkillEff[i].transform.LookAt(_Targets[i]._ActorObj.transform, Vector3.up);
                        iTween.MoveTo(_SkillEff[i], iTween.Hash("time", _SkillData._BeattackTime[0], "position", _Targets[i]._ActorObj.transform.position, "easetype", iTween.EaseType.linear));
                    }
                }
                for(int i=0; i < _SkillData._BeattackTime.Length; ++i)
                {
                    new Timer().Start(new TimerParam(_SkillData._BeattackTime[i], delegate
                    {
                        for (int j = 0; j < _Targets.Length; ++j)
                        {
                            _Targets[j].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                            _Targets[j].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);

                            //beattack effect
                            if (_BeattackEff[j] != null)
                            {
                                _BeattackEff[j].SetActive(false);
                                _BeattackEff[j].SetActive(true);
                            }
                        }
                    }));
                }
                for(int i=0; i < _SkillData._EmitNumTime.Length; ++i)
                {
                    new Timer().Start(new TimerParam(_SkillData._EmitNumTime[i], delegate
                    {
                        for (int j = 0; j < _Targets.Length; ++j)
                        {
                            _Targets[j].UpdateValue(_Actions[j].ActionParam, -1);
                            int disValue = _Actions[j].ActionParam / _SkillData._EmitNumTime.Length;
                            if(disValue == 0)
                            {
                                disValue = _Actions[j].ActionParam > 0? 1: -1;
                            }
                            _Targets[j].PopContent(disValue);
                        }
                    }));
                }
                // 技能表受击时间播放受击目标受击动作和受击特效
                new Timer().Start(new TimerParam(_SkillData._TotalTime, delegate
                {
                    HandleBuff();
                    // 远程类技能根据TotalTime 总时间回归复原
                    for (int i = 0; i < _SkillEff.Length; ++i)
                    {
                        GameObject.Destroy(_SkillEff[i]);
                    }
                    Clear();
                    _Caster.Stop();
                    _Caster.Reset();
                }));
            }));
        }

        return true;
    }

    void HandleBuff()
    {
        Actor target = null;
        for(int i=0; i < _Actions.Length; ++i)
        {
            target = Battle.GetActor(_Actions[i].InstId);
            if (target != null)
            {
                if (_Actions [i].BuffAdd == null)
                {
                    Debug.LogWarning(" _Actions [i].BuffAdd is null!");
                }
                else
                {
                    for(int j=0; j < _Actions[i].BuffAdd.Length; ++j)
                    {
                        if (_Actions [i].BuffAdd [j].Change == 1)
                            target.AddBuff(_Actions [i].BuffAdd [j].BuffId);
                        else if(_Actions [i].BuffAdd [j].Change == 0)
                            target.RemoveBuff(_Actions [i].BuffAdd [j].BuffId);
                    }
                }

                if (_Actions [i].Dead)
                {
                    target.Play(Define.ANIMATION_PLAYER_ACTION_DEAD);
                    new Timer().Start(target.ClipLength(Define.ANIMATION_PLAYER_ACTION_DEAD) + 1f, delegate {
                        Battle.DelActor(target._RealPosInScene);
                    });
                }
            }
        }

        if (_SkillBuff != null)
        {
            for(int i=0; i < _SkillBuff.Length; ++i)
            {
                if (_SkillBuff [i].Change == 1)
                    _Caster.AddBuff(_SkillBuff [i].BuffId);
                else if(_SkillBuff [i].Change == 0)
                    _Caster.RemoveBuff(_SkillBuff [i].BuffId);
            }
        }
    }

    void Clear()
    {
        _IsCasting = false;
        _IsCasted = true;

        if (_CastEff != null)
            GameObject.Destroy(_CastEff);

        if (_SkillEff != null)
        {
            for(int i=0; i < _SkillEff.Length; ++i)
            {
                GameObject.Destroy(_SkillEff[i]);
            }
        }

        if (_BeattackEff != null)
        {
            for(int i=0; i < _BeattackEff.Length; ++i)
            {
                GameObject.Destroy(_BeattackEff[i]);
            }
        }
    }
}
