package prpc

// enum CPropertyType
const (
	CPT_MIN            = 0
	CPT_HP             = 1
	CPT_CHP            = 2
	CPT_ATK            = 3
	CPT_DEF            = 4
	CPT_MAGIC_ATK      = 5
	CPT_MAGIC_DEF      = 6
	CPT_AGILE          = 7
	CPT_KILL           = 8
	CPT_CRIT           = 9
	CPT_COUNTER_ATTACK = 10
	CPT_SPUTTERING     = 11
	CPT_DOUBLE_HIT     = 12
	CPT_RECOVERY       = 13
	CPT_REFLEX         = 14
	CPT_SUCK_BLOOD     = 15
	CPT_INCANTER       = 16
	CPT_RESISTANCE     = 17
	CPT_MAX            = 18
)
const (
	K_CPT_MIN            = "CPT_MIN"
	K_CPT_HP             = "CPT_HP"
	K_CPT_CHP            = "CPT_CHP"
	K_CPT_ATK            = "CPT_ATK"
	K_CPT_DEF            = "CPT_DEF"
	K_CPT_MAGIC_ATK      = "CPT_MAGIC_ATK"
	K_CPT_MAGIC_DEF      = "CPT_MAGIC_DEF"
	K_CPT_AGILE          = "CPT_AGILE"
	K_CPT_KILL           = "CPT_KILL"
	K_CPT_CRIT           = "CPT_CRIT"
	K_CPT_COUNTER_ATTACK = "CPT_COUNTER_ATTACK"
	K_CPT_SPUTTERING     = "CPT_SPUTTERING"
	K_CPT_DOUBLE_HIT     = "CPT_DOUBLE_HIT"
	K_CPT_RECOVERY       = "CPT_RECOVERY"
	K_CPT_REFLEX         = "CPT_REFLEX"
	K_CPT_SUCK_BLOOD     = "CPT_SUCK_BLOOD"
	K_CPT_INCANTER       = "CPT_INCANTER"
	K_CPT_RESISTANCE     = "CPT_RESISTANCE"
	K_CPT_MAX            = "CPT_MAX"
)

func ToName_CPropertyType(id int) string {
	switch id {
	case CPT_MIN:
		return "CPT_MIN"
	case CPT_HP:
		return "CPT_HP"
	case CPT_CHP:
		return "CPT_CHP"
	case CPT_ATK:
		return "CPT_ATK"
	case CPT_DEF:
		return "CPT_DEF"
	case CPT_MAGIC_ATK:
		return "CPT_MAGIC_ATK"
	case CPT_MAGIC_DEF:
		return "CPT_MAGIC_DEF"
	case CPT_AGILE:
		return "CPT_AGILE"
	case CPT_KILL:
		return "CPT_KILL"
	case CPT_CRIT:
		return "CPT_CRIT"
	case CPT_COUNTER_ATTACK:
		return "CPT_COUNTER_ATTACK"
	case CPT_SPUTTERING:
		return "CPT_SPUTTERING"
	case CPT_DOUBLE_HIT:
		return "CPT_DOUBLE_HIT"
	case CPT_RECOVERY:
		return "CPT_RECOVERY"
	case CPT_REFLEX:
		return "CPT_REFLEX"
	case CPT_SUCK_BLOOD:
		return "CPT_SUCK_BLOOD"
	case CPT_INCANTER:
		return "CPT_INCANTER"
	case CPT_RESISTANCE:
		return "CPT_RESISTANCE"
	case CPT_MAX:
		return "CPT_MAX"
	default:
		return ""
	}
}
func ToId_CPropertyType(name string) int {
	switch name {
	case "CPT_MIN":
		return CPT_MIN
	case "CPT_HP":
		return CPT_HP
	case "CPT_CHP":
		return CPT_CHP
	case "CPT_ATK":
		return CPT_ATK
	case "CPT_DEF":
		return CPT_DEF
	case "CPT_MAGIC_ATK":
		return CPT_MAGIC_ATK
	case "CPT_MAGIC_DEF":
		return CPT_MAGIC_DEF
	case "CPT_AGILE":
		return CPT_AGILE
	case "CPT_KILL":
		return CPT_KILL
	case "CPT_CRIT":
		return CPT_CRIT
	case "CPT_COUNTER_ATTACK":
		return CPT_COUNTER_ATTACK
	case "CPT_SPUTTERING":
		return CPT_SPUTTERING
	case "CPT_DOUBLE_HIT":
		return CPT_DOUBLE_HIT
	case "CPT_RECOVERY":
		return CPT_RECOVERY
	case "CPT_REFLEX":
		return CPT_REFLEX
	case "CPT_SUCK_BLOOD":
		return CPT_SUCK_BLOOD
	case "CPT_INCANTER":
		return CPT_INCANTER
	case "CPT_RESISTANCE":
		return CPT_RESISTANCE
	case "CPT_MAX":
		return CPT_MAX
	default:
		return -1
	}
}
