package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_Unit struct {
	sync.Mutex
	UnitId      int32           //0
	InstId      int64           //1
	Level       int32           //2
	Race        int32           //3
	IProperties []int32         //4
	CProperties []float32       //5
	Equipments  []COM_ItemInst  //6
	Skills      []COM_UnitSkill //7
}

func (this *COM_Unit) SetUnitId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.UnitId = value
}
func (this *COM_Unit) GetUnitId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.UnitId
}
func (this *COM_Unit) SetInstId(value int64) {
	this.Lock()
	defer this.Unlock()
	this.InstId = value
}
func (this *COM_Unit) GetInstId() int64 {
	this.Lock()
	defer this.Unlock()
	return this.InstId
}
func (this *COM_Unit) SetLevel(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Level = value
}
func (this *COM_Unit) GetLevel() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Level
}
func (this *COM_Unit) SetRace(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Race = value
}
func (this *COM_Unit) GetRace() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Race
}
func (this *COM_Unit) SetIProperties(value []int32) {
	this.Lock()
	defer this.Unlock()
	this.IProperties = value
}
func (this *COM_Unit) GetIProperties() []int32 {
	this.Lock()
	defer this.Unlock()
	return this.IProperties
}
func (this *COM_Unit) SetCProperties(value []float32) {
	this.Lock()
	defer this.Unlock()
	this.CProperties = value
}
func (this *COM_Unit) GetCProperties() []float32 {
	this.Lock()
	defer this.Unlock()
	return this.CProperties
}
func (this *COM_Unit) SetEquipments(value []COM_ItemInst) {
	this.Lock()
	defer this.Unlock()
	this.Equipments = value
}
func (this *COM_Unit) GetEquipments() []COM_ItemInst {
	this.Lock()
	defer this.Unlock()
	return this.Equipments
}
func (this *COM_Unit) SetSkills(value []COM_UnitSkill) {
	this.Lock()
	defer this.Unlock()
	this.Skills = value
}
func (this *COM_Unit) GetSkills() []COM_UnitSkill {
	this.Lock()
	defer this.Unlock()
	return this.Skills
}
func (this *COM_Unit) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.UnitId != 0)
	mask.writeBit(this.InstId != 0)
	mask.writeBit(this.Level != 0)
	mask.writeBit(this.Race != 0)
	mask.writeBit(len(this.IProperties) != 0)
	mask.writeBit(len(this.CProperties) != 0)
	mask.writeBit(len(this.Equipments) != 0)
	mask.writeBit(len(this.Skills) != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize UnitId
	{
		if this.UnitId != 0 {
			err := write(buffer, this.UnitId)
			if err != nil {
				return err
			}
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
	// serialize Level
	{
		if this.Level != 0 {
			err := write(buffer, this.Level)
			if err != nil {
				return err
			}
		}
	}
	// serialize Race
	{
		if this.Race != 0 {
			err := write(buffer, this.Race)
			if err != nil {
				return err
			}
		}
	}
	// serialize IProperties
	if len(this.IProperties) != 0 {
		{
			err := write(buffer, uint(len(this.IProperties)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.IProperties {
			err := write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize CProperties
	if len(this.CProperties) != 0 {
		{
			err := write(buffer, uint(len(this.CProperties)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.CProperties {
			err := write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize Equipments
	if len(this.Equipments) != 0 {
		{
			err := write(buffer, uint(len(this.Equipments)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.Equipments {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize Skills
	if len(this.Skills) != 0 {
		{
			err := write(buffer, uint(len(this.Skills)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.Skills {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Unit) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize UnitId
	if mask.readBit() {
		err := read(buffer, &this.UnitId)
		if err != nil {
			return err
		}
	}
	// deserialize InstId
	if mask.readBit() {
		err := read(buffer, &this.InstId)
		if err != nil {
			return err
		}
	}
	// deserialize Level
	if mask.readBit() {
		err := read(buffer, &this.Level)
		if err != nil {
			return err
		}
	}
	// deserialize Race
	if mask.readBit() {
		err := read(buffer, &this.Race)
		if err != nil {
			return err
		}
	}
	// deserialize IProperties
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.IProperties = make([]int32, size)
		for i, _ := range this.IProperties {
			err := read(buffer, &this.IProperties[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize CProperties
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.CProperties = make([]float32, size)
		for i, _ := range this.CProperties {
			err := read(buffer, &this.CProperties[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize Equipments
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.Equipments = make([]COM_ItemInst, size)
		for i, _ := range this.Equipments {
			err := this.Equipments[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize Skills
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.Skills = make([]COM_UnitSkill, size)
		for i, _ := range this.Skills {
			err := this.Skills[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Unit) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
