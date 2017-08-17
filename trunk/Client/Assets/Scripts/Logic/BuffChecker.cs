using System.Collections.Generic;

public class BuffChecker {

    public bool _IsChecking;
    public bool _IsChecked;
    // buff主体
    Actor _Actor;

    List<COM_BattleBuffAction> _BuffCheck;

    COM_BattleBuff[] _BuffChange;

    public BuffChecker(Actor actor, COM_BattleBuffAction[] buffCheck, COM_BattleBuff[] buffChange)
    {
        _Actor = actor;
        _BuffCheck = new List<COM_BattleBuffAction>(buffCheck);
        _BuffChange = buffChange;
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
            HandleBuff();
            return;
        }

        BuffData data = BuffData.GetData(_BuffCheck [0].BuffId);
        float maxTime = 0f;
        if (maxTime < data._AnimTime + _Actor.ClipLength(data._Anim))
            maxTime = data._AnimTime + _Actor.ClipLength(data._Anim);
        if (maxTime < data._EffectTime)
            maxTime = data._EffectTime;
        if (maxTime < data._EmitTime)
            maxTime = data._EmitTime;

        new Timer().Start(new TimerParam(data._AnimTime, delegate
        {
            _Actor.Play(data._Anim);
        }), new TimerParam(data._EffectTime, delegate
        {
            
        }), new TimerParam(data._EmitTime, delegate
        {
            _Actor.PopContent(_BuffCheck[0].BuffData);
        }), new TimerParam(maxTime, delegate
        {
            if(_BuffCheck [0].Dead)
            {
                float deadTime = _Actor.ClipLength(Define.ANIMATION_PLAYER_ACTION_DEAD);
                _Actor.Play(Define.ANIMATION_PLAYER_ACTION_DEAD);
                new Timer().Start(deadTime, delegate
                {
                    Clear();
                });
            }
            else
            {
                _BuffCheck.RemoveAt(0);
                Check();
            }
        }));
    }

    void HandleBuff()
    {
        if (_Actor == null)
            return;

        for(int i=0; i < _BuffChange.Length; ++i)
        {
            if(_BuffChange[i].Change)
                _Actor.AddBuff(_BuffChange[i].BuffId);
            else
                _Actor.RemoveBuff(_BuffChange[i].BuffId);
        }
        Clear();
    }

    void Clear()
    {
        _IsChecking = false;
        _IsChecked = true;
    }
}
