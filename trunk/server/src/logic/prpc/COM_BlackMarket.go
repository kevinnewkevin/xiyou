package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BlackMarket struct{
  sync.Mutex
  RefreshTime int64  //0
  RefreshNum int32  //1
  ShopItems []COM_MlackMarketItemData  //2
}
func (this *COM_BlackMarket)SetRefreshTime(value int64) {
  this.Lock()
  defer this.Unlock()
  this.RefreshTime = value
}
func (this *COM_BlackMarket)GetRefreshTime() int64 {
  this.Lock()
  defer this.Unlock()
  return this.RefreshTime
}
func (this *COM_BlackMarket)SetRefreshNum(value int32) {
  this.Lock()
  defer this.Unlock()
  this.RefreshNum = value
}
func (this *COM_BlackMarket)GetRefreshNum() int32 {
  this.Lock()
  defer this.Unlock()
  return this.RefreshNum
}
func (this *COM_BlackMarket)SetShopItems(value []COM_MlackMarketItemData) {
  this.Lock()
  defer this.Unlock()
  this.ShopItems = value
}
func (this *COM_BlackMarket)GetShopItems() []COM_MlackMarketItemData {
  this.Lock()
  defer this.Unlock()
  return this.ShopItems
}
func (this *COM_BlackMarket)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.RefreshTime!=0)
  mask.writeBit(this.RefreshNum!=0)
  mask.writeBit(len(this.ShopItems) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize RefreshTime
  {
    if(this.RefreshTime!=0){
      err := write(buffer,this.RefreshTime)
      if err != nil{
        return err
      }
    }
  }
  // serialize RefreshNum
  {
    if(this.RefreshNum!=0){
      err := write(buffer,this.RefreshNum)
      if err != nil{
        return err
      }
    }
  }
  // serialize ShopItems
  if len(this.ShopItems) != 0{
    {
      err := write(buffer,uint(len(this.ShopItems)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.ShopItems {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BlackMarket)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize RefreshTime
  if mask.readBit() {
    err := read(buffer,&this.RefreshTime)
    if err != nil{
      return err
    }
  }
  // deserialize RefreshNum
  if mask.readBit() {
    err := read(buffer,&this.RefreshNum)
    if err != nil{
      return err
    }
  }
  // deserialize ShopItems
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.ShopItems = make([]COM_MlackMarketItemData,size)
    for i,_ := range this.ShopItems{
      err := this.ShopItems[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BlackMarket)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
