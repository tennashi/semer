package domain

type ScrumTeam struct {
	ID          uint64
	DevTeam     []User
	ProdOwner   *User
	ScrumMaster *User
}
