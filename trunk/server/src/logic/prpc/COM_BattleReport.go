package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BattleReport struct {
	UnitList   []COM_BattleUnit   //0
	ActionList []COM_BattleAction //1
}

func (this *COM_BattleReport) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.UnitList) != 0)
	mask.WriteBit(len(this.ActionList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize UnitList
	if len(this.UnitList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.UnitList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.UnitList {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// serialize ActionList
	if len(this.ActionList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.ActionList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.ActionList {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleReport) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize UnitList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.UnitList = make([]COM_BattleUnit, size)
		for i, _ := range this.UnitList {
			err := this.UnitList[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	// deserialize ActionList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.ActionList = make([]COM_BattleAction, size)
		for i, _ := range this.ActionList {
			err := this.ActionList[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
