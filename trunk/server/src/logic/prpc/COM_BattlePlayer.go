package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattlePlayer struct{
  sync.Mutex
  Player COM_Player  //0
  MaxPoint int32  //1
  CurPoint int32  //2
  BattlePosition []COM_BattlePosition  //3
}
func (this *COM_BattlePlayer)SetPlayer(value COM_Player) {
  this.Lock()
  defer this.Unlock()
  this.Player = value
}
func (this *COM_BattlePlayer)GetPlayer() COM_Player {
  this.Lock()
  defer this.Unlock()
  return this.Player
}
func (this *COM_BattlePlayer)SetMaxPoint(value int32) {
  this.Lock()
  defer this.Unlock()
  this.MaxPoint = value
}
func (this *COM_BattlePlayer)GetMaxPoint() int32 {
  this.Lock()
  defer this.Unlock()
  return this.MaxPoint
}
func (this *COM_BattlePlayer)SetCurPoint(value int32) {
  this.Lock()
  defer this.Unlock()
  this.CurPoint = value
}
func (this *COM_BattlePlayer)GetCurPoint() int32 {
  this.Lock()
  defer this.Unlock()
  return this.CurPoint
}
func (this *COM_BattlePlayer)SetBattlePosition(value []COM_BattlePosition) {
  this.Lock()
  defer this.Unlock()
  this.BattlePosition = value
}
func (this *COM_BattlePlayer)GetBattlePosition() []COM_BattlePosition {
  this.Lock()
  defer this.Unlock()
  return this.BattlePosition
}
func (this *COM_BattlePlayer)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := NewMask1(1)
  mask.WriteBit(true) //Player
  mask.WriteBit(this.MaxPoint!=0)
  mask.WriteBit(this.CurPoint!=0)
  mask.WriteBit(len(this.BattlePosition) != 0)
  {
    err := Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Player
  {
    err := this.Player.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize MaxPoint
  {
    if(this.MaxPoint!=0){
      err := Write(buffer,this.MaxPoint)
      if err != nil{
        return err
      }
    }
  }
  // serialize CurPoint
  {
    if(this.CurPoint!=0){
      err := Write(buffer,this.CurPoint)
      if err != nil{
        return err
      }
    }
  }
  // serialize BattlePosition
  if len(this.BattlePosition) != 0{
    {
      err := Write(buffer,uint(len(this.BattlePosition)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BattlePosition {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattlePlayer)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Player
  if mask.ReadBit() {
    err := this.Player.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize MaxPoint
  if mask.ReadBit() {
    err := Read(buffer,&this.MaxPoint)
    if err != nil{
      return err
    }
  }
  // deserialize CurPoint
  if mask.ReadBit() {
    err := Read(buffer,&this.CurPoint)
    if err != nil{
      return err
    }
  }
  // deserialize BattlePosition
  if mask.ReadBit() {
    var size uint
    err := Read(buffer,&size)
    if err != nil{
      return err
    }
    this.BattlePosition = make([]COM_BattlePosition,size)
    for i,_ := range this.BattlePosition{
      err := this.BattlePosition[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattlePlayer)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
