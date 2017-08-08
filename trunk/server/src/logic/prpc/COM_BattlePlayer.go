package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_BattlePlayer struct{
  Player COM_Player  //0
  MaxPoint int32  //1
  CurPoint int32  //2
  BattlePosition []COM_BattlePosition  //3
}
func (this *COM_BattlePlayer)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(true) //Player
  mask.WriteBit(this.MaxPoint!=0)
  mask.WriteBit(this.CurPoint!=0)
  mask.WriteBit(len(this.BattlePosition) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Player
  {
    err := this.Player.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize MaxPoint
  {
    if(this.MaxPoint!=0){
      err := prpc.Write(buffer,this.MaxPoint)
      if err != nil{
        return err
      }
    }
  }
  // serialize CurPoint
  {
    if(this.CurPoint!=0){
      err := prpc.Write(buffer,this.CurPoint)
      if err != nil{
        return err
      }
    }
  }
  // serialize BattlePosition
  if len(this.BattlePosition) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.BattlePosition)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BattlePosition {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattlePlayer)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Player
  if mask.ReadBit() {
    err := this.Player.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize MaxPoint
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.MaxPoint)
    if err != nil{
      return err
    }
  }
  // deserialize CurPoint
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.CurPoint)
    if err != nil{
      return err
    }
  }
  // deserialize BattlePosition
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.BattlePosition = make([]COM_BattlePosition,size)
    for i,_ := range this.BattlePosition{
      err := this.BattlePosition[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
