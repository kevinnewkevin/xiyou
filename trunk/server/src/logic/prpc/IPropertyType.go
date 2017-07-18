package prpc

// enum IPropertyType
const (
	IPT_MIN        = 0
	IPT_PHYLE      = 1
	IPT_TITLE      = 2
	IPT_EXPERIENCE = 3
	IPT_LEVEL      = 4
	IPT_COPPER     = 5
	IPT_SILVER     = 6
	IPT_GOLD       = 7
	IPT_MAX        = 8
)

func ToName_IPropertyType(id int) string {
	switch id {
	case IPT_MIN:
		return "IPT_MIN"
	case IPT_PHYLE:
		return "IPT_PHYLE"
	case IPT_TITLE:
		return "IPT_TITLE"
	case IPT_EXPERIENCE:
		return "IPT_EXPERIENCE"
	case IPT_LEVEL:
		return "IPT_LEVEL"
	case IPT_COPPER:
		return "IPT_COPPER"
	case IPT_SILVER:
		return "IPT_SILVER"
	case IPT_GOLD:
		return "IPT_GOLD"
	case IPT_MAX:
		return "IPT_MAX"
	default:
		return ""
	}
}
func ToId_IPropertyType(name string) int {
	switch name {
	case "IPT_MIN":
		return IPT_MIN
	case "IPT_PHYLE":
		return IPT_PHYLE
	case "IPT_TITLE":
		return IPT_TITLE
	case "IPT_EXPERIENCE":
		return IPT_EXPERIENCE
	case "IPT_LEVEL":
		return IPT_LEVEL
	case "IPT_COPPER":
		return IPT_COPPER
	case "IPT_SILVER":
		return IPT_SILVER
	case "IPT_GOLD":
		return IPT_GOLD
	case "IPT_MAX":
		return IPT_MAX
	default:
		return -1
	}
}
