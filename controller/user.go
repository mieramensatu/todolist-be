package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/model"
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
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id_user")

	if err := repository.DeleteUserById(db, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func PromoteUserToAdmin(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	user := c.Locals("user").(*model.Users)

	// Hanya admin yang bisa mengakses endpoint ini
	if user.IdRole != 1 {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	// Dapatkan id user yang ingin di-promote dari body request
	var requestBody struct {
		IDUser uint `json:"id_user"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	// Panggil repository untuk melakukan promosi user ke admin
	if err := repository.PromoteUserToAdmin(db, requestBody.IDUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User promoted to admin successfully",
	})
}