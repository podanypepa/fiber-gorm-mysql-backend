package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	return app
}
