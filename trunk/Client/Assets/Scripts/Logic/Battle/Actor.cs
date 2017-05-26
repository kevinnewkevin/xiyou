using UnityEngine;
using System.Collections;

public class Actor {

    //场上角色的Obj
    public GameObject _ActorObj;

    ulong _InstID;

    Animator _Animator;

    const float MOVE_SPEED = 4f;

    bool _IsRunning;

    public delegate void CallBackHandler();

    public Actor(GameObject go, Vector3 pos, ulong instid)
    {
        if (go == null)
        {
            Debug.LogWarning("Actor obj is null.");
            return;
        }
        _ActorObj = go;
        _InstID = instid;
        //_ActorObj.transform.position = pos;
        Init();
    }

    void Init()
    {
        _Animator = _ActorObj.GetComponent<Animator>();
        if (_Animator == null)
            Debug.LogWarning("Actor " + _ActorObj.name + " has no Animator launched.");
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

    int LaunchHandler(CallBackHandler callback)
    {
        if (_ActorObj == null)
            return -1;

        GameObjectHandler handler = _ActorObj.GetComponent<GameObjectHandler>();
        if (handler == null)
            handler = _ActorObj.AddComponent<GameObjectHandler>();
        return handler.LaunchHandler(callback);
    }

    public string IsPlaying
    {
        get
        {
            if (_ActorObj == null)
                return "";

            if (_Animator == null)
                return "";

            AnimatorClipInfo[] aci = _Animator.GetCurrentAnimatorClipInfo(0);
            if (aci.Length == 1)
                return aci [0].clip.name;
            return "";
        }
    }

    //播放某个动作
    public void Play(string action, bool bVal = false, float fVal = 0f)
    {
        if (_Animator == null)
            return;

        string[] info = action.Split(new char[]{'_'}, System.StringSplitOptions.RemoveEmptyEntries);
        if (info == null || info.Length < 2)
        {
            Debug.LogWarning("Wrong animation transition parameters " + action);
            return;
        }
        switch(info[0])
        {
            case "b":
                _Animator.SetBool(action, bVal);
                break;
            case "f":
                _Animator.SetFloat(action, fVal);
                break;
            case "t":
                _Animator.SetTrigger(action);
                break;
        }
    }

    //Hud操作
    public void PopContent()
    {

    }

    public long InstID
    {
        get{ return InstID; }
    }

    public void Fini()
    {

    }
}
