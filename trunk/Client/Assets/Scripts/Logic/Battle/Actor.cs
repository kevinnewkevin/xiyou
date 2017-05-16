using UnityEngine;

public class Actor {

    //场上角色的Obj
    public GameObject _ActorObj;

    Animator _Animator;

    public Actor(GameObject go)
    {
        if (go == null)
        {
            Debug.LogWarning("Actor obj is null.");
            return;
        }
        _ActorObj = go;
        Init();
    }

    void Init()
    {
        _Animator = _ActorObj.GetComponent<Animator>();
        if (_Animator == null)
            Debug.LogWarning("Actor " + _ActorObj.name + " has no Animator launched.");
    }

    //移动到场上某位置
    public void MoveTo()
    {
        if (_ActorObj == null)
            return;

        
    }

    //播放某个动作
    public void Play(string action)
    {
        if (_Animator == null)
            return;

        _Animator.Play(action);
    }

    //Hud操作
    public void PopContent()
    {

    }

    public void Fini()
    {

    }
}
