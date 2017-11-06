package sqlite3

import (
	"database/sql/driver"
	"errors"
	"database/sql"
	"logic/log"
)

type (
	SQLiteDriver struct{}
)
/*
 *
 * */
func (d SQLiteDriver) Open(dsn string) (driver.Conn, error) {
	var (
		pDB  uintptr
		nRet int
	)

	if !sqlite3_threadsafe() {
		log.Error("SQLite thread unsafe ")
	}

	log.Info("SQLite VERSION(%s|%d) SOURCEID(%s)", sqlite3_libversion(),sqlite3_libversion_number(),sqlite3_sourceid())

	nRet = sqlite3_open(dsn, &pDB)

	if nRet != SQLITE_OK {
		return nil, errors.New(sqlite3_errmsg(pDB))
	}

	nRet = sqlite3_busy_timeout(pDB,SQLITE_BUSY_TIMEOUT)
	if nRet != SQLITE_OK {
		sqlite3_close(pDB)
		return nil, errors.New(sqlite3_errmsg(pDB))
	}
	return &SQLiteConn{pDB: pDB, nBusyTimeoutMs: SQLITE_BUSY_TIMEOUT}, nil
}


func init() {
	sql.Register("sqlite3", &SQLiteDriver{})
}
