package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_Unit struct{
  UnitId int32  //0
  InstId int64  //1
  Level int32  //2
  IProperties []int32  //3
  CProperties []float32  //4
  Equipments []COM_ItemInst  //5
}
func (this *COM_Unit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.UnitId!=0)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.Level!=0)
  mask.WriteBit(len(this.IProperties) != 0)
  mask.WriteBit(len(this.CProperties) != 0)
  mask.WriteBit(len(this.Equipments) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize UnitId
  {
    if(this.UnitId!=0){
      err := prpc.Write(buffer,this.UnitId)
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
  // serialize Level
  {
    if(this.Level!=0){
      err := prpc.Write(buffer,this.Level)
      if err != nil{
        return err
      }
    }
  }
  // serialize IProperties
  if len(this.IProperties) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.IProperties)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.IProperties {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  // serialize CProperties
  if len(this.CProperties) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.CProperties)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.CProperties {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  // serialize Equipments
  if len(this.Equipments) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.Equipments)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.Equipments {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_Unit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize UnitId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.UnitId)
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
  // deserialize Level
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Level)
    if err != nil{
      return err
    }
  }
  // deserialize IProperties
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.IProperties = make([]int32,size)
    for i,_ := range this.IProperties{
      err := prpc.Read(buffer,&this.IProperties[i])
      if err != nil{
        return err
      }
    }
  }
  // deserialize CProperties
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.CProperties = make([]float32,size)
    for i,_ := range this.CProperties{
      err := prpc.Read(buffer,&this.CProperties[i])
      if err != nil{
        return err
      }
    }
  }
  // deserialize Equipments
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.Equipments = make([]COM_ItemInst,size)
    for i,_ := range this.Equipments{
      err := this.Equipments[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
