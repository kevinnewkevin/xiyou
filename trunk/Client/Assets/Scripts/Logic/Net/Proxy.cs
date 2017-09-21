
class Proxy : ICOM_ServerToClientProxy
{
    public bool ErrorMessage(int err)
    {
        LuaManager.Call("global.lua", "ErrorMessage", err);
        return true;
    }

    public bool LoginOK(ref COM_AccountInfo info)
    {
        if (true/* new account */)
        {
            UIManager.Hide("denglu");
            UIManager.Show("xuanren");
        }
        else
        {
            // onboard
        }
        return true;
    }


    public bool CreatePlayerOK(ref COM_Player player)
    {
        GamePlayer.Init(player);
        SceneLoader.LoadScene(Define.SCENE_MAIN);
        return true;
    }

    public bool JoinBattleOk(int side, int battleid, ref int[] opponentCards)
    {
        Battle.Init(side, battleid, opponentCards);
        SceneLoader.LoadScene(Define.SCENE_BATTLE);
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

    public bool PromoteUnitOK()
	{
        UIManager.Show("qianghuachenggong");
        UIManager.SetDirty("qianghuachenggong");
        UIManager.SetDirty("xiangxiziliao");
        UIManager.SetDirty("paiku");
		return true;
	}

	public bool RequestChapterStarRewardOK()
	{
        UIManager.SetDirty("jiehun");
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
}