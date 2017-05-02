using System.Collections.Generic;

public class TimerManager {

    static public List<Timer> _Timers = new List<Timer>();

    static public void Update()
    {
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

    static public void Add(Timer timer)
    {
        _Timers.Add(timer);
    }
}
