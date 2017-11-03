
class Proxy : ICOM_ServerToClientProxy
{
    public bool ErrorMessage(int err)
    {
        LuaManager.Call("global.lua", "ErrorMessage", err);
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
        return true;
    }

    int delaySide = 0;
    int delayBattleId = 0;
    int[] delayOppo = null;
    COM_BattleUnit[] delayUnits = null;
    public bool JoinBattleOk(int side, int battleid, ref int[] opponentCards, ref COM_BattleUnit[] units)
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
                SceneLoader.LoadScene(Define.RandomBattleScene);
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
                    SceneLoader.LoadScene(Define.RandomBattleScene);
            });
        }
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
		return true;
	}
	public bool UpdateBagItem(ref COM_ItemInst inst)
	{
		BagSystem.UpdateItem(inst);
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
}