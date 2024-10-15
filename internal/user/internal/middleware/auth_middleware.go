package middleware

import (
	"log"
	"strings"

	"github.com/achmad-dev/internal/user/internal/service"
	"github.com/achmad-dev/internal/user/internal/util"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(secret string, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		claims, err := util.ValidateToken(tokenString, secret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		_, err = userService.GetUserByUsername(c.Context(), claims.Username)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		c.Locals("username", claims.Username)
		return c.Next()
	}
}

func AuthAdminMiddleware(secret string, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header missing")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			log.Println("Token missing after Bearer")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		claims, err := util.ValidateToken(tokenString, secret)
		if err != nil {
			log.Printf("Token validation failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		user, err := userService.GetUserByUsername(c.Context(), claims.Username)
		if err != nil {
			log.Printf("User not found: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		if user.Role != "admin" {
			log.Println("User is not an admin")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		c.Locals("username", claims.Username)
		return c.Next()
	}
}
