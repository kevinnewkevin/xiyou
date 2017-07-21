using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Test : MonoBehaviour {

    public GameObject ltz;
    public GameObject dest;
    Actor actor;
	// Use this for initialization
	void Start () {
        //actor = new Actor(ltz);
        //AssetLoader.LoadAsset(PathDefine.PLAYER_ASSET_PATH + "longtaizi");
	}
	
	// Update is called once per frame
	void Update () {
        //TimerManager.Update();

//        if(Input.GetMouseButtonUp(0))
//        {
//            Ray ray = Camera.main.ScreenPointToRay(Input.mousePosition);
//            RaycastHit hit;
//            if (Physics.Raycast(ray, out hit))  
//            {
//                actor.MoveTo(new Vector3(hit.point.x, -14f, hit.point.z), delegate
//                {
//                    Debug.Log("Reached There.");
//                });
//            }
//        }
	}
    const string PLAYER_AB_PATH = "pc/player/";
    void OnGUI()
    {
        if (Input.GetKeyUp(KeyCode.E))
        {
//            actor.MoveTo(dest.transform.position, delegate
//            {
//                Debug.Log("Reached There.");
//            });


        }
    }
}
