public class COM_BattleResult{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(Money!=0);
    w.Write(mask.Bytes);
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
    //D Money
    if(mask.ReadBit()){
      if(!r.Read(ref Money)){
        return false;
      }
    }
    return true;
  }
  public int Money;
}
