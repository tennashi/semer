package api

import (
	"context"

	"github.com/tennashi/semer/proto/gen/data/v1"
	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
	"github.com/tennashi/semer/usecase"
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
	team, err := usecase.FindTeamByID(req.TeamId)
	if err != nil {
		return nil, err
	}

	prodOwner := &data.User{
		Id: team.ProdOwner.ID,
	}
	scrumMaster := &data.User{
		Id: team.ScrumMaster.ID,
	}
	devTeam := []*data.User{}
	for _, member := range team.DevTeam {
		m := &data.User{
			Id: member.ID,
		}
		devTeam = append(devTeam, m)
	}

	scrumTeam := &data.ScrumTeam{
		Id:          1,
		ProdOwner:   prodOwner,
		ScrumMaster: scrumMaster,
		DevTeam:     devTeam,
	}

	return &scrum_team.FindTeamByIDResponse{
		Team: scrumTeam,
	}, nil
}
