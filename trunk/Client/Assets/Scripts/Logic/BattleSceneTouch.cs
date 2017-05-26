using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BattleSceneTouch : MonoBehaviour {

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
        if(Input.GetMouseButtonUp(0))
        {
            Ray ray = Camera.main.ScreenPointToRay(Input.mousePosition);
            RaycastHit hit;
            if (Physics.Raycast(ray, out hit, 10000f))
            {
                if (hit.transform.CompareTag("Point"))
                {
                    PointHandle handler = hit.transform.GetComponent<PointHandle>();
                    if (handler != null)
                        handler.Excute();
                }
            }
        }
	}
}
