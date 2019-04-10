package main

import (
	"time"

	"github.com/pattack/fcjanstun/team"
	"github.com/pattack/rcss"
)

func main() {
	if srv, err := rcss.NewServer("127.0.0.1:6000"); err != nil {
		panic(err)
	} else if err := srv.Join(team.NewTeam("Janstun")); err != nil {
		panic(err)
	}

	// TODO: Remove this. Wait for stop signal and do graceful shutdown instead
	<-time.After(time.Hour)
}
