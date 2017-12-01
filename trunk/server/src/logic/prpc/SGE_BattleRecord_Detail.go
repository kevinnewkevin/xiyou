package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type SGE_BattleRecord_Detail struct{
  sync.Mutex
  Detail []COM_BattleRecord_Detail  //0
}
func (this *SGE_BattleRecord_Detail)SetDetail(value []COM_BattleRecord_Detail) {
  this.Lock()
  defer this.Unlock()
  this.Detail = value
}
func (this *SGE_BattleRecord_Detail)GetDetail() []COM_BattleRecord_Detail {
  this.Lock()
  defer this.Unlock()
  return this.Detail
}
func (this *SGE_BattleRecord_Detail)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.Detail) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Detail
  if len(this.Detail) != 0{
    {
      err := write(buffer,uint(len(this.Detail)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Detail {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *SGE_BattleRecord_Detail)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Detail
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Detail = make([]COM_BattleRecord_Detail,size)
    for i,_ := range this.Detail{
      err := this.Detail[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *SGE_BattleRecord_Detail)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
