package game

/*
extern int __loadfile(void*);
extern int __DefineCards(void*);
extern int __GetStrings(void*);
extern int __GetFriend(void*);
extern int __GetMainFriend(void*);
extern int __GetFriends(void*);
extern int __GetTarget(void*);
extern int __GetMainTarget(void*);
extern int __CheckUnitDead(void*);
extern int __GetTargets(void*);
extern int __GetTargetsAround(void*);
extern int __GetUnitProperty(void*);
extern int __ChangeCptProperty(void*);
extern int __ChangeIptProperty(void*);
extern int __AddSheld(void*);
extern int __PopSheld(void*);
extern int __DamageSheld(void*);
extern int __ChangeSpecial(void*);
extern int __PopSpec(void*);
extern int __GetSpecial(void*);
extern int __GetOneSpecial(void*);
extern int __GetCheckSpec(void*);
extern int __GetBuffLockId(void*);
extern int __Attack(void*);
extern int __Cure(void*);
extern int __GetTime(void*);
extern int __GetCrit(void*);
extern int __AddBuff(void*);
extern int __AddSkillBuff(void*);
extern int __HasBuff(void*);
extern int __HasDebuff(void*);
extern int __BuffMintsHp(void*);
extern int __BuffCureHp(void*);
extern int __BuffUpdate(void*);
extern int __BuffChangeStillData(void*);
extern int __BuffChangeData(void*);
extern int __GetCalcMagicDef(void*);
extern int __GetUnitMtk(void*);
extern int __GetUnitAtk(void*);
extern int __GetCalcDef(void*);
extern int __GetUnitDamage(void*);
extern int __GetMagicDamage(void*);
extern int __ClacSheld(void*);
extern int __TargetOver(void*);
extern int __TargetOn(void*);
extern int __PopAllBuffByDebuff(void*);
extern int __PopAllBuffBybuff(void*);
extern int __GetUnitSheld(void*);
extern int __FrontTarget(void*);
extern int __LineTraget(void*);
extern int __BackTarget(void*);
extern int __GetSpecialData(void*);
extern int __GetOneSheld(void*);
extern int __ClacStrongPer(void*);
extern int __ClacWeakPer(void*);
extern int __ChangeBuffTimes(void*);
extern int __GetMyUnitIProperty(void*);
extern int __AddMyUnitEnergy(void*);

*/
import "C"
import (
	"unsafe"
	"logic/lua"

	"fmt"
	"time"
	"logic/prpc"
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
	_L.LoadApi(C.__DefineCards,"DefineCards","Server")
	_L.LoadApi(C.__GetStrings,"GetStrings","Player")
	_L.LoadApi(C.__GetFriend,"GetFriend","Player")
	_L.LoadApi(C.__GetMainFriend,"GetMainFriend","Player")
	_L.LoadApi(C.__CheckUnitDead,"CheckUnitDead","Player")
	_L.LoadApi(C.__GetFriends,"GetFriends","Player")
	_L.LoadApi(C.__GetTarget,"GetTarget","Player")
	_L.LoadApi(C.__GetMainTarget,"GetMainTarget","Player")
	_L.LoadApi(C.__GetTargets,"GetTargets","Player")
	_L.LoadApi(C.__GetTargetsAround,"GetTargetsAround","Player")
	_L.LoadApi(C.__GetUnitProperty,"GetUnitProperty","Player")
	_L.LoadApi(C.__ChangeCptProperty,"ChangeUnitProperty","Player")
	_L.LoadApi(C.__ChangeIptProperty,"ChangeIptProperty","Player")
	_L.LoadApi(C.__AddSheld,"AddSheld","Player")
	_L.LoadApi(C.__PopSheld,"PopSheld","Player")
	_L.LoadApi(C.__DamageSheld,"DownSheld","Player")
	_L.LoadApi(C.__ChangeSpecial,"ChangeSpecial","Player")
	_L.LoadApi(C.__PopSpec,"PopSpec","Player")
	_L.LoadApi(C.__GetSpecial,"GetSpecial","Player")
	_L.LoadApi(C.__GetOneSpecial,"GetOneSpecial","Player")
	_L.LoadApi(C.__GetCheckSpec,"GetCheckSpec","Player")
	_L.LoadApi(C.__GetBuffLockId,"GetBuffLockId","Player")
	_L.LoadApi(C.__GetUnitMtk,"GetUnitMtk","Player")
	_L.LoadApi(C.__GetCalcMagicDef,"GetCalcMagicDef","Player")
	_L.LoadApi(C.__GetUnitAtk,"GetUnitAtk","Player")
	_L.LoadApi(C.__GetCalcDef,"GetCalcDef","Player")
	_L.LoadApi(C.__GetUnitDamage,"GetUnitDamage","Player")
	_L.LoadApi(C.__GetMagicDamage,"GetMagicDamage","Player")
	_L.LoadApi(C.__ClacSheld,"ClacSheld","Player")
	_L.LoadApi(C.__PopAllBuffByDebuff,"PopAllBuffByDebuff","Player")
	_L.LoadApi(C.__PopAllBuffBybuff,"PopAllBuffBybuff","Player")
	_L.LoadApi(C.__GetUnitSheld,"GetUnitSheld","Player")
	_L.LoadApi(C.__FrontTarget,"FrontTarget","Player")
	_L.LoadApi(C.__LineTraget,"LineTraget","Player")
	_L.LoadApi(C.__BackTarget,"BackTarget","Player")
	_L.LoadApi(C.__GetOneSheld,"GetOneSheld","Player")
	_L.LoadApi(C.__GetSpecialData,"GetSpecialData","Player")
	_L.LoadApi(C.__ClacWeakPer,"ClacWeakPer","Player")
	_L.LoadApi(C.__ClacStrongPer,"ClacStrongPer","Player")
	_L.LoadApi(C.__ChangeBuffTimes,"ChangeBuffTimes","Player")
	_L.LoadApi(C.__GetMyUnitIProperty,"GetMyUnitIProperty","Player")
	_L.LoadApi(C.__AddMyUnitEnergy,"AddMyUnitEnergy","Player")

	_L.LoadApi(C.__Attack,"Attack","Battle")
	_L.LoadApi(C.__Cure,"Cure","Battle")
	_L.LoadApi(C.__GetCrit,"GetCrit","Battle")
	_L.LoadApi(C.__AddBuff,"AddBuff","Battle")
	_L.LoadApi(C.__AddSkillBuff,"AddSkillBuff","Battle")
	_L.LoadApi(C.__HasBuff,"HasBuff","Battle")
	_L.LoadApi(C.__HasDebuff,"HasDebuff","Battle")
	_L.LoadApi(C.__BuffMintsHp,"BuffMintsHp","Battle")
	_L.LoadApi(C.__BuffCureHp,"BuffCureHp","Battle")
	_L.LoadApi(C.__BuffUpdate,"BuffUpdate","Battle")
	_L.LoadApi(C.__BuffChangeStillData,"BuffChangeStillData","Battle")
	_L.LoadApi(C.__BuffChangeData,"BuffChangeData","Battle")
	_L.LoadApi(C.__TargetOver,"TargetOver","Battle")
	_L.LoadApi(C.__TargetOn,"TargetOn","Battle")

	_L.LoadApi(C.__GetTime,"GetTime","os")
	_L.LoadFile(_R + "main.lua")

}



//export __loadfile
func __loadfile(p unsafe.Pointer) C.int {

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

//export __DefineCards
func __DefineCards(p unsafe.Pointer) C.int {

	L := lua.GetLuaState(p)
	idx:= 1

	cards := L.ToString(idx)
	SetDefaultUnits(cards)

	//fmt.Println("__GetStrings")`

	return 0
}

//export __GetTarget
func __GetTarget(p unsafe.Pointer) C.int {

	fmt.Println("__GetTarget")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	//fmt.Println(battleid, uid)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(uid))

	t_id := battle.SelectNearTarget(unit.InstId)

	fmt.Println("__GetTarget end ,", t_id)

	L.PushInteger(int(t_id))

	return 1
}

//export __GetMainTarget
func __GetMainTarget(p unsafe.Pointer) C.int {

	fmt.Println("__GetMainTarget")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	//fmt.Println(battleid, uid)

	battle := FindBattle(int64(battleid))

	var t_id int64
	t_id = battle.selectMainUnit(int64(uid), false)

	fmt.Println("__GetMainTarget end ,", t_id)

	L.PushInteger(int(t_id))

	return 1
}
//export __CheckUnitDead
func __CheckUnitDead(p unsafe.Pointer) C.int {

	fmt.Println("__CheckUnitDead")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	//fmt.Println(battleid, uid)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(uid))

	var t_id bool
	if unit.IsDead() {
		t_id = true
	}else{
		t_id = false
	}
	L.PushBoolean(t_id)

	return 1
}
//export __GetMainFriend
func __GetMainFriend(p unsafe.Pointer) C.int {

	fmt.Println("__GetMainFriend")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	//fmt.Println(battleid, uid)

	battle := FindBattle(int64(battleid))

	var t_id int64
	t_id = battle.selectMainUnit(int64(uid), true)

	fmt.Println("__GetMainFriend end ,", t_id)

	L.PushInteger(int(t_id))

	return 1
}

//export __GetFriend
func __GetFriend(p unsafe.Pointer) C.int {

	//fmt.Println("__GetFriend")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	uid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(uid))

	t_id := battle.SelectNearFriend(unit.InstId)

	fmt.Println("__GetFriend end ,", t_id)

	L.PushInteger(int(t_id))

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

//export __ChangeCptProperty
func __ChangeCptProperty(p unsafe.Pointer) C.int {   //加Cpt减属性值

	//fmt.Println("__ChangeCptProperty")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	data := L.ToInteger(idx)
	idx ++
	property := L.ToString(idx)

	fmt.Println("__ChangeCptProperty", battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	battle.ChangeCptProperty(int64(unitid), int32(data), property)

	return 0
}
//export __ChangeIptProperty
func __ChangeIptProperty(p unsafe.Pointer) C.int {   //加 IPT减属性值

	//fmt.Println("__ChangeIptProperty")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	data := L.ToInteger(idx)
	idx ++
	property := L.ToString(idx)

	fmt.Println("__ChangeIptProperty", battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	battle.ChangeIptProperty(int64(unitid), int32(data), property)

	return 0
}
//export __AddSheld
func __AddSheld(p unsafe.Pointer) C.int {   //加护盾

	fmt.Println("__AddSheld")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	buffinstid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	unit.AddSpec("BF_SHELDNUM", int32(buffinstid))

	buff := unit.SelectBuff(int32(buffinstid))

	unit.VirtualHp += int32(buff.Data)

	return 0
}

//export __PopSheld
func __PopSheld(p unsafe.Pointer) C.int {   //减护盾

	fmt.Println("__PopSheld")

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

	unit.VirtualHp -= int32(buff.Data)

	if unit.VirtualHp < 0 {
		unit.VirtualHp = 0
	}

	return 0
}

//export __DamageSheld
func __DamageSheld(p unsafe.Pointer) C.int {   //减護盾值

	fmt.Println("__DamageSheld")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	damage := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))
	needdel := []int32{}

	for _, buffid := range unit.Special[prpc.BF_SHELDNUM] {
		buff := unit.SelectBuff(buffid)
		if buff == nil {
			needdel = append(needdel, buffid)
			continue
		}
		if damage > int(buff.Data) {
			damage -= int(buff.Data)
			buff.Over = true
			needdel = append(needdel, buffid)
			unit.VirtualHp -= buff.Data
			continue
		}
		unit.VirtualHp -= int32(damage)
	}

	for _, buffid := range needdel {
		unit.PopSpec("BF_SHELDNUM", buffid)
	}

	if unit.VirtualHp < 0 {
		unit.VirtualHp = 0
	}

	return 0
}

//export __ClacSheld
func __ClacSheld(p unsafe.Pointer) C.int {   //减伤

	fmt.Println("__ClacSheld")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	//fmt.Println("__ClacSheld",battleid,unitid)
	unit := battle.SelectOneUnit(int64(unitid))

	ClacSheld := unit.ClacSheldPer(int32(battle.Round))

	//fmt.Println("__ClacSheld1111",unit,ClacSheld)

	L.PushNumber(float64(ClacSheld))

	return 1
}

//export __ClacStrongPer
func __ClacStrongPer(p unsafe.Pointer) C.int {   //增输出伤比

	fmt.Println("__ClacStrongPer")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	//fmt.Println("__ClacStrongPer",battleid,unitid)
	unit := battle.SelectOneUnit(int64(unitid))

	ClacSheld := unit.ClacStrongPer(int32(battle.Round))

	//fmt.Println("__ClacStrongPer",unit,ClacSheld)

	L.PushNumber(float64(ClacSheld))

	return 1
}
//export __ClacWeakPer
func __ClacWeakPer(p unsafe.Pointer) C.int {   //增承受伤比

	fmt.Println("__ClacStrongPer")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	//fmt.Println("__ClacWeakPer",battleid,unitid)
	unit := battle.SelectOneUnit(int64(unitid))

	ClacSheld := unit.ClacWeakPer(int32(battle.Round))

	//fmt.Println("__ClacWeakPer",unit,ClacSheld)

	L.PushNumber(float64(ClacSheld))

	return 1
}

//export __ChangeSpecial
func __ChangeSpecial(p unsafe.Pointer) C.int {  //判断有无这个属性，有替换，么加上

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

	unit.AddSpec(spec, int32(buffinstid))

	return 0
}
//export __PopSpec
func __PopSpec(p unsafe.Pointer) C.int {  //

	fmt.Println("__PopSpec")

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

	fmt.Println("__PopSpec , unit:", unit)
	if unit == nil {
		return 0
	}
	unit.PopSpec(spec, int32(buffinstid))

	return 0
}
//export __GetSpecial
func  __GetSpecial(p unsafe.Pointer) C.int { //獲取spec相对应的buffid

	fmt.Println("__GetSpecial")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx++
	unitid := L.ToInteger(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffid := unit.GetSpecial(spec)

	L.NewTable()

	for i:=1;i<len(buffid);i++{

		L.PushInteger(i+1)
		L.PushInteger(int(buffid[i]))
		L.SetTable(-3)

	}

	return 1

}
//export __GetOneSpecial
func  __GetOneSpecial(p unsafe.Pointer) C.int { //獲取spec相对应的buffid  实例id

	fmt.Println("__GetOneSpecial")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx++
	unitid := L.ToInteger(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffid := unit.GetOneSpecial(spec, battle.Round)

	L.PushInteger(int(buffid))

	return 1

}

//export __GetSpecialData
func  __GetSpecialData(p unsafe.Pointer) C.int { //獲取spec相对应的buffid s数值

	fmt.Println("__GetSpecialData")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx++
	unitid := L.ToInteger(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffids := unit.GetSpecial(spec)
	fmt.Println("aaaaaaaaaaaaa", buffids)

	var data int32

	for _, buffid := range buffids {

		buff := unit.SelectBuff(buffid)

		data += buff.Data

	}

	L.PushInteger(int(data))

	return 1

}

//export __GetCheckSpec
func __GetCheckSpec(p unsafe.Pointer) C.int { //是否有特殊效果的buff

	fmt.Println("__GetCheckSpec")

	L := lua.GetLuaState(p)

	idx := 1

	battleid := L.ToInteger(idx)
	idx++
	unitid := L.ToInteger(idx)
	idx++
	spec := L.ToString(idx)

	//fmt.Println("__GetCheckSpec ",battleid,unitid)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	//fmt.Println("__GetCheckSpec 1111",battleid,unitid)

	_bool := unit.CheckSpec(spec, battle.Round)

	//fmt.Println("__GetCheckSpec 22222",_bool)

	L.PushBoolean(_bool)

	return 1
	
}

//export __GetBuffLockId
func __GetBuffLockId(p unsafe.Pointer) C.int { //是否有特殊效果的buff

	fmt.Println("__GetBuffLockId")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx++
	unitid := L.ToInteger(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffids := unit.GetSpecial(spec)
	fmt.Println("__GetBuffLockId", buffids)

	var data int32

	for _, buffid := range buffids {

		buff := unit.SelectBuff(buffid)

		if buff.Data == 0 {
			continue
		}
		data = buff.Data
		break
	}

	L.PushInteger(int(data))

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
//export __GetTargetsAround
func __GetTargetsAround(p unsafe.Pointer) C.int {

	//fmt.Println("__GetTargetsAround")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)


	////fmt.Println("4444444444", battleid, unitid, num)

	battle := FindBattle(int64(battleid))

	ls := battle.SelectAroundTraget(int64(unitid))

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

	ls := battle.SelectMoreFriend(int64(unitid), num)
	fmt.Println("__GetFriends", ls)

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

//export __FrontTarget
func __FrontTarget(p unsafe.Pointer) C.int {		//获取前排人数

	fmt.Println("__FrontTarget")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	FrontTarget := battle.SelectFrontTarget(int(unit.Camp))

	fmt.Println(unitid, "Front", len(FrontTarget), "info", FrontTarget)

	//L.PushInteger(int(num))

	L.NewTable()
	//L.PushInteger(-1)
	//L.RawSetI(-2, 0)

	for i :=0; i < len(FrontTarget); i++ {
		L.PushInteger(i + 1)
		L.PushInteger(int(FrontTarget[i]))
		L.SetTable(-3)
	}

	return 1
}

//export __LineTraget
func __LineTraget(p unsafe.Pointer) C.int {		//获取纵排人数

	fmt.Println("__LineTraget")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	LineTraget := battle.SelectLineTraget(int64(unitid))

	fmt.Println(unitid, "line", len(LineTraget), "info", LineTraget)

	//L.PushInteger(int(num))

	L.NewTable()
	//L.PushInteger(-1)
	//L.RawSetI(-2, 0)

	for i :=0; i < len(LineTraget); i++ {
		L.PushInteger(i + 1)
		L.PushInteger(int(LineTraget[i]))
		L.SetTable(-3)
	}

	return 1
}

//export __BackTarget
func __BackTarget(p unsafe.Pointer) C.int {		//获取后排人数

	fmt.Println("__BackTarget")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	BackTarget := battle.SelectBackTarget(int(unit.Camp))

	fmt.Println(unitid, "Back", len(BackTarget), "info", BackTarget)

	//L.PushInteger(int(num))

	L.NewTable()
	//L.PushInteger(-1)
	//L.RawSetI(-2, 0)

	for i :=0; i < len(BackTarget); i++ {
		L.PushInteger(i + 1)
		L.PushInteger(int(BackTarget[i]))
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

	return 0
}

//export __Cure
func __Cure(p unsafe.Pointer) C.int {

	fmt.Println("__Cure")

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

	fmt.Println("6666666666666", battleid, target, crit, damage)

	battle.AddHp(int64(target), int32(damage), int32(crit))

	return 0
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

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

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

	return 0
}

//export __HasBuff
func __HasBuff(p unsafe.Pointer) C.int {

	//fmt.Println("__HasBuff")

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

	//fmt.Println("__HasDebuff")

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

//export __AddSkillBuff
func __AddSkillBuff(p unsafe.Pointer) C.int {
	//fmt.Println("__AddSkillBuff")

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

	battle.AddSkillBuff(int64(casterid),int64(target),int32(buffid),int32(data))

	return 0
}

//export __BuffMintsHp
func __BuffMintsHp(p unsafe.Pointer) C.int {  //掉血

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
	//fmt.Println("__BuffMintsHp, ", buff.IsOver(battle.Round), !buff.IsOver(battle.Round))

	battle.BuffMintsHp(buff.CasterId, buff.Owner.InstId, buff.BuffId, buff.Data, !buff.IsOver(battle.Round))

	return 0
}

//export __BuffCureHp
func __BuffCureHp(p unsafe.Pointer) C.int {   //回血buff

	fmt.Println("__BuffCureHp")

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
	//fmt.Println("__BuffCureHp, ", buff.IsOver(battle.Round), !buff.IsOver(battle.Round))

	battle.BuffAddHp(buff.Owner.InstId, buff.BuffId, buff.Data, !buff.IsOver(battle.Round))

	return 0
}

//export __PopAllBuffByDebuff
func __PopAllBuffByDebuff(p unsafe.Pointer) C.int {		//驱散所有负面效果    返回负面buff数量

	fmt.Println("__PopAllBuffByDebuff")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	num := unit.PopAllBuffByDebuff()

	L.PushInteger(num)

	return 1
}

//export __PopAllBuffBybuff
func __PopAllBuffBybuff(p unsafe.Pointer) C.int {		//驱散所有增益buff效果

	fmt.Println("__PopAllBuffBybuff")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	num := unit.PopAllBuffByBuff()

	L.PushInteger(num)

	return 1
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//export __GetUnitDamage
func __GetUnitDamage(p unsafe.Pointer) C.int {    //物理  伤害

	fmt.Println("__GetUnitDamage")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	casterid := L.ToInteger(idx)
	idx ++
	targetid := L.ToInteger(idx)

	fmt.Println("battleid",battleid)
	fmt.Println("targetid",targetid)

	battle := FindBattle(int64(battleid))
	fmt.Println("battle",battle)

	caster := battle.SelectOneUnit(int64(casterid))
	fmt.Println("caster",caster)

	target := battle.SelectOneUnit(int64(targetid))
	fmt.Println("target",target)

	finaldamage := CalcDamage(caster, target)

	fmt.Println(finaldamage)

	L.PushNumber(float64(finaldamage))

	return 1
}

//export __GetMagicDamage
func __GetMagicDamage(p unsafe.Pointer) C.int {    //法术   伤害

	fmt.Println("__GetMagicDamage")


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

	finaldamage := CalcMagicDamage(caster, target)

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
		for _, buffid := range u.Special[prpc.BF_SHELDNUM] {
			buff := u.SelectBuff(buffid)
			buff.Over = true
		}
	}

	L.PushInteger(int(sheld))

	return 1
}

//export __GetOneSheld
func __GetOneSheld(p unsafe.Pointer) C.int {	// 获取场上单个玩家护盾数值

	fmt.Println("__GetOneSheld")

	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	sheld := unit.VirtualHp

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
func __GetUnitAtk(p unsafe.Pointer) C.int {		//获取减伤百分比  物理强度

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

	fmt.Println("__GetCalcDef")

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
func __GetUnitMtk(p unsafe.Pointer) C.int {		//获取减伤百分比   法术强度

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

	fmt.Println("__GetCalcMagicDef")

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
//export __TargetOver
func __TargetOver(p unsafe.Pointer) C.int {		//结束后

	fmt.Println("__TargetOver")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	battle.TargetOver()

	return 0
}
//export __TargetOn
func __TargetOn(p unsafe.Pointer) C.int {		//开始前清理数据

	fmt.Println("__TargetOn")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))

	battle.TargetOn()

	return 0
}

//export __ChangeBuffTimes
func __ChangeBuffTimes(p unsafe.Pointer) C.int {		//开始前清理数据

	fmt.Println("__ChangeBuffTimes")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(unitid))

	unit.ChangeBuffTimes(battle.Round)

	return 0
}

//export __BuffUpdate
func __BuffUpdate(p unsafe.Pointer) C.int {		//开始前清理数据

	fmt.Println("__BuffUpdate")

	L := lua.GetLuaState(p)

	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	spe := L.ToString(idx)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(unitid))

	unit.MustUpdateBuff(spe, battle.Round)

	//buff := unit.SelectBuff(int32(buffinstid))
	//buff.MustUpdate()

	return 0
}
//export __BuffChangeStillData
func __BuffChangeStillData(p unsafe.Pointer) C.int {
	//开始前清理数据
	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	new_data := L.ToInteger(idx)

	fmt.Println("__BuffChangeStillData")
	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(unitid))

	for _, buff := range unit.Allbuff {
		if buff.BuffKind != kKindStill {
			continue
		}

		buff.Data =int32(new_data)
	}

	//buff := unit.SelectBuff(int32(buffinstid))
	//buff.MustUpdate()

	return 0
}
//export __BuffChangeData
func __BuffChangeData(p unsafe.Pointer) C.int {
	//开始前清理数据
	L := lua.GetLuaState(p)
	idx := 1
	battleid := L.ToInteger(idx)
	idx ++
	unitid := L.ToInteger(idx)
	idx ++
	new_data := L.ToInteger(idx)

	fmt.Println("__BuffChangeData")
	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(unitid))

	for _, buff := range unit.Allbuff {
		if buff.BuffType != kTypeBuff {
			continue
		}

		buff.Data = buff.Data + buff.Data*(int32(new_data))

		fmt.Println("__BuffChangeData  battleid",battleid,"unitid",unitid,"buff.Data",buff.Data)
	}

	//buff := unit.SelectBuff(int32(buffinstid))
	//buff.MustUpdate()

	return 0
}

////////////////////////////////////////////////////////////////////////////////////////////
//export __GetMyUnitIProperty
func __GetMyUnitIProperty(p unsafe.Pointer) C.int {
	L := lua.GetLuaState(p)

	idx := 1
	casterId := L.ToLong(idx)
	idx ++
	ipc := L.ToString(idx)

	player := FindPlayerByInstId(casterId)
	if player == nil {
		fmt.Println("__GetMyUnitIProperty FindPlayerByInstId==nil",casterId)
		return 1
	}

	pd := prpc.ToId_IPropertyType(ipc)

	val := player.MyUnit.GetIProperty(int32(pd))

	L.PushInteger(int(val))

	return 1
}
//export __AddMyUnitEnergy
func __AddMyUnitEnergy(p unsafe.Pointer) C.int {

	L := lua.GetLuaState(p)

	idx := 1
	casterId := L.ToLong(idx)
	idx ++
	val := L.ToInteger(idx)

	player := FindPlayerByInstId(casterId)
	if player == nil {
		fmt.Println("__AddMyUnitEnergy FindPlayerByInstId==nil",casterId)
		return 0
	}

	player.SetMyEnergy(int32(val),true)

	return 0
}



