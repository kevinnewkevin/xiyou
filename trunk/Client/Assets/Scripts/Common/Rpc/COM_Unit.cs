public class COM_Unit{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(UnitId!=0);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(IProperties!=null&&IProperties.Length!=0);
    mask.WriteBit(CProperties!=null&&CProperties.Length!=0);
    mask.WriteBit(Equipments!=null&&Equipments.Length!=0);
    w.Write(mask.Bytes);
    //S UnitId
    if(UnitId!=0){
      w.Write(UnitId);
    }
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S IProperties
    if(IProperties!=null&&IProperties.Length!=0){
      w.WriteSize(IProperties.Length);
      for(int i=0; i<IProperties.Length; ++i){
        w.Write(IProperties[i]);
      }
    }
    //S CProperties
    if(CProperties!=null&&CProperties.Length!=0){
      w.WriteSize(CProperties.Length);
      for(int i=0; i<CProperties.Length; ++i){
        w.Write(CProperties[i]);
      }
    }
    //S Equipments
    if(Equipments!=null&&Equipments.Length!=0){
      w.WriteSize(Equipments.Length);
      for(int i=0; i<Equipments.Length; ++i){
        Equipments[i].Serialize(w);
      }
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
    //D IProperties
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      IProperties = new int[size];
      for(int i=0; i<size; ++i){
        if(!r.Read(ref IProperties[i])){
          return false;
        }
      }
    }
    //D CProperties
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      CProperties = new float[size];
      for(int i=0; i<size; ++i){
        if(!r.Read(ref CProperties[i])){
          return false;
        }
      }
    }
    //D Equipments
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      Equipments = new COM_ItemInst[size];
      for(int i=0; i<size; ++i){
        Equipments[i] = new COM_ItemInst();
        if(!Equipments[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public int UnitId;
  public long InstId;
  public int[] IProperties;
  public float[] CProperties;
  public COM_ItemInst[] Equipments;
}
