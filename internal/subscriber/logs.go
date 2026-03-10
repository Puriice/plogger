package subscriber

import (
	"context"
	"encoding/json"

	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/plogger/internal/repository"
	"github.com/puriice/plogger/pkg/sdk/plog"
)

type LogsSubscriber struct {
	repo repository.LoggerRepository
}

func NewLogsSubscriber(repo repository.LoggerRepository) LogsSubscriber {
	return LogsSubscriber{
		repo: repo,
	}
}

func (s LogsSubscriber) onCreate(log plog.Log) error {
	return s.repo.CreateLog(context.Background(), *log.ProjectId, *log.Type, *log.Message)
}

func (s LogsSubscriber) Subscribe(context context.Context, listener *messaging.RabbitListener) error {
	return listener.Subscribe(context, func(b []byte) error {
		var log plog.LogEvent

		err := json.Unmarshal(b, &log)

		if err != nil {
			return err
		}

		if !log.IsValid() {
			return plog.ErrUnprocessableBody
		}

		switch log.EventType {
		case plog.LogCreate:
			return s.onCreate(log.Log)
		default:
			return plog.ErrUnknownEvent
		}
	})
}
