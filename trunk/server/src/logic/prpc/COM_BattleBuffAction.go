package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleBuffAction struct{
  sync.Mutex
  BuffData int32  //0
  Dead bool  //1
  BuffChange COM_BattleBuff  //2
}
func (this *COM_BattleBuffAction)SetBuffData(value int32) {
  this.Lock()
  defer this.Unlock()
  this.BuffData = value
}
func (this *COM_BattleBuffAction)GetBuffData() int32 {
  this.Lock()
  defer this.Unlock()
  return this.BuffData
}
func (this *COM_BattleBuffAction)SetDead(value bool) {
  this.Lock()
  defer this.Unlock()
  this.Dead = value
}
func (this *COM_BattleBuffAction)GetDead() bool {
  this.Lock()
  defer this.Unlock()
  return this.Dead
}
func (this *COM_BattleBuffAction)SetBuffChange(value COM_BattleBuff) {
  this.Lock()
  defer this.Unlock()
  this.BuffChange = value
}
func (this *COM_BattleBuffAction)GetBuffChange() COM_BattleBuff {
  this.Lock()
  defer this.Unlock()
  return this.BuffChange
}
func (this *COM_BattleBuffAction)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := NewMask1(1)
  mask.WriteBit(this.BuffData!=0)
  mask.WriteBit(this.Dead)
  mask.WriteBit(true) //BuffChange
  {
    err := Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize BuffData
  {
    if(this.BuffData!=0){
      err := Write(buffer,this.BuffData)
      if err != nil{
        return err
      }
    }
  }
  // serialize Dead
  {
  }
  // serialize BuffChange
  {
    err := this.BuffChange.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_BattleBuffAction)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize BuffData
  if mask.ReadBit() {
    err := Read(buffer,&this.BuffData)
    if err != nil{
      return err
    }
  }
  // deserialize Dead
  this.Dead = mask.ReadBit();
  // deserialize BuffChange
  if mask.ReadBit() {
    err := this.BuffChange.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_BattleBuffAction)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
