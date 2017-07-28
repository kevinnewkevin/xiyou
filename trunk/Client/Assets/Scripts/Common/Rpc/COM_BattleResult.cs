public class COM_BattleResult{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(Win!=0);
    mask.WriteBit(Money!=0);
    w.Write(mask.Bytes);
    //S Win
    if(Win!=0){
      w.Write(Win);
    }
    //S Money
    if(Money!=0){
      w.Write(Money);
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
    return true;
  }
  public int Win;
  public int Money;
}
