public class COM_BattleUnit{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(UnitId!=0);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(Position!=0);
    mask.WriteBit(HP!=0);
    mask.WriteBit(Camp!=0);
    mask.WriteBit(Name!=null&&Name.Length!=0&&Name!="");
    w.Write(mask.Bytes);
    //S UnitId
    if(UnitId!=0){
      w.Write(UnitId);
    }
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S Position
    if(Position!=0){
      w.Write(Position);
    }
    //S HP
    if(HP!=0){
      w.Write(HP);
    }
    //S Camp
    if(Camp!=0){
      w.Write((byte)Camp);
    }
    //S Name
    if(Name!=null&&Name.Length!=0&&Name!=""){
      w.Write(Name);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D UnitId
    if(mask.ReadBit()){
      if(!r.Read(ref UnitId)){
        return false;
      }
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
    //D HP
    if(mask.ReadBit()){
      if(!r.Read(ref HP)){
        return false;
      }
    }
    //D Camp
    if(mask.ReadBit()){
      byte enumer = 0XFF;
      if(!r.Read(ref enumer) || enumer >= 3 ){
        return false;
      }
      Camp = enumer;
    }
    //D Name
    if(mask.ReadBit()){
      if(!r.Read(ref Name)){
        return false;
      }
    }
    return true;
  }
  public int UnitId;
  public long InstId;
  public int Position;
  public int HP;
  public int Camp;
  public string Name;
}
