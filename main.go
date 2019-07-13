package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type scrumTeam struct {
}

func (st *scrumTeam) FindTeamByID(
	ctx context.Context,
	req *scrum_team.FindTeamByIDRequest,
) (*scrum_team.FindTeamByIDResponse, error) {
	return &scrum_team.FindTeamByIDResponse{
		Res: "hello",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	server := grpc.NewServer()

	scrum_team.RegisterScrumTeamServer(server, new(scrumTeam))

	reflection.Register(server)
	server.Serve(lis)
}
