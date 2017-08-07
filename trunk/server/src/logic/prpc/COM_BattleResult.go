package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_BattleResult struct{
  Win int32  //0
  Money int32  //1
}
func (this *COM_BattleResult)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.Win!=0)
  mask.WriteBit(this.Money!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Win
  {
    if(this.Win!=0){
      err := prpc.Write(buffer,this.Win)
      if err != nil{
        return err
      }
    }
  }
  // serialize Money
  {
    if(this.Money!=0){
      err := prpc.Write(buffer,this.Money)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleResult)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Win
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Win)
    if err != nil{
      return err
    }
  }
  // deserialize Money
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Money)
    if err != nil{
      return err
    }
  }
  return nil
}
