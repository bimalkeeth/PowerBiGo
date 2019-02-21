package contracts

type Groups struct {
	ID                    string `json:"id"`
	IsReadOnly            bool   `json:"isReadOnly"`
	IsOnDedicatedCapacity bool   `json:"isOnDedicatedCapacity"`
	Name                  string `json:"name"`
}

type GroupData struct {
	Context    string   `json:"@odata.context"`
	Count      int      `json:"@odata.count"`
	GroupsData []Groups `json:"value"`
}
