using System.Collections;
using System.Collections.Generic;
using UnityEngine;

[RequireComponent(typeof(BoxCollider))]
public class NpcHandle : MonoBehaviour {

    int _ID;
    public int ID
    {
        set{ _ID = value; gameObject.tag = "Npc"; }
        get{ return _ID; }
    }

    public void Excute()
    {
        Debug.Log("Excute Npc Function " + ID.ToString());
        //TODO Test for wait battle.
        LuaManager.Call("Main.lua", "ExcuteNpc", ID);
        //SceneLoader.LoadScene(Define.SCENE_BATTLE);
    }
}
