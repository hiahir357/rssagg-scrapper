-- +goose Up
ALTER TABLE
    Feeds
ADD
    COLUMN last_fetchedat TIMESTAMP;

-- +goose Down
ALTER TABLE Feeds DROP COLUMN last_fetchedat;