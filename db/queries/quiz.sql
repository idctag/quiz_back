-- name: GetFullQuiz :one
SELECT 
  q.id AS quiz_id,
  q.name AS quiz_name,
  p.id AS phase_id,
  p.name AS phase_name,
  qs.id AS phase_name,
  qs.text AS question_text,
  qs.types AS question_type,
  a.id AS answer_id,
  a.text AS answer_text,
  c.id AS choice_id,
  c.text AS choice_text
FROM quizzes q 
LEFT JOIN phases p ON q.id = p.quiz_id
LEFT JOIN questions qs ON p.id = qs.phase_id
LEFT JOIN answers a ON qs.id = a.question_id
LEFT JOIN choices c ON qs.id = c.question_id
WHERE q.id = $1;

-- name: GetQuiz :one
SELECT * FROM quizzes
WHERE id = $1 LIMIT 1;

-- name: ListQuizzes :many
SELECT * FROM quizzes
ORDER BY created_at ASC
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
