package prpc
// enum BattleState
const( 
  BS_START = 0
  BS_HANDLE = 1
  BS_SHOW = 2
  BS_END = 3
)
const( 
  K_BS_START = "BS_START"
  K_BS_HANDLE = "BS_HANDLE"
  K_BS_SHOW = "BS_SHOW"
  K_BS_END = "BS_END"
)
func ToName_BattleState( id int )string{
  switch(id){
    case BS_START:
      return "BS_START"
    case BS_HANDLE:
      return "BS_HANDLE"
    case BS_SHOW:
      return "BS_SHOW"
    case BS_END:
      return "BS_END"
    default:
      return ""
  }
}
func ToId_BattleState( name string ) int {
  switch(name){
    case "BS_START":
      return BS_START
    case "BS_HANDLE":
      return BS_HANDLE
    case "BS_SHOW":
      return BS_SHOW
    case "BS_END":
      return BS_END
    default:
      return -1
  }
}
