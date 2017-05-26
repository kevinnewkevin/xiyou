using UnityEngine;
using System.Collections.Generic;

public class Timer {

    public delegate void TimerCallBack();

    List<TimerParam> _TimerParams = new List<TimerParam>();

    public bool _IsDead;

	public void Start(float time, TimerCallBack callback)
    {
        _TimerParams.Add(new TimerParam(time, callback));
        TimerManager.Add(this);
    }

    public void NextFrame(TimerCallBack callback)
    {
        _TimerParams.Add(new TimerParam(2, callback));
        TimerManager.Add(this);
    }

    /// <summary>
    /// Time param, Handler Param, loop.
    /// </summary>
    /// <param name="callbackParam"></param>
    public void Start(params TimerParam[] callbackParam)
    {
        if(callbackParam == null)
            return;
        _TimerParams.AddRange(callbackParam);
        TimerManager.Add(this);
    }

    public void Update()
    {
        if (_TimerParams.Count == 0)
        {
            _IsDead = true;
            return;
        }

        for (int i = 0; i < _TimerParams.Count; )
        {
            _TimerParams[i]._Frame -= 1;
            _TimerParams[i]._Timer -= Time.deltaTime;
            if (_TimerParams[i]._Timer <= 0 && _TimerParams[i]._Frame <= 0)
            {
                if (_TimerParams[i]._Delegate != null)
                    _TimerParams[i]._Delegate();

                _TimerParams.RemoveAt(i);
            }
            else
            {
                ++i;
            }
        }
    }
}

public class TimerParam
{
    public TimerParam(float timer, Timer.TimerCallBack delegat)
    {
        _Timer = timer;
        _Delegate = delegat;
    }
    public TimerParam(int frame, Timer.TimerCallBack delegat)
    {
        _Frame = frame;
        _Delegate = delegat;
    }
    public float _Timer;
    public int _Frame;
    public Timer.TimerCallBack _Delegate;
}
