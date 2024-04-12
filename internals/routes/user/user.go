package user

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/itsmng/itsm-ng-backend/internal/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {
    user := router.Group("/user")

    user.Post("/", userHandler.CreateUser)

    user.Get("/", userHandler.GetUsers)

    user.Get("/:userId", userHandler.GetUser)

    user.Put("/:userId", userHandler.UpdateUser)

    user.Delete("/:userId", userHandler.DeleteUser)
}
