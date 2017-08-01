using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Skill {

    SkillData _SkillData;
    Actor _Caster;
    Actor[] _Targets;
    Vector3 _OriginPos;

    // skilldata member value;

	public Skill(int skillId, Actor caster, Actor[] targets)
    {
        // get skilldata by id

        _SkillData = SkillData.GetData(skillId);
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
            //cast effect
            GameObject effCast = AssetLoader.LoadAsset(_SkillData._CastEffect);
            effCast.transform.parent = _Caster._ActorObj.transform;
            effCast.transform.localPosition = Vector3.zero;
            effCast.transform.localScale = Vector3.one;
            effCast.SetActive(false);
            effCast.SetActive(true);

            float attackTime = _Caster.ClipLength(_SkillData._CastAnim);
            _Caster.MoveTo(_Targets[0].Forward, delegate
            {
                //clip name in skilldata
                _Caster.Play(_SkillData._CastAnim);
                _Caster.PlayQueue(Define.ANIMATION_PLAYER_ACTION_RUN);


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
                        GameObject effBeattack = AssetLoader.LoadAsset(_SkillData._BeattackEffect);
                        effBeattack.transform.parent = _Targets[i]._ActorObj.transform;
                        effBeattack.transform.localPosition = Vector3.zero;
                        effBeattack.transform.localScale = Vector3.one;
                        effBeattack.SetActive(false);
                        effBeattack.SetActive(true);
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
        }
        else
        {
            //clip name in skilldata
            _Caster.Play(_SkillData._CastAnim);
            _Caster.PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
            //cast effect
            GameObject effCast = AssetLoader.LoadAsset(_SkillData._CastEffect);
            effCast.transform.parent = _Caster._ActorObj.transform;
            effCast.transform.localPosition = Vector3.zero;
            effCast.transform.localScale = Vector3.one;
            effCast.SetActive(false);
            effCast.SetActive(true);
            //1.目标播受击动作和特效的时间
            //2.目标弹伤害数字的时间
            //3.技能总时间
            new Timer().Start(new TimerParam(_SkillData._BeattackTime, delegate
            {
                for (int i = 0; i < _Targets.Length; ++i)
                {
                    _Targets[i].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);

                    //beattack effect
                    GameObject effBeattack = AssetLoader.LoadAsset(_SkillData._BeattackEffect);
                    effBeattack.transform.parent = _Targets[i]._ActorObj.transform;
                    effBeattack.transform.localPosition = Vector3.zero;
                    effBeattack.transform.localScale = Vector3.one;
                    effBeattack.SetActive(false);
                    effBeattack.SetActive(true);
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
        }

        return true;
    }
}
