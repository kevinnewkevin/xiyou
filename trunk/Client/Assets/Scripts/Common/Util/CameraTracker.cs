using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;

public class CameraTracker : MonoBehaviour {

    private float _LookAtPos;

    private bool _CanMove;

    public float smoothTime = 0.01f;  //摄像机平滑移动的时间

    private Vector3 cameraVelocity = Vector3.zero;

    private Camera mainCamera;

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

    void Awake()
    {
        mainCamera = Camera.main;
    }

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
        if (!_CanMove)
            return;
        
        Vector3 pos = new Vector3(_LookAtPos, transform.position.y, transform.position.z);
        iTween.MoveTo(gameObject, iTween.Hash("speed", 4f, "position", pos, "oncomplete", "Moved", "easetype", iTween.EaseType.linear));
	}

    void Moved()
    {
        _CanMove = false;
    }
}
