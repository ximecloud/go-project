package http

import (
	"github.com/gofiber/fiber/v2"
	"project/internal/core/application/dto"
	"project/internal/port"
)

type UserHandler struct {
	user port.UserService
}

func (r *UserHandler) Route(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/create", r.Create)
}

func (r *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	user, err := r.user.Create(req)
	if err != nil {
		return handleError(c, err)
	}
	return handleSuccess(c, user)
}

func NewUserHandler(user port.UserService) *UserHandler {
	return &UserHandler{user: user}
}
