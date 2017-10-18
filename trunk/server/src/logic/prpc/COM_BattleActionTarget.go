package prpc
import(
  "bytes"
  "encoding/json"
  "suzuki/prpc"
)
type COM_BattleActionTarget struct{
  InstId int64  //0
  ActionType int  //1
  ActionParam int32  //2
  ActionParamExt string  //3
  Dead bool  //4
  ThrowCard COM_ThrowCard  //5
  BuffAdd []COM_BattleBuff  //6
}
func (this *COM_BattleActionTarget)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.ActionType!=0)
  mask.WriteBit(this.ActionParam!=0)
  mask.WriteBit(len(this.ActionParamExt) != 0)
  mask.WriteBit(this.Dead)
  mask.WriteBit(true) //ThrowCard
  mask.WriteBit(len(this.BuffAdd) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize InstId
  {
    if(this.InstId!=0){
      err := prpc.Write(buffer,this.InstId)
      if err != nil{
        return err
      }
    }
  }
  // serialize ActionType
  {
    if(this.ActionType!=0){
      err := prpc.Write(buffer,this.ActionType)
      if err != nil{
        return err
      }
    }
  }
  // serialize ActionParam
  {
    if(this.ActionParam!=0){
      err := prpc.Write(buffer,this.ActionParam)
      if err != nil{
        return err
      }
    }
  }
  // serialize ActionParamExt
  if len(this.ActionParamExt) != 0{
    err := prpc.Write(buffer,this.ActionParamExt)
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
      err := prpc.Write(buffer,uint(len(this.BuffAdd)))
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
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize InstId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize ActionType
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.ActionType)
    if err != nil{
      return err
    }
  }
  // deserialize ActionParam
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.ActionParam)
    if err != nil{
      return err
    }
  }
  // deserialize ActionParamExt
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.ActionParamExt)
    if err != nil{
      return err
    }
  }
  // deserialize Dead
  this.Dead = mask.ReadBit();
  // deserialize ThrowCard
  if mask.ReadBit() {
    err := this.ThrowCard.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize BuffAdd
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
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
