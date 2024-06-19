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
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	task, err := repository.GetTaskById(db, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve task",
		})
	}

	return c.JSON(task)
}

func GetUserTasks(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	user := c.Locals("user").(*model.Users)

	tasks, err := repository.GetTasksByUserId(db, user.IdUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user tasks",
		})
	}

	return c.JSON(tasks)
}

func InsertTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	user := c.Locals("user").(*model.Users)

	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	task.IdUser = user.IdUser

	if err := repository.InsertTask(db, task); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert task",
		})
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	var updatedTask model.Task
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := repository.UpdateTask(db, id, updatedTask); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task updated successfully",
	})
}

func DeleteTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	if err := repository.DeleteTask(db, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
