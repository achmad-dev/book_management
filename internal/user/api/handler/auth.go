package handler

import (
	"github.com/achmad-dev/internal/user/api/dto"
	"github.com/achmad-dev/internal/user/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}

type authHandlerImpl struct {
	userService service.UserService
}

// SignIn implements AuthHandler.
func (a *authHandlerImpl) SignIn(c *fiber.Ctx) error {
	var signInRequest dto.SignInRequest
	if err := c.BodyParser(&signInRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	token, err := a.userService.SignIn(c.Context(), signInRequest)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// SignUp implements AuthHandler.
func (a *authHandlerImpl) SignUp(c *fiber.Ctx) error {
	var signUpRequest dto.SignUpRequest
	if err := c.BodyParser(&signUpRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	err := a.userService.CreateUser(c.Context(), signUpRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return &authHandlerImpl{userService: userService}
}
