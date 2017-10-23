package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_LearnSkill struct{
  sync.Mutex
  Position int32  //0
  SkillID int32  //1
}
func (this *COM_LearnSkill)SetPosition(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Position = value
}
func (this *COM_LearnSkill)GetPosition() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Position
}
func (this *COM_LearnSkill)SetSkillID(value int32) {
  this.Lock()
  defer this.Unlock()
  this.SkillID = value
}
func (this *COM_LearnSkill)GetSkillID() int32 {
  this.Lock()
  defer this.Unlock()
  return this.SkillID
}
func (this *COM_LearnSkill)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := NewMask1(1)
  mask.WriteBit(this.Position!=0)
  mask.WriteBit(this.SkillID!=0)
  {
    err := Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize Position
  {
    if(this.Position!=0){
      err := Write(buffer,this.Position)
      if err != nil{
        return err
      }
    }
  }
  // serialize SkillID
  {
    if(this.SkillID!=0){
      err := Write(buffer,this.SkillID)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_LearnSkill)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Position
  if mask.ReadBit() {
    err := Read(buffer,&this.Position)
    if err != nil{
      return err
    }
  }
  // deserialize SkillID
  if mask.ReadBit() {
    err := Read(buffer,&this.SkillID)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_LearnSkill)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
