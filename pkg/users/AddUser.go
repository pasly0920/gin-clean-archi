package users

import (
	"gin-clean-archi/pkg/common/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

func PhoneValidation(fl validator.FieldLevel) bool {
	phoneRegex := regexp.MustCompile(`^\d{3}-\d{4}-\d{4}$`)
	return phoneRegex.MatchString(fl.Field().String())
}

type AddUserRequestBody struct {
	Username     string `json:"username" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Nickname     string `json:"nickname" validate:"required"`
	AvatarID     uint64 `json:"avatarId" validate:"required,number"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,phone"`
	StudentYear  int    `json:"studentYear" validate:"required,number"`
	StudentGroup string `json:"studentGroup" validate:"required"`
	MajorCode    int    `json:"majorCode" validate:"required,number"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := AddUserRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	validate.RegisterValidation("phone", PhoneValidation)

	//validate error
	if err := validate.Struct(body); err != nil {
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
