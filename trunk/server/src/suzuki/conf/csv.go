package conf

import (
	"sync"
	"strconv"
	"strings"
)

type CSV struct{
	sync.RWMutex
	data  [][]string
	names map[string]int
	ErrorColum  int
	ErrorLine 	int
}

func(this* CSV) set(row, column int, value string){
	if len(this.data) <= row{
		panic("CSV error row index overflow")
	}
	if this.data[row] == nil{
		panic("CSV error row is nil")
	}
	if len(this.data[row]) <= column{
		panic("CSV error column index overflow")
	}
	this.data[row][column] = value
}

func(this* CSV) get(row, column int) string{
	if len(this.data) <= row{
		panic("CSV error row index overflow")
	}
	if this.data[row] == nil{
		panic("CSV error row is nil")
	}
	if len(this.data[row]) <= column{
		panic("CSV error column index overflow")
	}
	return this.data[row][column]
}

func (this*CSV) Length() int {
	return len(this.data)
}

func (this *CSV) SetBool(row int, column string, value bool) {
	this.set(row, this.names[column], strconv.FormatBool(value))
}
func (this *CSV) SetInt(row int, column string, value int) {
	this.set(row, this.names[column], strconv.FormatInt(int64(value), 32))
}

func (this *CSV) SetInt64(row int, column string, value int64) {
	this.set(row, this.names[column], strconv.FormatInt(int64(value), 64))
}
func (this *CSV) SetFloat32(row int, column string, value float32) {
	this.set(row, this.names[column], strconv.FormatFloat(float64(value), 'E', -1, 32))
}
func (this *CSV) SetFloat64(row int, column string, value float64) {
	this.set(row, this.names[column], strconv.FormatFloat(float64(value), 'E', -1, 64))
}
func (this *CSV) SetString(row int, column string, value string) {
	this.set(row, this.names[column], value)
}
func (this *CSV) SetStrings(row int, column string, values []string) {
	this.set(row, this.names[column], strings.Join(values, ","))
}

func (this *CSV) GetBool(row int, column string) bool {
	return this.TryGetBool(row, column, false)
}
func (this *CSV) GetInt(row int, column string) int {
	return this.TryGetInt(row, column, 0)
}
func (this *CSV) GetInt32(row int, column string) int32 {
	return int32(this.TryGetInt(row, column, 0))
}

func (this *CSV) GetInt64(row int, column string) int64 {
	return this.TryGetInt64(row, column, 0)
}
func (this *CSV) GetFloat32(row int, column string) float32 {
	return this.TryGetFloat32(row, column, 0)
}
func (this *CSV) GetFloat64(row int, column string) float64 {
	return this.TryGetFloat64(row, column, 0)
}
func (this *CSV) GetString(row int, column string) string {
	return this.TryGetString(row, column, "")
}
func (this *CSV) GetStrings(row int, column string) []string {
	return this.TryGetStrings(row, column, []string{})
}

func (this *CSV) TryGetBool(row int, column string, defaultValue bool) bool {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValue
	}
	r, e := strconv.ParseBool(s)
	if e != nil {
		return defaultValue
	}
	return r
}
func (this *CSV) TryGetInt(row int, column string, defaultValue int) int {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValue
	}
	r, e := strconv.ParseInt(s, 10, 32)
	if e != nil {
		return defaultValue
	}
	return int(r)
}

func (this *CSV) TryGetInt64(row int, column string, defaultValue int64) int64 {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValue
	}
	r, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return defaultValue
	}
	return r
}
func (this *CSV) TryGetFloat32(row int, column string, defaultValue float32) float32 {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValue
	}
	r, e := strconv.ParseFloat(s, 32)
	if e != nil {
		return defaultValue
	}
	return float32(r)
}
func (this *CSV) TryGetFloat64(row int, column string, defaultValue float64) float64 {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValue
	}
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return defaultValue
	}
	return float64(r)
}
func (this *CSV) TryGetString(row int, column string, defaultValue string) string {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValue
	}
	return s
}
func (this *CSV) TryGetStrings(row int, column string, defaultValues []string) []string {
	s := this.get(row, this.names[column])
	if s == "" {
		return defaultValues
	}
	return strings.Split(s, ",")
}

func (this *CSV) LoadFile(fileName string) error {
	this.Lock()
	defer this.Unlock()
	err := this.parseFile(fileName)
	return err
}

func (this *CSV) LoadString(s string) error {
	this.Lock()
	defer this.Unlock()
	err := this.parseString(s)
	return err
}

func NewCSVFile(fileName string) (*CSV, error) {
	r := &CSV{
		sync.RWMutex{},
		[][]string{},
		map[string]int{},
		0,
		0,
 	}
	e := r.LoadFile(fileName)
	return r, e
}

func NewCSVString(s string) (*CSV, error) {
	r := &CSV{
		sync.RWMutex{},
		[][]string{},
		map[string]int{},
		0,
		0,
	}
	e := r.LoadString(s)
	return r, e
}
