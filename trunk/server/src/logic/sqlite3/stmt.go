package sqlite3

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type SQLiteStmt struct {
	pDB uintptr
	pVM uintptr
}

func (s *SQLiteStmt) Close() error {
	if s.pVM != 0 {
		nRet := sqlite3_finalize(s.pVM)
		s.pVM = 0
		if nRet != SQLITE_OK {
			return errors.New(sqlite3_errmsg(s.pDB))
		}
	}
	return nil
}

func (s *SQLiteStmt) NumInput() int {
	return sqlite3_bind_parameter_count(s.pVM)
}

func (s *SQLiteStmt) Exec(args []driver.Value) (driver.Result, error) {
	if s.NumInput() != len(args){
		return  nil, fmt.Errorf("SQL input field length not match current %d, input %d",s.NumInput(),len(args))
	}

	for i, v := range args{
		err := s.bind(i+1,v)
		if err != nil{
			return nil,err
		}
	}

	nRet := sqlite3_step(s.pVM)

	if nRet == SQLITE_DONE{
		return  &SQLiteResult{pVM:s.pVM,pDB:s.pDB}, nil
	}else if nRet == SQLITE_ROW {
		return  &SQLiteResult{pVM:s.pVM,pDB:s.pDB}, nil
	}else{
		nRet = sqlite3_finalize(s.pVM)
		return  nil, errors.New(sqlite3_errmsg(s.pDB))
	}
}

func (s *SQLiteStmt) Query(args []driver.Value) (driver.Rows, error) {
	if s.NumInput() != len(args){
		return  nil, fmt.Errorf("SQL input field length not match current %d, input %d",s.NumInput(),len(args))
	}

	for i, v := range args{
		err := s.bind(i+1,v)
		if err != nil{
			return nil,err
		}
	}
	return &SQLiteResult{pVM:s.pVM,pDB:s.pDB}, nil
	//nRet := sqlite3_step(s.pVM)
	//
	//if nRet == SQLITE_DONE{
	//	return   &SQLiteResult{pVM:s.pVM,pDB:s.pDB}, nil
	//}else if nRet == SQLITE_ROW {
	//	return  &SQLiteResult{pVM:s.pVM,pDB:s.pDB}, nil
	//}else{
	//	nRet = sqlite3_finalize(s.pVM)
	//	return  nil, errors.New(sqlite3_errmsg(s.pDB))
	//}
}

func (s *SQLiteStmt) bind(i int, v interface{}) error{
	nRet := SQLITE_OK

	switch v.(type) {
	case string:
		nRet = sqlite3_bind_text(s.pVM,i,v.(string),SQLITE_TRANSIENT)
		break
	case int:
		nRet = sqlite3_bind_int(s.pVM,i,v.(int))
		break
	case uint:
		nRet = sqlite3_bind_int(s.pVM,i,int(v.(uint)))
		break
	case int64:
		nRet = sqlite3_bind_int64(s.pVM,i,v.(int64))
		break
	case uint64:
		nRet = sqlite3_bind_int64(s.pVM,i, int64(v.(uint64)))
		break
	case float32:
		nRet =sqlite3_bind_double(s.pVM,i, float64(v.(float32)))
		break
	case float64:
		nRet =sqlite3_bind_double(s.pVM,i, v.(float64))
		break
	case []byte:
		nRet = sqlite3_bind_blob(s.pVM,i,v.([]byte),SQLITE_TRANSIENT)
		break
	case nil:
		nRet = sqlite3_bind_null(s.pVM,i)
		break
	default:
		return fmt.Errorf("SQLite bind type can not use %d ", i )
	}

	if nRet != SQLITE_OK{
		return errors.New(sqlite3_errmsg(s.pDB))
	}

	return nil
}