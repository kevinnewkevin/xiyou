package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_UnitInfo struct{
  InstId int64  //0
  UnitId int32  //1
  Level int32  //2
  HP int32  //3
  AGILE int32  //4
  ATK int32  //5
  DEF int32  //6
  MATK int32  //7
  MDEF int32  //8
}
func (this *COM_UnitInfo)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(2)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.UnitId!=0)
  mask.WriteBit(this.Level!=0)
  mask.WriteBit(this.HP!=0)
  mask.WriteBit(this.AGILE!=0)
  mask.WriteBit(this.ATK!=0)
  mask.WriteBit(this.DEF!=0)
  mask.WriteBit(this.MATK!=0)
  mask.WriteBit(this.MDEF!=0)
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
  // serialize UnitId
  {
    if(this.UnitId!=0){
      err := prpc.Write(buffer,this.UnitId)
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
  // serialize HP
  {
    if(this.HP!=0){
      err := prpc.Write(buffer,this.HP)
      if err != nil{
        return err
      }
    }
  }
  // serialize AGILE
  {
    if(this.AGILE!=0){
      err := prpc.Write(buffer,this.AGILE)
      if err != nil{
        return err
      }
    }
  }
  // serialize ATK
  {
    if(this.ATK!=0){
      err := prpc.Write(buffer,this.ATK)
      if err != nil{
        return err
      }
    }
  }
  // serialize DEF
  {
    if(this.DEF!=0){
      err := prpc.Write(buffer,this.DEF)
      if err != nil{
        return err
      }
    }
  }
  // serialize MATK
  {
    if(this.MATK!=0){
      err := prpc.Write(buffer,this.MATK)
      if err != nil{
        return err
      }
    }
  }
  // serialize MDEF
  {
    if(this.MDEF!=0){
      err := prpc.Write(buffer,this.MDEF)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_UnitInfo)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,2);
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
  // deserialize UnitId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.UnitId)
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
  // deserialize HP
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.HP)
    if err != nil{
      return err
    }
  }
  // deserialize AGILE
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.AGILE)
    if err != nil{
      return err
    }
  }
  // deserialize ATK
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.ATK)
    if err != nil{
      return err
    }
  }
  // deserialize DEF
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.DEF)
    if err != nil{
      return err
    }
  }
  // deserialize MATK
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.MATK)
    if err != nil{
      return err
    }
  }
  // deserialize MDEF
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.MDEF)
    if err != nil{
      return err
    }
  }
  return nil
}
