package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleSnape struct{
  sync.Mutex
  Camp int32  //0
  battleid int32  //1
  targetcards []int32  //2
  MainUnit []COM_BattleUnit  //3
  BattleField []COM_BattleUnit  //4
  HandCard []COM_BattleUnit  //5
}
func (this *COM_BattleSnape)SetCamp(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Camp = value
}
func (this *COM_BattleSnape)GetCamp() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Camp
}
func (this *COM_BattleSnape)Setbattleid(value int32) {
  this.Lock()
  defer this.Unlock()
  this.battleid = value
}
func (this *COM_BattleSnape)Getbattleid() int32 {
  this.Lock()
  defer this.Unlock()
  return this.battleid
}
func (this *COM_BattleSnape)Settargetcards(value []int32) {
  this.Lock()
  defer this.Unlock()
  this.targetcards = value
}
func (this *COM_BattleSnape)Gettargetcards() []int32 {
  this.Lock()
  defer this.Unlock()
  return this.targetcards
}
func (this *COM_BattleSnape)SetMainUnit(value []COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.MainUnit = value
}
func (this *COM_BattleSnape)GetMainUnit() []COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.MainUnit
}
func (this *COM_BattleSnape)SetBattleField(value []COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.BattleField = value
}
func (this *COM_BattleSnape)GetBattleField() []COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.BattleField
}
func (this *COM_BattleSnape)SetHandCard(value []COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.HandCard = value
}
func (this *COM_BattleSnape)GetHandCard() []COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.HandCard
}
func (this *COM_BattleSnape)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Camp!=0)
  mask.writeBit(this.battleid!=0)
  mask.writeBit(len(this.targetcards) != 0)
  mask.writeBit(len(this.MainUnit) != 0)
  mask.writeBit(len(this.BattleField) != 0)
  mask.writeBit(len(this.HandCard) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Camp
  {
    if(this.Camp!=0){
      err := write(buffer,this.Camp)
      if err != nil{
        return err
      }
    }
  }
  // serialize battleid
  {
    if(this.battleid!=0){
      err := write(buffer,this.battleid)
      if err != nil{
        return err
      }
    }
  }
  // serialize targetcards
  if len(this.targetcards) != 0{
    {
      err := write(buffer,uint(len(this.targetcards)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.targetcards {
      err := write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  // serialize MainUnit
  if len(this.MainUnit) != 0{
    {
      err := write(buffer,uint(len(this.MainUnit)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.MainUnit {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize BattleField
  if len(this.BattleField) != 0{
    {
      err := write(buffer,uint(len(this.BattleField)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BattleField {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize HandCard
  if len(this.HandCard) != 0{
    {
      err := write(buffer,uint(len(this.HandCard)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.HandCard {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleSnape)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Camp
  if mask.readBit() {
    err := read(buffer,&this.Camp)
    if err != nil{
      return err
    }
  }
  // deserialize battleid
  if mask.readBit() {
    err := read(buffer,&this.battleid)
    if err != nil{
      return err
    }
  }
  // deserialize targetcards
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.targetcards = make([]int32,size)
    for i,_ := range this.targetcards{
      err := read(buffer,&this.targetcards[i])
      if err != nil{
        return err
      }
    }
  }
  // deserialize MainUnit
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.MainUnit = make([]COM_BattleUnit,size)
    for i,_ := range this.MainUnit{
      err := this.MainUnit[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize BattleField
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.BattleField = make([]COM_BattleUnit,size)
    for i,_ := range this.BattleField{
      err := this.BattleField[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize HandCard
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.HandCard = make([]COM_BattleUnit,size)
    for i,_ := range this.HandCard{
      err := this.HandCard[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleSnape)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
