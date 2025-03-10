-- name: GetAnswerByQuestion :one
SELECT * FROM answers
WHERE question_id = $1 LIMIT 1;

-- name: GetAnswer :one
SELECT * FROM answers
WHERE id = $1 LIMIT 1;

-- name: ListAnswers :many
SELECT * FROM answers
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateAnswer :one
INSERT INTO answers (
  text, question_id
) VALUES (
  $1, $2 
)
RETURNING *;

-- name: UpdateAnswer :one
UPDATE answers
  set text = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAnswer :exec
DELETE FROM answers
WHERE id = $1;


