package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
  "suzuki/prpc"
)
type COM_LoginInfo struct{
  sync.Mutex
  Username string  //0
  Password string  //1
}
func (this *COM_LoginInfo)SetUsername(value string) {
  this.Lock()
  defer this.Unlock()
  this.Username = value
}
func (this *COM_LoginInfo)GetUsername() string {
  this.Lock()
  defer this.Unlock()
  return this.Username
}
func (this *COM_LoginInfo)SetPassword(value string) {
  this.Lock()
  defer this.Unlock()
  this.Password = value
}
func (this *COM_LoginInfo)GetPassword() string {
  this.Lock()
  defer this.Unlock()
  return this.Password
}
func (this *COM_LoginInfo)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(len(this.Username) != 0)
  mask.WriteBit(len(this.Password) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Username
  if len(this.Username) != 0{
    err := prpc.Write(buffer,this.Username)
    if err != nil {
      return err
    }
  }
  // serialize Password
  if len(this.Password) != 0{
    err := prpc.Write(buffer,this.Password)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_LoginInfo)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Username
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Username)
    if err != nil{
      return err
    }
  }
  // deserialize Password
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Password)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_LoginInfo)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
