using UnityEngine;

public class Timer {

    public delegate void TimerCallBack();

    float _Time;

    TimerCallBack _Callback;

    public bool _IsDead;

	public void Start(float time, TimerCallBack callback)
    {
        _Time = time;
        _Callback = callback;
        TimerManager.Add(this);
    }

    public void Update()
    {
        _Time -= Time.deltaTime;
        if (_Time <= 0)
        {
            if (_Callback != null)
                _Callback();
            _IsDead = true;
        }
    }
}
