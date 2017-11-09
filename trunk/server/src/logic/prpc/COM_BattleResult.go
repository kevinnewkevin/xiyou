package prpc

import (
	"bytes"
	"encoding/json"
	"sync"
)

type COM_BattleResult struct {
	sync.Mutex
	Win            int32          //0
	Money          int32          //1
	Exp            int32          //2
	KillMonsters   []int32        //3
	BattleRound    int32          //4
	MySelfDeathNum int32          //5
	BattleItems    []COM_ItemInst //6
}

func (this *COM_BattleResult) SetWin(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Win = value
}
func (this *COM_BattleResult) GetWin() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Win
}
func (this *COM_BattleResult) SetMoney(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Money = value
}
func (this *COM_BattleResult) GetMoney() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Money
}
func (this *COM_BattleResult) SetExp(value int32) {
	this.Lock()
	defer this.Unlock()
	this.Exp = value
}
func (this *COM_BattleResult) GetExp() int32 {
	this.Lock()
	defer this.Unlock()
	return this.Exp
}
func (this *COM_BattleResult) SetKillMonsters(value []int32) {
	this.Lock()
	defer this.Unlock()
	this.KillMonsters = value
}
func (this *COM_BattleResult) GetKillMonsters() []int32 {
	this.Lock()
	defer this.Unlock()
	return this.KillMonsters
}
func (this *COM_BattleResult) SetBattleRound(value int32) {
	this.Lock()
	defer this.Unlock()
	this.BattleRound = value
}
func (this *COM_BattleResult) GetBattleRound() int32 {
	this.Lock()
	defer this.Unlock()
	return this.BattleRound
}
func (this *COM_BattleResult) SetMySelfDeathNum(value int32) {
	this.Lock()
	defer this.Unlock()
	this.MySelfDeathNum = value
}
func (this *COM_BattleResult) GetMySelfDeathNum() int32 {
	this.Lock()
	defer this.Unlock()
	return this.MySelfDeathNum
}
func (this *COM_BattleResult) SetBattleItems(value []COM_ItemInst) {
	this.Lock()
	defer this.Unlock()
	this.BattleItems = value
}
func (this *COM_BattleResult) GetBattleItems() []COM_ItemInst {
	this.Lock()
	defer this.Unlock()
	return this.BattleItems
}
func (this *COM_BattleResult) Serialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask := newMask1(1)
	mask.writeBit(this.Win != 0)
	mask.writeBit(this.Money != 0)
	mask.writeBit(this.Exp != 0)
	mask.writeBit(len(this.KillMonsters) != 0)
	mask.writeBit(this.BattleRound != 0)
	mask.writeBit(this.MySelfDeathNum != 0)
	mask.writeBit(len(this.BattleItems) != 0)
	{
		err := write(buffer, mask.bytes())
		if err != nil {
			return err
		}
	}
	// serialize Win
	{
		if this.Win != 0 {
			err := write(buffer, this.Win)
			if err != nil {
				return err
			}
		}
	}
	// serialize Money
	{
		if this.Money != 0 {
			err := write(buffer, this.Money)
			if err != nil {
				return err
			}
		}
	}
	// serialize Exp
	{
		if this.Exp != 0 {
			err := write(buffer, this.Exp)
			if err != nil {
				return err
			}
		}
	}
	// serialize KillMonsters
	if len(this.KillMonsters) != 0 {
		{
			err := write(buffer, uint(len(this.KillMonsters)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.KillMonsters {
			err := write(buffer, value)
			if err != nil {
				return err
			}
		}
	}
	// serialize BattleRound
	{
		if this.BattleRound != 0 {
			err := write(buffer, this.BattleRound)
			if err != nil {
				return err
			}
		}
	}
	// serialize MySelfDeathNum
	{
		if this.MySelfDeathNum != 0 {
			err := write(buffer, this.MySelfDeathNum)
			if err != nil {
				return err
			}
		}
	}
	// serialize BattleItems
	if len(this.BattleItems) != 0 {
		{
			err := write(buffer, uint(len(this.BattleItems)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.BattleItems {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleResult) Deserialize(buffer *bytes.Buffer) error {
	this.Lock()
	defer this.Unlock()
	//field mask
	mask, err := newMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize Win
	if mask.readBit() {
		err := read(buffer, &this.Win)
		if err != nil {
			return err
		}
	}
	// deserialize Money
	if mask.readBit() {
		err := read(buffer, &this.Money)
		if err != nil {
			return err
		}
	}
	// deserialize Exp
	if mask.readBit() {
		err := read(buffer, &this.Exp)
		if err != nil {
			return err
		}
	}
	// deserialize KillMonsters
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.KillMonsters = make([]int32, size)
		for i, _ := range this.KillMonsters {
			err := read(buffer, &this.KillMonsters[i])
			if err != nil {
				return err
			}
		}
	}
	// deserialize BattleRound
	if mask.readBit() {
		err := read(buffer, &this.BattleRound)
		if err != nil {
			return err
		}
	}
	// deserialize MySelfDeathNum
	if mask.readBit() {
		err := read(buffer, &this.MySelfDeathNum)
		if err != nil {
			return err
		}
	}
	// deserialize BattleItems
	if mask.readBit() {
		var size uint
		err := read(buffer, &size)
		if err != nil {
			return err
		}
		this.BattleItems = make([]COM_ItemInst, size)
		for i, _ := range this.BattleItems {
			err := this.BattleItems[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleResult) String() string {
	b, e := json.Marshal(this)
	if e != nil {
		return e.Error()
	} else {
		return string(b)
	}
}
