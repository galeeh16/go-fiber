package route

import (
	"go-fiber-gorm/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "API With Go Fiber v2.42.0",
		})
	})

	r.Get("/users", controller.GetAllUser)
	r.Post("/users", controller.CreateUser)
}
