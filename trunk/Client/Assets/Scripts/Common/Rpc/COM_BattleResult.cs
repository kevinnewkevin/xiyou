public class COM_BattleResult{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(Win!=0);
    mask.WriteBit(Money!=0);
    mask.WriteBit(KillMonsters!=null&&KillMonsters.Length!=0);
    mask.WriteBit(BattleRound!=0);
    mask.WriteBit(MySelfDeathNum!=0);
    w.Write(mask.Bytes);
    //S Win
    if(Win!=0){
      w.Write(Win);
    }
    //S Money
    if(Money!=0){
      w.Write(Money);
    }
    //S KillMonsters
    if(KillMonsters!=null&&KillMonsters.Length!=0){
      w.WriteSize(KillMonsters.Length);
      for(int i=0; i<KillMonsters.Length; ++i){
        w.Write(KillMonsters[i]);
      }
    }
    //S BattleRound
    if(BattleRound!=0){
      w.Write(BattleRound);
    }
    //S MySelfDeathNum
    if(MySelfDeathNum!=0){
      w.Write(MySelfDeathNum);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D Win
    if(mask.ReadBit()){
      if(!r.Read(ref Win)){
        return false;
      }
    }
    //D Money
    if(mask.ReadBit()){
      if(!r.Read(ref Money)){
        return false;
      }
    }
    //D KillMonsters
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      KillMonsters = new int[size];
      for(int i=0; i<size; ++i){
        if(!r.Read(ref KillMonsters[i])){
          return false;
        }
      }
    }
    //D BattleRound
    if(mask.ReadBit()){
      if(!r.Read(ref BattleRound)){
        return false;
      }
    }
    //D MySelfDeathNum
    if(mask.ReadBit()){
      if(!r.Read(ref MySelfDeathNum)){
        return false;
      }
    }
    return true;
  }
  public int Win;
  public int Money;
  public int[] KillMonsters;
  public int BattleRound;
  public int MySelfDeathNum;
}
