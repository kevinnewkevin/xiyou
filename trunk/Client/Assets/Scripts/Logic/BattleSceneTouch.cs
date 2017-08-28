using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class BattleSceneTouch : MonoBehaviour {

    bool _IsPress;

    float _Premag;

	// Use this for initialization
	void Start () {
        Stage.inst.onTouchBegin.Add(OnTouchBegin);
        Stage.inst.onTouchMove.Add(OnTouchMove);
        Stage.inst.onTouchEnd.Add(OnTouchEnd);
	}

    void OnTouchBegin()
    {
        _IsPress = true;
        if (!Stage.isTouchOnUI)
        {
            RaycastHit hit;
            Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
            if (Physics.Raycast(ray, out hit))
            {
                if (hit.transform.CompareTag("Point"))
                {
                    PointHandle handler = hit.transform.GetComponent<PointHandle>();
                    if (handler != null)
                        handler.Excute();
                }
            }
        }
    }

    void OnTouchMove(EventContext context)
    {
        if (!Stage.isTouchOnUI)
        {
            if (_IsPress)
            {
                float crtmag = ((InputEvent)context.data).position.magnitude;
                bool anti = false;
                if (_Premag > crtmag)
                    anti = true;
                _Premag = crtmag;
                Camera.main.transform.RotateAround(Battle._Center, Vector3.up, crtmag * Time.deltaTime * 0.1f * (anti? -1f: 1f));
                Camera.main.transform.LookAt(Battle._Center);
            }
        }
    }

    void OnTouchEnd()
    {
        _IsPress = false;
        if (!Stage.isTouchOnUI)
        {

        }
    }

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnTouchBegin);
    }
}
