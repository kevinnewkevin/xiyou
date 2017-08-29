using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class MainSceneTouch : MonoBehaviour {

    bool _IsPress;
    bool _IsMoved;

    float _Premag;

	// Use this for initialization
	void Start () {
        World.InitPlayerActor();
        World.InitNpcActor();
        Stage.inst.onTouchBegin.Add(OnTouchBegin);
        Stage.inst.onTouchEnd.Add(OnTouchEnd);
        Stage.inst.onTouchMove.Add(OnTouchMove);
	}

    void OnTouchBegin(EventContext context)
    {
        if (Stage.isTouchOnUI)
            return;
        
        _Premag = ((InputEvent)context.data).position.magnitude;
        _IsPress = true;
        _IsMoved = false;
    }

    void OnTouchEnd()
    {
        _IsPress = false;
        if (_IsMoved)
            return;

        if (!Stage.isTouchOnUI)
        {
            RaycastHit hit;
            Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
            if (Physics.Raycast(ray, out hit))
            {
                if (hit.transform.name.Equals("Cube"))
                {
                    Camera.main.GetComponent<CameraTracker>().MoveToLookAt = hit.point.x;
                    World.PlayerActor.MoveTo(new Vector3(hit.point.x, -14.31f, hit.point.z), delegate {
                        World.PlayerActor.Stop();
                    });
                }

                if (hit.transform.CompareTag("Npc"))
                {
                    NpcHandle handler = hit.transform.GetComponent<NpcHandle>();
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
                Camera.main.transform.Translate(new Vector3(_Premag - crtmag, 0f, 0f) * Time.deltaTime * 0.5f);
                _Premag = crtmag;
                _IsMoved = true;
            }
        }
    }

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnTouchBegin);
        Stage.inst.onTouchEnd.Remove(OnTouchEnd);
        Stage.inst.onTouchMove.Remove(OnTouchMove);
    }
}
