using UnityEngine;
using System.Collections;
using FairyGUI;

public class Billboard : MonoBehaviour {

    UIPanel panel;
	// Use this for initialization
	void Start () {
        //panel = GetComponent<UIPanel>();
	}
	
	// Update is called once per frame
	void Update () {
//        transform.LookAt(Camera.main.transform);
//        transform.Rotate(Vector3.up, 180f);
        transform.LookAt(Camera.main.transform.position, -Vector3.up);
        transform.Rotate(Vector3.left, 180f);
	}
}
