﻿using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class SceneLoader
{
    static AsyncOperation asyncOper;
    static public void LoadScene(string sceneName)
    {
        CameraEffect.Fade(delegate{
            UIManager.HideAll();
            //UIManager.HideAll();
            //UIManager.Show("LoadingPanel");
            asyncOper = SceneManager.LoadSceneAsync(sceneName);
            CameraEffect.Continue();
            if(Battle.InBattle && Battle.CurrentState != Battle.BattleState.BS_Init)
                Battle.Fini();
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
