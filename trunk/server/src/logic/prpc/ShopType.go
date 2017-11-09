package prpc

// enum ShopType
const (
	SHT_Min         = 0
	SHT_Card        = 1
	SHT_BuyCopper   = 2
	SHT_BlackMarket = 3
	SHT_Max         = 4
)
const (
	K_SHT_Min         = "SHT_Min"
	K_SHT_Card        = "SHT_Card"
	K_SHT_BuyCopper   = "SHT_BuyCopper"
	K_SHT_BlackMarket = "SHT_BlackMarket"
	K_SHT_Max         = "SHT_Max"
)

func ToName_ShopType(id int) string {
	switch id {
	case SHT_Min:
		return "SHT_Min"
	case SHT_Card:
		return "SHT_Card"
	case SHT_BuyCopper:
		return "SHT_BuyCopper"
	case SHT_BlackMarket:
		return "SHT_BlackMarket"
	case SHT_Max:
		return "SHT_Max"
	default:
		return ""
	}
}
func ToId_ShopType(name string) int {
	switch name {
	case "SHT_Min":
		return SHT_Min
	case "SHT_Card":
		return SHT_Card
	case "SHT_BuyCopper":
		return SHT_BuyCopper
	case "SHT_BlackMarket":
		return SHT_BlackMarket
	case "SHT_Max":
		return SHT_Max
	default:
		return -1
	}
}
