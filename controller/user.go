package controller

import (
	"net/http"
	"strconv"

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
    idParam := c.Query("id_user")
    idUser, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    db := c.Locals("db").(*gorm.DB)
	if err := repository.DeleteUserById(db, strconv.FormatUint(idUser, 10)); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete user",
        })
    }

    return c.JSON(fiber.Map{
        "message": "User successfully deleted",
    })
}

func PromoteUserToAdmin(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	idUser := c.Query("id_user")

	id, err := strconv.Atoi(idUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := repository.GetUserById(db, uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if user.IdRole == 1 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "User is already an admin",
		})
	}

	if err := repository.PromoteUserToAdmin(db, user.IdUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to promote user to admin",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User promoted to admin successfully",
	})
}

func GetUserById(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	idUser := c.Query("id_user")

	id, err := strconv.Atoi(idUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := repository.GetUserById(db, uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(http.StatusOK).JSON(user)
}