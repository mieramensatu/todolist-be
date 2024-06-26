package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
)

// GetAllRole gets all roles
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

// GetRoleById gets role by id_role
func GetRoleById(c *fiber.Ctx) error {
    idParam := c.Query("id_role")
	idRole, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid role ID",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	role, err := repository.GetRoleById(db, strconv.FormatUint(idRole, 10))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Role not found",
		})
	}

	return c.JSON(role)
}