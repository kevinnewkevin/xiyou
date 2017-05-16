using UnityEngine;
using System.Collections.Generic;

public class Timer {

    public delegate void TimerCallBack();

    List<float> _Time = new List<float>();

    List<TimerCallBack> _Callback = new List<TimerCallBack>();

    public bool _IsDead;

	public void Start(float time, TimerCallBack callback)
    {
        _Time.Add(time);
        _Callback.Add(callback);
        TimerManager.Add(this);
    }

    /// <summary>
    /// Time param, Handler Param, loop.
    /// </summary>
    /// <param name="callbackParam"></param>
    public void Start(params object[] callbackParam)
    {
        if(callbackParam == null)
            return;

        if (callbackParam.Length % 2 == 1)
            return;

        for (int i = 0; i < callbackParam.Length; i+=2)
        {
            if (callbackParam[i] is float == false)
                return;

            _Time.Add((float)callbackParam[i]);

            if (callbackParam[i + 1] is TimerCallBack == false)
                return;
            
            _Callback.Add((TimerCallBack)callbackParam[i+1]);
        }

        TimerManager.Add(this);
    }

    public void Update()
    {
        if (_Callback.Count != _Time.Count)
        {
            _IsDead = true;
            return;
        }

        for (int i = 0; i < _Callback.Count; )
        {
            _Time[i] -= Time.deltaTime;
            if (_Time[i] <= 0)
            {
                if (_Callback[i] != null)
                    _Callback[i]();

                _Callback.RemoveAt(i);
                _Time.RemoveAt(i);
            }
            else
            {
                ++i;
            }
        }

        if (_Callback.Count == 0 || _Time.Count == 0)
        {
            _IsDead = true;
            return;
        }
    }
}
