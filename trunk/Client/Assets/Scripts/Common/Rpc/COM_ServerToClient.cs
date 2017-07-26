 class COM_ServerToClient_ErrorMessage{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(id!=0);
    w.Write(mask.Bytes);
    //S id
    if(id!=0){
      w.Write((byte)id);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D id
    if(mask.ReadBit()){
      byte enumer = 0XFF;
      if(!r.Read(ref enumer) || enumer >= 5 ){
        return false;
      }
      id = enumer;
    }
    return true;
  }
  public int id;
}
 class COM_ServerToClient_LoginOK{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(true); // info
    w.Write(mask.Bytes);
    //S info
    {
      info.Serialize(w);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D info
    if(mask.ReadBit()){
      if(info==null){
        info = new COM_AccountInfo();
      }
      if(!info.Deserialize(r)){
        return false;
      }
    }
    return true;
  }
  public COM_AccountInfo info;
}
 class COM_ServerToClient_CreatePlayerOK{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(true); // player
    w.Write(mask.Bytes);
    //S player
    {
      player.Serialize(w);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D player
    if(mask.ReadBit()){
      if(player==null){
        player = new COM_Player();
      }
      if(!player.Deserialize(r)){
        return false;
      }
    }
    return true;
  }
  public COM_Player player;
}
 class COM_ServerToClient_JoinBattleOk{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(Camp!=0);
    w.Write(mask.Bytes);
    //S Camp
    if(Camp!=0){
      w.Write(Camp);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D Camp
    if(mask.ReadBit()){
      if(!r.Read(ref Camp)){
        return false;
      }
    }
    return true;
  }
  public int Camp;
}
 class COM_ServerToClient_SetBattleUnitOK{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(instId!=0);
    w.Write(mask.Bytes);
    //S instId
    if(instId!=0){
      w.Write(instId);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D instId
    if(mask.ReadBit()){
      if(!r.Read(ref instId)){
        return false;
      }
    }
    return true;
  }
  public long instId;
}
 class COM_ServerToClient_BattleReport{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(true); // report
    w.Write(mask.Bytes);
    //S report
    {
      report.Serialize(w);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D report
    if(mask.ReadBit()){
      if(report==null){
        report = new COM_BattleReport();
      }
      if(!report.Deserialize(r)){
        return false;
      }
    }
    return true;
  }
  public COM_BattleReport report;
}
 class COM_ServerToClient_BattleExit{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(true); // result
    w.Write(mask.Bytes);
    //S result
    {
      result.Serialize(w);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D result
    if(mask.ReadBit()){
      if(result==null){
        result = new COM_BattleResult();
      }
      if(!result.Deserialize(r)){
        return false;
      }
    }
    return true;
  }
  public COM_BattleResult result;
}
public abstract class COM_ServerToClientStub {
  public void ErrorMessage(int id ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)0);
    _0.id = id;
    _0.Serialize(w);
    MethodEnd();
  }
  public void LoginOK(COM_AccountInfo info ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)1);
    _1.info = info;
    _1.Serialize(w);
    MethodEnd();
  }
  public void CreatePlayerOK(COM_Player player ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)2);
    _2.player = player;
    _2.Serialize(w);
    MethodEnd();
  }
  public void JoinBattleOk(int Camp ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)3);
    _3.Camp = Camp;
    _3.Serialize(w);
    MethodEnd();
  }
  public void SetupBattleOK(){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)4);
    MethodEnd();
  }
  public void SetBattleUnitOK(long instId ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)5);
    _5.instId = instId;
    _5.Serialize(w);
    MethodEnd();
  }
  public void BattleReport(COM_BattleReport report ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)6);
    _6.report = report;
    _6.Serialize(w);
    MethodEnd();
  }
  public void BattleExit(COM_BattleResult result ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)7);
    _7.result = result;
    _7.Serialize(w);
    MethodEnd();
  }
  public abstract IWriter MethodBegin();
  public abstract void MethodEnd();
  
  COM_ServerToClient_ErrorMessage _0 = new COM_ServerToClient_ErrorMessage();
  COM_ServerToClient_LoginOK _1 = new COM_ServerToClient_LoginOK();
  COM_ServerToClient_CreatePlayerOK _2 = new COM_ServerToClient_CreatePlayerOK();
  COM_ServerToClient_JoinBattleOk _3 = new COM_ServerToClient_JoinBattleOk();
  COM_ServerToClient_SetBattleUnitOK _5 = new COM_ServerToClient_SetBattleUnitOK();
  COM_ServerToClient_BattleReport _6 = new COM_ServerToClient_BattleReport();
  COM_ServerToClient_BattleExit _7 = new COM_ServerToClient_BattleExit();
};
public interface ICOM_ServerToClientProxy {
  bool ErrorMessage( int id ); // 0
  bool LoginOK(ref COM_AccountInfo info ); // 1
  bool CreatePlayerOK(ref COM_Player player ); // 2
  bool JoinBattleOk( int Camp ); // 3
  bool SetupBattleOK(); // 4
  bool SetBattleUnitOK( long instId ); // 5
  bool BattleReport(ref COM_BattleReport report ); // 6
  bool BattleExit(ref COM_BattleResult result ); // 7
}
public class COM_ServerToClientDispatcher {
  public static bool ErrorMessage(IReader r, ICOM_ServerToClientProxy p){ // 0
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_0.Deserialize(r)){
      return false;
    }
    return p.ErrorMessage( _0.id);
  }
  public static bool LoginOK(IReader r, ICOM_ServerToClientProxy p){ // 1
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_1.Deserialize(r)){
      return false;
    }
    return p.LoginOK(ref _1.info);
  }
  public static bool CreatePlayerOK(IReader r, ICOM_ServerToClientProxy p){ // 2
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_2.Deserialize(r)){
      return false;
    }
    return p.CreatePlayerOK(ref _2.player);
  }
  public static bool JoinBattleOk(IReader r, ICOM_ServerToClientProxy p){ // 3
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_3.Deserialize(r)){
      return false;
    }
    return p.JoinBattleOk( _3.Camp);
  }
  public static bool SetupBattleOK(IReader r, ICOM_ServerToClientProxy p){ // 4
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    return p.SetupBattleOK();
  }
  public static bool SetBattleUnitOK(IReader r, ICOM_ServerToClientProxy p){ // 5
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_5.Deserialize(r)){
      return false;
    }
    return p.SetBattleUnitOK( _5.instId);
  }
  public static bool BattleReport(IReader r, ICOM_ServerToClientProxy p){ // 6
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_6.Deserialize(r)){
      return false;
    }
    return p.BattleReport(ref _6.report);
  }
  public static bool BattleExit(IReader r, ICOM_ServerToClientProxy p){ // 7
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_7.Deserialize(r)){
      return false;
    }
    return p.BattleExit(ref _7.result);
  }
  public static bool Execute(IReader r, ICOM_ServerToClientProxy p){
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    ushort pid = 0XFFFF;
    if(!r.Read(ref pid)){
      return false;
    }
    switch(pid){
      case 0 :
        return ErrorMessage(r,p);
      case 1 :
        return LoginOK(r,p);
      case 2 :
        return CreatePlayerOK(r,p);
      case 3 :
        return JoinBattleOk(r,p);
      case 4 :
        return SetupBattleOK(r,p);
      case 5 :
        return SetBattleUnitOK(r,p);
      case 6 :
        return BattleReport(r,p);
      case 7 :
        return BattleExit(r,p);
      default:
        return false;
    }
  }
  
  static COM_ServerToClient_ErrorMessage _0 = new COM_ServerToClient_ErrorMessage();
  static COM_ServerToClient_LoginOK _1 = new COM_ServerToClient_LoginOK();
  static COM_ServerToClient_CreatePlayerOK _2 = new COM_ServerToClient_CreatePlayerOK();
  static COM_ServerToClient_JoinBattleOk _3 = new COM_ServerToClient_JoinBattleOk();
  static COM_ServerToClient_SetBattleUnitOK _5 = new COM_ServerToClient_SetBattleUnitOK();
  static COM_ServerToClient_BattleReport _6 = new COM_ServerToClient_BattleReport();
  static COM_ServerToClient_BattleExit _7 = new COM_ServerToClient_BattleExit();
}
