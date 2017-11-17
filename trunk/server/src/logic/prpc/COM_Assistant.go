package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_Assistant struct{
  sync.Mutex
  Id int32  //0
  ItemId int32  //1
  CrtCount int32  //2
  MaxCount int32  //3
  PlayerName string  //4
  IsAssistanted []int64  //5
}
func (this *COM_Assistant)SetId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Id = value
}
func (this *COM_Assistant)GetId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Id
}
func (this *COM_Assistant)SetItemId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ItemId = value
}
func (this *COM_Assistant)GetItemId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ItemId
}
func (this *COM_Assistant)SetCrtCount(value int32) {
  this.Lock()
  defer this.Unlock()
  this.CrtCount = value
}
func (this *COM_Assistant)GetCrtCount() int32 {
  this.Lock()
  defer this.Unlock()
  return this.CrtCount
}
func (this *COM_Assistant)SetMaxCount(value int32) {
  this.Lock()
  defer this.Unlock()
  this.MaxCount = value
}
func (this *COM_Assistant)GetMaxCount() int32 {
  this.Lock()
  defer this.Unlock()
  return this.MaxCount
}
func (this *COM_Assistant)SetPlayerName(value string) {
  this.Lock()
  defer this.Unlock()
  this.PlayerName = value
}
func (this *COM_Assistant)GetPlayerName() string {
  this.Lock()
  defer this.Unlock()
  return this.PlayerName
}
func (this *COM_Assistant)SetIsAssistanted(value []int64) {
  this.Lock()
  defer this.Unlock()
  this.IsAssistanted = value
}
func (this *COM_Assistant)GetIsAssistanted() []int64 {
  this.Lock()
  defer this.Unlock()
  return this.IsAssistanted
}
func (this *COM_Assistant)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Id!=0)
  mask.writeBit(this.ItemId!=0)
  mask.writeBit(this.CrtCount!=0)
  mask.writeBit(this.MaxCount!=0)
  mask.writeBit(len(this.PlayerName) != 0)
  mask.writeBit(len(this.IsAssistanted) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Id
  {
    if(this.Id!=0){
      err := write(buffer,this.Id)
      if err != nil{
        return err
      }
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
  // serialize CrtCount
  {
    if(this.CrtCount!=0){
      err := write(buffer,this.CrtCount)
      if err != nil{
        return err
      }
    }
  }
  // serialize MaxCount
  {
    if(this.MaxCount!=0){
      err := write(buffer,this.MaxCount)
      if err != nil{
        return err
      }
    }
  }
  // serialize PlayerName
  if len(this.PlayerName) != 0{
    err := write(buffer,this.PlayerName)
    if err != nil {
      return err
    }
  }
  // serialize IsAssistanted
  if len(this.IsAssistanted) != 0{
    {
      err := write(buffer,uint(len(this.IsAssistanted)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.IsAssistanted {
      err := write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_Assistant)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Id
  if mask.readBit() {
    err := read(buffer,&this.Id)
    if err != nil{
      return err
    }
  }
  // deserialize ItemId
  if mask.readBit() {
    err := read(buffer,&this.ItemId)
    if err != nil{
      return err
    }
  }
  // deserialize CrtCount
  if mask.readBit() {
    err := read(buffer,&this.CrtCount)
    if err != nil{
      return err
    }
  }
  // deserialize MaxCount
  if mask.readBit() {
    err := read(buffer,&this.MaxCount)
    if err != nil{
      return err
    }
  }
  // deserialize PlayerName
  if mask.readBit() {
    err := read(buffer,&this.PlayerName)
    if err != nil{
      return err
    }
  }
  // deserialize IsAssistanted
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.IsAssistanted = make([]int64,size)
    for i,_ := range this.IsAssistanted{
      err := read(buffer,&this.IsAssistanted[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_Assistant)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
