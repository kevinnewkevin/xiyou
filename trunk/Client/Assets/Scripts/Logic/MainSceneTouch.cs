using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class MainSceneTouch : MonoBehaviour {

    bool isMove;
    bool isPress;
    float preX;
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
        
        isPress = true;
        preX = ((FairyGUI.InputEvent)context.data).x;
    }

    void OnTouchEnd()
    {
        isPress = false;
        if (isMove)
        {
            isMove = false;
            return;
        }
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
        if (Stage.isTouchOnUI)
        {
            isPress = false;
            isMove = false;
            return;
        }

        if (!isPress)
            return;
        
        float dir = ((FairyGUI.InputEvent)context.data).x - preX;
        Camera.main.GetComponent<CameraTracker>().MoveToLookAt += dir;
        preX = ((FairyGUI.InputEvent)context.data).x;
        isMove = true;
        Debug.Log(dir);
    }

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnTouchBegin);
        Stage.inst.onTouchEnd.Remove(OnTouchEnd);
        Stage.inst.onTouchMove.Remove(OnTouchMove);
    }
}
