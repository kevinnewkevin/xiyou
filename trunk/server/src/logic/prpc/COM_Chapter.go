package prpc
import(
  "bytes"
  "encoding/json"
  "suzuki/prpc"
)
type COM_Chapter struct{
  ChapterId int32  //0
  SmallChapters []COM_SmallChapter  //1
  StarReward []int32  //2
}
func (this *COM_Chapter)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.ChapterId!=0)
  mask.WriteBit(len(this.SmallChapters) != 0)
  mask.WriteBit(len(this.StarReward) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize ChapterId
  {
    if(this.ChapterId!=0){
      err := prpc.Write(buffer,this.ChapterId)
      if err != nil{
        return err
      }
    }
  }
  // serialize SmallChapters
  if len(this.SmallChapters) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.SmallChapters)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.SmallChapters {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize StarReward
  if len(this.StarReward) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.StarReward)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.StarReward {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_Chapter)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize ChapterId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.ChapterId)
    if err != nil{
      return err
    }
  }
  // deserialize SmallChapters
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.SmallChapters = make([]COM_SmallChapter,size)
    for i,_ := range this.SmallChapters{
      err := this.SmallChapters[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize StarReward
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.StarReward = make([]int32,size)
    for i,_ := range this.StarReward{
      err := prpc.Read(buffer,&this.StarReward[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_Chapter)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
