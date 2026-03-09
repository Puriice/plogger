package main

import (
	"log"

	"github.com/puriice/golibs/pkg/db"
	"github.com/puriice/golibs/pkg/env"
	"github.com/puriice/golibs/pkg/server"
	"github.com/puriice/plogger/pkg/logger"
)

func main() {
	env.Init()

	host := env.Get("HOST", "")
	port := env.Get("PORT", "8082")
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(host, port, database)

	logger.RegisterRoute(server)

	server.Start()
}
