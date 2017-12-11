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
	
	static public bool _CancelMatch = false;
	static public void StartMatching(int teamId) 
	{
		new Timer ().Start (3f, delegate {
			if(!_CancelMatch)
				NetWoking.S.StartMatching(teamId);
			_CancelMatch = false;
		});
	}

	static public void StopMatching()
	{
		_CancelMatch = true;
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

    static public COM_LoginInfo _LoginInfo;
    static public void Login(string account, string password)  
    {
        //if has channel get username from channel;
        //if debug fetch mobile phone info instead.

        _LoginInfo = new COM_LoginInfo();
        _LoginInfo.Username = account;
        _LoginInfo.Password = password;
        NetWoking.S.Login(_LoginInfo);
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

    static public void SendChat(COM_Chat chat)
    {
        NetWoking.S.SendChat (chat);
    }

    static public void PlayAudio(long id)
    {
        NetWoking.S.RequestAudio(id);
    }
	
	static public void SerchFriendByName(string name)
	{
		NetWoking.S.SerchFriendByName (name);	
	}
	static public void SerchFriendRandom()
	{
		NetWoking.S.SerchFriendRandom ();	
	}
	static public void ProcessingFriend(string name)
	{
		NetWoking.S.ProcessingFriend (name);	
	}
	static public void DeleteFriend(int instid)
	{
		NetWoking.S.DeleteFriend (instid);	
	}
	static public void ApplicationFriend(string name)
	{
		NetWoking.S.ApplicationFriend (name);	
	}
	static public void AddEnemy(long instId)
	{
		NetWoking.S.AddEnemy (instId);	
	}

	static public void DeleteEnemy(long instId)
	{
		NetWoking.S.DeleteEnemy (instId);	
	}
	static public void QueryPlayerInfo(long instId)
	{
		NetWoking.S.QueryPlayerInfo (instId);	
	}

    static public void NeedAssistantItem(int itemId)
    {
        NetWoking.S.NeedAssistantItem(itemId);
        GamePlayer.AddCoolDown("AssistantCoolDown", TimerManager.GetTimeStamp());
    }

    static public void Assistant(int assId)
    {
        NetWoking.S.AssistantItem(assId);
    }
	
	static public void CreateGuild(string name)
	{
		NetWoking.S.CreateGuild(name);
	}
	static public void RequestJoinGuild(int guid)
	{
		NetWoking.S.RequestJoinGuild(guid);
	}
	static public void KickOut(long guid)
	{
		NetWoking.S.KickOut(guid);
	}
    static public void AcceptRequestGuild(long playerId) 							
	{
		NetWoking.S.AcceptRequestGuild(playerId);
	}
    static public void RefuseRequestGuild(long playerId)								
	{
		NetWoking.S.RefuseRequestGuild(playerId);
	}
	static public void ChangeMemberPosition(long targetId , int job)			
	{
		NetWoking.S.ChangeMemberPosition(targetId,job);
	}
	static public void QueryGuildList()												
	{
		NetWoking.S.QueryGuildList();
	}
	static public void QueryGuildDetails(int guildid)								
	{
		NetWoking.S.QueryGuildDetails(guildid);
	}
	static public void QueryGuildData()
	{
		NetWoking.S.QueryGuildData();
	}
	static public void ChangeJoinGuildFlag(bool isFlag,int require)
	{
		NetWoking.S.ChangeJoinGuildFlag(isFlag,require);
	}
    static public void LeaveGuild()
    {
        NetWoking.S.LeaveGuild();
    }
    static public void RandChapter()
    {
        NetWoking.S.RandChapter();
    }
	static public void GetMailItem(int id)
	{
		NetWoking.S.GetMailItem(id);
	}
	static public void ReadMail(int id)
	{
		NetWoking.S.ReadMail(id);
	}
		static public void DelMail(int id)
		{
				NetWoking.S.DelMail(id);
		}
    static public void RequestSelfRecordData()
    {
        BattleRecordSystem.MirrorPlayerId = GamePlayer._InstID;
        UIManager.Show("luxiang");
        BattleRecordSystem.CacheSimpleData(GamePlayer.MyRecord);
    }

    static public void RequestPlayerRecordData(long instid)
    {
        BattleRecordSystem.MirrorPlayerId = instid;
        UIManager.Show("luxiang");
        NetWoking.S.QueryPlayerRecordDetail(instid);
    }

	static public void RequestFBRecordData(int smallChapterId)
    {
        UIManager.Show("guankaluxiang");
		NetWoking.S.QueryCheckpointRecordDetail(smallChapterId);
    }

    static public void RequestRecord(long recordid)
    {
        BattleRecordSystem.RequestRecord(recordid);
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

    static public void CheckBattleTurn()
    {
        NetWoking.S.QueryBattleRound(Battle._Turn);
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

    static public FairyGUI.GoWrapper GetEffectAssetGameObject(string assetPath)
    {
        if (string.IsNullOrEmpty(assetPath))
            return new FairyGUI.GoWrapper(new GameObject());
        GameObject go = AssetLoader.LoadAsset(assetPath);
        go.transform.position = new Vector3(0f, 0f, 1500f);
        go.transform.localScale = Vector3.one;
        go.SetActive(false);
        go.SetActive(true);
        return new FairyGUI.GoWrapper(go);
    }

    static public FairyGUI.GoWrapper GetAssetGameObject(string assetPath, bool canRot, float distance = 1500f, float scale = 1f)
    {
        if (string.IsNullOrEmpty(assetPath))
            return new FairyGUI.GoWrapper(new GameObject());
        GameObject go = AssetLoader.LoadAsset(assetPath);
        if(canRot)
            go.AddComponent<ActorRotate>();
        go.transform.position = new Vector3(0f, 0f, distance);
        go.transform.localScale = Vector3.one * Define.GetFloat("UIModelScale") * scale;
        iTween.ShakePosition(go, new Vector3(0.01f, 0f, 0f), 0.02f);
        ParticalScale ps = go.GetComponent<ParticalScale>();
        if(ps != null)
            ps.scanleSize = Define.GetFloat("UIModelScale") * scale * FairyGUI.Stage.inst.cachedTransform.localScale.x;
        go.transform.Rotate(Vector3.up, 180f);
        go.SetActive(false);
        go.SetActive(true);
        return new FairyGUI.GoWrapper(go);
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
        iTween.ShakePosition(go, new Vector3(0.01f, 0f, 0f), 0.02f);
        ParticalScale ps = go.GetComponent<ParticalScale>();
        if(ps != null)
            ps.scanleSize = Define.GetFloat("UIModelScale") * FairyGUI.Stage.inst.cachedTransform.localScale.x;
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

    static public void BlackGameObject(FairyGUI.GGraph holder)
    {
        if (holder.displayObject == null)
            return;

        if (holder.displayObject.gameObject == null)
            return;

        SkinnedMeshRenderer[] smr = holder.displayObject.gameObject.GetComponentsInChildren<SkinnedMeshRenderer>();
        for(int i=0; i < smr.Length; ++i)
        {
            if (smr [i].material != null)
                smr [i].material.SetColor("_Color", Color.black);
        }
    }

	static public void ColorGameObject(FairyGUI.GGraph holder,float R,float G,float B)
	{
		if (holder.displayObject == null)
			return;

		if (holder.displayObject.gameObject == null)
			return;

		SkinnedMeshRenderer[] smr = holder.displayObject.gameObject.GetComponentsInChildren<SkinnedMeshRenderer>();
		for(int i=0; i < smr.Length; ++i)
		{
			if (smr [i].sharedMaterial != null)
				smr [i].sharedMaterial.SetColor("_Color", new Color(R,G,B));
		}
	}

    static public void WhiteGameObject(FairyGUI.GGraph holder)
    {
        if (holder.displayObject == null)
            return;

        if (holder.displayObject.gameObject == null)
            return;

        SkinnedMeshRenderer[] smr = holder.displayObject.gameObject.GetComponentsInChildren<SkinnedMeshRenderer>();
        for(int i=0; i < smr.Length; ++i)
        {
            if (smr [i].sharedMaterial != null)
                smr [i].sharedMaterial.SetColor("_Color", Color.white);
        }
    }

    static public void UnloadAsset(string assetPath)
    {
        if (string.IsNullOrEmpty(assetPath))
            return;

        AssetLoader.UnloadAsset(assetPath);
    }

    static public void ForceUnloadAsset(string assetPath)
    {
        if (string.IsNullOrEmpty(assetPath))
            return;

        AssetLoader.UnloadAsset(assetPath, true);
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
            if (Battle._IsRecord)
                return false;
            
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
            if (Battle._IsRecord)
                return false;
            
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
            if (Battle._IsRecord)
                return false;
            
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

    static public string RemoveString(string origin, string toRemove)
    {
        return origin.Remove(origin.IndexOf(toRemove));
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
        for (int i = 0; i < UIManager._NoPopIgnoreUI.Count; ++i)
        {
            if (UIManager.IsShow(UIManager._NoPopIgnoreUI [i]))
                return;
        }
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

    static public ResourceUpdate ResUpdate
    {
        get
        {
            ResourceUpdate ru = GameObject.FindObjectOfType<ResourceUpdate>();
            if (ru == null)
            {
                GameObject go = new GameObject();
                ru = go.AddComponent<ResourceUpdate>();
            }
            return ru;
        }
    }

    static public void ReturnToLogin()
    {
        SceneLoader.LoadScene(Define.SCENE_LOGIN);
    }

    static public string _ServerIP;
    static public List<ServInfo> _ServList;
    static public void CheckServer()
    {
        _ServList = new List<ServInfo>();
        string servers = Define.GetStr("DebugServerAddress");
        string[] servs = servers.Split(new char[]{'|'}, System.StringSplitOptions.RemoveEmptyEntries);
        ServInfo si = null;
        for(int i=0; i < servs.Length; ++i)
        {
            string[] servsDetail = servs [i].Split(new char[]{'#'}, System.StringSplitOptions.RemoveEmptyEntries);
            si = new ServInfo();
            si.serverName = servsDetail[0];
            si.serverIP = servsDetail[1];
            _ServList.Add(si);
        }
        UIManager.SetDirty("denglu");
    }

    #endregion

    static public Dictionary<string, List<string>> _AssetToDelete = new Dictionary<string, List<string>>();
    static public void AddToDelete(string ui, string url)
    {
        if (!_AssetToDelete.ContainsKey(ui))
            _AssetToDelete.Add(ui, new List<string>());

        if(!_AssetToDelete [ui].Contains(url))
            _AssetToDelete [ui].Add(url);
    }

    static public void ClearToDeleteAsset(string ui)
    {
        if (!_AssetToDelete.ContainsKey(ui))
            return;
        
        for(int i=0; i < _AssetToDelete[ui].Count; ++i)
        {
            AssetLoader.UnloadAsset(_AssetToDelete[ui][i], false);
        }
        _AssetToDelete [ui].Clear();
        _AssetToDelete.Remove(ui);
    }

    static public Dictionary<string, List<string>> _SceneUI = new Dictionary<string, List<string>>();
    static public void RegHoldUI(string sceneName, string ui)
    {
        if (!_SceneUI.ContainsKey(sceneName))
            _SceneUI.Add(sceneName, new List<string>());

		if(!_SceneUI [sceneName].Contains(ui))
			_SceneUI [sceneName].Add(ui);
    }

    static public void ClearHoldUI(string sceneName)
    {
        if (!_SceneUI.ContainsKey(sceneName))
            return;
        
        _SceneUI.Remove(sceneName);
    }

    static public string ConvertToChineseNumber(int albNum)
    {
        return Define.ConvertInt(albNum.ToString());
    }

    static public string ChangeColor(string content, string color)
    {
        return LuaManager.CallGlobal("ChangeColor", content, color)[0].ToString();
    }

    static public bool LongIsEqual(long val, long val2)
    {
        return val == val2;
    }

    static public FairyGUI.SwipeGesture SwipeGesture(FairyGUI.GObject holder)
    {
        return new FairyGUI.SwipeGesture(holder);
    }

    static public object[] CallGlobalFunction(string funcName, object pa1)
    {
        return LuaManager.CallGlobal(funcName, pa1);
    }

    static public object[] CallGlobalFunction(string funcName, object pa1, object pa2)
    {
        return LuaManager.CallGlobal(funcName, pa1, pa2);
    }

    static public object[] CallGlobalFunction(string funcName, object pa1, object pa2, object pa3)
    {
        return LuaManager.CallGlobal(funcName, pa1, pa2, pa3);
    }

    static public object[] CallGlobalFunction(string funcName, object pa1, object pa2, object pa3, object pa4)
    {
        return LuaManager.CallGlobal(funcName, pa1, pa2, pa3, pa4);
    }
}

public class ServInfo
{
    public string serverName;
    public string serverIP;
}