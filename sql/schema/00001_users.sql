-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL PRIMARY KEY, 
    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'), 
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW() AT TIME ZONE 'UTC'),
    name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down 
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd