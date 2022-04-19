package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Index GET /
func Index(c *fiber.Ctx) error {
	return c.JSON(struct {
		Status string `json:""`
		TS     time.Time
	}{
		Status: "ok",
		TS:     time.Now(),
	})
}
