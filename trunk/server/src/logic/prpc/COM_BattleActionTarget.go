package prpc

import (
	"bytes"
	"suzuki/prpc"
)

type COM_BattleActionTarget struct {
	InstId         int64 //0
	ActionType     int   //1
	ActionParam    int32 //2
	ActionParamExt int32 //3
}

func (this *COM_BattleActionTarget) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.InstId != 0)
	mask.WriteBit(this.ActionType != 0)
	mask.WriteBit(this.ActionParam != 0)
	mask.WriteBit(this.ActionParamExt != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize InstId
	{
		if this.InstId != 0 {
			err := prpc.Write(buffer, this.InstId)
			if err != nil {
				return err
			}
		}
	}
	// serialize ActionType
	{
		if this.ActionType != 0 {
			err := prpc.Write(buffer, this.ActionType)
			if err != nil {
				return err
			}
		}
	}
	// serialize ActionParam
	{
		if this.ActionParam != 0 {
			err := prpc.Write(buffer, this.ActionParam)
			if err != nil {
				return err
			}
		}
	}
	// serialize ActionParamExt
	{
		if this.ActionParamExt != 0 {
			err := prpc.Write(buffer, this.ActionParamExt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_BattleActionTarget) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize InstId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.InstId)
		if err != nil {
			return err
		}
	}
	// deserialize ActionType
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ActionType)
		if err != nil {
			return err
		}
	}
	// deserialize ActionParam
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ActionParam)
		if err != nil {
			return err
		}
	}
	// deserialize ActionParamExt
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.ActionParamExt)
		if err != nil {
			return err
		}
	}
	return nil
}
