package sqlite3

import "errors"

type SQLiteTx struct{
	pDB uintptr
}

func (t *SQLiteTx)Commit() error{
	nRet := sqlite3_exec(t.pDB,"COMMIT;");
	if nRet != SQLITE_OK {
		return errors.New(sqlite3_errmsg(t.pDB))
	}
	return nil
}
func (t *SQLiteTx)Rollback() error{
	nRet := sqlite3_exec(t.pDB,"ROLLBACK;");
	if nRet != SQLITE_OK {
		return errors.New(sqlite3_errmsg(t.pDB))
	}
	return nil
}