package prpc
import(
  "bytes"
  "errors"
  "suzuki/prpc"
)
type COM_ClientToServer_Login struct{
  info COM_LoginInfo  //0
}
type COM_ClientToServer_CreatePlayer struct{
  tempId int32  //0
  playerName string  //1
}
type COM_ClientToServer_AddBattleUnit struct{
  instId int64  //0
  groupId int32  //1
}
type COM_ClientToServer_PopBattleUnit struct{
  instId int64  //0
  groupId int32  //1
}
type COM_ClientToServer_SetBattleUnit struct{
  instId int64  //0
  groupId int32  //1
  isBattle bool  //2
}
type COM_ClientToServer_DelUnitGroup struct{
  groupId int32  //0
}
type COM_ClientToServer_SetupBattle struct{
  positionList []COM_BattlePosition  //0
}
type COM_ClientToServer_RequestChapterData struct{
  chapterId int32  //0
}
type COM_ClientToServer_ChallengeSmallChapter struct{
  smallChapterId int32  //0
}
type COM_ClientToServerStub struct{
  Sender prpc.StubSender
}
type COM_ClientToServerProxy interface{
  Login(info COM_LoginInfo ) error // 0
  CreatePlayer(tempId int32, playerName string ) error // 1
  AddBattleUnit(instId int64, groupId int32 ) error // 2
  PopBattleUnit(instId int64, groupId int32 ) error // 3
  SetBattleUnit(instId int64, groupId int32, isBattle bool ) error // 4
  DelUnitGroup(groupId int32 ) error // 5
  JoinBattle() error // 6
  SetupBattle(positionList []COM_BattlePosition ) error // 7
  RequestChapterData(chapterId int32 ) error // 8
  ChallengeSmallChapter(smallChapterId int32 ) error // 9
}
func (this *COM_ClientToServer_Login)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //info
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize info
  {
    err := this.info.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_Login)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.ReadBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_CreatePlayer)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.tempId!=0)
  mask.WriteBit(len(this.playerName) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize tempId
  {
    if(this.tempId!=0){
      err := prpc.Write(buffer,this.tempId)
      if err != nil{
        return err
      }
    }
  }
  // serialize playerName
  if len(this.playerName) != 0{
    err := prpc.Write(buffer,this.playerName)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_CreatePlayer)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize tempId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.tempId)
    if err != nil{
      return err
    }
  }
  // deserialize playerName
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.playerName)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_AddBattleUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instId!=0)
  mask.WriteBit(this.groupId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := prpc.Write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := prpc.Write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_AddBattleUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize groupId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_PopBattleUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instId!=0)
  mask.WriteBit(this.groupId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := prpc.Write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := prpc.Write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_PopBattleUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize groupId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SetBattleUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instId!=0)
  mask.WriteBit(this.groupId!=0)
  mask.WriteBit(this.isBattle)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := prpc.Write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := prpc.Write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  // serialize isBattle
  {
  }
  return nil
}
func (this *COM_ClientToServer_SetBattleUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize groupId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  // deserialize isBattle
  this.isBattle = mask.ReadBit();
  return nil
}
func (this *COM_ClientToServer_DelUnitGroup)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.groupId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := prpc.Write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_DelUnitGroup)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize groupId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SetupBattle)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(len(this.positionList) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize positionList
  if len(this.positionList) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.positionList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.positionList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_SetupBattle)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize positionList
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.positionList = make([]COM_BattlePosition,size)
    for i,_ := range this.positionList{
      err := this.positionList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestChapterData)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.chapterId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize chapterId
  {
    if(this.chapterId!=0){
      err := prpc.Write(buffer,this.chapterId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestChapterData)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize chapterId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.chapterId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_ChallengeSmallChapter)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.smallChapterId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize smallChapterId
  {
    if(this.smallChapterId!=0){
      err := prpc.Write(buffer,this.smallChapterId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_ChallengeSmallChapter)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize smallChapterId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.smallChapterId)
    if err != nil{
      return err
    }
  }
  return nil
}
func(this* COM_ClientToServerStub)Login(info COM_LoginInfo ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(0))
  if err != nil{
    return err
  }
  _0 := COM_ClientToServer_Login{}
  _0.info = info;
  err = _0.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)CreatePlayer(tempId int32, playerName string ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(1))
  if err != nil{
    return err
  }
  _1 := COM_ClientToServer_CreatePlayer{}
  _1.tempId = tempId;
  _1.playerName = playerName;
  err = _1.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)AddBattleUnit(instId int64, groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(2))
  if err != nil{
    return err
  }
  _2 := COM_ClientToServer_AddBattleUnit{}
  _2.instId = instId;
  _2.groupId = groupId;
  err = _2.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)PopBattleUnit(instId int64, groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(3))
  if err != nil{
    return err
  }
  _3 := COM_ClientToServer_PopBattleUnit{}
  _3.instId = instId;
  _3.groupId = groupId;
  err = _3.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SetBattleUnit(instId int64, groupId int32, isBattle bool ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(4))
  if err != nil{
    return err
  }
  _4 := COM_ClientToServer_SetBattleUnit{}
  _4.instId = instId;
  _4.groupId = groupId;
  _4.isBattle = isBattle;
  err = _4.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)DelUnitGroup(groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(5))
  if err != nil{
    return err
  }
  _5 := COM_ClientToServer_DelUnitGroup{}
  _5.groupId = groupId;
  err = _5.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)JoinBattle() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(6))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SetupBattle(positionList []COM_BattlePosition ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(7))
  if err != nil{
    return err
  }
  _7 := COM_ClientToServer_SetupBattle{}
  _7.positionList = positionList;
  err = _7.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)RequestChapterData(chapterId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(8))
  if err != nil{
    return err
  }
  _8 := COM_ClientToServer_RequestChapterData{}
  _8.chapterId = chapterId;
  err = _8.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)ChallengeSmallChapter(smallChapterId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(9))
  if err != nil{
    return err
  }
  _9 := COM_ClientToServer_ChallengeSmallChapter{}
  _9.smallChapterId = smallChapterId;
  err = _9.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func Bridging_COM_ClientToServer_Login(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _0 := COM_ClientToServer_Login{}
  err := _0.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.Login(_0.info)
}
func Bridging_COM_ClientToServer_CreatePlayer(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _1 := COM_ClientToServer_CreatePlayer{}
  err := _1.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.CreatePlayer(_1.tempId,_1.playerName)
}
func Bridging_COM_ClientToServer_AddBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _2 := COM_ClientToServer_AddBattleUnit{}
  err := _2.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.AddBattleUnit(_2.instId,_2.groupId)
}
func Bridging_COM_ClientToServer_PopBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _3 := COM_ClientToServer_PopBattleUnit{}
  err := _3.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.PopBattleUnit(_3.instId,_3.groupId)
}
func Bridging_COM_ClientToServer_SetBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _4 := COM_ClientToServer_SetBattleUnit{}
  err := _4.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SetBattleUnit(_4.instId,_4.groupId,_4.isBattle)
}
func Bridging_COM_ClientToServer_DelUnitGroup(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _5 := COM_ClientToServer_DelUnitGroup{}
  err := _5.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DelUnitGroup(_5.groupId)
}
func Bridging_COM_ClientToServer_JoinBattle(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  return p.JoinBattle()
}
func Bridging_COM_ClientToServer_SetupBattle(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _7 := COM_ClientToServer_SetupBattle{}
  err := _7.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SetupBattle(_7.positionList)
}
func Bridging_COM_ClientToServer_RequestChapterData(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _8 := COM_ClientToServer_RequestChapterData{}
  err := _8.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RequestChapterData(_8.chapterId)
}
func Bridging_COM_ClientToServer_ChallengeSmallChapter(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _9 := COM_ClientToServer_ChallengeSmallChapter{}
  err := _9.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ChallengeSmallChapter(_9.smallChapterId)
}
func COM_ClientToServerDispatch(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil {
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  pid := uint16(0XFFFF)
  err := prpc.Read(buffer,&pid)
  if err != nil{
    return err
  }
  switch(pid){
    case 0 :
      return Bridging_COM_ClientToServer_Login(buffer,p);
    case 1 :
      return Bridging_COM_ClientToServer_CreatePlayer(buffer,p);
    case 2 :
      return Bridging_COM_ClientToServer_AddBattleUnit(buffer,p);
    case 3 :
      return Bridging_COM_ClientToServer_PopBattleUnit(buffer,p);
    case 4 :
      return Bridging_COM_ClientToServer_SetBattleUnit(buffer,p);
    case 5 :
      return Bridging_COM_ClientToServer_DelUnitGroup(buffer,p);
    case 6 :
      return Bridging_COM_ClientToServer_JoinBattle(buffer,p);
    case 7 :
      return Bridging_COM_ClientToServer_SetupBattle(buffer,p);
    case 8 :
      return Bridging_COM_ClientToServer_RequestChapterData(buffer,p);
    case 9 :
      return Bridging_COM_ClientToServer_ChallengeSmallChapter(buffer,p);
    default:
      return errors.New(prpc.NoneDispatchMatchError)
  }
}
