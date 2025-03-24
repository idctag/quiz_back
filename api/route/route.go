package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/idctag/quiz_back/api/controller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// Quiz
	quiz := api.Group("/quiz")
	quiz.Get("/:id", controller.GetQuiz)
	quiz.Get("/", controller.GetQuizzes)
	quiz.Post("/", controller.CreateQuiz)
}
