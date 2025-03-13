package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/idctag/quiz_back/api/route"
	"github.com/idctag/quiz_back/db"
)

func init() {
	db.ConnectDB()
}

func main() {
	app := fiber.New()
	route.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
