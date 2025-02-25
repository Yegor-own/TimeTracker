package handler

import (
	"errors"
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
// @Description	gets user by id/name/surname/...
// @Param 		user_id 		query int 	 false 	"used for finding user by id"
// @Param 		user_name 		query string false 	"used for finding user by name"
// @Param 		user_surname 	query string false 	"used for finding user by surname"
// @Param 		user_patronymic query string false 	"used for finding user by patronymic"
// @Param 		user_address 	query string false 	"used for finding user by address"
// @Param 		user_passport 	query int 	 false 	"used for finding user by passport"
// @Accept		json
// @Produce		json
// @Success		200	{object}	Output	"user data"
// @Failure		400	{object}	Output	"error"
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
// @Description	getting users
// @Param 		limit 	query int true 	"max items on page"
// @Param 		page 	query int true 	"starts from 1"
// @Accept		json
// @Produce		json
// @Success		200	{object}	Output	"list of users"
// @Failure		400	{object}	Output	"error"
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

		result := map[string]interface{}{}
		db.Model(&model.User{}).First(&result)
		fId, ok := result["id"].(uint)
		if !ok {
			return handleErrors(c, errors.New("first ID inst int"))
		}
		x := ((page - 1) * limit) + int(fId) - 1

		var users []model.User

		db.Where("id > ?", x).Limit(limit).Find(&users)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  users,
			Error: "",
		})
	}
}

// LabourCosts godoc
//
//	@Description	gives tracks
//	@Param 			user_id query int 		true 	"used for finding tracks"
//	@Param 			from 	query string 	true 	"used for left edge of dates"
//	@Param 			to 		query string 	false 	"used for right edge of dates"
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	Output	"list of tracks"
//	@Failure		400		{object}	Output	"error"
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

		db.Select("(stop - start) as period, tracks .*").Where("stop > ? AND start < ? AND user_id = ?", from, to, userId).Order("period DESC").Find(&tracks)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  tracks,
			Error: "",
		})
	}
}

// CreateUser godoc
//
// @Description	create user
// @Param 		user body model.UserCreate true "used to set user"
// @Accept		json
// @Produce		json
// @Success		200	{object} 	Output "success response"
// @Failure		400	{object}	Output	"error"
// @Router		/api/user/createUser 	[post]
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

// UpdateUser godoc
//
// @Description	update user by id
// @Param 		user body model.UserUpdate true "used to update user"
// @Accept		json
// @Produce		json
// @Success		200	{object} 	Output "success response"
// @Failure		400	{object}	Output "error"
// @Router		/api/user/updateUser 	[patch]
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

// DeleteUser godoc
//
// @Description	delete user by id
// @Param 		user body model.UserDelete true "used to delete user"
// @Accept		json
// @Produce		json
// @Success		200	{object} Output
// @Failure		400	{object} Output	"error"
// @Router		/api/user/deleteUser 	[delete]
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
