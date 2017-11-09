package sqlite3

const (
	SQLITE_OK = iota /* Successful result */
	/* beginning-of-error-codes */
	SQLITE_ERROR            /* SQL error or missing database */
	SQLITE_INTERNAL         /* Internal logic error in SQLite */
	SQLITE_PERM             /* Access permission denied */
	SQLITE_ABORT            /* Callback routine requested an abort */
	SQLITE_BUSY             /* The database file is locked */
	SQLITE_LOCKED           /* A table in the database is locked */
	SQLITE_NOMEM            /* A malloc() failed */
	SQLITE_READONLY         /* Attempt to write a readonly database */
	SQLITE_INTERRUPT        /* Operation terminated by sqlite3_interrupt()*/
	SQLITE_IOERR            /* Some kind of disk I/O error occurred */
	SQLITE_CORRUPT          /* The database disk image is malformed */
	SQLITE_NOTFOUND         /* Unknown opcode in sqlite3_file_control() */
	SQLITE_FULL             /* Insertion failed because database is full */
	SQLITE_CANTOPEN         /* Unable to open the database file */
	SQLITE_PROTOCOL         /* Database lock protocol error */
	SQLITE_EMPTY            /* Database is empty */
	SQLITE_SCHEMA           /* The database schema changed */
	SQLITE_TOOBIG           /* String or BLOB exceeds size limit */
	SQLITE_CONSTRAINT       /* Abort due to constraint violation */
	SQLITE_MISMATCH         /* Data type mismatch */
	SQLITE_MISUSE           /* Library used incorrectly */
	SQLITE_NOLFS            /* Uses OS features not supported on host */
	SQLITE_AUTH             /* Authorization denied */
	SQLITE_FORMAT           /* Auxiliary database format error */
	SQLITE_RANGE            /* 2nd parameter to sqlite3_bind out of range */
	SQLITE_NOTADB           /* File opened that is not a database file */
	SQLITE_NOTICE           /* Notifications from sqlite3_log() */
	SQLITE_WARNING          /* Warnings from sqlite3_log() */
	SQLITE_ROW        = 100 /* sqlite3_step() has another row ready */
	SQLITE_DONE       = 101 /* sqlite3_step() has finished executing */
	/* end-of-error-codes */

	/*
	** CAPI3REF: Flags For File Open Operations
	**
	** These bit values are intended for use in the
	** 3rd parameter to the [sqlite3_open_v2()] interface and
	** in the 4th parameter to the [sqlite3_vfs.xOpen] method.
	 */
	SQLITE_OPEN_READONLY       = 0x00000001 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_READWRITE      = 0x00000002 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_CREATE         = 0x00000004 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_DELETEONCLOSE  = 0x00000008 /* VFS only */
	SQLITE_OPEN_EXCLUSIVE      = 0x00000010 /* VFS only */
	SQLITE_OPEN_AUTOPROXY      = 0x00000020 /* VFS only */
	SQLITE_OPEN_URI            = 0x00000040 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_MEMORY         = 0x00000080 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_MAIN_DB        = 0x00000100 /* VFS only */
	SQLITE_OPEN_TEMP_DB        = 0x00000200 /* VFS only */
	SQLITE_OPEN_TRANSIENT_DB   = 0x00000400 /* VFS only */
	SQLITE_OPEN_MAIN_JOURNAL   = 0x00000800 /* VFS only */
	SQLITE_OPEN_TEMP_JOURNAL   = 0x00001000 /* VFS only */
	SQLITE_OPEN_SUBJOURNAL     = 0x00002000 /* VFS only */
	SQLITE_OPEN_MASTER_JOURNAL = 0x00004000 /* VFS only */
	SQLITE_OPEN_NOMUTEX        = 0x00008000 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_FULLMUTEX      = 0x00010000 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_SHAREDCACHE    = 0x00020000 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_PRIVATECACHE   = 0x00040000 /* Ok for sqlite3_open_v2() */
	SQLITE_OPEN_WAL            = 0x00080000 /* VFS only */

	SQLITE_BUSY_TIMEOUT = 60000 // 60 seconds

	SQLITE_INTEGER = 1
	SQLITE_FLOAT   = 2
	SQLITE_TEXT    = 3
	SQLITE_BLOB    = 4
	SQLITE_NULL    = 5

	SQLITE_DECL_TYPE_TIMESTAMP = "timestamp"
	SQLITE_DECL_TYPE_DATETIME  = "datetime"
	SQLITE_DECL_TYPE_DATE      = "date"
	SQLITE_DECL_TYPE_BOOLEAN   = "boolean"

	SQLITE_STATIC = 0
	SQLITE_TRANSIENT = -1
)
