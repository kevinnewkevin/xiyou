using System.Collections;
using System.Collections.Generic;
using UnityEngine;

[RequireComponent(typeof(BoxCollider))]
public class NpcHandle : MonoBehaviour {

    int _ID;
    public bool _DisableTrigger;
    public int ID
    {
        set{ _ID = value; gameObject.tag = "Npc"; }
        get{ return _ID; }
    }

    void Start()
    {
        BoxCollider bc = GetComponent<BoxCollider>();
        if (bc != null)
        {
            bc.center = new Vector3(0f, 0.75f, 0f);
        }
        bc.enabled = !_DisableTrigger;
    }

    public void Excute()
    {
        LuaManager.CallGlobal("ExcuteNpc", ID);
    }
}
