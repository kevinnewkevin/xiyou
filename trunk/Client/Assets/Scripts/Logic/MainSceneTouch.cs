using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class MainSceneTouch : MonoBehaviour {

    bool _IsPress;
    bool _IsMoved;

    float _Premag;

    Timer _DestLightTimer;

    GameObject _DestLight;

	// Use this for initialization
	void Start () {
        World.InitPlayerActor(GamePlayer._Data);
        World.InitNpcActor();
        Stage.inst.onTouchBegin.Add(OnTouchBegin);
        Stage.inst.onTouchEnd.Add(OnTouchEnd);
        Stage.inst.onTouchMove.Add(OnTouchMove);

        UIManager.Show("zhujiemian");
	}

    void OnTouchBegin(EventContext context)
    {
        if (World._DisableMainSceneOperate)
        {
            bool isTouchOnNpc = false;
            RaycastHit hit;
            Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
            if (Physics.Raycast(ray, out hit))
                isTouchOnNpc = hit.transform.CompareTag("Npc");

            if (Stage.isTouchOnUI || !isTouchOnNpc)
                Camera.main.GetComponent<CameraTracker>().CancelFocus();

            if (isTouchOnNpc)
            {
                NpcHandle handler = hit.transform.GetComponent<NpcHandle>();
                if (handler != null)
                    handler.Excute();
            }
            return;
        }
        
        if (Stage.isTouchOnUI)
            return;
        
        _Premag = ((InputEvent)context.data).position.magnitude;
        _IsPress = true;
        _IsMoved = false;
    }

    void OnTouchEnd()
    {
        if (World._DisableMainSceneOperate)
            return;
        
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
                    if (_DestLight == null)
                    {
                        string dlpath = Define.GetStr("DestLight");
                        _DestLight = AssetLoader.LoadAsset(dlpath);
                    }
                    else
                        _DestLight.SetActive(false);
                    Vector3 dest = new Vector3(hit.point.x, World._GroudHeight, hit.point.z);
                    _DestLight.transform.position = dest;
                    _DestLight.SetActive(true);
                    Camera.main.GetComponent<CameraTracker>().MoveToLookAt = hit.point.x;
                    World.PlayerActor.MoveTo(dest, delegate {
                        World.PlayerActor.Stop();
                    });

                    if (_DestLightTimer == null || _DestLightTimer._IsDead)
                    {
                        _DestLightTimer = new Timer();
                        _DestLightTimer.Start(1f, delegate
                        {
                            if (_DestLight != null)
                                _DestLight.SetActive(false);
                        });
                    }
                    else
                    {
                        _DestLightTimer.Reset(1f, delegate {
                            if (_DestLight != null)
                                _DestLight.SetActive(false);
                        });
                    }
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
        if (World._DisableMainSceneOperate)
            return;
        
        if (!Stage.isTouchOnUI)
        {
            if (_IsPress)
            {
                float crtmag = ((InputEvent)context.data).position.magnitude;
                if (Mathf.Abs(_Premag - crtmag) < 20)//消除抖动
                    return;
                
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
