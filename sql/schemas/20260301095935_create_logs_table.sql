-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE logs.logs (
	id				UUID 		UNIQUE NOT NULL DEFAULT gen_random_uuid(),
	project_id		UUID		NOT NULL,
	message			TEXT		NOT NULL DEFAULT '',
	created_at		TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY(id),
	FOREIGN KEY(project_id) REFERENCES projects.projects(id) ON UPDATE CASCADE ON DELETE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE logs;
-- +goose StatementEnd
