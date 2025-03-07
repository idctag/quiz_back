-- name: ListChoices :many
SELECT * FROM choices
ORDER BY question_id DESC
LIMIT $1 OFFSET $2;

-- name: ListChoicesByQuestion :many
SELECT * FROM choices
WHERE question_id = $1
ORDER BY id DESC;

-- name: CreateChoice :one
INSERT INTO choices (
  text, question_id
) VALUES (
  $1, $2 
)
RETURNING *;

-- name: UpdateChoice :one
UPDATE choices
  set text = $2
WHERE id = $1
RETURNING *;

-- name: DeleteChoice :exec
DELETE FROM choices
WHERE id = $1;

