using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class MainSceneTouch : MonoBehaviour {

    bool isMove;
    bool isPress;
    float preX;

    enum slideVector { nullVector, left, right };  

    private Vector2 lastPos;//上一个位置  

    private Vector2 currentPos;//下一个位置  

    private slideVector lastVector = slideVector.nullVector;//当前滑动方向  

    private slideVector currentVector = slideVector.nullVector;//当前滑动方向  

    private float timer;//时间计数器  

    public float offsetTime = 0.01f;//判断的时间间隔  

    bool isMoveScreen;

	// Use this for initialization
	void Start () {
        World.InitPlayerActor();
        World.InitNpcActor();
//        Stage.inst.onTouchBegin.Add(OnTouchBegin);
//        Stage.inst.onTouchEnd.Add(OnTouchEnd);
//        Stage.inst.onTouchMove.Add(OnTouchMove);
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

    void Update()
    {
        if (Input.GetMouseButtonDown (0))
        {  
            lastPos = Input.mousePosition;  
            currentPos = Input.mousePosition;  
            timer = 0;  
        }
        if (Input.GetMouseButton (0))
        {  
            currentPos = Input.mousePosition;  
            if (currentPos.x < lastPos.x)
            {
                Camera.main.GetComponent<CameraTracker>().MoveToLookAt -= 1f;
                isMoveScreen = true;
            }   
            if (currentPos.x > lastPos.x)
            {
                Camera.main.GetComponent<CameraTracker>().MoveToLookAt += 1f;
                isMoveScreen = true;
            }

            lastPos = currentPos;  
        }
        if (Input.GetMouseButtonUp (0))
        { 
            if (isMoveScreen)
            {
                isMoveScreen = false;
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
            currentVector = slideVector.nullVector;  
        }  
    }

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnTouchBegin);
        Stage.inst.onTouchEnd.Remove(OnTouchEnd);
        Stage.inst.onTouchMove.Remove(OnTouchMove);
    }

//    void OnGUI()
//    {
//        if (GUILayout.Button("加buff"))
//            World.PlayerActor.AddBuff(1);
//
//        if (GUILayout.Button("减buff"))
//            World.PlayerActor.RemoveBuff(1);
//    }
}
