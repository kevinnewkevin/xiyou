package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleActionTarget struct{
  sync.Mutex
  InstId int64  //0
  ActionType int  //1
  ActionParam int32  //2
  ActionParamExt string  //3
  Dead bool  //4
  ThrowCard COM_ThrowCard  //5
  BuffAdd []COM_BattleBuff  //6
}
func (this *COM_BattleActionTarget)SetInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.InstId = value
}
func (this *COM_BattleActionTarget)GetInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.InstId
}
func (this *COM_BattleActionTarget)SetActionType(value int) {
  this.Lock()
  defer this.Unlock()
  this.ActionType = value
}
func (this *COM_BattleActionTarget)GetActionType() int {
  this.Lock()
  defer this.Unlock()
  return this.ActionType
}
func (this *COM_BattleActionTarget)SetActionParam(value int32) {
  this.Lock()
  defer this.Unlock()
  this.ActionParam = value
}
func (this *COM_BattleActionTarget)GetActionParam() int32 {
  this.Lock()
  defer this.Unlock()
  return this.ActionParam
}
func (this *COM_BattleActionTarget)SetActionParamExt(value string) {
  this.Lock()
  defer this.Unlock()
  this.ActionParamExt = value
}
func (this *COM_BattleActionTarget)GetActionParamExt() string {
  this.Lock()
  defer this.Unlock()
  return this.ActionParamExt
}
func (this *COM_BattleActionTarget)SetDead(value bool) {
  this.Lock()
  defer this.Unlock()
  this.Dead = value
}
func (this *COM_BattleActionTarget)GetDead() bool {
  this.Lock()
  defer this.Unlock()
  return this.Dead
}
func (this *COM_BattleActionTarget)SetThrowCard(value COM_ThrowCard) {
  this.Lock()
  defer this.Unlock()
  this.ThrowCard = value
}
func (this *COM_BattleActionTarget)GetThrowCard() COM_ThrowCard {
  this.Lock()
  defer this.Unlock()
  return this.ThrowCard
}
func (this *COM_BattleActionTarget)SetBuffAdd(value []COM_BattleBuff) {
  this.Lock()
  defer this.Unlock()
  this.BuffAdd = value
}
func (this *COM_BattleActionTarget)GetBuffAdd() []COM_BattleBuff {
  this.Lock()
  defer this.Unlock()
  return this.BuffAdd
}
func (this *COM_BattleActionTarget)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.InstId!=0)
  mask.writeBit(this.ActionType!=0)
  mask.writeBit(this.ActionParam!=0)
  mask.writeBit(len(this.ActionParamExt) != 0)
  mask.writeBit(this.Dead)
  mask.writeBit(true) //ThrowCard
  mask.writeBit(len(this.BuffAdd) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
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
  // serialize ActionType
  {
    if(this.ActionType!=0){
      err := write(buffer,this.ActionType)
      if err != nil{
        return err
      }
    }
  }
  // serialize ActionParam
  {
    if(this.ActionParam!=0){
      err := write(buffer,this.ActionParam)
      if err != nil{
        return err
      }
    }
  }
  // serialize ActionParamExt
  if len(this.ActionParamExt) != 0{
    err := write(buffer,this.ActionParamExt)
    if err != nil {
      return err
    }
  }
  // serialize Dead
  {
  }
  // serialize ThrowCard
  {
    err := this.ThrowCard.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize BuffAdd
  if len(this.BuffAdd) != 0{
    {
      err := write(buffer,uint(len(this.BuffAdd)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BuffAdd {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleActionTarget)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize InstId
  if mask.readBit() {
    err := read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize ActionType
  if mask.readBit() {
    err := read(buffer,&this.ActionType)
    if err != nil{
      return err
    }
  }
  // deserialize ActionParam
  if mask.readBit() {
    err := read(buffer,&this.ActionParam)
    if err != nil{
      return err
    }
  }
  // deserialize ActionParamExt
  if mask.readBit() {
    err := read(buffer,&this.ActionParamExt)
    if err != nil{
      return err
    }
  }
  // deserialize Dead
  this.Dead = mask.readBit();
  // deserialize ThrowCard
  if mask.readBit() {
    err := this.ThrowCard.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize BuffAdd
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.BuffAdd = make([]COM_BattleBuff,size)
    for i,_ := range this.BuffAdd{
      err := this.BuffAdd[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleActionTarget)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
