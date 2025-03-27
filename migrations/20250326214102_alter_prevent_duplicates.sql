-- +goose Up
-- +goose StatementBegin
ALTER TABLE songs ADD CONSTRAINT unique_song UNIQUE(group_name, song);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP CONSTRAINT unique_song;
-- +goose StatementEnd
