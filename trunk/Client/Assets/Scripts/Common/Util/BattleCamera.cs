using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BattleCamera : MonoBehaviour {

    Camera _MainCamera;

    public GameObject _Target;
    Vector3 _OriginPos;
    Quaternion _OriginQuat;

    public bool _IsPlaying;

	// Use this for initialization
	void Start () {
        _MainCamera = Camera.main;
	}

    public void Feature(GameObject target, string type)
    {
        if (!type.Equals("1"))
            return;

        _IsPlaying = true;

        _Target = target;
        _MainCamera.transform.parent = target.transform;
        string[] devPos = Define.GetStr("BattleCamera_plus").Split(new char[]{','}, System.StringSplitOptions.RemoveEmptyEntries);
        Vector3 plusPos = new Vector3(float.Parse(devPos[0]), float.Parse(devPos[1]), float.Parse(devPos[2]));
        _MainCamera.transform.localPosition = plusPos;
        _MainCamera.transform.LookAt(target.transform.position);
        _MainCamera.transform.localScale = Vector3.one;
    }

    void Update()
    {
        if (_IsPlaying)
            return;
        
        _OriginPos = _MainCamera.transform.position;
        _OriginQuat = _MainCamera.transform.rotation;
    }

    public void Reset()
    {
        _MainCamera.transform.parent = null;
        _MainCamera.transform.position = _OriginPos;
        _MainCamera.transform.rotation = _OriginQuat;
        _MainCamera.transform.localScale = Vector3.one;
        _IsPlaying = false;
    }
}
