package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/puriice/httplibs/pkg/pgutils"
	"github.com/puriice/plogger/internal/model"
)

type LoggerRepository struct {
	db *pgxpool.Pool
}

func NewLoggerRepository(db *pgxpool.Pool) *LoggerRepository {
	return &LoggerRepository{
		db: db,
	}
}

func (r *LoggerRepository) GetLogByProject(context context.Context, projectId string) ([]model.Log, error) {
	rows, err := r.db.Query(context, "SELECT id, type, message, created_at FROM logs.logs WHERE project_id = $1 ORDER BY created_at", projectId)

	if err != nil {
		return nil, err
	}

	logs := make([]model.Log, 0)

	for rows.Next() {
		var log model.Log

		rows.Scan(&log.Identifier, &log.Type, &log.Message, &log.CreatedAt)

		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *LoggerRepository) CreateLog(context context.Context, projectId string, msgType string, message string) error {
	cmdTag, err := r.db.Exec(context, "INSERT INTO logs.logs (project_id, type, message) VALUES ($1, $2, $3)", projectId, msgType, message)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() != 1 {
		return pgutils.ErrNoRowsAffected
	}

	return nil
}
