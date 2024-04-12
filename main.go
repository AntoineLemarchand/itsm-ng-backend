package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsmng/itsm-ng-backend/database"
	"github.com/itsmng/itsm-ng-backend/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

    router.SetupRoutes(app)

	app.Listen(":3000")
}
