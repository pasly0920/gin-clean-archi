package users

import (
	"gin-clean-archi/pkg/common/model"
	"github.com/gofiber/fiber/v2"
)

type AddUserRequestBody struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	AvatarID     uint64 `json:"avatarId"`
	PhoneNumber  string `json:"phoneNumber"`
	StudentYear  int    `json:"studentYear"`
	StudentGroup string `json:"studentGroup"`
	MajorCode    int    `json:"majorCode"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := AddUserRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user := h.getUserFromBody(body)

	// insert new db entry
	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func (h handler) getUserFromBody(body AddUserRequestBody) model.User {
	var user model.User

	user.Username = body.Username
	user.Name = body.Name
	user.Nickname = body.Nickname
	user.AvatarID = body.AvatarID
	user.PhoneNumber = body.PhoneNumber
	user.StudentYear = body.StudentYear
	user.StudentGroup = body.StudentGroup
	user.MajorCode = body.MajorCode

	return user
}
