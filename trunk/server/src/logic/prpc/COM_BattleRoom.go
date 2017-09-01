package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_BattleRoom struct{
  InstId int64  //0
  Status int32  //1
  PlayerList []COM_BattlePlayer  //2
  Target COM_BattlePlayer  //3
  Bout int32  //4
  TurnMove int32  //5
  NextPlayer COM_BattlePlayer  //6
}
func (this *COM_BattleRoom)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.Status!=0)
  mask.WriteBit(len(this.PlayerList) != 0)
  mask.WriteBit(true) //Target
  mask.WriteBit(this.Bout!=0)
  mask.WriteBit(this.TurnMove!=0)
  mask.WriteBit(true) //NextPlayer
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
  // serialize Status
  {
    if(this.Status!=0){
      err := prpc.Write(buffer,this.Status)
      if err != nil{
        return err
      }
    }
  }
  // serialize PlayerList
  if len(this.PlayerList) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.PlayerList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.PlayerList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize Target
  {
    err := this.Target.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  // serialize Bout
  {
    if(this.Bout!=0){
      err := prpc.Write(buffer,this.Bout)
      if err != nil{
        return err
      }
    }
  }
  // serialize TurnMove
  {
    if(this.TurnMove!=0){
      err := prpc.Write(buffer,this.TurnMove)
      if err != nil{
        return err
      }
    }
  }
  // serialize NextPlayer
  {
    err := this.NextPlayer.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_BattleRoom)Deserialize(buffer *bytes.Buffer) error{
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
  // deserialize Status
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Status)
    if err != nil{
      return err
    }
  }
  // deserialize PlayerList
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.PlayerList = make([]COM_BattlePlayer,size)
    for i,_ := range this.PlayerList{
      err := this.PlayerList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize Target
  if mask.ReadBit() {
    err := this.Target.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  // deserialize Bout
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Bout)
    if err != nil{
      return err
    }
  }
  // deserialize TurnMove
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.TurnMove)
    if err != nil{
      return err
    }
  }
  // deserialize NextPlayer
  if mask.ReadBit() {
    err := this.NextPlayer.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
