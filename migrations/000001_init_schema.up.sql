CREATE TYPE question_types AS ENUM ('img', 'audio' );

CREATE TABLE "quizzes" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "phases" (
  "id" BIGSERIAL PRIMARY KEY,
  "quiz_id" BIGINT NOT NULL,
  "name" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "questions" (
  "id" BIGSERIAL PRIMARY KEY,
  "phase_id" BIGINT NOT NULL,
  "text" TEXT NOT NULL,
  "types" question_types NOT NULL,
  "img_url" TEXT,
  "audio_url" TEXT,
  "is_multiple_choice" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "choices" (
  "id" BIGSERIAL PRIMARY KEY,
  "question_id" BIGINT NOT NULL,
  "text" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "answers" (
  "id" BIGSERIAL PRIMARY KEY,
  "question_id" BIGINT NOT NULL,
  "text" TEXT NOT NULL,
  "is_correct" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

ALTER TABLE "phases" ADD FOREIGN KEY ("quiz_id") REFERENCES "quizzes" ("id") ON DELETE CASCADE;

ALTER TABLE "questions" ADD FOREIGN KEY ("phase_id") REFERENCES "phases" ("id") ON DELETE CASCADE;
 
ALTER TABLE "answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id") ON DELETE CASCADE;

ALTER TABLE "choices" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id") ON DELETE CASCADE;

-- Existing Indexes for Foreign Keys
CREATE INDEX idx_phases_quiz_id ON "phases" ("quiz_id");
CREATE INDEX idx_questions_phase_id ON "questions" ("phase_id");
CREATE INDEX idx_choices_question_id ON "choices" ("question_id");
CREATE INDEX idx_answers_question_id ON "answers" ("question_id");

-- New Indexes for Timestamps
CREATE INDEX idx_quizzes_created_at ON "quizzes" ("created_at");
CREATE INDEX idx_phases_created_at ON "phases" ("created_at");
CREATE INDEX idx_questions_created_at ON "questions" ("created_at");
CREATE INDEX idx_choices_created_at ON "choices" ("created_at");
CREATE INDEX idx_answers_created_at ON "answers" ("created_at");
