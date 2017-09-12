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
//            UIManager.DisposeAll();
            UIManager.HideAll();
            //UIManager.Show("LoadingPanel");
            asyncOper = SceneManager.LoadSceneAsync(sceneName);
            if(Battle.InBattle && Battle.CurrentState != Battle.BattleState.BS_Init)
                Battle.Fini();

            if(Battle.CurrentState == Battle.BattleState.BS_Init)
            {
                if(Battle._BattleId != 0)
                {
                    BattleData bData = BattleData.GetData(Battle._BattleId);
                    EntityData eData = null;
                    DisplayData dData = null;
                    SkillData sData = null;
                    for(int i=0; i < bData._Monsters.Length; ++i)
                    {
                        eData = EntityData.GetData(bData._Monsters[i]);
                        if(eData != null)
                        {
//                            dData = DisplayData.GetData(eData._DisplayId);
//                            if(dData != null)
//                                AssetLoader.LaunchBundle(dData._AssetPath);

                            for(int j=0; j < eData._Skills.Length; ++j)
                            {
                                sData = SkillData.GetData(eData._Skills[j]);
                                if(sData != null)
                                {
                                    AssetLoader.LaunchBundle(sData._CastEffect);
                                    AssetLoader.LaunchBundle(sData._BeattackEffect);
                                    AssetLoader.LaunchBundle(sData._SkillEffect);
                                }
                            }
                        }
                    }
                }
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
