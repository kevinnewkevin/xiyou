package sqlite3

import "fmt"

type SQLiteError struct{
	code int
	desc string
}

func (e SQLiteError) Error() string{
	return fmt.Sprintf("CODE:%d | DESC:%s",e.code,e.desc)
}
