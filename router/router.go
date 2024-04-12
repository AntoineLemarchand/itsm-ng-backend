package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	userRoutes "github.com/itsmng/itsm-ng-backend/internals/routes/user"
    _ "github.com/itsmng/itsm-ng-backend/docs"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api", logger.New())

    app.Get("/swagger/*", swagger.HandlerDefault)

    userRoutes.SetupUserRoutes(api)
}
