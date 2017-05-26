using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PointHandle : MonoBehaviour {

    int _Pos;
    public void Init(int pos)
    {
        _Pos = pos;
    }

    public void Excute()
    {
        Battle.OperateSetActor(_Pos);
    }
}
