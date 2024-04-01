-- name: CreateFeed :one

INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many 
SELECT * FROM Feeds;

-- name: GetNextFeedsToFetch :many 
SELECT * FROM Feeds 
ORDER BY LAST_FETCHED_AT NULLS FIRST 
LIMIT $1;

-- name: MarkAsFetched :one 
UPDATE feeds 
SET 
updated_at = CURRENT_TIMESTAMP, 
last_fetched_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
