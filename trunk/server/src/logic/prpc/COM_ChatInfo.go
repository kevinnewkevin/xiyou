package prpc
import(
  "bytes"
  "encoding/json"
)
type COM_ChatInfo struct{
  COM_Chat
  AudioId int32  //0
  AssetId uint16  //1
  PlayerName string  //2
}
func (this *COM_ChatInfo)SetAudioId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.AudioId = value
}
func (this *COM_ChatInfo)GetAudioId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.AudioId
}
func (this *COM_ChatInfo)SetAssetId(value uint16) {
  this.Lock()
  defer this.Unlock()
  this.AssetId = value
}
func (this *COM_ChatInfo)GetAssetId() uint16 {
  this.Lock()
  defer this.Unlock()
  return this.AssetId
}
func (this *COM_ChatInfo)SetPlayerName(value string) {
  this.Lock()
  defer this.Unlock()
  this.PlayerName = value
}
func (this *COM_ChatInfo)GetPlayerName() string {
  this.Lock()
  defer this.Unlock()
  return this.PlayerName
}
func (this *COM_ChatInfo)Serialize(buffer *bytes.Buffer) error {
  {
    err := this.COM_Chat.Serialize(buffer);
    if err != nil {
      return err
    }
  }
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.AudioId!=0)
  mask.writeBit(this.AssetId!=0)
  mask.writeBit(len(this.PlayerName) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize AudioId
  {
    if(this.AudioId!=0){
      err := write(buffer,this.AudioId)
      if err != nil{
        return err
      }
    }
  }
  // serialize AssetId
  {
    if(this.AssetId!=0){
      err := write(buffer,this.AssetId)
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
  return nil
}
func (this *COM_ChatInfo)Deserialize(buffer *bytes.Buffer) error{
  {
    this.COM_Chat.Deserialize(buffer);
  }
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize AudioId
  if mask.readBit() {
    err := read(buffer,&this.AudioId)
    if err != nil{
      return err
    }
  }
  // deserialize AssetId
  if mask.readBit() {
    err := read(buffer,&this.AssetId)
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
  return nil
}
func (this *COM_ChatInfo)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
