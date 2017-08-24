using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class World {

    const int MAIN_SCENE_ID = 1;

    static Actor _PlayerActor;

    static Actor[] _NpcActors;

    static public void InitPlayerActor(/* player struct */)
    {
        // load player Asset for gameobject
        DisplayData display = DisplayData.GetData(1);
        GameObject go = AssetLoader.LoadAsset(display._AssetPath);
        _PlayerActor = new Actor(go, new Vector3(20f, -14.3f, 4f), GamePlayer._InstID, GamePlayer._Name, "", null);
        Camera.main.GetComponent<CameraTracker>().MoveToLookAt = go.transform.position.x;
    }

    static public void InitNpcActor()
    {
        SceneData scene = SceneData.GetData(MAIN_SCENE_ID);
        NpcData npc = null;
        _NpcActors = new Actor[scene._NpcId.Length];
        for(int i=0; i < scene._NpcId.Length; ++i)
        {
            npc = NpcData.GetData(scene._NpcId[i]);
            string assetPath = DisplayData.GetData(npc._Display)._AssetPath;
            // load player Asset for gameobject
            GameObject go = AssetLoader.LoadAsset(assetPath);
            Debug.Log(npc._QuestID);
            _NpcActors[i] = new Actor(go, npc._Position, npc._Id, "土地", "", npc._QuestID);
            NpcHandle npcHandler = go.AddComponent<NpcHandle>();
            npcHandler.ID = npc._Id;
        }
    }

    static public Actor PlayerActor
    {
        get
        {
            return _PlayerActor;
        }
    }

    static public void Update()
    {
        if(_PlayerActor != null)
            _PlayerActor.Update();

        if (_NpcActors != null)
        {
            for (int i = 0; i < _NpcActors.Length; ++i)
            {
                if (_NpcActors[i] == null)
                    continue;

                _NpcActors[i].Update();
            }
        }
    }
}
