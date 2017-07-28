
class Proxy : ICOM_ServerToClientProxy
{
    public bool ErrorMessage(int err)
    {
        return true;
    }

    public bool LoginOK(ref COM_AccountInfo info)
    {
        return true;
    }


    public bool CreatePlayerOK(ref COM_Player player)
    {
        UnityEngine.Debug.Log("createplayerok");
        GamePlayer.Init(player);
        World.InitPlayerActor();
        World.InitNpcActor();
        return true;
    }

    public bool JoinBattleOk(int side)
    {
        SceneLoader.LoadScene(Define.SCENE_BATTLE);
        Battle.Init(side);
        return true;
    }

    public bool BattleReport(ref COM_BattleReport report)
    {
        Battle._BattleReport = report;
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
}