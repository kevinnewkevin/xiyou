package prpc

// enum CPropertyType
const (
	CPT_MIN            = 0
	CPT_HP             = 1
	CPT_ATK            = 2
	CPT_DEF            = 3
	CPT_MAGIC_ATK      = 4
	CPT_MAGIC_DEF      = 5
	CPT_AGILE          = 6
	CPT_KILL           = 7
	CPT_CRIT           = 8
	CPT_COUNTER_ATTACK = 9
	CPT_SPUTTERING     = 10
	CPT_DOUBLE_HIT     = 11
	CPT_RECOVERY       = 12
	CPT_REFLEX         = 13
	CPT_SUCK_BLOOD     = 14
	CPT_INCANTER       = 15
	CPT_RESISTANCE     = 16
	CPT_MAX            = 17
)

func ToName_CPropertyType(id int) string {
	switch id {
	case CPT_MIN:
		return "CPT_MIN"
	case CPT_HP:
		return "CPT_HP"
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
