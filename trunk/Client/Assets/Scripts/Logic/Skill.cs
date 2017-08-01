using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Skill {

    SkillData _SkillData;
    Actor _Caster;
    Actor[] _Targets;
    Vector3 _OriginPos;

    GameObject _CastEff;
    GameObject[] _SkillEff;
    GameObject[] _BeattackEff;

    // skilldata member value;

	public Skill(int skillId, Actor caster, Actor[] targets)
    {
        // get skilldata by id

        _SkillData = SkillData.GetData(skillId);

        _CastEff = AssetLoader.LoadAsset(_SkillData._CastEffect);
        if (_CastEff != null)
        {
            _CastEff.transform.parent = caster._ActorObj.transform;
            _CastEff.transform.localPosition = Vector3.zero;
            _CastEff.transform.localScale = Vector3.one;
            _CastEff.SetActive(false);
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
    }

    public bool Cast()
    {
        Debug.Log("Skill Cast");
        if (_Caster == null)
            return false;

        if (_Targets == null || _Targets.Length == 0)
            return false;

        // judge whether is melee skill
        if (_SkillData._IsMelee)
        {
            //cast anim
            _Caster.Play(_SkillData._CastAnim);
            //cast effect
            if (_CastEff != null)
                _CastEff.SetActive(true);

            float attackTime = _Caster.ClipLength(_SkillData._AttackAnim);
            new Timer().Start(new TimerParam(_SkillData._CastTime, delegate
            {
                for (int i = 0; i < _SkillEff.Length; ++i)
                {
                    //skill effect
                    if (_SkillEff[i] != null)
                        _SkillEff[i].SetActive(true);
                }

                _Caster.Play(Define.ANIMATION_PLAYER_ACTION_RUN);
                _Caster.MoveTo(_Targets[0].Forward, delegate
                {
                    _Caster.Stop();
                    _Caster.Play(_SkillData._AttackAnim);
                    //1.目标播受击动作和特效的时间
                    //2.目标弹伤害数字的时间
                    //3.施法者回归时间
                    new Timer().Start(new TimerParam(_SkillData._BeattackTime, delegate
                    {
                        for (int i = 0; i < _Targets.Length; ++i)
                        {
                            _Targets[i].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                            _Targets[i].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);

                            //beattack effect
                            if (_BeattackEff[i] != null)
                                _BeattackEff[i].SetActive(true);
                        }
                    }), new TimerParam(_SkillData._EmitNumTime, delegate
                    {
                        for (int i = 0; i < _Targets.Length; ++i)
                        {
                            _Targets[i].PopContent();
                        }
                    }), new TimerParam(attackTime, delegate
                    {
                        _Caster.MoveTo(_OriginPos, delegate {
                            Battle._ReportIsPlaying = false;
                            _Caster.Stop();
                        });
                    }));
                });
            }));
        }
        else
        {
            //clip name in skilldata
            _Caster.Play(_SkillData._CastAnim);
            _Caster.PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
            //cast effect
            if (_CastEff != null)
                _CastEff.SetActive(true);
            //1.目标播受击动作和特效的时间
            //2.目标弹伤害数字的时间
            //3.技能总时间
            new Timer().Start(new TimerParam(_SkillData._CastTime, delegate
            {
                for (int i = 0; i < _SkillEff.Length; ++i)
                {
                    //skill effect
                    if (_SkillEff[i] != null)
                        _SkillEff[i].SetActive(true);
                }

                if(_SkillData._Motion == SkillData.MotionType.MT_Fly)
                {
                    for (int i = 0; i < _SkillEff.Length; ++i)
                    {
                        iTween.MoveTo(_SkillEff[i], iTween.Hash("time", _SkillData._BeattackTime, "position", _Targets[i]._ActorObj.transform.position, "easetype", iTween.EaseType.linear));
                    }
                }

                new Timer().Start(new TimerParam(_SkillData._BeattackTime, delegate
                {
                    for (int i = 0; i < _SkillEff.Length; ++i)
                    {
                        GameObject.Destroy(_SkillEff[i]);
                    }

                    for (int i = 0; i < _Targets.Length; ++i)
                    {
                        _Targets[i].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                        _Targets[i].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);

                        //beattack effect
                        if (_BeattackEff[i] != null)
                            _BeattackEff[i].SetActive(true);
                    }
                }), new TimerParam(_SkillData._EmitNumTime, delegate
                {
                    for (int i = 0; i < _Targets.Length; ++i)
                    {
                        _Targets[i].PopContent();
                    }
                }), new TimerParam(_SkillData._TotalTime, delegate
                {
                    Battle._ReportIsPlaying = false;
                    _Caster.Stop();
                }));
            }));
        }

        return true;
    }
}
