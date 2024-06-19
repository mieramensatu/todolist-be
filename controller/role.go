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
	db := c.Locals("db").(*gorm.DB)

	id := c.Params("id_role")
	role, err := repository.GetRoleById(db, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve role",
		})
	}

	return c.JSON(role)
}
