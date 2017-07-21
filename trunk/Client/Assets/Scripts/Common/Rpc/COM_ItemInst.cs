public class COM_ItemInst{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(ItemId!=0);
    mask.WriteBit(InstId!=0);
    w.Write(mask.Bytes);
    //S ItemId
    if(ItemId!=0){
      w.Write(ItemId);
    }
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D ItemId
    if(mask.ReadBit()){
      if(!r.Read(ref ItemId)){
        return false;
      }
    }
    //D InstId
    if(mask.ReadBit()){
      if(!r.Read(ref InstId)){
        return false;
      }
    }
    return true;
  }
  public int ItemId;
  public long InstId;
}
