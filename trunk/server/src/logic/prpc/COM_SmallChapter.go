package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_SmallChapter struct{
  SmallChapterId int32  //0
  Star1 bool  //1
  Star2 bool  //2
  Star3 bool  //3
}
func (this *COM_SmallChapter)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.SmallChapterId!=0)
  mask.WriteBit(this.Star1)
  mask.WriteBit(this.Star2)
  mask.WriteBit(this.Star3)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize SmallChapterId
  {
    if(this.SmallChapterId!=0){
      err := prpc.Write(buffer,this.SmallChapterId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Star1
  {
  }
  // serialize Star2
  {
  }
  // serialize Star3
  {
  }
  return nil
}
func (this *COM_SmallChapter)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize SmallChapterId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.SmallChapterId)
    if err != nil{
      return err
    }
  }
  // deserialize Star1
  this.Star1 = mask.ReadBit();
  // deserialize Star2
  this.Star2 = mask.ReadBit();
  // deserialize Star3
  this.Star3 = mask.ReadBit();
  return nil
}
