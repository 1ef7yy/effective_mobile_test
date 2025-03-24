-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS songs(
    id bigint GENERATED ALWAYS AS IDENTITY,
    group_name TEXT NOT NULL,
    song TEXT NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE songs;
-- +goose StatementEnd
