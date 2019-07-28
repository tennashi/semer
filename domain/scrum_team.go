package domain

// ScrumTeam represents a scrum team.
type ScrumTeam struct {
	ID          uint64
	DevTeam     []User
	ProdOwner   *User
	ScrumMaster *User
}

type ScrumTeamRepository interface {
	FindByID(teamID uint64) (*ScrumTeam, error)
}
