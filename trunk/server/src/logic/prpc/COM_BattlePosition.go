package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_BattlePosition struct {
	sync.Mutex
	InstId   int64 //0
	Position int32 //1
}

func (this *COM_BattlePosition) SetInstId(value int64) {
	this.Lock()
	defer this.Unlock()
	this.InstId = value
}
func (this *COM_BattlePosition) GetInstId() int64 {
	this.Lock()
	defer this.Unlock()
	return this.InstId
}
func (this *COM_BattlePosition) SetPosition(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Position = value
}
func (this *COM_BattlePosition) GetPosition() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Position
}
func (this *COM_BattlePosition) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.InstId != 0)
	mask.writeBit(this.Position != 0)
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
	// serialize Position
	{
		if this.Position != 0 {
			err := write(buffer, this.Position)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattlePosition) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize Position
	if mask.readBit() {
		err := read(buffer, &this.Position)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_BattlePosition) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
