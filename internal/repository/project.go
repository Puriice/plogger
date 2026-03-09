package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/puriice/golibs/pkg/pgutils"
	"github.com/puriice/plogger/internal/model"
)

type PostgresProjectRepository struct {
	db *pgxpool.Pool
}

func NewPostgresProjectRepository(db *pgxpool.Pool) PostgresProjectRepository {
	return PostgresProjectRepository{
		db: db,
	}
}

func (r PostgresProjectRepository) CreateProject(context context.Context, project model.Project) error {
	_, err := r.db.Exec(context, "INSERT INTO projects (id, name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", project.ID, project.Name)

	return err
}

func (r PostgresProjectRepository) UpdateProject(context context.Context, project model.Project) error {
	cmdTag, err := r.db.Exec(context, "UPDATE projects SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", project.Name, project.ID)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return pgutils.ErrNoRowsAffected
	}

	return nil
}

func (r PostgresProjectRepository) DeleteProject(context context.Context, id string) error {
	cmdTag, err := r.db.Exec(context, "UPDATE projects SET is_deleted = true, updated_at = CURRENT_TIMESTAMP WHERE id = $1", id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return pgutils.ErrNoRowsAffected
	}

	return nil
}
