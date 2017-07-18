package prpc

import (
	"bytes"
	"encoding/binary"
)

func ReadSizeType(buffer *bytes.Buffer, s *uint) error {
	*s = uint(0)
	b := uint8(0)
	err := binary.Read(buffer, binary.LittleEndian, &b)
	if err != nil {
		return err
	}
	n := uint(b) & 0XC0 >> 6
	*s = uint(b) & 0X3F
	for i := uint(0); i < n; i++ {
		err := binary.Read(buffer, binary.LittleEndian, &b)
		if err != nil {
			return err
		}
		*s = (*s << 8) | uint(b)
	}
	return nil
}

func ReadString(buffer *bytes.Buffer, v *string) error {
	s := uint(0)
	err := ReadSizeType(buffer, &s)
	if err != nil {
		return err
	}
	if s == 0 {
		return nil
	}
	b := make([]byte, s)
	err = binary.Read(buffer, binary.LittleEndian, &b)
	if err != nil {
		return err
	}
	*v = string(b)
	return nil
}

func Read(buffer *bytes.Buffer, i interface{}) error {
	switch i.(type) {
	case *int:
		{
			var value int8
			err := binary.Read(buffer, binary.LittleEndian, &value)
			*(i.(*int)) = int(value)
			return err
		}
	case *uint:
		return ReadSizeType(buffer, i.(*uint))
	case *string:
		return ReadString(buffer, i.(*string))
	default:
		return binary.Read(buffer, binary.LittleEndian, i)
	}
}
