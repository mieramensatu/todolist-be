package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
)

func GetAllRole(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	roles, err := repository.GetAllRole(db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve roles",
		})
	}

	return c.JSON(roles)
}

func GetRoleById(c *fiber.Ctx) error {
	var data struct {
		IdRole string `json:"id_role"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	role, err := repository.GetRoleById(db, data.IdRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(role)
}