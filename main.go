package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database sebelum init fiber
	database.DatabaseInit()

	// jalankan migration setelah konek ke database
	migration.RunMigration()

	// init fiber
	app := fiber.New()

	// init base route
	route.RouteInit(app)

	app.Listen(":3000")
}
