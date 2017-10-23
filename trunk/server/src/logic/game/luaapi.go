package game


import (

	"time"
	"logic/prpc"
	"logic/std"
	"github.com/yuin/gopher-lua"
)

var (
	_L *lua.LState
	_R string

	EnvIntegers			= make(map[string]int)
	EnvStrings			= make(map[string]string)
	GlobalIntegers		= make(map[string]int)
	GlobalStrings		= make(map[string]string)
)


func LoadFile(L *lua.LState, file string){
	err := L.DoFile(file)

	if err != nil {
		std.LogFatal("Load lua file %s", err.Error())
	}
}

func CallLuaFunc(fn string, inParams []interface{}, outParams *[]interface{}) error{

	p := lua.P{
		Fn:_L.GetGlobal(fn),
		NRet:len(*outParams),
		Protect:true,
	}

	a := make([]lua.LValue,len(inParams))

	for i, arg := range inParams {
		switch arg.(type) {
		case int:
			a[i] = lua.LNumber(arg.(int))
			break
		case int64:
			a[i] = lua.LNumber(arg.(int64))
			break
		case float64:
			a[i] = lua.LNumber(arg.(float64))
			break
		case string:
			a[i] = lua.LString(arg.(string))
			break
		case bool:
			a[i] = lua.LBool(arg.(bool))
		default:
			panic("cant not use lua params")
			break
		}
	}


	e := _L.CallByParam(p,a...)

	if e != nil{
		return  e
	}

	resLen := len(*outParams)
	for i := 0; i < resLen; i++ {
		idx := i-resLen
		switch (*outParams)[i].(type) {
		case int:
			(*outParams)[i] = _L.ToInt(idx)
			break
		case int64:
			(*outParams)[i] = _L.ToInt64(idx)
			break
		case float64:
			(*outParams)[i] = _L.ToNumber(idx)
			break
		case string:
			(*outParams)[i] = _L.ToString(idx)
			break
		case bool:
			(*outParams)[i] = _L.ToBool(idx)
		default:
			panic("cant not use lua params rsuly")
			break
		}
	}

	_L.Pop(resLen)

	return  nil
}

func __luaLog(L *lua.LState) int {
	str := L.ToString(-1)
	std.LogInfo("{{ %s",str)
	return 0
}

func __luaError(L *lua.LState) int {
	str := L.ToString(-1)
	std.LogError("{{ %s",str)
	return 0
}

func __luaFatal(L *lua.LState) int {
	str := L.ToString(-1)
	std.LogFatal("{{ %s",str)
	return 0
}

var sysMod = map[string]lua.LGFunction{
	"log"                :__luaLog                ,
	"err"                :__luaError              ,
	"fatal"              :__luaFatal              ,
	"loadfile"           :__loadfile              ,
	"DefineCards"        :__DefineCards           ,
	"GetTime"            :__GetTime               ,
}

var playerMod = map[string]lua.LGFunction{
	"GetStrings"         :__GetStrings            ,
	"GetFriend"          :__GetFriend             ,
	"GetMainFriend"      :__GetMainFriend         ,
	"CheckUnitDead"      :__CheckUnitDead         ,
	"GetFriends"         :__GetFriends            ,
	"GetTarget"          :__GetTarget             ,
	"GetMainTarget"      :__GetMainTarget         ,
	"RandomTarget"       :__GetRandomTarget       ,
	"GetTargets"         :__GetTargets            ,
	"GetTargetsAround"   :__GetTargetsAround      ,
	"GetTargetsRandom"   :__GetTargetsRandom      ,
	"GetUnitProperty"    :__GetUnitProperty       ,
	"ChangeUnitProperty" :__ChangeCptProperty     ,
	"ChangeIptProperty"  :__ChangeIptProperty     ,
	"AddSheld"           :__AddSheld              ,
	"PopSheld"           :__PopSheld              ,
	"DownSheld"          :__DamageSheld           ,
	"ChangeSpecial"      :__ChangeSpecial         ,
	"PopSpec"            :__PopSpec               ,
	"GetSpecial"         :__GetSpecial            ,
	"GetOneSpecial"      :__GetOneSpecial         ,
	"GetCheckSpec"       :__GetCheckSpec          ,
	"GetBuffLockId"      :__GetBuffLockId         ,
	"GetUnitMtk"         :__GetUnitMtk            ,
	"GetCalcMagicDef"    :__GetCalcMagicDef       ,
	"GetUnitAtk"         :__GetUnitAtk            ,
	"GetCalcDef"         :__GetCalcDef            ,
	"GetUnitDamage"      :__GetUnitDamage         ,
	"GetMagicDamage"     :__GetMagicDamage        ,
	"ClacSheld"          :__ClacSheld             ,
	"PopAllBuffByDebuff" :__PopAllBuffByDebuff    ,
	"PopAllBuffBybuff"   :__PopAllBuffBybuff      ,
	"GetUnitSheld"       :__GetUnitSheld          ,
	"FrontTarget"        :__FrontTarget           ,
	"LineTraget"         :__LineTraget            ,
	"BackTarget"         :__BackTarget            ,
	"GetOneSheld"        :__GetOneSheld           ,
	"GetSpecialData"     :__GetSpecialData        ,
	"ClacWeakPer"        :__ClacWeakPer           ,
	"ClacStrongPer"      :__ClacStrongPer         ,
	"ChangeBuffTimes"    :__ChangeBuffTimes       ,
	"GetMyUnitIProperty" :__GetMyUnitIProperty    ,
	"AddMyUnitEnergy"    :__AddMyUnitEnergy       ,
	"ThrowCard"          :__ThrowCard             ,
	"Throw"              :__Throw                 ,

}

var battleMod = map[string]lua.LGFunction{
	"Attack"             :__Attack                ,
	"Cure"               :__Cure                  ,
	"GetCrit"            :__GetCrit               ,
	"AddBuff"            :__AddBuff               ,
	"AddSkillBuff"       :__AddSkillBuff          ,
	"HasBuff"            :__HasBuff               ,
	"HasDebuff"          :__HasDebuff             ,
	"BuffMintsHp"        :__BuffMintsHp           ,
	"BuffCureHp"         :__BuffCureHp            ,
	"BuffUpdate"         :__BuffUpdate            ,
	"BuffChangeStillData":__BuffChangeStillData   ,
	"BuffChangeData"     :__BuffChangeData        ,
	"TargetOver"         :__TargetOver            ,
	"TargetOn"           :__TargetOn              ,

}


var envMod = map[string]lua.LGFunction{
	"setEnvString"       :__setEnvString          ,
	"setEnvInt"          :__setEnvInt             ,
}

var globalMod = map[string]lua.LGFunction{
	"setGlobalString"    :__setGlobalString       ,
	"setGlobalInt"       :__setGlobalInt          ,

}

func luaOpenSys(L*lua.LState) int {
	mod := _L.RegisterModule("sys", sysMod).(*lua.LTable)
	_L.Push(mod)
	return 1
}

func luaOpenEnv(L*lua.LState) int {
	mod := _L.RegisterModule("Env", envMod).(*lua.LTable)
	_L.Push(mod)
	return 1
}

func luaOpenGlo(L*lua.LState) int {
	mod := _L.RegisterModule("Global", globalMod).(*lua.LTable)
	_L.Push(mod)
	return 1
}

func luaOpenPla(L*lua.LState) int {
	mod := _L.RegisterModule("Player", playerMod).(*lua.LTable)
	_L.Push(mod)
	return 1
}

func luaOpenBatt(L*lua.LState) int {
	mod := _L.RegisterModule("Battle", battleMod).(*lua.LTable)
	_L.Push(mod)
	return 1
}

type luaLib struct {
	libName string
	libFunc lua.LGFunction
}

var luaLibs = []luaLib{
	luaLib{"env", luaOpenSys},
	luaLib{"sys", luaOpenEnv},
	luaLib{"global", luaOpenGlo},
	luaLib{"player", luaOpenPla},
	luaLib{"battle", luaOpenBatt},
}



func luaOpnApi(){
	for _, lib := range luaLibs {
		_L.Push(_L.NewFunction(lib.libFunc))
		_L.Push(lua.LString(lib.libName))
		_L.Call(1, 0)
	}

}

func InitLua(r string){
	_R = r
	_L = lua.NewState()
	_L.OpenLibs()
	luaOpnApi()

	LoadFile(_L,_R + "main.lua")

	LoadFile(_L,_R + "env.lua")
	LoadFile(_L,_R + "serGlobal.lua")
}


//////////////////////////////////////////////////////////////////////////////////////

//export __setEnvString
func __setEnvString(L* lua.LState) int {
	idx := 1
	s1 := L.ToString(idx)
	idx++
	s2 := L.ToString(idx)
	//fmt.Println("Test Env Set String ===>",s1,s2)
	EnvStrings[s1] = s2

	return 0
}

//export __setEnvInt
func __setEnvInt(L* lua.LState) int {
	idx := 1
	s1 := L.ToString(idx)
	idx++
	s2 := L.ToInt(idx)
	//fmt.Println("Test Env Set Int ===>",s1,s2)
	EnvIntegers[s1] = s2

	return 0
}

//export __setGlobalString
func __setGlobalString(L* lua.LState) int {
	
	idx := 1
	s1 := L.ToString(idx)
	idx++
	s2 := L.ToString(idx)
	GlobalStrings[s1] = s2

	return 0
}

//export __setGlobalInt
func __setGlobalInt(L* lua.LState) int {
	
	idx := 1
	s1 := L.ToString(idx)
	idx++
	s2 := L.ToInt(idx)
	GlobalIntegers[s1] = s2

	return 0
}

func GetEnvString(val string) string {
	return EnvStrings[val]
}

func GetEnvInt(val string) int {
	return EnvIntegers[val]
}

func GetGlobalString(val string) string {
	return GlobalStrings[val]
}

func GetGlobalInt(val string) int {
	return GlobalIntegers[val]
}

//////////////////////////////////////////////////////////////////////////////////////

//export __loadfile
func __loadfile(L* lua.LState) int {

	
	fileName := L.ToString(-1)
	LoadFile(L, _R + fileName)
	return 0
}

//export __GetStrings
func __GetStrings(L* lua.LState) int {

	
	idx:= 1

	i := L.ToInt(idx)

	std.LogInfo("__GetStrings", int32(i))

	//std.LogInfo("__GetStrings")

	return 0
}

//export __DefineCards
func __DefineCards(L* lua.LState) int {

	
	idx:= 1

	cards := L.ToString(idx)
	SetDefaultUnits(cards)

	//std.LogInfo("__GetStrings")`

	return 0
}

//export __GetTarget
func __GetTarget(L* lua.LState) int { //获取 敌方单个目标

	std.LogInfo("__GetTarget")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	uid := L.ToInt(idx)

	//std.LogInfo(battleid, uid)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(uid))

	t_id := battle.SelectNearTarget(unit.InstId)

	std.LogInfo("__GetTarget end ,", t_id)

	L.Push(lua.LNumber(t_id))

	return 1
}

//export __GetMainTarget
func __GetMainTarget(L* lua.LState) int {// 获取 敌方主角目标

	std.LogInfo("__GetMainTarget")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	uid := L.ToInt(idx)

	//std.LogInfo(battleid, uid)

	battle := FindBattle(int64(battleid))

	var t_id int64
	t_id = battle.selectMainUnit(int64(uid), false)

	std.LogInfo("__GetMainTarget end ,", t_id)

	L.Push(lua.LNumber(t_id))

	return 1
}

//export __GetRandomTarget
func __GetRandomTarget(L* lua.LState) int {// 获取 敌方主角目标

	std.LogInfo("__GetRandomTarget")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	uid := L.ToInt(idx)

	//std.LogInfo(battleid, uid)

	battle := FindBattle(int64(battleid))

	var t_id int64
	t_id = battle.SelectOneTarget(int64(uid))

	std.LogInfo("__GetMainTarget end ,", t_id)

	L.Push(lua.LNumber(t_id))

	return 1
}


//export __CheckUnitDead
func __CheckUnitDead(L* lua.LState) int {//判断是否死亡

	std.LogInfo("__CheckUnitDead")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	uid := L.ToInt(idx)

	//std.LogInfo(battleid, uid)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(uid))

	var t_id bool
	if unit.IsDead() {
		t_id = true
	}else{
		t_id = false
	}
	L.Push(lua.LBool(t_id))

	return 1
}
//export __GetMainFriend
func __GetMainFriend(L* lua.LState) int {  //友方主角
	std.LogInfo("__GetMainFriend")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	uid := L.ToInt(idx)

	//std.LogInfo(battleid, uid)

	battle := FindBattle(int64(battleid))

	var t_id int64
	t_id = battle.selectMainUnit(int64(uid), true)

	std.LogInfo("__GetMainFriend end ,", t_id)

	L.Push(lua.LNumber(t_id))

	return 1
}

//export __GetFriend
func __GetFriend(L* lua.LState) int {//友方单个目标

	//std.LogInfo("__GetFriend")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	uid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(uid))

	t_id := battle.SelectNearFriend(unit.InstId)

	std.LogInfo("__GetFriend end ,", t_id)

	L.Push(lua.LNumber(t_id))

	return 1
}

//export __GetUnitProperty
func __GetUnitProperty(L* lua.LState) int {//获取属性值

	//std.LogInfo("__GetUnitProperty")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	property := L.ToString(idx)

	//std.LogInfo(battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	date := battle.GetUnitProperty(int64(unitid), property)

	L.Push(lua.LNumber(date))

	return 1
}

//export __ChangeCptProperty
func __ChangeCptProperty(L* lua.LState) int {   //加Cpt减属性值

	//std.LogInfo("__ChangeCptProperty")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	data := L.ToInt(idx)
	idx ++
	property := L.ToString(idx)

	std.LogInfo("__ChangeCptProperty", battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	battle.ChangeCptProperty(int64(unitid), int32(data), property)

	return 0
}
//export __ChangeIptProperty
func __ChangeIptProperty(L* lua.LState) int {   //加 IPT减属性值

	//std.LogInfo("__ChangeIptProperty")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	data := L.ToInt(idx)
	idx ++
	property := L.ToString(idx)

	std.LogInfo("__ChangeIptProperty", battleid, unitid, property)

	battle := FindBattle(int64(battleid))

	battle.ChangeIptProperty(int64(unitid), int32(data), property)

	return 0
}
//export __AddSheld
func __AddSheld(L* lua.LState) int {   //加护盾

	std.LogInfo("__AddSheld")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	buffinstid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	unit.AddSpec("BF_SHELDNUM", int32(buffinstid))

	buff := unit.SelectBuff(int32(buffinstid))

	unit.VirtualHp += int32(buff.Data)

	return 0
}

//export __PopSheld
func __PopSheld(L* lua.LState) int {   //减护盾

	std.LogInfo("__PopSheld")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	buffinstid := L.ToInt(idx)

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
func __DamageSheld(L* lua.LState) int {   //减護盾值

	std.LogInfo("__DamageSheld")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	damage := L.ToInt(idx)

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
func __ClacSheld(L* lua.LState) int {   //减伤

	std.LogInfo("__ClacSheld")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	//std.LogInfo("__ClacSheld",battleid,unitid)
	unit := battle.SelectOneUnit(int64(unitid))

	ClacSheld := unit.ClacSheldPer(int32(battle.Round))

	//std.LogInfo("__ClacSheld1111",unit,ClacSheld)

	L.Push(lua.LNumber(ClacSheld))

	return 1
}

//export __ClacStrongPer
func __ClacStrongPer(L* lua.LState) int {   //增输出伤比

	std.LogInfo("__ClacStrongPer")

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	//std.LogInfo("__ClacStrongPer",battleid,unitid)
	unit := battle.SelectOneUnit(int64(unitid))

	ClacSheld := unit.ClacStrongPer(int32(battle.Round))

	//std.LogInfo("__ClacStrongPer",unit,ClacSheld)

	L.Push(lua.LNumber(ClacSheld))

	return 1
}
//export __ClacWeakPer
func __ClacWeakPer(L* lua.LState) int {   //增承受伤比

	std.LogInfo("__ClacStrongPer")

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	//std.LogInfo("__ClacWeakPer",battleid,unitid)
	unit := battle.SelectOneUnit(int64(unitid))

	ClacSheld := unit.ClacWeakPer(int32(battle.Round))

	//std.LogInfo("__ClacWeakPer",unit,ClacSheld)

	L.Push(lua.LNumber(ClacSheld))

	return 1
}

//export __ChangeSpecial
func __ChangeSpecial(L* lua.LState) int {  //判断有无这个属性，有替换，么加上

	std.LogInfo("__ChangeSpecial")

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	buffinstid := L.ToInt(idx)
	idx ++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	unit.AddSpec(spec, int32(buffinstid))

	return 0
}
//export __PopSpec
func __PopSpec(L* lua.LState) int {  //删除buff

	std.LogInfo("__PopSpec")

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	buffinstid := L.ToInt(idx)
	idx ++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	std.LogInfo("__PopSpec , unit:", unit)
	if unit == nil {
		return 0
	}
	unit.PopSpec(spec, int32(buffinstid))

	return 0
}
//export __GetSpecial
func  __GetSpecial(L* lua.LState) int { //獲取spec相对应的buffid

	std.LogInfo("__GetSpecial")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx++
	unitid := L.ToInt(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffid := unit.GetSpecial(spec)

	table := L.NewTable()

	for i:=1;i<len(buffid);i++{
		table.Insert(i,lua.LNumber(buffid[i]))
	}

	L.Push(table)

	return 1

}
//export __GetOneSpecial
func  __GetOneSpecial(L* lua.LState) int { //獲取spec相对应的buffid  实例id

	std.LogInfo("__GetOneSpecial")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx++
	unitid := L.ToInt(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffid := unit.GetOneSpecial(spec, battle.Round)

	L.Push(lua.LNumber(buffid))

	return 1

}

//export __GetSpecialData
func  __GetSpecialData(L* lua.LState) int { //獲取spec相对应的buffid s数值

	std.LogInfo("__GetSpecialData")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx++
	unitid := L.ToInt(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffids := unit.GetSpecial(spec)
	std.LogInfo("aaaaaaaaaaaaa", buffids)

	var data int32

	for _, buffid := range buffids {

		buff := unit.SelectBuff(buffid)

		data += buff.Data

	}

	L.Push(lua.LNumber(data))

	return 1

}

//export __GetCheckSpec
func __GetCheckSpec(L* lua.LState) int { //是否有特殊效果的buff

	std.LogInfo("__GetCheckSpec")

	

	idx := 1

	battleid := L.ToInt(idx)
	idx++
	unitid := L.ToInt(idx)
	idx++
	spec := L.ToString(idx)

	//std.LogInfo("__GetCheckSpec ",battleid,unitid)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	//std.LogInfo("__GetCheckSpec 1111",battleid,unitid)

	_bool := unit.CheckSpec(spec, battle.Round)

	//std.LogInfo("__GetCheckSpec 22222",_bool)

	L.Push(lua.LBool(_bool))

	return 1
	
}

//export __GetBuffLockId
func __GetBuffLockId(L* lua.LState) int { //是否有特殊效果的buff

	std.LogInfo("__GetBuffLockId")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx++
	unitid := L.ToInt(idx)
	idx++
	spec := L.ToString(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buffids := unit.GetSpecial(spec)
	std.LogInfo("__GetBuffLockId", buffids)

	var data int32

	for _, buffid := range buffids {

		buff := unit.SelectBuff(buffid)

		if buff.Data == 0 {
			continue
		}
		data = buff.Data
		break
	}

	L.Push(lua.LNumber(data))

	return 1

}

//export __GetTargets
func __GetTargets(L* lua.LState) int {  //获取敌方多个目标

	//std.LogInfo("__GetTargets")


	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	num := L.ToInt(idx)

	////std.LogInfo("4444444444", battleid, unitid, num)

	battle := FindBattle(int64(battleid))

	ls := battle.SelectMoreTarget(int64(unitid), num)


	table := L.NewTable()

	for i:=0;i<len(ls);i++{
		table.Insert(i+1,lua.LNumber(ls[i]))
	}

	L.Push(table)

	return 1
}
//export __GetTargetsAround
func __GetTargetsAround(L* lua.LState) int {  //溅射目标

	//std.LogInfo("__GetTargetsAround")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)


	////std.LogInfo("4444444444", battleid, unitid, num)

	battle := FindBattle(int64(battleid))

	ls := battle.SelectAroundTraget(int64(unitid))


	//L.Push(-1)
	//L.RawSetI(-2, 0)

	table := L.NewTable()

	for i:=0;i<len(ls);i++{
		table.Insert(i+1,lua.LNumber(ls[i]))
	}

	L.Push(table)

	return 1
}

//export __GetTargetsRandom
func __GetTargetsRandom(L* lua.LState) int {  //溅射目标

	//std.LogInfo("__GetTargetsRandom")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	targetnum := L.ToInt(idx)


	////std.LogInfo("4444444444", battleid, unitid, num)

	battle := FindBattle(int64(battleid))

	ls := battle.SelectRandomTarget(int64(unitid), int32(targetnum))
	//std.LogInfo("4444444444",ls)
	table := L.NewTable()

	for i:=0;i<len(ls);i++{
		table.Insert(i+1,lua.LNumber(ls[i]))
	}

	L.Push(table)

	return 1
}

//export __GetFriends
func __GetFriends(L* lua.LState) int {

	//std.LogInfo("__GetTargets")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	num := L.ToInt(idx)

	////std.LogInfo("4444444444", battleid, unitid, num)

	battle := FindBattle(int64(battleid))

	ls := battle.SelectMoreFriend(int64(unitid), num)
	std.LogInfo("__GetFriends", ls)

	table := L.NewTable()

	for i:=0;i<len(ls);i++{
		table.Insert(i+1,lua.LNumber(ls[i]))
	}

	L.Push(table)

	return 1
}

//export __FrontTarget
func __FrontTarget(L* lua.LState) int {		//获取前排人数

	std.LogInfo("__FrontTarget")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	FrontTarget := battle.SelectFrontTarget(int(unit.Camp))

	std.LogInfo(string(unitid), "Front", len(FrontTarget), "info", FrontTarget)

	//L.Push(int(num))


	table := L.NewTable()

	for i:=0;i<len(FrontTarget);i++{
		table.Insert(i+1,lua.LNumber(FrontTarget[i]))
	}

	L.Push(table)

	return 1
}

//export __LineTraget
func __LineTraget(L* lua.LState) int {		//获取纵排人数

	std.LogInfo("__LineTraget")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	lineTraget := battle.SelectLineTraget(int64(unitid))

	std.LogInfo(string(unitid), "line", len(lineTraget), "info", lineTraget)

	//L.Push(int(num))

	table := L.NewTable()

	for i:=0;i<len(lineTraget);i++{
		table.Insert(i+1,lua.LNumber(lineTraget[i]))
	}

	L.Push(table)

	return 1
}

//export __BackTarget
func __BackTarget(L* lua.LState) int {		//获取后排人数

	std.LogInfo("__BackTarget")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	BackTarget := battle.SelectBackTarget(int(unit.Camp))

	std.LogInfo(string(unitid), "Back", len(BackTarget), "info", BackTarget)

	//L.Push(int(num))

	table := L.NewTable()

	for i:=0;i<len(BackTarget);i++{
		table.Insert(i+1,lua.LNumber(BackTarget[i]))
	}

	L.Push(table)

	return 1
}

//export __Attack
func __Attack(L* lua.LState) int {

	//std.LogInfo("__Attack battleid")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	caster := L.ToInt(idx)
	idx ++
	target := L.ToInt(idx)
	idx ++
	damage := L.ToInt(idx)
	idx ++
	crit := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	battle.MintsHp(int64(caster), int64(target), int32(damage), int32(crit))

	//std.LogInfo("55555555555555", battleid, caster, target, crit, damage)

	return 0
}

//export __Cure
func __Cure(L* lua.LState) int {

	std.LogInfo("__Cure")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	target := L.ToInt(idx)
	idx ++
	damage := L.ToInt(idx)
	idx ++
	crit := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	std.LogInfo("6666666666666", battleid, target, crit, damage)

	battle.AddHp(int64(target), int32(damage), int32(crit))

	return 0
}

//export __GetCrit
func __GetCrit(L* lua.LState) int {   //判断暴击

	
	idx := 1
	skillid := L.ToInt(idx)

	crit := IsCrit(int32(skillid))

	L.Push(lua.LNumber(crit))

	return 1
}

//export __GetTime
func __GetTime(L* lua.LState) int {

	//std.LogInfo("__GetTime")

	

	time_unix := time.Now().Unix()

	L.Push(lua.LNumber(time_unix))

	return 1
}

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

//export __AddBuff
func __AddBuff(L* lua.LState) int {

	//std.LogInfo("__AddBuff")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	casterid := L.ToInt(idx)
	idx ++
	target := L.ToInt(idx)
	idx ++
	buffid := L.ToInt(idx)
	idx ++
	data := L.ToInt(idx)

	battle := FindBattle(int64(battleid))


	battle.AddBuff(int64(casterid), int64(target), int32(buffid), int32(data))

	return 0
}

//export __HasBuff
func __HasBuff(L* lua.LState) int {  //是否有增益buff

	//std.LogInfo("__HasBuff")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	target := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	//battle.AddBuff(int64(casterid), int64(target), int32(buffid), int32(data))
	has := battle.HasBuff(int64(target))

	L.Push(lua.LBool(has))

	return 1
}

//export __HasDebuff
func __HasDebuff(L* lua.LState) int { //是否delbuff

	//std.LogInfo("__HasDebuff")

	
	idx := 1
	battleid := L.ToInt(idx)

	idx ++
	target := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	//battle.AddBuff(int64(casterid), int64(target), int32(buffid), int32(data))
	has := battle.HasDebuff(int64(target))

	L.Push(lua.LBool(has))

	return 1
}

//export __AddSkillBuff
func __AddSkillBuff(L* lua.LState) int {
	//std.LogInfo("__AddSkillBuff")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	casterid := L.ToInt(idx)
	idx ++
	target := L.ToInt(idx)
	idx ++
	buffid := L.ToInt(idx)
	idx ++
	data := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	battle.AddSkillBuff(int64(casterid),int64(target),int32(buffid),int32(data))

	return 0
}

//export __BuffMintsHp
func __BuffMintsHp(L* lua.LState) int {  //掉血

	std.LogInfo("__BuffMintsHp")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	buffinstid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buff := unit.SelectBuff(int32(buffinstid))
	//std.LogInfo("__BuffMintsHp, ", buff.IsOver(battle.Round), !buff.IsOver(battle.Round))

	battle.BuffMintsHp(buff.CasterId, buff.Owner.InstId, buff.BuffId, buff.Data, !buff.IsOver(battle.Round))

	return 0
}

//export __BuffCureHp
func __BuffCureHp(L* lua.LState) int {   //回血buff

	std.LogInfo("__BuffCureHp")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	buffinstid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	buff := unit.SelectBuff(int32(buffinstid))
	//std.LogInfo("__BuffCureHp, ", buff.IsOver(battle.Round), !buff.IsOver(battle.Round))

	battle.BuffAddHp(buff.Owner.InstId, buff.BuffId, buff.Data, !buff.IsOver(battle.Round))

	return 0
}

//export __PopAllBuffByDebuff
func __PopAllBuffByDebuff(L* lua.LState) int {		//驱散所有负面效果    返回负面buff数量

	std.LogInfo("__PopAllBuffByDebuff")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	num := unit.PopAllBuffByDebuff()

	L.Push(lua.LNumber(num))

	return 1
}

//export __PopAllBuffBybuff
func __PopAllBuffBybuff(L* lua.LState) int {		//驱散所有增益buff效果

	std.LogInfo("__PopAllBuffBybuff")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	num := unit.PopAllBuffByBuff()

	L.Push(lua.LNumber(num))

	return 1
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//export __GetUnitDamage
func __GetUnitDamage(L* lua.LState) int {    //物理  伤害

	std.LogInfo("__GetUnitDamage")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	casterid := L.ToInt(idx)
	idx ++
	targetid := L.ToInt(idx)

	std.LogInfo("battleid",battleid)
	std.LogInfo("targetid",targetid)

	battle := FindBattle(int64(battleid))
	std.LogInfo("battle",battle)

	caster := battle.SelectOneUnit(int64(casterid))
	std.LogInfo("caster",caster)

	target := battle.SelectOneUnit(int64(targetid))
	std.LogInfo("target",target)

	finaldamage := CalcDamage(caster, target)

	std.LogInfo(string(int(finaldamage)))

	L.Push(lua.LNumber(finaldamage))

	return 1
}

//export __GetMagicDamage
func __GetMagicDamage(L* lua.LState) int {    //法术   伤害

	std.LogInfo("__GetMagicDamage")


	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	casterid := L.ToInt(idx)
	idx ++
	targetid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	caster := battle.SelectOneUnit(int64(casterid))

	target := battle.SelectOneUnit(int64(targetid))

	finaldamage := CalcMagicDamage(caster, target)

	L.Push(lua.LNumber(finaldamage))

	return 1
}

//export __GetUnitSheld
func __GetUnitSheld(L* lua.LState) int {	// 获取场上所有玩家护盾数值

	std.LogInfo("__GetUnitSheld")

	
	idx := 1
	battleid := L.ToInt(idx)

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

	L.Push(lua.LNumber(sheld))

	return 1
}

//export __GetOneSheld
func __GetOneSheld(L* lua.LState) int {	// 获取场上单个玩家护盾数值

	std.LogInfo("__GetOneSheld")

	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	sheld := unit.VirtualHp

	L.Push(lua.LNumber(sheld))

	return 1
}

//export __GetUnitSheldPer
func __GetUnitSheldPer(L* lua.LState) int {		//获取减伤百分比

	std.LogInfo("__GetUnitSheldPer")

	
	L.Push(lua.LNumber(1))

	return 1
}

//export __GetUnitAtk
func __GetUnitAtk(L* lua.LState) int {		//获取减伤百分比  物理强度

	std.LogInfo("__GetUnitAtk")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	atk := CalcAtk(unit)
	L.Push(lua.LNumber(atk))

	return 1
}


//export __GetCalcDef
func __GetCalcDef(L* lua.LState) int {		//获取减伤百分比  物理防御

	std.LogInfo("__GetCalcDef")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	atk := CalcDef(unit)

	L.Push(lua.LNumber(atk))

	return 1
}
//export __GetUnitMtk
func __GetUnitMtk(L* lua.LState) int {		//获取减伤百分比   法术强度

	std.LogInfo("__GetUnitMtk")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	mtk := CalcMagicAtk(unit)

	L.Push(lua.LNumber(mtk))

	return 1
}
//export __GetCalcMagicDef
func __GetCalcMagicDef(L* lua.LState) int {		//获取减伤百分比   法术 防御

	std.LogInfo("__GetCalcMagicDef")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	unit := battle.SelectOneUnit(int64(unitid))

	mtk := CalcMagicDef(unit)

	L.Push(lua.LNumber(mtk))

	return 1
}
//export __TargetOver
func __TargetOver(L* lua.LState) int {		//结束后

	std.LogInfo("__TargetOver")

	

	idx := 1
	battleid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	battle.TargetOver()

	return 0
}
//export __TargetOn
func __TargetOn(L* lua.LState) int {		//开始前清理数据

	std.LogInfo("__TargetOn")

	

	idx := 1
	battleid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))

	battle.TargetOn()

	return 0
}

//export __ChangeBuffTimes
func __ChangeBuffTimes(L* lua.LState) int {		//开始前清理数据

	std.LogInfo("__ChangeBuffTimes")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)

	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(unitid))

	unit.ChangeBuffTimes(battle.Round)

	return 0
}

//export __BuffUpdate
func __BuffUpdate(L* lua.LState) int {		//开始前清理数据

	std.LogInfo("__BuffUpdate")

	

	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
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
func __BuffChangeStillData(L* lua.LState) int {
	//开始前清理数据
	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	new_data := L.ToInt(idx)

	std.LogInfo("__BuffChangeStillData")
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
func __BuffChangeData(L* lua.LState) int {
	//开始前清理数据
	
	idx := 1
	battleid := L.ToInt(idx)
	idx ++
	unitid := L.ToInt(idx)
	idx ++
	new_data := L.ToInt(idx)

	std.LogInfo("__BuffChangeData")
	battle := FindBattle(int64(battleid))
	unit := battle.SelectOneUnit(int64(unitid))

	for _, buff := range unit.Allbuff {
		if buff.BuffType != kTypeBuff {
			continue
		}

		buff.Data = buff.Data + buff.Data*(int32(new_data))

		std.LogInfo("__BuffChangeData  battleid",battleid,"unitid",unitid,"buff.Data",buff.Data)
	}

	//buff := unit.SelectBuff(int32(buffinstid))
	//buff.MustUpdate()

	return 0
}

////////////////////////////////////////////////////////////////////////////////////////////
//export __GetMyUnitIProperty
func __GetMyUnitIProperty(L* lua.LState) int {
	

	idx := 1
	casterId := L.ToInt64(idx)
	idx ++
	ipc := L.ToString(idx)

	player := FindPlayerByInstId(casterId)
	if player == nil {
		std.LogInfo("__GetMyUnitIProperty FindPlayerByInstId==nil",casterId)
		return 1
	}

	pd := prpc.ToId_IPropertyType(ipc)

	val := player.MyUnit.GetIProperty(int32(pd))

	L.Push(lua.LNumber(val))

	return 1
}
//export __AddMyUnitEnergy
func __AddMyUnitEnergy(L* lua.LState) int {

	

	idx := 1
	casterId := L.ToInt64(idx)
	idx ++
	val := L.ToInt(idx)

	player := FindPlayerByInstId(casterId)
	if player == nil {
		std.LogInfo("__AddMyUnitEnergy FindPlayerByInstId==nil",casterId)
		return 0
	}

	player.SetMyEnergy(int32(val),true)

	return 0
}


//export __ThrowCard
func __ThrowCard(L* lua.LState) int {  //获取要删除的卡牌

	

	idx := 1
	battleid := L.ToInt64(idx)
	idx ++
	unitid := L.ToInt64(idx)
	idx ++
	target := L.ToInt64(idx)

	battle := FindBattle(battleid)

	throwCard, entid, level := battle.SelectThrowCard(unitid)

	battle.ThrowCard(target, throwCard, entid, level)

	L.Push(lua.LNumber(throwCard))

	return 1
}
//export __Throw
func __Throw(L* lua.LState) int { //删除指定卡牌

	

	idx := 1
	battleid := L.ToInt64(idx)
	idx ++
	unitid := L.ToInt64(idx)
	idx ++
	throw := L.ToInt64(idx)

	battle := FindBattle(battleid)

	battle.Throw(unitid, throw)

	return 0
}



