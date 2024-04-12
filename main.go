package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/itsmng/itsm-ng-backend/database"
	"github.com/itsmng/itsm-ng-backend/router"
)

// @title ITSM-NG API
// @version 1.0
// @description ITSM-NG API Documentation

func main() {
	app := fiber.New()

    app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
