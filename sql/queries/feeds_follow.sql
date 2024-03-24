-- name: FollowFeed :one

INSERT INTO feeds_users 
(id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UnfollowFeed :exec

DELETE FROM feeds_users WHERE id = $1 AND user_id = $2; 


-- name: GetFeedFollowsForUser :many

SELECT * FROM feeds_users WHERE user_id = $1;
