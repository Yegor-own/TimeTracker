package handler

import (
	"fmt"
	"net/http"
	"time-tracker/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TaskCreate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		task := model.Task{
			Name:        input.Name,
			Description: input.Description,
		}

		res := db.Create(&task)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  task,
			Error: "",
		})
	}
}

func UpdateTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TaskUpdate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		task := model.Task{
			ID: input.ID,
		}
		res := db.First(&task)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		if input.Name != nil && *input.Name != "" {
			task.Name = *input.Name
		}
		if input.Description != nil && *input.Description != "" {
			task.Description = *input.Description
		}

		res = db.Save(&task)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  task,
			Error: "",
		})
	}
}

func DeliteTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TaskDelete
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		task := model.Task{
			ID: input.ID,
		}

		res := db.Delete(&task)
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
