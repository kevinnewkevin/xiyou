package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_GuildMember struct{
  sync.Mutex
  GuildId int32  //0
  RoleId int64  //1
  RoleName string  //2
  Level int32  //3
  UnitId int32  //4
  Job int  //5
  TianTiVal int32  //6
  IsOnline bool  //7
}
func (this *COM_GuildMember)SetGuildId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildId = value
}
func (this *COM_GuildMember)GetGuildId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildId
}
func (this *COM_GuildMember)SetRoleId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.RoleId = value
}
func (this *COM_GuildMember)GetRoleId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.RoleId
}
func (this *COM_GuildMember)SetRoleName(value string) {
  this.Lock()
  defer this.Unlock()
  this.RoleName = value
}
func (this *COM_GuildMember)GetRoleName() string {
  this.Lock()
  defer this.Unlock()
  return this.RoleName
}
func (this *COM_GuildMember)SetLevel(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *COM_GuildMember)GetLevel() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *COM_GuildMember)SetUnitId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.UnitId = value
}
func (this *COM_GuildMember)GetUnitId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.UnitId
}
func (this *COM_GuildMember)SetJob(value int) {
  this.Lock()
  defer this.Unlock()
  this.Job = value
}
func (this *COM_GuildMember)GetJob() int {
  this.Lock()
  defer this.Unlock()
  return this.Job
}
func (this *COM_GuildMember)SetTianTiVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TianTiVal = value
}
func (this *COM_GuildMember)GetTianTiVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TianTiVal
}
func (this *COM_GuildMember)SetIsOnline(value bool) {
  this.Lock()
  defer this.Unlock()
  this.IsOnline = value
}
func (this *COM_GuildMember)GetIsOnline() bool {
  this.Lock()
  defer this.Unlock()
  return this.IsOnline
}
func (this *COM_GuildMember)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.GuildId!=0)
  mask.writeBit(this.RoleId!=0)
  mask.writeBit(len(this.RoleName) != 0)
  mask.writeBit(this.Level!=0)
  mask.writeBit(this.UnitId!=0)
  mask.writeBit(this.Job!=0)
  mask.writeBit(this.TianTiVal!=0)
  mask.writeBit(this.IsOnline)
  {
    err := write(buffer,mask.bytes())
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
  // serialize RoleId
  {
    if(this.RoleId!=0){
      err := write(buffer,this.RoleId)
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
  // serialize Level
  {
    if(this.Level!=0){
      err := write(buffer,this.Level)
      if err != nil{
        return err
      }
    }
  }
  // serialize UnitId
  {
    if(this.UnitId!=0){
      err := write(buffer,this.UnitId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Job
  {
    if(this.Job!=0){
      err := write(buffer,this.Job)
      if err != nil{
        return err
      }
    }
  }
  // serialize TianTiVal
  {
    if(this.TianTiVal!=0){
      err := write(buffer,this.TianTiVal)
      if err != nil{
        return err
      }
    }
  }
  // serialize IsOnline
  {
  }
  return nil
}
func (this *COM_GuildMember)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize GuildId
  if mask.readBit() {
    err := read(buffer,&this.GuildId)
    if err != nil{
      return err
    }
  }
  // deserialize RoleId
  if mask.readBit() {
    err := read(buffer,&this.RoleId)
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
  // deserialize Level
  if mask.readBit() {
    err := read(buffer,&this.Level)
    if err != nil{
      return err
    }
  }
  // deserialize UnitId
  if mask.readBit() {
    err := read(buffer,&this.UnitId)
    if err != nil{
      return err
    }
  }
  // deserialize Job
  if mask.readBit() {
    err := read(buffer,&this.Job)
    if err != nil{
      return err
    }
  }
  // deserialize TianTiVal
  if mask.readBit() {
    err := read(buffer,&this.TianTiVal)
    if err != nil{
      return err
    }
  }
  // deserialize IsOnline
  this.IsOnline = mask.readBit();
  return nil
}
func (this *COM_GuildMember)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
