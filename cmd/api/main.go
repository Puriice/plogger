package main

import (
	"log"

	"github.com/puriice/httplibs/pkg/db"
	"github.com/puriice/httplibs/pkg/env"
	"github.com/puriice/httplibs/pkg/server"
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
