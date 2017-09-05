package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_ItemInst struct{
  ItemId int32  //0
  InstId int64  //1
  Stack_ int32  //2
}
func (this *COM_ItemInst)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.ItemId!=0)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.Stack_!=0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize ItemId
  {
    if(this.ItemId!=0){
      err := prpc.Write(buffer,this.ItemId)
      if err != nil{
        return err
      }
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
  // serialize Stack_
  {
    if(this.Stack_!=0){
      err := prpc.Write(buffer,this.Stack_)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ItemInst)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize ItemId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.ItemId)
    if err != nil{
      return err
    }
  }
  // deserialize InstId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize Stack_
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Stack_)
    if err != nil{
      return err
    }
  }
  return nil
}
