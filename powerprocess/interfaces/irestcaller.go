package interfaces

type IRestCall interface {
   GetAllGroups()(map[string]interface{},error)
   GetGroupById(id string)(interface{},error)
   GetGroupByName(name string)(map[string]interface{},error)
}
