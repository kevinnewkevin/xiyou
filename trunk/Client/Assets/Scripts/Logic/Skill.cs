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
            float attackTime = _Caster.ClipLength(Define.ANIMATION_PLAYER_ACTION_ATTACK);
            _Caster.MoveTo(_Targets[0].Forward, delegate
            {
                //clip name in skilldata
                _Caster.Play(Define.ANIMATION_PLAYER_ACTION_ATTACK);
                _Caster.PlayQueue(Define.ANIMATION_PLAYER_ACTION_RUN);

                //1.目标播受击动作的时间
                //2.目标播受击特效的时间
                //3.目标弹伤害数字的时间
                new Timer().Start(new TimerParam(0.3f, delegate
                {
                    for (int i = 0; i < _Targets.Length; ++i)
                    {
                        _Targets[i].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                        _Targets[i].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
                    }
                }), new TimerParam(1f, delegate
                {

                }), new TimerParam(1f, delegate
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
            _Caster.Play("");

            //1.目标播受击动作的时间
            //2.目标播受击特效的时间
            //3.目标弹伤害数字的时间
            new Timer().Start(new TimerParam(1f, delegate
            {
                for (int i = 0; i < _Targets.Length; ++i)
                {
                    _Targets[i].Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                }
            }), new TimerParam(1f, delegate
            {

            }), new TimerParam(1f, delegate
            {
                for (int i = 0; i < _Targets.Length; ++i)
                {
                    _Targets[i].PopContent();
                }
            }));
        }

        return true;
    }
}
