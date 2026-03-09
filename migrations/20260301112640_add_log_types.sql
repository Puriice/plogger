-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE log_types (
	type	TEXT	UNIQUE NOT NULL,
	PRIMARY KEY(type)
);

INSERT INTO log_types VALUES ('FATAL'), ('ERROR'), ('WARN'), ('INFO');

ALTER TABLE logs
ADD COLUMN type TEXT NOT NULL DEFAULT 'INFO';

ALTER TABLE logs
ADD CONSTRAINT type_fk
FOREIGN KEY (type) REFERENCES log_types(type) ON UPDATE CASCADE ON DELETE SET DEFAULT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE logs 
DROP CONSTRAINT type_fk;

ALTER TABLE logs
DROP COLUMN type;

DROP TABLE log_types;
-- +goose StatementEnd
