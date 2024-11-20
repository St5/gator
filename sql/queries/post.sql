-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostByUrl :one
SELECT * FROM posts WHERE url = $1;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts INNER JOIN feeds ON posts.feed_id = feeds.id 
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1 LIMIT $2; 
