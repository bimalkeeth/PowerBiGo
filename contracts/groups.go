package contracts

type Groups struct {
	ID string `json:"id"`
	IsReadOnly bool `json:"isReadOnly"`
	IsOnDedicatedCapacity bool `json:"isOnDedicatedCapacity"`
	Name string `json:"name"`
}