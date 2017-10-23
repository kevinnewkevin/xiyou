package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleBuff struct{
  sync.Mutex
  BuffId int32  //0
  Change int32  //1
}
func (this *COM_BattleBuff)SetBuffId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.BuffId = value
}
func (this *COM_BattleBuff)GetBuffId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.BuffId
}
func (this *COM_BattleBuff)SetChange(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Change = value
}
func (this *COM_BattleBuff)GetChange() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Change
}
func (this *COM_BattleBuff)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.BuffId!=0)
  mask.writeBit(this.Change!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize BuffId
  {
    if(this.BuffId!=0){
      err := write(buffer,this.BuffId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Change
  {
    if(this.Change!=0){
      err := write(buffer,this.Change)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleBuff)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize BuffId
  if mask.readBit() {
    err := read(buffer,&this.BuffId)
    if err != nil{
      return err
    }
  }
  // deserialize Change
  if mask.readBit() {
    err := read(buffer,&this.Change)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_BattleBuff)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
