package users

import (
	"gin-clean-archi/pkg/common/model"
	"github.com/gofiber/fiber/v2"
)

type UpdateUserRequestBody struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	AvatarID     uint64 `json:"avatarId"`
	PhoneNumber  string `json:"phoneNumber"`
	StudentYear  int    `json:"studentYear"`
	StudentGroup string `json:"studentGroup"`
	MajorCode    int    `json:"majorCode"`
}

func updateUserFields(user *model.User, body UpdateUserRequestBody) {
	user.Username = body.Username
	user.Name = body.Name
	user.Nickname = body.Nickname
	user.AvatarID = body.AvatarID
	user.PhoneNumber = body.PhoneNumber
	user.StudentYear = body.StudentYear
	user.StudentGroup = body.StudentGroup
	user.MajorCode = body.MajorCode
}

func (h handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateUserRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user model.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	updateUserFields(&user, body)

	// save user
	h.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(&user)
}
