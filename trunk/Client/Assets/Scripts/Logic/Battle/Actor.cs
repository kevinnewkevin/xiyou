using UnityEngine;
using System.Collections;
using FairyGUI;

public class Actor {

    //场上角色的Obj
    public GameObject _ActorObj;

    public GameObject _Headbar;

    Transform _Pos;

    long _InstID;

    //Animator _Animator;
    Animation _Animation;

    const float MOVE_SPEED = 4f;

    bool _IsRunning;

    public delegate void CallBackHandler();

    public Actor(GameObject go, Vector3 pos, long instid)
    {
        if (go == null)
        {
            Debug.LogWarning("Actor obj is null.");
            return;
        }
        _ActorObj = go;
        _InstID = instid;
        _ActorObj.transform.position = pos;
        Init();
    }

    public Actor(GameObject go, Transform pos, long instid)
    {
        if (go == null)
        {
            Debug.LogWarning("Actor obj is null.");
            return;
        }
        _ActorObj = go;
        _InstID = instid;
        _Pos = pos;
        _ActorObj.transform.position = _Pos.position;
        _ActorObj.transform.rotation = _Pos.rotation;
        Init();
    }

    void Init()
    {
        _Animation = _ActorObj.GetComponent<Animation>();

        UIPackage.AddPackage("UI/EmitNumbers");
        Define.LaunchUIBundle("xuetiao");
        _Headbar = new GameObject();
        _Headbar.AddComponent<Billboard>();
        UIPanel headbarpanel = _Headbar.AddComponent<UIPanel>();
        headbarpanel.componentName = "xuetiao_com";
        headbarpanel.packageName = "xuetiao";
        _Headbar.transform.parent = _ActorObj.transform;
        _Headbar.transform.localScale = Vector3.one;
        _Headbar.transform.localPosition = new Vector3(0f, 2f, 0f);
    }

    //移动到场上某位置
    public void MoveTo(Vector3 position, CallBackHandler moveToCallback)
    {
        if (_ActorObj == null)
            return;

        _ActorObj.transform.LookAt(position);

        //为角色obj添加回调脚本和事件
        int param = LaunchHandler(moveToCallback);

        //Tween position
        iTween.MoveTo(_ActorObj, iTween.Hash("speed", MOVE_SPEED, "position", position, "oncomplete", "HandlerFunction", "oncompleteparams", param.ToString(), "easetype", iTween.EaseType.linear));

        if (!_IsRunning)
            Play(Define.ANIMATION_PLAYER_ACTION_RUN);

        _IsRunning = true;
    }

    public void Stop()
    {
        if (_IsRunning)
            Play(Define.ANIMATION_PLAYER_ACTION_IDLE);

        _IsRunning = false;
    }

    public void Reset()
    {
        if(_Pos != null)
            _ActorObj.transform.rotation = _Pos.rotation;
    }

    int LaunchHandler(CallBackHandler callback)
    {
        if (_ActorObj == null)
            return -1;

        GameObjectHandler handler = _ActorObj.GetComponent<GameObjectHandler>();
        if (handler == null)
            handler = _ActorObj.AddComponent<GameObjectHandler>();
        return handler.LaunchHandler(callback);
    }

    public bool IsPlaying
    {
        get
        {
            if (_ActorObj == null)
                return false;
        
            if (_Animation == null)
                return false;
        
            return _Animation.isPlaying;
        }
    }

    //播放某个动作
    public void Play(string action)
    {
        if (_Animation == null)
            return;

        if (_Animation.IsPlaying(action))
            _Animation.Rewind(action);
        else
            _Animation.CrossFade(action);
    }

    public float ClipLength(string clipName)
    {
        if (_Animation == null)
            return 0f;

        AnimationClip clip = _Animation.GetClip(clipName);
        if (clip == null)
            return 0f;

        return clip.length;
    }

    public void PlayQueue(string action)
    {
        if (_Animation == null)
            return;

        _Animation.CrossFadeQueued(action);
    }

    public Vector3 Forward
    {
        get{ return _ActorObj.transform.position + _ActorObj.transform.forward; }
    }

    //Hud操作
    public void PopContent(int value)
    {
        EmitManager.inst.Emit(_ActorObj.transform, 0, value, UnityEngine.Random.Range(0, 10) == 5);
    }

    public long InstID
    {
        get{ return _InstID; }
    }

    public void Fini()
    {

    }
}
