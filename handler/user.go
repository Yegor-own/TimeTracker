package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"time-tracker/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// LabourCosts godoc
//
//	@Summary		labour costs
//	@Description	gives tracks
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Success		200	{string}	string	"pong"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Failure		500	{string}	string	"ok"
//	@Router			/examples/ping [get]
func LabourCosts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			return handleErrors(c, err)
		}

		from, err := time.Parse("2006-01-02 15:04:05", c.Query("from"))
		if err != nil {
			return handleErrors(c, err)
		}

		to, err := time.Parse("2006-01-02 15:04:05", c.Query("to"))
		if err != nil {
			to = time.Now()
		}

		var tracks []model.Track

		db.Find(&tracks, "stop > ? and start < ? and user_id = ?", from, to, userId)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  tracks,
			Error: "",
		})
	}
}

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
			Passport:   input.Passport,
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
			ID: input.ID,
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
		if input.Passport != nil && *input.Passport != 0 {
			user.Passport = *input.Passport
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
			ID: input.ID,
		}

		res := db.Delete(&user)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  fmt.Sprintf("user with ID:%v deleted successfuly", input.ID),
			Error: "",
		})
	}
}
