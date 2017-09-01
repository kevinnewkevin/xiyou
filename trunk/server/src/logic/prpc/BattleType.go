package prpc
// enum BattleType
const( 
  BT_MIN = 0
  BT_PVP = 1
  BT_PVE = 2
  BT_MAX = 3
)
const( 
  K_BT_MIN = "BT_MIN"
  K_BT_PVP = "BT_PVP"
  K_BT_PVE = "BT_PVE"
  K_BT_MAX = "BT_MAX"
)
func ToName_BattleType( id int )string{
  switch(id){
    case BT_MIN:
      return "BT_MIN"
    case BT_PVP:
      return "BT_PVP"
    case BT_PVE:
      return "BT_PVE"
    case BT_MAX:
      return "BT_MAX"
    default:
      return ""
  }
}
func ToId_BattleType( name string ) int {
  switch(name){
    case "BT_MIN":
      return BT_MIN
    case "BT_PVP":
      return BT_PVP
    case "BT_PVE":
      return BT_PVE
    case "BT_MAX":
      return BT_MAX
    default:
      return -1
  }
}
