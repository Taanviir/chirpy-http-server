// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: chirpy_red.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const updateChirpyRedStatus = `-- name: UpdateChirpyRedStatus :exec
UPDATE users SET is_chirpy_red = true
WHERE id = $1
`

func (q *Queries) UpdateChirpyRedStatus(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateChirpyRedStatus, id)
	return err
}
