using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class BattleSceneTouch : MonoBehaviour {

	// Use this for initialization
	void Start () {
        Stage.inst.onTouchBegin.Add(OnTouchBegin);
	}

    void OnTouchBegin()
    {
        if (!Stage.isTouchOnUI)
        {
            RaycastHit hit;
            Ray ray = Camera.main.ScreenPointToRay(new Vector2(Stage.inst.touchPosition.x, Screen.height - Stage.inst.touchPosition.y));
            if (Physics.Raycast(ray, out hit))
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

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnTouchBegin);
    }
}
