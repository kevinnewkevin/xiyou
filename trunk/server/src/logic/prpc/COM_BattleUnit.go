package prpc
import(
  "bytes"
  "sync"
  "suzuki/prpc"
)
type COM_BattleUnit struct{
  sync.Mutex
  UnitId int32  //0
  InstId int64  //1
  Position int32  //2
  HP int32  //3
  CHP int32  //4
  Level int32  //5
  Name string  //6
}
func (this *COM_BattleUnit)SetUnitId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.UnitId = value
}
func (this *COM_BattleUnit)GetUnitId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.UnitId
}
func (this *COM_BattleUnit)SetInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.InstId = value
}
func (this *COM_BattleUnit)GetInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.InstId
}
func (this *COM_BattleUnit)SetPosition(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Position = value
}
func (this *COM_BattleUnit)GetPosition() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Position
}
func (this *COM_BattleUnit)SetHP(value int32) {
  this.Lock()
  defer this.Unlock()
  this.HP = value
}
func (this *COM_BattleUnit)GetHP() int32 {
  this.Lock()
  defer this.Unlock()
  return this.HP
}
func (this *COM_BattleUnit)SetCHP(value int32) {
  this.Lock()
  defer this.Unlock()
  this.CHP = value
}
func (this *COM_BattleUnit)GetCHP() int32 {
  this.Lock()
  defer this.Unlock()
  return this.CHP
}
func (this *COM_BattleUnit)SetLevel(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *COM_BattleUnit)GetLevel() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *COM_BattleUnit)SetName(value string) {
  this.Lock()
  defer this.Unlock()
  this.Name = value
}
func (this *COM_BattleUnit)GetName() string {
  this.Lock()
  defer this.Unlock()
  return this.Name
}
func (this *COM_BattleUnit)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.UnitId!=0)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.Position!=0)
  mask.WriteBit(this.HP!=0)
  mask.WriteBit(this.CHP!=0)
  mask.WriteBit(this.Level!=0)
  mask.WriteBit(len(this.Name) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize UnitId
  {
    if(this.UnitId!=0){
      err := prpc.Write(buffer,this.UnitId)
      if err != nil{
        return err
      }
    }
  }
  // serialize InstId
  {
    if(this.InstId!=0){
      err := prpc.Write(buffer,this.InstId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Position
  {
    if(this.Position!=0){
      err := prpc.Write(buffer,this.Position)
      if err != nil{
        return err
      }
    }
  }
  // serialize HP
  {
    if(this.HP!=0){
      err := prpc.Write(buffer,this.HP)
      if err != nil{
        return err
      }
    }
  }
  // serialize CHP
  {
    if(this.CHP!=0){
      err := prpc.Write(buffer,this.CHP)
      if err != nil{
        return err
      }
    }
  }
  // serialize Level
  {
    if(this.Level!=0){
      err := prpc.Write(buffer,this.Level)
      if err != nil{
        return err
      }
    }
  }
  // serialize Name
  if len(this.Name) != 0{
    err := prpc.Write(buffer,this.Name)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_BattleUnit)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize UnitId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.UnitId)
    if err != nil{
      return err
    }
  }
  // deserialize InstId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize Position
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Position)
    if err != nil{
      return err
    }
  }
  // deserialize HP
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.HP)
    if err != nil{
      return err
    }
  }
  // deserialize CHP
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.CHP)
    if err != nil{
      return err
    }
  }
  // deserialize Level
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Level)
    if err != nil{
      return err
    }
  }
  // deserialize Name
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Name)
    if err != nil{
      return err
    }
  }
  return nil
}
