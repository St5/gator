-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE name = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeeds :many
SELECT f.*, u.name FROM feeds as f LEFT JOIN users as u ON f.user_id = u.id;

-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id,feed_id)
    VALUES (
         $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id 
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many

SELECT f.*, u.name 
FROM feeds as f
INNER JOIN feed_follows as ff ON f.id = ff.feed_id
INNER JOIN users as u ON ff.user_id = u.id
WHERE ff.user_id = $1;