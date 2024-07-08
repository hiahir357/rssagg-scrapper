-- name: CreateFeed :one
INSERT INTO
    Feeds (id, created_at, updated_at, name, url, user_id)
VALUES
    ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: GetFeeds :many
SELECT * from Feeds;

-- name: GetNextFeedsToFecth :many
SELECT * FROM Feeds
ORDER BY last_fetchedat ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE Feeds
SET last_fetchedat = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;