-- name: GetPhase :one
SELECT * FROM phases
WHERE id = $1 LIMIT 2;

-- name: ListPhases :many
SELECT * FROM phases
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreatePhase :one
INSERT INTO phases (
  name, quiz_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdatePhase :one
UPDATE phases
  set name = $2
WHERE id = $1
RETURNING *;

-- name: DeletePhase :exec
DELETE FROM phases
WHERE id = $1;

