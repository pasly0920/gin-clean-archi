package users

import (
	"gin-clean-archi/pkg/common/model"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user model.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete book from db
	h.DB.Delete(&user)

	return c.SendStatus(fiber.StatusOK)
}
