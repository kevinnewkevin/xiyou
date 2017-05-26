using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;

public class CameraTracker : MonoBehaviour {

    public GameObject _FollowTarget;

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
        if (!_FollowTarget)
            return;
        
        Vector3 pos = new Vector3(_FollowTarget.transform.position.x, transform.position.y, transform.position.z);
        iTween.MoveTo(gameObject, pos, 8f);
	}
}
