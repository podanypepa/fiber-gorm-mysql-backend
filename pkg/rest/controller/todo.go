package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/database"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/model/todo"
)

// TodoGet GET /todo
func TodoGet(c *fiber.Ctx) error {
	id := c.Params("id")
	t := &todo.TODO{}
	if err := todo.Read(database.DB, t, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{
		Status: "ok",
		Data:   &[]todo.TODO{*t},
	})
}

func TodoGetAll(c *fiber.Ctx) error {
	t := &[]todo.TODO{}
	if err := todo.ReadAll(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{
		Status: "ok",
		Data:   t,
	})
}

// TodoPost POST /todo
func TodoPost(c *fiber.Ctx) error {
	t := &todo.TODO{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := todo.Create(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}

// TodoPut PUT /todo
func TodoPut(c *fiber.Ctx) error {
	t := &todo.TODO{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := todo.Update(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}

// TodoDel DELETE /todo/:id
func TodoDel(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := todo.DeleteByID(database.DB, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}
