package contracts

type GroupCreateRequest struct {
	GroupName string
	Members   []MembersRights
}
