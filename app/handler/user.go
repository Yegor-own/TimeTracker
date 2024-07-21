package handler

import (
	"fmt"
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
			Data:  user,
			Error: "",
		})
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.UserUpdate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		user := model.User{
			ID: input.UserID,
		}
		res := db.First(&user)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		if input.Name != nil && *input.Name != "" {
			user.Name = *input.Name
		}
		if input.Surname != nil && *input.Surname != "" {
			user.Surname = *input.Surname
		}
		if input.Patronymic != nil && *input.Patronymic != "" {
			user.Patronymic = input.Patronymic
		}
		if input.Address != nil && *input.Address != "" {
			user.Address = *input.Address
		}

		res = db.Save(&user)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  user,
			Error: "",
		})
	}
}

func DeliteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.UserDelete
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		user := model.User{
			ID: input.UserID,
		}

		res := db.Delete(&user)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  fmt.Sprintf("user with ID:%v deleted successfuly", input.UserID),
			Error: "",
		})
	}
}
