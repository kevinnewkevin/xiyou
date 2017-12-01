package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_MailItem struct{
  sync.Mutex
  ItemId int32  //0
  ItemStack int32  //1
}
func (this *COM_MailItem)SetItemId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ItemId = value
}
func (this *COM_MailItem)GetItemId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ItemId
}
func (this *COM_MailItem)SetItemStack(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ItemStack = value
}
func (this *COM_MailItem)GetItemStack() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ItemStack
}
func (this *COM_MailItem)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.ItemId!=0)
  mask.writeBit(this.ItemStack!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize ItemId
  {
    if(this.ItemId!=0){
      err := write(buffer,this.ItemId)
      if err != nil{
        return err
      }
    }
  }
  // serialize ItemStack
  {
    if(this.ItemStack!=0){
      err := write(buffer,this.ItemStack)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_MailItem)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize ItemId
  if mask.readBit() {
    err := read(buffer,&this.ItemId)
    if err != nil{
      return err
    }
  }
  // deserialize ItemStack
  if mask.readBit() {
    err := read(buffer,&this.ItemStack)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_MailItem)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
