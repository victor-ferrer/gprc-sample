package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}
func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	if err := h.db.Ping(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "unhealthy",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
func (h *HealthHandler) UpCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ready",
	})
}
