public class COM_BattleReport{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(UnitList!=null&&UnitList.Length!=0);
    mask.WriteBit(ActionList!=null&&ActionList.Length!=0);
    w.Write(mask.Bytes);
    //S UnitList
    if(UnitList!=null&&UnitList.Length!=0){
      w.WriteSize(UnitList.Length);
      for(int i=0; i<UnitList.Length; ++i){
        UnitList[i].Serialize(w);
      }
    }
    //S ActionList
    if(ActionList!=null&&ActionList.Length!=0){
      w.WriteSize(ActionList.Length);
      for(int i=0; i<ActionList.Length; ++i){
        ActionList[i].Serialize(w);
      }
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D UnitList
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      UnitList = new COM_BattleUnit[size];
      for(int i=0; i<size; ++i){
        UnitList[i] = new COM_BattleUnit();
        if(!UnitList[i].Deserialize(r)){
          return false;
        }
      }
    }
    //D ActionList
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      ActionList = new COM_BattleAction[size];
      for(int i=0; i<size; ++i){
        ActionList[i] = new COM_BattleAction();
        if(!ActionList[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public COM_BattleUnit[] UnitList;
  public COM_BattleAction[] ActionList;
}
