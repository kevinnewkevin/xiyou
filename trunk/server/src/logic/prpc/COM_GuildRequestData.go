package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_GuildRequestData struct{
  sync.Mutex
  RoleId int64  //0
  Level int32  //1
  RoleName string  //2
  Time int64  //3
  UnitId int32  //4
  TianTiVal int32  //5
}
func (this *COM_GuildRequestData)SetRoleId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.RoleId = value
}
func (this *COM_GuildRequestData)GetRoleId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.RoleId
}
func (this *COM_GuildRequestData)SetLevel(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *COM_GuildRequestData)GetLevel() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *COM_GuildRequestData)SetRoleName(value string) {
  this.Lock()
  defer this.Unlock()
  this.RoleName = value
}
func (this *COM_GuildRequestData)GetRoleName() string {
  this.Lock()
  defer this.Unlock()
  return this.RoleName
}
func (this *COM_GuildRequestData)SetTime(value int64) {
  this.Lock()
  defer this.Unlock()
  this.Time = value
}
func (this *COM_GuildRequestData)GetTime() int64 {
  this.Lock()
  defer this.Unlock()
  return this.Time
}
func (this *COM_GuildRequestData)SetUnitId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.UnitId = value
}
func (this *COM_GuildRequestData)GetUnitId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.UnitId
}
func (this *COM_GuildRequestData)SetTianTiVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TianTiVal = value
}
func (this *COM_GuildRequestData)GetTianTiVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TianTiVal
}
func (this *COM_GuildRequestData)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.RoleId!=0)
  mask.writeBit(this.Level!=0)
  mask.writeBit(len(this.RoleName) != 0)
  mask.writeBit(this.Time!=0)
  mask.writeBit(this.UnitId!=0)
  mask.writeBit(this.TianTiVal!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
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
  // serialize Level
  {
    if(this.Level!=0){
      err := write(buffer,this.Level)
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
  // serialize Time
  {
    if(this.Time!=0){
      err := write(buffer,this.Time)
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
  // serialize TianTiVal
  {
    if(this.TianTiVal!=0){
      err := write(buffer,this.TianTiVal)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_GuildRequestData)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize RoleId
  if mask.readBit() {
    err := read(buffer,&this.RoleId)
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
  // deserialize RoleName
  if mask.readBit() {
    err := read(buffer,&this.RoleName)
    if err != nil{
      return err
    }
  }
  // deserialize Time
  if mask.readBit() {
    err := read(buffer,&this.Time)
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
  // deserialize TianTiVal
  if mask.readBit() {
    err := read(buffer,&this.TianTiVal)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_GuildRequestData)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
