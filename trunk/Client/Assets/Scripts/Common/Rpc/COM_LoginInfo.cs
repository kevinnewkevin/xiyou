public class COM_LoginInfo{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(Username!=null&&Username.Length!=0&&Username!="");
    mask.WriteBit(Password!=null&&Password.Length!=0&&Password!="");
    w.Write(mask.Bytes);
    //S Username
    if(Username!=null&&Username.Length!=0&&Username!=""){
      w.Write(Username);
    }
    //S Password
    if(Password!=null&&Password.Length!=0&&Password!=""){
      w.Write(Password);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D Username
    if(mask.ReadBit()){
      if(!r.Read(ref Username)){
        return false;
      }
    }
    //D Password
    if(mask.ReadBit()){
      if(!r.Read(ref Password)){
        return false;
      }
    }
    return true;
  }
  public string Username;
  public string Password;
}
