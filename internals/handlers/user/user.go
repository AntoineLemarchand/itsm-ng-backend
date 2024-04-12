package userHandler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "github.com/itsmng/itsm-ng-backend/database"
    "github.com/itsmng/itsm-ng-backend/internals/model"
)

func GetUsers(c *fiber.Ctx) error {
    db := database.DB
    users := []model.User{}

    db.Find(&users)

    if len(users) == 0 {
        return c.Status(404).JSON(fiber.Map{ "status": "error", "message": "No users present", "data": nil})
    }

    return c.JSON(fiber.Map{ "status": "success", "message": "Users Found", "data": users})
}


// @Summary         Create a new user
// @Description     Create a new user
// @Tags            user
// @Accept          json
// @Produce         json
// @Param           user body           model.AddUser   true    "User object that needs to be created"
// @Success         200                 {object}        model.UserResponse
// @Failure         400                 {object}        httputil.HTTPError
// @Failure         500                 {object}        httputil.HTTPError
// @Router          /user [post]
func CreateUser(c *fiber.Ctx) error {
    db := database.DB
    user := new(model.User)

    err := c.BodyParser(user)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{ "status": "error", "message": "Review your input", "data": err})
    }

    user.ID = uuid.New()
    err = db.Create(&user).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{ "status": "error", "message": "Could not create note", "data": err})
    }

    return c.JSON(fiber.Map{ "status": "success", "message": "Created User", "data": user})
}

func GetUser(c *fiber.Ctx) error {
    db := database.DB
    var user model.User

    id := c.Params("userId")

    db.Find(&user, "id = ?", id)

    if user.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{ "status": "error", "message": "No note present", "data": nil})
    }

    return c.JSON(fiber.Map{ "status": "success", "message": "User Found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
    type updateUser struct {
        Name    string `json:"name"`
        Email string `json:"email"`
        Password     string `json:"Password"`
    }
    db := database.DB
    var user model.User

    id := c.Params("userId")

    db.Find(&user, "id = ?", id)

    if user.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
    }

    var updateNoteData updateUser
    err := c.BodyParser(&updateNoteData)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }

    user.Name = updateNoteData.Name
    user.Email = updateNoteData.Email
    user.Password = updateNoteData.Password

    db.Save(&user)

    return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
    db := database.DB
    var user model.User

    // Read the param noteId
    id := c.Params("userId")

    // Find the note with the given Id
    db.Find(&user, "id = ?", id)

    // If no such note present return an error
    if user.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
    }

    // Delete the note and return error if encountered
    err := db.Delete(&user, "id = ?", id).Error

    if err != nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
    }

    // Return success message
    return c.JSON(fiber.Map{"status": "success", "message": "Deleted user"})
}
