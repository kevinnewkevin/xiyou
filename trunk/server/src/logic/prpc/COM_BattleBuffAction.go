package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_BattleBuffAction struct {
	sync.Mutex
	BuffData   int32          //0
	Dead       bool           //1
	BuffChange COM_BattleBuff //2
}

func (this *COM_BattleBuffAction) SetBuffData(value int32) {
	this.Lock()
	defer this.Unlock()
	this.BuffData = value
}
func (this *COM_BattleBuffAction) GetBuffData() int32 {
	this.Lock()
	defer this.Unlock()
	return this.BuffData
}
func (this *COM_BattleBuffAction) SetDead(value bool) {
	this.Lock()
	defer this.Unlock()
	this.Dead = value
}
func (this *COM_BattleBuffAction) GetDead() bool {
	this.Lock()
	defer this.Unlock()
	return this.Dead
}
func (this *COM_BattleBuffAction) SetBuffChange(value COM_BattleBuff) {
	this.Lock()
	defer this.Unlock()
	this.BuffChange = value
}
func (this *COM_BattleBuffAction) GetBuffChange() COM_BattleBuff {
	this.Lock()
	defer this.Unlock()
	return this.BuffChange
}
func (this *COM_BattleBuffAction) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.BuffData != 0)
	mask.writeBit(this.Dead)
	mask.writeBit(true) //BuffChange
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize BuffData
	{
		if this.BuffData != 0 {
			err := write(buffer, this.BuffData)
			if err != nil {
				return err
			}
		}
	}
	// serialize Dead
	{
	}
	// serialize BuffChange
	{
		err := this.BuffChange.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_BattleBuffAction) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize BuffData
	if mask.readBit() {
		err := read(buffer, &this.BuffData)
		if err != nil {
			return err
		}
	}
	// deserialize Dead
	this.Dead = mask.readBit()
	// deserialize BuffChange
	if mask.readBit() {
		err := this.BuffChange.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_BattleBuffAction) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
