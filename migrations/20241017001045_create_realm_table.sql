-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS realms (
    id VARCHAR NOT NULL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS realms;
-- +goose StatementEnd
