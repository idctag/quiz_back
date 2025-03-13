-- Drop indexes first (no dependencies)
DROP INDEX IF EXISTS idx_answers_created_at;
DROP INDEX IF EXISTS idx_choices_created_at;
DROP INDEX IF EXISTS idx_questions_created_at;
DROP INDEX IF EXISTS idx_phases_created_at;
DROP INDEX IF EXISTS idx_quizzes_created_at;

DROP INDEX IF EXISTS idx_answers_question_id;
DROP INDEX IF EXISTS idx_choices_question_id;
DROP INDEX IF EXISTS idx_questions_phase_id;
DROP INDEX IF EXISTS idx_phases_quiz_id;

-- Drop tables (in reverse order of creation to respect foreign key dependencies)
DROP TABLE IF EXISTS "answers";
DROP TABLE IF EXISTS "choices";
DROP TABLE IF EXISTS "questions";
DROP TABLE IF EXISTS "phases";
DROP TABLE IF EXISTS "quizzes";

-- Drop the ENUM type last (after tables that use it are gone)
DROP TYPE IF EXISTS question_types;
