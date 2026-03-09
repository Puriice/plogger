package logger

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/puriice/plogger/internal/repository"
	"github.com/puriice/plogger/internal/subscriber"
	"github.com/puriice/pproject/pkg/sdk/pproject"
)

func SubscribeToProject(context context.Context, service *pproject.ProjectService, db *pgxpool.Pool) error {
	projectRepository := repository.NewPostgresProjectRepository(db)
	subscriber := subscriber.NewProjectSubscriber(projectRepository)

	listener, err := service.NewListener("plogger.projects")

	if err != nil {
		return err
	}

	subscriber.RegisterSubscriber(*listener)

	return listener.Subscribe(context)
}
