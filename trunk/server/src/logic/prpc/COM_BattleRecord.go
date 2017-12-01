package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleRecord struct{
  sync.Mutex
  Battleid int32  //0
  Round int32  //1
  Type int32  //2
  Winner int64  //3
  Players []COM_ReportCamp  //4
  DefinePos []COM_BattleUnit  //5
  Report []COM_BattleReport  //6
}
func (this *COM_BattleRecord)SetBattleid(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Battleid = value
}
func (this *COM_BattleRecord)GetBattleid() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Battleid
}
func (this *COM_BattleRecord)SetRound(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Round = value
}
func (this *COM_BattleRecord)GetRound() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Round
}
func (this *COM_BattleRecord)SetType(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Type = value
}
func (this *COM_BattleRecord)GetType() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Type
}
func (this *COM_BattleRecord)SetWinner(value int64) {
  this.Lock()
  defer this.Unlock()
  this.Winner = value
}
func (this *COM_BattleRecord)GetWinner() int64 {
  this.Lock()
  defer this.Unlock()
  return this.Winner
}
func (this *COM_BattleRecord)SetPlayers(value []COM_ReportCamp) {
  this.Lock()
  defer this.Unlock()
  this.Players = value
}
func (this *COM_BattleRecord)GetPlayers() []COM_ReportCamp {
  this.Lock()
  defer this.Unlock()
  return this.Players
}
func (this *COM_BattleRecord)SetDefinePos(value []COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.DefinePos = value
}
func (this *COM_BattleRecord)GetDefinePos() []COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.DefinePos
}
func (this *COM_BattleRecord)SetReport(value []COM_BattleReport) {
  this.Lock()
  defer this.Unlock()
  this.Report = value
}
func (this *COM_BattleRecord)GetReport() []COM_BattleReport {
  this.Lock()
  defer this.Unlock()
  return this.Report
}
func (this *COM_BattleRecord)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Battleid!=0)
  mask.writeBit(this.Round!=0)
  mask.writeBit(this.Type!=0)
  mask.writeBit(this.Winner!=0)
  mask.writeBit(len(this.Players) != 0)
  mask.writeBit(len(this.DefinePos) != 0)
  mask.writeBit(len(this.Report) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Battleid
  {
    if(this.Battleid!=0){
      err := write(buffer,this.Battleid)
      if err != nil{
        return err
      }
    }
  }
  // serialize Round
  {
    if(this.Round!=0){
      err := write(buffer,this.Round)
      if err != nil{
        return err
      }
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
  // serialize Winner
  {
    if(this.Winner!=0){
      err := write(buffer,this.Winner)
      if err != nil{
        return err
      }
    }
  }
  // serialize Players
  if len(this.Players) != 0{
    {
      err := write(buffer,uint(len(this.Players)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Players {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize DefinePos
  if len(this.DefinePos) != 0{
    {
      err := write(buffer,uint(len(this.DefinePos)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.DefinePos {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize Report
  if len(this.Report) != 0{
    {
      err := write(buffer,uint(len(this.Report)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Report {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleRecord)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Battleid
  if mask.readBit() {
    err := read(buffer,&this.Battleid)
    if err != nil{
      return err
    }
  }
  // deserialize Round
  if mask.readBit() {
    err := read(buffer,&this.Round)
    if err != nil{
      return err
    }
  }
  // deserialize Type
  if mask.readBit() {
    err := read(buffer,&this.Type)
    if err != nil{
      return err
    }
  }
  // deserialize Winner
  if mask.readBit() {
    err := read(buffer,&this.Winner)
    if err != nil{
      return err
    }
  }
  // deserialize Players
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Players = make([]COM_ReportCamp,size)
    for i,_ := range this.Players{
      err := this.Players[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize DefinePos
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.DefinePos = make([]COM_BattleUnit,size)
    for i,_ := range this.DefinePos{
      err := this.DefinePos[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize Report
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Report = make([]COM_BattleReport,size)
    for i,_ := range this.Report{
      err := this.Report[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleRecord)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
