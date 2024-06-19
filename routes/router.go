// routes/router.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/controller"
	"github.com/mieramensatu/todolist-be/middleware"
)

func SetupTaskRoutes(app *fiber.App) {
	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.LoginUser)

	// Endpoint untuk cek sudah login atau belum
	app.Get("/getme", middleware.Auth(), controller.GetMe)

	// Middleware AdminOnly untuk membatasi akses hanya ke admin
	adminOnly := middleware.AdminOnly()

	// Protected routes for admin only
	admin := app.Group("")
	admin.Use(middleware.Auth()) // Ensure all routes in admin group require authentication

	admin.Get("/tasks", adminOnly, controller.GetAllTask)
	admin.Get("/users", adminOnly, controller.GetAllUsers)
	admin.Get("/roles", adminOnly, controller.GetAllRole)
	admin.Get("/role/get/:id_role", adminOnly, controller.GetRoleById)
	admin.Delete("/user/delete/:id_user", adminOnly, controller.DeleteUserById)
	admin.Post("/promoteuser", adminOnly, controller.PromoteUserToAdmin)

	// Non-admin specific routes
	app.Get("/tasks", middleware.Auth(), controller.GetUserTasks)
	app.Get("/task/get/:id_task", middleware.Auth(), controller.GetTaskById)
	app.Post("/task/insert", middleware.Auth(), controller.InsertTask)
	app.Put("/task/update", middleware.Auth(), controller.UpdateTask)
	app.Delete("/task/delete/:id_task", middleware.Auth(), controller.DeleteTask)
}
