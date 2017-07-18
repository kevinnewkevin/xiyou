package prpc

import (
	"bytes"
	"encoding/binary"
)

func WriteSizeType(buffer *bytes.Buffer, s uint) error {
	b := bytes.NewBuffer(nil)
	err := binary.Write(b, binary.LittleEndian, uint32(s))
	if err != nil {
		return err
	}
	n := 0
	if s <= 0X3F {
		n = 0
	} else if s <= 0X3FFF {
		n = 1
	} else if s <= 0X3FFFFF {
		n = 2
	} else if s <= 0X3FFFFFFF {
		n = 3
	}
	b2 := b.Bytes()
	b2[n] = byte(int(b2[n]) | (n << 6))
	for i := n; i >= 0; i-- {
		err := binary.Write(buffer, binary.LittleEndian, b2[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteString(buffer *bytes.Buffer, v string) error {
	b := bytes.NewBufferString(v)
	err := WriteSizeType(buffer, uint(b.Len()))
	if err != nil {
		return err
	}
	return binary.Write(buffer, binary.LittleEndian, b.Bytes())
}

func Write(buffer *bytes.Buffer, i interface{}) error {
	switch i.(type) {
	case int:
		return binary.Write(buffer, binary.LittleEndian, int8(i.(int)))
	case uint:
		return WriteSizeType(buffer, i.(uint))
	case string:
		return WriteString(buffer, i.(string))

	default:
		return binary.Write(buffer, binary.LittleEndian, i)
	}

}
