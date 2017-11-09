package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_UnitInfo struct {
	sync.Mutex
	InstId int64 //0
	UnitId int32 //1
	Level  int32 //2
	HP     int32 //3
	AGILE  int32 //4
	ATK    int32 //5
	DEF    int32 //6
	MATK   int32 //7
	MDEF   int32 //8
}

func (this *COM_UnitInfo) SetInstId(value int64) {
	this.Lock()
	defer this.Unlock()
	this.InstId = value
}
func (this *COM_UnitInfo) GetInstId() int64 {
	this.Lock()
	defer this.Unlock()
	return this.InstId
}
func (this *COM_UnitInfo) SetUnitId(value int32) {
	this.Lock()
	defer this.Unlock()
	this.UnitId = value
}
func (this *COM_UnitInfo) GetUnitId() int32 {
	this.Lock()
	defer this.Unlock()
	return this.UnitId
}
func (this *COM_UnitInfo) SetLevel(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Level = value
}
func (this *COM_UnitInfo) GetLevel() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Level
}
func (this *COM_UnitInfo) SetHP(value int32) {
	this.Lock()
	defer this.Unlock()
	this.HP = value
}
func (this *COM_UnitInfo) GetHP() int32 {
	this.Lock()
	defer this.Unlock()
	return this.HP
}
func (this *COM_UnitInfo) SetAGILE(value int32) {
	this.Lock()
	defer this.Unlock()
	this.AGILE = value
}
func (this *COM_UnitInfo) GetAGILE() int32 {
	this.Lock()
	defer this.Unlock()
	return this.AGILE
}
func (this *COM_UnitInfo) SetATK(value int32) {
	this.Lock()
	defer this.Unlock()
	this.ATK = value
}
func (this *COM_UnitInfo) GetATK() int32 {
	this.Lock()
	defer this.Unlock()
	return this.ATK
}
func (this *COM_UnitInfo) SetDEF(value int32) {
	this.Lock()
	defer this.Unlock()
	this.DEF = value
}
func (this *COM_UnitInfo) GetDEF() int32 {
	this.Lock()
	defer this.Unlock()
	return this.DEF
}
func (this *COM_UnitInfo) SetMATK(value int32) {
	this.Lock()
	defer this.Unlock()
	this.MATK = value
}
func (this *COM_UnitInfo) GetMATK() int32 {
	this.Lock()
	defer this.Unlock()
	return this.MATK
}
func (this *COM_UnitInfo) SetMDEF(value int32) {
	this.Lock()
	defer this.Unlock()
	this.MDEF = value
}
func (this *COM_UnitInfo) GetMDEF() int32 {
	this.Lock()
	defer this.Unlock()
	return this.MDEF
}
func (this *COM_UnitInfo) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(2)
	mask.writeBit(this.InstId != 0)
	mask.writeBit(this.UnitId != 0)
	mask.writeBit(this.Level != 0)
	mask.writeBit(this.HP != 0)
	mask.writeBit(this.AGILE != 0)
	mask.writeBit(this.ATK != 0)
	mask.writeBit(this.DEF != 0)
	mask.writeBit(this.MATK != 0)
	mask.writeBit(this.MDEF != 0)
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
	// serialize UnitId
	{
		if this.UnitId != 0 {
			err := write(buffer, this.UnitId)
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
	// serialize HP
	{
		if this.HP != 0 {
			err := write(buffer, this.HP)
			if err != nil {
				return err
			}
		}
	}
	// serialize AGILE
	{
		if this.AGILE != 0 {
			err := write(buffer, this.AGILE)
			if err != nil {
				return err
			}
		}
	}
	// serialize ATK
	{
		if this.ATK != 0 {
			err := write(buffer, this.ATK)
			if err != nil {
				return err
			}
		}
	}
	// serialize DEF
	{
		if this.DEF != 0 {
			err := write(buffer, this.DEF)
			if err != nil {
				return err
			}
		}
	}
	// serialize MATK
	{
		if this.MATK != 0 {
			err := write(buffer, this.MATK)
			if err != nil {
				return err
			}
		}
	}
	// serialize MDEF
	{
		if this.MDEF != 0 {
			err := write(buffer, this.MDEF)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_UnitInfo) Deserialize(buffer *bytes.Buffer) error {
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
	// deserialize UnitId
	if mask.readBit() {
		err := read(buffer, &this.UnitId)
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
	// deserialize HP
	if mask.readBit() {
		err := read(buffer, &this.HP)
		if err != nil {
			return err
		}
	}
	// deserialize AGILE
	if mask.readBit() {
		err := read(buffer, &this.AGILE)
		if err != nil {
			return err
		}
	}
	// deserialize ATK
	if mask.readBit() {
		err := read(buffer, &this.ATK)
		if err != nil {
			return err
		}
	}
	// deserialize DEF
	if mask.readBit() {
		err := read(buffer, &this.DEF)
		if err != nil {
			return err
		}
	}
	// deserialize MATK
	if mask.readBit() {
		err := read(buffer, &this.MATK)
		if err != nil {
			return err
		}
	}
	// deserialize MDEF
	if mask.readBit() {
		err := read(buffer, &this.MDEF)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_UnitInfo) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
