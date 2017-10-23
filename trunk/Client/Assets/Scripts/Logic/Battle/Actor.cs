using UnityEngine;
using System.Collections.Generic;
using FairyGUI;

public class Actor {

    //场上角色的Obj
    public GameObject _ActorObj;

    public string _Name;

    public string _Title;

    public HeadBar _Headbar;

    public int _RealPosInScene;

    public int _CrtValue;   //当前血量
    public int _MaxValue;  //最大血量

    public float _IncreaseValue;

    public int _StrLv;

    public int _DisplayID;

    public int _EntityID;

    //Animator _Animator;
    Animation _Animation;

    float MOVE_SPEED_BATTLE = 4f;
    float MOVE_SPEED_WORLD = 4f;

    Transform _Pos;

    CallBackHandler _MoveCallBack;

    Vector3 _PrePos;

    long _InstID;

    bool _IsRunning;

    // 身上持有的buff列表
    List<int> _BuffList;

    // 身上持有的任务列表
    List<int> _QuestList;

    Vector3 _DestPosition;

    Vector3 _ForwardPlus;

    public int _BornTurn;

    public delegate void CallBackHandler();

    public Actor(GameObject go, Vector3 pos, long instid, string name, string title, int[] questes, int displayid)
    {
        if (go == null)
        {
            Debug.LogWarning("Actor obj is null.");
            return;
        }
        _ActorObj = go;
        _InstID = instid;
        _ActorObj.transform.position = pos;
        _Name = name;
        _Title = title;
        _DisplayID = displayid;
        if (questes != null)
        {
            _QuestList = new List<int>(questes);
        }
        Init(false);
    }

    public Actor(GameObject go, Transform pos, long instid, int realPos, int crtHp, int maxHp, int entityid, int strLv)
    {
        if (go == null)
        {
            Debug.LogWarning("Actor obj is null.");
            return;
        }
        _ActorObj = go;
        _InstID = instid;
        _Pos = pos;
        _StrLv = strLv;
        _RealPosInScene = realPos;
        _CrtValue = crtHp;
        _MaxValue = maxHp;
        _EntityID = entityid;
        EntityData eData = EntityData.GetData(entityid);
        if (eData != null)
        {
            _DisplayID = eData._DisplayId;
            _Name = eData._Name;
        }
        _ActorObj.transform.position = _Pos.position;
        _ActorObj.transform.rotation = _Pos.rotation;
        Init();
    }

    void Init(bool normal = true)
    {
        _Animation = _ActorObj.GetComponent<Animation>();

        _Headbar = new HeadBar(this, normal? 0: 1);
        _Headbar.Visible = InstID != 0;

        MOVE_SPEED_BATTLE = Define.GetFloat("MoveSpeed_InBattle");
        MOVE_SPEED_WORLD = Define.GetFloat("MoveSpeed_InWorld");

        AdjustHeadBarScale();

        _BornTurn = Battle._Turn;
    }

    public void UpdateValue(int value, int maxValue)
    {
        _CrtValue += value;
        if(maxValue != -1)
            _MaxValue = maxValue;
        _Headbar._IsDirty = true;
    }

    public void SetValue(int value, int maxValue)
    {
        _CrtValue = value;
        _MaxValue = maxValue;
        _Headbar._IsDirty = true;
    }

    public void Update()
    {
        if (_Headbar != null)
            _Headbar.Update();

        if (!Battle.InBattle)
        {
            if (_IsRunning == false)
            {
                if (!IsPlay(Define.ANIMATION_PLAYER_ACTION_IDLE))
                    Play(Define.ANIMATION_PLAYER_ACTION_IDLE);
            }
            else
            {
                if (!IsPlay(Define.ANIMATION_PLAYER_ACTION_WALK))
                    Play(Define.ANIMATION_PLAYER_ACTION_WALK);
            }
        }

        if (_ActorObj != null)
        {
            AdjustHeadBarScale();
        }
    }

    //移动到场上某位置
    public void MoveTo(Vector3 position, CallBackHandler moveToCallback)
    {
        if (_ActorObj == null)
            return;

        _MoveCallBack = moveToCallback;
        _DestPosition = position;
        _ActorObj.transform.LookAt(position);

        //为角色obj添加回调脚本和事件
        int param = LaunchHandler(moveToCallback);

        string moveAnim;
        float speed;
        if (Battle.InBattle)
        {
            moveAnim = Define.ANIMATION_PLAYER_ACTION_RUN;
            speed = MOVE_SPEED_BATTLE;
        }
        else
        {
            moveAnim = Define.ANIMATION_PLAYER_ACTION_WALK;
            speed = MOVE_SPEED_WORLD;
        }

        //Tween position
        iTween.MoveTo(_ActorObj, iTween.Hash("speed", speed, "position", position, "oncomplete", "HandlerFunction", "oncompleteparams", param.ToString(), "easetype", iTween.EaseType.linear));

//        _ActorObj.transform.position = _DestPosition;
//        if (_MoveCallBack != null)
//            _MoveCallBack();

        if (!_IsRunning)
            Play(moveAnim);

        _IsRunning = true;
    }

    public void Stop()
    {
        if(_IsRunning)
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

        if(_Animation.GetClip(action) == null)
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

    public bool IsPlay(string clipName)
    {
        if (_Animation == null)
            return false;

        return _Animation.IsPlaying(clipName);
    }

    public void PlayQueue(string action)
    {
        if (_Animation == null)
            return;

        _Animation.CrossFadeQueued(action);
    }

    public Vector3 Forward(float ajax)
    {
        return _ActorObj.transform.position + _ActorObj.transform.forward * ajax;
    }

    public float ForwardAjax
    {
        get
        {
            DisplayData dData = DisplayData.GetData(_DisplayID);
            return dData._Distance;
        }
    }

    public void DisplaySkill(string skillName)
    {
        _Headbar.DisplaySkill(skillName);
    }

    //Hud操作
    public void PopContent(int value, string special = "")
    {
        _Headbar.PopContent(value, special);
    }

    public void AddBuff(int buffid)
    {
        if (_BuffList == null)
            _BuffList = new List<int>();
//
//        if (!_BuffList.Contains(buffid) /* || 可叠加 读表 */)
            _BuffList.Add(buffid);

        _Headbar._IsDirty = true;
    }

    public void RemoveBuff(int buffid)
    {
        if (_BuffList == null)
            return;

        if (_BuffList.Contains(buffid))
            _BuffList.Remove(buffid);

        _Headbar._IsDirty = true;
    }

    public bool HasQuest
    {
        get
        {
            if(_QuestList == null)
                return false;
            return _QuestList.Count != 0; 
        }
    }

    public List<int> BuffList
    {
        get
        {
            return _BuffList;
        }
    }

    public int BuffCount
    {
        get
        {
            if (_BuffList == null)
                return 0;
            return _BuffList.Count;
        }
    }

    public long InstID
    {
        get{ return _InstID; }
        set{ _InstID = value;}
    }

    void AdjustHeadBarScale()
    {
        if(_ActorObj == null)
            return;
        
        if (_Headbar == null)
            return;

        float dis = Vector3.Distance(Camera.main.transform.position, _ActorObj.transform.position);
        // 14 ===> 1

        float adjScale = dis * 0.0005f;
        _Headbar._Headbar.transform.localScale = new Vector3(adjScale, adjScale, adjScale);
    }

    public void Fini(bool clear = false)
    {
        DisplayData dData = DisplayData.GetData(_DisplayID);
        if (dData == null)
            Debug.LogWarning("There is MEMORY WASTED! " + _DisplayID);
        else
            AssetLoader.UnloadAsset(dData._AssetPath, clear);
        GameObject.Destroy(_ActorObj);
    }
}
