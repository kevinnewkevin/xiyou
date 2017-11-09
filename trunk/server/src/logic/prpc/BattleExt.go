package prpc
// enum BattleExt
const( 
  BE_Crit = 0
  BE_Miss = 1
  BE_3 = 2
  BE_4 = 3
  BE_5 = 4
  BE_6 = 5
  BE_7 = 6
  BE_MAX = 7
)
const( 
  K_BE_Crit = "BE_Crit"
  K_BE_Miss = "BE_Miss"
  K_BE_3 = "BE_3"
  K_BE_4 = "BE_4"
  K_BE_5 = "BE_5"
  K_BE_6 = "BE_6"
  K_BE_7 = "BE_7"
  K_BE_MAX = "BE_MAX"
)
func ToName_BattleExt( id int )string{
  switch(id){
    case BE_Crit:
      return "BE_Crit"
    case BE_Miss:
      return "BE_Miss"
    case BE_3:
      return "BE_3"
    case BE_4:
      return "BE_4"
    case BE_5:
      return "BE_5"
    case BE_6:
      return "BE_6"
    case BE_7:
      return "BE_7"
    case BE_MAX:
      return "BE_MAX"
    default:
      return ""
  }
}
func ToId_BattleExt( name string ) int {
  switch(name){
    case "BE_Crit":
      return BE_Crit
    case "BE_Miss":
      return BE_Miss
    case "BE_3":
      return BE_3
    case "BE_4":
      return BE_4
    case "BE_5":
      return BE_5
    case "BE_6":
      return BE_6
    case "BE_7":
      return BE_7
    case "BE_MAX":
      return BE_MAX
    default:
      return -1
  }
}
