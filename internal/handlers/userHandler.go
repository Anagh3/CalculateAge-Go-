package handlers

import (
	"strconv"
	"time"

	"github.com/Anagh3/go-backend/internal/models"
	"github.com/Anagh3/go-backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	Service  *service.UserService
	Validate *validator.Validate
	Logger   *zap.Logger
}

func NewUserHandler(s *service.UserService, v *validator.Validate, l *zap.Logger) *UserHandler {
	return &UserHandler{Service: s, Validate: v, Logger: l}
}

// POST /users
func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var body models.UserInput
	if err := c.BodyParser(&body); err != nil {
		h.Logger.Error("JSON parse error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}

	if err := h.Validate.Struct(body); err != nil {
		h.Logger.Warn("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dob, _ := time.Parse("2006-01-02", body.Dob)
	if err := h.Service.AddUser(c.Context(), body.Name, dob); err != nil {
		h.Logger.Error("Add user failed", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.Logger.Info("User created", zap.String("name", body.Name))
	return c.SendStatus(fiber.StatusCreated)
}

// GET /users
func (h *UserHandler) GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := h.Service.ListUsers(c.Context())
	if err != nil {
		h.Logger.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GET /users/:id
func (h *UserHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		h.Logger.Warn("Invalid user ID", zap.String("id", idParam))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.Service.GetUser(c.Context(), id)
	if err != nil {
		h.Logger.Warn("User not found", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

// PUT /users/:id
func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	var body models.UserInput
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}

	if err := h.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dob, _ := time.Parse("2006-01-02", body.Dob)
	if err := h.Service.UpdateUser(c.Context(), id, body.Name, dob); err != nil {
		h.Logger.Error("Update user failed", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.Logger.Info("User updated", zap.Uint64("id", id))
	return c.SendStatus(fiber.StatusOK)
}

// DELETE /users/:id
func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	if err := h.Service.DeleteUser(c.Context(), id); err != nil {
		h.Logger.Error("Delete user failed", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.Logger.Info("User deleted", zap.Uint64("id", id))
	return c.SendStatus(fiber.StatusNoContent)
}
