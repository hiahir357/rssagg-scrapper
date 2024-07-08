-- +goose Up
CREATE TABLE FeedFollows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES Feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE FeedFollows;