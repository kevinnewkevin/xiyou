package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_Friend struct{
  sync.Mutex
  InstId int64  //0
  Name string  //1
  Level int32  //2
  DisplayID int32  //3
}
func (this *COM_Friend)SetInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.InstId = value
}
func (this *COM_Friend)GetInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.InstId
}
func (this *COM_Friend)SetName(value string) {
  this.Lock()
  defer this.Unlock()
  this.Name = value
}
func (this *COM_Friend)GetName() string {
  this.Lock()
  defer this.Unlock()
  return this.Name
}
func (this *COM_Friend)SetLevel(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *COM_Friend)GetLevel() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *COM_Friend)SetDisplayID(value int32) {
  this.Lock()
  defer this.Unlock()
  this.DisplayID = value
}
func (this *COM_Friend)GetDisplayID() int32 {
  this.Lock()
  defer this.Unlock()
  return this.DisplayID
}
func (this *COM_Friend)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.InstId!=0)
  mask.writeBit(len(this.Name) != 0)
  mask.writeBit(this.Level!=0)
  mask.writeBit(this.DisplayID!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize InstId
  {
    if(this.InstId!=0){
      err := write(buffer,this.InstId)
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
  // serialize Level
  {
    if(this.Level!=0){
      err := write(buffer,this.Level)
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
func (this *COM_Friend)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize InstId
  if mask.readBit() {
    err := read(buffer,&this.InstId)
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
  // deserialize Level
  if mask.readBit() {
    err := read(buffer,&this.Level)
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
func (this *COM_Friend)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
