package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	idParam := c.Params("id_user")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	if err := repository.PromoteUserToAdmin(db, uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to promote user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User promoted to admin successfully",
	})
}
