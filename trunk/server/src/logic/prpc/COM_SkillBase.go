package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_SkillBase struct{
  SkillIdx int32  //0
  SkillId int32  //1
}
func (this *COM_SkillBase)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.SkillIdx!=0)
  mask.WriteBit(this.SkillId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize SkillIdx
  {
    if(this.SkillIdx!=0){
      err := prpc.Write(buffer,this.SkillIdx)
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
  return nil
}
func (this *COM_SkillBase)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize SkillIdx
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.SkillIdx)
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
  return nil
}
