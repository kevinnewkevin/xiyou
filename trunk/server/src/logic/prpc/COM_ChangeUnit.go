package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_ChangeUnit struct{
  sync.Mutex
  Unit COM_BattleUnit  //0
  Status bool  //1
}
func (this *COM_ChangeUnit)SetUnit(value COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.Unit = value
}
func (this *COM_ChangeUnit)GetUnit() COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.Unit
}
func (this *COM_ChangeUnit)SetStatus(value bool) {
  this.Lock()
  defer this.Unlock()
  this.Status = value
}
func (this *COM_ChangeUnit)GetStatus() bool {
  this.Lock()
  defer this.Unlock()
  return this.Status
}
func (this *COM_ChangeUnit)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //Unit
  mask.writeBit(this.Status)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Unit
  {
    err := this.Unit.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize Status
  {
  }
  return nil
}
func (this *COM_ChangeUnit)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Unit
  if mask.readBit() {
    err := this.Unit.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize Status
  this.Status = mask.readBit();
  return nil
}
func (this *COM_ChangeUnit)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
