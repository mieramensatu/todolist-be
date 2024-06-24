package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/model"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
	"net/http"
)

func GetAllTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	user := c.Locals("user").(*model.Users)

	tasks, err := repository.GetAllTasksByUserId(db, user.IdUser, user.IdRole == 1)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve tasks",
		})
	}

	return c.JSON(tasks)
}


func GetTaskById(c *fiber.Ctx) error {
	var data struct {
		IdTask string `json:"id_task"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	task, err := repository.GetTaskById(db, data.IdTask)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(task)
}

func GetUserTasks(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.Users)
	db := c.Locals("db").(*gorm.DB)

	tasks, err := repository.GetTasksByUserId(db, user.IdUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(tasks)
}

func InsertTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.Users)
	db := c.Locals("db").(*gorm.DB)

	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task.IdUser = user.IdUser

	if err := repository.InsertTask(db, task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	var data struct {
		IdTask string       `json:"id_task"`
		Task   model.Task `json:"task"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	if err := repository.UpdateTask(db, data.IdTask, data.Task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task updated successfully",
	})
}

func DeleteTask(c *fiber.Ctx) error {
	var data struct {
		IdTask string `json:"id_task"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db := c.Locals("db").(*gorm.DB)
	if err := repository.DeleteTask(db, data.IdTask); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}