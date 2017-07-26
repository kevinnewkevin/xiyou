package prpc

// enum BattlePosotion
const (
	BP_RED_1  = 0
	BP_RED_2  = 1
	BP_RED_3  = 2
	BP_RED_4  = 3
	BP_RED_5  = 4
	BP_RED_6  = 5
	BP_BLUE_1 = 6
	BP_BLUE_2 = 7
	BP_BLUE_3 = 8
	BP_BLUE_4 = 9
	BP_BLUE_5 = 10
	BP_BLUE_6 = 11
	BP_MAX    = 12
)
const (
	K_BP_RED_1  = "BP_RED_1"
	K_BP_RED_2  = "BP_RED_2"
	K_BP_RED_3  = "BP_RED_3"
	K_BP_RED_4  = "BP_RED_4"
	K_BP_RED_5  = "BP_RED_5"
	K_BP_RED_6  = "BP_RED_6"
	K_BP_BLUE_1 = "BP_BLUE_1"
	K_BP_BLUE_2 = "BP_BLUE_2"
	K_BP_BLUE_3 = "BP_BLUE_3"
	K_BP_BLUE_4 = "BP_BLUE_4"
	K_BP_BLUE_5 = "BP_BLUE_5"
	K_BP_BLUE_6 = "BP_BLUE_6"
	K_BP_MAX    = "BP_MAX"
)

func ToName_BattlePosotion(id int) string {
	switch id {
	case BP_RED_1:
		return "BP_RED_1"
	case BP_RED_2:
		return "BP_RED_2"
	case BP_RED_3:
		return "BP_RED_3"
	case BP_RED_4:
		return "BP_RED_4"
	case BP_RED_5:
		return "BP_RED_5"
	case BP_RED_6:
		return "BP_RED_6"
	case BP_BLUE_1:
		return "BP_BLUE_1"
	case BP_BLUE_2:
		return "BP_BLUE_2"
	case BP_BLUE_3:
		return "BP_BLUE_3"
	case BP_BLUE_4:
		return "BP_BLUE_4"
	case BP_BLUE_5:
		return "BP_BLUE_5"
	case BP_BLUE_6:
		return "BP_BLUE_6"
	case BP_MAX:
		return "BP_MAX"
	default:
		return ""
	}
}
func ToId_BattlePosotion(name string) int {
	switch name {
	case "BP_RED_1":
		return BP_RED_1
	case "BP_RED_2":
		return BP_RED_2
	case "BP_RED_3":
		return BP_RED_3
	case "BP_RED_4":
		return BP_RED_4
	case "BP_RED_5":
		return BP_RED_5
	case "BP_RED_6":
		return BP_RED_6
	case "BP_BLUE_1":
		return BP_BLUE_1
	case "BP_BLUE_2":
		return BP_BLUE_2
	case "BP_BLUE_3":
		return BP_BLUE_3
	case "BP_BLUE_4":
		return BP_BLUE_4
	case "BP_BLUE_5":
		return BP_BLUE_5
	case "BP_BLUE_6":
		return BP_BLUE_6
	case "BP_MAX":
		return BP_MAX
	default:
		return -1
	}
}
