using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;

public class CameraTracker : MonoBehaviour {

    private float _LookAtPos;

    private bool _CanMove;

    float _Speed;

    public float smoothTime = 0.01f;  //摄像机平滑移动的时间

    private Vector3 cameraVelocity = Vector3.zero;

    private Camera mainCamera;

    Vector3 _PlusPos;
    GameObject _FocusObject;
    Vector3 _PosSnapShot;
    Quaternion _RotSnapShot;

    public float MoveToLookAt
    {
        set
        {
            _LookAtPos = value;
            _CanMove = true;
        }
        get
        {
            return transform.position.x;
        }
    }

    public void Focus(GameObject go)
    {
        _FocusObject = go;
        // disable obj's trigger
        BoxCollider bc = go.GetComponent<BoxCollider>();
        if (bc != null)
            bc.enabled = false;

        // snapshot preTransform info
        _PosSnapShot = transform.position;
        _RotSnapShot = transform.rotation;

        World._DisableMainSceneOperate = true;

        Vector3 destPos = go.transform.position + _PlusPos;
        iTween.MoveTo(gameObject, iTween.Hash("time", 0.6f, "position", destPos, "oncomplete", "Focused", "easetype", iTween.EaseType.linear));
    }

    public void CancelFocus()
    {
        if (_FocusObject != null)
        {
            BoxCollider bc = _FocusObject.GetComponent<BoxCollider>();
            if (bc != null)
                bc.enabled = true;
        }
        
        iTween.MoveTo(gameObject, iTween.Hash("time", 0.6f, "position", _PosSnapShot, "oncomplete", "CancelFocused", "easetype", iTween.EaseType.linear));
    }

    void Awake()
    {
        mainCamera = Camera.main;

        string[] devPos = Define.GetStr("WorldCamera_focusPlus").Split(new char[]{','}, System.StringSplitOptions.RemoveEmptyEntries);
        _PlusPos = new Vector3(float.Parse(devPos[0]), float.Parse(devPos[1]), float.Parse(devPos[2]));
    }

	// Use this for initialization
	void Start () {
        _Speed = Define.GetFloat("MoveSpeed_InWorld");
	}
	
	// Update is called once per frame
	void Update () {
        if (!_CanMove)
            return;
        
        Vector3 pos = new Vector3(_LookAtPos, transform.position.y, transform.position.z);
        iTween.MoveTo(gameObject, iTween.Hash("speed", _Speed, "position", pos, "oncomplete", "Moved", "easetype", iTween.EaseType.linear));
	}

    void Moved()
    {
        _CanMove = false;
    }

    void Focused()
    {
        
    }

    void CancelFocused()
    {
        transform.rotation = _RotSnapShot;
        World._DisableMainSceneOperate = false;
    }
}
