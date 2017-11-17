package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_GuildDetails struct{
  sync.Mutex
  GuildId int32  //0
  GuildName string  //1
  MasterName string  //2
  MemberNum int32  //3
  GuildVal int32  //4
  IsRatify bool  //5
  Require int32  //6
  Contribution int32  //7
  Member []COM_GuildMember  //8
}
func (this *COM_GuildDetails)SetGuildId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildId = value
}
func (this *COM_GuildDetails)GetGuildId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildId
}
func (this *COM_GuildDetails)SetGuildName(value string) {
  this.Lock()
  defer this.Unlock()
  this.GuildName = value
}
func (this *COM_GuildDetails)GetGuildName() string {
  this.Lock()
  defer this.Unlock()
  return this.GuildName
}
func (this *COM_GuildDetails)SetMasterName(value string) {
  this.Lock()
  defer this.Unlock()
  this.MasterName = value
}
func (this *COM_GuildDetails)GetMasterName() string {
  this.Lock()
  defer this.Unlock()
  return this.MasterName
}
func (this *COM_GuildDetails)SetMemberNum(value int32) {
  this.Lock()
  defer this.Unlock()
  this.MemberNum = value
}
func (this *COM_GuildDetails)GetMemberNum() int32 {
  this.Lock()
  defer this.Unlock()
  return this.MemberNum
}
func (this *COM_GuildDetails)SetGuildVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.GuildVal = value
}
func (this *COM_GuildDetails)GetGuildVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.GuildVal
}
func (this *COM_GuildDetails)SetIsRatify(value bool) {
  this.Lock()
  defer this.Unlock()
  this.IsRatify = value
}
func (this *COM_GuildDetails)GetIsRatify() bool {
  this.Lock()
  defer this.Unlock()
  return this.IsRatify
}
func (this *COM_GuildDetails)SetRequire(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Require = value
}
func (this *COM_GuildDetails)GetRequire() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Require
}
func (this *COM_GuildDetails)SetContribution(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Contribution = value
}
func (this *COM_GuildDetails)GetContribution() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Contribution
}
func (this *COM_GuildDetails)SetMember(value []COM_GuildMember) {
  this.Lock()
  defer this.Unlock()
  this.Member = value
}
func (this *COM_GuildDetails)GetMember() []COM_GuildMember {
  this.Lock()
  defer this.Unlock()
  return this.Member
}
func (this *COM_GuildDetails)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(2)
  mask.writeBit(this.GuildId!=0)
  mask.writeBit(len(this.GuildName) != 0)
  mask.writeBit(len(this.MasterName) != 0)
  mask.writeBit(this.MemberNum!=0)
  mask.writeBit(this.GuildVal!=0)
  mask.writeBit(this.IsRatify)
  mask.writeBit(this.Require!=0)
  mask.writeBit(this.Contribution!=0)
  mask.writeBit(len(this.Member) != 0)
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
  // serialize Member
  if len(this.Member) != 0{
    {
      err := write(buffer,uint(len(this.Member)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Member {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_GuildDetails)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,2);
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
  // deserialize Member
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Member = make([]COM_GuildMember,size)
    for i,_ := range this.Member{
      err := this.Member[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_GuildDetails)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
