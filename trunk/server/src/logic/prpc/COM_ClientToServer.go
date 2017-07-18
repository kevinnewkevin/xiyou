package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type COM_ClientToServer_Login struct {
	info COM_LoginInfo //0
}
type COM_ClientToServer_CreatePlayer struct {
	tempId     int32  //0
	playerName string //1
}
type COM_ClientToServer_SetBattleUnit struct {
	instId int64 //0
}
type COM_ClientToServer_SetupBattle struct {
	positionList []COM_BattlePosition //0
}
type COM_ClientToServerStub struct {
	Sender prpc.StubSender
}
type COM_ClientToServerProxy interface {
	Login(info COM_LoginInfo) error                      // 0
	CreatePlayer(tempId int32, playerName string) error  // 1
	SetBattleUnit(instId int64) error                    // 2
	JoinBattle() error                                   // 3
	SetupBattle(positionList []COM_BattlePosition) error // 4
}

func (this *COM_ClientToServer_Login) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ClientToServer_Login) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ClientToServer_CreatePlayer) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.tempId != 0)
	mask.WriteBit(len(this.playerName) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize tempId
	{
		if this.tempId != 0 {
			err := prpc.Write(buffer, this.tempId)
			if err != nil {
				return err
			}
		}
	}
	// serialize playerName
	if len(this.playerName) != 0 {
		err := prpc.Write(buffer, this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ClientToServer_CreatePlayer) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize tempId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.tempId)
		if err != nil {
			return err
		}
	}
	// deserialize playerName
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.playerName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ClientToServer_SetBattleUnit) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ClientToServer_SetBattleUnit) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ClientToServer_SetupBattle) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(len(this.positionList) != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize positionList
	if len(this.positionList) != 0 {
		{
			err := prpc.Write(buffer, uint(len(this.positionList)))
			if err != nil {
				return err
			}
		}
		for _, value := range this.positionList {
			err := value.Serialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ClientToServer_SetupBattle) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize positionList
	if mask.ReadBit() {
		var size uint
		err := prpc.Read(buffer, &size)
		if err != nil {
			return err
		}
		this.positionList = make([]COM_BattlePosition, size)
		for i, _ := range this.positionList {
			err := this.positionList[i].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ClientToServerStub) Login(info COM_LoginInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := COM_ClientToServer_Login{}
	_0.info = info
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ClientToServerStub) CreatePlayer(tempId int32, playerName string) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := COM_ClientToServer_CreatePlayer{}
	_1.tempId = tempId
	_1.playerName = playerName
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ClientToServerStub) SetBattleUnit(instId int64) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := COM_ClientToServer_SetBattleUnit{}
	_2.instId = instId
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ClientToServerStub) JoinBattle() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ClientToServerStub) SetupBattle(positionList []COM_BattlePosition) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	_4 := COM_ClientToServer_SetupBattle{}
	_4.positionList = positionList
	err = _4.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_COM_ClientToServer_Login(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := COM_ClientToServer_Login{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.Login(_0.info)
}
func Bridging_COM_ClientToServer_CreatePlayer(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := COM_ClientToServer_CreatePlayer{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.CreatePlayer(_1.tempId, _1.playerName)
}
func Bridging_COM_ClientToServer_SetBattleUnit(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := COM_ClientToServer_SetBattleUnit{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.SetBattleUnit(_2.instId)
}
func Bridging_COM_ClientToServer_JoinBattle(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.JoinBattle()
}
func Bridging_COM_ClientToServer_SetupBattle(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_4 := COM_ClientToServer_SetupBattle{}
	err := _4.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.SetupBattle(_4.positionList)
}
func COM_ClientToServerDispatch(buffer *bytes.Buffer, p COM_ClientToServerProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	pid := uint16(0XFFFF)
	err := prpc.Read(buffer, &pid)
	if err != nil {
		return err
	}
	switch pid {
	case 0:
		return Bridging_COM_ClientToServer_Login(buffer, p)
	case 1:
		return Bridging_COM_ClientToServer_CreatePlayer(buffer, p)
	case 2:
		return Bridging_COM_ClientToServer_SetBattleUnit(buffer, p)
	case 3:
		return Bridging_COM_ClientToServer_JoinBattle(buffer, p)
	case 4:
		return Bridging_COM_ClientToServer_SetupBattle(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
