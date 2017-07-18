package prpc

// enum CampType
const (
	CT_RED  = 0
	CT_BLUE = 1
	CT_MAX  = 2
)

func ToName_CampType(id int) string {
	switch id {
	case CT_RED:
		return "CT_RED"
	case CT_BLUE:
		return "CT_BLUE"
	case CT_MAX:
		return "CT_MAX"
	default:
		return ""
	}
}
func ToId_CampType(name string) int {
	switch name {
	case "CT_RED":
		return CT_RED
	case "CT_BLUE":
		return CT_BLUE
	case "CT_MAX":
		return CT_MAX
	default:
		return -1
	}
}
