package prpc
import(
  "bytes"
  "encoding/json"
)
type SGE_DBPlayer struct{
  COM_Player
  PlayerId int64  //0
  Username string  //1
  LoginTime int64  //2
  LogoutTime int64  //3
  BattleGroupIdx int32  //4
  GenItemMaxGuid int64  //5
  AssistantId int32  //6
  BagItemList []COM_ItemInst  //7
  BlackMarketData COM_BlackMarket  //8
  ChapterPondId int32  //9
}
func (this *SGE_DBPlayer)SetPlayerId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.PlayerId = value
}
func (this *SGE_DBPlayer)GetPlayerId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.PlayerId
}
func (this *SGE_DBPlayer)SetUsername(value string) {
  this.Lock()
  defer this.Unlock()
  this.Username = value
}
func (this *SGE_DBPlayer)GetUsername() string {
  this.Lock()
  defer this.Unlock()
  return this.Username
}
func (this *SGE_DBPlayer)SetLoginTime(value int64) {
  this.Lock()
  defer this.Unlock()
  this.LoginTime = value
}
func (this *SGE_DBPlayer)GetLoginTime() int64 {
  this.Lock()
  defer this.Unlock()
  return this.LoginTime
}
func (this *SGE_DBPlayer)SetLogoutTime(value int64) {
  this.Lock()
  defer this.Unlock()
  this.LogoutTime = value
}
func (this *SGE_DBPlayer)GetLogoutTime() int64 {
  this.Lock()
  defer this.Unlock()
  return this.LogoutTime
}
func (this *SGE_DBPlayer)SetBattleGroupIdx(value int32) {
  this.Lock()
  defer this.Unlock()
  this.BattleGroupIdx = value
}
func (this *SGE_DBPlayer)GetBattleGroupIdx() int32 {
  this.Lock()
  defer this.Unlock()
  return this.BattleGroupIdx
}
func (this *SGE_DBPlayer)SetGenItemMaxGuid(value int64) {
  this.Lock()
  defer this.Unlock()
  this.GenItemMaxGuid = value
}
func (this *SGE_DBPlayer)GetGenItemMaxGuid() int64 {
  this.Lock()
  defer this.Unlock()
  return this.GenItemMaxGuid
}
func (this *SGE_DBPlayer)SetAssistantId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.AssistantId = value
}
func (this *SGE_DBPlayer)GetAssistantId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.AssistantId
}
func (this *SGE_DBPlayer)SetBagItemList(value []COM_ItemInst) {
  this.Lock()
  defer this.Unlock()
  this.BagItemList = value
}
func (this *SGE_DBPlayer)GetBagItemList() []COM_ItemInst {
  this.Lock()
  defer this.Unlock()
  return this.BagItemList
}
func (this *SGE_DBPlayer)SetBlackMarketData(value COM_BlackMarket) {
  this.Lock()
  defer this.Unlock()
  this.BlackMarketData = value
}
func (this *SGE_DBPlayer)GetBlackMarketData() COM_BlackMarket {
  this.Lock()
  defer this.Unlock()
  return this.BlackMarketData
}
func (this *SGE_DBPlayer)SetChapterPondId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ChapterPondId = value
}
func (this *SGE_DBPlayer)GetChapterPondId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ChapterPondId
}
func (this *SGE_DBPlayer)Serialize(buffer *bytes.Buffer) error {
  {
    err := this.COM_Player.Serialize(buffer);
    if err != nil {
      return err
    }
  }
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(2)
  mask.writeBit(this.PlayerId!=0)
  mask.writeBit(len(this.Username) != 0)
  mask.writeBit(this.LoginTime!=0)
  mask.writeBit(this.LogoutTime!=0)
  mask.writeBit(this.BattleGroupIdx!=0)
  mask.writeBit(this.GenItemMaxGuid!=0)
  mask.writeBit(this.AssistantId!=0)
  mask.writeBit(len(this.BagItemList) != 0)
  mask.writeBit(true) //BlackMarketData
  mask.writeBit(this.ChapterPondId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize PlayerId
  {
    if(this.PlayerId!=0){
      err := write(buffer,this.PlayerId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Username
  if len(this.Username) != 0{
    err := write(buffer,this.Username)
    if err != nil {
      return err
    }
  }
  // serialize LoginTime
  {
    if(this.LoginTime!=0){
      err := write(buffer,this.LoginTime)
      if err != nil{
        return err
      }
    }
  }
  // serialize LogoutTime
  {
    if(this.LogoutTime!=0){
      err := write(buffer,this.LogoutTime)
      if err != nil{
        return err
      }
    }
  }
  // serialize BattleGroupIdx
  {
    if(this.BattleGroupIdx!=0){
      err := write(buffer,this.BattleGroupIdx)
      if err != nil{
        return err
      }
    }
  }
  // serialize GenItemMaxGuid
  {
    if(this.GenItemMaxGuid!=0){
      err := write(buffer,this.GenItemMaxGuid)
      if err != nil{
        return err
      }
    }
  }
  // serialize AssistantId
  {
    if(this.AssistantId!=0){
      err := write(buffer,this.AssistantId)
      if err != nil{
        return err
      }
    }
  }
  // serialize BagItemList
  if len(this.BagItemList) != 0{
    {
      err := write(buffer,uint(len(this.BagItemList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BagItemList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize BlackMarketData
  {
    err := this.BlackMarketData.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize ChapterPondId
  {
    if(this.ChapterPondId!=0){
      err := write(buffer,this.ChapterPondId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *SGE_DBPlayer)Deserialize(buffer *bytes.Buffer) error{
  {
    this.COM_Player.Deserialize(buffer);
  }
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,2);
  if err != nil{
    return err
  }
  // deserialize PlayerId
  if mask.readBit() {
    err := read(buffer,&this.PlayerId)
    if err != nil{
      return err
    }
  }
  // deserialize Username
  if mask.readBit() {
    err := read(buffer,&this.Username)
    if err != nil{
      return err
    }
  }
  // deserialize LoginTime
  if mask.readBit() {
    err := read(buffer,&this.LoginTime)
    if err != nil{
      return err
    }
  }
  // deserialize LogoutTime
  if mask.readBit() {
    err := read(buffer,&this.LogoutTime)
    if err != nil{
      return err
    }
  }
  // deserialize BattleGroupIdx
  if mask.readBit() {
    err := read(buffer,&this.BattleGroupIdx)
    if err != nil{
      return err
    }
  }
  // deserialize GenItemMaxGuid
  if mask.readBit() {
    err := read(buffer,&this.GenItemMaxGuid)
    if err != nil{
      return err
    }
  }
  // deserialize AssistantId
  if mask.readBit() {
    err := read(buffer,&this.AssistantId)
    if err != nil{
      return err
    }
  }
  // deserialize BagItemList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.BagItemList = make([]COM_ItemInst,size)
    for i,_ := range this.BagItemList{
      err := this.BagItemList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize BlackMarketData
  if mask.readBit() {
    err := this.BlackMarketData.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize ChapterPondId
  if mask.readBit() {
    err := read(buffer,&this.ChapterPondId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *SGE_DBPlayer)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
