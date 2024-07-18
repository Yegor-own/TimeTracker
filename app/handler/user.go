package handler

import (
	"net/http"
	"time-tracker/app/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.UserCreate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		user := model.User{
			Name:       input.Name,
			Surname:    input.Surname,
			Patronymic: input.Patronymic,
			Address:    input.Address,
		}

		res := db.Create(&user)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  "user create successful",
			Error: "",
		})
	}
}
