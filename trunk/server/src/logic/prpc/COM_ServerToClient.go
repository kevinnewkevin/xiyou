package prpc
import(
  "bytes"
  "errors"
  "suzuki/prpc"
)
type COM_ServerToClient_ErrorMessage struct{
  sync.Mutex
  id int  //0
}
type COM_ServerToClient_LoginOK struct{
  sync.Mutex
  info COM_AccountInfo  //0
}
type COM_ServerToClient_CreatePlayerOK struct{
  sync.Mutex
  player COM_Player  //0
}
type COM_ServerToClient_JoinBattleOk struct{
  sync.Mutex
  Camp int32  //0
  battleid int32  //1
  targetcards []int32  //2
  MainUnit []COM_BattleUnit  //3
}
type COM_ServerToClient_SetBattleUnitOK struct{
  sync.Mutex
  instId int64  //0
}
type COM_ServerToClient_BattleReport struct{
  sync.Mutex
  report COM_BattleReport  //0
}
type COM_ServerToClient_BattleExit struct{
  sync.Mutex
  result COM_BattleResult  //0
}
type COM_ServerToClient_OpenChapter struct{
  sync.Mutex
  data COM_Chapter  //0
}
type COM_ServerToClient_SycnChapterData struct{
  sync.Mutex
  data COM_Chapter  //0
}
type COM_ServerToClient_InitBagItems struct{
  sync.Mutex
  items []COM_ItemInst  //0
}
type COM_ServerToClient_AddBagItem struct{
  sync.Mutex
  item COM_ItemInst  //0
}
type COM_ServerToClient_UpdateBagItem struct{
  sync.Mutex
  item COM_ItemInst  //0
}
type COM_ServerToClient_DeleteItemOK struct{
  sync.Mutex
  instId int64  //0
}
type COM_ServerToClient_UpdateTiantiVal struct{
  sync.Mutex
  curVal int32  //0
}
type COM_ServerToClient_UpdateUnitIProperty struct{
  sync.Mutex
  instid int64  //0
  iType int32  //1
  value int32  //2
}
type COM_ServerToClient_UpdateUnitCProperty struct{
  sync.Mutex
  instid int64  //0
  cType int32  //1
  value float32  //2
}
type COM_ServerToClient_EquipSkillOK struct{
  sync.Mutex
  skillIndex int32  //0
  skillID int32  //1
}
type COM_ServerToClient_SkillUpdateOK struct{
  sync.Mutex
  skillIndex int32  //0
  skillID int32  //1
  skillpos int32  //2
}
type COM_ServerToClient_BuyShopItemOK struct{
  sync.Mutex
  items []COM_ItemInst  //0
}
type COM_ServerToClient_AddNewUnit struct{
  sync.Mutex
  unit COM_Unit  //0
}
type COM_ServerToClientStub struct{
  Sender prpc.StubSender
}
type COM_ServerToClientProxy interface{
  ErrorMessage(id int ) error // 0
  LoginOK(info COM_AccountInfo ) error // 1
  CreatePlayerOK(player COM_Player ) error // 2
  JoinBattleOk(Camp int32, battleid int32, targetcards []int32, MainUnit []COM_BattleUnit ) error // 3
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
  BuyShopItemOK(items []COM_ItemInst ) error // 21
  AddNewUnit(unit COM_Unit ) error // 22
}
func (this *COM_ServerToClient_ErrorMessage)Setid(value int) {
  this.Lock()
  defer this.Unlock()
  this.id = value
}
func (this *COM_ServerToClient_ErrorMessage)Getid() int {
  this.Lock()
  defer this.Unlock()
  return this.id
}
func (this *COM_ServerToClient_ErrorMessage)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_LoginOK)Setinfo(value COM_AccountInfo) {
  this.Lock()
  defer this.Unlock()
  this.info = value
}
func (this *COM_ServerToClient_LoginOK)Getinfo() COM_AccountInfo {
  this.Lock()
  defer this.Unlock()
  return this.info
}
func (this *COM_ServerToClient_LoginOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_CreatePlayerOK)Setplayer(value COM_Player) {
  this.Lock()
  defer this.Unlock()
  this.player = value
}
func (this *COM_ServerToClient_CreatePlayerOK)Getplayer() COM_Player {
  this.Lock()
  defer this.Unlock()
  return this.player
}
func (this *COM_ServerToClient_CreatePlayerOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_JoinBattleOk)SetCamp(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Camp = value
}
func (this *COM_ServerToClient_JoinBattleOk)GetCamp() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Camp
}
func (this *COM_ServerToClient_JoinBattleOk)Setbattleid(value int32) {
  this.Lock()
  defer this.Unlock()
  this.battleid = value
}
func (this *COM_ServerToClient_JoinBattleOk)Getbattleid() int32 {
  this.Lock()
  defer this.Unlock()
  return this.battleid
}
func (this *COM_ServerToClient_JoinBattleOk)Settargetcards(value []int32) {
  this.Lock()
  defer this.Unlock()
  this.targetcards = value
}
func (this *COM_ServerToClient_JoinBattleOk)Gettargetcards() []int32 {
  this.Lock()
  defer this.Unlock()
  return this.targetcards
}
func (this *COM_ServerToClient_JoinBattleOk)SetMainUnit(value []COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.MainUnit = value
}
func (this *COM_ServerToClient_JoinBattleOk)GetMainUnit() []COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.MainUnit
}
func (this *COM_ServerToClient_JoinBattleOk)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.Camp!=0)
  mask.WriteBit(this.battleid!=0)
  mask.WriteBit(len(this.targetcards) != 0)
  mask.WriteBit(len(this.MainUnit) != 0)
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
  // serialize MainUnit
  if len(this.MainUnit) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.MainUnit)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.MainUnit {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
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
  // deserialize MainUnit
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.MainUnit = make([]COM_BattleUnit,size)
    for i,_ := range this.MainUnit{
      err := this.MainUnit[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK)SetinstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.instId = value
}
func (this *COM_ServerToClient_SetBattleUnitOK)GetinstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.instId
}
func (this *COM_ServerToClient_SetBattleUnitOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_BattleReport)Setreport(value COM_BattleReport) {
  this.Lock()
  defer this.Unlock()
  this.report = value
}
func (this *COM_ServerToClient_BattleReport)Getreport() COM_BattleReport {
  this.Lock()
  defer this.Unlock()
  return this.report
}
func (this *COM_ServerToClient_BattleReport)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_BattleExit)Setresult(value COM_BattleResult) {
  this.Lock()
  defer this.Unlock()
  this.result = value
}
func (this *COM_ServerToClient_BattleExit)Getresult() COM_BattleResult {
  this.Lock()
  defer this.Unlock()
  return this.result
}
func (this *COM_ServerToClient_BattleExit)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_OpenChapter)Setdata(value COM_Chapter) {
  this.Lock()
  defer this.Unlock()
  this.data = value
}
func (this *COM_ServerToClient_OpenChapter)Getdata() COM_Chapter {
  this.Lock()
  defer this.Unlock()
  return this.data
}
func (this *COM_ServerToClient_OpenChapter)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_SycnChapterData)Setdata(value COM_Chapter) {
  this.Lock()
  defer this.Unlock()
  this.data = value
}
func (this *COM_ServerToClient_SycnChapterData)Getdata() COM_Chapter {
  this.Lock()
  defer this.Unlock()
  return this.data
}
func (this *COM_ServerToClient_SycnChapterData)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_InitBagItems)Setitems(value []COM_ItemInst) {
  this.Lock()
  defer this.Unlock()
  this.items = value
}
func (this *COM_ServerToClient_InitBagItems)Getitems() []COM_ItemInst {
  this.Lock()
  defer this.Unlock()
  return this.items
}
func (this *COM_ServerToClient_InitBagItems)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_AddBagItem)Setitem(value COM_ItemInst) {
  this.Lock()
  defer this.Unlock()
  this.item = value
}
func (this *COM_ServerToClient_AddBagItem)Getitem() COM_ItemInst {
  this.Lock()
  defer this.Unlock()
  return this.item
}
func (this *COM_ServerToClient_AddBagItem)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_UpdateBagItem)Setitem(value COM_ItemInst) {
  this.Lock()
  defer this.Unlock()
  this.item = value
}
func (this *COM_ServerToClient_UpdateBagItem)Getitem() COM_ItemInst {
  this.Lock()
  defer this.Unlock()
  return this.item
}
func (this *COM_ServerToClient_UpdateBagItem)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_DeleteItemOK)SetinstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.instId = value
}
func (this *COM_ServerToClient_DeleteItemOK)GetinstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.instId
}
func (this *COM_ServerToClient_DeleteItemOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_UpdateTiantiVal)SetcurVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.curVal = value
}
func (this *COM_ServerToClient_UpdateTiantiVal)GetcurVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.curVal
}
func (this *COM_ServerToClient_UpdateTiantiVal)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_UpdateUnitIProperty)Setinstid(value int64) {
  this.Lock()
  defer this.Unlock()
  this.instid = value
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Getinstid() int64 {
  this.Lock()
  defer this.Unlock()
  return this.instid
}
func (this *COM_ServerToClient_UpdateUnitIProperty)SetiType(value int32) {
  this.Lock()
  defer this.Unlock()
  this.iType = value
}
func (this *COM_ServerToClient_UpdateUnitIProperty)GetiType() int32 {
  this.Lock()
  defer this.Unlock()
  return this.iType
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Setvalue(value int32) {
  this.Lock()
  defer this.Unlock()
  this.value = value
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Getvalue() int32 {
  this.Lock()
  defer this.Unlock()
  return this.value
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_UpdateUnitCProperty)Setinstid(value int64) {
  this.Lock()
  defer this.Unlock()
  this.instid = value
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Getinstid() int64 {
  this.Lock()
  defer this.Unlock()
  return this.instid
}
func (this *COM_ServerToClient_UpdateUnitCProperty)SetcType(value int32) {
  this.Lock()
  defer this.Unlock()
  this.cType = value
}
func (this *COM_ServerToClient_UpdateUnitCProperty)GetcType() int32 {
  this.Lock()
  defer this.Unlock()
  return this.cType
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Setvalue(value float32) {
  this.Lock()
  defer this.Unlock()
  this.value = value
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Getvalue() float32 {
  this.Lock()
  defer this.Unlock()
  return this.value
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_EquipSkillOK)SetskillIndex(value int32) {
  this.Lock()
  defer this.Unlock()
  this.skillIndex = value
}
func (this *COM_ServerToClient_EquipSkillOK)GetskillIndex() int32 {
  this.Lock()
  defer this.Unlock()
  return this.skillIndex
}
func (this *COM_ServerToClient_EquipSkillOK)SetskillID(value int32) {
  this.Lock()
  defer this.Unlock()
  this.skillID = value
}
func (this *COM_ServerToClient_EquipSkillOK)GetskillID() int32 {
  this.Lock()
  defer this.Unlock()
  return this.skillID
}
func (this *COM_ServerToClient_EquipSkillOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_SkillUpdateOK)SetskillIndex(value int32) {
  this.Lock()
  defer this.Unlock()
  this.skillIndex = value
}
func (this *COM_ServerToClient_SkillUpdateOK)GetskillIndex() int32 {
  this.Lock()
  defer this.Unlock()
  return this.skillIndex
}
func (this *COM_ServerToClient_SkillUpdateOK)SetskillID(value int32) {
  this.Lock()
  defer this.Unlock()
  this.skillID = value
}
func (this *COM_ServerToClient_SkillUpdateOK)GetskillID() int32 {
  this.Lock()
  defer this.Unlock()
  return this.skillID
}
func (this *COM_ServerToClient_SkillUpdateOK)Setskillpos(value int32) {
  this.Lock()
  defer this.Unlock()
  this.skillpos = value
}
func (this *COM_ServerToClient_SkillUpdateOK)Getskillpos() int32 {
  this.Lock()
  defer this.Unlock()
  return this.skillpos
}
func (this *COM_ServerToClient_SkillUpdateOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_BuyShopItemOK)Setitems(value []COM_ItemInst) {
  this.Lock()
  defer this.Unlock()
  this.items = value
}
func (this *COM_ServerToClient_BuyShopItemOK)Getitems() []COM_ItemInst {
  this.Lock()
  defer this.Unlock()
  return this.items
}
func (this *COM_ServerToClient_BuyShopItemOK)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_BuyShopItemOK)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
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
func (this *COM_ServerToClient_AddNewUnit)Setunit(value COM_Unit) {
  this.Lock()
  defer this.Unlock()
  this.unit = value
}
func (this *COM_ServerToClient_AddNewUnit)Getunit() COM_Unit {
  this.Lock()
  defer this.Unlock()
  return this.unit
}
func (this *COM_ServerToClient_AddNewUnit)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //unit
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize unit
  {
    err := this.unit.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_AddNewUnit)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize unit
  if mask.ReadBit() {
    err := this.unit.Deserialize(buffer)
    if err != nil{
      return err
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
func(this* COM_ServerToClientStub)JoinBattleOk(Camp int32, battleid int32, targetcards []int32, MainUnit []COM_BattleUnit ) error {
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
  _3.MainUnit = MainUnit;
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
func(this* COM_ServerToClientStub)BuyShopItemOK(items []COM_ItemInst ) error {
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
func(this* COM_ServerToClientStub)AddNewUnit(unit COM_Unit ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  err := prpc.Write(buffer,uint16(22))
  if err != nil{
    return err
  }
  _22 := COM_ServerToClient_AddNewUnit{}
  _22.unit = unit;
  err = _22.Serialize(buffer)
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
  return p.JoinBattleOk(_3.Camp,_3.battleid,_3.targetcards,_3.MainUnit)
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
func Bridging_COM_ServerToClient_AddNewUnit(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(prpc.NoneBufferError)
  }
  if p == nil {
    return errors.New(prpc.NoneProxyError)
  }
  _22 := COM_ServerToClient_AddNewUnit{}
  err := _22.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.AddNewUnit(_22.unit)
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
    case 22 :
      return Bridging_COM_ServerToClient_AddNewUnit(buffer,p);
    default:
      return errors.New(prpc.NoneDispatchMatchError)
  }
}
