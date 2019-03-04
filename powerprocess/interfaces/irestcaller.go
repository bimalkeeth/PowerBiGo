package interfaces

import con "PowerBiGo/contracts"

type IRestCall interface {
	GetAllGroups(credential *con.Credential) (*[]con.Groups, error)
	GetGroupById(credential *con.Credential, id string) (*con.Groups, error)
	GetGroupByName(credential *con.Credential, name string) (*[]con.Groups, error)
	CreateGroup(credential *con.Credential, request con.GroupCreateRequest) (bool, error)
}
