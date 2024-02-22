package main

import (
	"day06/internal/db"
	"day06/internal/handlers"
	"day06/internal/logger"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	lgr := logger.New()
	defer lgr.Sync()

	app := fiber.New()
	database := db.New(lgr)
	handler := handlers.New(database, lgr)
	handler.InitRoutes(app)

	log.Fatal(app.Listen(":8888"))
}
