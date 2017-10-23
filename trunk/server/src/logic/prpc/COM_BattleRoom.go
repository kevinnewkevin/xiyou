package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleRoom struct{
  sync.Mutex
  InstId int64  //0
  Status int32  //1
  PlayerList []COM_BattlePlayer  //2
  Target COM_BattlePlayer  //3
  Bout int32  //4
  TurnMove int32  //5
  NextPlayer COM_BattlePlayer  //6
}
func (this *COM_BattleRoom)SetInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.InstId = value
}
func (this *COM_BattleRoom)GetInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.InstId
}
func (this *COM_BattleRoom)SetStatus(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Status = value
}
func (this *COM_BattleRoom)GetStatus() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Status
}
func (this *COM_BattleRoom)SetPlayerList(value []COM_BattlePlayer) {
  this.Lock()
  defer this.Unlock()
  this.PlayerList = value
}
func (this *COM_BattleRoom)GetPlayerList() []COM_BattlePlayer {
  this.Lock()
  defer this.Unlock()
  return this.PlayerList
}
func (this *COM_BattleRoom)SetTarget(value COM_BattlePlayer) {
  this.Lock()
  defer this.Unlock()
  this.Target = value
}
func (this *COM_BattleRoom)GetTarget() COM_BattlePlayer {
  this.Lock()
  defer this.Unlock()
  return this.Target
}
func (this *COM_BattleRoom)SetBout(value int32) {
  this.Lock()
  defer this.Unlock()
  this.Bout = value
}
func (this *COM_BattleRoom)GetBout() int32 {
  this.Lock()
  defer this.Unlock()
  return this.Bout
}
func (this *COM_BattleRoom)SetTurnMove(value int32) {
  this.Lock()
  defer this.Unlock()
  this.TurnMove = value
}
func (this *COM_BattleRoom)GetTurnMove() int32 {
  this.Lock()
  defer this.Unlock()
  return this.TurnMove
}
func (this *COM_BattleRoom)SetNextPlayer(value COM_BattlePlayer) {
  this.Lock()
  defer this.Unlock()
  this.NextPlayer = value
}
func (this *COM_BattleRoom)GetNextPlayer() COM_BattlePlayer {
  this.Lock()
  defer this.Unlock()
  return this.NextPlayer
}
func (this *COM_BattleRoom)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := NewMask1(1)
  mask.WriteBit(this.InstId!=0)
  mask.WriteBit(this.Status!=0)
  mask.WriteBit(len(this.PlayerList) != 0)
  mask.WriteBit(true) //Target
  mask.WriteBit(this.Bout!=0)
  mask.WriteBit(this.TurnMove!=0)
  mask.WriteBit(true) //NextPlayer
  {
    err := Write(buffer,mask.Bytes())
    if err != nil {
      return err
    }
  }
  // serialize InstId
  {
    if(this.InstId!=0){
      err := Write(buffer,this.InstId)
      if err != nil{
        return err
      }
    }
  }
  // serialize Status
  {
    if(this.Status!=0){
      err := Write(buffer,this.Status)
      if err != nil{
        return err
      }
    }
  }
  // serialize PlayerList
  if len(this.PlayerList) != 0{
    {
      err := Write(buffer,uint(len(this.PlayerList)))
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
      err := Write(buffer,this.Bout)
      if err != nil{
        return err
      }
    }
  }
  // serialize TurnMove
  {
    if(this.TurnMove!=0){
      err := Write(buffer,this.TurnMove)
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
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= NewMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize InstId
  if mask.ReadBit() {
    err := Read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize Status
  if mask.ReadBit() {
    err := Read(buffer,&this.Status)
    if err != nil{
      return err
    }
  }
  // deserialize PlayerList
  if mask.ReadBit() {
    var size uint
    err := Read(buffer,&size)
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
    err := Read(buffer,&this.Bout)
    if err != nil{
      return err
    }
  }
  // deserialize TurnMove
  if mask.ReadBit() {
    err := Read(buffer,&this.TurnMove)
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
func (this *COM_BattleRoom)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
