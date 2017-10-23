package prpc

import (
	"bytes"
	"encoding/binary"
)

const (
	kFieldMaskOffset = 3
	kFieldMaskMagic  = 7
	kFieldMaskMasker = 128
)

type (
	Mask struct {
		m []byte
		p uint
	}
)

func (this *Mask) WriteBit(b bool) {
	if b {
		this.m[this.p>>kFieldMaskOffset] |= (kFieldMaskMasker >> (this.p & kFieldMaskMagic))
	}
	this.p++
}

func (this *Mask) ReadBit() bool {
	p := this.p
	this.p++
	return this.m[p>>kFieldMaskOffset]&(kFieldMaskMasker>>(p&kFieldMaskMagic)) != 0
}

func (this *Mask) Bytes() *[]byte {

	return &this.m
}

func NewMask0(buffer *bytes.Buffer, s int) (*Mask, error) {
	m := Mask{make([]byte, s), 0}
	err := binary.Read(buffer, binary.LittleEndian, &(m.m))
	if err != nil {
		return &m, err
	}
	return &m, nil
}

func NewMask1(s int) *Mask {
	return &Mask{make([]byte, s), 0}
}
