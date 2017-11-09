package prpc

// enum BuffSpecial
const (
	BF_MIN        = 0
	BF_JUMP       = 1
	BF_CRIT       = 2
	BF_UNDEAD     = 3
	BF_SHELD      = 4
	BF_SHELDNUM   = 5
	BF_WEAK       = 6
	BF_STRONG     = 7
	BF_THORNS     = 8
	BF_UNCURE     = 9
	BF_AOE        = 10
	BF_COMBO      = 11
	BF_ATKCURE    = 12
	BF_LOCK       = 13
	BF_UNDAMAGE   = 14
	BF_FRIENDLOCK = 15
	BF_MAX        = 16
)
const (
	K_BF_MIN        = "BF_MIN"
	K_BF_JUMP       = "BF_JUMP"
	K_BF_CRIT       = "BF_CRIT"
	K_BF_UNDEAD     = "BF_UNDEAD"
	K_BF_SHELD      = "BF_SHELD"
	K_BF_SHELDNUM   = "BF_SHELDNUM"
	K_BF_WEAK       = "BF_WEAK"
	K_BF_STRONG     = "BF_STRONG"
	K_BF_THORNS     = "BF_THORNS"
	K_BF_UNCURE     = "BF_UNCURE"
	K_BF_AOE        = "BF_AOE"
	K_BF_COMBO      = "BF_COMBO"
	K_BF_ATKCURE    = "BF_ATKCURE"
	K_BF_LOCK       = "BF_LOCK"
	K_BF_UNDAMAGE   = "BF_UNDAMAGE"
	K_BF_FRIENDLOCK = "BF_FRIENDLOCK"
	K_BF_MAX        = "BF_MAX"
)

func ToName_BuffSpecial(id int) string {
	switch id {
	case BF_MIN:
		return "BF_MIN"
	case BF_JUMP:
		return "BF_JUMP"
	case BF_CRIT:
		return "BF_CRIT"
	case BF_UNDEAD:
		return "BF_UNDEAD"
	case BF_SHELD:
		return "BF_SHELD"
	case BF_SHELDNUM:
		return "BF_SHELDNUM"
	case BF_WEAK:
		return "BF_WEAK"
	case BF_STRONG:
		return "BF_STRONG"
	case BF_THORNS:
		return "BF_THORNS"
	case BF_UNCURE:
		return "BF_UNCURE"
	case BF_AOE:
		return "BF_AOE"
	case BF_COMBO:
		return "BF_COMBO"
	case BF_ATKCURE:
		return "BF_ATKCURE"
	case BF_LOCK:
		return "BF_LOCK"
	case BF_UNDAMAGE:
		return "BF_UNDAMAGE"
	case BF_FRIENDLOCK:
		return "BF_FRIENDLOCK"
	case BF_MAX:
		return "BF_MAX"
	default:
		return ""
	}
}
func ToId_BuffSpecial(name string) int {
	switch name {
	case "BF_MIN":
		return BF_MIN
	case "BF_JUMP":
		return BF_JUMP
	case "BF_CRIT":
		return BF_CRIT
	case "BF_UNDEAD":
		return BF_UNDEAD
	case "BF_SHELD":
		return BF_SHELD
	case "BF_SHELDNUM":
		return BF_SHELDNUM
	case "BF_WEAK":
		return BF_WEAK
	case "BF_STRONG":
		return BF_STRONG
	case "BF_THORNS":
		return BF_THORNS
	case "BF_UNCURE":
		return BF_UNCURE
	case "BF_AOE":
		return BF_AOE
	case "BF_COMBO":
		return BF_COMBO
	case "BF_ATKCURE":
		return BF_ATKCURE
	case "BF_LOCK":
		return BF_LOCK
	case "BF_UNDAMAGE":
		return BF_UNDAMAGE
	case "BF_FRIENDLOCK":
		return BF_FRIENDLOCK
	case "BF_MAX":
		return BF_MAX
	default:
		return -1
	}
}
