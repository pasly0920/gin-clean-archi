package main

import (
	"gin-clean-archi/pkg/users"
	"log"

	"gin-clean-archi/pkg/common/config"
	"gin-clean-archi/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)
	app.Use(recover())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("")
	})

	users.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
