using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class AssetLoader {

	public GameObject LoadAsset(string path)
    {
#if UNITY_EDITOR
        Object obj = Resources.Load(path);
        if (obj == null)
            return null;
        return GameObject.Instantiate(obj) as GameObject;
#else
        return null;
#endif
    }
}
