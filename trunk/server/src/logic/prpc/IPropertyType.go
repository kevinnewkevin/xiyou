package prpc
// enum IPropertyType
const( 
  IPT_MIN = 0
  IPT_PHYLE = 1
  IPT_TITLE = 2
  IPT_EXPERIENCE = 3
  IPT_LEVEL = 4
  IPT_COPPER = 5
  IPT_SILVER = 6
  IPT_GOLD = 7
  IPT_MAX = 8
)
const( 
  K_IPT_MIN = "IPT_MIN"
  K_IPT_PHYLE = "IPT_PHYLE"
  K_IPT_TITLE = "IPT_TITLE"
  K_IPT_EXPERIENCE = "IPT_EXPERIENCE"
  K_IPT_LEVEL = "IPT_LEVEL"
  K_IPT_COPPER = "IPT_COPPER"
  K_IPT_SILVER = "IPT_SILVER"
  K_IPT_GOLD = "IPT_GOLD"
  K_IPT_MAX = "IPT_MAX"
)
func ToName_IPropertyType( id int )string{
  switch(id){
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
func ToId_IPropertyType( name string ) int {
  switch(name){
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
