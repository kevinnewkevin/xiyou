package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_MlackMarketItemData struct {
	sync.Mutex
	IsBuy  bool  //0
	ItemId int32 //1
}

func (this *COM_MlackMarketItemData) SetIsBuy(value bool) {
	this.Lock()
	defer this.Unlock()
	this.IsBuy = value
}
func (this *COM_MlackMarketItemData) GetIsBuy() bool {
	this.Lock()
	defer this.Unlock()
	return this.IsBuy
}
func (this *COM_MlackMarketItemData) SetItemId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.ItemId = value
}
func (this *COM_MlackMarketItemData) GetItemId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.ItemId
}
func (this *COM_MlackMarketItemData) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.IsBuy)
	mask.writeBit(this.ItemId != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize IsBuy
	{
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
	return nil
}
func (this *COM_MlackMarketItemData) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize IsBuy
	this.IsBuy = mask.readBit()
	// deserialize ItemId
	if mask.readBit() {
		err := read(buffer, &this.ItemId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_MlackMarketItemData) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
