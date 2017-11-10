package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type SGE_DBTopUnit struct{
  sync.Mutex
  Name string  //0
  Level int32  //1
  TianTi int32  //2
  DisplayID int32  //3
}
func (this *SGE_DBTopUnit)SetName(value string) {
  this.Lock()
  defer this.Unlock()
  this.Name = value
}
func (this *SGE_DBTopUnit)GetName() string {
  this.Lock()
  defer this.Unlock()
  return this.Name
}
func (this *SGE_DBTopUnit)SetLevel(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *SGE_DBTopUnit)GetLevel() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *SGE_DBTopUnit)SetTianTi(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TianTi = value
}
func (this *SGE_DBTopUnit)GetTianTi() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TianTi
}
func (this *SGE_DBTopUnit)SetDisplayID(value int32) {
  this.Lock()
  defer this.Unlock()
  this.DisplayID = value
}
func (this *SGE_DBTopUnit)GetDisplayID() int32 {
  this.Lock()
  defer this.Unlock()
  return this.DisplayID
}
func (this *SGE_DBTopUnit)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.Name) != 0)
  mask.writeBit(this.Level!=0)
  mask.writeBit(this.TianTi!=0)
  mask.writeBit(this.DisplayID!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Name
  if len(this.Name) != 0{
    err := write(buffer,this.Name)
    if err != nil {
      return err
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
  // serialize TianTi
  {
    if(this.TianTi!=0){
      err := write(buffer,this.TianTi)
      if err != nil{
        return err
      }
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
  return nil
}
func (this *SGE_DBTopUnit)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Name
  if mask.readBit() {
    err := read(buffer,&this.Name)
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
  // deserialize TianTi
  if mask.readBit() {
    err := read(buffer,&this.TianTi)
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
  return nil
}
func (this *SGE_DBTopUnit)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
