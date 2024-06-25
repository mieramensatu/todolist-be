package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/model"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var user model.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := repository.CreateUser(db, &user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	user, err := repository.GetUserByUsername(db, request.Username)
	if err != nil || !repository.CheckPasswordHash(request.Password, user.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	token, err := repository.GenerateToken(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	return c.JSON(fiber.Map{
		
		"token": token,
	})
}

func GetMe(c *fiber.Ctx) error {
    user := c.Locals("user").(*model.Users)
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "user": user,
    })
}