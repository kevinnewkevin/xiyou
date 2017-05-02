public class Timer {

    public delegate void TimerCallBack();

    float _Time;

    TimerCallBack _Callback;

	public void Start(float time, TimerCallBack callback)
    {
        _Time = time;
        _Callback = callback;
    }

    public void Update()
    {

    }
}
