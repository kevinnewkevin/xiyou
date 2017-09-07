using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class BattleSceneTouch : MonoBehaviour {

    bool _IsPress;

    float _PreVmag, _PreHmag;

	// Use this for initialization
	void Start () {
        Stage.inst.onTouchBegin.Add(OnTouchBegin);
        Stage.inst.onTouchMove.Add(OnTouchMove);
        Stage.inst.onTouchEnd.Add(OnTouchEnd);
	}

    void OnTouchBegin()
    {
        if (Stage.isTouchOnUI)
            return;

        if (Battle._BattleCamera._IsPlaying)
            return;
        
        _IsPress = true;
        RaycastHit hit;
        Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
        if (Physics.Raycast(ray, out hit))
        {
            if (hit.transform.CompareTag("Point"))
            {
                PointHandle handler = hit.transform.GetComponent<PointHandle>();
                if (handler != null)
                    handler.Excute();
                _IsPress = false;
            }
        }
    }

    void OnTouchMove(EventContext context)
    {
        if (!Stage.isTouchOnUI && !Battle._BattleCamera._IsPlaying)
        {
            if (_IsPress)
            {
                Vector2 pos = ((InputEvent)context.data).position;
                Vector2 hPos = new Vector2(pos.x, 0f);
                Vector2 vPos = new Vector2(0f, pos.y);
                float crtHmag = hPos.magnitude;
                float crtVmag = vPos.magnitude;
                bool hanti = false;
                if (_PreHmag > crtHmag)
                    hanti = true;

                bool vanti = false;
                if (_PreVmag > crtVmag)
                    vanti = true;

                float h = Mathf.Abs(_PreHmag - crtHmag);
                float v = Mathf.Abs(_PreVmag - crtVmag);
                if (h > v)
                    Camera.main.transform.RotateAround(Battle._Center, Camera.main.transform.up, crtVmag * Time.deltaTime * 0.1f * (hanti ? -1f : 1f));
                else if (h < v)
                {
                    float gap = Camera.main.transform.position.y - Battle._Center.y;
                    if(gap >= 6f && vanti)
                        Camera.main.transform.RotateAround(Battle._Center, Camera.main.transform.right, crtHmag * Time.deltaTime * 0.05f * (vanti ? -1f : 1f));
                    else if(gap <= 2f && !vanti)
                        Camera.main.transform.RotateAround(Battle._Center, Camera.main.transform.right, crtHmag * Time.deltaTime * 0.05f * (vanti ? -1f : 1f));
                    else if(gap < 6f && gap > 2f)
                        Camera.main.transform.RotateAround(Battle._Center, Camera.main.transform.right, crtHmag * Time.deltaTime * 0.05f * (vanti ? -1f : 1f));
                }
                Camera.main.transform.LookAt(Battle._Center);

                _PreHmag = crtHmag;
                _PreVmag = crtVmag;
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
