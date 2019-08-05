package dto

import (
	"github.com/tennashi/semer/domain"
	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
)

func ToScrumTeamPB(st *domain.ScrumTeam) *scrum_team.ScrumTeam {
	prodOwner := &scrum_team.ScrumTeam_User{
		Id: st.ProdOwner.ID,
	}
	scrumMaster := &scrum_team.ScrumTeam_User{
		Id: st.ScrumMaster.ID,
	}
	devTeam := make([]*scrum_team.ScrumTeam_User, len(st.DevTeam))
	for i, member := range st.DevTeam {
		m := &scrum_team.ScrumTeam_User{
			Id: member.ID,
		}
		devTeam[i] = m
	}

	return &scrum_team.ScrumTeam{
		Id:          st.ID,
		ProdOwner:   prodOwner,
		ScrumMaster: scrumMaster,
		DevTeam:     devTeam,
	}
}
