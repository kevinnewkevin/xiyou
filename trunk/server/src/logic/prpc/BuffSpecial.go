package prpc
// enum BuffSpecial
const( 
  BF_MIN = 0
  BF_JUMP = 1
  BF_CRIT = 2
  BF_UNDEAD = 3
  BF_SHELD = 4
  BF_WEAK = 5
  BF_STRONG = 6
  BF_MAX = 7
)
const( 
  K_BF_MIN = "BF_MIN"
  K_BF_JUMP = "BF_JUMP"
  K_BF_CRIT = "BF_CRIT"
  K_BF_UNDEAD = "BF_UNDEAD"
  K_BF_SHELD = "BF_SHELD"
  K_BF_WEAK = "BF_WEAK"
  K_BF_STRONG = "BF_STRONG"
  K_BF_MAX = "BF_MAX"
)
func ToName_BuffSpecial( id int )string{
  switch(id){
    case BF_MIN:
      return "BF_MIN"
    case BF_JUMP:
      return "BF_JUMP"
    case BF_CRIT:
      return "BF_CRIT"
    case BF_UNDEAD:
      return "BF_UNDEAD"
    case BF_SHELD:
      return "BF_SHELD"
    case BF_WEAK:
      return "BF_WEAK"
    case BF_STRONG:
      return "BF_STRONG"
    case BF_MAX:
      return "BF_MAX"
    default:
      return ""
  }
}
func ToId_BuffSpecial( name string ) int {
  switch(name){
    case "BF_MIN":
      return BF_MIN
    case "BF_JUMP":
      return BF_JUMP
    case "BF_CRIT":
      return BF_CRIT
    case "BF_UNDEAD":
      return BF_UNDEAD
    case "BF_SHELD":
      return BF_SHELD
    case "BF_WEAK":
      return BF_WEAK
    case "BF_STRONG":
      return BF_STRONG
    case "BF_MAX":
      return BF_MAX
    default:
      return -1
  }
}
