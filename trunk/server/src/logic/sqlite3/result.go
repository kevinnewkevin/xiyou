package sqlite3

import (
	"errors"
	"database/sql/driver"
	"io"
	"logic/log"
)

type SQLiteResult struct {
	pVM      uintptr
	pDB 	 uintptr
	colNames []string
}


func (r *SQLiteResult)LastInsertId() (int64, error){
	if r.pDB == 0{
		return 0, errors.New("SQLite invalid")
	}
	return sqlite3_last_insert_rowid(r.pDB), nil
}

// RowsAffected returns the number of rows affected by the
// query.
func (r *SQLiteResult)RowsAffected() (int64, error){
	if r.pDB == 0{
		return 0, errors.New("SQLite invalid")
	}
	return sqlite3_changes(r.pDB), nil
}

func (r *SQLiteResult) Columns() []string {
	if len(r.colNames) == 0 {
		len := sqlite3_column_count(r.pVM)
		for i := 0; i < len; i++ {
			r.colNames = append(r.colNames, sqlite3_column_name(r.pVM, i))
		}

	}

	return r.colNames
}

func (r *SQLiteResult) Close() error {
	return nil
}

func (r *SQLiteResult) Next(dest []driver.Value) error {

	nRet := sqlite3_step(r.pVM)



	if nRet == SQLITE_DONE{
		return  io.EOF
	}
	if nRet != SQLITE_ROW{
		log.Info("ROW")
		nRet = sqlite3_reset(r.pVM)
		if nRet != SQLITE_OK {
			return sqlite3_lasterror(r.pDB)
		}
		return nil
	}

	for i := range dest {

		switch sqlite3_column_type(r.pVM,i) {
		case SQLITE_INTEGER:
			switch  sqlite3_column_decltype(r.pVM,i) {
			case SQLITE_DECL_TYPE_TIMESTAMP, SQLITE_DECL_TYPE_DATETIME, SQLITE_DECL_TYPE_DATE:
				break
			case SQLITE_DECL_TYPE_BOOLEAN:
				dest[i] = sqlite3_column_int(r.pVM, i) > 0
				break
			default:
				dest[i] = sqlite3_column_int64(r.pVM, i)
				break
			}
			break
		case SQLITE_FLOAT:
			dest[i] = sqlite3_column_double(r.pVM, i)
			break
		case SQLITE_TEXT:
			dest[i] = sqlite3_column_text(r.pVM, i)
			break
		case SQLITE_BLOB:
			dest[i] = sqlite3_column_blob(r.pVM, i)
			break
		case SQLITE_NULL:
			dest[i] = nil
			break

		}
	}

	return nil
}
