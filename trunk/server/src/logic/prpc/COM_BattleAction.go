package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_BattleAction struct{
  InstId int64  //0
  SkillId int32  //1
  TargetList []COM_BattleActionTarget  //2
}
func (this *COM_BattleAction)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.SkillId!=0)
  mask.WriteBit(len(this.TargetList) != 0)
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
  // serialize SkillId
  {
    if(this.SkillId!=0){
      err := prpc.Write(buffer,this.SkillId)
      if err != nil{
        return err
      }
    }
  }
  // serialize TargetList
  if len(this.TargetList) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.TargetList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.TargetList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleAction)Deserialize(buffer *bytes.Buffer) error{
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
  // deserialize SkillId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.SkillId)
    if err != nil{
      return err
    }
  }
  // deserialize TargetList
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.TargetList = make([]COM_BattleActionTarget,size)
    for i,_ := range this.TargetList{
      err := this.TargetList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
