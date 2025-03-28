// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: phase.sql

package sqlc

import (
	"context"
)

const createPhase = `-- name: CreatePhase :one
INSERT INTO phases (
  name, quiz_id
) VALUES (
  $1, $2
)
RETURNING id, quiz_id, name, created_at
`

type CreatePhaseParams struct {
	Name   string `json:"name"`
	QuizID int64  `json:"quiz_id"`
}

func (q *Queries) CreatePhase(ctx context.Context, arg CreatePhaseParams) (Phase, error) {
	row := q.db.QueryRow(ctx, createPhase, arg.Name, arg.QuizID)
	var i Phase
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const deletePhase = `-- name: DeletePhase :exec
DELETE FROM phases
WHERE id = $1
`

func (q *Queries) DeletePhase(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePhase, id)
	return err
}

const getPhase = `-- name: GetPhase :one
SELECT id, quiz_id, name, created_at FROM phases
WHERE id = $1 LIMIT 2
`

func (q *Queries) GetPhase(ctx context.Context, id int64) (Phase, error) {
	row := q.db.QueryRow(ctx, getPhase, id)
	var i Phase
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const listPhases = `-- name: ListPhases :many
SELECT id, quiz_id, name, created_at FROM phases
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListPhasesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPhases(ctx context.Context, arg ListPhasesParams) ([]Phase, error) {
	rows, err := q.db.Query(ctx, listPhases, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Phase
	for rows.Next() {
		var i Phase
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Name,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePhase = `-- name: UpdatePhase :one
UPDATE phases
  set name = $2
WHERE id = $1
RETURNING id, quiz_id, name, created_at
`

type UpdatePhaseParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdatePhase(ctx context.Context, arg UpdatePhaseParams) (Phase, error) {
	row := q.db.QueryRow(ctx, updatePhase, arg.ID, arg.Name)
	var i Phase
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}
