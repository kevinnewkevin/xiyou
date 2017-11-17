package prpc
// enum ModifyListFlag
const( 
  MLF_Add = 0
  MLF_Delete = 1
  MLF_ChangeOnline = 2
  MLF_ChangeOffline = 3
  MLF_ChangePosition = 4
  MLF_ChangeContribution = 5
  MLF_ChangeLevel = 6
  MLF_ChangeTianTiVal = 7
)
const( 
  K_MLF_Add = "MLF_Add"
  K_MLF_Delete = "MLF_Delete"
  K_MLF_ChangeOnline = "MLF_ChangeOnline"
  K_MLF_ChangeOffline = "MLF_ChangeOffline"
  K_MLF_ChangePosition = "MLF_ChangePosition"
  K_MLF_ChangeContribution = "MLF_ChangeContribution"
  K_MLF_ChangeLevel = "MLF_ChangeLevel"
  K_MLF_ChangeTianTiVal = "MLF_ChangeTianTiVal"
)
func ToName_ModifyListFlag( id int )string{
  switch(id){
    case MLF_Add:
      return "MLF_Add"
    case MLF_Delete:
      return "MLF_Delete"
    case MLF_ChangeOnline:
      return "MLF_ChangeOnline"
    case MLF_ChangeOffline:
      return "MLF_ChangeOffline"
    case MLF_ChangePosition:
      return "MLF_ChangePosition"
    case MLF_ChangeContribution:
      return "MLF_ChangeContribution"
    case MLF_ChangeLevel:
      return "MLF_ChangeLevel"
    case MLF_ChangeTianTiVal:
      return "MLF_ChangeTianTiVal"
    default:
      return ""
  }
}
func ToId_ModifyListFlag( name string ) int {
  switch(name){
    case "MLF_Add":
      return MLF_Add
    case "MLF_Delete":
      return MLF_Delete
    case "MLF_ChangeOnline":
      return MLF_ChangeOnline
    case "MLF_ChangeOffline":
      return MLF_ChangeOffline
    case "MLF_ChangePosition":
      return MLF_ChangePosition
    case "MLF_ChangeContribution":
      return MLF_ChangeContribution
    case "MLF_ChangeLevel":
      return MLF_ChangeLevel
    case "MLF_ChangeTianTiVal":
      return MLF_ChangeTianTiVal
    default:
      return -1
  }
}
