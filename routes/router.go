package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/controller"
	"github.com/mieramensatu/todolist-be/middleware"
)

func SetupTaskRoutes(app *fiber.App) {
	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.LoginUser)

	app.Get("/getme", middleware.Auth(), controller.GetMe)

	adminOnly := middleware.AdminOnly()

	admin := app.Group("")
	admin.Use(middleware.Auth())

	admin.Get("/tasks", adminOnly, controller.GetAllTask)
	admin.Get("/users", adminOnly, controller.GetAllUsers)
	admin.Get("/roles", adminOnly, controller.GetAllRole)
	admin.Get("/role/get", adminOnly, controller.GetRoleById)        
	admin.Delete("/user/delete", adminOnly, controller.DeleteUserById)  
	admin.Post("/promoteuser", adminOnly, controller.PromoteUserToAdmin)

	app.Get("/task", middleware.Auth(), controller.GetUserTasks)
	app.Get("/task/get", middleware.Auth(), controller.GetTaskById)  
	app.Post("/task/insert", middleware.Auth(), controller.InsertTask)
	app.Put("/task/update", middleware.Auth(), controller.UpdateTask) 
	app.Delete("/task/delete", middleware.Auth(), controller.DeleteTask)
}
