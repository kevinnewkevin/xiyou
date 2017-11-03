using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class BattleSceneTouch : MonoBehaviour {

    bool _IsPress;
    bool _WantPress;

    float _PreVmag, _PreHmag;

    float _PressTimer;

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

        if (Battle.op != null && Battle.op._IsPlaying)
            return;

        if (!_WantPress)
            _WantPress = true;
        
//        _IsPress = true;
    }

    void OnTouchMove(EventContext context)
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
                Camera.main.transform.RotateAround(Battle._Center, Camera.main.transform.up, crtVmag * Time.deltaTime * 0.05f * (hanti ? -1f : 1f));
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

        if (DragDrop.inst.dragging)
        {
            RaycastHit hit;
            Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
            if (Physics.Raycast(ray, out hit))
            {
                if (hit.transform.CompareTag("Point"))
                {
                    PointHandle handler = hit.transform.GetComponent<PointHandle>();
                    if (handler != null)
                    {
                        DragDrop.inst.dragAgentGraph.visible = false;
                        Battle.SimSetActor(handler._Pos);
                    }
                }
            }
            else
            {
                DragDrop.inst.dragAgentGraph.visible = true;
                Battle.ClearSimActor();
            }
        }
    }

    void OnTouchEnd()
    {
        if (_WantPress)
        {
            _WantPress = false;
            _PressTimer = 0;
        }

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
        _IsPress = false;
    }

    void Update()
    {
        if (_WantPress)
        {
            _PressTimer += Time.deltaTime;
            if (_PressTimer > 0.2f)
            {
                _WantPress = false;
                _IsPress = true;
                _PressTimer = 0;
            }
        }
    }

    public bool IsMoving
    {
        get{ return _IsPress; }
    }

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnTouchBegin);
    }
}
