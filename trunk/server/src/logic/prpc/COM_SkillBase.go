package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_SkillBase struct{
  sync.Mutex
  SkillIdx int32  //0
  SkillId int32  //1
}
func (this *COM_SkillBase)SetSkillIdx(value int32) {
  this.Lock()
  defer this.Unlock()
  this.SkillIdx = value
}
func (this *COM_SkillBase)GetSkillIdx() int32 {
  this.Lock()
  defer this.Unlock()
  return this.SkillIdx
}
func (this *COM_SkillBase)SetSkillId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.SkillId = value
}
func (this *COM_SkillBase)GetSkillId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.SkillId
}
func (this *COM_SkillBase)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.SkillIdx!=0)
  mask.writeBit(this.SkillId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize SkillIdx
  {
    if(this.SkillIdx!=0){
      err := write(buffer,this.SkillIdx)
      if err != nil{
        return err
      }
    }
  }
  // serialize SkillId
  {
    if(this.SkillId!=0){
      err := write(buffer,this.SkillId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_SkillBase)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize SkillIdx
  if mask.readBit() {
    err := read(buffer,&this.SkillIdx)
    if err != nil{
      return err
    }
  }
  // deserialize SkillId
  if mask.readBit() {
    err := read(buffer,&this.SkillId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_SkillBase)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
