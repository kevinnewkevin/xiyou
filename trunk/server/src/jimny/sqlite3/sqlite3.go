package sqlite3


/*
#include <stddef.h>
#include <stdlib.h>
#cgo windows LDFLAGS: -L. sqlite3.dll
#cgo linux LDFLAGS:  -lsqlite3
//sqlite3_aggregate_context
//sqlite3_aggregate_count
//sqlite3_auto_extension
//sqlite3_backup_finish
//sqlite3_backup_init
//sqlite3_backup_pagecount
//sqlite3_backup_remaining
//sqlite3_backup_step
extern int sqlite3_bind_blob(void* pStmt, int i, const void* bValue, int len, ptrdiff_t xDel);
//extern sqlite3_bind_blob64
extern int sqlite3_bind_double(void* pStmt, int i, double rValue);
extern int sqlite3_bind_int(void* pStmt, int i,  int iValue);
extern int sqlite3_bind_int64(void* pStmt, int i,  long long value);
extern int sqlite3_bind_null(void* pStmt, int i);
extern int sqlite3_bind_parameter_count(void* pStmt);
extern int sqlite3_bind_parameter_index(void* pStmt, const void* zName);
//sqlite3_bind_parameter_name
//sqlite3_bind_pointer
extern int sqlite3_bind_text(void* pDB, int i, const void* zData, int nData, ptrdiff_t xDel);
//sqlite3_bind_text16
//sqlite3_bind_text64
//sqlite3_bind_value
//sqlite3_bind_zeroblob
//sqlite3_bind_zeroblob64
//sqlite3_blob_bytes
//sqlite3_blob_close
//sqlite3_blob_open
//sqlite3_blob_read
//sqlite3_blob_reopen
//sqlite3_blob_write
//sqlite3_busy_handler
extern int sqlite3_busy_timeout(void* pDB, int ms);
//sqlite3_cancel_auto_extension
extern int sqlite3_changes(void *pDB);
extern int sqlite3_clear_bindings(void* pStmt);
extern int sqlite3_close(void* pDB);
//sqlite3_close_v2
//sqlite3_collation_needed
//sqlite3_collation_needed16
extern const void* sqlite3_column_blob(void* pStmt, int i);
extern int sqlite3_column_bytes(void* pStmt, int i);
//sqlite3_column_bytes16
extern int sqlite3_column_count(void* pStmt);
//sqlite3_column_database_name
//sqlite3_column_database_name16
extern const char* sqlite3_column_decltype(void* pStmt, int i);
//sqlite3_column_decltype16
extern double sqlite3_column_double(void* pStmt, int i);
extern int sqlite3_column_int(void* pStmt, int i);
extern long long sqlite3_column_int64(void* pStmt, int i);
extern const char* sqlite3_column_name(void* pStmt, int i);
//sqlite3_column_name16
//sqlite3_column_origin_name
//sqlite3_column_origin_name16
//sqlite3_column_table_name
//sqlite3_column_table_name16
extern const char* sqlite3_column_text(void* pStmt, int i);
//sqlite3_column_text16
extern int sqlite3_column_type(void* pStmt, int i);
//sqlite3_column_value
//sqlite3_commit_hook
//sqlite3_compileoption_get
//sqlite3_compileoption_used
//sqlite3_complete
//sqlite3_complete16
//sqlite3_config
//sqlite3_context_db_handle
//sqlite3_create_collation
//sqlite3_create_collation_v2
//sqlite3_create_collation16
//sqlite3_create_function
//sqlite3_create_function_v2
//sqlite3_create_function16
//sqlite3_create_module
//sqlite3_create_module_v2
//sqlite3_data_count
//sqlite3_data_directory
//sqlite3_db_cacheflush
//sqlite3_db_config
//sqlite3_db_filename
//sqlite3_db_handle
//sqlite3_db_mutex
//sqlite3_db_readonly
//sqlite3_db_release_memory
//sqlite3_db_status
//sqlite3_declare_vtab
//sqlite3_enable_load_extension
//sqlite3_enable_shared_cache
extern int sqlite3_errcode(void* pDB);
extern const char* sqlite3_errmsg(void* pDB);
//sqlite3_errmsg16
//sqlite3_errstr
extern int sqlite3_exec(void* pDB, const void* sql, void* callback, void* args, void* pzMsg);
//sqlite3_expanded_sql
//sqlite3_expired
//sqlite3_extended_errcode
//sqlite3_extended_result_codes
//sqlite3_file_control
extern int sqlite3_finalize(void* pStmt);
//sqlite3_free
//sqlite3_free_table
//sqlite3_fts5_may_be_corrupt
//sqlite3_get_autocommit
//sqlite3_get_auxdata
//sqlite3_get_table
//sqlite3_global_recover
//sqlite3_initialize
//sqlite3_interrupt
extern long long sqlite3_last_insert_rowid(void *pDB);
extern const char* sqlite3_libversion();
extern int sqlite3_libversion_number();
//sqlite3_limit
//sqlite3_load_extension
//sqlite3_log
//sqlite3_malloc
//sqlite3_malloc64
//sqlite3_memory_alarm
//sqlite3_memory_highwater
//sqlite3_memory_used
//sqlite3_mprintf
//sqlite3_msize
//sqlite3_mutex_alloc
//sqlite3_mutex_enter
//sqlite3_mutex_free
//sqlite3_mutex_leave
//sqlite3_mutex_try
//sqlite3_next_stmt
extern int sqlite3_open(const void *filename, void *ppDB );
extern int sqlite3_open_v2(const void* filename, void *ppDB, int flags, const void* zVfs);
extern int sqlite3_open16(const void* filename, void *ppDB);
//sqlite3_os_end
//sqlite3_os_init
//sqlite3_overload_function
//sqlite3_prepare
extern int sqlite3_prepare_v2(void* pDB, const void* zSQL, int nByte, void *ppStmt, const void *pzTail);
//sqlite3_prepare_v3
//sqlite3_prepare16
//sqlite3_prepare16_v2
//sqlite3_prepare16_v3
//sqlite3_profile
//sqlite3_progress_handler
//sqlite3_randomness
//sqlite3_realloc
//sqlite3_realloc64
//sqlite3_release_memory
extern int sqlite3_reset(void* pStmt);
//sqlite3_reset_auto_extension
//sqlite3_result_blob
//sqlite3_result_blob64
//sqlite3_result_double
//sqlite3_result_error
//sqlite3_result_error_code
//sqlite3_result_error_nomem
//sqlite3_result_error_toobig
//sqlite3_result_error16
//sqlite3_result_int
//sqlite3_result_int64
//sqlite3_result_null
//sqlite3_result_pointer
//sqlite3_result_subtype
//sqlite3_result_text
//sqlite3_result_text16
//sqlite3_result_text16be
//sqlite3_result_text16le
//sqlite3_result_text64
//sqlite3_result_value
//sqlite3_result_zeroblob
//sqlite3_result_zeroblob64
//sqlite3_rollback_hook
//sqlite3_rtree_geometry_callback
//sqlite3_rtree_query_callback
//sqlite3_set_authorizer
//sqlite3_set_auxdata
//sqlite3_set_last_insert_rowid
//sqlite3_shutdown
//sqlite3_sleep
//sqlite3_snprintf
//sqlite3_soft_heap_limit
//sqlite3_soft_heap_limit64
extern const char* sqlite3_sourceid();
//sqlite3_sql
//sqlite3_status
//sqlite3_status64
extern int sqlite3_step(void* pStmt);
//sqlite3_stmt_busy
//sqlite3_stmt_readonly
//sqlite3_stmt_status
//sqlite3_strglob
//sqlite3_stricmp
//sqlite3_strlike
//sqlite3_strnicmp
//sqlite3_system_errno
//sqlite3_table_column_metadata
//sqlite3_temp_directory
//sqlite3_test_control
//sqlite3_thread_cleanup
extern int sqlite3_threadsafe();
//sqlite3_total_changes
//sqlite3_trace
//sqlite3_trace_v2
//sqlite3_transfer_bindings
//sqlite3_update_hook
//sqlite3_uri_boolean
//sqlite3_uri_int64
//sqlite3_uri_parameter
//sqlite3_user_data
//sqlite3_value_blob
//sqlite3_value_bytes
//sqlite3_value_bytes16
//sqlite3_value_double
//sqlite3_value_dup
//sqlite3_value_free
//sqlite3_value_int
//sqlite3_value_int64
//sqlite3_value_numeric_type
//sqlite3_value_pointer
//sqlite3_value_subtype
//sqlite3_value_text
//sqlite3_value_text16
//sqlite3_value_text16be
//sqlite3_value_text16le
//sqlite3_value_type
//sqlite3_version
//sqlite3_vfs_find
//sqlite3_vfs_register
//sqlite3_vfs_unregister
//sqlite3_vmprintf
//sqlite3_vsnprintf
//sqlite3_vtab_config
//sqlite3_vtab_on_conflict
//sqlite3_wal_autocheckpoint
//sqlite3_wal_checkpoint
//sqlite3_wal_checkpoint_v2
//sqlite3_wal_hook
//sqlite3_win32_is_nt
//sqlite3_win32_mbcs_to_utf8
//sqlite3_win32_mbcs_to_utf8_v2
//sqlite3_win32_set_directory
//sqlite3_win32_sleep
//sqlite3_win32_unicode_to_utf8
//sqlite3_win32_utf8_to_mbcs
//sqlite3_win32_utf8_to_mbcs_v2
//sqlite3_win32_utf8_to_unicode
//sqlite3_win32_write_debug
*/
import "C"
import (
	"unsafe"
)

func sqlite3_libversion()string{
	return C.GoString(C.sqlite3_libversion())
}

func sqlite3_sourceid()string{
	return C.GoString(C.sqlite3_sourceid())
}

func sqlite3_libversion_number()int{
	return int(C.sqlite3_libversion_number())
}

func sqlite3_threadsafe() bool {
	return C.sqlite3_threadsafe() != 0
}

func sqlite3_busy_timeout(pDB uintptr, ms int )int{
	return int(C.sqlite3_busy_timeout(unsafe.Pointer(pDB),C.int(ms)))
}

func sqlite3_lasterror(pDB uintptr) error {
	return SQLiteError{code:sqlite3_errcode(pDB),desc:sqlite3_errmsg(pDB)}
}

func sqlite3_open(filename string, ppDB *uintptr) int {
	szFilename := unsafe.Pointer(C.CString(filename))
	defer C.free(szFilename)
	return int(C.sqlite3_open(szFilename, unsafe.Pointer(ppDB)))
}

func sqlite3_close(pDB uintptr) int {
	return int(C.sqlite3_close(unsafe.Pointer(pDB)))
}

func sqlite3_errcode(pDB uintptr) int {
	return int(C.sqlite3_errcode(unsafe.Pointer(pDB)))
}

func sqlite3_errmsg(pDB uintptr) string {
	szError := C.sqlite3_errmsg(unsafe.Pointer(pDB))
	return C.GoString(szError)
}

func sqlite3_prepare_v2(pDB uintptr, sql string, nByte int, ppVM *uintptr, pzTail *uintptr) int {
	szSQL := unsafe.Pointer(C.CString(sql))
	defer C.free(szSQL)
	return int(C.sqlite3_prepare_v2(unsafe.Pointer(pDB), szSQL, C.int(nByte), unsafe.Pointer(ppVM), unsafe.Pointer(pzTail)))
}

func sqlite3_exec(pDB uintptr, sql string) int {
	szSQL := unsafe.Pointer(C.CString(sql))
	defer C.free(szSQL)
	return int(C.sqlite3_exec(unsafe.Pointer(pDB),szSQL,nil,nil,nil))
}

func sqlite3_finalize(pVM uintptr) int {
	return int(C.sqlite3_finalize(unsafe.Pointer(pVM)))
}

func sqlite3_reset(pVM uintptr) int{
	return int(C.sqlite3_reset(unsafe.Pointer(pVM)))
}

func sqlite3_bind_blob(pVM uintptr, i int, blob []byte,x int64) int{
	return int(C.sqlite3_bind_blob(unsafe.Pointer(pVM),C.int(i),unsafe.Pointer(&blob[0]),C.int(len(blob)),C.ptrdiff_t(x)))
}

func sqlite3_bind_double(pVM uintptr, i int, v float64) int{
	return  int(C.sqlite3_bind_double(unsafe.Pointer(pVM),C.int(i),C.double(v)))
}

func sqlite3_bind_int(pVM uintptr, i int, v int) int {
	return int(C.sqlite3_bind_int(unsafe.Pointer(pVM),C.int(i),C.int(v)))
}

func sqlite3_bind_int64(pVM uintptr, i int, v int64) int {
	return int(C.sqlite3_bind_int64(unsafe.Pointer(pVM),C.int(i),C.longlong (v)))
}

func sqlite3_bind_null(pVM uintptr, i int ) int{
	return int(C.sqlite3_bind_null(unsafe.Pointer(pVM),C.int(i)))
}

func sqlite3_bind_text(pVM uintptr, i int, text string, x int64)int{
	var (
		b []byte
		n = len(text)
	)
	if n == 0{
		b = []byte{0}
	}else{
		b = []byte(text)
	}

	return int(C.sqlite3_bind_text(unsafe.Pointer(pVM),C.int(i),unsafe.Pointer(&b[0]),C.int(n),C.ptrdiff_t(x)))
}

func sqlite3_bind_parameter_count(pVM uintptr) int {
	return int(C.sqlite3_bind_parameter_count(unsafe.Pointer(pVM)))
}
func sqlite3_bind_parameter_index(pVM uintptr, name string ) int{
	szName := unsafe.Pointer(C.CString(name))
	defer C.free(szName)
	return int(C.sqlite3_bind_parameter_index(unsafe.Pointer(pVM),szName))
}

func sqlite3_step(pVM uintptr) int{
	return  int(C.sqlite3_step(unsafe.Pointer(pVM)))
}

func sqlite3_last_insert_rowid(pDB uintptr) int64{
	return int64(C.sqlite3_last_insert_rowid(unsafe.Pointer(pDB)))
}

func sqlite3_changes(pDB uintptr)int64{
	return  int64(C.sqlite3_changes(unsafe.Pointer(pDB)))
}

func sqlite3_column_count(pVM uintptr)int{
	return int(C.sqlite3_column_count(unsafe.Pointer(pVM)))
}

func sqlite3_column_name(pVM uintptr, i int)string{
	return C.GoString(C.sqlite3_column_name(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_decltype(pVM uintptr, i int) string{
	return C.GoString(C.sqlite3_column_decltype(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_type(pVM uintptr, i int) int{
	return int(C.sqlite3_column_type(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_int(pVM uintptr, i int)int{
	return int(C.sqlite3_column_int(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_int64(pVM uintptr, i int)int64{
	return int64(C.sqlite3_column_int64(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_double(pVM uintptr, i int)float64{
	return float64(C.sqlite3_column_double(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_bytes(pVM uintptr, i int)int{
	return int(C.sqlite3_column_bytes(unsafe.Pointer(pVM), C.int(i)))
}

func sqlite3_column_blob(pVM uintptr, i int) []byte{
	n := sqlite3_column_bytes(pVM,i)
	p := (*byte)(C.sqlite3_column_blob(unsafe.Pointer(pVM), C.int(i)))
	return (*[1 << 30]byte)(unsafe.Pointer(p))[0:n]
}

func sqlite3_column_text(pVM uintptr, i int)string{
	n := sqlite3_column_bytes(pVM,i)
	return C.GoStringN(C.sqlite3_column_text(unsafe.Pointer(pVM), C.int(i)),C.int(n))
}