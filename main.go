package main

import (
	"fmt"
	"net"
	"os"

	"github.com/tennashi/semer/api"
	"github.com/tennashi/semer/proto/gen/services/scrum_team/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	server := grpc.NewServer()

	st := api.NewScrumTeamServer()
	scrum_team.RegisterScrumTeamServer(server, st)

	reflection.Register(server)
	server.Serve(lis)
}
