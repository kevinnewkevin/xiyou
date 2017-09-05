using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OpraActor
{
    public GameObject _Go;
    public ActorData _Data;
    public OpraActor(ActorData data, GameObject go)
    {
        _Data = data;
        _Go = go;
    }
}

public class OpraClip
{
    public string _ClipName;
    public AnimationData _Data;
    public OpraClip(AnimationData data, string clipName)
    {
        _Data = data;
        _ClipName = clipName;
    }
}

public class OpraSystem : MonoBehaviour {

    Animation _Animation;
    Queue<OpraClip> _ClipQue;
    List<OpraActor> _ActorList;

    public OpraClip _CrtScreen;

    public bool _IsPlaying;
	// Use this for initialization
	void Start () {
        if (_Animation == null)
            _Animation = GetComponent<Animation>();

        _IsPlaying = false;

        DontDestroyOnLoad(this);
	}
	
	// Update is called once per frame
	void Update () {
		
	}

    public void Begin(int[] ids)
    {
        AnimationData aData = null;
        ActorData acData = null;
        for(int i=0; i < ids.Length; ++i)
        {
            aData = AnimationData.GetData(ids[i]);
            LaunchClip(aData);

            for (int j = 0; j < aData._Actors.Length; ++j)
            {
                acData = ActorData.GetData(aData._Actors[j]);
                LaunchAsset(acData);
            }
        }
    }

    public void Play()
    {
        if (_CrtScreen != null)
            UnloadClip(_CrtScreen._Data);
        
        if (_Animation == null)
        {
            ReturnToBattle();
            return;
        }

        if (_ClipQue.Count <= 0)
        {
            ReturnToBattle();
            return;
        }

        _CrtScreen = _ClipQue.Dequeue();
        _Animation.Play(_CrtScreen._ClipName);
        SetActorAction(_CrtScreen._Data._Id);
        _IsPlaying = true;
        new Timer().Start(_Animation.GetClip(_CrtScreen._ClipName).length, delegate {
            AnimationData aData = AnimationData.GetData(_CrtScreen._Data._Id);
            if(aData == null)
            {
                Play();
                return;
            }

            if(aData._Talks == null || aData._Talks.Length == 0)
            {
                Play();
            }
            else
            {
                UIManager.Show("duibaikuang");
            }
        });
    }

    void LaunchClip(AnimationData data)
    {
        if(_ClipQue == null)
            _ClipQue = new Queue<OpraClip>();
        AnimationClip clip = AssetLoader.LoadClip(data._ClipName);
        _Animation.AddClip(clip, clip.name);
        _ClipQue.Enqueue(new OpraClip(data, clip.name));
    }

    void LaunchAsset(ActorData data)
    {
        if(_ActorList == null)
            _ActorList = new List<OpraActor>();
        GameObject obj = AssetLoader.LoadAsset(data._Asset);
        obj.name = obj.name.Remove(obj.name.IndexOf("(Clone)"));
        obj.transform.parent = transform;
        obj.transform.localScale = Vector3.one;
        obj.SetActive(false);
        _ActorList.Add(new OpraActor(data, obj));
    }

    void UnloadClip(AnimationData data)
    {
        AssetLoader.UnloadAsset(data._ClipName, true);
    }

    void UnloadAsset(ActorData data)
    {
        AssetLoader.UnloadAsset(data._Asset, true);
    }

    void SetActorAction(int animationid)
    {
        AnimationData aData = AnimationData.GetData(animationid);
        if (aData == null)
            return;

        OpraActor actor;
        ActorData acData;
        AnimHolder aHolder;
        for(int i=0; i < aData._Actors.Length; ++i)
        {
            actor = GetActor(aData._Actors[i]);
            acData = ActorData.GetData(actor._Data._Id);
            for(int j=0; j < acData._Actions.Length; ++j)
            {
                aHolder = actor._Go.GetComponent<AnimHolder>();
                if(aHolder == null)
                    aHolder = actor._Go.AddComponent<AnimHolder>();
                aHolder.Add(acData._Actions[j].Key, acData._Actions[j].Value);
            }
        }
    }

    OpraActor GetActor(int id)
    {
        for(int i=0; i < _ActorList.Count; ++i)
        {
            if (_ActorList [i]._Data._Id == id)
                return _ActorList [i];
        }
        return null;
    }

    void ReturnToBattle()
    {
        _IsPlaying = false;
        for (int i = 0; i < _ActorList.Count; ++i)
        {
            UnloadAsset(_ActorList [i]._Data);
            GameObject.Destroy(_ActorList[i]._Go);
        }
        CameraEffect.Fade(delegate {
            Battle.CurrentState = Battle.BattleState.BS_Oper;
            UIManager.Show("BattlePanel");
            CameraEffect.Continue();
        });
    }
}