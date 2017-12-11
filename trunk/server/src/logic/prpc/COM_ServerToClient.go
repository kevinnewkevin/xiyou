package prpc
import(
  "bytes"
  "errors"
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
  MainUnit []COM_BattleUnit  //3
  battlefield string  //4
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
  items []COM_ItemInst  //0
}
type COM_ServerToClient_AddNewUnit struct{
  unit COM_Unit  //0
}
type COM_ServerToClient_SycnBlackMarkte struct{
  data COM_BlackMarket  //0
}
type COM_ServerToClient_ReceiveChat struct{
  info COM_Chat  //0
}
type COM_ServerToClient_RequestAudioOk struct{
  audioId int64  //0
  content []uint8  //1
}
type COM_ServerToClient_RecvTopList struct{
  TopList []COM_TopUnit  //0
  MyRank int32  //1
}
type COM_ServerToClient_RecvFriendTopList struct{
  TopList []COM_TopUnit  //0
  MyRank int32  //1
}
type COM_ServerToClient_SerchFriendInfo struct{
  info COM_Friend  //0
}
type COM_ServerToClient_FriendInfo struct{
  info []COM_Friend  //0
}
type COM_ServerToClient_ApplyFriend struct{
  info COM_Friend  //0
}
type COM_ServerToClient_RecvFriend struct{
  info COM_Friend  //0
}
type COM_ServerToClient_DelFriend struct{
  instid int64  //0
}
type COM_ServerToClient_RecvEnemy struct{
  info COM_Friend  //0
}
type COM_ServerToClient_DelEnemy struct{
  instid int64  //0
}
type COM_ServerToClient_QueryPlayerInfoOK struct{
  Info COM_PlayerInfo  //0
}
type COM_ServerToClient_UpdateGuildAssistant struct{
  info COM_Assistant  //0
  donator string  //1
}
type COM_ServerToClient_SycnGuildAssistant struct{
  infos []COM_Assistant  //0
}
type COM_ServerToClient_LeaveGuildOk struct{
  who string  //0
  isKick bool  //1
}
type COM_ServerToClient_InitGuildData struct{
  guild COM_Guild  //0
}
type COM_ServerToClient_InitGuildMemberList struct{
  members []COM_GuildMember  //0
}
type COM_ServerToClient_ModifyGuildMemberList struct{
  member COM_GuildMember  //0
  flag int  //1
}
type COM_ServerToClient_QueryGuildListResult struct{
  guildList []COM_GuildViewerData  //0
}
type COM_ServerToClient_QueryGuildDetailsResult struct{
  info COM_GuildDetails  //0
}
type COM_ServerToClient_AppendMail struct{
  mails []COM_Mail  //0
}
type COM_ServerToClient_DelMailOK struct{
  mailId int32  //0
}
type COM_ServerToClient_UpdateMailOk struct{
  mail COM_Mail  //0
}
type COM_ServerToClient_QueryBattleRecordOK struct{
  record COM_BattleRecord  //0
}
type COM_ServerToClient_QueryRecordDetailOK struct{
  recordDetials []COM_BattleRecord_Detail  //0
}
type COM_ServerToClient_JoinBattleOk_back struct{
  round int32  //0
  state int32  //1
  second int32  //2
  battlesnape COM_BattleSnape  //3
}
type COM_ServerToClientStub struct{
  Sender StubSender
}
type COM_ServerToClientProxy interface{
  ErrorMessage(id int ) error // 0
  LoginOK(info COM_AccountInfo ) error // 1
  CreatePlayerOK(player COM_Player ) error // 2
  JoinBattleOk(Camp int32, battleid int32, targetcards []int32, MainUnit []COM_BattleUnit, battlefield string ) error // 3
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
  SycnBlackMarkte(data COM_BlackMarket ) error // 23
  ReceiveChat(info COM_Chat ) error // 24
  RequestAudioOk(audioId int64, content []uint8 ) error // 25
  RecvTopList(TopList []COM_TopUnit, MyRank int32 ) error // 26
  RecvFriendTopList(TopList []COM_TopUnit, MyRank int32 ) error // 27
  SerchFriendInfo(info COM_Friend ) error // 28
  FriendInfo(info []COM_Friend ) error // 29
  ApplyFriend(info COM_Friend ) error // 30
  RecvFriend(info COM_Friend ) error // 31
  DelFriend(instid int64 ) error // 32
  RecvEnemy(info COM_Friend ) error // 33
  DelEnemy(instid int64 ) error // 34
  QueryPlayerInfoOK(Info COM_PlayerInfo ) error // 35
  UpdateGuildAssistant(info COM_Assistant, donator string ) error // 36
  SycnGuildAssistant(infos []COM_Assistant ) error // 37
  CreateGuildOK() error // 38
  DelGuildOK() error // 39
  LeaveGuildOk(who string, isKick bool ) error // 40
  InitGuildData(guild COM_Guild ) error // 41
  InitGuildMemberList(members []COM_GuildMember ) error // 42
  ModifyGuildMemberList(member COM_GuildMember, flag int ) error // 43
  QueryGuildListResult(guildList []COM_GuildViewerData ) error // 44
  QueryGuildDetailsResult(info COM_GuildDetails ) error // 45
  JoinGuildOk() error // 46
  AppendMail(mails []COM_Mail ) error // 47
  DelMailOK(mailId int32 ) error // 48
  UpdateMailOk(mail COM_Mail ) error // 49
  QueryBattleRecordOK(record COM_BattleRecord ) error // 50
  QueryRecordDetailOK(recordDetials []COM_BattleRecord_Detail ) error // 51
  JoinBattleOk_back(round int32, state int32, second int32, battlesnape COM_BattleSnape ) error // 52
}
func (this *COM_ServerToClient_ErrorMessage)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.id!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize id
  {
    if(this.id!=0){
      err := write(buffer,this.id)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_ErrorMessage)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize id
  if mask.readBit() {
    err := read(buffer,&this.id)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_LoginOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_CreatePlayerOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //player
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize player
  if mask.readBit() {
    err := this.player.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Camp!=0)
  mask.writeBit(this.battleid!=0)
  mask.writeBit(len(this.targetcards) != 0)
  mask.writeBit(len(this.MainUnit) != 0)
  mask.writeBit(len(this.battlefield) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Camp
  {
    if(this.Camp!=0){
      err := write(buffer,this.Camp)
      if err != nil{
        return err
      }
    }
  }
  // serialize battleid
  {
    if(this.battleid!=0){
      err := write(buffer,this.battleid)
      if err != nil{
        return err
      }
    }
  }
  // serialize targetcards
  if len(this.targetcards) != 0{
    {
      err := write(buffer,uint(len(this.targetcards)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.targetcards {
      err := write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  // serialize MainUnit
  if len(this.MainUnit) != 0{
    {
      err := write(buffer,uint(len(this.MainUnit)))
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
  // serialize battlefield
  if len(this.battlefield) != 0{
    err := write(buffer,this.battlefield)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Camp
  if mask.readBit() {
    err := read(buffer,&this.Camp)
    if err != nil{
      return err
    }
  }
  // deserialize battleid
  if mask.readBit() {
    err := read(buffer,&this.battleid)
    if err != nil{
      return err
    }
  }
  // deserialize targetcards
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.targetcards = make([]int32,size)
    for i,_ := range this.targetcards{
      err := read(buffer,&this.targetcards[i])
      if err != nil{
        return err
      }
    }
  }
  // deserialize MainUnit
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
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
  // deserialize battlefield
  if mask.readBit() {
    err := read(buffer,&this.battlefield)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BattleReport)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //report
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize report
  if mask.readBit() {
    err := this.report.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BattleExit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //result
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize result
  if mask.readBit() {
    err := this.result.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_OpenChapter)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //data
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize data
  if mask.readBit() {
    err := this.data.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SycnChapterData)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //data
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize data
  if mask.readBit() {
    err := this.data.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_InitBagItems)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.items) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize items
  if len(this.items) != 0{
    {
      err := write(buffer,uint(len(this.items)))
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize items
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
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
  mask := newMask1(1)
  mask.writeBit(true) //item
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize item
  if mask.readBit() {
    err := this.item.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateBagItem)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //item
  {
    err := write(buffer,mask.bytes())
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
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize item
  if mask.readBit() {
    err := this.item.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_DeleteItemOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_DeleteItemOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateTiantiVal)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.curVal!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize curVal
  {
    if(this.curVal!=0){
      err := write(buffer,this.curVal)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateTiantiVal)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize curVal
  if mask.readBit() {
    err := read(buffer,&this.curVal)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instid!=0)
  mask.writeBit(this.iType!=0)
  mask.writeBit(this.value!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  // serialize iType
  {
    if(this.iType!=0){
      err := write(buffer,this.iType)
      if err != nil{
        return err
      }
    }
  }
  // serialize value
  {
    if(this.value!=0){
      err := write(buffer,this.value)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitIProperty)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.readBit() {
    err := read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  // deserialize iType
  if mask.readBit() {
    err := read(buffer,&this.iType)
    if err != nil{
      return err
    }
  }
  // deserialize value
  if mask.readBit() {
    err := read(buffer,&this.value)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instid!=0)
  mask.writeBit(this.cType!=0)
  mask.writeBit(this.value!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  // serialize cType
  {
    if(this.cType!=0){
      err := write(buffer,this.cType)
      if err != nil{
        return err
      }
    }
  }
  // serialize value
  {
    if(this.value!=0){
      err := write(buffer,this.value)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateUnitCProperty)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.readBit() {
    err := read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  // deserialize cType
  if mask.readBit() {
    err := read(buffer,&this.cType)
    if err != nil{
      return err
    }
  }
  // deserialize value
  if mask.readBit() {
    err := read(buffer,&this.value)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_EquipSkillOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.skillIndex!=0)
  mask.writeBit(this.skillID!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize skillIndex
  {
    if(this.skillIndex!=0){
      err := write(buffer,this.skillIndex)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillID
  {
    if(this.skillID!=0){
      err := write(buffer,this.skillID)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_EquipSkillOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize skillIndex
  if mask.readBit() {
    err := read(buffer,&this.skillIndex)
    if err != nil{
      return err
    }
  }
  // deserialize skillID
  if mask.readBit() {
    err := read(buffer,&this.skillID)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SkillUpdateOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.skillIndex!=0)
  mask.writeBit(this.skillID!=0)
  mask.writeBit(this.skillpos!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize skillIndex
  {
    if(this.skillIndex!=0){
      err := write(buffer,this.skillIndex)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillID
  {
    if(this.skillID!=0){
      err := write(buffer,this.skillID)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillpos
  {
    if(this.skillpos!=0){
      err := write(buffer,this.skillpos)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_SkillUpdateOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize skillIndex
  if mask.readBit() {
    err := read(buffer,&this.skillIndex)
    if err != nil{
      return err
    }
  }
  // deserialize skillID
  if mask.readBit() {
    err := read(buffer,&this.skillID)
    if err != nil{
      return err
    }
  }
  // deserialize skillpos
  if mask.readBit() {
    err := read(buffer,&this.skillpos)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_BuyShopItemOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.items) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize items
  if len(this.items) != 0{
    {
      err := write(buffer,uint(len(this.items)))
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
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize items
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
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
func (this *COM_ServerToClient_AddNewUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //unit
  {
    err := write(buffer,mask.bytes())
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
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize unit
  if mask.readBit() {
    err := this.unit.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SycnBlackMarkte)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //data
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_SycnBlackMarkte)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize data
  if mask.readBit() {
    err := this.data.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_ReceiveChat)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_ReceiveChat)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_RequestAudioOk)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.audioId!=0)
  mask.writeBit(len(this.content) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize audioId
  {
    if(this.audioId!=0){
      err := write(buffer,this.audioId)
      if err != nil{
        return err
      }
    }
  }
  // serialize content
  if len(this.content) != 0{
    {
      err := write(buffer,uint(len(this.content)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.content {
      err := write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_RequestAudioOk)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize audioId
  if mask.readBit() {
    err := read(buffer,&this.audioId)
    if err != nil{
      return err
    }
  }
  // deserialize content
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.content = make([]uint8,size)
    for i,_ := range this.content{
      err := read(buffer,&this.content[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_RecvTopList)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.TopList) != 0)
  mask.writeBit(this.MyRank!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize TopList
  if len(this.TopList) != 0{
    {
      err := write(buffer,uint(len(this.TopList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.TopList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize MyRank
  {
    if(this.MyRank!=0){
      err := write(buffer,this.MyRank)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_RecvTopList)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize TopList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.TopList = make([]COM_TopUnit,size)
    for i,_ := range this.TopList{
      err := this.TopList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize MyRank
  if mask.readBit() {
    err := read(buffer,&this.MyRank)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_RecvFriendTopList)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.TopList) != 0)
  mask.writeBit(this.MyRank!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize TopList
  if len(this.TopList) != 0{
    {
      err := write(buffer,uint(len(this.TopList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.TopList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize MyRank
  {
    if(this.MyRank!=0){
      err := write(buffer,this.MyRank)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_RecvFriendTopList)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize TopList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.TopList = make([]COM_TopUnit,size)
    for i,_ := range this.TopList{
      err := this.TopList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize MyRank
  if mask.readBit() {
    err := read(buffer,&this.MyRank)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SerchFriendInfo)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_SerchFriendInfo)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_FriendInfo)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.info) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize info
  if len(this.info) != 0{
    {
      err := write(buffer,uint(len(this.info)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.info {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_FriendInfo)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.info = make([]COM_Friend,size)
    for i,_ := range this.info{
      err := this.info[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_ApplyFriend)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_ApplyFriend)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_RecvFriend)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_RecvFriend)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_DelFriend)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instid!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_DelFriend)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.readBit() {
    err := read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_RecvEnemy)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_RecvEnemy)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_DelEnemy)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instid!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_DelEnemy)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.readBit() {
    err := read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryPlayerInfoOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //Info
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Info
  {
    err := this.Info.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryPlayerInfoOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Info
  if mask.readBit() {
    err := this.Info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateGuildAssistant)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  mask.writeBit(len(this.donator) != 0)
  {
    err := write(buffer,mask.bytes())
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
  // serialize donator
  if len(this.donator) != 0{
    err := write(buffer,this.donator)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateGuildAssistant)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize donator
  if mask.readBit() {
    err := read(buffer,&this.donator)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_SycnGuildAssistant)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.infos) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize infos
  if len(this.infos) != 0{
    {
      err := write(buffer,uint(len(this.infos)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.infos {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_SycnGuildAssistant)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize infos
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.infos = make([]COM_Assistant,size)
    for i,_ := range this.infos{
      err := this.infos[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_LeaveGuildOk)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.who) != 0)
  mask.writeBit(this.isKick)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize who
  if len(this.who) != 0{
    err := write(buffer,this.who)
    if err != nil {
      return err
    }
  }
  // serialize isKick
  {
  }
  return nil
}
func (this *COM_ServerToClient_LeaveGuildOk)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize who
  if mask.readBit() {
    err := read(buffer,&this.who)
    if err != nil{
      return err
    }
  }
  // deserialize isKick
  this.isKick = mask.readBit();
  return nil
}
func (this *COM_ServerToClient_InitGuildData)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //guild
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize guild
  {
    err := this.guild.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_InitGuildData)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize guild
  if mask.readBit() {
    err := this.guild.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_InitGuildMemberList)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.members) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize members
  if len(this.members) != 0{
    {
      err := write(buffer,uint(len(this.members)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.members {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_InitGuildMemberList)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize members
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.members = make([]COM_GuildMember,size)
    for i,_ := range this.members{
      err := this.members[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_ModifyGuildMemberList)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //member
  mask.writeBit(this.flag!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize member
  {
    err := this.member.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize flag
  {
    if(this.flag!=0){
      err := write(buffer,this.flag)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_ModifyGuildMemberList)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize member
  if mask.readBit() {
    err := this.member.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize flag
  if mask.readBit() {
    err := read(buffer,&this.flag)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryGuildListResult)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.guildList) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize guildList
  if len(this.guildList) != 0{
    {
      err := write(buffer,uint(len(this.guildList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.guildList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryGuildListResult)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize guildList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.guildList = make([]COM_GuildViewerData,size)
    for i,_ := range this.guildList{
      err := this.guildList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryGuildDetailsResult)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
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
func (this *COM_ServerToClient_QueryGuildDetailsResult)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_AppendMail)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.mails) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize mails
  if len(this.mails) != 0{
    {
      err := write(buffer,uint(len(this.mails)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.mails {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_AppendMail)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize mails
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.mails = make([]COM_Mail,size)
    for i,_ := range this.mails{
      err := this.mails[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_DelMailOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.mailId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize mailId
  {
    if(this.mailId!=0){
      err := write(buffer,this.mailId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_DelMailOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize mailId
  if mask.readBit() {
    err := read(buffer,&this.mailId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateMailOk)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //mail
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize mail
  {
    err := this.mail.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_UpdateMailOk)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize mail
  if mask.readBit() {
    err := this.mail.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryBattleRecordOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //record
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize record
  {
    err := this.record.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryBattleRecordOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize record
  if mask.readBit() {
    err := this.record.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryRecordDetailOK)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.recordDetials) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize recordDetials
  if len(this.recordDetials) != 0{
    {
      err := write(buffer,uint(len(this.recordDetials)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.recordDetials {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_QueryRecordDetailOK)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize recordDetials
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.recordDetials = make([]COM_BattleRecord_Detail,size)
    for i,_ := range this.recordDetials{
      err := this.recordDetials[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk_back)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.round!=0)
  mask.writeBit(this.state!=0)
  mask.writeBit(this.second!=0)
  mask.writeBit(true) //battlesnape
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize round
  {
    if(this.round!=0){
      err := write(buffer,this.round)
      if err != nil{
        return err
      }
    }
  }
  // serialize state
  {
    if(this.state!=0){
      err := write(buffer,this.state)
      if err != nil{
        return err
      }
    }
  }
  // serialize second
  {
    if(this.second!=0){
      err := write(buffer,this.second)
      if err != nil{
        return err
      }
    }
  }
  // serialize battlesnape
  {
    err := this.battlesnape.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ServerToClient_JoinBattleOk_back)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize round
  if mask.readBit() {
    err := read(buffer,&this.round)
    if err != nil{
      return err
    }
  }
  // deserialize state
  if mask.readBit() {
    err := read(buffer,&this.state)
    if err != nil{
      return err
    }
  }
  // deserialize second
  if mask.readBit() {
    err := read(buffer,&this.second)
    if err != nil{
      return err
    }
  }
  // deserialize battlesnape
  if mask.readBit() {
    err := this.battlesnape.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func(this* COM_ServerToClientStub)ErrorMessage(id int ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(0))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(1))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(2))
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
func(this* COM_ServerToClientStub)JoinBattleOk(Camp int32, battleid int32, targetcards []int32, MainUnit []COM_BattleUnit, battlefield string ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(3))
  if err != nil{
    return err
  }
  _3 := COM_ServerToClient_JoinBattleOk{}
  _3.Camp = Camp;
  _3.battleid = battleid;
  _3.targetcards = targetcards;
  _3.MainUnit = MainUnit;
  _3.battlefield = battlefield;
  err = _3.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SetupBattleOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(4))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SetBattleUnitOK(instId int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(5))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(6))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(7))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(8))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(9))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(10))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(11))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(12))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(13))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(14))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(15))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(16))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(17))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RequestChapterStarRewardOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(18))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)EquipSkillOK(skillIndex int32, skillID int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(19))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(20))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(21))
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
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(22))
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
func(this* COM_ServerToClientStub)SycnBlackMarkte(data COM_BlackMarket ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(23))
  if err != nil{
    return err
  }
  _23 := COM_ServerToClient_SycnBlackMarkte{}
  _23.data = data;
  err = _23.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)ReceiveChat(info COM_Chat ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(24))
  if err != nil{
    return err
  }
  _24 := COM_ServerToClient_ReceiveChat{}
  _24.info = info;
  err = _24.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RequestAudioOk(audioId int64, content []uint8 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(25))
  if err != nil{
    return err
  }
  _25 := COM_ServerToClient_RequestAudioOk{}
  _25.audioId = audioId;
  _25.content = content;
  err = _25.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RecvTopList(TopList []COM_TopUnit, MyRank int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(26))
  if err != nil{
    return err
  }
  _26 := COM_ServerToClient_RecvTopList{}
  _26.TopList = TopList;
  _26.MyRank = MyRank;
  err = _26.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RecvFriendTopList(TopList []COM_TopUnit, MyRank int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(27))
  if err != nil{
    return err
  }
  _27 := COM_ServerToClient_RecvFriendTopList{}
  _27.TopList = TopList;
  _27.MyRank = MyRank;
  err = _27.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SerchFriendInfo(info COM_Friend ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(28))
  if err != nil{
    return err
  }
  _28 := COM_ServerToClient_SerchFriendInfo{}
  _28.info = info;
  err = _28.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)FriendInfo(info []COM_Friend ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(29))
  if err != nil{
    return err
  }
  _29 := COM_ServerToClient_FriendInfo{}
  _29.info = info;
  err = _29.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)ApplyFriend(info COM_Friend ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(30))
  if err != nil{
    return err
  }
  _30 := COM_ServerToClient_ApplyFriend{}
  _30.info = info;
  err = _30.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RecvFriend(info COM_Friend ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(31))
  if err != nil{
    return err
  }
  _31 := COM_ServerToClient_RecvFriend{}
  _31.info = info;
  err = _31.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)DelFriend(instid int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(32))
  if err != nil{
    return err
  }
  _32 := COM_ServerToClient_DelFriend{}
  _32.instid = instid;
  err = _32.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)RecvEnemy(info COM_Friend ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(33))
  if err != nil{
    return err
  }
  _33 := COM_ServerToClient_RecvEnemy{}
  _33.info = info;
  err = _33.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)DelEnemy(instid int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(34))
  if err != nil{
    return err
  }
  _34 := COM_ServerToClient_DelEnemy{}
  _34.instid = instid;
  err = _34.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)QueryPlayerInfoOK(Info COM_PlayerInfo ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(35))
  if err != nil{
    return err
  }
  _35 := COM_ServerToClient_QueryPlayerInfoOK{}
  _35.Info = Info;
  err = _35.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)UpdateGuildAssistant(info COM_Assistant, donator string ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(36))
  if err != nil{
    return err
  }
  _36 := COM_ServerToClient_UpdateGuildAssistant{}
  _36.info = info;
  _36.donator = donator;
  err = _36.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)SycnGuildAssistant(infos []COM_Assistant ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(37))
  if err != nil{
    return err
  }
  _37 := COM_ServerToClient_SycnGuildAssistant{}
  _37.infos = infos;
  err = _37.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)CreateGuildOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(38))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)DelGuildOK() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(39))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)LeaveGuildOk(who string, isKick bool ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(40))
  if err != nil{
    return err
  }
  _40 := COM_ServerToClient_LeaveGuildOk{}
  _40.who = who;
  _40.isKick = isKick;
  err = _40.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)InitGuildData(guild COM_Guild ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(41))
  if err != nil{
    return err
  }
  _41 := COM_ServerToClient_InitGuildData{}
  _41.guild = guild;
  err = _41.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)InitGuildMemberList(members []COM_GuildMember ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(42))
  if err != nil{
    return err
  }
  _42 := COM_ServerToClient_InitGuildMemberList{}
  _42.members = members;
  err = _42.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)ModifyGuildMemberList(member COM_GuildMember, flag int ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(43))
  if err != nil{
    return err
  }
  _43 := COM_ServerToClient_ModifyGuildMemberList{}
  _43.member = member;
  _43.flag = flag;
  err = _43.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)QueryGuildListResult(guildList []COM_GuildViewerData ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(44))
  if err != nil{
    return err
  }
  _44 := COM_ServerToClient_QueryGuildListResult{}
  _44.guildList = guildList;
  err = _44.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)QueryGuildDetailsResult(info COM_GuildDetails ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(45))
  if err != nil{
    return err
  }
  _45 := COM_ServerToClient_QueryGuildDetailsResult{}
  _45.info = info;
  err = _45.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)JoinGuildOk() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(46))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)AppendMail(mails []COM_Mail ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(47))
  if err != nil{
    return err
  }
  _47 := COM_ServerToClient_AppendMail{}
  _47.mails = mails;
  err = _47.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)DelMailOK(mailId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(48))
  if err != nil{
    return err
  }
  _48 := COM_ServerToClient_DelMailOK{}
  _48.mailId = mailId;
  err = _48.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)UpdateMailOk(mail COM_Mail ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(49))
  if err != nil{
    return err
  }
  _49 := COM_ServerToClient_UpdateMailOk{}
  _49.mail = mail;
  err = _49.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)QueryBattleRecordOK(record COM_BattleRecord ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(50))
  if err != nil{
    return err
  }
  _50 := COM_ServerToClient_QueryBattleRecordOK{}
  _50.record = record;
  err = _50.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)QueryRecordDetailOK(recordDetials []COM_BattleRecord_Detail ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(51))
  if err != nil{
    return err
  }
  _51 := COM_ServerToClient_QueryRecordDetailOK{}
  _51.recordDetials = recordDetials;
  err = _51.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ServerToClientStub)JoinBattleOk_back(round int32, state int32, second int32, battlesnape COM_BattleSnape ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(52))
  if err != nil{
    return err
  }
  _52 := COM_ServerToClient_JoinBattleOk_back{}
  _52.round = round;
  _52.state = state;
  _52.second = second;
  _52.battlesnape = battlesnape;
  err = _52.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func Bridging_COM_ServerToClient_ErrorMessage(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _3 := COM_ServerToClient_JoinBattleOk{}
  err := _3.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.JoinBattleOk(_3.Camp,_3.battleid,_3.targetcards,_3.MainUnit,_3.battlefield)
}
func Bridging_COM_ServerToClient_SetupBattleOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.SetupBattleOK()
}
func Bridging_COM_ServerToClient_SetBattleUnitOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.PromoteUnitOK()
}
func Bridging_COM_ServerToClient_RequestChapterStarRewardOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.RequestChapterStarRewardOK()
}
func Bridging_COM_ServerToClient_EquipSkillOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
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
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _22 := COM_ServerToClient_AddNewUnit{}
  err := _22.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.AddNewUnit(_22.unit)
}
func Bridging_COM_ServerToClient_SycnBlackMarkte(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _23 := COM_ServerToClient_SycnBlackMarkte{}
  err := _23.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SycnBlackMarkte(_23.data)
}
func Bridging_COM_ServerToClient_ReceiveChat(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _24 := COM_ServerToClient_ReceiveChat{}
  err := _24.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ReceiveChat(_24.info)
}
func Bridging_COM_ServerToClient_RequestAudioOk(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _25 := COM_ServerToClient_RequestAudioOk{}
  err := _25.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RequestAudioOk(_25.audioId,_25.content)
}
func Bridging_COM_ServerToClient_RecvTopList(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _26 := COM_ServerToClient_RecvTopList{}
  err := _26.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RecvTopList(_26.TopList,_26.MyRank)
}
func Bridging_COM_ServerToClient_RecvFriendTopList(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _27 := COM_ServerToClient_RecvFriendTopList{}
  err := _27.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RecvFriendTopList(_27.TopList,_27.MyRank)
}
func Bridging_COM_ServerToClient_SerchFriendInfo(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _28 := COM_ServerToClient_SerchFriendInfo{}
  err := _28.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SerchFriendInfo(_28.info)
}
func Bridging_COM_ServerToClient_FriendInfo(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _29 := COM_ServerToClient_FriendInfo{}
  err := _29.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.FriendInfo(_29.info)
}
func Bridging_COM_ServerToClient_ApplyFriend(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _30 := COM_ServerToClient_ApplyFriend{}
  err := _30.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ApplyFriend(_30.info)
}
func Bridging_COM_ServerToClient_RecvFriend(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _31 := COM_ServerToClient_RecvFriend{}
  err := _31.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RecvFriend(_31.info)
}
func Bridging_COM_ServerToClient_DelFriend(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _32 := COM_ServerToClient_DelFriend{}
  err := _32.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DelFriend(_32.instid)
}
func Bridging_COM_ServerToClient_RecvEnemy(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _33 := COM_ServerToClient_RecvEnemy{}
  err := _33.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RecvEnemy(_33.info)
}
func Bridging_COM_ServerToClient_DelEnemy(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _34 := COM_ServerToClient_DelEnemy{}
  err := _34.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DelEnemy(_34.instid)
}
func Bridging_COM_ServerToClient_QueryPlayerInfoOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _35 := COM_ServerToClient_QueryPlayerInfoOK{}
  err := _35.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.QueryPlayerInfoOK(_35.Info)
}
func Bridging_COM_ServerToClient_UpdateGuildAssistant(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _36 := COM_ServerToClient_UpdateGuildAssistant{}
  err := _36.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.UpdateGuildAssistant(_36.info,_36.donator)
}
func Bridging_COM_ServerToClient_SycnGuildAssistant(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _37 := COM_ServerToClient_SycnGuildAssistant{}
  err := _37.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SycnGuildAssistant(_37.infos)
}
func Bridging_COM_ServerToClient_CreateGuildOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.CreateGuildOK()
}
func Bridging_COM_ServerToClient_DelGuildOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.DelGuildOK()
}
func Bridging_COM_ServerToClient_LeaveGuildOk(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _40 := COM_ServerToClient_LeaveGuildOk{}
  err := _40.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.LeaveGuildOk(_40.who,_40.isKick)
}
func Bridging_COM_ServerToClient_InitGuildData(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _41 := COM_ServerToClient_InitGuildData{}
  err := _41.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.InitGuildData(_41.guild)
}
func Bridging_COM_ServerToClient_InitGuildMemberList(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _42 := COM_ServerToClient_InitGuildMemberList{}
  err := _42.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.InitGuildMemberList(_42.members)
}
func Bridging_COM_ServerToClient_ModifyGuildMemberList(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _43 := COM_ServerToClient_ModifyGuildMemberList{}
  err := _43.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ModifyGuildMemberList(_43.member,_43.flag)
}
func Bridging_COM_ServerToClient_QueryGuildListResult(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _44 := COM_ServerToClient_QueryGuildListResult{}
  err := _44.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.QueryGuildListResult(_44.guildList)
}
func Bridging_COM_ServerToClient_QueryGuildDetailsResult(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _45 := COM_ServerToClient_QueryGuildDetailsResult{}
  err := _45.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.QueryGuildDetailsResult(_45.info)
}
func Bridging_COM_ServerToClient_JoinGuildOk(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.JoinGuildOk()
}
func Bridging_COM_ServerToClient_AppendMail(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _47 := COM_ServerToClient_AppendMail{}
  err := _47.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.AppendMail(_47.mails)
}
func Bridging_COM_ServerToClient_DelMailOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _48 := COM_ServerToClient_DelMailOK{}
  err := _48.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DelMailOK(_48.mailId)
}
func Bridging_COM_ServerToClient_UpdateMailOk(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _49 := COM_ServerToClient_UpdateMailOk{}
  err := _49.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.UpdateMailOk(_49.mail)
}
func Bridging_COM_ServerToClient_QueryBattleRecordOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _50 := COM_ServerToClient_QueryBattleRecordOK{}
  err := _50.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.QueryBattleRecordOK(_50.record)
}
func Bridging_COM_ServerToClient_QueryRecordDetailOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _51 := COM_ServerToClient_QueryRecordDetailOK{}
  err := _51.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.QueryRecordDetailOK(_51.recordDetials)
}
func Bridging_COM_ServerToClient_JoinBattleOk_back(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _52 := COM_ServerToClient_JoinBattleOk_back{}
  err := _52.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.JoinBattleOk_back(_52.round,_52.state,_52.second,_52.battlesnape)
}
func COM_ServerToClientDispatch(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
  if buffer == nil {
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  pid := uint16(0XFFFF)
  err := read(buffer,&pid)
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
    case 23 :
      return Bridging_COM_ServerToClient_SycnBlackMarkte(buffer,p);
    case 24 :
      return Bridging_COM_ServerToClient_ReceiveChat(buffer,p);
    case 25 :
      return Bridging_COM_ServerToClient_RequestAudioOk(buffer,p);
    case 26 :
      return Bridging_COM_ServerToClient_RecvTopList(buffer,p);
    case 27 :
      return Bridging_COM_ServerToClient_RecvFriendTopList(buffer,p);
    case 28 :
      return Bridging_COM_ServerToClient_SerchFriendInfo(buffer,p);
    case 29 :
      return Bridging_COM_ServerToClient_FriendInfo(buffer,p);
    case 30 :
      return Bridging_COM_ServerToClient_ApplyFriend(buffer,p);
    case 31 :
      return Bridging_COM_ServerToClient_RecvFriend(buffer,p);
    case 32 :
      return Bridging_COM_ServerToClient_DelFriend(buffer,p);
    case 33 :
      return Bridging_COM_ServerToClient_RecvEnemy(buffer,p);
    case 34 :
      return Bridging_COM_ServerToClient_DelEnemy(buffer,p);
    case 35 :
      return Bridging_COM_ServerToClient_QueryPlayerInfoOK(buffer,p);
    case 36 :
      return Bridging_COM_ServerToClient_UpdateGuildAssistant(buffer,p);
    case 37 :
      return Bridging_COM_ServerToClient_SycnGuildAssistant(buffer,p);
    case 38 :
      return Bridging_COM_ServerToClient_CreateGuildOK(buffer,p);
    case 39 :
      return Bridging_COM_ServerToClient_DelGuildOK(buffer,p);
    case 40 :
      return Bridging_COM_ServerToClient_LeaveGuildOk(buffer,p);
    case 41 :
      return Bridging_COM_ServerToClient_InitGuildData(buffer,p);
    case 42 :
      return Bridging_COM_ServerToClient_InitGuildMemberList(buffer,p);
    case 43 :
      return Bridging_COM_ServerToClient_ModifyGuildMemberList(buffer,p);
    case 44 :
      return Bridging_COM_ServerToClient_QueryGuildListResult(buffer,p);
    case 45 :
      return Bridging_COM_ServerToClient_QueryGuildDetailsResult(buffer,p);
    case 46 :
      return Bridging_COM_ServerToClient_JoinGuildOk(buffer,p);
    case 47 :
      return Bridging_COM_ServerToClient_AppendMail(buffer,p);
    case 48 :
      return Bridging_COM_ServerToClient_DelMailOK(buffer,p);
    case 49 :
      return Bridging_COM_ServerToClient_UpdateMailOk(buffer,p);
    case 50 :
      return Bridging_COM_ServerToClient_QueryBattleRecordOK(buffer,p);
    case 51 :
      return Bridging_COM_ServerToClient_QueryRecordDetailOK(buffer,p);
    case 52 :
      return Bridging_COM_ServerToClient_JoinBattleOk_back(buffer,p);
    default:
      return errors.New(NoneDispatchMatchError)
  }
}
