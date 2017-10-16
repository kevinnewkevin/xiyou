package lua

/*
extern int __panic(void*);
extern int __err(void*);
extern int __msg(void*);
extern int __log(void*);
extern int __test(void*);
*/
import "C"
import (
	"fmt"
	"unsafe"
	"logic/std"
)

//export __panic
func __panic(p unsafe.Pointer) C.int {
	L := GetLuaState(p)
	errorDescC := L.ToString(-1)
	//defer C.free(errorDescC)
	panic(errorDescC)
	return 0
}

//export __err
func __err(p unsafe.Pointer) C.int {
	L := GetLuaState(p)
	idx := 1
	s := L.ToString(idx)
	fmt.Println("LUA {{", s)
	return 0
}

//export __msg
func __msg(p unsafe.Pointer) C.int {
	L := GetLuaState(p)
	idx := 1
	s := L.ToString(idx)
	fmt.Println("LUA [[", s)
	return 0
}

//export __log
func __log(p unsafe.Pointer) C.int {
	L := GetLuaState(p)
	idx := 1
	s := L.ToString(idx)
	std.LogDebug("LUA >> %s", s)
	return 0
}

//export __test
func __test(p unsafe.Pointer) C.int {
	L := GetLuaState(p)
	idx := 1
	v1 := L.ToInteger(idx)
	idx++
	v2 := L.ToInteger(idx)

	idx = 0
	L.PushInteger(v1 + v2)
	idx++
	return C.int(idx)
}

func (this *LuaState) OpenSys() {
	this.OpenPanic(C.__panic)

	this.LoadApi(C.__err, "err", "sys")
	this.LoadApi(C.__log, "log", "sys")
	this.LoadApi(C.__msg, "msg", "sys")
	this.LoadApi(C.__test, "test", "sys")

}
