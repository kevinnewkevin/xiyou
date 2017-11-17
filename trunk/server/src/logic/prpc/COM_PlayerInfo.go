package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_PlayerInfo struct{
  sync.Mutex
  UnitLIst []COM_Unit  //0
  Level int32  //1
  Name string  //2
  DisplayID int32  //3
  ClanName string  //4
  TiatiVal int32  //5
  TiatiRank int32  //6
  IsOnline bool  //7
}
func (this *COM_PlayerInfo)SetUnitLIst(value []COM_Unit) {
  this.Lock()
  defer this.Unlock()
  this.UnitLIst = value
}
func (this *COM_PlayerInfo)GetUnitLIst() []COM_Unit {
  this.Lock()
  defer this.Unlock()
  return this.UnitLIst
}
func (this *COM_PlayerInfo)SetLevel(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *COM_PlayerInfo)GetLevel() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *COM_PlayerInfo)SetName(value string) {
  this.Lock()
  defer this.Unlock()
  this.Name = value
}
func (this *COM_PlayerInfo)GetName() string {
  this.Lock()
  defer this.Unlock()
  return this.Name
}
func (this *COM_PlayerInfo)SetDisplayID(value int32) {
  this.Lock()
  defer this.Unlock()
  this.DisplayID = value
}
func (this *COM_PlayerInfo)GetDisplayID() int32 {
  this.Lock()
  defer this.Unlock()
  return this.DisplayID
}
func (this *COM_PlayerInfo)SetClanName(value string) {
  this.Lock()
  defer this.Unlock()
  this.ClanName = value
}
func (this *COM_PlayerInfo)GetClanName() string {
  this.Lock()
  defer this.Unlock()
  return this.ClanName
}
func (this *COM_PlayerInfo)SetTiatiVal(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TiatiVal = value
}
func (this *COM_PlayerInfo)GetTiatiVal() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TiatiVal
}
func (this *COM_PlayerInfo)SetTiatiRank(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TiatiRank = value
}
func (this *COM_PlayerInfo)GetTiatiRank() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TiatiRank
}
func (this *COM_PlayerInfo)SetIsOnline(value bool) {
  this.Lock()
  defer this.Unlock()
  this.IsOnline = value
}
func (this *COM_PlayerInfo)GetIsOnline() bool {
  this.Lock()
  defer this.Unlock()
  return this.IsOnline
}
func (this *COM_PlayerInfo)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.UnitLIst) != 0)
  mask.writeBit(this.Level!=0)
  mask.writeBit(len(this.Name) != 0)
  mask.writeBit(this.DisplayID!=0)
  mask.writeBit(len(this.ClanName) != 0)
  mask.writeBit(this.TiatiVal!=0)
  mask.writeBit(this.TiatiRank!=0)
  mask.writeBit(this.IsOnline)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize UnitLIst
  if len(this.UnitLIst) != 0{
    {
      err := write(buffer,uint(len(this.UnitLIst)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.UnitLIst {
      err := value.Serialize(buffer)
      if err != nil {
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
  // serialize Name
  if len(this.Name) != 0{
    err := write(buffer,this.Name)
    if err != nil {
      return err
    }
  }
  // serialize DisplayID
  {
    if(this.DisplayID!=0){
      err := write(buffer,this.DisplayID)
      if err != nil{
        return err
      }
    }
  }
  // serialize ClanName
  if len(this.ClanName) != 0{
    err := write(buffer,this.ClanName)
    if err != nil {
      return err
    }
  }
  // serialize TiatiVal
  {
    if(this.TiatiVal!=0){
      err := write(buffer,this.TiatiVal)
      if err != nil{
        return err
      }
    }
  }
  // serialize TiatiRank
  {
    if(this.TiatiRank!=0){
      err := write(buffer,this.TiatiRank)
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
func (this *COM_PlayerInfo)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize UnitLIst
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.UnitLIst = make([]COM_Unit,size)
    for i,_ := range this.UnitLIst{
      err := this.UnitLIst[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize Level
  if mask.readBit() {
    err := read(buffer,&this.Level)
    if err != nil{
      return err
    }
  }
  // deserialize Name
  if mask.readBit() {
    err := read(buffer,&this.Name)
    if err != nil{
      return err
    }
  }
  // deserialize DisplayID
  if mask.readBit() {
    err := read(buffer,&this.DisplayID)
    if err != nil{
      return err
    }
  }
  // deserialize ClanName
  if mask.readBit() {
    err := read(buffer,&this.ClanName)
    if err != nil{
      return err
    }
  }
  // deserialize TiatiVal
  if mask.readBit() {
    err := read(buffer,&this.TiatiVal)
    if err != nil{
      return err
    }
  }
  // deserialize TiatiRank
  if mask.readBit() {
    err := read(buffer,&this.TiatiRank)
    if err != nil{
      return err
    }
  }
  // deserialize IsOnline
  this.IsOnline = mask.readBit();
  return nil
}
func (this *COM_PlayerInfo)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
