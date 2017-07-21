public class COM_BattleActionTarget{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(ActionType!=0);
    mask.WriteBit(ActionParam!=0);
    mask.WriteBit(ActionParamExt!=0);
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
    if(ActionParamExt!=0){
      w.Write(ActionParamExt);
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
    return true;
  }
  public long InstId;
  public int ActionType;
  public int ActionParam;
  public int ActionParamExt;
}
