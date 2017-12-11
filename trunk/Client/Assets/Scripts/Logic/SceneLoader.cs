using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class SceneLoader
{
    static AsyncOperation asyncOper;
#if !EDITOR_MODE
    static string _PreScene = "";
#endif
    static public void LoadScene(string sceneName)
    {
        CameraEffect.Fade(delegate{
            #if !EDITOR_MODE
                if(!string.IsNullOrEmpty(_PreScene) || (!string.IsNullOrEmpty(_PreScene) && !_PreScene.Equals(sceneName)))
                    AssetLoader.UnloadAsset("Scene/" + _PreScene);

                if(string.IsNullOrEmpty(_PreScene) || !_PreScene.Equals(sceneName))
                    AssetLoader.LoadAssetBundle("Scene/" + sceneName);
                _PreScene = sceneName;
            #endif

            UIManager.HideAll();
            asyncOper = SceneManager.LoadSceneAsync(sceneName);
            if(Battle.InBattle && Battle.CurrentState != Battle.BattleState.BS_Init)
                Battle.Fini();

            if(Battle.CurrentState == Battle.BattleState.BS_Init)
            {
                Battle.LaunchBundle();
            }
            AudioSystem.PlayBackground(SceneData.GetMusicData(sceneName));
        }, delegate {
            if(Battle.CurrentState == Battle.BattleState.BS_Init)
            {
                Battle.FadedCallback();
            }
        });
    }

    static public void LoadSceneSync(string sceneName)
    {
        #if !EDITOR_MODE
        if(!string.IsNullOrEmpty(_PreScene) || (!string.IsNullOrEmpty(_PreScene) && !_PreScene.Equals(sceneName)))
        AssetLoader.UnloadAsset("Scene/" + _PreScene);

        if(string.IsNullOrEmpty(_PreScene) || !_PreScene.Equals(sceneName))
        AssetLoader.LoadAssetBundle("Scene/" + sceneName);
        _PreScene = sceneName;
        #endif

        if(Battle.InBattle && Battle.CurrentState != Battle.BattleState.BS_Init)
            Battle.Fini();

        if(Battle.CurrentState == Battle.BattleState.BS_Init)
        {
            Battle.LaunchBundle();
        }
        AudioSystem.PlayBackground(SceneData.GetMusicData(sceneName));

        SceneManager.LoadScene(sceneName);

        if(Battle.CurrentState == Battle.BattleState.BS_Init)
        {
            Battle.FadedCallback();
        }
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
