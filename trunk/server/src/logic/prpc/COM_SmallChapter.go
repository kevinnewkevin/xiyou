package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_SmallChapter struct{
  sync.Mutex
  SmallChapterId int32  //0
  Star1 bool  //1
  Star2 bool  //2
  Star3 bool  //3
}
func (this *COM_SmallChapter)SetSmallChapterId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.SmallChapterId = value
}
func (this *COM_SmallChapter)GetSmallChapterId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.SmallChapterId
}
func (this *COM_SmallChapter)SetStar1(value bool) {
  this.Lock()
  defer this.Unlock()
  this.Star1 = value
}
func (this *COM_SmallChapter)GetStar1() bool {
  this.Lock()
  defer this.Unlock()
  return this.Star1
}
func (this *COM_SmallChapter)SetStar2(value bool) {
  this.Lock()
  defer this.Unlock()
  this.Star2 = value
}
func (this *COM_SmallChapter)GetStar2() bool {
  this.Lock()
  defer this.Unlock()
  return this.Star2
}
func (this *COM_SmallChapter)SetStar3(value bool) {
  this.Lock()
  defer this.Unlock()
  this.Star3 = value
}
func (this *COM_SmallChapter)GetStar3() bool {
  this.Lock()
  defer this.Unlock()
  return this.Star3
}
func (this *COM_SmallChapter)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.SmallChapterId!=0)
  mask.writeBit(this.Star1)
  mask.writeBit(this.Star2)
  mask.writeBit(this.Star3)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize SmallChapterId
  {
    if(this.SmallChapterId!=0){
      err := write(buffer,this.SmallChapterId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Star1
  {
  }
  // serialize Star2
  {
  }
  // serialize Star3
  {
  }
  return nil
}
func (this *COM_SmallChapter)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize SmallChapterId
  if mask.readBit() {
    err := read(buffer,&this.SmallChapterId)
    if err != nil{
      return err
    }
  }
  // deserialize Star1
  this.Star1 = mask.readBit();
  // deserialize Star2
  this.Star2 = mask.readBit();
  // deserialize Star3
  this.Star3 = mask.readBit();
  return nil
}
func (this *COM_SmallChapter)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
