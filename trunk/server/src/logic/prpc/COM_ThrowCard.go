package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_ThrowCard struct{
  InstId int64  //0
  EntityId int32  //1
  Level int32  //2
}
func (this *COM_ThrowCard)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.EntityId!=0)
  mask.WriteBit(this.Level!=0)
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
  // serialize EntityId
  {
    if(this.EntityId!=0){
      err := prpc.Write(buffer,this.EntityId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Level
  {
    if(this.Level!=0){
      err := prpc.Write(buffer,this.Level)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ThrowCard)Deserialize(buffer *bytes.Buffer) error{
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
  // deserialize EntityId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.EntityId)
    if err != nil{
      return err
    }
  }
  // deserialize Level
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Level)
    if err != nil{
      return err
    }
  }
  return nil
}
