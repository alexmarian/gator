-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6) RETURNING *;

-- name: GetFeedByName :one
SELECT *
FROM feeds
WHERE name = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetAllFeedsWithUserNames :many
SELECT f.id, f.created_at, f.updated_at, f.name, f.url, f.user_id, u.name as user_name
FROM feeds f
    inner join users u on f.user_id = u.id;



