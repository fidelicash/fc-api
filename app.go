package main

import (
	"github.com/fidelicash/fc-api/server"
	"github.com/fidelicash/fc-api/user"
)

func main() {
	user.Migrate()
	server.Listen()
}
