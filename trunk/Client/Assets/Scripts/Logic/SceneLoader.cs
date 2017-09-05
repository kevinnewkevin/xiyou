using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class SceneLoader
{
    static AsyncOperation asyncOper;
    static public void LoadScene(string sceneName)
    {
        CameraEffect.Fade(delegate{
            UIManager.DisposeAll();
            //UIManager.HideAll();
            //UIManager.Show("LoadingPanel");
            asyncOper = SceneManager.LoadSceneAsync(sceneName);
            if (!sceneName.Equals(Define.SCENE_BATTLE))
            {
                UIManager.Show("zhujiemian");
            }
            CameraEffect.Continue();
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
