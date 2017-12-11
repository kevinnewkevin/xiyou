
class Proxy : ICOM_ServerToClientProxy
{
    public bool ErrorMessage(int err)
    {
        LuaManager.CallGlobal("ErrorMessage", err);
        return true;
    }

    public bool LoginOK(ref COM_AccountInfo info)
    {
        if (info.MyPlayer.InstId == 0)
        {
            SceneLoader.LoadScene(Define.SCENE_CREATE);
        }
        else
        {
            // onboard
            CreatePlayerOK(ref info.MyPlayer);
        }
        return true;
    }


    public bool CreatePlayerOK(ref COM_Player player)
    {
        GamePlayer.Init(player);
        SceneLoader.LoadScene(Define.SCENE_MAIN);
        RankSystem.Sync();
        return true;
    }

    int delaySide = 0;
    int delayBattleId = 0;
    int[] delayOppo = null;
    COM_BattleUnit[] delayUnits = null;
    public bool JoinBattleOk(int side, int battleid, ref int[] opponentCards, ref COM_BattleUnit[] units, ref string battleSceneName)
    {
        if (delayBattleId != 0)
            return true;

        float delayTime = Proxy4Lua.NextBattleDelay;
        if (delayTime == 0f)
        {
            Battle.Init(side, battleid, opponentCards, units);
            BattleData bd = BattleData.GetData(battleid);
            if (bd != null)
                SceneLoader.LoadScene(bd._SceneName);
            else
                SceneLoader.LoadScene(battleSceneName);
        }
        else
        {
            Proxy4Lua._ReadyToJoinBattle = true;
            delaySide = side;
            delayBattleId = battleid;
            delayOppo = opponentCards;
            delayUnits = units;
            new Timer().Start(delayTime, delegate {
                Battle.Init(delaySide, delayBattleId, delayOppo, delayUnits);
                delaySide = 0;
                delayBattleId = 0;
                delayOppo = null;
                delayUnits = null;
                BattleData bd = BattleData.GetData(delayBattleId);
                if (bd != null)
                    SceneLoader.LoadScene(bd._SceneName);
                else
                    SceneLoader.LoadScene(battleSceneName);
            });
        }
        return true;
    }

    public bool JoinBattleOk_back(int turn, int state, int second, ref COM_BattleSnape snap)
    {
        float delayTime = Proxy4Lua.NextBattleDelay;
        Battle.Init(snap.Camp, snap.battleid, snap.targetcards, snap.MainUnit, false, turn, second, state);
        BattleData bd = BattleData.GetData(snap.battleid);
        if (bd != null)
            SceneLoader.LoadScene(bd._SceneName);
        else
            SceneLoader.LoadScene(snap.SceneName);
        return true;
    }

    public bool AddNewUnit(ref COM_Unit u)
	{
		GamePlayer.AddCard (u);
		GamePlayer.showNewCard = true;	
		GamePlayer.newCard = u;
        return true;
    }

    public bool BattleReport(ref COM_BattleReport report)
    {
        Battle.BattleReport = report;
        return true;
    }

    public bool BattleExit(ref COM_BattleResult result)
    {
        Battle.Result = result;
        return true;
    }

    public bool SetupBattleOK()
    {
        return true;
    }

    public bool SetBattleUnitOK(long instid)
    {
        return true;
    }

    public bool BattleSetupSuccess(ref COM_BattleUnit[] cards)
    {
        return true;
    }

    public bool EquipSkillOK(int idx, int skillid)
    {
        GamePlayer.UpdateEquipedSkill(idx, skillid);
        return true;
    }

    public bool SkillUpdateOK(int rsId, int skillid, int idx)
    {
        RoleSkillData.SetData(rsId, skillid);
        if(idx != -1)
            EquipSkillOK(idx, skillid);
        UIManager.Show("jinengshengji");
        UIManager.SetDirty("jineng");
        return true;
    }

	public bool SycnChapterData(ref COM_Chapter chapter)
	{
        JieHunSystem.instance.UpdateChapterData(chapter);
		return true;
	}

	public bool InitBagItems(ref COM_ItemInst[] items)
	{
		BagSystem.Init (items);
		return true;
	}

	public bool AddBagItem(ref COM_ItemInst inst)
	{
		BagSystem.AddItem(inst);
        LuaManager.CallGlobal("GainItem", inst);
		return true;
	}
	public bool UpdateBagItem(ref COM_ItemInst inst)
	{
        int preNum = BagSystem.GetItemMaxNum(inst.ItemId);
		BagSystem.UpdateItem(inst);
        int crtNum = BagSystem.GetItemMaxNum(inst.ItemId);
        if (crtNum > preNum)
        {
            COM_ItemInst tcii = new COM_ItemInst();
            tcii.ItemId = inst.ItemId;
            tcii.Stack = crtNum - preNum;
            LuaManager.CallGlobal("GainItem", tcii);
        }
		return true;
	}
	
	public bool UpdateTiantiVal(int value)
	{
		GamePlayer._TianTiVal = value;
		return true;
	}

	public bool DeleteItemOK(long instId) 
	{
		BagSystem.DelItem (instId);			
		return true;
	}

	public bool OpenChapter(ref COM_Chapter chapter)
	{
        UnityEngine.Debug.Log("has new one");
		JieHunSystem.instance.AddChapterData(chapter);		
		return true;
	}

    public bool PromoteUnitOK()
	{
        UIManager.Show("qianghuachenggong");
        UIManager.SetDirty("qianghuachenggong");
        UIManager.SetDirty("xiangxiziliao");
        UIManager.SetDirty("paiku");
		UIManager.SetDirty("qiecuo");
		return true;
	}

	public bool RequestChapterStarRewardOK()
	{
		UIParamHolder.Set("showChaptersDrop", true);
		JieHunSystem.instance.UpdataChapterRewardData (JieHunSystem.instance.chapterID, JieHunSystem.instance.chapterBox);
		UIManager.SetDirty("xiaoguanka");
		UIManager.Show("baowu");
		return true;
	}

    public bool UpdateUnitIProperty(long instId, int type, int vaule)
    {
        GamePlayer.UpdateUnitIProperty(instId, type, vaule);
        return true;
    }

    public bool UpdateUnitCProperty(long instId, int type, float vaule)
    {
        GamePlayer.UpdateUnitCProperty(instId, type, vaule);
        return true;
    }

	public bool BuyShopItemOK (ref COM_ItemInst[] items)
	{
		ShopSystem.BuyItems = items;
		UIManager.SetDirty("cangbaoge");
	
			if (ShopSystem.buyType == 1000 || ShopSystem.buyType == 1001 || ShopSystem.buyType == 1002)
					UIManager.Show ("kaikabao");
			else 
			{
					Proxy4Lua.PopMsg("购买成功");
			}

		return true;
	}

	public bool	SycnBlackMarkte(ref COM_BlackMarket black)
	{
		ShopSystem.BlackMarket = black;
		UIManager.SetDirty("cangbaoge");

		return true;
	}

    public bool ReceiveChat(ref COM_Chat chat)
    {
        if (FriendSystem.IsInBlack(chat.PlayerInstId))
            return true;
        
		if (chat.Type == 3) 
		{
			//FriendSystem.chatFriend (chat.PlayerInstId, chat);
			FriendSystem.chatFriendStr (chat.PlayerName, chat);
			FriendSystem.AddLatelyFriend (chat.PlayerInstId);
			FriendSystem.AddNewCahtList(chat.PlayerName);
			UIManager.SetDirty ("haoyou");
			UIManager.SetDirty ("zhujiemian");
		} 
		//else 
		//{
			ChatSystem.AddMsg (chat);
		//}
        return true;
    }

    public bool RequestAudioOk(long id, ref byte[] data)
    {
//        YYSystem.PlayRecord(data);
//        ChatSystem.SetRecord(id, data);
        return true;
    }

    public bool RecvTopList(ref COM_TopUnit[] rankList, int myRank)
    {
        UnityEngine.Debug.Log("RecvTopList");
        if(rankList != null)
            RankSystem._AllRank = new System.Collections.Generic.List<COM_TopUnit>(rankList);
        RankSystem._MyAllRank = myRank;
        UIManager.SetDirty("paihangbang");
        UIManager.SetDirty("zhujiemian_paihang");
        return true;
    }

    public bool RecvFriendTopList(ref COM_TopUnit[] rankList, int myRank)
    {
        UnityEngine.Debug.Log("RecvFriendTopList");
        if(rankList != null)
            RankSystem._FirendRank = new System.Collections.Generic.List<COM_TopUnit>(rankList);
        RankSystem._MyFirendRank = myRank;
        UIManager.SetDirty("paihangbang");
        UIManager.SetDirty("zhujiemian_paihang");
        return true;
    }

	public bool	FriendInfo(ref COM_Friend[] friends)
	{
		FriendSystem.randomFriends = friends;
		UIManager.SetDirty("haoyou");

		return true;
	}

	public bool	ApplyFriend(ref COM_Friend friend)
	{
		FriendSystem.ApplyFriend(friend);
		UIManager.SetDirty("haoyou");
		UIManager.SetDirty ("zhujiemian");
		return true;
	}

	public bool	RecvFriend(ref COM_Friend friend)
	{
		FriendSystem.AddFriend (friend);
		UIManager.SetDirty("haoyou");
		LuaManager.CallGlobal("AddFriend",friend.Name);
		return true;
	}

	public bool	DelFriend(long id)
	{
		FriendSystem.DelFriend (id);
		UIManager.SetDirty("haoyou");
		return true;
	}

	public bool	SerchFriendInfo(ref COM_Friend friend)
	{
		FriendSystem.findFriend = friend;
		UIManager.SetDirty("haoyou");
		return true;
	}

	public bool	RecvEnemy(ref COM_Friend friend)
	{
		FriendSystem.AddBlack (friend);
		UIManager.SetDirty("haoyou");
		FriendSystem.DelLatelyFriend (friend.InstId);
		return true;
	}

	public bool	DelEnemy(long id)
	{
		FriendSystem.DelBlack (id);
		UIManager.SetDirty("haoyou");
		return true;
	}

	public bool QueryPlayerInfoOK(ref COM_PlayerInfo info)
	{
		FriendSystem.friendInfo = info; 
		UIManager.Show("wanjiaxinxi");
		return true;
	}

    public bool UpdateGuildAssistant(ref COM_Assistant info, ref string whoAssMe)
    {
        ChatSystem.UpdateAss(info);
        if (!string.IsNullOrEmpty(whoAssMe))
        {
            if (!GamePlayer._Name.Equals(whoAssMe))
            {
                LuaManager.CallGlobal("WhoAssistantMe", whoAssMe, info.ItemId);
            }
        }
        return true;
    }

    public bool SycnGuildAssistant(ref COM_Assistant[] infos)
    {
        if (infos != null)
        {
            for(int i=0; i < infos.Length; ++i)
            {
                ChatSystem.UpdateAss(infos[i]);
            }
        }
        return true;
    }

		public bool CreateGuildOK()
		{
			UIManager.Hide("squadList");
			UIManager.Show("squad");
            LuaManager.CallGlobal("CreateGuild");
			return true;
		}

		public bool DelGuildOK()
		{
			return true;
		}

		public bool LeaveGuildOk(ref string str,bool b)
		{
			if (str == GamePlayer._Name) 
			{
				GamePlayer._iGuildId = 0;
				GuildSystem.myGuild = null;
                LuaManager.CallGlobal("LeaveGuild");
				UIManager.Hide("squad");	
			}

			GuildSystem.LeaveGuildMember (str, b);
			UIManager.SetDirty("squad");
			return true;
		}

		public bool InitGuildData(ref COM_Guild data)
		{
			GuildSystem.myGuild = data;
			GamePlayer._iGuildId = data.GuildId; 
			UIManager.SetDirty("squadList");
			UIManager.SetDirty("squad");
            UIManager.SetDirty("squadSetting");
			return true;
		}

		public bool InitGuildMemberList(ref COM_GuildMember[] data)
		{
			GuildSystem.InitGuildMember (data);
			UIManager.SetDirty("squadList");
				UIManager.SetDirty("squad");
			return true;
		}

		public bool ModifyGuildMemberList(ref COM_GuildMember data,int num)
		{
			GuildSystem.UpdateGuildMember(data);
			if (num == 0) 
			{
				GuildSystem.AddGuildMember(data);
			}
			else if (num == 1) 
			{
				GuildSystem.DelGuildMember(data);
			}

			UIManager.SetDirty("squadList");
			UIManager.SetDirty("squad");
			return true;
		}

		public bool QueryGuildListResult(ref COM_GuildViewerData[] data)
		{
			GuildSystem.InitViewer(data);
			UIManager.SetDirty("squadList");
				UIManager.SetDirty("squad");
			return true;
		}

		public bool QueryGuildDetailsResult(ref COM_GuildDetails data)
		{
			GuildSystem.searchData = data;
			UIManager.SetDirty("squadList");
				UIManager.SetDirty("squad");
			return true;
		}
		public bool JoinGuildOk()
		{
			if (UIManager.IsShow ("squadList")) 
			{
				UIManager.Hide("squadList");	
				UIManager.Show("squad");
			}	
			return true;
		}
		public bool  TianTiSeasonDrop(ref COM_Award item)
		{
			return true;
		}

    public bool QueryBattleRecordOK(ref COM_BattleRecord br)
    {
        BattleRecordSystem.SetRecord(br);
        return true;
    }

    public bool QueryRecordDetailOK(ref COM_BattleRecord_Detail[] rds)
    {
        if(rds == null)
            rds = new COM_BattleRecord_Detail[0];
        BattleRecordSystem.CacheSimpleData(rds);
        return true;
    }

	public bool	AppendMail(ref COM_Mail[] mails)
	{
		MailSystem.AppendMail (mails);
		UIManager.SetDirty("youxiang");
		UIManager.SetDirty("zhujiemian");
		return true;
	}

	public bool DelMailOK(int id)
	{
		MailSystem.DelMail (id);
		UIManager.SetDirty("youxiang");
		UIManager.SetDirty("zhujiemian");
		return true;
	}


	public bool UpdateMailOk(ref COM_Mail mail)
	{
		MailSystem.UpdateMail (mail);
		UIManager.SetDirty("youxiang");
		UIManager.SetDirty("zhujiemian");
		return true;
	}
}