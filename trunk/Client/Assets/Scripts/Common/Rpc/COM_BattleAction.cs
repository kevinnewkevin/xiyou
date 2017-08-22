public class COM_BattleAction{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(BuffList!=null&&BuffList.Length!=0);
    mask.WriteBit(SkillId!=0);
    mask.WriteBit(SkillBuff!=null&&SkillBuff.Length!=0);
    mask.WriteBit(TargetList!=null&&TargetList.Length!=0);
    w.Write(mask.Bytes);
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S BuffList
    if(BuffList!=null&&BuffList.Length!=0){
      w.WriteSize(BuffList.Length);
      for(int i=0; i<BuffList.Length; ++i){
        BuffList[i].Serialize(w);
      }
    }
    //S SkillId
    if(SkillId!=0){
      w.Write(SkillId);
    }
    //S SkillBuff
    if(SkillBuff!=null&&SkillBuff.Length!=0){
      w.WriteSize(SkillBuff.Length);
      for(int i=0; i<SkillBuff.Length; ++i){
        SkillBuff[i].Serialize(w);
      }
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
    //D BuffList
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      BuffList = new COM_BattleBuffAction[size];
      for(int i=0; i<size; ++i){
        BuffList[i] = new COM_BattleBuffAction();
        if(!BuffList[i].Deserialize(r)){
          return false;
        }
      }
    }
    //D SkillId
    if(mask.ReadBit()){
      if(!r.Read(ref SkillId)){
        return false;
      }
    }
    //D SkillBuff
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      SkillBuff = new COM_BattleBuff[size];
      for(int i=0; i<size; ++i){
        SkillBuff[i] = new COM_BattleBuff();
        if(!SkillBuff[i].Deserialize(r)){
          return false;
        }
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
  public COM_BattleBuffAction[] BuffList;
  public int SkillId;
  public COM_BattleBuff[] SkillBuff;
  public COM_BattleActionTarget[] TargetList;
}
