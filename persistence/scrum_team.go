package persistence

import "github.com/tennashi/semer/domain"

type scrumTeam struct{}

func NewScrumTeam() domain.ScrumTeamRepository {
	return &scrumTeam{}
}

func (st *scrumTeam) FindByID(teamID uint64) (*domain.ScrumTeam, error) {
	team := &domain.ScrumTeam{
		ID:          teamID,
		DevTeam:     []domain.User{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
		ProdOwner:   &domain.User{ID: 1},
		ScrumMaster: &domain.User{ID: 2},
	}
	return team, nil
}
