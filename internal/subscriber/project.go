package subscriber

import (
	"context"
	"log"

	"github.com/puriice/plogger/internal/model"
	"github.com/puriice/plogger/internal/repository"
	pProjectModel "github.com/puriice/pproject/pkg/model"
	"github.com/puriice/pproject/pkg/sdk/pproject"
)

type ProjectSubscriber struct {
	repo repository.ProjectRepository
}

func NewProjectSubscriber(repo repository.ProjectRepository) ProjectSubscriber {
	return ProjectSubscriber{
		repo: repo,
	}
}

func (l ProjectSubscriber) onCreate(project *pProjectModel.Project) {
	log.Printf("Project create event received. ID: %s", *project.ID)
	prjModel := model.Project{
		ID:   *project.ID,
		Name: *project.Name,
	}

	err := l.repo.CreateProject(context.Background(), prjModel)

	if err != nil {
		log.Println(err)
	}
}

func (l ProjectSubscriber) onUpdate(project *pProjectModel.Project) {
	log.Printf("Project update event received. ID: %s", *project.ID)
	prjModel := model.Project{
		ID:   *project.ID,
		Name: *project.Name,
	}

	err := l.repo.UpdateProject(context.Background(), prjModel)

	if err != nil {
		log.Println(err)
	}
}

func (l ProjectSubscriber) onDelete(id string) {
	log.Printf("Project delete event received. ID: %s", id)

	err := l.repo.DeleteProject(context.Background(), id)

	if err != nil {
		log.Println(err)
	}
}

func (l ProjectSubscriber) RegisterSubscriber(listener pproject.ProjectListener) {
	listener.OnCreate(l.onCreate)
	listener.OnUpdate(l.onUpdate)
	listener.OnDelete(l.onDelete)
}
