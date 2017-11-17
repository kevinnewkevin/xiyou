package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type SGE_DBGuildAssistant struct{
  sync.Mutex
  Id int32  //0
  RoleName string  //1
  GuildId int32  //2
  ItemId int32  //3
  CrtCount int32  //4
  MaxCount int32  //5
  CatchNum int32  //6
  Donator []int64  //7
}
func (this *SGE_DBGuildAssistant)SetId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Id = value
}
func (this *SGE_DBGuildAssistant)GetId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Id
}
func (this *SGE_DBGuildAssistant)SetRoleName(value string) {
  this.Lock()
  defer this.Unlock()
  this.RoleName = value
}
func (this *SGE_DBGuildAssistant)GetRoleName() string {
  this.Lock()
  defer this.Unlock()
  return this.RoleName
}
func (this *SGE_DBGuildAssistant)SetGuildId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildId = value
}
func (this *SGE_DBGuildAssistant)GetGuildId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildId
}
func (this *SGE_DBGuildAssistant)SetItemId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ItemId = value
}
func (this *SGE_DBGuildAssistant)GetItemId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ItemId
}
func (this *SGE_DBGuildAssistant)SetCrtCount(value int32) {
  this.Lock()
  defer this.Unlock()
  this.CrtCount = value
}
func (this *SGE_DBGuildAssistant)GetCrtCount() int32 {
  this.Lock()
  defer this.Unlock()
  return this.CrtCount
}
func (this *SGE_DBGuildAssistant)SetMaxCount(value int32) {
  this.Lock()
  defer this.Unlock()
  this.MaxCount = value
}
func (this *SGE_DBGuildAssistant)GetMaxCount() int32 {
  this.Lock()
  defer this.Unlock()
  return this.MaxCount
}
func (this *SGE_DBGuildAssistant)SetCatchNum(value int32) {
  this.Lock()
  defer this.Unlock()
  this.CatchNum = value
}
func (this *SGE_DBGuildAssistant)GetCatchNum() int32 {
  this.Lock()
  defer this.Unlock()
  return this.CatchNum
}
func (this *SGE_DBGuildAssistant)SetDonator(value []int64) {
  this.Lock()
  defer this.Unlock()
  this.Donator = value
}
func (this *SGE_DBGuildAssistant)GetDonator() []int64 {
  this.Lock()
  defer this.Unlock()
  return this.Donator
}
func (this *SGE_DBGuildAssistant)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Id!=0)
  mask.writeBit(len(this.RoleName) != 0)
  mask.writeBit(this.GuildId!=0)
  mask.writeBit(this.ItemId!=0)
  mask.writeBit(this.CrtCount!=0)
  mask.writeBit(this.MaxCount!=0)
  mask.writeBit(this.CatchNum!=0)
  mask.writeBit(len(this.Donator) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Id
  {
    if(this.Id!=0){
      err := write(buffer,this.Id)
      if err != nil{
        return err
      }
    }
  }
  // serialize RoleName
  if len(this.RoleName) != 0{
    err := write(buffer,this.RoleName)
    if err != nil {
      return err
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
  // serialize ItemId
  {
    if(this.ItemId!=0){
      err := write(buffer,this.ItemId)
      if err != nil{
        return err
      }
    }
  }
  // serialize CrtCount
  {
    if(this.CrtCount!=0){
      err := write(buffer,this.CrtCount)
      if err != nil{
        return err
      }
    }
  }
  // serialize MaxCount
  {
    if(this.MaxCount!=0){
      err := write(buffer,this.MaxCount)
      if err != nil{
        return err
      }
    }
  }
  // serialize CatchNum
  {
    if(this.CatchNum!=0){
      err := write(buffer,this.CatchNum)
      if err != nil{
        return err
      }
    }
  }
  // serialize Donator
  if len(this.Donator) != 0{
    {
      err := write(buffer,uint(len(this.Donator)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Donator {
      err := write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *SGE_DBGuildAssistant)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Id
  if mask.readBit() {
    err := read(buffer,&this.Id)
    if err != nil{
      return err
    }
  }
  // deserialize RoleName
  if mask.readBit() {
    err := read(buffer,&this.RoleName)
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
  // deserialize ItemId
  if mask.readBit() {
    err := read(buffer,&this.ItemId)
    if err != nil{
      return err
    }
  }
  // deserialize CrtCount
  if mask.readBit() {
    err := read(buffer,&this.CrtCount)
    if err != nil{
      return err
    }
  }
  // deserialize MaxCount
  if mask.readBit() {
    err := read(buffer,&this.MaxCount)
    if err != nil{
      return err
    }
  }
  // deserialize CatchNum
  if mask.readBit() {
    err := read(buffer,&this.CatchNum)
    if err != nil{
      return err
    }
  }
  // deserialize Donator
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Donator = make([]int64,size)
    for i,_ := range this.Donator{
      err := read(buffer,&this.Donator[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *SGE_DBGuildAssistant)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
