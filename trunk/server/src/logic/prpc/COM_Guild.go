package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_Guild struct{
  sync.Mutex
  CreateTime int64  //0
  GuildId int32  //1
  Master int64  //2
  MasterName string  //3
  GuildName string  //4
  GuildVal int32  //5
  IsRatify bool  //6
  Require int32  //7
  Contribution int32  //8
  RequestList []COM_GuildRequestData  //9
}
func (this *COM_Guild)SetCreateTime(value int64) {
  this.Lock()
  defer this.Unlock()
  this.CreateTime = value
}
func (this *COM_Guild)GetCreateTime() int64 {
  this.Lock()
  defer this.Unlock()
  return this.CreateTime
}
func (this *COM_Guild)SetGuildId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildId = value
}
func (this *COM_Guild)GetGuildId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildId
}
func (this *COM_Guild)SetMaster(value int64) {
  this.Lock()
  defer this.Unlock()
  this.Master = value
}
func (this *COM_Guild)GetMaster() int64 {
  this.Lock()
  defer this.Unlock()
  return this.Master
}
func (this *COM_Guild)SetMasterName(value string) {
  this.Lock()
  defer this.Unlock()
  this.MasterName = value
}
func (this *COM_Guild)GetMasterName() string {
  this.Lock()
  defer this.Unlock()
  return this.MasterName
}
func (this *COM_Guild)SetGuildName(value string) {
  this.Lock()
  defer this.Unlock()
  this.GuildName = value
}
func (this *COM_Guild)GetGuildName() string {
  this.Lock()
  defer this.Unlock()
  return this.GuildName
}
func (this *COM_Guild)SetGuildVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildVal = value
}
func (this *COM_Guild)GetGuildVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildVal
}
func (this *COM_Guild)SetIsRatify(value bool) {
  this.Lock()
  defer this.Unlock()
  this.IsRatify = value
}
func (this *COM_Guild)GetIsRatify() bool {
  this.Lock()
  defer this.Unlock()
  return this.IsRatify
}
func (this *COM_Guild)SetRequire(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Require = value
}
func (this *COM_Guild)GetRequire() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Require
}
func (this *COM_Guild)SetContribution(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Contribution = value
}
func (this *COM_Guild)GetContribution() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Contribution
}
func (this *COM_Guild)SetRequestList(value []COM_GuildRequestData) {
  this.Lock()
  defer this.Unlock()
  this.RequestList = value
}
func (this *COM_Guild)GetRequestList() []COM_GuildRequestData {
  this.Lock()
  defer this.Unlock()
  return this.RequestList
}
func (this *COM_Guild)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(2)
  mask.writeBit(this.CreateTime!=0)
  mask.writeBit(this.GuildId!=0)
  mask.writeBit(this.Master!=0)
  mask.writeBit(len(this.MasterName) != 0)
  mask.writeBit(len(this.GuildName) != 0)
  mask.writeBit(this.GuildVal!=0)
  mask.writeBit(this.IsRatify)
  mask.writeBit(this.Require!=0)
  mask.writeBit(this.Contribution!=0)
  mask.writeBit(len(this.RequestList) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize CreateTime
  {
    if(this.CreateTime!=0){
      err := write(buffer,this.CreateTime)
      if err != nil{
        return err
      }
    }
  }
  // serialize GuildId
  {
    if(this.GuildId!=0){
      err := write(buffer,this.GuildId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Master
  {
    if(this.Master!=0){
      err := write(buffer,this.Master)
      if err != nil{
        return err
      }
    }
  }
  // serialize MasterName
  if len(this.MasterName) != 0{
    err := write(buffer,this.MasterName)
    if err != nil {
      return err
    }
  }
  // serialize GuildName
  if len(this.GuildName) != 0{
    err := write(buffer,this.GuildName)
    if err != nil {
      return err
    }
  }
  // serialize GuildVal
  {
    if(this.GuildVal!=0){
      err := write(buffer,this.GuildVal)
      if err != nil{
        return err
      }
    }
  }
  // serialize IsRatify
  {
  }
  // serialize Require
  {
    if(this.Require!=0){
      err := write(buffer,this.Require)
      if err != nil{
        return err
      }
    }
  }
  // serialize Contribution
  {
    if(this.Contribution!=0){
      err := write(buffer,this.Contribution)
      if err != nil{
        return err
      }
    }
  }
  // serialize RequestList
  if len(this.RequestList) != 0{
    {
      err := write(buffer,uint(len(this.RequestList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.RequestList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_Guild)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,2);
  if err != nil{
    return err
  }
  // deserialize CreateTime
  if mask.readBit() {
    err := read(buffer,&this.CreateTime)
    if err != nil{
      return err
    }
  }
  // deserialize GuildId
  if mask.readBit() {
    err := read(buffer,&this.GuildId)
    if err != nil{
      return err
    }
  }
  // deserialize Master
  if mask.readBit() {
    err := read(buffer,&this.Master)
    if err != nil{
      return err
    }
  }
  // deserialize MasterName
  if mask.readBit() {
    err := read(buffer,&this.MasterName)
    if err != nil{
      return err
    }
  }
  // deserialize GuildName
  if mask.readBit() {
    err := read(buffer,&this.GuildName)
    if err != nil{
      return err
    }
  }
  // deserialize GuildVal
  if mask.readBit() {
    err := read(buffer,&this.GuildVal)
    if err != nil{
      return err
    }
  }
  // deserialize IsRatify
  this.IsRatify = mask.readBit();
  // deserialize Require
  if mask.readBit() {
    err := read(buffer,&this.Require)
    if err != nil{
      return err
    }
  }
  // deserialize Contribution
  if mask.readBit() {
    err := read(buffer,&this.Contribution)
    if err != nil{
      return err
    }
  }
  // deserialize RequestList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.RequestList = make([]COM_GuildRequestData,size)
    for i,_ := range this.RequestList{
      err := this.RequestList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_Guild)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
