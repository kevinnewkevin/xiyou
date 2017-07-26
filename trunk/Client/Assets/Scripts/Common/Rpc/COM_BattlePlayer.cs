public class COM_BattlePlayer{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(true); // Player
    mask.WriteBit(MaxPoint!=0);
    mask.WriteBit(CurPoint!=0);
    mask.WriteBit(BattlePosition!=null&&BattlePosition.Length!=0);
    w.Write(mask.Bytes);
    //S Player
    {
      Player.Serialize(w);
    }
    //S MaxPoint
    if(MaxPoint!=0){
      w.Write(MaxPoint);
    }
    //S CurPoint
    if(CurPoint!=0){
      w.Write(CurPoint);
    }
    //S BattlePosition
    if(BattlePosition!=null&&BattlePosition.Length!=0){
      w.WriteSize(BattlePosition.Length);
      for(int i=0; i<BattlePosition.Length; ++i){
        BattlePosition[i].Serialize(w);
      }
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D Player
    if(mask.ReadBit()){
      if(Player==null){
        Player = new COM_Player();
      }
      if(!Player.Deserialize(r)){
        return false;
      }
    }
    //D MaxPoint
    if(mask.ReadBit()){
      if(!r.Read(ref MaxPoint)){
        return false;
      }
    }
    //D CurPoint
    if(mask.ReadBit()){
      if(!r.Read(ref CurPoint)){
        return false;
      }
    }
    //D BattlePosition
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      BattlePosition = new COM_BattlePosition[size];
      for(int i=0; i<size; ++i){
        BattlePosition[i] = new COM_BattlePosition();
        if(!BattlePosition[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public COM_Player Player;
  public int MaxPoint;
  public int CurPoint;
  public COM_BattlePosition[] BattlePosition;
}
