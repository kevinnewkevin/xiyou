package prpc

// enum ErrorNo
const (
	EN_MIN                        = 0
	EN_CREATE_PLAYER_SAME_NAME    = 1
	EN_CREATE_PLAYER_ILLEGAL_NAME = 2
	EN_CREATE_PLAYER_SECOND_TIME  = 3
	EN_MAX                        = 4
)

func ToName_ErrorNo(id int) string {
	switch id {
	case EN_MIN:
		return "EN_MIN"
	case EN_CREATE_PLAYER_SAME_NAME:
		return "EN_CREATE_PLAYER_SAME_NAME"
	case EN_CREATE_PLAYER_ILLEGAL_NAME:
		return "EN_CREATE_PLAYER_ILLEGAL_NAME"
	case EN_CREATE_PLAYER_SECOND_TIME:
		return "EN_CREATE_PLAYER_SECOND_TIME"
	case EN_MAX:
		return "EN_MAX"
	default:
		return ""
	}
}
func ToId_ErrorNo(name string) int {
	switch name {
	case "EN_MIN":
		return EN_MIN
	case "EN_CREATE_PLAYER_SAME_NAME":
		return EN_CREATE_PLAYER_SAME_NAME
	case "EN_CREATE_PLAYER_ILLEGAL_NAME":
		return EN_CREATE_PLAYER_ILLEGAL_NAME
	case "EN_CREATE_PLAYER_SECOND_TIME":
		return EN_CREATE_PLAYER_SECOND_TIME
	case "EN_MAX":
		return EN_MAX
	default:
		return -1
	}
}
