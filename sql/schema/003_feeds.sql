-- +goose Up
CREATE TABLE Feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE Feeds;