package game

/*
extern int __loadfile(void*);
extern int __GetStrings(void*);
extern int __GetFriend(void*);
extern int __GetFriends(void*);
extern int __GetTarget(void*);
extern int __GetTargets(void*);
extern int __GetUnitProperty(void*);
extern int __ChangeUnitProperty(void*);
extern int __ChangeSheld(void*);
extern int __ChangeSpecial(void*);
extern int __Attack(void*);
extern int __Cure(void*);
extern int __GetTime(void*);
extern int __GetCrit(void*);
extern int __AddBuff(void*);
extern int __HasBuff(void*);
extern int __HasDebuff(void*);
extern int __BuffMintsHp(void*);
extern int __GetCalcMagicDef(void*);
extern int __GetUnitMtk(void*);
extern int __GetUnitAtk(void*);
extern int __GetCalcDef(void*);


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
	_L.LoadApi(C.__GetFriend,"GetFriend","Player")
	_L.LoadApi(C.__GetFriends,"GetFriends","Player")
	_L.LoadApi(C.__GetTarget,"GetTarget","Player")
	_L.LoadApi(C.__GetTargets,"GetTargets","Player")
	_L.LoadApi(C.__GetUnitProperty,"GetUnitProperty","Player")
	_L.LoadApi(C.__ChangeUnitProperty,"ChangeUnitProperty","Player")
	_L.LoadApi(C.__ChangeSheld,"ChangeSheld","Player")
	_L.LoadApi(C.__ChangeSpecial,"ChangeSpecial","Player")
	_L.LoadApi(C.__GetUnitMtk,"GetUnitMtk","Player")
	_L.LoadApi(C.__GetCalcMagicDef,"GetCalcMagicDef","Player")
	_L.LoadApi(C.__GetUnitAtk,"GetUnitAtk","Player")
	_L.LoadApi(C.__GetCalcDef,"GetCalcDef","Player")

	_L.LoadApi(C.__Attack,"Attack","Battle")
	_L.LoadApi(C.__Cure,"Cure","Battle")
	_L.LoadApi(C.__GetCrit,"GetCrit","Battle")
	_L.LoadApi(C.__AddBuff,"AddBuff","Battle")
	_L.LoadApi(C.__HasBuff,"HasBuff","Battle")
	_L.LoadApi(C.__HasDebuff,"HasDebuff","Battle")
	_L.LoadApi(C.__BuffMintsHp,"BuffMintsHp","Battle")

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

	L := lua.GetLuaState(p)
	idx:= 1

	i := L.ToInteger(idx)

	fmt.Println("__GetStrings", int32(i))

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

//export __GetFriend
func __GetFriend(p unsafe.Pointer) C.int {

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
		if u.Camp != unit.Camp {
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

//export __ChangeUnitProperty
func __ChangeUnitProperty(p unsafe.Pointer) C.int {

	//fmt.Println("__ChangeUnitProperty")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	data := L.ToInteger(idx)
	idx ++
	property := L.ToString(idx)

	fmt.Println("__ChangeUnitProperty", battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	battle.ChangeUnitProperty(int64(unitid), int32(data), property)

	return 1
}

//export __ChangeSheld
func __ChangeSheld(p unsafe.Pointer) C.int {

	fmt.Println("__ChangeSheld")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	data := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	unit.VirtualHp += int32(data)

	return 1
}

//export __ChangeSpecial
func __ChangeSpecial(p unsafe.Pointer) C.int {

	fmt.Println("__ChangeSpecial")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	buffinstid := L.ToInteger(idx)
	idx ++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	unit.ChangeSpec(spec, int32(buffinstid))

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

//export __GetFriends
func __GetFriends(p unsafe.Pointer) C.int {

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

//export __HasBuff
func __HasBuff(p unsafe.Pointer) C.int {

	//fmt.Println("__AddBuff")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	target := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	//battle.AddBuff(int64(casterid), int64(target), int32(buffid), int32(data))
	has := battle.HasBuff(int64(target))

	L.PushBoolean(has)

	return 1
}

//export __HasDebuff
func __HasDebuff(p unsafe.Pointer) C.int {

	//fmt.Println("__AddBuff")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	target := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	//battle.AddBuff(int64(casterid), int64(target), int32(buffid), int32(data))
	has := battle.HasDebuff(int64(target))

	L.PushBoolean(has)

	return 1
}

//export __BuffMintsHp
func __BuffMintsHp(p unsafe.Pointer) C.int {

	fmt.Println("__BuffMintsHp")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	buffinstid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buff := unit.SelectBuff(int32(buffinstid))
	fmt.Println("__BuffMintsHp, ", buff.IsOver(battle.Round), !buff.IsOver(battle.Round))

	battle.BuffMintsHp(buff.CasterId, buff.Owner.InstId, buff.BuffId, buff.Data, !buff.IsOver(battle.Round))

	return 1
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//export __GetUnitDamage
func __GetUnitDamage(p unsafe.Pointer) C.int {    //物理   伤害

	fmt.Println("__GetUnitDamage")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	casterid := L.ToInteger(idx)
	idx ++
	targetid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	caster := battle.SelectOneUnit(int64(casterid))

	target := battle.SelectOneUnit(int64(targetid))

	finaldamage := CalcDamage(caster, target)

	L.PushNumber(float64(finaldamage))


	return 1
}

//export __GetUnitSheld
func __GetUnitSheld(p unsafe.Pointer) C.int {	// 获取场上所有玩家护盾数值

	fmt.Println("__GetUnitSheld")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	var sheld int32
	for _, u := range battle.Units {
		if u == nil {
			continue
		}
		sheld += u.VirtualHp
	}

	L.PushInteger(int(sheld))

	return 1
}

//export __GetUnitSheldPer
func __GetUnitSheldPer(p unsafe.Pointer) C.int {		//获取减伤百分比

	fmt.Println("__GetUnitSheldPer")

	L := lua.GetLuaState(p)
	L.PushInteger(1)

	return 1
}

//export __GetUnitAtk
func __GetUnitAtk(p unsafe.Pointer) C.int {		//获取减伤百分比  物理

	fmt.Println("__GetUnitAtk")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	atk := CalcAtk(unit)

	L.PushInteger(int(atk))

	return 1
}


//export __GetCalcDef
func __GetCalcDef(p unsafe.Pointer) C.int {		//获取减伤百分比  物理防御

	fmt.Println("__GetUnitAtk")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	atk := CalcDef(unit)

	L.PushInteger(int(atk))

	return 1
}
//export __GetUnitMtk
func __GetUnitMtk(p unsafe.Pointer) C.int {		//获取减伤百分比   法术

	fmt.Println("__GetUnitMtk")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	mtk := CalcMagicAtk(unit)

	L.PushInteger(int(mtk))

	return 1
}
//export __GetCalcMagicDef
func __GetCalcMagicDef(p unsafe.Pointer) C.int {		//获取减伤百分比   法术 防御

	fmt.Println("__GetUnitMtk")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	mtk := CalcMagicDef(unit)

	L.PushInteger(int(mtk))

	return 1
}




