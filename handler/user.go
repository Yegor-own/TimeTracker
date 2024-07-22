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

// GetUser godoc
//
// @Summary		labour costs
// @Description	gives track
// @Param 		user_id 		query int 	 false 	"used for finding user by id"
// @Param 		user_name 		query string false 	"used for finding user by name"
// @Param 		user_surname 	query string false 	"used for finding user by surname"
// @Param 		user_patronymic query string false 	"used for finding user by patronymic"
// @Param 		user_address 	query string false 	"used for finding user by address"
// @Param 		user_passport 	query int 	 false 	"used for finding user by passport"
// @Accept		json
// @Produce		json
// @Success		200	{string}	json	"list of tracks"
// @Failure		400	{string}	json	"error"
// @Router		/api/user/getUser 	[get]
func GetUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := 0
		if c.Query("user_id") != "" {
			id, _ = strconv.Atoi(c.Query("user_id"))
		}

		name := c.Query("user_name")
		surname := c.Query("user_surname")
		patronymic := c.Query("user_patronymic")
		address := c.Query("user_address")

		passport := 0
		if c.Query("user_passport") != "" {
			passport, _ = strconv.Atoi(c.Query("user_passport"))
		}

		user := model.User{
			ID:         uint(id),
			Name:       name,
			Surname:    surname,
			Patronymic: &patronymic,
			Address:    address,
			Passport:   passport,
		}

		db.Find(&user)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  user,
			Error: "",
		})
	}
}

// GetUsers godoc
//
// @Summary		"labour costs"
// @Description	"gives track"
// @Param 		limit 	query int true 	"max items on page"
// @Param 		page 	query int true 	"starts from 1"
// @Accept		json
// @Produce		json
// @Success		200	{string}	json	"list of tracks"
// @Failure		400	{string}	json	"error"
// @Router		/api/user/getUsers 	[get]
func GetUsers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			return handleErrors(c, err)
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			return handleErrors(c, err)
		}
		x := (page - 1) * limit

		var users []model.User

		db.Where("id > ? limit ?", x, limit).Find(&users)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  users,
			Error: "",
		})
	}
}

// LabourCosts godoc
//
//	@Summary		labour costs
//	@Description	gives tracks
//	@Param user_id query int true "used for finding tracks"
//	@Param from query string true "used for left edge of dates"
//	@Param to query string false "used for right edge of dates"
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json	"list of tracks"
//	@Failure		400	{string}	json	"error"
//	@Router			/api/user/labourCosts [get]
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

		db.Where("stop > ? AND start < ? AND user_id = ?", from, to, userId).Find(&tracks)

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
