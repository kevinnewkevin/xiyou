package prpc
// enum IPropertyType
const( 
  IPT_MIN = 0
  IPT_HP = 1
  IPT_PHYLE = 2
  IPT_TITLE = 3
  IPT_EXPERIENCE = 4
  IPT_LEVEL = 5
  IPT_COPPER = 6
  IPT_SILVER = 7
  IPT_GOLD = 8
  IPT_PROMOTE = 9
  IPT_MAX = 10
)
const( 
  K_IPT_MIN = "IPT_MIN"
  K_IPT_HP = "IPT_HP"
  K_IPT_PHYLE = "IPT_PHYLE"
  K_IPT_TITLE = "IPT_TITLE"
  K_IPT_EXPERIENCE = "IPT_EXPERIENCE"
  K_IPT_LEVEL = "IPT_LEVEL"
  K_IPT_COPPER = "IPT_COPPER"
  K_IPT_SILVER = "IPT_SILVER"
  K_IPT_GOLD = "IPT_GOLD"
  K_IPT_PROMOTE = "IPT_PROMOTE"
  K_IPT_MAX = "IPT_MAX"
)
func ToName_IPropertyType( id int )string{
  switch(id){
    case IPT_MIN:
      return "IPT_MIN"
    case IPT_HP:
      return "IPT_HP"
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
    case IPT_PROMOTE:
      return "IPT_PROMOTE"
    case IPT_MAX:
      return "IPT_MAX"
    default:
      return ""
  }
}
func ToId_IPropertyType( name string ) int {
  switch(name){
    case "IPT_MIN":
      return IPT_MIN
    case "IPT_HP":
      return IPT_HP
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
    case "IPT_PROMOTE":
      return IPT_PROMOTE
    case "IPT_MAX":
      return IPT_MAX
    default:
      return -1
  }
}
