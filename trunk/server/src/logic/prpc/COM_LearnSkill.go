package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_LearnSkill struct{
  Position int32  //0
  SkillID int32  //1
}
func (this *COM_LearnSkill)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.Position!=0)
  mask.WriteBit(this.SkillID!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Position
  {
    if(this.Position!=0){
      err := prpc.Write(buffer,this.Position)
      if err != nil{
        return err
      }
    }
  }
  // serialize SkillID
  {
    if(this.SkillID!=0){
      err := prpc.Write(buffer,this.SkillID)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_LearnSkill)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Position
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Position)
    if err != nil{
      return err
    }
  }
  // deserialize SkillID
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.SkillID)
    if err != nil{
      return err
    }
  }
  return nil
}
