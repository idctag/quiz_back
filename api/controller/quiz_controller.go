package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/idctag/quiz_back/db"
	sqlc "github.com/idctag/quiz_back/db/models"
)

var (
	timeOut = 5 * time.Second
	ctx     = context.Background()
)

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
	return c.JSON("Hello create quiz endpoint")
}
