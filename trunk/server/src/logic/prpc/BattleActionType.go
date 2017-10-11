package prpc
// enum BattleActionType
const( 
  BAT_MIN = 0
  BAT_CRIT = 1
  BAT_SUCK = 2
  BAT_RECOVERY = 3
  BAT_ADD_STATE = 4
  BAT_DEL_STATE = 5
  BAT_MAX = 6
)
const( 
  K_BAT_MIN = "BAT_MIN"
  K_BAT_CRIT = "BAT_CRIT"
  K_BAT_SUCK = "BAT_SUCK"
  K_BAT_RECOVERY = "BAT_RECOVERY"
  K_BAT_ADD_STATE = "BAT_ADD_STATE"
  K_BAT_DEL_STATE = "BAT_DEL_STATE"
  K_BAT_MAX = "BAT_MAX"
)
func ToName_BattleActionType( id int )string{
  switch(id){
    case BAT_MIN:
      return "BAT_MIN"
    case BAT_CRIT:
      return "BAT_CRIT"
    case BAT_SUCK:
      return "BAT_SUCK"
    case BAT_RECOVERY:
      return "BAT_RECOVERY"
    case BAT_ADD_STATE:
      return "BAT_ADD_STATE"
    case BAT_DEL_STATE:
      return "BAT_DEL_STATE"
    case BAT_MAX:
      return "BAT_MAX"
    default:
      return ""
  }
}
func ToId_BattleActionType( name string ) int {
  switch(name){
    case "BAT_MIN":
      return BAT_MIN
    case "BAT_CRIT":
      return BAT_CRIT
    case "BAT_SUCK":
      return BAT_SUCK
    case "BAT_RECOVERY":
      return BAT_RECOVERY
    case "BAT_ADD_STATE":
      return BAT_ADD_STATE
    case "BAT_DEL_STATE":
      return BAT_DEL_STATE
    case "BAT_MAX":
      return BAT_MAX
    default:
      return -1
  }
}
