using UnityEngine;
using System.Collections.Generic;

public class Timer {

    public delegate void TimerCallBack1();
    public delegate void TimerCallBack2(object oparam);

    List<TimerParam> _TimerParams = new List<TimerParam>();

    public bool _IsDead;

    public void Start(float time, TimerCallBack1 callback)
    {
        _TimerParams.Add(new TimerParam(time, callback));
        TimerManager.Add(this);
    }

    public void NextFrame(TimerCallBack1 callback)
    {
        _TimerParams.Add(new TimerParam(2, callback));
        TimerManager.Add(this);
    }

    public void Start(float time, TimerCallBack2 callback, object oparam = null)
    {
        _TimerParams.Add(new TimerParam(time, callback, oparam));
        TimerManager.Add(this);
    }

    public void NextFrame(TimerCallBack2 callback, object oparam = null)
    {
        _TimerParams.Add(new TimerParam(2, callback, oparam));
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

    public void Reset(float time, TimerCallBack1 callback)
    {
        _TimerParams.Clear();
        _TimerParams.Add(new TimerParam(time, callback));
    }

    public void Reset(float time, TimerCallBack2 callback, object oparam = null)
    {
        _TimerParams.Clear();
        _TimerParams.Add(new TimerParam(time, callback, oparam));
    }


    public void Reset(params TimerParam[] callbackParam)
    {
        if(callbackParam == null)
            return;

        _TimerParams.Clear();
        _TimerParams.AddRange(callbackParam);
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
                if (_TimerParams[i]._Delegate1 != null)
                    _TimerParams[i]._Delegate1();
                if (_TimerParams[i]._Delegate2 != null)
                    _TimerParams[i]._Delegate2(_TimerParams[i]._OParam);

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
    public TimerParam(float timer, Timer.TimerCallBack1 delegat)
    {
        _Timer = timer;
        _Delegate1 = delegat;
    }
    public TimerParam(int frame, Timer.TimerCallBack1 delegat)
    {
        _Frame = frame;
        _Delegate1 = delegat;
    }

    public TimerParam(float timer, Timer.TimerCallBack2 delegat, object oparam = null)
    {
        _Timer = timer;
        _Delegate2 = delegat;
        _OParam = oparam;
    }
    public TimerParam(int frame, Timer.TimerCallBack2 delegat, object oparam = null)
    {
        _Frame = frame;
        _Delegate2 = delegat;
        _OParam = oparam;
    }
    public float _Timer;
    public int _Frame;
    public Timer.TimerCallBack1 _Delegate1;
    public Timer.TimerCallBack2 _Delegate2;
    public object _OParam;
}
