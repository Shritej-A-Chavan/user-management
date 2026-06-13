package handler

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"go-task/internal/logger"
	"go-task/internal/models"
	"go-task/internal/service"
	"go-task/internal/validator"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.UserRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn(
			"invalid request body",
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	validationErrors := validator.ValidateStruct(req)

	if validationErrors != nil {
		logger.Log.Warn(
			"user validation failed",
			zap.Any("errors", validationErrors),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validationErrors,
		})
	}

	user, err := h.service.CreateUser(
		c.Context(),
		req.Name,
		req.Dob,
	)

	if err != nil {
		logger.Log.Warn(
			"user creation failed",
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Log.Info(
		"user created successfully",
		zap.Int64("id", user.ID),
		zap.String("name", user.Name),
	)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		logger.Log.Warn(
			"invalid user id",
			zap.String("id", c.Params("id")),
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	user, err := h.service.GetUserByID(
		c.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Log.Warn(
				"user not found",
				zap.Int64("id", id),
			)

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		logger.Log.Error(
			"failed to fetch user",
			zap.Int64("id", id),
			zap.Error(err),
		)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	logger.Log.Info(
		"user fetched successfully",
		zap.Int64("id", user.ID),
	)

	return c.JSON(models.UserResponseWithAge{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  service.CalculateAge(user.Dob),
	})
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(
		c.Query("limit", "10"),
	)

	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(
		c.Query("page", "1"),
	)

	if err != nil || page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	users, err := h.service.GetUsers(
		c.Context(),
		int32(limit),
		int32(offset),
	)

	if err != nil {
		logger.Log.Error(
			"failed to fetch users",
			zap.Error(err),
		)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}

	var response []models.UserResponseWithAge

	for _, user := range users {
		response = append(response, models.UserResponseWithAge{
			ID:   user.ID,
			Name: user.Name,
			Dob:  user.Dob.Format("2006-01-02"),
			Age:  service.CalculateAge(user.Dob),
		})
	}

	logger.Log.Info(
		"users fetched successfully",
		zap.Int("page", page),
		zap.Int("limit", limit),
		zap.Int("count", len(response)),
	)

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"count": len(response),
		"data":  response,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		logger.Log.Warn(
			"invalid user id",
			zap.String("id", c.Params("id")),
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	var req models.UserRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn(
			"invalid request body",
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	validationErrors := validator.ValidateStruct(req)

	if validationErrors != nil {
		logger.Log.Warn(
			"user validation failed",
			zap.Any("errors", validationErrors),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validationErrors,
		})
	}

	user, err := h.service.UpdateUser(
		c.Context(),
		id,
		req.Name,
		req.Dob,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Log.Warn(
				"user not found",
				zap.Int64("id", id),
			)

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		logger.Log.Warn(
			"user update failed",
			zap.Int64("id", id),
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Log.Info(
		"user updated successfully",
		zap.Int64("id", user.ID),
	)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		logger.Log.Warn(
			"invalid user id",
			zap.String("id", c.Params("id")),
			zap.Error(err),
		)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	err = h.service.DeleteUser(
		c.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Log.Warn(
				"user not found",
				zap.Int64("id", id),
			)

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		logger.Log.Error(
			"failed to delete user",
			zap.Int64("id", id),
			zap.Error(err),
		)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}

	logger.Log.Info(
		"user deleted successfully",
		zap.Int64("id", id),
	)

	return c.SendStatus(fiber.StatusNoContent)
}
