package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mieramensatu/todolist-be/config"
	"github.com/mieramensatu/todolist-be/routes"
)

func main() {
	// Membuat aplikasi Fiber
	app := fiber.New()

	// Koneksi ke database
	db := config.CreateDBConnection()

	// Middleware untuk logging
	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	// Middleware CORS
	app.Use(cors.New(cors.Config{
    	AllowOrigins: "*",
    	AllowMethods: "GET, POST, PUT, DELETE",
    	AllowHeaders: "*",
	}))

	// Middleware untuk menyimpan koneksi database dalam context Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Setup routes
	routes.SetupTaskRoutes(app)

	// Menjalankan server Fiber pada port 6969
	app.Listen(":6969")
}
