package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/puriice/golibs/pkg/db"
	"github.com/puriice/golibs/pkg/env"
	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/golibs/pkg/server"
	"github.com/puriice/plogger/pkg/logger"
	"github.com/puriice/pproject/pkg/sdk/pproject"
)

func main() {
	env.Init()

	projectBroker, err := messaging.NewRabbitMQ(env.Get("amqp_url", "amqp://guest:guest@localhost/"), pproject.ExchangeName)

	if err != nil {
		log.Fatal(err)
	}

	defer projectBroker.Shutdown()

	projectService := pproject.NewService("", projectBroker)

	host := env.Get("HOST", "")
	port := env.Get("PORT", "8082")
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(host, port, database)

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := logger.SubscribeToProject(ctx, projectService, database); err != nil {
			log.Println(err)
		}
	}()
	go func() {
		<-sig

		log.Println("shutting down...")
		cancel()
	}()

	logger.RegisterRoute(server)

	server.Start()
}
