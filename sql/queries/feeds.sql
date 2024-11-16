-- name: CreateFeed :one
INSERT INTO feeds( id, created_at, updated_at, title, url, user_id)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: GetFeeds :many
SELECT * FROM feeds;


-- name: DeleteFeed :one
DELETE FROM feeds WHERE id=$1 AND user_id=$2
RETURNING id;


-- name: GetNextFeedToFetchFrom :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;


-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at = NOW(), updated_at = NOW() WHERE id=$1
RETURNING *;