package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/model"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
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
    idParam := c.Params("id_task")
    idTask, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid task ID",
        })
    }

    db := c.Locals("db").(*gorm.DB)
    task, err := repository.GetTaskById(db, strconv.FormatUint(idTask, 10))
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.JSON(task)
}

// GetUserTasks gets tasks by user_id
func GetUserTasks(c *fiber.Ctx) error {
    user := c.Locals("user").(*model.Users)
    db := c.Locals("db").(*gorm.DB)

    tasks, err := repository.GetTasksByUserId(db, user.IdUser)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Tasks not found",
        })
    }

    return c.JSON(tasks)
}

func InsertTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	user := c.Locals("user").(*model.Users)

	var task model.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Validate due_date format
	if _, err := time.Parse("2006-01-02", task.DueDate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid due_date format. Use YYYY-MM-DD",
		})
	}

	task.IdUser = user.IdUser
	if err := repository.InsertTask(db, &task); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert task",
		})
	}

	return c.Status(http.StatusCreated).JSON(task)
}

// UpdateTask updates task by id_task
func UpdateTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	idTask := c.Query("id_task")

	var updatedTask model.Task
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Validate due_date format
	if _, err := time.Parse("2006-01-02", updatedTask.DueDate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid due_date format. Use YYYY-MM-DD",
		})
	}

	if err := repository.UpdateTask(db, idTask, updatedTask); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
	})
}

// DeleteTask deletes task by id_task
func DeleteTask(c *fiber.Ctx) error {
    idParam := c.Query("id_task")
    idTask, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid task ID",
        })
    }

    db := c.Locals("db").(*gorm.DB)
	if err := repository.DeleteTask(db, strconv.FormatUint(idTask, 10)); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete task",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Task successfully deleted",
    })
}