package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_GuildViewerData struct{
  sync.Mutex
  GuildId int32  //0
  GuildName string  //1
  MasterName string  //2
  MemberNum int32  //3
  GuildVal int32  //4
}
func (this *COM_GuildViewerData)SetGuildId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildId = value
}
func (this *COM_GuildViewerData)GetGuildId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildId
}
func (this *COM_GuildViewerData)SetGuildName(value string) {
  this.Lock()
  defer this.Unlock()
  this.GuildName = value
}
func (this *COM_GuildViewerData)GetGuildName() string {
  this.Lock()
  defer this.Unlock()
  return this.GuildName
}
func (this *COM_GuildViewerData)SetMasterName(value string) {
  this.Lock()
  defer this.Unlock()
  this.MasterName = value
}
func (this *COM_GuildViewerData)GetMasterName() string {
  this.Lock()
  defer this.Unlock()
  return this.MasterName
}
func (this *COM_GuildViewerData)SetMemberNum(value int32) {
  this.Lock()
  defer this.Unlock()
  this.MemberNum = value
}
func (this *COM_GuildViewerData)GetMemberNum() int32 {
  this.Lock()
  defer this.Unlock()
  return this.MemberNum
}
func (this *COM_GuildViewerData)SetGuildVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildVal = value
}
func (this *COM_GuildViewerData)GetGuildVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildVal
}
func (this *COM_GuildViewerData)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.GuildId!=0)
  mask.writeBit(len(this.GuildName) != 0)
  mask.writeBit(len(this.MasterName) != 0)
  mask.writeBit(this.MemberNum!=0)
  mask.writeBit(this.GuildVal!=0)
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
  // serialize GuildName
  if len(this.GuildName) != 0{
    err := write(buffer,this.GuildName)
    if err != nil {
      return err
    }
  }
  // serialize MasterName
  if len(this.MasterName) != 0{
    err := write(buffer,this.MasterName)
    if err != nil {
      return err
    }
  }
  // serialize MemberNum
  {
    if(this.MemberNum!=0){
      err := write(buffer,this.MemberNum)
      if err != nil{
        return err
      }
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
  return nil
}
func (this *COM_GuildViewerData)Deserialize(buffer *bytes.Buffer) error{
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
  // deserialize GuildName
  if mask.readBit() {
    err := read(buffer,&this.GuildName)
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
  // deserialize MemberNum
  if mask.readBit() {
    err := read(buffer,&this.MemberNum)
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
  return nil
}
func (this *COM_GuildViewerData)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
