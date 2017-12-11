using System.Collections.Generic;

public class BuffChecker {

    public bool _IsChecking;
    public bool _IsChecked;
    // buff主体
    Actor _Actor;

    List<COM_BattleBuffAction> _BuffCheck;

    public BuffChecker(Actor actor, COM_BattleBuffAction[] buffCheck)
    {
        _Actor = actor;
        if(buffCheck != null)
            _BuffCheck = new List<COM_BattleBuffAction>(buffCheck);
    }
    // 结算
    public void Check()
    {
        if (_Actor == null)
        {
            Clear();
            return;
        }

        if (_BuffCheck == null || _BuffCheck.Count <= 0)
        {
            Clear();
            return;
        }

        BuffData data = BuffData.GetData(_BuffCheck [0].BuffChange.BuffId);
        float maxTime = 0f;
        if (maxTime < data._AnimTime + _Actor.ClipLength(Define.ANIMATION_PLAYER_ACTION_BEATTACK))
            maxTime = data._AnimTime + _Actor.ClipLength(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
        if (maxTime < data._EffectTime)
            maxTime = data._EffectTime;
        if (maxTime < data._EmitTime)
            maxTime = data._EmitTime;

        new Timer().Start(new TimerParam(data._AnimTime, delegate
        {
            if(_Actor != null)
            {
                if(_BuffCheck != null && _BuffCheck.Count > 0 && _BuffCheck[0].BuffData < 0)
                {
                    _Actor.Play(Define.ANIMATION_PLAYER_ACTION_BEATTACK);
                    _Actor.PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
                }
            }
            else
                Clear();
        }), new TimerParam(data._EffectTime, delegate
        {
            
        }), new TimerParam(data._EmitTime, delegate
        {
            if(_Actor != null)
            {
                if(_BuffCheck != null && _BuffCheck.Count > 0 && _BuffCheck[0].BuffData != 0)
                {
                    _Actor.UpdateValue(_BuffCheck[0].BuffData, -1/*_BuffCheck[0].BuffMax*/);
                    _Actor.PopContent(_BuffCheck[0].BuffData);
                }
            }
            else
                Clear();
        }), new TimerParam(maxTime, delegate
        {
            if(_Actor != null)
            {
                if(_BuffCheck != null && _BuffCheck.Count > 0 && _BuffCheck [0].Dead)
                {
                    float deadTime = _Actor.ClipLength(Define.ANIMATION_PLAYER_ACTION_DEAD);
                    _Actor.Play(Define.ANIMATION_PLAYER_ACTION_DEAD);
                    new Timer().Start(deadTime, (object actor) =>
                    {
                        Clear();
                        new Timer().Start(((Actor)actor).ClipLength(Define.ANIMATION_PLAYER_ACTION_DEAD) + 1f, (object ac) => {
                            if(ac != null)
                                Battle.DelActor(((Actor)ac)._RealPosInScene);
                        }, actor);
                    }, _Actor);
                }
                else
                {
                    HandleBuff();
                }
            }
            else
                Clear();
        }));
    }

    void HandleBuff()
    {
        if (_Actor == null)
        {
            Clear();
            return;
        }

        if (_BuffCheck == null || _BuffCheck.Count <= 0)
        {
            Clear();
            return;
        }

        if (_BuffCheck[0].BuffChange.Change == 1)
            _Actor.AddBuff(_BuffCheck[0].BuffChange.BuffId);
        else if (_BuffCheck[0].BuffChange.Change == 0)
            _Actor.RemoveBuff(_BuffCheck[0].BuffChange.BuffId);

        _BuffCheck.RemoveAt(0);
        Check();
    }

    void Clear()
    {
        _IsChecking = false;
        _IsChecked = true;
    }
}
