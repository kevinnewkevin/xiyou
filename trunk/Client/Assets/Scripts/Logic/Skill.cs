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

    // 释放特效缓存
    GameObject _CastEff;

    // 技能特效缓存
    GameObject[] _SkillEff;

    // 受击特效缓存
    GameObject[] _BeattackEff;

    public Skill(int skillId, Actor caster, Actor[] targets, COM_BattleActionTarget[] actionTargets)
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

        _SkillEff = new GameObject[targets.Length];
        if (_SkillData._Motion == SkillData.MotionType.MT_Target)
        {
            for (int i = 0; i < targets.Length; ++i)
            {
                _SkillEff [i] = AssetLoader.LoadAsset(_SkillData._SkillEffect);
                if (_SkillEff [i] != null)
                {
                    if (_SkillData._Motion == SkillData.MotionType.MT_Target)
                    {
                        _SkillEff [i].transform.parent = targets [i]._ActorObj.transform;
                        _SkillEff [i].transform.localPosition = Vector3.zero;
                        _SkillEff [i].transform.localScale = Vector3.one;
                    }

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
                                _Targets[j].UpdateValue(_Actions[j].ActionParam / _SkillData._EmitNumTime.Length, -1);
                                _Targets[j].PopContent(_Actions[j].ActionParam / _SkillData._EmitNumTime.Length);
                            }
                        }));
                    }
                    new Timer().Start(new TimerParam(attackTime, delegate
                    {
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
                            _Targets[j].UpdateValue(_Actions[j].ActionParam / _SkillData._EmitNumTime.Length, -1);
                            _Targets[j].PopContent(_Actions[j].ActionParam / _SkillData._EmitNumTime.Length);
                        }
                    }));
                }
                // 技能表受击时间播放受击目标受击动作和受击特效
                new Timer().Start(new TimerParam(_SkillData._TotalTime, delegate
                {
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

    void Clear()
    {
        _IsCasting = false;
        _IsCasted = true;
    }
}
