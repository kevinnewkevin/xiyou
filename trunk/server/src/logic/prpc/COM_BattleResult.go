package prpc
import(
  "bytes"
  "suzuki/prpc"
)
type COM_BattleResult struct{
  Win int32  //0
  Money int32  //1
  KillMonsters []int32  //2
  BattleRound int32  //3
  MySelfDeathNum int32  //4
  BattleItems []COM_ItemInst  //5
}
func (this *COM_BattleResult)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.Win!=0)
  mask.WriteBit(this.Money!=0)
  mask.WriteBit(len(this.KillMonsters) != 0)
  mask.WriteBit(this.BattleRound!=0)
  mask.WriteBit(this.MySelfDeathNum!=0)
  mask.WriteBit(len(this.BattleItems) != 0)
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
  // serialize KillMonsters
  if len(this.KillMonsters) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.KillMonsters)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.KillMonsters {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  // serialize BattleRound
  {
    if(this.BattleRound!=0){
      err := prpc.Write(buffer,this.BattleRound)
      if err != nil{
        return err
      }
    }
  }
  // serialize MySelfDeathNum
  {
    if(this.MySelfDeathNum!=0){
      err := prpc.Write(buffer,this.MySelfDeathNum)
      if err != nil{
        return err
      }
    }
  }
  // serialize BattleItems
  if len(this.BattleItems) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.BattleItems)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BattleItems {
      err := value.Serialize(buffer)
      if err != nil {
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
  // deserialize KillMonsters
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.KillMonsters = make([]int32,size)
    for i,_ := range this.KillMonsters{
      err := prpc.Read(buffer,&this.KillMonsters[i])
      if err != nil{
        return err
      }
    }
  }
  // deserialize BattleRound
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.BattleRound)
    if err != nil{
      return err
    }
  }
  // deserialize MySelfDeathNum
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.MySelfDeathNum)
    if err != nil{
      return err
    }
  }
  // deserialize BattleItems
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.BattleItems = make([]COM_ItemInst,size)
    for i,_ := range this.BattleItems{
      err := this.BattleItems[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
