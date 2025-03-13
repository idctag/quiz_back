package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/idctag/quiz_back/api/controller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	// Quiz
	api.Get("/quiz", controller.GetQuizzes)
}
