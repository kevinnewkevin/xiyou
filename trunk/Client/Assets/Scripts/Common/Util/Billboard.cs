using UnityEngine;
using System.Collections;
using FairyGUI;

public class Billboard : MonoBehaviour {


    Camera _MainCam;
	// Use this for initialization
	void Start () {
        _MainCam = Camera.main;
	}
	
	// Update is called once per frame
	void Update () {
//        transform.LookAt(Camera.main.transform.position, -Vector3.up);
//        transform.Rotate(Vector3.left, 180f);

        transform.rotation = _MainCam.transform.rotation;
	}
}
