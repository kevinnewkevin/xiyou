public class COM_BattleRoom{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(Status!=0);
    mask.WriteBit(PlayerList!=null&&PlayerList.Length!=0);
    mask.WriteBit(true); // Target
    mask.WriteBit(Bout!=0);
    mask.WriteBit(TurnMove!=0);
    mask.WriteBit(true); // NextPlayer
    w.Write(mask.Bytes);
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S Status
    if(Status!=0){
      w.Write(Status);
    }
    //S PlayerList
    if(PlayerList!=null&&PlayerList.Length!=0){
      w.WriteSize(PlayerList.Length);
      for(int i=0; i<PlayerList.Length; ++i){
        PlayerList[i].Serialize(w);
      }
    }
    //S Target
    {
      Target.Serialize(w);
    }
    //S Bout
    if(Bout!=0){
      w.Write(Bout);
    }
    //S TurnMove
    if(TurnMove!=0){
      w.Write(TurnMove);
    }
    //S NextPlayer
    {
      NextPlayer.Serialize(w);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D InstId
    if(mask.ReadBit()){
      if(!r.Read(ref InstId)){
        return false;
      }
    }
    //D Status
    if(mask.ReadBit()){
      if(!r.Read(ref Status)){
        return false;
      }
    }
    //D PlayerList
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      PlayerList = new COM_BattlePlayer[size];
      for(int i=0; i<size; ++i){
        PlayerList[i] = new COM_BattlePlayer();
        if(!PlayerList[i].Deserialize(r)){
          return false;
        }
      }
    }
    //D Target
    if(mask.ReadBit()){
      if(Target==null){
        Target = new COM_BattlePlayer();
      }
      if(!Target.Deserialize(r)){
        return false;
      }
    }
    //D Bout
    if(mask.ReadBit()){
      if(!r.Read(ref Bout)){
        return false;
      }
    }
    //D TurnMove
    if(mask.ReadBit()){
      if(!r.Read(ref TurnMove)){
        return false;
      }
    }
    //D NextPlayer
    if(mask.ReadBit()){
      if(NextPlayer==null){
        NextPlayer = new COM_BattlePlayer();
      }
      if(!NextPlayer.Deserialize(r)){
        return false;
      }
    }
    return true;
  }
  public long InstId;
  public int Status;
  public COM_BattlePlayer[] PlayerList;
  public COM_BattlePlayer Target;
  public int Bout;
  public int TurnMove;
  public COM_BattlePlayer NextPlayer;
}
