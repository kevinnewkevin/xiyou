package game

/*
extern int __loadfile(void*);
extern int __GetStrings(void*);
extern int __GetTarget(void*);
extern int __GetTargets(void*);
extern int __GetUnitProperty(void*);
extern int __Attack(void*);
*/
import "C"
import (
	"unsafe"
	"logic/lua"

	"fmt"
)

var (
	_L *lua.LuaState
	_R string
)

func InitLua(r string){
	_R = r
	_L = &lua.LuaState{}
	_L.Open()

	_L.OpenSys()
	_L.LoadApi(C.__loadfile,"loadfile","sys")
	_L.LoadApi(C.__GetStrings,"GetStrings","Player")
	_L.LoadApi(C.__GetTarget,"GetTarget","Player")
	_L.LoadApi(C.__GetTargets,"GetTargets","Player")
	_L.LoadApi(C.__GetUnitProperty,"GetUnitProperty","Player")
	_L.LoadApi(C.__Attack,"Attack","Battle")
	_L.LoadFile(_R + "main.lua")

}



//export __loadfile
func __loadfile(p unsafe.Pointer) C.int {

	fmt.Println("ppppppppp",p)
	L := lua.GetLuaState(p)
	fileName := L.ToString(-1)
	L.LoadFile(_R + fileName)
	return 0
}

//export __GetStrings
func __GetStrings(p unsafe.Pointer) C.int {

	fmt.Println("__GetStrings")

	return 0
}

//export __GetTarget
func __GetTarget(p unsafe.Pointer) C.int {

	fmt.Println("__GetTargets")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	fmt.Println(battleid, uid)

	//battle := FindBattle(int64(battleid))
	//
	//for _, u := range battle.Units {
	//	if u == nil {
	//		continue
	//	}
	//	if u.Owner.BattleCamp == 2 {
	//		continue
	//	}
	//}

	L.PushInteger(11111)

	return 1
}


//export __GetUnitProperty
func __GetUnitProperty(p unsafe.Pointer) C.int {

	fmt.Println("__GetUnitProperty")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	property := L.ToInteger(idx)

	fmt.Println(battleid, unitid, property)

	L.PushInteger(50)

	return 1
}


//export __GetTargets
func __GetTargets(p unsafe.Pointer) C.int {

	fmt.Println("__GetTargets")

	L := lua.GetLuaState(p)
	battleid := L.ToInteger(-1)

	fmt.Println(battleid)

	ls := []int{1,2,3,4,5}

	L.CreateTable(5, 0)
	L.PushInteger(-1)
	L.RawSetI(-2, 0)

	for i :=0; i < len(ls); i++ {
		L.PushInteger(ls[i])
		L.RawSetI(-2, i+1)
	}

	return 1
}

//export __Attack
func __Attack(p unsafe.Pointer) C.int {

	fmt.Println("__Attack battleid, t, damage, true")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	target := L.ToInteger(idx)
	idx ++
	damage := L.ToInteger(idx)
	idx ++
	crit := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	battle.MintsHp(int64(target), int32(damage), int32(crit))

	fmt.Println(battleid, target, crit, battle, damage)

	return 1
}

