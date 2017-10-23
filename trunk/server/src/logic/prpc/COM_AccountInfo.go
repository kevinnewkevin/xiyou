package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_AccountInfo struct{
  sync.Mutex
  SessionCode string  //0
  MyPlayer COM_Player  //1
}
func (this *COM_AccountInfo)SetSessionCode(value string) {
  this.Lock()
  defer this.Unlock()
  this.SessionCode = value
}
func (this *COM_AccountInfo)GetSessionCode() string {
  this.Lock()
  defer this.Unlock()
  return this.SessionCode
}
func (this *COM_AccountInfo)SetMyPlayer(value COM_Player) {
  this.Lock()
  defer this.Unlock()
  this.MyPlayer = value
}
func (this *COM_AccountInfo)GetMyPlayer() COM_Player {
  this.Lock()
  defer this.Unlock()
  return this.MyPlayer
}
func (this *COM_AccountInfo)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.SessionCode) != 0)
  mask.writeBit(true) //MyPlayer
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize SessionCode
  if len(this.SessionCode) != 0{
    err := write(buffer,this.SessionCode)
    if err != nil {
      return err
    }
  }
  // serialize MyPlayer
  {
    err := this.MyPlayer.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_AccountInfo)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize SessionCode
  if mask.readBit() {
    err := read(buffer,&this.SessionCode)
    if err != nil{
      return err
    }
  }
  // deserialize MyPlayer
  if mask.readBit() {
    err := this.MyPlayer.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_AccountInfo)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
