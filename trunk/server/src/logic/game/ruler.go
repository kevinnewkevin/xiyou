package game

import (
	"logic/prpc"
	"math/rand"
)

//游戏规则文件

const (
	kAtkValue      = 240
	kDefvalue      = 240
	kMagicAtkValue = 240
	kMagicDefvalue = 240
)

var kDefenseRandom = []float32{0, 0.1, 0.2, 0.3, 0.4, 0.5}

//计算物理攻击力
func CalcAtk(unit *GameUnit) float32 {
	currAtk := unit.GetCProperty(prpc.CPT_ATK)
	if currAtk < 240 {
		return currAtk
	} else {
		return kAtkValue + 3*(currAtk-kAtkValue)*0.1
	}
}

//计算物理防御力
func CalcDef(unit *GameUnit) float32 {
	currDef := unit.GetCProperty(prpc.CPT_DEF)
	if currDef < 240 {
		return currDef
	} else {
		return kDefvalue + 3*(currDef-kDefvalue)*0.1
	}
}

//计算法术攻击力
func CalcMagicAtk(unit *GameUnit) float32 {
	currAtk := unit.GetCProperty(prpc.CPT_MAGIC_ATK)
	if currAtk < 240 {
		return currAtk
	} else {
		return kMagicAtkValue + 3*(currAtk-kMagicAtkValue)*0.1
	}
}

//计算法术防御力
func CalcMagicDef(unit *GameUnit) float32 {
	currDef := unit.GetCProperty(prpc.CPT_MAGIC_DEF)
	if currDef < 240 {
		return currDef
	} else {
		return kMagicDefvalue + 3*(currDef-kMagicDefvalue)*0.1
	}
}

//计算攻击数值
func CalcDamage(attacker, victim *GameUnit) float32 {
	currAtk := CalcAtk(attacker)
	currDef := CalcDef(victim)

	damage := (currAtk * currAtk) * 3 / (currAtk + currDef*3)

	return damage
}

//计算法术攻击数值
func CalcMagicDamage(attacker, victim *GameUnit) float32 {
	currAtk := CalcMagicAtk(attacker)
	currDef := CalcMagicDef(victim)

	damage := (currAtk * currAtk) * 3 / (currAtk + currDef*3)

	return damage
}

//计算防御状态受击数量
func CalcDefenseDamage(attacker, victim *GameUnit) float32 {
	damage := CalcDamage(attacker, victim)
	index := rand.Intn(len(kDefenseRandom))
	damage = damage * kDefenseRandom[index]
	if damage < 1 {
		damage = 1
	}
	return damage
}

//计算防御状态受击数量
func CalcDefenseMagicDamage(attacker, victim *GameUnit) float32 {
	damage := CalcMagicDamage(attacker, victim)
	index := rand.Intn(len(kDefenseRandom))
	damage = damage * kDefenseRandom[index]
	if damage < 1 {
		damage = 1
	}
	return damage
}
