public class COM_BattleAction{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(SkillId!=0);
    mask.WriteBit(TargetList!=null&&TargetList.Length!=0);
    w.Write(mask.Bytes);
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S SkillId
    if(SkillId!=0){
      w.Write(SkillId);
    }
    //S TargetList
    if(TargetList!=null&&TargetList.Length!=0){
      w.WriteSize(TargetList.Length);
      for(int i=0; i<TargetList.Length; ++i){
        TargetList[i].Serialize(w);
      }
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
    //D SkillId
    if(mask.ReadBit()){
      if(!r.Read(ref SkillId)){
        return false;
      }
    }
    //D TargetList
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      TargetList = new COM_BattleActionTarget[size];
      for(int i=0; i<size; ++i){
        TargetList[i] = new COM_BattleActionTarget();
        if(!TargetList[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public long InstId;
  public int SkillId;
  public COM_BattleActionTarget[] TargetList;
}
