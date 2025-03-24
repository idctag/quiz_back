package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/idctag/quiz_back/api/params"
	"github.com/idctag/quiz_back/db"
	sqlc "github.com/idctag/quiz_back/db/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	timeOut  = 5 * time.Second
	validate = validator.New()
)

func GetQuiz(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	quiz, err := sqlc.New(db.DB).GetQuiz(ctx, int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(quiz)
}

func GetQuizzes(c fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), timeOut)
	defer cancle()

	limitStr := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}

	offsterStr := c.Query("offset", "0")
	offset, err := strconv.Atoi(offsterStr)
	if err != nil || offset < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}

	arg := sqlc.ListQuizzesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	quizzes, err := sqlc.New(db.DB).ListQuizzes(ctx, arg)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to fetch quizzes",
		})
	}

	return c.JSON(quizzes)
}

func CreateQuiz(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	// Parse request body
	var req params.CreateQuizRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	// Start transaction
	tx, err := db.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer tx.Rollback(ctx)

	quiz, err := sqlc.New(tx).CreateQuiz(ctx, req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create quiz",
		})
	}

	for _, phaseData := range req.Phases {
		phaseParams := sqlc.CreatePhaseParams{
			QuizID: quiz.ID,
			Name:   phaseData.Name,
		}
		phase, err := sqlc.New(tx).CreatePhase(ctx, phaseParams)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create phase",
			})
		}

		for _, qData := range phaseData.Questions {
			questionParams := sqlc.CreateQuestionParams{
				PhaseID: phase.ID,
				Text:    qData.Text,
				Types:   sqlc.QuestionTypes(qData.Type),
				ImgUrl: pgtype.Text{
					String: qData.ImgURL,
				},
				AudioUrl: pgtype.Text{
					String: qData.AudioURL,
				},
			}
			question, err := sqlc.New(tx).CreateQuestion(ctx, questionParams)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to create question",
				})
			}

			// Insert Answers
			for _, answerText := range qData.Answers {
				answerParams := sqlc.CreateAnswerParams{
					QuestionID: question.ID,
					Text:       answerText,
				}
				if _, err := sqlc.New(tx).CreateAnswer(ctx, answerParams); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to create answer",
					})
				}
			}

			// Insert Choices
			for _, choiceText := range qData.Choices {
				choiceParams := sqlc.CreateChoiceParams{
					QuestionID: question.ID,
					Text:       choiceText,
				}
				if _, err := sqlc.New(tx).CreateChoice(ctx, choiceParams); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to create choice",
					})
				}
			}
		}
	}
	// Commit transaciton
	if err := tx.Commit(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Quiz created Successfully",
		"id":      quiz.ID,
	})
}
