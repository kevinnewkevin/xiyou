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
 class COM_ClientToServer_AddBattleUnit{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(instId!=0);
    mask.WriteBit(groupId!=0);
    w.Write(mask.Bytes);
    //S instId
    if(instId!=0){
      w.Write(instId);
    }
    //S groupId
    if(groupId!=0){
      w.Write(groupId);
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
    //D groupId
    if(mask.ReadBit()){
      if(!r.Read(ref groupId)){
        return false;
      }
    }
    return true;
  }
  public long instId;
  public int groupId;
}
 class COM_ClientToServer_PopBattleUnit{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(instId!=0);
    mask.WriteBit(groupId!=0);
    w.Write(mask.Bytes);
    //S instId
    if(instId!=0){
      w.Write(instId);
    }
    //S groupId
    if(groupId!=0){
      w.Write(groupId);
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
    //D groupId
    if(mask.ReadBit()){
      if(!r.Read(ref groupId)){
        return false;
      }
    }
    return true;
  }
  public long instId;
  public int groupId;
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
 class COM_ClientToServer_RequestChapterData{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(chapterId!=0);
    w.Write(mask.Bytes);
    //S chapterId
    if(chapterId!=0){
      w.Write(chapterId);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D chapterId
    if(mask.ReadBit()){
      if(!r.Read(ref chapterId)){
        return false;
      }
    }
    return true;
  }
  public int chapterId;
}
 class COM_ClientToServer_ChallengeSmallChapter{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(smallChapterId!=0);
    w.Write(mask.Bytes);
    //S smallChapterId
    if(smallChapterId!=0){
      w.Write(smallChapterId);
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D smallChapterId
    if(mask.ReadBit()){
      if(!r.Read(ref smallChapterId)){
        return false;
      }
    }
    return true;
  }
  public int smallChapterId;
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
  public void AddBattleUnit(long instId, int groupId ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)2);
    _2.instId = instId;
    _2.groupId = groupId;
    _2.Serialize(w);
    MethodEnd();
  }
  public void PopBattleUnit(long instId, int groupId ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)3);
    _3.instId = instId;
    _3.groupId = groupId;
    _3.Serialize(w);
    MethodEnd();
  }
  public void JoinBattle(){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)4);
    MethodEnd();
  }
  public void SetupBattle(COM_BattlePosition[] positionList ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)5);
    _5.positionList = positionList;
    _5.Serialize(w);
    MethodEnd();
  }
  public void RequestChapterData(int chapterId ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)6);
    _6.chapterId = chapterId;
    _6.Serialize(w);
    MethodEnd();
  }
  public void ChallengeSmallChapter(int smallChapterId ){
    IWriter w = MethodBegin();
    if(w==null){
      return;
    }
    w.Write((ushort)7);
    _7.smallChapterId = smallChapterId;
    _7.Serialize(w);
    MethodEnd();
  }
  public abstract IWriter MethodBegin();
  public abstract void MethodEnd();
  
  COM_ClientToServer_Login _0 = new COM_ClientToServer_Login();
  COM_ClientToServer_CreatePlayer _1 = new COM_ClientToServer_CreatePlayer();
  COM_ClientToServer_AddBattleUnit _2 = new COM_ClientToServer_AddBattleUnit();
  COM_ClientToServer_PopBattleUnit _3 = new COM_ClientToServer_PopBattleUnit();
  COM_ClientToServer_SetupBattle _5 = new COM_ClientToServer_SetupBattle();
  COM_ClientToServer_RequestChapterData _6 = new COM_ClientToServer_RequestChapterData();
  COM_ClientToServer_ChallengeSmallChapter _7 = new COM_ClientToServer_ChallengeSmallChapter();
};
public interface ICOM_ClientToServerProxy {
  bool Login(ref COM_LoginInfo info ); // 0
  bool CreatePlayer( int tempId, ref string playerName ); // 1
  bool AddBattleUnit( long instId,  int groupId ); // 2
  bool PopBattleUnit( long instId,  int groupId ); // 3
  bool JoinBattle(); // 4
  bool SetupBattle(ref COM_BattlePosition[] positionList ); // 5
  bool RequestChapterData( int chapterId ); // 6
  bool ChallengeSmallChapter( int smallChapterId ); // 7
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
  public static bool AddBattleUnit(IReader r, ICOM_ClientToServerProxy p){ // 2
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_2.Deserialize(r)){
      return false;
    }
    return p.AddBattleUnit( _2.instId, _2.groupId);
  }
  public static bool PopBattleUnit(IReader r, ICOM_ClientToServerProxy p){ // 3
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_3.Deserialize(r)){
      return false;
    }
    return p.PopBattleUnit( _3.instId, _3.groupId);
  }
  public static bool JoinBattle(IReader r, ICOM_ClientToServerProxy p){ // 4
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    return p.JoinBattle();
  }
  public static bool SetupBattle(IReader r, ICOM_ClientToServerProxy p){ // 5
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_5.Deserialize(r)){
      return false;
    }
    return p.SetupBattle(ref _5.positionList);
  }
  public static bool RequestChapterData(IReader r, ICOM_ClientToServerProxy p){ // 6
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_6.Deserialize(r)){
      return false;
    }
    return p.RequestChapterData( _6.chapterId);
  }
  public static bool ChallengeSmallChapter(IReader r, ICOM_ClientToServerProxy p){ // 7
    if(r==null){
      return false;
    }
    if(p==null){
      return false;
    }
    if(!_7.Deserialize(r)){
      return false;
    }
    return p.ChallengeSmallChapter( _7.smallChapterId);
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
        return AddBattleUnit(r,p);
      case 3 :
        return PopBattleUnit(r,p);
      case 4 :
        return JoinBattle(r,p);
      case 5 :
        return SetupBattle(r,p);
      case 6 :
        return RequestChapterData(r,p);
      case 7 :
        return ChallengeSmallChapter(r,p);
      default:
        return false;
    }
  }
  
  static COM_ClientToServer_Login _0 = new COM_ClientToServer_Login();
  static COM_ClientToServer_CreatePlayer _1 = new COM_ClientToServer_CreatePlayer();
  static COM_ClientToServer_AddBattleUnit _2 = new COM_ClientToServer_AddBattleUnit();
  static COM_ClientToServer_PopBattleUnit _3 = new COM_ClientToServer_PopBattleUnit();
  static COM_ClientToServer_SetupBattle _5 = new COM_ClientToServer_SetupBattle();
  static COM_ClientToServer_RequestChapterData _6 = new COM_ClientToServer_RequestChapterData();
  static COM_ClientToServer_ChallengeSmallChapter _7 = new COM_ClientToServer_ChallengeSmallChapter();
}
