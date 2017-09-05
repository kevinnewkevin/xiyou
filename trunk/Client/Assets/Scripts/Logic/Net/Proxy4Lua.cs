using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Proxy4Lua {
    #region 网络协议
    static public void BattleSetup()
    {
        Battle.BattleSetup();
    }

    static public void BattleJoin()
    {
        NetWoking.S.JoinBattle();
    }
	
	static public void RequestChapterData(int id)
	{
		NetWoking.S.RequestChapterData(id);
	}

	static public void ChallengeSmallChapter(int id)
	{
		NetWoking.S.ChallengeSmallChapter(id);
		UnityEngine.Debug.Log("ChallengeSmallChapter0000");
	}

	static public void StartMatching(int teamId)
	{
		NetWoking.S.StartMatching(teamId);
	}

    static public void CreatePlayer(int templateId, string nickName)  
    {
        NetWoking.S.CreatePlayer(templateId, nickName);
    }
    #endregion

    #region 内部接口
    static public void SelectCard4Ready(int idx)
    {
        if (idx >= Battle._HandCards.Count || idx < 0)
            return;
        
        Battle._SelectedHandCardInstID = Battle._HandCards[idx].InstId;
        Battle.SwitchPoint(true);
    }

    static public FairyGUI.GoWrapper GetAssetGameObject(string assetPath)
    {
        GameObject go = AssetLoader.LoadAsset(assetPath);
        go.transform.position = new Vector3(0f, 0f, 1000f);
        go.transform.localScale = Vector3.one * Define.GetFloat("UIModelScale");
        go.transform.Rotate(Vector3.up, 180f);
        return new FairyGUI.GoWrapper(go);
    }

    #endregion
}
