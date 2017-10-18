package prpc
import(
  "bytes"
  "encoding/json"
  "suzuki/prpc"
)
type SGE_DBPlayer struct{
  COM_Player
  PlayerId int64  //0
  Username string  //1
  BattleGroupIdx int32  //2
  BagItemList []COM_ItemInst  //3
}
func (this *SGE_DBPlayer)Serialize(buffer *bytes.Buffer) error {
  {
    err := this.COM_Player.Serialize(buffer);
    if err != nil {
      return err
    }
  }
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.PlayerId!=0)
  mask.WriteBit(len(this.Username) != 0)
  mask.WriteBit(this.BattleGroupIdx!=0)
  mask.WriteBit(len(this.BagItemList) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize PlayerId
  {
    if(this.PlayerId!=0){
      err := prpc.Write(buffer,this.PlayerId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Username
  if len(this.Username) != 0{
    err := prpc.Write(buffer,this.Username)
    if err != nil {
      return err
    }
  }
  // serialize BattleGroupIdx
  {
    if(this.BattleGroupIdx!=0){
      err := prpc.Write(buffer,this.BattleGroupIdx)
      if err != nil{
        return err
      }
    }
  }
  // serialize BagItemList
  if len(this.BagItemList) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.BagItemList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BagItemList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *SGE_DBPlayer)Deserialize(buffer *bytes.Buffer) error{
  {
    this.COM_Player.Deserialize(buffer);
  }
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize PlayerId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.PlayerId)
    if err != nil{
      return err
    }
  }
  // deserialize Username
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.Username)
    if err != nil{
      return err
    }
  }
  // deserialize BattleGroupIdx
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.BattleGroupIdx)
    if err != nil{
      return err
    }
  }
  // deserialize BagItemList
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.BagItemList = make([]COM_ItemInst,size)
    for i,_ := range this.BagItemList{
      err := this.BagItemList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *SGE_DBPlayer)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
