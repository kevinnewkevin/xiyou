package game

/*
extern int __loadfile(void*);
extern int __GetStrings(void*);
extern int __GetTarget(void*);
extern int __GetTargets(void*);
extern int __GetUnitProperty(void*);
extern int __Attack(void*);
extern int __Cure(void*);
extern int __GetTime(void*);
extern int __GetCrit(void*);
extern int __AddBuff(void*);
*/
import "C"
import (
	"unsafe"
	"logic/lua"

	"fmt"
	"time"
)

var (
	_L *lua.LuaState
	_R string
)

func InitLua(r string){
	_R = r
	_L = &lua.LuaState{}
	_L.Open()
	_L.OpenLibs()
	_L.OpenSys()
	_L.LoadApi(C.__loadfile,"loadfile","sys")

	_L.LoadApi(C.__GetStrings,"GetStrings","Player")
	_L.LoadApi(C.__GetTarget,"GetTarget","Player")
	_L.LoadApi(C.__GetTargets,"GetTargets","Player")
	_L.LoadApi(C.__GetUnitProperty,"GetUnitProperty","Player")

	_L.LoadApi(C.__Attack,"Attack","Battle")
	_L.LoadApi(C.__Cure,"Cure","Battle")
	_L.LoadApi(C.__GetCrit,"GetCrit","Battle")
	_L.LoadApi(C.__AddBuff,"AddBuff","Battle")

	_L.LoadApi(C.__GetTime,"GetTime","os")
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

	//fmt.Println("__GetStrings")

	return 0
}

//export __GetTarget
func __GetTarget(p unsafe.Pointer) C.int {

	//fmt.Println("__GetTargets")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	//fmt.Println(battleid, uid)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(uid))

	t_id := 0
	for _, u := range battle.Units {
		if u == nil {
			continue
		}
		if u.Camp == unit.Camp {
			continue
		}
		t_id = int(u.InstId)
		break
	}

	//fmt.Println("__GetTargets end ,", t_id)

	L.PushInteger(t_id)

	return 1
}


//export __GetUnitProperty
func __GetUnitProperty(p unsafe.Pointer) C.int {

	//fmt.Println("__GetUnitProperty")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	property := L.ToString(idx)

	//fmt.Println(battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	date := battle.GetUnitProperty(int64(unitid), property)

	L.PushInteger(date)

	return 1
}


//export __GetTargets
func __GetTargets(p unsafe.Pointer) C.int {

	//fmt.Println("__GetTargets")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	num := L.ToInteger(idx)

	////fmt.Println("4444444444", battleid, unitid, num)

	battle := FindBattle(int64(battleid))

	ls := battle.SelectMoreTarget(int64(unitid), num)

	L.NewTable()
	//L.PushInteger(-1)
	//L.RawSetI(-2, 0)

	for i :=0; i < len(ls); i++ {
		L.PushInteger(i + 1)
		L.PushInteger(int(ls[i]))
		L.SetTable(-3)
	}

	return 1
}

//export __Attack
func __Attack(p unsafe.Pointer) C.int {

	//fmt.Println("__Attack battleid")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	caster := L.ToInteger(idx)
	idx ++
	target := L.ToInteger(idx)
	idx ++
	damage := L.ToInteger(idx)
	idx ++
	crit := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	battle.MintsHp(int64(caster), int64(target), int32(damage), int32(crit))

	//fmt.Println("55555555555555", battleid, caster, target, crit, damage)

	return 1
}

//export __Cure
func __Cure(p unsafe.Pointer) C.int {

	//fmt.Println("__Cure")

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

	battle.AddHp(int64(target), int32(damage), int32(crit))

	//fmt.Println("6666666666666", battleid, target, crit, damage)

	return 1
}

//export __GetCrit
func __GetCrit(p unsafe.Pointer) C.int {

	L := lua.GetLuaState(p)
	idx := 1
	skillid := L.ToInteger(idx)

	crit := IsCrit(int32(skillid))

	L.PushInteger(crit)

	return 1
}

//export __GetTime
func __GetTime(p unsafe.Pointer) C.int {

	//fmt.Println("__GetTime")

	L := lua.GetLuaState(p)

	time_unix := time.Now().Unix()

	L.PushInteger(int(time_unix))

	return 1
}

//export __AddBuff
func __AddBuff(p unsafe.Pointer) C.int {

	//fmt.Println("__AddBuff")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	casterid := L.ToInteger(idx)
	idx ++
	target := L.ToInteger(idx)
	idx ++
	buffid := L.ToInteger(idx)
	idx ++
	data := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	battle.AddBuff(int64(casterid), int64(target), int32(buffid), int32(data))

	return 1
}

