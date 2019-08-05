package usecase

import "github.com/tennashi/semer/domain"

type ScrumTeam struct {
	stRepo domain.ScrumTeamRepository
}

func NewScrumTeam(stRepo domain.ScrumTeamRepository) *ScrumTeam {
	return &ScrumTeam{
		stRepo: stRepo,
	}
}

// FindTeamByID returns the scrum team.
func (st *ScrumTeam) FindTeamByID(teamID uint64) (*domain.ScrumTeam, error) {
	return st.stRepo.FindByID(teamID)
}
