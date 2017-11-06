package sqlite3

import (
	"database/sql/driver"
	"errors"
)

//import "database/sql/driver"

type SQLiteConn struct {
	pDB            uintptr
	nBusyTimeoutMs int
	bForeignKeys	bool
	bRecursiveTriggers bool
}

func (c *SQLiteConn) Prepare(query string) (driver.Stmt, error) {
	pVM, err := c.compile(query)
	if err != nil {
		return nil, err
	}
	return &SQLiteStmt{pDB: c.pDB, pVM: pVM}, nil
}

func (c *SQLiteConn) Begin() (driver.Tx, error) {
	nRet := sqlite3_exec(c.pDB,"COMMIT;");
	if nRet != SQLITE_OK {
		return nil, errors.New(sqlite3_errmsg(c.pDB))
	}
	return &SQLiteTx{pDB:c.pDB},nil
}

func (c *SQLiteConn) Close() error {
	if c.pDB != 0 {
		if sqlite3_close(c.pDB) == SQLITE_OK {
			c.pDB = 0
		} else {
			return errors.New(sqlite3_errmsg(c.pDB))
		}
	}
	return nil
}

func (c *SQLiteConn) compile(sql string) (uintptr, error) {
	var (
		zTail uintptr
		pVM   uintptr
		nRet  int
	)

	nRet = sqlite3_prepare_v2(c.pDB, sql, -1, &pVM, &zTail)

	if nRet != SQLITE_OK {
		return 0, errors.New(sqlite3_errmsg(c.pDB))
	}
	return pVM, nil
}
