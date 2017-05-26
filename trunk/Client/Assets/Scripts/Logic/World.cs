using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class World {

    const int MAIN_SCENE_ID = 1;

    static Actor _PlayerActor;

    static public void InitPlayerActor(/* player struct */)
    {
        // load player Asset for gameobject
        GameObject go = AssetLoader.LoadAsset(PathDefine.PLAYER_ASSET_PATH + "longtaizi");
        _PlayerActor = new Actor(go, new Vector3(20f, -14.3f, 4f), GamePlayer._InstID);
        Camera.main.GetComponent<CameraTracker>()._FollowTarget = go;
    }

    static public void InitNpcActor()
    {
        SceneData scene = SceneData.GetData(MAIN_SCENE_ID);
        NpcData npc = null;
        for(int i=0; i < scene._NpcId.Length; ++i)
        {
            npc = NpcData.GetData(scene._NpcId[i]);
            string assetPath = DisplayData.GetData(npc._Display)._AssetPath;
            // load player Asset for gameobject
            GameObject go = AssetLoader.LoadAsset(assetPath);
            NpcHandle npcHandler = go.AddComponent<NpcHandle>();
            npcHandler.ID = npc._Id;
            go.transform.position = npc._Position;
        }
    }

    static public Actor PlayerActor
    {
        get
        {
            return _PlayerActor;
        }
    }
}
