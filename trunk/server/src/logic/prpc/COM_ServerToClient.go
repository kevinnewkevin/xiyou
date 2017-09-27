package prpc
import(
  "bytes"
  "errors"
  "suzuki/prpc"
)
type COM_ServerToClient_ErrorMessage struct{
  id int  //0
}
type COM_ServerToClient_LoginOK struct{
  info COM_AccountInfo  //0
}
type COM_ServerToClient_CreatePlayerOK struct{
  player COM_Player  //0
}
type COM_ServerToClient_JoinBattleOk struct{
  Camp int32  //0
  battleid int32  //1
  targetcards []int32  //2
}
type COM_ServerToClient_SetBattleUnitOK struct{
  instId int64  //0
}
type COM_ServerToClient_BattleReport struct{
  report COM_BattleReport  //0
}
type COM_ServerToClient_BattleExit struct{
  result COM_BattleResult  //0
}
type COM_ServerToClient_OpenChapter struct{
  data COM_Chapter  //0
}
type COM_ServerToClient_SycnChapterData struct{
  data COM_Chapter  //0
}
type COM_ServerToClient_InitBagItems struct{
  items []COM_ItemInst  //0
}
type COM_ServerToClient_AddBagItem struct{
  item COM_ItemInst  //0
}
type COM_ServerToClient_UpdateBagItem struct{
  item COM_ItemInst  //0
}
type COM_ServerToClient_DeleteItemOK struct{
  instId int64  //0
}
type COM_ServerToClient_UpdateTiantiVal struct{
  curVal int32  //0
}
type COM_ServerToClient_UpdateUnitIProperty struct{
  instid int64  //0
  iType int32  //1
  value int32  //2
}
type COM_ServerToClient_UpdateUnitCProperty struct{
  instid int64  //0
  cType int32  //1
  value float32  //2
}
type COM_ServerToClient_EquipSkillOK struct{
  skillIndex int32  //0
  skillID int32  //1
}
type COM_ServerToClient_SkillUpdateOK struct{
  skillIndex int32  //0
  skillID int32  //1
  skillpos int32  //2
}
type COM_ServerToClient_BuyShopItemOK struct{
  items []int32  //0
}
type COM_ServerToClientStub struct{
  Sender prpc.StubSender
}
type COM_ServerToClientProxy interface{
  ErrorMessage(id int ) error // 0
  LoginOK(info COM_AccountInfo ) error // 1
  CreatePlayerOK(player COM_Player ) error // 2
  JoinBattleOk(Camp int32, battleid int32, targetcards []int32 ) error // 3
  SetupBattleOK() error // 4
  SetBattleUnitOK(instId int64 ) error // 5
  BattleReport(report COM_BattleReport ) error // 6
  BattleExit(result COM_BattleResult ) error // 7
  OpenChapter(data COM_Chapter ) error // 8
  SycnChapterData(data COM_Chapter ) error // 9
  InitBagItems(items []COM_ItemInst ) error // 10
  AddBagItem(item COM_ItemInst ) error // 11
  UpdateBagItem(item COM_ItemInst ) error // 12
  DeleteItemOK(instId int64 ) error // 13
  UpdateTiantiVal(curVal int32 ) error // 14
  UpdateUnitIProperty(instid int64, iType int32, value int32 ) error // 15
  UpdateUnitCProperty(instid int64, cType int32, value float32 ) error // 16
  PromoteUnitOK() error // 17
  RequestChapterStarRewardOK() error // 18
  EquipSkillOK(skillIndex int32, skillID int32 ) error // 19
  SkillUpdateOK(skillIndex int32, skillID int32, skillpos int32 ) error // 20
  BuyShopItemOK(items []int32 ) error // 21
}
func (this *COM_ServerToClient_ErrorMessage)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.id!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize id
  {
    if(this.id!=0){
      err := prpc.Write(buffer,this.id)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_ErrorMessage)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize id
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.id)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_LoginOK)Serialize(buffer *bytes.Buffer) error {
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
func (this *COM_ServerToClient_LoginOK)Deserialize(buffer *bytes.Buffer) error{
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
func (this *COM_ServerToClient_CreatePlayerOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //player
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize player
  {
    err := this.player.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_CreatePlayerOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize player
  if mask.ReadBit() {
    err := this.player.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.Camp!=0)
  mask.WriteBit(this.battleid!=0)
  mask.WriteBit(len(this.targetcards) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Camp
  {
    if(this.Camp!=0){
      err := prpc.Write(buffer,this.Camp)
      if err != nil{
        return err
      }
    }
  }
  // serialize battleid
  {
    if(this.battleid!=0){
      err := prpc.Write(buffer,this.battleid)
      if err != nil{
        return err
      }
    }
  }
  // serialize targetcards
  if len(this.targetcards) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.targetcards)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.targetcards {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Camp
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Camp)
    if err != nil{
      return err
    }
  }
  // deserialize battleid
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.battleid)
    if err != nil{
      return err
    }
  }
  // deserialize targetcards
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.targetcards = make([]int32,size)
    for i,_ := range this.targetcards{
      err := prpc.Read(buffer,&this.targetcards[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instId!=0)
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
  return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK)Deserialize(buffer *bytes.Buffer) error{
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
  return nil
}
func (this *COM_ServerToClient_BattleReport)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //report
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize report
  {
    err := this.report.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BattleReport)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize report
  if mask.ReadBit() {
    err := this.report.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BattleExit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //result
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize result
  {
    err := this.result.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BattleExit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize result
  if mask.ReadBit() {
    err := this.result.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_OpenChapter)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //data
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize data
  {
    err := this.data.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_OpenChapter)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize data
  if mask.ReadBit() {
    err := this.data.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SycnChapterData)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //data
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize data
  {
    err := this.data.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SycnChapterData)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize data
  if mask.ReadBit() {
    err := this.data.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_InitBagItems)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(len(this.items) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize items
  if len(this.items) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.items)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.items {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_InitBagItems)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize items
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.items = make([]COM_ItemInst,size)
    for i,_ := range this.items{
      err := this.items[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_AddBagItem)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //item
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize item
  {
    err := this.item.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_AddBagItem)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize item
  if mask.ReadBit() {
    err := this.item.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateBagItem)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //item
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize item
  {
    err := this.item.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateBagItem)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize item
  if mask.ReadBit() {
    err := this.item.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_DeleteItemOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instId!=0)
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
  return nil
}
func (this *COM_ServerToClient_DeleteItemOK)Deserialize(buffer *bytes.Buffer) error{
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
  return nil
}
func (this *COM_ServerToClient_UpdateTiantiVal)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.curVal!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize curVal
  {
    if(this.curVal!=0){
      err := prpc.Write(buffer,this.curVal)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateTiantiVal)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize curVal
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.curVal)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instid!=0)
  mask.WriteBit(this.iType!=0)
  mask.WriteBit(this.value!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := prpc.Write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  // serialize iType
  {
    if(this.iType!=0){
      err := prpc.Write(buffer,this.iType)
      if err != nil{
        return err
      }
    }
  }
  // serialize value
  {
    if(this.value!=0){
      err := prpc.Write(buffer,this.value)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  // deserialize iType
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.iType)
    if err != nil{
      return err
    }
  }
  // deserialize value
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.value)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.instid!=0)
  mask.WriteBit(this.cType!=0)
  mask.WriteBit(this.value!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := prpc.Write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  // serialize cType
  {
    if(this.cType!=0){
      err := prpc.Write(buffer,this.cType)
      if err != nil{
        return err
      }
    }
  }
  // serialize value
  {
    if(this.value!=0){
      err := prpc.Write(buffer,this.value)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  // deserialize cType
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.cType)
    if err != nil{
      return err
    }
  }
  // deserialize value
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.value)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_EquipSkillOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.skillIndex!=0)
  mask.WriteBit(this.skillID!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize skillIndex
  {
    if(this.skillIndex!=0){
      err := prpc.Write(buffer,this.skillIndex)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillID
  {
    if(this.skillID!=0){
      err := prpc.Write(buffer,this.skillID)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_EquipSkillOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize skillIndex
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.skillIndex)
    if err != nil{
      return err
    }
  }
  // deserialize skillID
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.skillID)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SkillUpdateOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.skillIndex!=0)
  mask.WriteBit(this.skillID!=0)
  mask.WriteBit(this.skillpos!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize skillIndex
  {
    if(this.skillIndex!=0){
      err := prpc.Write(buffer,this.skillIndex)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillID
  {
    if(this.skillID!=0){
      err := prpc.Write(buffer,this.skillID)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillpos
  {
    if(this.skillpos!=0){
      err := prpc.Write(buffer,this.skillpos)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_SkillUpdateOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize skillIndex
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.skillIndex)
    if err != nil{
      return err
    }
  }
  // deserialize skillID
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.skillID)
    if err != nil{
      return err
    }
  }
  // deserialize skillpos
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.skillpos)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BuyShopItemOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(len(this.items) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize items
  if len(this.items) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.items)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.items {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_BuyShopItemOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize items
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.items = make([]int32,size)
    for i,_ := range this.items{
      err := prpc.Read(buffer,&this.items[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func(this* COM_ServerToClientStub)ErrorMessage(id int ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(0))
  if err != nil{
    return err
  }
  _0 := COM_ServerToClient_ErrorMessage{}
  _0.id = id;
  err = _0.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)LoginOK(info COM_AccountInfo ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(1))
  if err != nil{
    return err
  }
  _1 := COM_ServerToClient_LoginOK{}
  _1.info = info;
  err = _1.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)CreatePlayerOK(player COM_Player ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(2))
  if err != nil{
    return err
  }
  _2 := COM_ServerToClient_CreatePlayerOK{}
  _2.player = player;
  err = _2.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)JoinBattleOk(Camp int32, battleid int32, targetcards []int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(3))
  if err != nil{
    return err
  }
  _3 := COM_ServerToClient_JoinBattleOk{}
  _3.Camp = Camp;
  _3.battleid = battleid;
  _3.targetcards = targetcards;
  err = _3.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SetupBattleOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(4))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SetBattleUnitOK(instId int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(5))
  if err != nil{
    return err
  }
  _5 := COM_ServerToClient_SetBattleUnitOK{}
  _5.instId = instId;
  err = _5.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)BattleReport(report COM_BattleReport ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(6))
  if err != nil{
    return err
  }
  _6 := COM_ServerToClient_BattleReport{}
  _6.report = report;
  err = _6.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)BattleExit(result COM_BattleResult ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(7))
  if err != nil{
    return err
  }
  _7 := COM_ServerToClient_BattleExit{}
  _7.result = result;
  err = _7.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)OpenChapter(data COM_Chapter ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(8))
  if err != nil{
    return err
  }
  _8 := COM_ServerToClient_OpenChapter{}
  _8.data = data;
  err = _8.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SycnChapterData(data COM_Chapter ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(9))
  if err != nil{
    return err
  }
  _9 := COM_ServerToClient_SycnChapterData{}
  _9.data = data;
  err = _9.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)InitBagItems(items []COM_ItemInst ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(10))
  if err != nil{
    return err
  }
  _10 := COM_ServerToClient_InitBagItems{}
  _10.items = items;
  err = _10.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)AddBagItem(item COM_ItemInst ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(11))
  if err != nil{
    return err
  }
  _11 := COM_ServerToClient_AddBagItem{}
  _11.item = item;
  err = _11.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)UpdateBagItem(item COM_ItemInst ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(12))
  if err != nil{
    return err
  }
  _12 := COM_ServerToClient_UpdateBagItem{}
  _12.item = item;
  err = _12.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)DeleteItemOK(instId int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(13))
  if err != nil{
    return err
  }
  _13 := COM_ServerToClient_DeleteItemOK{}
  _13.instId = instId;
  err = _13.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)UpdateTiantiVal(curVal int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(14))
  if err != nil{
    return err
  }
  _14 := COM_ServerToClient_UpdateTiantiVal{}
  _14.curVal = curVal;
  err = _14.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)UpdateUnitIProperty(instid int64, iType int32, value int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(15))
  if err != nil{
    return err
  }
  _15 := COM_ServerToClient_UpdateUnitIProperty{}
  _15.instid = instid;
  _15.iType = iType;
  _15.value = value;
  err = _15.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)UpdateUnitCProperty(instid int64, cType int32, value float32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(16))
  if err != nil{
    return err
  }
  _16 := COM_ServerToClient_UpdateUnitCProperty{}
  _16.instid = instid;
  _16.cType = cType;
  _16.value = value;
  err = _16.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)PromoteUnitOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(17))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RequestChapterStarRewardOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(18))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)EquipSkillOK(skillIndex int32, skillID int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(19))
  if err != nil{
    return err
  }
  _19 := COM_ServerToClient_EquipSkillOK{}
  _19.skillIndex = skillIndex;
  _19.skillID = skillID;
  err = _19.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SkillUpdateOK(skillIndex int32, skillID int32, skillpos int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(20))
  if err != nil{
    return err
  }
  _20 := COM_ServerToClient_SkillUpdateOK{}
  _20.skillIndex = skillIndex;
  _20.skillID = skillID;
  _20.skillpos = skillpos;
  err = _20.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)BuyShopItemOK(items []int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(21))
  if err != nil{
    return err
  }
  _21 := COM_ServerToClient_BuyShopItemOK{}
  _21.items = items;
  err = _21.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func Bridging_COM_ServerToClient_ErrorMessage(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _0 := COM_ServerToClient_ErrorMessage{}
  err := _0.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ErrorMessage(_0.id)
}
func Bridging_COM_ServerToClient_LoginOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _1 := COM_ServerToClient_LoginOK{}
  err := _1.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.LoginOK(_1.info)
}
func Bridging_COM_ServerToClient_CreatePlayerOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _2 := COM_ServerToClient_CreatePlayerOK{}
  err := _2.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.CreatePlayerOK(_2.player)
}
func Bridging_COM_ServerToClient_JoinBattleOk(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _3 := COM_ServerToClient_JoinBattleOk{}
  err := _3.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.JoinBattleOk(_3.Camp,_3.battleid,_3.targetcards)
}
func Bridging_COM_ServerToClient_SetupBattleOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  return p.SetupBattleOK()
}
func Bridging_COM_ServerToClient_SetBattleUnitOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _5 := COM_ServerToClient_SetBattleUnitOK{}
  err := _5.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SetBattleUnitOK(_5.instId)
}
func Bridging_COM_ServerToClient_BattleReport(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _6 := COM_ServerToClient_BattleReport{}
  err := _6.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.BattleReport(_6.report)
}
func Bridging_COM_ServerToClient_BattleExit(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _7 := COM_ServerToClient_BattleExit{}
  err := _7.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.BattleExit(_7.result)
}
func Bridging_COM_ServerToClient_OpenChapter(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _8 := COM_ServerToClient_OpenChapter{}
  err := _8.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.OpenChapter(_8.data)
}
func Bridging_COM_ServerToClient_SycnChapterData(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _9 := COM_ServerToClient_SycnChapterData{}
  err := _9.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SycnChapterData(_9.data)
}
func Bridging_COM_ServerToClient_InitBagItems(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _10 := COM_ServerToClient_InitBagItems{}
  err := _10.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.InitBagItems(_10.items)
}
func Bridging_COM_ServerToClient_AddBagItem(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _11 := COM_ServerToClient_AddBagItem{}
  err := _11.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.AddBagItem(_11.item)
}
func Bridging_COM_ServerToClient_UpdateBagItem(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _12 := COM_ServerToClient_UpdateBagItem{}
  err := _12.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.UpdateBagItem(_12.item)
}
func Bridging_COM_ServerToClient_DeleteItemOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _13 := COM_ServerToClient_DeleteItemOK{}
  err := _13.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DeleteItemOK(_13.instId)
}
func Bridging_COM_ServerToClient_UpdateTiantiVal(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _14 := COM_ServerToClient_UpdateTiantiVal{}
  err := _14.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.UpdateTiantiVal(_14.curVal)
}
func Bridging_COM_ServerToClient_UpdateUnitIProperty(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _15 := COM_ServerToClient_UpdateUnitIProperty{}
  err := _15.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.UpdateUnitIProperty(_15.instid,_15.iType,_15.value)
}
func Bridging_COM_ServerToClient_UpdateUnitCProperty(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _16 := COM_ServerToClient_UpdateUnitCProperty{}
  err := _16.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.UpdateUnitCProperty(_16.instid,_16.cType,_16.value)
}
func Bridging_COM_ServerToClient_PromoteUnitOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  return p.PromoteUnitOK()
}
func Bridging_COM_ServerToClient_RequestChapterStarRewardOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  return p.RequestChapterStarRewardOK()
}
func Bridging_COM_ServerToClient_EquipSkillOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _19 := COM_ServerToClient_EquipSkillOK{}
  err := _19.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.EquipSkillOK(_19.skillIndex,_19.skillID)
}
func Bridging_COM_ServerToClient_SkillUpdateOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _20 := COM_ServerToClient_SkillUpdateOK{}
  err := _20.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SkillUpdateOK(_20.skillIndex,_20.skillID,_20.skillpos)
}
func Bridging_COM_ServerToClient_BuyShopItemOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _21 := COM_ServerToClient_BuyShopItemOK{}
  err := _21.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.BuyShopItemOK(_21.items)
}
func COM_ServerToClientDispatch(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
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
      return Bridging_COM_ServerToClient_ErrorMessage(buffer,p);
    case 1 :
      return Bridging_COM_ServerToClient_LoginOK(buffer,p);
    case 2 :
      return Bridging_COM_ServerToClient_CreatePlayerOK(buffer,p);
    case 3 :
      return Bridging_COM_ServerToClient_JoinBattleOk(buffer,p);
    case 4 :
      return Bridging_COM_ServerToClient_SetupBattleOK(buffer,p);
    case 5 :
      return Bridging_COM_ServerToClient_SetBattleUnitOK(buffer,p);
    case 6 :
      return Bridging_COM_ServerToClient_BattleReport(buffer,p);
    case 7 :
      return Bridging_COM_ServerToClient_BattleExit(buffer,p);
    case 8 :
      return Bridging_COM_ServerToClient_OpenChapter(buffer,p);
    case 9 :
      return Bridging_COM_ServerToClient_SycnChapterData(buffer,p);
    case 10 :
      return Bridging_COM_ServerToClient_InitBagItems(buffer,p);
    case 11 :
      return Bridging_COM_ServerToClient_AddBagItem(buffer,p);
    case 12 :
      return Bridging_COM_ServerToClient_UpdateBagItem(buffer,p);
    case 13 :
      return Bridging_COM_ServerToClient_DeleteItemOK(buffer,p);
    case 14 :
      return Bridging_COM_ServerToClient_UpdateTiantiVal(buffer,p);
    case 15 :
      return Bridging_COM_ServerToClient_UpdateUnitIProperty(buffer,p);
    case 16 :
      return Bridging_COM_ServerToClient_UpdateUnitCProperty(buffer,p);
    case 17 :
      return Bridging_COM_ServerToClient_PromoteUnitOK(buffer,p);
    case 18 :
      return Bridging_COM_ServerToClient_RequestChapterStarRewardOK(buffer,p);
    case 19 :
      return Bridging_COM_ServerToClient_EquipSkillOK(buffer,p);
    case 20 :
      return Bridging_COM_ServerToClient_SkillUpdateOK(buffer,p);
    case 21 :
      return Bridging_COM_ServerToClient_BuyShopItemOK(buffer,p);
    default:
      return errors.New(prpc.NoneDispatchMatchError)
  }
}
