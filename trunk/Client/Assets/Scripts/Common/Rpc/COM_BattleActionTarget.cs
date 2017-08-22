public class COM_BattleActionTarget{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(ActionType!=0);
    mask.WriteBit(ActionParam!=0);
    mask.WriteBit(ActionParamExt!=null&&ActionParamExt.Length!=0&&ActionParamExt!="");
    mask.WriteBit(!Dead);
    mask.WriteBit(BuffAdd!=null&&BuffAdd.Length!=0);
    w.Write(mask.Bytes);
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S ActionType
    if(ActionType!=0){
      w.Write((byte)ActionType);
    }
    //S ActionParam
    if(ActionParam!=0){
      w.Write(ActionParam);
    }
    //S ActionParamExt
    if(ActionParamExt!=null&&ActionParamExt.Length!=0&&ActionParamExt!=""){
      w.Write(ActionParamExt);
    }
    //S Dead
    {
    }
    //S BuffAdd
    if(BuffAdd!=null&&BuffAdd.Length!=0){
      w.WriteSize(BuffAdd.Length);
      for(int i=0; i<BuffAdd.Length; ++i){
        BuffAdd[i].Serialize(w);
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
    //D ActionType
    if(mask.ReadBit()){
      byte enumer = 0XFF;
      if(!r.Read(ref enumer) || enumer >= 7 ){
        return false;
      }
      ActionType = enumer;
    }
    //D ActionParam
    if(mask.ReadBit()){
      if(!r.Read(ref ActionParam)){
        return false;
      }
    }
    //D ActionParamExt
    if(mask.ReadBit()){
      if(!r.Read(ref ActionParamExt)){
        return false;
      }
    }
    //D Dead
    {
      Dead = mask.ReadBit();
    }
    //D BuffAdd
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      BuffAdd = new COM_BattleBuff[size];
      for(int i=0; i<size; ++i){
        BuffAdd[i] = new COM_BattleBuff();
        if(!BuffAdd[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public long InstId;
  public int ActionType;
  public int ActionParam;
  public string ActionParamExt;
  public bool Dead;
  public COM_BattleBuff[] BuffAdd;
}
