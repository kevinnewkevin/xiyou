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
        LuaManager.Call("global.lua", "ExcuteNpc", ID);
    }
}
