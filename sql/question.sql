-- name: GetQuestion :one
SELECT * FROM questions
WHERE id = $1 LIMIT 1;

-- name: ListQuestions :many
SELECT * FROM questions
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateQuestion :one
INSERT INTO questions (
  phase_id, text, types, img_url, audio_url, is_multiple_choice 
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateQuestion :one
UPDATE questions
  set text = $2,
  types = $3,
  img_url = $4,
  audio_url = $4,
  is_multiple_choice = $5
WHERE id = $1
RETURNING *;

-- name: DeleteQuestion :exec
DELETE FROM phases
WHERE id = $1;

