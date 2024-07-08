-- name: CreatePost :one
INSERT INTO
    Posts (
        id, 
        created_at, 
        updated_at, 
        title, 
        description,
        published_at,
        url, 
        feed_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING *;

-- name: GetPostsForUser :many
SELECT Posts.* FROM Posts
JOIN FeedFollows ON Posts.feed_id = FeedFollows.feed_id
WHERE FeedFollows.user_id = $1
ORDER BY Posts.published_at DESC
LIMIT $2;