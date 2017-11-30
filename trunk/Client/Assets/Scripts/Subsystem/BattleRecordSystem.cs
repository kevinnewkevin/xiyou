using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BattleRecordSystem {

    static Queue<COM_BattleReport> _Records;
    static COM_BattleResult _Result;
    static public COM_BattleRecord_Detail[] _BrDetail;

    //当前录像的视角玩家
    static public long MirrorPlayerId;

    //请求一个录像
    static public void RequestRecord(long recordid)
    {
        if (Battle.InBattle)
            return;
        //NetWork
        NetWoking.S.QueryBattleRecord(recordid);
    }

    //初始化录像
    static public void SetRecord(COM_BattleRecord br)
    {
        if (br == null || br.Players == null || br.Players.Length == 0)
            return;
        
        _Records = new Queue<COM_BattleReport>();
        for(int i=0; i < br.Report.Length; ++i)
        {
            _Records.Enqueue(br.Report[i]);
        }
        _Result = new COM_BattleResult();
        int side = 0;
        for(int i=0; i < br.Players.Length; ++i)
        {
            if (br.Players [i].InstId == br.Winner && MirrorPlayerId == br.Players [i].InstId)
                _Result.Win = 1;

            if (br.Players [i].InstId == MirrorPlayerId)
                side = (int)br.Players [i].Camp;
        }

        List<COM_BattleUnit> bus = new List<COM_BattleUnit>();
        List<int> opponentCards = new List<int>();
        for(int i=0; i < br.Players.Length; ++i)
        {
            bus.Add(br.Players [i].MainUnit);

            if (br.Players [i].Units == null)
                continue;
            
            for(int j=0; j < br.Players[i].Units.Length; ++j)
            {
                opponentCards.Add(br.Players[i].Units[j].UnitId);
            }
        }

        Battle.Init(side, br.Battleid, opponentCards.ToArray(), bus.ToArray(), true);
        BattleData bd = BattleData.GetData(br.Battleid);
        if (bd != null)
            SceneLoader.LoadScene(bd._SceneName);
        else
            SceneLoader.LoadScene(Define.RandomBattleScene);
    }

    //获取下一回合情况
    static public void LaunchOperate()
    {
        if (_Records != null && _Records.Count > 0)
            Battle.BattleReport = _Records.Dequeue();

        if(_Records != null && _Records.Count == 0)
            Battle.Result = _Result;
    }

    //录像简略信息
    static public void CacheSimpleData(ref COM_BattleRecord_Detail[] brdetail)
    {
        //window waiting close

        if (brdetail == null || brdetail.Length == 0)
        {
            
            return;
        }

        bool checkOk = false;
        if (MirrorPlayerId != 0)
        {
            for (int i = 0; i < brdetail [0].Players.Length; ++i)
            {
                for(int j=0; j < brdetail [0].Players.Length; ++j)
                {
                    if (brdetail [0].Players [j].InstId == MirrorPlayerId)
                    {
                        checkOk = true;
                        break;
                    }
                }
            }
        }
        if (!checkOk)
            return;

        if (brdetail[0].Battleid != 0)
            MirrorPlayerId = brdetail[0].Players [0].InstId;
        
        _BrDetail = brdetail;

        UIManager.SetDirty("luxiang");
        UIManager.SetDirty("guankaluxiang");
    }
}
