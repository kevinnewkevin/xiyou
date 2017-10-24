using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class SceneLoader
{
    static AsyncOperation asyncOper;
    static string _PreScene = "";
    static public void LoadScene(string sceneName)
    {
        CameraEffect.Fade(delegate{
            #if !EDITOR_MODE
            if(!string.IsNullOrEmpty(_PreScene) || (!string.IsNullOrEmpty(_PreScene) && !_PreScene.Equals(sceneName)))
            AssetLoader.UnloadAsset("Scene/" + _PreScene);
            #endif

            UIManager.HideAll();
            #if !EDITOR_MODE
                if(string.IsNullOrEmpty(_PreScene) || !_PreScene.Equals(sceneName))
                    AssetLoader.LoadAsset("Scene/" + sceneName);
                _PreScene = sceneName;
            #endif
            asyncOper = SceneManager.LoadSceneAsync(sceneName);
            if(Battle.InBattle && Battle.CurrentState != Battle.BattleState.BS_Init)
                Battle.Fini();

            if(Battle.CurrentState == Battle.BattleState.BS_Init)
            {
                Battle.LaunchBundle();
            }
        }, delegate {
            if(Battle.CurrentState == Battle.BattleState.BS_Init)
            {
                Battle.FadedCallback();
            }
        });
    }

    static public void Update()
    {
        if (asyncOper != null)
        {
            if (asyncOper.isDone)
            {
                asyncOper = null;
            }
        }
    }

    static public float Progress
    {
        get
        {
            if (asyncOper != null)
            {
                return asyncOper.progress;
            }
            return 0f;
        }
    }
}
