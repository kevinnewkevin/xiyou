using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class CreateRoleSceneTouch : MonoBehaviour {

    Actor[] _Actors;

    Vector3 _MalePos, _FemalePos, _ToPos;
    float _MaleRotY, _FemaleRotY;

    string _DefaultAnim;
    string _SelectAnim;

    public int _SelectRole;

	// Use this for initialization
	void Start () {
        _SelectRole = -1;
        Stage.inst.onTouchBegin.Add(OnTouchBegin);

        string[] posStr = Define.GetStr("CreateMalePos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _MalePos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));
        _MaleRotY = Define.GetFloat("CreateMaleRotY");

        posStr = Define.GetStr("CreateFemalePos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _FemalePos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));
        _FemaleRotY = Define.GetFloat("CreateFemaleRot");

        posStr = Define.GetStr("CreateSelectPos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _ToPos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));

        _DefaultAnim = Define.GetStr("CreateDefaultClip");
        _SelectAnim = Define.GetStr("CreateSelectClip");

        _Actors = new Actor[2];
        EntityData eData = EntityData.GetData(Define.MALE_ID);
        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        GameObject actorObj = AssetLoader.LoadAsset(dData._AssetPathDetail);
        actorObj.transform.Rotate(Vector3.up, _MaleRotY);
        _Actors [0] = new Actor(actorObj, _MalePos, 0, "", "", null, dData._Id);
        _Actors [0].Play(_DefaultAnim);
        NpcHandle npcHandler = actorObj.AddComponent<NpcHandle>();
        npcHandler.ID = 0;
        npcHandler._DisableTrigger = true;

        eData = EntityData.GetData(Define.FEMALE_ID);
        dData = DisplayData.GetData(eData._DisplayId);
        actorObj = AssetLoader.LoadAsset(dData._AssetPathDetail);
        actorObj.transform.Rotate(Vector3.up, _FemaleRotY);
        _Actors [1] = new Actor(actorObj, _FemalePos, 0, "", "", null, dData._Id);
        _Actors [1].Play(_DefaultAnim);
        npcHandler = actorObj.AddComponent<NpcHandle>();
        npcHandler.ID = 1;
        npcHandler._DisableTrigger = true;
	}

    public Actor GetRole(int roleid)
    {
        return _Actors[roleid];
    }

    public void SelectMale()
    {
        if (_SelectRole == 0)
            return;

        _Actors [0]._ActorObj.SetActive(true);
        _Actors [1]._ActorObj.SetActive(false);

        Proxy4Lua.FocusSelectRoleObject(0);

        if (_Actors [0] == null)
            return;

        _Actors [0].Play(_SelectAnim);
        _Actors [0].PlayQueue(_DefaultAnim);
        _SelectRole = 0;
        UIManager.SetDirty("xuanren");
    }

    public void SelectFemale()
    {
        if (_SelectRole == 1)
            return;

        _Actors [1]._ActorObj.SetActive(true);
        _Actors [0]._ActorObj.SetActive(false);

        Proxy4Lua.FocusSelectRoleObject(1);

        if (_Actors [1] == null)
            return;

        _Actors [1].Play(_SelectAnim);
        _Actors [1].PlayQueue(_DefaultAnim);
        _SelectRole = 1;
        UIManager.SetDirty("xuanren");
    }

    void OnTouchBegin(EventContext context)
    {
        bool isTouchOnNpc = false;
        RaycastHit hit;
        Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
        if (Physics.Raycast(ray, out hit))
            isTouchOnNpc = hit.transform.CompareTag("Npc");

        if (isTouchOnNpc)
        {
            NpcHandle handler = hit.transform.GetComponent<NpcHandle>();
            if (handler != null)
                handler.Excute();
        }
    }
	
    void OnDestroy()
    {
        _SelectRole = -1;

        EntityData eData = EntityData.GetData(Define.MALE_ID);
        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        AssetLoader.UnloadAsset(dData._AssetPathDetail, true);

        eData = EntityData.GetData(Define.FEMALE_ID);
        dData = DisplayData.GetData(eData._DisplayId);
        AssetLoader.UnloadAsset(dData._AssetPathDetail, true);
    }
}
