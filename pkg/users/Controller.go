package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/api/user")
	routes.Get("/:id", h.GetUser)
	routes.Delete("/:id", h.DeleteUser)
	routes.Post("/", h.AddUser)
	routes.Post("/:id", h.UpdateUser)
}
