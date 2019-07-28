package usecase

import "github.com/tennashi/semer/domain"

type scrumTeam struct {
	scrumTeamRepo domain.ScrumTeamRepository
}

// FindTeamByID returns the scrum team.
func (st *scrumTeam) FindTeamByID(teamID uint64) (*domain.ScrumTeam, error) {
	return st.scrumTeamRepo.FindByID(teamID)
}
