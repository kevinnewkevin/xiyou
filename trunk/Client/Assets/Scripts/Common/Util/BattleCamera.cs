using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BattleCamera : MonoBehaviour {

    Camera _MainCamera;

    public GameObject _Target;

    public bool _Melee;
    public bool _Range;

    public Vector3 _MeleeTargetPlus;
    public float _MeleeTargetTime;

    public Vector3 _RangeTargetPlus = new Vector3(-4.8f, 2.7f, 8f);
    public float _RangeTargetTime = 1f;

    Vector3 _OriginPos;
    Quaternion _OriginQuat;

	// Use this for initialization
	void Start () {
        _MainCamera = Camera.main;
        _OriginPos = _MainCamera.transform.position;
        _OriginQuat = _MainCamera.transform.rotation;
	}
	
    public void Melee(GameObject target)
    {
        _Target = target;
        iTween.MoveTo(_MainCamera.gameObject, iTween.Hash("time", _MeleeTargetTime, "position", target.transform.position + _MeleeTargetPlus, "looktarget", target.transform.position, "easetype", iTween.EaseType.linear));
    }

    public void Range(GameObject target)
    {
        _Target = target;
        iTween.MoveTo(_MainCamera.gameObject, iTween.Hash("time", _RangeTargetTime, "position", target.transform.position + _RangeTargetPlus, "looktarget", target.transform.position, "easetype", iTween.EaseType.linear));
    }

    public void Reset()
    {
        iTween.MoveTo(_MainCamera.gameObject, iTween.Hash("time", _RangeTargetTime, "position", _OriginPos, "looktarget", _Target.transform.position, "easetype", iTween.EaseType.linear));
        _MainCamera.transform.rotation = _OriginQuat;
    }
}
