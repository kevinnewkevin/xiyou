package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_ItemInst struct {
	sync.Mutex
	ItemId int32 //0
	InstId int64 //1
	Stack  int32 //2
}

func (this *COM_ItemInst) SetItemId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.ItemId = value
}
func (this *COM_ItemInst) GetItemId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.ItemId
}
func (this *COM_ItemInst) SetInstId(value int64) {
	this.Lock()
	defer this.Unlock()
	this.InstId = value
}
func (this *COM_ItemInst) GetInstId() int64 {
	this.Lock()
	defer this.Unlock()
	return this.InstId
}
func (this *COM_ItemInst) SetStack(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Stack = value
}
func (this *COM_ItemInst) GetStack() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Stack
}
func (this *COM_ItemInst) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.ItemId != 0)
	mask.writeBit(this.InstId != 0)
	mask.writeBit(this.Stack != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize ItemId
	{
		if this.ItemId != 0 {
			err := write(buffer, this.ItemId)
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
	// serialize Stack
	{
		if this.Stack != 0 {
			err := write(buffer, this.Stack)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ItemInst) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize ItemId
	if mask.readBit() {
		err := read(buffer, &this.ItemId)
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
	// deserialize Stack
	if mask.readBit() {
		err := read(buffer, &this.Stack)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ItemInst) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
