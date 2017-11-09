package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_ThrowCard struct {
	sync.Mutex
	InstId   int64 //0
	EntityId int32 //1
	Level    int32 //2
}

func (this *COM_ThrowCard) SetInstId(value int64) {
	this.Lock()
	defer this.Unlock()
	this.InstId = value
}
func (this *COM_ThrowCard) GetInstId() int64 {
	this.Lock()
	defer this.Unlock()
	return this.InstId
}
func (this *COM_ThrowCard) SetEntityId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.EntityId = value
}
func (this *COM_ThrowCard) GetEntityId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.EntityId
}
func (this *COM_ThrowCard) SetLevel(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Level = value
}
func (this *COM_ThrowCard) GetLevel() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Level
}
func (this *COM_ThrowCard) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.InstId != 0)
	mask.writeBit(this.EntityId != 0)
	mask.writeBit(this.Level != 0)
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
	// serialize EntityId
	{
		if this.EntityId != 0 {
			err := write(buffer, this.EntityId)
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
	return nil
}
func (this *COM_ThrowCard) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
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
	// deserialize EntityId
	if mask.readBit() {
		err := read(buffer, &this.EntityId)
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
	return nil
}
func (this *COM_ThrowCard) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
