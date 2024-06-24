package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	users, err := repository.GetAllUsers(db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(users)
}

func DeleteUserById(c *fiber.Ctx) error {
	var data struct {
		IdUser string `json:"id_user"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	if err := repository.DeleteUserById(db, data.IdUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func PromoteUserToAdmin(c *fiber.Ctx) error {
	var data struct {
		IdUser uint `json:"id_user"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	user, err := repository.GetUserById(db, data.IdUser)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if user.IdRole == 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User is already an admin",
		})
	}

	if err := repository.PromoteUserToAdmin(db, data.IdUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User promoted to admin successfully",
	})
}