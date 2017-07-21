 class COM_ClientToServer_Login{
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
        info = new COM_LoginInfo();
      }
      if(!info.Deserialize(r)){
        return false;
      }
    }
    return true;
  }
  public COM_LoginInfo info;
}
 class COM_ClientToServer_CreatePlayer{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(tempId!=0);
    mask.WriteBit(playerName!=null&&playerName.Length!=0&&playerName!="");
    w.Write(mask.Bytes);
    //S tempId
    if(tempId!=0){
      w.Write(tempId);
    }
    //S playerName
    if(playerName!=null&&playerName.Length!=0&&playerName!=""){
      w.Write(playerName);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D tempId
    if(mask.ReadBit()){
      if(!r.Read(ref tempId)){
        return false;
      }
    }
    //D playerName
    if(mask.ReadBit()){
      if(!r.Read(ref playerName)){
        return false;
      }
    }
    return true;
  }
  public int tempId;
  public string playerName;
}
 class COM_ClientToServer_SetBattleUnit{
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
 class COM_ClientToServer_SetupBattle{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(positionList!=null&&positionList.Length!=0);
    w.Write(mask.Bytes);
    //S positionList
    if(positionList!=null&&positionList.Length!=0){
      w.WriteSize(positionList.Length);
      for(int i=0; i<positionList.Length; ++i){
        positionList[i].Serialize(w);
      }
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D positionList
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      positionList = new COM_BattlePosition[size];
      for(int i=0; i<size; ++i){
        positionList[i] = new COM_BattlePosition();
        if(!positionList[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public COM_BattlePosition[] positionList;
}
public abstract class COM_ClientToServerStub {
  public void Login(COM_LoginInfo info ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)0);
    _0.info = info;
    _0.Serialize(w);
    MethodEnd();
  }
  public void CreatePlayer(int tempId, string playerName ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)1);
    _1.tempId = tempId;
    _1.playerName = playerName;
    _1.Serialize(w);
    MethodEnd();
  }
  public void SetBattleUnit(long instId ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)2);
    _2.instId = instId;
    _2.Serialize(w);
    MethodEnd();
  }
  public void JoinBattle(){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)3);
    MethodEnd();
  }
  public void SetupBattle(COM_BattlePosition[] positionList ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)4);
    _4.positionList = positionList;
    _4.Serialize(w);
    MethodEnd();
  }
  public abstract IWriter MethodBegin();
  public abstract void MethodEnd();
  
  COM_ClientToServer_Login _0 = new COM_ClientToServer_Login();
  COM_ClientToServer_CreatePlayer _1 = new COM_ClientToServer_CreatePlayer();
  COM_ClientToServer_SetBattleUnit _2 = new COM_ClientToServer_SetBattleUnit();
  COM_ClientToServer_SetupBattle _4 = new COM_ClientToServer_SetupBattle();
};
public interface ICOM_ClientToServerProxy {
  bool Login(ref COM_LoginInfo info ); // 0
  bool CreatePlayer( int tempId, ref string playerName ); // 1
  bool SetBattleUnit( long instId ); // 2
  bool JoinBattle(); // 3
  bool SetupBattle(ref COM_BattlePosition[] positionList ); // 4
}
public class COM_ClientToServerDispatcher {
  public static bool Login(IReader r, ICOM_ClientToServerProxy p){ // 0
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_0.Deserialize(r)){
      return false;
    }
    return p.Login(ref _0.info);
  }
  public static bool CreatePlayer(IReader r, ICOM_ClientToServerProxy p){ // 1
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_1.Deserialize(r)){
      return false;
    }
    return p.CreatePlayer( _1.tempId,ref _1.playerName);
  }
  public static bool SetBattleUnit(IReader r, ICOM_ClientToServerProxy p){ // 2
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_2.Deserialize(r)){
      return false;
    }
    return p.SetBattleUnit( _2.instId);
  }
  public static bool JoinBattle(IReader r, ICOM_ClientToServerProxy p){ // 3
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    return p.JoinBattle();
  }
  public static bool SetupBattle(IReader r, ICOM_ClientToServerProxy p){ // 4
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_4.Deserialize(r)){
      return false;
    }
    return p.SetupBattle(ref _4.positionList);
  }
  public static bool Execute(IReader r, ICOM_ClientToServerProxy p){
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
        return Login(r,p);
      case 1 :
        return CreatePlayer(r,p);
      case 2 :
        return SetBattleUnit(r,p);
      case 3 :
        return JoinBattle(r,p);
      case 4 :
        return SetupBattle(r,p);
      default:
        return false;
    }
  }
  
  static COM_ClientToServer_Login _0 = new COM_ClientToServer_Login();
  static COM_ClientToServer_CreatePlayer _1 = new COM_ClientToServer_CreatePlayer();
  static COM_ClientToServer_SetBattleUnit _2 = new COM_ClientToServer_SetBattleUnit();
  static COM_ClientToServer_SetupBattle _4 = new COM_ClientToServer_SetupBattle();
}
