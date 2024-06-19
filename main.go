package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mieramensatu/todolist-be/config"
	"github.com/mieramensatu/todolist-be/routes"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Koneksi ke database
	db := config.CreateDBConnection()

	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Menyimpan koneksi database dalam context Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Menetapkan rute untuk handler buku
	routes.SetupTaskRoutes(app)

	// Menjalankan server Fiber
	app.Listen(":3000")
}
