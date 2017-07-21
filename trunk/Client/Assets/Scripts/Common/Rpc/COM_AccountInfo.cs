public class COM_AccountInfo{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(SessionCode!=null&&SessionCode.Length!=0&&SessionCode!="");
    w.Write(mask.Bytes);
    //S SessionCode
    if(SessionCode!=null&&SessionCode.Length!=0&&SessionCode!=""){
      w.Write(SessionCode);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D SessionCode
    if(mask.ReadBit()){
      if(!r.Read(ref SessionCode)){
        return false;
      }
    }
    return true;
  }
  public string SessionCode;
}
