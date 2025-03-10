-- name: GetQuiz :one
SELECT * FROM quizzes
WHERE id = $1 LIMIT 1;

-- name: ListQuizzes :many
SELECT * FROM quizzes
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateQuiz :one
INSERT INTO quizzes (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateQuiz :one
UPDATE quizzes
  set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteQuiz :exec
DELETE FROM quizzes
WHERE id = $1;
