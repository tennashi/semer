package main

import (
	"os"

	"github.com/tennashi/semer/server"
)

func main() {
	os.Exit(server.Run())
}
