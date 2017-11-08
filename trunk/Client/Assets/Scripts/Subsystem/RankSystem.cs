using System.Collections.Generic;

public class RankSystem {

    static public List<COM_TopUnit> _AllRank;
    static public List<COM_TopUnit> _FirendRank;

    static public int _MyFirendRank;
    static public int _MyAllRank;

    static public void Init()
    {
        _AllRank = new List<COM_TopUnit>();
        _FirendRank = new List<COM_TopUnit>();
    }
}
