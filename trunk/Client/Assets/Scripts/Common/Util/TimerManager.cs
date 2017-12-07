using System.Collections.Generic;
using UnityEngine;
using System;

public class TimerManager {

    static float _Ticker;

    static public List<Timer> _Timers = new List<Timer>();

    static public List<string> _CountDownKeys = new List<string>();
    static public List<long> _CountDownValues = new List<long>();

    static public void Update()
    {
        _Ticker += Time.deltaTime;
        if (_Ticker >= 1f)
        {
            for(int i=0; i < _CountDownValues.Count; ++i)
            {
                _CountDownValues [i] -= 1;
            }
            _Ticker = _Ticker - 1f;
        }

        for (int i = 0; i < _Timers.Count; )
        {
            if (_Timers[i]._IsDead)
            {
                _Timers.RemoveAt(i);
            }
            else
            {
                _Timers[i].Update();
                ++i;
            }
        }
    }

    static public void AddCountDown(string key, long timeStamp)
    {
        if (_CountDownKeys.Contains(key))
            _CountDownValues [_CountDownKeys.IndexOf(key)] = timeStamp;
        else
        {
            _CountDownKeys.Add(key);
            _CountDownValues.Add(timeStamp);
        }
    }

    static public string GetCountDown(string key)
    {
        long t = 0;
        long h, m, s;
        
        if (_CountDownKeys.Contains(key))
            t = _CountDownValues [_CountDownKeys.IndexOf(key)];
        h = t / 3600;
        m = (t - h * 3600) / 60;
        s = (t - h * 3600 - m * 60);
        return string.Format("{0}:{1}:{2}", h, m, s);
    }

    static public int GetCountDownSecond(string key)
    {
        if (_CountDownKeys.Contains(key))
            return (int)_CountDownValues [_CountDownKeys.IndexOf(key)];
        return 0;
    }

    public static long LeftTimeInSecond(long originTimeStamp)
    {
        TimeSpan ts = DateTime.UtcNow - new DateTime(1970, 1, 1, 0, 0, 0, 0);
        long left = (long)ts.TotalSeconds - originTimeStamp;
        return left;
    }

    public static long GetTimeStamp()
    {
        TimeSpan ts = DateTime.UtcNow - new DateTime(1970, 1, 1, 0, 0, 0, 0);
        return (long)ts.TotalSeconds;
    }

    public static string TimeAgo(long timeStamp)
    {
        long gap = GetTimeStamp() - timeStamp;
        DateTime dt = new DateTime(gap, DateTimeKind.Utc);
        if (dt.Month > 0)
            return dt.Month + "个月前";
        if (dt.Day > 0)
            return dt.Day + "天前";
        if (dt.Hour > 0)
            return dt.Hour + "小时前";
        if (dt.Minute > 0)
            return dt.Minute + "分钟前";
        return "";
    }

    public static void SetTickerGap(long timegap)
    {
        for(int i=0; i < _CountDownValues.Count; ++i)
        {
            _CountDownValues [i] -= timegap;
        }
    }

    static public void Add(Timer timer)
    {
        _Timers.Add(timer);
    }
}
