package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_Mail struct{
  sync.Mutex
  MailId int32  //0
  Mailtype int  //1
  MailTimestamp int64  //2
  SendPlayerName string  //3
  RecvPlayerName string  //4
  Title string  //5
  Content string  //6
  Copper int32  //7
  Gold int32  //8
  Hero int32  //9
  Items []COM_MailItem  //10
  IsRead bool  //11
}
func (this *COM_Mail)SetMailId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.MailId = value
}
func (this *COM_Mail)GetMailId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.MailId
}
func (this *COM_Mail)SetMailtype(value int) {
  this.Lock()
  defer this.Unlock()
  this.Mailtype = value
}
func (this *COM_Mail)GetMailtype() int {
  this.Lock()
  defer this.Unlock()
  return this.Mailtype
}
func (this *COM_Mail)SetMailTimestamp(value int64) {
  this.Lock()
  defer this.Unlock()
  this.MailTimestamp = value
}
func (this *COM_Mail)GetMailTimestamp() int64 {
  this.Lock()
  defer this.Unlock()
  return this.MailTimestamp
}
func (this *COM_Mail)SetSendPlayerName(value string) {
  this.Lock()
  defer this.Unlock()
  this.SendPlayerName = value
}
func (this *COM_Mail)GetSendPlayerName() string {
  this.Lock()
  defer this.Unlock()
  return this.SendPlayerName
}
func (this *COM_Mail)SetRecvPlayerName(value string) {
  this.Lock()
  defer this.Unlock()
  this.RecvPlayerName = value
}
func (this *COM_Mail)GetRecvPlayerName() string {
  this.Lock()
  defer this.Unlock()
  return this.RecvPlayerName
}
func (this *COM_Mail)SetTitle(value string) {
  this.Lock()
  defer this.Unlock()
  this.Title = value
}
func (this *COM_Mail)GetTitle() string {
  this.Lock()
  defer this.Unlock()
  return this.Title
}
func (this *COM_Mail)SetContent(value string) {
  this.Lock()
  defer this.Unlock()
  this.Content = value
}
func (this *COM_Mail)GetContent() string {
  this.Lock()
  defer this.Unlock()
  return this.Content
}
func (this *COM_Mail)SetCopper(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Copper = value
}
func (this *COM_Mail)GetCopper() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Copper
}
func (this *COM_Mail)SetGold(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Gold = value
}
func (this *COM_Mail)GetGold() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Gold
}
func (this *COM_Mail)SetHero(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Hero = value
}
func (this *COM_Mail)GetHero() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Hero
}
func (this *COM_Mail)SetItems(value []COM_MailItem) {
  this.Lock()
  defer this.Unlock()
  this.Items = value
}
func (this *COM_Mail)GetItems() []COM_MailItem {
  this.Lock()
  defer this.Unlock()
  return this.Items
}
func (this *COM_Mail)SetIsRead(value bool) {
  this.Lock()
  defer this.Unlock()
  this.IsRead = value
}
func (this *COM_Mail)GetIsRead() bool {
  this.Lock()
  defer this.Unlock()
  return this.IsRead
}
func (this *COM_Mail)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(2)
  mask.writeBit(this.MailId!=0)
  mask.writeBit(this.Mailtype!=0)
  mask.writeBit(this.MailTimestamp!=0)
  mask.writeBit(len(this.SendPlayerName) != 0)
  mask.writeBit(len(this.RecvPlayerName) != 0)
  mask.writeBit(len(this.Title) != 0)
  mask.writeBit(len(this.Content) != 0)
  mask.writeBit(this.Copper!=0)
  mask.writeBit(this.Gold!=0)
  mask.writeBit(this.Hero!=0)
  mask.writeBit(len(this.Items) != 0)
  mask.writeBit(this.IsRead)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize MailId
  {
    if(this.MailId!=0){
      err := write(buffer,this.MailId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Mailtype
  {
    if(this.Mailtype!=0){
      err := write(buffer,this.Mailtype)
      if err != nil{
        return err
      }
    }
  }
  // serialize MailTimestamp
  {
    if(this.MailTimestamp!=0){
      err := write(buffer,this.MailTimestamp)
      if err != nil{
        return err
      }
    }
  }
  // serialize SendPlayerName
  if len(this.SendPlayerName) != 0{
    err := write(buffer,this.SendPlayerName)
    if err != nil {
      return err
    }
  }
  // serialize RecvPlayerName
  if len(this.RecvPlayerName) != 0{
    err := write(buffer,this.RecvPlayerName)
    if err != nil {
      return err
    }
  }
  // serialize Title
  if len(this.Title) != 0{
    err := write(buffer,this.Title)
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
  // serialize Copper
  {
    if(this.Copper!=0){
      err := write(buffer,this.Copper)
      if err != nil{
        return err
      }
    }
  }
  // serialize Gold
  {
    if(this.Gold!=0){
      err := write(buffer,this.Gold)
      if err != nil{
        return err
      }
    }
  }
  // serialize Hero
  {
    if(this.Hero!=0){
      err := write(buffer,this.Hero)
      if err != nil{
        return err
      }
    }
  }
  // serialize Items
  if len(this.Items) != 0{
    {
      err := write(buffer,uint(len(this.Items)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Items {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize IsRead
  {
  }
  return nil
}
func (this *COM_Mail)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,2);
  if err != nil{
    return err
  }
  // deserialize MailId
  if mask.readBit() {
    err := read(buffer,&this.MailId)
    if err != nil{
      return err
    }
  }
  // deserialize Mailtype
  if mask.readBit() {
    err := read(buffer,&this.Mailtype)
    if err != nil{
      return err
    }
  }
  // deserialize MailTimestamp
  if mask.readBit() {
    err := read(buffer,&this.MailTimestamp)
    if err != nil{
      return err
    }
  }
  // deserialize SendPlayerName
  if mask.readBit() {
    err := read(buffer,&this.SendPlayerName)
    if err != nil{
      return err
    }
  }
  // deserialize RecvPlayerName
  if mask.readBit() {
    err := read(buffer,&this.RecvPlayerName)
    if err != nil{
      return err
    }
  }
  // deserialize Title
  if mask.readBit() {
    err := read(buffer,&this.Title)
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
  // deserialize Copper
  if mask.readBit() {
    err := read(buffer,&this.Copper)
    if err != nil{
      return err
    }
  }
  // deserialize Gold
  if mask.readBit() {
    err := read(buffer,&this.Gold)
    if err != nil{
      return err
    }
  }
  // deserialize Hero
  if mask.readBit() {
    err := read(buffer,&this.Hero)
    if err != nil{
      return err
    }
  }
  // deserialize Items
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Items = make([]COM_MailItem,size)
    for i,_ := range this.Items{
      err := this.Items[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize IsRead
  this.IsRead = mask.readBit();
  return nil
}
func (this *COM_Mail)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
