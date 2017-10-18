package prpc
import(
  "bytes"
  "encoding/json"
  "suzuki/prpc"
)
type COM_UnitSkill struct{
  Pos int32  //0
  SkillId int32  //1
}
func (this *COM_UnitSkill)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.Pos!=0)
  mask.WriteBit(this.SkillId!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Pos
  {
    if(this.Pos!=0){
      err := prpc.Write(buffer,this.Pos)
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
func (this *COM_UnitSkill)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Pos
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Pos)
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
func (this *COM_UnitSkill)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
