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
	}

	static public void StartMatching(int teamId)
	{
		NetWoking.S.StartMatching(teamId);
	}

	static public void StopMatching()
	{
		NetWoking.S.StopMatching();
	}

    static public void RequestChapterStarReward(int chapterId, int star)
    {
        NetWoking.S.RequestChapterStarReward(chapterId, star);
    }

	static public void ResolveItem(long instId,int num)
	{
		NetWoking.S.ResolveItem(instId, num);
	}

    static public void  PromoteUnit(long instId)
    {
        NetWoking.S.PromoteUnit(instId);
    }

    static public void Login(string account, string password)  
    {
        //if has channel get username from channel;
        //if debug fetch mobile phone info instead.

        COM_LoginInfo info = new COM_LoginInfo();
        info.Username = account;
        info.Password = password;
        NetWoking.S.Login(info);
    }

    static public void CreatePlayer(int templateId, string nickName)  
    {
        NetWoking.S.CreatePlayer(templateId, nickName);
    }

	static public void DeleteItem(long instId,int stack)
	{
		NetWoking.S.DeleteItem(instId,stack);
	}

    static public void UpdateSkill(int idx, int skillid)
    {
        NetWoking.S.SkillUpdate(idx, skillid);
    }

	static public void BuyShopItem(int shopId)
	{
		NetWoking.S.BuyShopItem (shopId);	
	}

    static public void EquipSkill(int idx, int skillid)
    {
        COM_LearnSkill ls = new COM_LearnSkill();
        ls.SkillID = skillid;
        ls.Position = idx;
        NetWoking.S.EquipSkill(ls);
    }
	
	static public void RefreshBlackMarkte()
	{
		NetWoking.S.RefreshBlackMarkte ();
	}



    #endregion

    #region 内部接口

    static public bool _ReadyToJoinBattle = false;
    static public bool ReadyToJoinBattle
    {
        get
        {
            bool tmp = _ReadyToJoinBattle;
            _ReadyToJoinBattle = false;
            return tmp;
        }
    }

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

    static public bool SameCardSelected(int idx)
    {
        if (idx >= Battle._HandCards.Count || idx < 0)
            return false;
        
        return Battle._SelectedHandCardInstID == Battle._HandCards [idx].InstId;
    }

    static public FairyGUI.GoWrapper GetAssetGameObject(string assetPath, bool canRot = false)
    {
        if (string.IsNullOrEmpty(assetPath))
            return new FairyGUI.GoWrapper(new GameObject());
        GameObject go = AssetLoader.LoadAsset(assetPath);
        if(canRot)
            go.AddComponent<ActorRotate>();
        go.transform.position = new Vector3(0f, 0f, 1500f);
        go.transform.localScale = Vector3.one * Define.GetFloat("UIModelScale");
        go.transform.Rotate(Vector3.up, 180f);
        go.SetActive(false);
		go.SetActive(true);
        return new FairyGUI.GoWrapper(go);
    }

    static public FairyGUI.GoWrapper GetAssetGameObject(string assetPath, float scale, float height, bool canRot = false)
    {
        if (string.IsNullOrEmpty(assetPath))
            return new FairyGUI.GoWrapper(new GameObject());
        GameObject go = AssetLoader.LoadAsset(assetPath);
        if(canRot)
            go.AddComponent<ActorRotate>();
        go.transform.position = new Vector3(0f, height, 1500f);
        go.transform.localScale = Vector3.one * scale;
        go.transform.Rotate(Vector3.up, 180f);
        go.SetActive(false);
        go.SetActive(true);
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

    static public void FocusSelectRoleObject(int roleid)
    {
        CreateRoleSceneTouch crst = GameObject.FindObjectOfType<CreateRoleSceneTouch>();
        if (crst == null)
            return;
        
        Actor role = crst.GetRole(roleid);
        if (role != null)
        {
            CameraTracker ct = Camera.main.GetComponent<CameraTracker>();
            if (ct == null)
                ct = Camera.main.gameObject.AddComponent<CameraTracker>();
            ct.Focus(role._ActorObj);
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

    static public void ShakeCamera(Vector3 amount, float time)
    {
        Camera[] cameras = GameObject.FindObjectsOfType<Camera>();
        for(int i=0; i < cameras.Length; ++i)
        {
            iTween.ShakePosition(cameras[i].gameObject, amount, time);
        }
    }

    static public SkillData GetCardInstSkillData(long instId, int idx)
    {
        COM_Unit unit = GamePlayer.GetCardByInstID(instId);
        if (unit == null)
            return null;

        for(int i=0; i < unit.Skills.Length; ++i)
        {
            if (unit.Skills [i].Pos == idx)
                return SkillData.GetData(unit.Skills [i].SkillId);
        }
        return null;
    }

    static public SkillData GetPlayerSkillData(int idx)
    {
        if (GamePlayer._Data.Skills == null)
            return null;
        
        for(int i=0; i < GamePlayer._Data.Skills.Length; ++i)
        {
            if (GamePlayer._Data.Skills [i].Pos == idx)
                return SkillData.GetData(GamePlayer._Data.Skills [i].SkillId);
        }
        return null;
    }

    static public int GetIndexBySkillID(int skillId)
    {
        if (GamePlayer._Data.Skills == null)
            return -1;
        
        for(int i=0; i < GamePlayer._Data.Skills.Length; ++i)
        {
            if (skillId == GamePlayer._Data.Skills [i].SkillId)
                return GamePlayer._Data.Skills [i].Pos;
        }
        return -1;
    }

    static public int CrtSelect()
    {
        CreateRoleSceneTouch crst = GameObject.FindObjectOfType<CreateRoleSceneTouch>();
        if (crst == null)
            return -1;

        return crst._SelectRole;
    }

    static public void SelectRole(int idx)
    {
        CreateRoleSceneTouch crst = GameObject.FindObjectOfType<CreateRoleSceneTouch>();
        if (crst == null)
            return;

        if (idx == 0)
            crst.SelectMale();
        else
            crst.SelectFemale();
    }

    static public string GetRoleDesc(int idx)
    {
        EntityData eData = EntityData.GetData(idx+1);
        if (eData != null)
            return eData._Desc;
        return "";
    }

    static public float DeltaTime
    {
        get
        {
            return Time.deltaTime;
        }
    }

    static public System.Collections.Generic.Queue<string> Message = new Queue<string>();
    static public void PopMsg(string msg)
    {
        Message.Enqueue(msg);
    }

    static float _JoinBattleDelay = 0f;
    static public float NextBattleDelay
    {
        set
        {
            _JoinBattleDelay = value;
        }
        get
        {
            float tmp = _JoinBattleDelay;
            _JoinBattleDelay = 0f;
            return tmp;
        }
    }

    #endregion
}
