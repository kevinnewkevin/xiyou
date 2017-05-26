using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class SceneLoader
{
    static public void LoadScene(string sceneName, bool keepUI = false)
    {
        SceneManager.LoadScene(sceneName, LoadSceneMode.Single);
        if (!keepUI)
        {
            UIManager.HideAll();
        }
        if(sceneName.Equals(Define.SCENE_BATTLE))
        {
            UIManager.Show("BattlePanel");
        }
    }
}
