package prpc
import(
  "bytes"
  "errors"
)
type COM_ClientToServer_Login struct{
  info COM_LoginInfo  //0
}
type COM_ClientToServer_CreatePlayer struct{
  tempId int32  //0
  playerName string  //1
}
type COM_ClientToServer_AddBattleUnit struct{
  instId int64  //0
  groupId int32  //1
}
type COM_ClientToServer_PopBattleUnit struct{
  instId int64  //0
  groupId int32  //1
}
type COM_ClientToServer_SetBattleUnit struct{
  instId int64  //0
  groupId int32  //1
  isBattle bool  //2
}
type COM_ClientToServer_DelUnitGroup struct{
  groupId int32  //0
}
type COM_ClientToServer_SetupBattle struct{
  positionList []COM_BattlePosition  //0
  skillid int32  //1
}
type COM_ClientToServer_RequestChapterData struct{
  chapterId int32  //0
}
type COM_ClientToServer_ChallengeSmallChapter struct{
  smallChapterId int32  //0
}
type COM_ClientToServer_RequestChapterStarReward struct{
  chapterId int32  //0
  star int32  //1
}
type COM_ClientToServer_StartMatching struct{
  groupId int32  //0
}
type COM_ClientToServer_DeleteItem struct{
  instId int64  //0
  stack int32  //1
}
type COM_ClientToServer_PromoteUnit struct{
  instId int64  //0
}
type COM_ClientToServer_EquipSkill struct{
  skillInfo COM_LearnSkill  //0
}
type COM_ClientToServer_SkillUpdate struct{
  skillIndex int32  //0
  skillID int32  //1
}
type COM_ClientToServer_BuyShopItem struct{
  shopId int32  //0
}
type COM_ClientToServer_ResolveItem struct{
  instId int64  //0
  num int32  //1
}
type COM_ClientToServer_NewPlayerGuide struct{
  Step uint64  //0
}
type COM_ClientToServer_SendChat struct{
  content COM_Chat  //0
}
type COM_ClientToServer_RequestAudio struct{
  audioId int64  //0
}
type COM_ClientToServer_SerchFriendByName struct{
  name string  //0
}
type COM_ClientToServer_ProcessingFriend struct{
  name string  //0
}
type COM_ClientToServer_DeleteFriend struct{
  instid int64  //0
}
type COM_ClientToServerStub struct{
  Sender StubSender
}
type COM_ClientToServerProxy interface{
  Login(info COM_LoginInfo ) error // 0
  CreatePlayer(tempId int32, playerName string ) error // 1
  AddBattleUnit(instId int64, groupId int32 ) error // 2
  PopBattleUnit(instId int64, groupId int32 ) error // 3
  SetBattleUnit(instId int64, groupId int32, isBattle bool ) error // 4
  DelUnitGroup(groupId int32 ) error // 5
  JoinBattle() error // 6
  SetupBattle(positionList []COM_BattlePosition, skillid int32 ) error // 7
  RequestChapterData(chapterId int32 ) error // 8
  ChallengeSmallChapter(smallChapterId int32 ) error // 9
  RequestChapterStarReward(chapterId int32, star int32 ) error // 10
  StartMatching(groupId int32 ) error // 11
  StopMatching() error // 12
  DeleteItem(instId int64, stack int32 ) error // 13
  PromoteUnit(instId int64 ) error // 14
  EquipSkill(skillInfo COM_LearnSkill ) error // 15
  SkillUpdate(skillIndex int32, skillID int32 ) error // 16
  BuyShopItem(shopId int32 ) error // 17
  ResolveItem(instId int64, num int32 ) error // 18
  RefreshBlackMarkte() error // 19
  NewPlayerGuide(Step uint64 ) error // 20
  SendChat(content COM_Chat ) error // 21
  RequestAudio(audioId int64 ) error // 22
  AllTopByPage() error // 23
  FriendTopByPage() error // 24
  SerchFriendByName(name string ) error // 25
  SerchFriendRandom() error // 26
  ProcessingFriend(name string ) error // 27
  DeleteFriend(instid int64 ) error // 28
}
func (this *COM_ClientToServer_Login)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //info
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize info
  {
    err := this.info.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_Login)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize info
  if mask.readBit() {
    err := this.info.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_CreatePlayer)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.tempId!=0)
  mask.writeBit(len(this.playerName) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize tempId
  {
    if(this.tempId!=0){
      err := write(buffer,this.tempId)
      if err != nil{
        return err
      }
    }
  }
  // serialize playerName
  if len(this.playerName) != 0{
    err := write(buffer,this.playerName)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_CreatePlayer)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize tempId
  if mask.readBit() {
    err := read(buffer,&this.tempId)
    if err != nil{
      return err
    }
  }
  // deserialize playerName
  if mask.readBit() {
    err := read(buffer,&this.playerName)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_AddBattleUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  mask.writeBit(this.groupId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_AddBattleUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize groupId
  if mask.readBit() {
    err := read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_PopBattleUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  mask.writeBit(this.groupId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_PopBattleUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize groupId
  if mask.readBit() {
    err := read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SetBattleUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  mask.writeBit(this.groupId!=0)
  mask.writeBit(this.isBattle)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  // serialize isBattle
  {
  }
  return nil
}
func (this *COM_ClientToServer_SetBattleUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize groupId
  if mask.readBit() {
    err := read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  // deserialize isBattle
  this.isBattle = mask.readBit();
  return nil
}
func (this *COM_ClientToServer_DelUnitGroup)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.groupId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_DelUnitGroup)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize groupId
  if mask.readBit() {
    err := read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SetupBattle)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.positionList) != 0)
  mask.writeBit(this.skillid!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize positionList
  if len(this.positionList) != 0{
    {
      err := write(buffer,uint(len(this.positionList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.positionList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize skillid
  {
    if(this.skillid!=0){
      err := write(buffer,this.skillid)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_SetupBattle)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize positionList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.positionList = make([]COM_BattlePosition,size)
    for i,_ := range this.positionList{
      err := this.positionList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize skillid
  if mask.readBit() {
    err := read(buffer,&this.skillid)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestChapterData)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.chapterId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize chapterId
  {
    if(this.chapterId!=0){
      err := write(buffer,this.chapterId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestChapterData)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize chapterId
  if mask.readBit() {
    err := read(buffer,&this.chapterId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_ChallengeSmallChapter)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.smallChapterId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize smallChapterId
  {
    if(this.smallChapterId!=0){
      err := write(buffer,this.smallChapterId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_ChallengeSmallChapter)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize smallChapterId
  if mask.readBit() {
    err := read(buffer,&this.smallChapterId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestChapterStarReward)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.chapterId!=0)
  mask.writeBit(this.star!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize chapterId
  {
    if(this.chapterId!=0){
      err := write(buffer,this.chapterId)
      if err != nil{
        return err
      }
    }
  }
  // serialize star
  {
    if(this.star!=0){
      err := write(buffer,this.star)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestChapterStarReward)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize chapterId
  if mask.readBit() {
    err := read(buffer,&this.chapterId)
    if err != nil{
      return err
    }
  }
  // deserialize star
  if mask.readBit() {
    err := read(buffer,&this.star)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_StartMatching)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.groupId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize groupId
  {
    if(this.groupId!=0){
      err := write(buffer,this.groupId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_StartMatching)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize groupId
  if mask.readBit() {
    err := read(buffer,&this.groupId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_DeleteItem)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  mask.writeBit(this.stack!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize stack
  {
    if(this.stack!=0){
      err := write(buffer,this.stack)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_DeleteItem)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize stack
  if mask.readBit() {
    err := read(buffer,&this.stack)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_PromoteUnit)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_PromoteUnit)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_EquipSkill)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //skillInfo
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize skillInfo
  {
    err := this.skillInfo.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_EquipSkill)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize skillInfo
  if mask.readBit() {
    err := this.skillInfo.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SkillUpdate)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.skillIndex!=0)
  mask.writeBit(this.skillID!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize skillIndex
  {
    if(this.skillIndex!=0){
      err := write(buffer,this.skillIndex)
      if err != nil{
        return err
      }
    }
  }
  // serialize skillID
  {
    if(this.skillID!=0){
      err := write(buffer,this.skillID)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_SkillUpdate)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize skillIndex
  if mask.readBit() {
    err := read(buffer,&this.skillIndex)
    if err != nil{
      return err
    }
  }
  // deserialize skillID
  if mask.readBit() {
    err := read(buffer,&this.skillID)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_BuyShopItem)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.shopId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize shopId
  {
    if(this.shopId!=0){
      err := write(buffer,this.shopId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_BuyShopItem)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize shopId
  if mask.readBit() {
    err := read(buffer,&this.shopId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_ResolveItem)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instId!=0)
  mask.writeBit(this.num!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instId
  {
    if(this.instId!=0){
      err := write(buffer,this.instId)
      if err != nil{
        return err
      }
    }
  }
  // serialize num
  {
    if(this.num!=0){
      err := write(buffer,this.num)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_ResolveItem)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instId
  if mask.readBit() {
    err := read(buffer,&this.instId)
    if err != nil{
      return err
    }
  }
  // deserialize num
  if mask.readBit() {
    err := read(buffer,&this.num)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_NewPlayerGuide)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.Step!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize Step
  {
    if(this.Step!=0){
      err := write(buffer,this.Step)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_NewPlayerGuide)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize Step
  if mask.readBit() {
    err := read(buffer,&this.Step)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SendChat)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(true) //content
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize content
  {
    err := this.content.Serialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SendChat)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize content
  if mask.readBit() {
    err := this.content.Deserialize(buffer)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestAudio)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.audioId!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize audioId
  {
    if(this.audioId!=0){
      err := write(buffer,this.audioId)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_RequestAudio)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize audioId
  if mask.readBit() {
    err := read(buffer,&this.audioId)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SerchFriendByName)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.name) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize name
  if len(this.name) != 0{
    err := write(buffer,this.name)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_SerchFriendByName)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize name
  if mask.readBit() {
    err := read(buffer,&this.name)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_ProcessingFriend)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(len(this.name) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize name
  if len(this.name) != 0{
    err := write(buffer,this.name)
    if err != nil {
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_ProcessingFriend)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize name
  if mask.readBit() {
    err := read(buffer,&this.name)
    if err != nil{
      return err
    }
  }
  return nil
}
func (this *COM_ClientToServer_DeleteFriend)Serialize(buffer *bytes.Buffer) error {
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.instid!=0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize instid
  {
    if(this.instid!=0){
      err := write(buffer,this.instid)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_ClientToServer_DeleteFriend)Deserialize(buffer *bytes.Buffer) error{
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize instid
  if mask.readBit() {
    err := read(buffer,&this.instid)
    if err != nil{
      return err
    }
  }
  return nil
}
func(this* COM_ClientToServerStub)Login(info COM_LoginInfo ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(0))
  if err != nil{
    return err
  }
  _0 := COM_ClientToServer_Login{}
  _0.info = info;
  err = _0.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)CreatePlayer(tempId int32, playerName string ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(1))
  if err != nil{
    return err
  }
  _1 := COM_ClientToServer_CreatePlayer{}
  _1.tempId = tempId;
  _1.playerName = playerName;
  err = _1.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)AddBattleUnit(instId int64, groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(2))
  if err != nil{
    return err
  }
  _2 := COM_ClientToServer_AddBattleUnit{}
  _2.instId = instId;
  _2.groupId = groupId;
  err = _2.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)PopBattleUnit(instId int64, groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(3))
  if err != nil{
    return err
  }
  _3 := COM_ClientToServer_PopBattleUnit{}
  _3.instId = instId;
  _3.groupId = groupId;
  err = _3.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SetBattleUnit(instId int64, groupId int32, isBattle bool ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(4))
  if err != nil{
    return err
  }
  _4 := COM_ClientToServer_SetBattleUnit{}
  _4.instId = instId;
  _4.groupId = groupId;
  _4.isBattle = isBattle;
  err = _4.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)DelUnitGroup(groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(5))
  if err != nil{
    return err
  }
  _5 := COM_ClientToServer_DelUnitGroup{}
  _5.groupId = groupId;
  err = _5.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)JoinBattle() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(6))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SetupBattle(positionList []COM_BattlePosition, skillid int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(7))
  if err != nil{
    return err
  }
  _7 := COM_ClientToServer_SetupBattle{}
  _7.positionList = positionList;
  _7.skillid = skillid;
  err = _7.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)RequestChapterData(chapterId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(8))
  if err != nil{
    return err
  }
  _8 := COM_ClientToServer_RequestChapterData{}
  _8.chapterId = chapterId;
  err = _8.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)ChallengeSmallChapter(smallChapterId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(9))
  if err != nil{
    return err
  }
  _9 := COM_ClientToServer_ChallengeSmallChapter{}
  _9.smallChapterId = smallChapterId;
  err = _9.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)RequestChapterStarReward(chapterId int32, star int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(10))
  if err != nil{
    return err
  }
  _10 := COM_ClientToServer_RequestChapterStarReward{}
  _10.chapterId = chapterId;
  _10.star = star;
  err = _10.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)StartMatching(groupId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(11))
  if err != nil{
    return err
  }
  _11 := COM_ClientToServer_StartMatching{}
  _11.groupId = groupId;
  err = _11.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)StopMatching() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(12))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)DeleteItem(instId int64, stack int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(13))
  if err != nil{
    return err
  }
  _13 := COM_ClientToServer_DeleteItem{}
  _13.instId = instId;
  _13.stack = stack;
  err = _13.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)PromoteUnit(instId int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(14))
  if err != nil{
    return err
  }
  _14 := COM_ClientToServer_PromoteUnit{}
  _14.instId = instId;
  err = _14.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)EquipSkill(skillInfo COM_LearnSkill ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(15))
  if err != nil{
    return err
  }
  _15 := COM_ClientToServer_EquipSkill{}
  _15.skillInfo = skillInfo;
  err = _15.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SkillUpdate(skillIndex int32, skillID int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(16))
  if err != nil{
    return err
  }
  _16 := COM_ClientToServer_SkillUpdate{}
  _16.skillIndex = skillIndex;
  _16.skillID = skillID;
  err = _16.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)BuyShopItem(shopId int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(17))
  if err != nil{
    return err
  }
  _17 := COM_ClientToServer_BuyShopItem{}
  _17.shopId = shopId;
  err = _17.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)ResolveItem(instId int64, num int32 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(18))
  if err != nil{
    return err
  }
  _18 := COM_ClientToServer_ResolveItem{}
  _18.instId = instId;
  _18.num = num;
  err = _18.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)RefreshBlackMarkte() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(19))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)NewPlayerGuide(Step uint64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(20))
  if err != nil{
    return err
  }
  _20 := COM_ClientToServer_NewPlayerGuide{}
  _20.Step = Step;
  err = _20.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SendChat(content COM_Chat ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(21))
  if err != nil{
    return err
  }
  _21 := COM_ClientToServer_SendChat{}
  _21.content = content;
  err = _21.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)RequestAudio(audioId int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(22))
  if err != nil{
    return err
  }
  _22 := COM_ClientToServer_RequestAudio{}
  _22.audioId = audioId;
  err = _22.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)AllTopByPage() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(23))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)FriendTopByPage() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(24))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SerchFriendByName(name string ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(25))
  if err != nil{
    return err
  }
  _25 := COM_ClientToServer_SerchFriendByName{}
  _25.name = name;
  err = _25.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)SerchFriendRandom() error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(26))
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)ProcessingFriend(name string ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(27))
  if err != nil{
    return err
  }
  _27 := COM_ClientToServer_ProcessingFriend{}
  _27.name = name;
  err = _27.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func(this* COM_ClientToServerStub)DeleteFriend(instid int64 ) error {
  buffer := this.Sender.MethodBegin()
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  err := write(buffer,uint16(28))
  if err != nil{
    return err
  }
  _28 := COM_ClientToServer_DeleteFriend{}
  _28.instid = instid;
  err = _28.Serialize(buffer)
  if err != nil{
    return err
  }
  return this.Sender.MethodEnd()
}
func Bridging_COM_ClientToServer_Login(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _0 := COM_ClientToServer_Login{}
  err := _0.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.Login(_0.info)
}
func Bridging_COM_ClientToServer_CreatePlayer(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _1 := COM_ClientToServer_CreatePlayer{}
  err := _1.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.CreatePlayer(_1.tempId,_1.playerName)
}
func Bridging_COM_ClientToServer_AddBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _2 := COM_ClientToServer_AddBattleUnit{}
  err := _2.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.AddBattleUnit(_2.instId,_2.groupId)
}
func Bridging_COM_ClientToServer_PopBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _3 := COM_ClientToServer_PopBattleUnit{}
  err := _3.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.PopBattleUnit(_3.instId,_3.groupId)
}
func Bridging_COM_ClientToServer_SetBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _4 := COM_ClientToServer_SetBattleUnit{}
  err := _4.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SetBattleUnit(_4.instId,_4.groupId,_4.isBattle)
}
func Bridging_COM_ClientToServer_DelUnitGroup(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _5 := COM_ClientToServer_DelUnitGroup{}
  err := _5.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DelUnitGroup(_5.groupId)
}
func Bridging_COM_ClientToServer_JoinBattle(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.JoinBattle()
}
func Bridging_COM_ClientToServer_SetupBattle(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _7 := COM_ClientToServer_SetupBattle{}
  err := _7.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SetupBattle(_7.positionList,_7.skillid)
}
func Bridging_COM_ClientToServer_RequestChapterData(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _8 := COM_ClientToServer_RequestChapterData{}
  err := _8.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RequestChapterData(_8.chapterId)
}
func Bridging_COM_ClientToServer_ChallengeSmallChapter(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _9 := COM_ClientToServer_ChallengeSmallChapter{}
  err := _9.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ChallengeSmallChapter(_9.smallChapterId)
}
func Bridging_COM_ClientToServer_RequestChapterStarReward(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _10 := COM_ClientToServer_RequestChapterStarReward{}
  err := _10.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RequestChapterStarReward(_10.chapterId,_10.star)
}
func Bridging_COM_ClientToServer_StartMatching(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _11 := COM_ClientToServer_StartMatching{}
  err := _11.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.StartMatching(_11.groupId)
}
func Bridging_COM_ClientToServer_StopMatching(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.StopMatching()
}
func Bridging_COM_ClientToServer_DeleteItem(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _13 := COM_ClientToServer_DeleteItem{}
  err := _13.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DeleteItem(_13.instId,_13.stack)
}
func Bridging_COM_ClientToServer_PromoteUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _14 := COM_ClientToServer_PromoteUnit{}
  err := _14.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.PromoteUnit(_14.instId)
}
func Bridging_COM_ClientToServer_EquipSkill(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _15 := COM_ClientToServer_EquipSkill{}
  err := _15.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.EquipSkill(_15.skillInfo)
}
func Bridging_COM_ClientToServer_SkillUpdate(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _16 := COM_ClientToServer_SkillUpdate{}
  err := _16.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SkillUpdate(_16.skillIndex,_16.skillID)
}
func Bridging_COM_ClientToServer_BuyShopItem(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _17 := COM_ClientToServer_BuyShopItem{}
  err := _17.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.BuyShopItem(_17.shopId)
}
func Bridging_COM_ClientToServer_ResolveItem(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _18 := COM_ClientToServer_ResolveItem{}
  err := _18.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ResolveItem(_18.instId,_18.num)
}
func Bridging_COM_ClientToServer_RefreshBlackMarkte(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.RefreshBlackMarkte()
}
func Bridging_COM_ClientToServer_NewPlayerGuide(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _20 := COM_ClientToServer_NewPlayerGuide{}
  err := _20.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.NewPlayerGuide(_20.Step)
}
func Bridging_COM_ClientToServer_SendChat(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _21 := COM_ClientToServer_SendChat{}
  err := _21.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SendChat(_21.content)
}
func Bridging_COM_ClientToServer_RequestAudio(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _22 := COM_ClientToServer_RequestAudio{}
  err := _22.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.RequestAudio(_22.audioId)
}
func Bridging_COM_ClientToServer_AllTopByPage(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.AllTopByPage()
}
func Bridging_COM_ClientToServer_FriendTopByPage(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.FriendTopByPage()
}
func Bridging_COM_ClientToServer_SerchFriendByName(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _25 := COM_ClientToServer_SerchFriendByName{}
  err := _25.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.SerchFriendByName(_25.name)
}
func Bridging_COM_ClientToServer_SerchFriendRandom(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  return p.SerchFriendRandom()
}
func Bridging_COM_ClientToServer_ProcessingFriend(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _27 := COM_ClientToServer_ProcessingFriend{}
  err := _27.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.ProcessingFriend(_27.name)
}
func Bridging_COM_ClientToServer_DeleteFriend(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil{
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  _28 := COM_ClientToServer_DeleteFriend{}
  err := _28.Deserialize(buffer)
  if err != nil{
    return err
  }
  return p.DeleteFriend(_28.instid)
}
func COM_ClientToServerDispatch(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
  if buffer == nil {
    return errors.New(NoneBufferError)
  }
  if p == nil {
    return errors.New(NoneProxyError)
  }
  pid := uint16(0XFFFF)
  err := read(buffer,&pid)
  if err != nil{
    return err
  }
  switch(pid){
    case 0 :
      return Bridging_COM_ClientToServer_Login(buffer,p);
    case 1 :
      return Bridging_COM_ClientToServer_CreatePlayer(buffer,p);
    case 2 :
      return Bridging_COM_ClientToServer_AddBattleUnit(buffer,p);
    case 3 :
      return Bridging_COM_ClientToServer_PopBattleUnit(buffer,p);
    case 4 :
      return Bridging_COM_ClientToServer_SetBattleUnit(buffer,p);
    case 5 :
      return Bridging_COM_ClientToServer_DelUnitGroup(buffer,p);
    case 6 :
      return Bridging_COM_ClientToServer_JoinBattle(buffer,p);
    case 7 :
      return Bridging_COM_ClientToServer_SetupBattle(buffer,p);
    case 8 :
      return Bridging_COM_ClientToServer_RequestChapterData(buffer,p);
    case 9 :
      return Bridging_COM_ClientToServer_ChallengeSmallChapter(buffer,p);
    case 10 :
      return Bridging_COM_ClientToServer_RequestChapterStarReward(buffer,p);
    case 11 :
      return Bridging_COM_ClientToServer_StartMatching(buffer,p);
    case 12 :
      return Bridging_COM_ClientToServer_StopMatching(buffer,p);
    case 13 :
      return Bridging_COM_ClientToServer_DeleteItem(buffer,p);
    case 14 :
      return Bridging_COM_ClientToServer_PromoteUnit(buffer,p);
    case 15 :
      return Bridging_COM_ClientToServer_EquipSkill(buffer,p);
    case 16 :
      return Bridging_COM_ClientToServer_SkillUpdate(buffer,p);
    case 17 :
      return Bridging_COM_ClientToServer_BuyShopItem(buffer,p);
    case 18 :
      return Bridging_COM_ClientToServer_ResolveItem(buffer,p);
    case 19 :
      return Bridging_COM_ClientToServer_RefreshBlackMarkte(buffer,p);
    case 20 :
      return Bridging_COM_ClientToServer_NewPlayerGuide(buffer,p);
    case 21 :
      return Bridging_COM_ClientToServer_SendChat(buffer,p);
    case 22 :
      return Bridging_COM_ClientToServer_RequestAudio(buffer,p);
    case 23 :
      return Bridging_COM_ClientToServer_AllTopByPage(buffer,p);
    case 24 :
      return Bridging_COM_ClientToServer_FriendTopByPage(buffer,p);
    case 25 :
      return Bridging_COM_ClientToServer_SerchFriendByName(buffer,p);
    case 26 :
      return Bridging_COM_ClientToServer_SerchFriendRandom(buffer,p);
    case 27 :
      return Bridging_COM_ClientToServer_ProcessingFriend(buffer,p);
    case 28 :
      return Bridging_COM_ClientToServer_DeleteFriend(buffer,p);
    default:
      return errors.New(NoneDispatchMatchError)
  }
}
