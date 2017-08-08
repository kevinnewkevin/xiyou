package prpc

// enum EquipmentSlot
const (
	ES_MIN      = 0
	ES_HEAD     = 1
	ES_BODY     = 2
	ES_FEET     = 3
	ES_ORNAMENT = 4
	ES_WEAPON   = 5
	ES_MAX      = 6
)
const (
	K_ES_MIN      = "ES_MIN"
	K_ES_HEAD     = "ES_HEAD"
	K_ES_BODY     = "ES_BODY"
	K_ES_FEET     = "ES_FEET"
	K_ES_ORNAMENT = "ES_ORNAMENT"
	K_ES_WEAPON   = "ES_WEAPON"
	K_ES_MAX      = "ES_MAX"
)

func ToName_EquipmentSlot(id int) string {
	switch id {
	case ES_MIN:
		return "ES_MIN"
	case ES_HEAD:
		return "ES_HEAD"
	case ES_BODY:
		return "ES_BODY"
	case ES_FEET:
		return "ES_FEET"
	case ES_ORNAMENT:
		return "ES_ORNAMENT"
	case ES_WEAPON:
		return "ES_WEAPON"
	case ES_MAX:
		return "ES_MAX"
	default:
		return ""
	}
}
func ToId_EquipmentSlot(name string) int {
	switch name {
	case "ES_MIN":
		return ES_MIN
	case "ES_HEAD":
		return ES_HEAD
	case "ES_BODY":
		return ES_BODY
	case "ES_FEET":
		return ES_FEET
	case "ES_ORNAMENT":
		return ES_ORNAMENT
	case "ES_WEAPON":
		return ES_WEAPON
	case "ES_MAX":
		return ES_MAX
	default:
		return -1
	}
}
