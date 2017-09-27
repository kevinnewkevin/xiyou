using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CreateRoleSceneTouch : MonoBehaviour {

    Actor[] _Actors;

    Vector3 _MalePos, _FemalePos, _ToPos;
    string _DefaultAnim;
    string _SelectAnim;

	// Use this for initialization
	void Start () {
        UIManager.Show("xuanren");

        string[] posStr = Define.GetStr("CreateMalePos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _MalePos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));

        posStr = Define.GetStr("CreateFemalePos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _FemalePos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));

        posStr = Define.GetStr("CreateSelectPos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _ToPos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));

        _DefaultAnim = Define.GetStr("CreateDefaultClip");
        _SelectAnim = Define.GetStr("CreateSelectClip");

        _Actors = new Actor[2];
        EntityData eData = EntityData.GetData(Define.MALE_ID);
        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        _Actors [0] = new Actor(AssetLoader.LoadAsset(dData._AssetPath), _MalePos, 0, "", "", null, dData._Id);
        _Actors [0].Play(_DefaultAnim);

        eData = EntityData.GetData(Define.FEMALE_ID);
        dData = DisplayData.GetData(eData._DisplayId);
        _Actors [1] = new Actor(AssetLoader.LoadAsset(dData._AssetPath), _FemalePos, 0, "", "", null, dData._Id);
        _Actors [1].Play(_DefaultAnim);
	}

    public void SelectMale()
    {
        if (_Actors [0] == null)
            return;

        _Actors [0].Play(_SelectAnim);
        _Actors [0].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
        _Actors [0].MoveTo(_ToPos, null);

        if (_Actors [1] == null)
            return;
        
        _Actors [1].Play(_SelectAnim);
        _Actors [1].PlayQueue(_DefaultAnim);
        _Actors [1].MoveTo(_FemalePos, null);
    }

    public void SelectFemale()
    {
        if (_Actors [1] == null)
            return;

        _Actors [1].Play(_SelectAnim);
        _Actors [1].PlayQueue(Define.ANIMATION_PLAYER_ACTION_IDLE);
        _Actors [1].MoveTo(_ToPos, null);

        if (_Actors [0] == null)
            return;

        _Actors [0].Play(_SelectAnim);
        _Actors [0].PlayQueue(_DefaultAnim);
        _Actors [0].MoveTo(_MalePos, null);
    }
	
	// Update is called once per frame
	void Update () {
		
	}
}
