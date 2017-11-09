package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_UnitSkill struct {
	sync.Mutex
	Pos     int32 //0
	SkillId int32 //1
}

func (this *COM_UnitSkill) SetPos(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Pos = value
}
func (this *COM_UnitSkill) GetPos() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Pos
}
func (this *COM_UnitSkill) SetSkillId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.SkillId = value
}
func (this *COM_UnitSkill) GetSkillId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.SkillId
}
func (this *COM_UnitSkill) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.Pos != 0)
	mask.writeBit(this.SkillId != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize Pos
	{
		if this.Pos != 0 {
			err := write(buffer, this.Pos)
			if err != nil {
				return err
			}
		}
	}
	// serialize SkillId
	{
		if this.SkillId != 0 {
			err := write(buffer, this.SkillId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_UnitSkill) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize Pos
	if mask.readBit() {
		err := read(buffer, &this.Pos)
		if err != nil {
			return err
		}
	}
	// deserialize SkillId
	if mask.readBit() {
		err := read(buffer, &this.SkillId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_UnitSkill) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
