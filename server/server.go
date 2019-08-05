package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tennashi/semer/api"
	"github.com/tennashi/semer/persistence"
	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run initialize and execute the gRPC server.
func Run() int {
	log.Println("server starting")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	server := grpc.NewServer()

	stRepo := persistence.NewScrumTeam()
	st := api.NewScrumTeamServer(stRepo)
	scrum_team.RegisterScrumTeamManagerServer(server, st)

	reflection.Register(server)

	log.Println("server started")
	server.Serve(lis)

	return 0
}
