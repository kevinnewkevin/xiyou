//#define	SHOW_MEMORY_USE

using UnityEngine;
using System;
using System.Collections;
 
 public class DisplayFPS : MonoBehaviour
 {
	public  float updateInterval = 0.5F;
 
	private float				accum   = 0;	// FPS accumulated over the interval
	private int					frames  = 0;	// Frames drawn over the interval
	private float				timeleft;		// Left time for current interval
	private	GUIStyle		labelStyle = new GUIStyle();
	private	GUIStyle		memStyle = new GUIStyle();
    private GUIStyle        pingStyle = new GUIStyle();
	private	string			szFPS = "FPS";
	//
	private	float			showProfileTime	=	0f;
	private	string			szMonoHeapSize	=	"";
	private	string			szMonoUsedSize = "";
	private	string			szUsedHeapSize = "";
	private	string			szTotalAllocatedMemory = "";
	private	string			szTotalUnusedReservedMemory = "";
	private	string			szTotalReservedMemory = "";
	private	string			szGCTotalMemory = "";	
	private	GUIStyle		profileStyle	=	new GUIStyle();
    //
	static	private	float	AverageFPS = 0f;
	static	private	bool	SampleFPS = false;
	static	private	int		SampleTimes = 0;
	static	private	float	TotalSampleFPS = 0f;

    string _time = string.Empty;

    static private float pingDelay = 0f;

    public delegate void TimeUpdateHandler(string time);
    public delegate void BatteryUpdateHandler(float battery);
    public delegate void PingUpdateHandler(float lag);

    public static event TimeUpdateHandler OnTimeUpdate;
	
	// Must call Begin/End sample
	static	public	float SampleAverageFPS
	{
		get	{ return AverageFPS; }
	}
	//
	static	public	void	BeginSampleAverageFPS()
	{
		SampleTimes		=	0;
		AverageFPS		=	0f;
		TotalSampleFPS	=	0f;
		SampleFPS		=	true;
	}
	//
	static	public	void	EndSampleAverageFPS()
	{
		SampleFPS	=	false;
		if( SampleTimes > 0 )
		{
			AverageFPS	=	TotalSampleFPS / (float)SampleTimes;
		}		
	}
	
	void Start()
	{
	    timeleft = updateInterval;  
		labelStyle.fontSize	=	20;
		memStyle.fontSize	=	20;
        pingStyle.fontSize = 20;
		profileStyle.fontSize	=	20;
		profileStyle.normal.textColor = Color.green;

        StartCoroutine("UpdataTime");
    }

    IEnumerator UpdataTime()
    {
        DateTime now = DateTime.Now;
        _time = string.Format("{0:HH:mm}", now);
       // yield return new WaitForSeconds(60f - now.Second);
        while (true)
        {
            now = DateTime.Now;
            _time = string.Format("{0:HH:mm}", now);
            yield return new WaitForSeconds(1f);
            if (OnTimeUpdate != null)
                OnTimeUpdate(_time);
        }
    }
	void Update()
	{
        timeleft -= Time.deltaTime;
        accum += Time.timeScale / Time.deltaTime;
	    ++frames;
	   
	    // Interval ended - update GUI text and start new interval
	    if( timeleft <= 0.0 )
	    {
	        // display two fractional digits (f2 format)
		    float fps = accum/frames;
			
			    szFPS = string.Format("{0:F2} FPS",fps);			
			    if(fps < 10)
				{
					labelStyle.normal.textColor = Color.red;
				}
			    else if(fps < 30)
				{
					labelStyle.normal.textColor = Color.yellow;
				}
			    else
				{
					labelStyle.normal.textColor = Color.green;
				}
				
			
		    timeleft += updateInterval;
		    accum = 0.0F;
		    frames = 0;
			
			//
			if( SampleFPS )
			{
				TotalSampleFPS	+= fps;
				SampleTimes++;
			}
        }
	}
	
	void OnGUI()
    {
        GUI.Label(new Rect(10, 2, 80, 20), szFPS, labelStyle);

        GUI.Label(new Rect(120, 2, 80, 20), string.Format("Ping:{0:N0}ms", pingDelay), pingStyle);

		showProfileTime	-= Time.deltaTime;
		if( showProfileTime < 0f )
		{
			showProfileTime	=	1f;
            szMonoHeapSize	= string.Format("MonoHeapSize : {0}Mb" , (UnityEngine.Profiling.Profiler.GetMonoHeapSizeLong() / 1048576).ToString());
            szMonoUsedSize = "MonoUsedSize : " + (UnityEngine.Profiling.Profiler.GetMonoUsedSizeLong() / 1048576).ToString();
            szUsedHeapSize = "UsedHeapSize : " + (UnityEngine.Profiling.Profiler.usedHeapSizeLong / 1048576).ToString();
            szTotalAllocatedMemory = "TotalAllocatedMem : " + (UnityEngine.Profiling.Profiler.GetTotalAllocatedMemoryLong() / 1048576).ToString();
            szTotalUnusedReservedMemory = "UnusedReservedMem : " + (UnityEngine.Profiling.Profiler.GetTotalUnusedReservedMemoryLong() / 1048576).ToString();
            szTotalReservedMemory = "TotalReservedMem : " + (UnityEngine.Profiling.Profiler.GetTotalReservedMemoryLong() / 1048576).ToString();
			szGCTotalMemory = "GC.TotalMem : " + (System.GC.GetTotalMemory(false) / 1048576).ToString();
		}
		
		GUI.Label( new Rect(10,20, 200, 20),  szMonoHeapSize, profileStyle );
		GUI.Label( new Rect(10,40, 200, 20),  szMonoUsedSize, profileStyle );
		GUI.Label( new Rect(10,60, 200, 20),  szUsedHeapSize, profileStyle );
		GUI.Label( new Rect(10,80, 200, 20),  szTotalAllocatedMemory, profileStyle );
		GUI.Label( new Rect(10,100, 200, 20),  szTotalUnusedReservedMemory, profileStyle );
		GUI.Label( new Rect(10,120, 200, 20),  szTotalReservedMemory, profileStyle );
		GUI.Label( new Rect(10,140, 200, 20),  szGCTotalMemory, profileStyle );

        GUI.Label(new Rect(10, 160, 200, 20), _time, profileStyle);
	}
 }