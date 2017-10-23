package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_MlackMarketItemData struct{
  sync.Mutex
  IsBuy bool  //0
  ItemId int32  //1
}
func (this *COM_MlackMarketItemData)SetIsBuy(value bool) {
  this.Lock()
  defer this.Unlock()
  this.IsBuy = value
}
func (this *COM_MlackMarketItemData)GetIsBuy() bool {
  this.Lock()
  defer this.Unlock()
  return this.IsBuy
}
func (this *COM_MlackMarketItemData)SetItemId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ItemId = value
}
func (this *COM_MlackMarketItemData)GetItemId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ItemId
}
func (this *COM_MlackMarketItemData)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := NewMask1(1)
  mask.WriteBit(this.IsBuy)
  mask.WriteBit(this.ItemId!=0)
  {
    err := Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize IsBuy
  {
  }
  // serialize ItemId
  {
    if(this.ItemId!=0){
      err := Write(buffer,this.ItemId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_MlackMarketItemData)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize IsBuy
  this.IsBuy = mask.ReadBit();
  // deserialize ItemId
  if mask.ReadBit() {
    err := Read(buffer,&this.ItemId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_MlackMarketItemData)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
