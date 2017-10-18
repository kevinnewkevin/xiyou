package prpc
import(
  "bytes"
  "encoding/json"
  "suzuki/prpc"
)
type COM_UnitGroup struct{
  GroupId int32  //0
  UnitList []int64  //1
}
func (this *COM_UnitGroup)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := prpc.NewMask1(1)
  mask.WriteBit(this.GroupId!=0)
  mask.WriteBit(len(this.UnitList) != 0)
  {
    err := prpc.Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize GroupId
  {
    if(this.GroupId!=0){
      err := prpc.Write(buffer,this.GroupId)
      if err != nil{
        return err
      }
    }
  }
  // serialize UnitList
  if len(this.UnitList) != 0{
    {
      err := prpc.Write(buffer,uint(len(this.UnitList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.UnitList {
      err := prpc.Write(buffer,value)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_UnitGroup)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= prpc.NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize GroupId
  if mask.ReadBit() {
    err := prpc.Read(buffer,&this.GroupId)
    if err != nil{
      return err
    }
  }
  // deserialize UnitList
  if mask.ReadBit() {
    var size uint
    err := prpc.Read(buffer,&size)
    if err != nil{
      return err
    }
    this.UnitList = make([]int64,size)
    for i,_ := range this.UnitList{
      err := prpc.Read(buffer,&this.UnitList[i])
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_UnitGroup)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
