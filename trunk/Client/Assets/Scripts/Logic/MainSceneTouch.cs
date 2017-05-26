using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class MainSceneTouch : MonoBehaviour {

	// Use this for initialization
	void Start () {
        NetWoking.S.CreatePlayer(1, "guowengui");
//        World.InitPlayerActor();
//        World.InitNpcActor();
	}
	
	// Update is called once per frame
	void Update () {
		if(Input.GetMouseButtonUp(0))
        {
            Ray ray = Camera.main.ScreenPointToRay(Input.mousePosition);
            RaycastHit hit;
            if (Physics.Raycast(ray, out hit, 10000f))
            {
                if (hit.transform.name.Equals("Cube"))
                {
                    World.PlayerActor.MoveTo(new Vector3(hit.point.x, -14.31f, hit.point.z), delegate {
                        World.PlayerActor.Stop();
                    });
                }

                if (hit.transform.CompareTag("Npc"))
                {
                    NpcHandle handler = hit.transform.GetComponent<NpcHandle>();
                    if (handler != null)
                        handler.Excute();
                }
            }
        }
	}
}
