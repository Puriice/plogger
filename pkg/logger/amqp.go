package logger

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/plogger/internal/constant"
	"github.com/puriice/plogger/internal/repository"
	"github.com/puriice/plogger/internal/subscriber"
	"github.com/puriice/plogger/pkg/sdk/plog"
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

func SubscribeToLogs(context context.Context, broker *messaging.RabbitBroker, db *pgxpool.Pool) error {
	listener, err := broker.NewListener(constant.QueueName, plog.LogEvents[:]...)

	if err != nil {
		return err
	}

	loggerRepository := repository.NewPostgresLoggerRepository(db)
	subscriber := subscriber.NewLogsSubscriber(loggerRepository)

	return subscriber.Subscribe(context, listener)
}
