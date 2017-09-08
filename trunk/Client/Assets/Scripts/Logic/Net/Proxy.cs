
class Proxy : ICOM_ServerToClientProxy
{
    public bool ErrorMessage(int err)
    {
        LuaManager.Call("global.lua", "ErrorMessage", err);
        return true;
    }

    public bool LoginOK(ref COM_AccountInfo info)
    {
        return true;
    }


    public bool CreatePlayerOK(ref COM_Player player)
    {
        GamePlayer.Init(player);
        SceneLoader.LoadScene(Define.SCENE_MAIN);
        return true;
    }

    public bool JoinBattleOk(int side, int battleid)
    {
        Battle.Init(side, battleid);
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
        Battle.SetResult = (Battle.BattleResult)result.Win;
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
		JieHunSystem.instance.ChapterData = chapter;	
		return true;
	}

	public bool InitBagItems(ref COM_ItemInst[] items)
	{
		BagSystem.Init (items);
		return true;
	}

	public bool AddBagItem(ref COM_ItemInst inst)
	{
	//	BagSystem.AddItem(inst);
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
}