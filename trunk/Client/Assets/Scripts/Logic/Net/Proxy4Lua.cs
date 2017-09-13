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

	static public void StopMatching()
	{
		NetWoking.S.StopMatching();
	}

    static public void CreatePlayer(int templateId, string nickName)  
    {
        NetWoking.S.CreatePlayer(templateId, nickName);
    }

	static public void DeleteItem(long instId,int stack)
	{
		NetWoking.S.DeleteItem(instId,stack);
	}
    #endregion

    #region 内部接口

    static public bool ReconnectServer()
    {
        return NetWoking.ReConnect();
    }

    static public void SelectCard4Ready(int idx)
    {
        if (idx >= Battle._HandCards.Count || idx < 0)
            return;
        
        Battle._SelectedHandCardInstID = Battle._HandCards[idx].InstId;
        Battle.SwitchPoint(true);
    }

    static public FairyGUI.GoWrapper GetAssetGameObject(string assetPath)
    {
        if (string.IsNullOrEmpty(assetPath))
            return new FairyGUI.GoWrapper(new GameObject());
        GameObject go = AssetLoader.LoadAsset(assetPath);
        go.transform.position = new Vector3(0f, 0f, 1000f);
        go.transform.localScale = Vector3.one * Define.GetFloat("UIModelScale");
        go.transform.Rotate(Vector3.up, 180f);
        return new FairyGUI.GoWrapper(go);
    }

    static public FairyGUI.GoWrapper GetAssetGameObject(string assetPath, float scale, float height)
    {
        if (string.IsNullOrEmpty(assetPath))
            return new FairyGUI.GoWrapper(new GameObject());
        GameObject go = AssetLoader.LoadAsset(assetPath);
        go.transform.position = new Vector3(0f, height, 1000f);
        go.transform.localScale = Vector3.one * scale;
        go.transform.Rotate(Vector3.up, 180f);
        return new FairyGUI.GoWrapper(go);
    }

    static public void UnloadAsset(string assetPath)
    {
        if (string.IsNullOrEmpty(assetPath))
            return;

        AssetLoader.UnloadAsset(assetPath);
    }

    static public int[] GetTalk()
    {
        OpraSystem os = GameObject.FindObjectOfType<OpraSystem>();
        if (os == null)
            return null;

        if (os._CrtScreen == null)
            return null;

        if (os._CrtScreen._Data == null)
            return null;
        
        return os._CrtScreen._Data._Talks;
    }

    static public int GetTalkNum()
    {
        OpraSystem os = GameObject.FindObjectOfType<OpraSystem>();
        if (os == null)
            return 0;

        if (os._CrtScreen == null)
            return 0;

        if (os._CrtScreen._Data == null)
            return 0;

        if (os._CrtScreen._Data._Talks == null)
            return 0;

        return os._CrtScreen._Data._Talks.Length;
    }

    static public void ContinueOpra()
    {
        OpraSystem os = GameObject.FindObjectOfType<OpraSystem>();
        if (os == null)
            return;

        os.Play();
    }

    static public void FocusNpcObject(int npcid)
    {
        Actor npc = World.GetNpc(npcid);
        if (npc != null)
        {
            CameraTracker ct = Camera.main.GetComponent<CameraTracker>();
            if (ct != null)
                ct.Focus(npc._ActorObj);
        }
    }

    static public bool IsAchieve1
    {
        get
        {
            if (Battle._Result == null)
                return false;

            CheckpointData cpData = CheckpointData.GetDataByBattleID(Battle._BattleId);
            if (cpData == null)
                return false;

            if(Battle._Result.KillMonsters == null)
                return false;

            for(int i=0; i < Battle._Result.KillMonsters.Length; ++i)
            {
                if (Battle._Result.KillMonsters [i] == cpData._Star1Need)
                    return true;
            }
            return false;
        }
    }

    static public bool IsAchieve2
    {
        get
        {
            if (Battle._Result == null)
                return false;

            CheckpointData cpData = CheckpointData.GetDataByBattleID(Battle._BattleId);
            if (cpData == null)
                return false;

            if(Battle._Result.BattleRound <= cpData._Star2Need)
                return true;

            return false;
        }
    }

    static public bool IsAchieve3
    {
        get
        {
            if (Battle._Result == null)
                return false;

            CheckpointData cpData = CheckpointData.GetDataByBattleID(Battle._BattleId);
            if (cpData == null)
                return false;

            if(Battle._Result.MySelfDeathNum <= cpData._Star3Need)
                return true;
            return false;
        }
    }

    #endregion
}
