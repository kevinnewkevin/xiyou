package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_Chat struct{
  sync.Mutex
  Type int8  //0
  PlayerInstId int64  //1
  PlayerName string  //2
  HeadIcon string  //3
  Level string  //4
  Content string  //5
  AudioUrl string  //6
  Audio []uint8  //7
  AudioLen int32  //8
}
func (this *COM_Chat)SetType(value int8) {
  this.Lock()
  defer this.Unlock()
  this.Type = value
}
func (this *COM_Chat)GetType() int8 {
  this.Lock()
  defer this.Unlock()
  return this.Type
}
func (this *COM_Chat)SetPlayerInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.PlayerInstId = value
}
func (this *COM_Chat)GetPlayerInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.PlayerInstId
}
func (this *COM_Chat)SetPlayerName(value string) {
  this.Lock()
  defer this.Unlock()
  this.PlayerName = value
}
func (this *COM_Chat)GetPlayerName() string {
  this.Lock()
  defer this.Unlock()
  return this.PlayerName
}
func (this *COM_Chat)SetHeadIcon(value string) {
  this.Lock()
  defer this.Unlock()
  this.HeadIcon = value
}
func (this *COM_Chat)GetHeadIcon() string {
  this.Lock()
  defer this.Unlock()
  return this.HeadIcon
}
func (this *COM_Chat)SetLevel(value string) {
  this.Lock()
  defer this.Unlock()
  this.Level = value
}
func (this *COM_Chat)GetLevel() string {
  this.Lock()
  defer this.Unlock()
  return this.Level
}
func (this *COM_Chat)SetContent(value string) {
  this.Lock()
  defer this.Unlock()
  this.Content = value
}
func (this *COM_Chat)GetContent() string {
  this.Lock()
  defer this.Unlock()
  return this.Content
}
func (this *COM_Chat)SetAudioUrl(value string) {
  this.Lock()
  defer this.Unlock()
  this.AudioUrl = value
}
func (this *COM_Chat)GetAudioUrl() string {
  this.Lock()
  defer this.Unlock()
  return this.AudioUrl
}
func (this *COM_Chat)SetAudio(value []uint8) {
  this.Lock()
  defer this.Unlock()
  this.Audio = value
}
func (this *COM_Chat)GetAudio() []uint8 {
  this.Lock()
  defer this.Unlock()
  return this.Audio
}
func (this *COM_Chat)SetAudioLen(value int32) {
  this.Lock()
  defer this.Unlock()
  this.AudioLen = value
}
func (this *COM_Chat)GetAudioLen() int32 {
  this.Lock()
  defer this.Unlock()
  return this.AudioLen
}
func (this *COM_Chat)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(2)
  mask.writeBit(this.Type!=0)
  mask.writeBit(this.PlayerInstId!=0)
  mask.writeBit(len(this.PlayerName) != 0)
  mask.writeBit(len(this.HeadIcon) != 0)
  mask.writeBit(len(this.Level) != 0)
  mask.writeBit(len(this.Content) != 0)
  mask.writeBit(len(this.AudioUrl) != 0)
  mask.writeBit(len(this.Audio) != 0)
  mask.writeBit(this.AudioLen!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Type
  {
    if(this.Type!=0){
      err := write(buffer,this.Type)
      if err != nil{
        return err
      }
    }
  }
  // serialize PlayerInstId
  {
    if(this.PlayerInstId!=0){
      err := write(buffer,this.PlayerInstId)
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
  // serialize HeadIcon
  if len(this.HeadIcon) != 0{
    err := write(buffer,this.HeadIcon)
    if err != nil {
      return err
    }
  }
  // serialize Level
  if len(this.Level) != 0{
    err := write(buffer,this.Level)
    if err != nil {
      return err
    }
  }
  // serialize Content
  if len(this.Content) != 0{
    err := write(buffer,this.Content)
    if err != nil {
      return err
    }
  }
  // serialize AudioUrl
  if len(this.AudioUrl) != 0{
    err := write(buffer,this.AudioUrl)
    if err != nil {
      return err
    }
  }
  // serialize Audio
  if len(this.Audio) != 0{
    {
      err := write(buffer,uint(len(this.Audio)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Audio {
      err := write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  // serialize AudioLen
  {
    if(this.AudioLen!=0){
      err := write(buffer,this.AudioLen)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_Chat)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,2);
  if err != nil{
    return err
  }
  // deserialize Type
  if mask.readBit() {
    err := read(buffer,&this.Type)
    if err != nil{
      return err
    }
  }
  // deserialize PlayerInstId
  if mask.readBit() {
    err := read(buffer,&this.PlayerInstId)
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
  // deserialize HeadIcon
  if mask.readBit() {
    err := read(buffer,&this.HeadIcon)
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
  // deserialize Content
  if mask.readBit() {
    err := read(buffer,&this.Content)
    if err != nil{
      return err
    }
  }
  // deserialize AudioUrl
  if mask.readBit() {
    err := read(buffer,&this.AudioUrl)
    if err != nil{
      return err
    }
  }
  // deserialize Audio
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Audio = make([]uint8,size)
    for i,_ := range this.Audio{
      err := read(buffer,&this.Audio[i])
      if err != nil{
        return err
      }
    }
  }
  // deserialize AudioLen
  if mask.readBit() {
    err := read(buffer,&this.AudioLen)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_Chat)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
