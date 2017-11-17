package prpc
// enum GuildJob
const( 
  GJ_None = 0
  GJ_People = 1
  GJ_VicePremier = 2
  GJ_Premier = 3
  GJ_Max = 4
)
const( 
  K_GJ_None = "GJ_None"
  K_GJ_People = "GJ_People"
  K_GJ_VicePremier = "GJ_VicePremier"
  K_GJ_Premier = "GJ_Premier"
  K_GJ_Max = "GJ_Max"
)
func ToName_GuildJob( id int )string{
  switch(id){
    case GJ_None:
      return "GJ_None"
    case GJ_People:
      return "GJ_People"
    case GJ_VicePremier:
      return "GJ_VicePremier"
    case GJ_Premier:
      return "GJ_Premier"
    case GJ_Max:
      return "GJ_Max"
    default:
      return ""
  }
}
func ToId_GuildJob( name string ) int {
  switch(name){
    case "GJ_None":
      return GJ_None
    case "GJ_People":
      return GJ_People
    case "GJ_VicePremier":
      return GJ_VicePremier
    case "GJ_Premier":
      return GJ_Premier
    case "GJ_Max":
      return GJ_Max
    default:
      return -1
  }
}
