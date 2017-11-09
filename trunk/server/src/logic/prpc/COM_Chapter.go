package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_Chapter struct {
	sync.Mutex
	ChapterId     int32              //0
	SmallChapters []COM_SmallChapter //1
	StarReward    []int32            //2
}

func (this *COM_Chapter) SetChapterId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.ChapterId = value
}
func (this *COM_Chapter) GetChapterId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.ChapterId
}
func (this *COM_Chapter) SetSmallChapters(value []COM_SmallChapter) {
	this.Lock()
	defer this.Unlock()
	this.SmallChapters = value
}
func (this *COM_Chapter) GetSmallChapters() []COM_SmallChapter {
	this.Lock()
	defer this.Unlock()
	return this.SmallChapters
}
func (this *COM_Chapter) SetStarReward(value []int32) {
	this.Lock()
	defer this.Unlock()
	this.StarReward = value
}
func (this *COM_Chapter) GetStarReward() []int32 {
	this.Lock()
	defer this.Unlock()
	return this.StarReward
}
func (this *COM_Chapter) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.ChapterId != 0)
	mask.writeBit(len(this.SmallChapters) != 0)
	mask.writeBit(len(this.StarReward) != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize ChapterId
	{
		if this.ChapterId != 0 {
			err := write(buffer, this.ChapterId)
			if err != nil {
				return err
			}
		}
	}
	// serialize SmallChapters
	if len(this.SmallChapters) != 0 {
		{
			err := write(buffer, uint(len(this.SmallChapters)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.SmallChapters {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize StarReward
	if len(this.StarReward) != 0 {
		{
			err := write(buffer, uint(len(this.StarReward)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.StarReward {
			err := write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Chapter) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize ChapterId
	if mask.readBit() {
		err := read(buffer, &this.ChapterId)
		if err != nil {
			return err
		}
	}
	// deserialize SmallChapters
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.SmallChapters = make([]COM_SmallChapter, size)
		for i, _ := range this.SmallChapters {
			err := this.SmallChapters[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize StarReward
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.StarReward = make([]int32, size)
		for i, _ := range this.StarReward {
			err := read(buffer, &this.StarReward[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_Chapter) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
