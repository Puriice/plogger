-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE projects (
	id			UUID		NOT NULL,
	name		TEXT		NOT NULL,
	is_deleted	BOOLEAN		NOT NULL DEFAULT FALSE,
	PRIMARY KEY(id)
);
ALTER TABLE logs 
ADD CONSTRAINT project_id_fk
FOREIGN KEY (project_id) REFERENCES projects(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE logs
DROP CONSTRAINT project_id_fk;

DROP TABLE projects;
-- +goose StatementEnd
