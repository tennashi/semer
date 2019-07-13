package api

import (
	"context"

	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
)

type scrumTeam struct {
}

func NewScrumTeamServer() scrum_team.ScrumTeamServer {
	return &scrumTeam{}
}

func (st *scrumTeam) FindTeamByID(
	ctx context.Context,
	req *scrum_team.FindTeamByIDRequest,
) (*scrum_team.FindTeamByIDResponse, error) {
	return &scrum_team.FindTeamByIDResponse{
		Res: "hello",
	}, nil
}
