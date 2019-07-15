package api

import (
	"context"

	"github.com/tennashi/semer/domain"
	"github.com/tennashi/semer/proto/gen/data/v1"
	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
	"github.com/tennashi/semer/usecase"
)

type scrumTeam struct {
}

// NewScrumTeamServer returns the ScrumTeamServer interface.
func NewScrumTeamServer() scrum_team.ScrumTeamServer {
	return &scrumTeam{}
}

func transformPB(st *domain.ScrumTeam) *data.ScrumTeam {
	prodOwner := &data.User{
		Id: st.ProdOwner.ID,
	}
	scrumMaster := &data.User{
		Id: st.ScrumMaster.ID,
	}
	devTeam := make([]*data.User, len(st.DevTeam))
	for _, member := range st.DevTeam {
		m := &data.User{
			Id: member.ID,
		}
		devTeam = append(devTeam, m)
	}

	return &data.ScrumTeam{
		Id:          st.ID,
		ProdOwner:   prodOwner,
		ScrumMaster: scrumMaster,
		DevTeam:     devTeam,
	}
}

func (st *scrumTeam) FindTeamByID(
	ctx context.Context,
	req *scrum_team.FindTeamByIDRequest,
) (*scrum_team.FindTeamByIDResponse, error) {
	team, err := usecase.FindTeamByID(req.TeamId)
	if err != nil {
		return nil, err
	}

	res := transformPB(team)

	return &scrum_team.FindTeamByIDResponse{
		Team: res,
	}, nil
}
