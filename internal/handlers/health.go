package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h HealthHandler) UpCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "up",
	})
}

func (h HealthHandler) HealthCheck(c *fiber.Ctx) error {
	// TODO: Actully check the DB
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
