package prpc
import(
  "bytes"
  "encoding/binary"
)
const(
  NoneDispatchMatchError = "NoneDispatchMatchError"
  NoneMethodError = "NoneMethodError"
  NoneBufferError = "NoneBufferError"
  NoneProxyError = "NoneProxyError"
  kFieldMaskOffset = 3
  kFieldMaskMagic  = 7
  kFieldMaskMasker = 128
)
type(
  StubSender interface{
    MethodBegin() *bytes.Buffer
    MethodEnd() error
  }
  mask struct{
    m []byte
    p uint
  }
)
func (this *mask) writeBit(b bool) {
  if b{
    this.m[this.p>>kFieldMaskOffset] |= (kFieldMaskMasker >> (this.p & kFieldMaskMagic))
  }
  this.p++
}
func (this *mask) readBit() bool {
  p := this.p
  this.p++
  return this.m[p>>kFieldMaskOffset]&(kFieldMaskMasker>>(p&kFieldMaskMagic)) != 0
}
func (this *mask) bytes() *[]byte { return &this.m }
func newMask0(buffer *bytes.Buffer, s int) (*mask, error) {
  m := mask{make([]byte, s), 0}
  err := binary.Read(buffer, binary.LittleEndian, &(m.m))
  if err != nil {
    return &m, err
  }
  return &m, nil
}
func newMask1(s int) *mask { return &mask{make([]byte, s), 0} }
func writeSize(buffer *bytes.Buffer, s uint) error {
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
func writeString(buffer *bytes.Buffer, v string) error {
  b := bytes.NewBufferString(v)
  err := writeSize(buffer, uint(b.Len()))
  if err != nil {
    return err
  }
  return binary.Write(buffer, binary.LittleEndian, b.Bytes())
}
func write(buffer *bytes.Buffer, i interface{}) error {
  switch i.(type) {
    case int:
      return binary.Write(buffer, binary.LittleEndian, int8(i.(int)))
    case uint:
      return writeSize(buffer, i.(uint))
    case string:
      return writeString(buffer, i.(string))
    default:
      return binary.Write(buffer, binary.LittleEndian, i)
  }
}
func readSize(buffer *bytes.Buffer, s *uint) error {
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
func readString(buffer *bytes.Buffer, v *string) error {
  s := uint(0)
  err := readSize(buffer, &s)
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
func read(buffer *bytes.Buffer, i interface{}) error {
  switch i.(type) {
    case *int:
      {
        var value int8
        err := binary.Read(buffer, binary.LittleEndian, &value)
        *(i.(*int)) = int(value)
        return err
      }
    case *uint:
      return readSize(buffer, i.(*uint))
    case *string:
      return readString(buffer, i.(*string))
    default:
      return binary.Read(buffer, binary.LittleEndian, i)
  }
}
