package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleRecord_Detail struct{
  sync.Mutex
  ReportId int64  //0
  Battleid int32  //1
  Winner int64  //2
  Players []COM_ReportCamp  //3
}
func (this *COM_BattleRecord_Detail)SetReportId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.ReportId = value
}
func (this *COM_BattleRecord_Detail)GetReportId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.ReportId
}
func (this *COM_BattleRecord_Detail)SetBattleid(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Battleid = value
}
func (this *COM_BattleRecord_Detail)GetBattleid() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Battleid
}
func (this *COM_BattleRecord_Detail)SetWinner(value int64) {
  this.Lock()
  defer this.Unlock()
  this.Winner = value
}
func (this *COM_BattleRecord_Detail)GetWinner() int64 {
  this.Lock()
  defer this.Unlock()
  return this.Winner
}
func (this *COM_BattleRecord_Detail)SetPlayers(value []COM_ReportCamp) {
  this.Lock()
  defer this.Unlock()
  this.Players = value
}
func (this *COM_BattleRecord_Detail)GetPlayers() []COM_ReportCamp {
  this.Lock()
  defer this.Unlock()
  return this.Players
}
func (this *COM_BattleRecord_Detail)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.ReportId!=0)
  mask.writeBit(this.Battleid!=0)
  mask.writeBit(this.Winner!=0)
  mask.writeBit(len(this.Players) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize ReportId
  {
    if(this.ReportId!=0){
      err := write(buffer,this.ReportId)
      if err != nil{
        return err
      }
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
  return nil
}
func (this *COM_BattleRecord_Detail)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize ReportId
  if mask.readBit() {
    err := read(buffer,&this.ReportId)
    if err != nil{
      return err
    }
  }
  // deserialize Battleid
  if mask.readBit() {
    err := read(buffer,&this.Battleid)
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
  return nil
}
func (this *COM_BattleRecord_Detail)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
