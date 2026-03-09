package repository

import (
	"context"

	"github.com/puriice/plogger/internal/model"
)

type LoggerRepository interface {
	GetLogByProject(context context.Context, projectId string) ([]model.Log, error)
	CreateLog(context context.Context, projectId string, msgType string, message string) error
}

type ProjectRepository interface {
	CreateProject(context context.Context, project model.Project) error
	UpdateProject(context context.Context, project model.Project) error
	DeleteProject(context context.Context, id string) error
}
