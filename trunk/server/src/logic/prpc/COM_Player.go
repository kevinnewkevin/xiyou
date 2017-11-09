package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_Player struct {
	sync.Mutex
	InstId           int64           //0
	Name             string          //1
	Unit             COM_Unit        //2
	Employees        []COM_Unit      //3
	Chapters         []COM_Chapter   //4
	UnitGroup        []COM_UnitGroup //5
	TianTiVal        int32           //6
	TianTiRank       int32           //7
	FriendTianTiRank int32           //8
	Guide            uint64          //9
	SkillBase        []COM_SkillBase //10
	Friends          []COM_Friend    //11
	Enemys           []COM_Friend    //12
}

func (this *COM_Player) SetInstId(value int64) {
	this.Lock()
	defer this.Unlock()
	this.InstId = value
}
func (this *COM_Player) GetInstId() int64 {
	this.Lock()
	defer this.Unlock()
	return this.InstId
}
func (this *COM_Player) SetName(value string) {
	this.Lock()
	defer this.Unlock()
	this.Name = value
}
func (this *COM_Player) GetName() string {
	this.Lock()
	defer this.Unlock()
	return this.Name
}
func (this *COM_Player) SetUnit(value COM_Unit) {
	this.Lock()
	defer this.Unlock()
	this.Unit = value
}
func (this *COM_Player) GetUnit() COM_Unit {
	this.Lock()
	defer this.Unlock()
	return this.Unit
}
func (this *COM_Player) SetEmployees(value []COM_Unit) {
	this.Lock()
	defer this.Unlock()
	this.Employees = value
}
func (this *COM_Player) GetEmployees() []COM_Unit {
	this.Lock()
	defer this.Unlock()
	return this.Employees
}
func (this *COM_Player) SetChapters(value []COM_Chapter) {
	this.Lock()
	defer this.Unlock()
	this.Chapters = value
}
func (this *COM_Player) GetChapters() []COM_Chapter {
	this.Lock()
	defer this.Unlock()
	return this.Chapters
}
func (this *COM_Player) SetUnitGroup(value []COM_UnitGroup) {
	this.Lock()
	defer this.Unlock()
	this.UnitGroup = value
}
func (this *COM_Player) GetUnitGroup() []COM_UnitGroup {
	this.Lock()
	defer this.Unlock()
	return this.UnitGroup
}
func (this *COM_Player) SetTianTiVal(value int32) {
	this.Lock()
	defer this.Unlock()
	this.TianTiVal = value
}
func (this *COM_Player) GetTianTiVal() int32 {
	this.Lock()
	defer this.Unlock()
	return this.TianTiVal
}
func (this *COM_Player) SetTianTiRank(value int32) {
	this.Lock()
	defer this.Unlock()
	this.TianTiRank = value
}
func (this *COM_Player) GetTianTiRank() int32 {
	this.Lock()
	defer this.Unlock()
	return this.TianTiRank
}
func (this *COM_Player) SetFriendTianTiRank(value int32) {
	this.Lock()
	defer this.Unlock()
	this.FriendTianTiRank = value
}
func (this *COM_Player) GetFriendTianTiRank() int32 {
	this.Lock()
	defer this.Unlock()
	return this.FriendTianTiRank
}
func (this *COM_Player) SetGuide(value uint64) {
	this.Lock()
	defer this.Unlock()
	this.Guide = value
}
func (this *COM_Player) GetGuide() uint64 {
	this.Lock()
	defer this.Unlock()
	return this.Guide
}
func (this *COM_Player) SetSkillBase(value []COM_SkillBase) {
	this.Lock()
	defer this.Unlock()
	this.SkillBase = value
}
func (this *COM_Player) GetSkillBase() []COM_SkillBase {
	this.Lock()
	defer this.Unlock()
	return this.SkillBase
}
func (this *COM_Player) SetFriends(value []COM_Friend) {
	this.Lock()
	defer this.Unlock()
	this.Friends = value
}
func (this *COM_Player) GetFriends() []COM_Friend {
	this.Lock()
	defer this.Unlock()
	return this.Friends
}
func (this *COM_Player) SetEnemys(value []COM_Friend) {
	this.Lock()
	defer this.Unlock()
	this.Enemys = value
}
func (this *COM_Player) GetEnemys() []COM_Friend {
	this.Lock()
	defer this.Unlock()
	return this.Enemys
}
func (this *COM_Player) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(2)
	mask.writeBit(this.InstId != 0)
	mask.writeBit(len(this.Name) != 0)
	mask.writeBit(true) //Unit
	mask.writeBit(len(this.Employees) != 0)
	mask.writeBit(len(this.Chapters) != 0)
	mask.writeBit(len(this.UnitGroup) != 0)
	mask.writeBit(this.TianTiVal != 0)
	mask.writeBit(this.TianTiRank != 0)
	mask.writeBit(this.FriendTianTiRank != 0)
	mask.writeBit(this.Guide != 0)
	mask.writeBit(len(this.SkillBase) != 0)
	mask.writeBit(len(this.Friends) != 0)
	mask.writeBit(len(this.Enemys) != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize InstId
	{
		if this.InstId != 0 {
			err := write(buffer, this.InstId)
			if err != nil {
				return err
			}
		}
	}
	// serialize Name
	if len(this.Name) != 0 {
		err := write(buffer, this.Name)
		if err != nil {
			return err
		}
	}
	// serialize Unit
	{
		err := this.Unit.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	// serialize Employees
	if len(this.Employees) != 0 {
		{
			err := write(buffer, uint(len(this.Employees)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.Employees {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize Chapters
	if len(this.Chapters) != 0 {
		{
			err := write(buffer, uint(len(this.Chapters)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.Chapters {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize UnitGroup
	if len(this.UnitGroup) != 0 {
		{
			err := write(buffer, uint(len(this.UnitGroup)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.UnitGroup {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize TianTiVal
	{
		if this.TianTiVal != 0 {
			err := write(buffer, this.TianTiVal)
			if err != nil {
				return err
			}
		}
	}
	// serialize TianTiRank
	{
		if this.TianTiRank != 0 {
			err := write(buffer, this.TianTiRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize FriendTianTiRank
	{
		if this.FriendTianTiRank != 0 {
			err := write(buffer, this.FriendTianTiRank)
			if err != nil {
				return err
			}
		}
	}
	// serialize Guide
	{
		if this.Guide != 0 {
			err := write(buffer, this.Guide)
			if err != nil {
				return err
			}
		}
	}
	// serialize SkillBase
	if len(this.SkillBase) != 0 {
		{
			err := write(buffer, uint(len(this.SkillBase)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.SkillBase {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize Friends
	if len(this.Friends) != 0 {
		{
			err := write(buffer, uint(len(this.Friends)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.Friends {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize Enemys
	if len(this.Enemys) != 0 {
		{
			err := write(buffer, uint(len(this.Enemys)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.Enemys {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Player) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 2)
	if err != nil {
		return err
	}
	// deserialize InstId
	if mask.readBit() {
		err := read(buffer, &this.InstId)
		if err != nil {
			return err
		}
	}
	// deserialize Name
	if mask.readBit() {
		err := read(buffer, &this.Name)
		if err != nil {
			return err
		}
	}
	// deserialize Unit
	if mask.readBit() {
		err := this.Unit.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	// deserialize Employees
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.Employees = make([]COM_Unit, size)
		for i, _ := range this.Employees {
			err := this.Employees[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize Chapters
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.Chapters = make([]COM_Chapter, size)
		for i, _ := range this.Chapters {
			err := this.Chapters[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize UnitGroup
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.UnitGroup = make([]COM_UnitGroup, size)
		for i, _ := range this.UnitGroup {
			err := this.UnitGroup[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize TianTiVal
	if mask.readBit() {
		err := read(buffer, &this.TianTiVal)
		if err != nil {
			return err
		}
	}
	// deserialize TianTiRank
	if mask.readBit() {
		err := read(buffer, &this.TianTiRank)
		if err != nil {
			return err
		}
	}
	// deserialize FriendTianTiRank
	if mask.readBit() {
		err := read(buffer, &this.FriendTianTiRank)
		if err != nil {
			return err
		}
	}
	// deserialize Guide
	if mask.readBit() {
		err := read(buffer, &this.Guide)
		if err != nil {
			return err
		}
	}
	// deserialize SkillBase
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.SkillBase = make([]COM_SkillBase, size)
		for i, _ := range this.SkillBase {
			err := this.SkillBase[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize Friends
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.Friends = make([]COM_Friend, size)
		for i, _ := range this.Friends {
			err := this.Friends[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize Enemys
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.Enemys = make([]COM_Friend, size)
		for i, _ := range this.Enemys {
			err := this.Enemys[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Player) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
