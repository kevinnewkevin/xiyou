using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Chaos : MonoBehaviour {

	// Use this for initialization
	void Start () {
        AppEntry ae = GameObject.FindObjectOfType<AppEntry>();
        GameObject go;
        if (ae == null)
        {
            go = new GameObject("AppEntry");
            go.AddComponent<AppEntry>();
            go.AddComponent<DisplayFPS>();
        }

        OpraSystem os = GameObject.FindObjectOfType<OpraSystem>();
        if (os == null)
        {
            go = new GameObject("OpraSystem");
            go.AddComponent<OpraSystem>();
            go.AddComponent<Animation>();
        }

        if (Camera.main == null)
        {
            go = new GameObject("Camera");
            go.AddComponent<Camera>();
            go.AddComponent<AudioListener>();
            go.tag = "MainCamera";
        }
	}
	
	// Update is called once per frame
	void Update () {
		
	}
}
