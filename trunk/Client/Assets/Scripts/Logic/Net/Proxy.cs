
class Proxy : protocol.COM_ServerToClient.Proxy
{
    public bool ErrorMessage(int err, string msg)
    {
        return true;
    }

    public bool LoginSuccess(protocol.COM_AccountInfo info)
    {
        return true;
    }


    public bool CreatePlayerSuccess(protocol.COM_PlayerInstance player)
    {
        GamePlayer.Init(player);
        World.InitPlayerActor();
        World.InitNpcActor();
         return true;
    }

    public bool BattleEnter()
    {
        SceneLoader.LoadScene(Define.SCENE_BATTLE);
        Battle.Init();
        return true;
    }

    public bool BattleReport(protocol.COM_BattleReport report)
    {
        Battle._BattleReport = report;
        return true;
    }

    public bool BattleExit(protocol.COM_BattleResult result)
    {
        Battle._Result = Battle.BattleResult.BR_Win;
        return true;
    }

    public bool SetBattleEmployeeSuccess(ulong instid)
    {
        return true;
    }

    public bool BattleSetupSuccess(System.Collections.Generic.List<protocol.COM_BattleUnit> cards)
    {
        return true;
    }
}