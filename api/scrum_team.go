package api

import (
	"context"

	"github.com/tennashi/semer/api/dto"
	"github.com/tennashi/semer/domain"
	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
	"github.com/tennashi/semer/usecase"
)

type ScrumTeam struct {
	stUC *usecase.ScrumTeam
}

// NewScrumTeamServer returns the ScrumTeamServer interface.
func NewScrumTeamServer(repo domain.ScrumTeamRepository) *ScrumTeam {
	uc := usecase.NewScrumTeam(repo)
	return &ScrumTeam{stUC: uc}
}

func (st *ScrumTeam) FindTeamByID(
	ctx context.Context,
	req *scrum_team.FindTeamByIDRequest,
) (*scrum_team.FindTeamByIDResponse, error) {
	team, err := st.stUC.FindTeamByID(req.TeamId)
	if err != nil {
		return nil, err
	}

	res := dto.ToScrumTeamPB(team)

	return &scrum_team.FindTeamByIDResponse{
		Team: res,
	}, nil
}
