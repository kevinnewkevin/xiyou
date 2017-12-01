package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_ReportCamp struct{
  sync.Mutex
  Camp int8  //0
  InstId int64  //1
  Name string  //2
  TianTi int32  //3
  MainUnit COM_BattleUnit  //4
  Units []COM_BattleUnit  //5
}
func (this *COM_ReportCamp)SetCamp(value int8) {
  this.Lock()
  defer this.Unlock()
  this.Camp = value
}
func (this *COM_ReportCamp)GetCamp() int8 {
  this.Lock()
  defer this.Unlock()
  return this.Camp
}
func (this *COM_ReportCamp)SetInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.InstId = value
}
func (this *COM_ReportCamp)GetInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.InstId
}
func (this *COM_ReportCamp)SetName(value string) {
  this.Lock()
  defer this.Unlock()
  this.Name = value
}
func (this *COM_ReportCamp)GetName() string {
  this.Lock()
  defer this.Unlock()
  return this.Name
}
func (this *COM_ReportCamp)SetTianTi(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TianTi = value
}
func (this *COM_ReportCamp)GetTianTi() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TianTi
}
func (this *COM_ReportCamp)SetMainUnit(value COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.MainUnit = value
}
func (this *COM_ReportCamp)GetMainUnit() COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.MainUnit
}
func (this *COM_ReportCamp)SetUnits(value []COM_BattleUnit) {
  this.Lock()
  defer this.Unlock()
  this.Units = value
}
func (this *COM_ReportCamp)GetUnits() []COM_BattleUnit {
  this.Lock()
  defer this.Unlock()
  return this.Units
}
func (this *COM_ReportCamp)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Camp!=0)
  mask.writeBit(this.InstId!=0)
  mask.writeBit(len(this.Name) != 0)
  mask.writeBit(this.TianTi!=0)
  mask.writeBit(true) //MainUnit
  mask.writeBit(len(this.Units) != 0)
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
  // serialize InstId
  {
    if(this.InstId!=0){
      err := write(buffer,this.InstId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Name
  if len(this.Name) != 0{
    err := write(buffer,this.Name)
    if err != nil {
      return err
    }
  }
  // serialize TianTi
  {
    if(this.TianTi!=0){
      err := write(buffer,this.TianTi)
      if err != nil{
        return err
      }
    }
  }
  // serialize MainUnit
  {
    err := this.MainUnit.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize Units
  if len(this.Units) != 0{
    {
      err := write(buffer,uint(len(this.Units)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Units {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_ReportCamp)Deserialize(buffer *bytes.Buffer) error{
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
  // deserialize InstId
  if mask.readBit() {
    err := read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize Name
  if mask.readBit() {
    err := read(buffer,&this.Name)
    if err != nil{
      return err
    }
  }
  // deserialize TianTi
  if mask.readBit() {
    err := read(buffer,&this.TianTi)
    if err != nil{
      return err
    }
  }
  // deserialize MainUnit
  if mask.readBit() {
    err := this.MainUnit.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize Units
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.Units = make([]COM_BattleUnit,size)
    for i,_ := range this.Units{
      err := this.Units[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ReportCamp)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
