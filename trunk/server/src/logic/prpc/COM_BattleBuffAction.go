package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_BattleBuffAction struct{
  BuffData int32  //0
  Dead bool  //1
  BuffChange COM_BattleBuff  //2
}
func (this *COM_BattleBuffAction)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.BuffData!=0)
  mask.WriteBit(this.Dead)
  mask.WriteBit(true) //BuffChange
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize BuffData
  {
    if(this.BuffData!=0){
      err := prpc.Write(buffer,this.BuffData)
      if err != nil{
        return err
      }
    }
  }
  // serialize Dead
  {
  }
  // serialize BuffChange
  {
    err := this.BuffChange.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_BattleBuffAction)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize BuffData
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.BuffData)
    if err != nil{
      return err
    }
  }
  // deserialize Dead
  this.Dead = mask.ReadBit();
  // deserialize BuffChange
  if mask.ReadBit() {
    err := this.BuffChange.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
