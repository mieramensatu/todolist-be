package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mieramensatu/todolist-be/model"
	"github.com/mieramensatu/todolist-be/repository"
	"gorm.io/gorm"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})
		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired JWT",
			})
		}

		claims := token.Claims.(*model.JWTClaims)
		db := c.Locals("db").(*gorm.DB)
		user, err := repository.GetUserById(db, claims.IdUser)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		c.Locals("user", user)

		// Check if user is admin (id_role == 1) to authorize access to certain endpoints
		if user.IdRole != 1 {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: Admin access required",
			})
		}

		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.Users)

		if user.IdRole != 1 { // Ubah sesuai dengan role admin di database Anda
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"error": "Akses ditolak",
			})
		}

		return c.Next()
	}
}