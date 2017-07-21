public class COM_BattlePosition{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(Position!=0);
    w.Write(mask.Bytes);
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S Position
    if(Position!=0){
      w.Write(Position);
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
    //D Position
    if(mask.ReadBit()){
      if(!r.Read(ref Position)){
        return false;
      }
    }
    return true;
  }
  public long InstId;
  public int Position;
}
