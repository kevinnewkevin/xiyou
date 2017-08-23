using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PointHandle : MonoBehaviour {

    GameObject _Light;
    int _Pos;
    public void Init(int pos)
    {
        _Pos = pos;
        string pl = Define.GetStr("PointLight");
        _Light = AssetLoader.LoadAsset(pl);
        _Light.transform.parent = transform;
        _Light.transform.localPosition = Vector3.zero;
        _Light.transform.localScale = Vector3.one;
    }

    public void Excute()
    {
        Battle.OperateSetActor(_Pos);
    }
}
