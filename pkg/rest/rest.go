package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)
	app.Get("/todo/:id", controller.TodoGet)
	app.Get("/todo", controller.TodoGetAll)
	app.Post("/todo", controller.TodoPost)
	app.Put("/todo", controller.TodoPut)
	app.Delete("/todo", controller.TodoDel)
	app.Get("/config", controller.Config)

	return app
}
