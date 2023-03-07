package users

import (
	"gin-clean-archi/pkg/common/model"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user model.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
