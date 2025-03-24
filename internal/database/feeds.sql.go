// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6) RETURNING id, created_at, updated_at, name, url, user_id
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}

const getAllFeedsWithUserNames = `-- name: GetAllFeedsWithUserNames :many
SELECT f.id, f.created_at, f.updated_at, f.name, f.url, f.user_id, u.name as user_name
FROM feeds f
inner join users u on f.user_id = u.id
`

type GetAllFeedsWithUserNamesRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
	UserName  string
}

func (q *Queries) GetAllFeedsWithUserNames(ctx context.Context) ([]GetAllFeedsWithUserNamesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllFeedsWithUserNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllFeedsWithUserNamesRow
	for rows.Next() {
		var i GetAllFeedsWithUserNamesRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedByName = `-- name: GetFeedByName :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE name = $1
`

func (q *Queries) GetFeedByName(ctx context.Context, name string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByName, name)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}
