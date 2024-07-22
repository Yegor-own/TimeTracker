package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time-tracker/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetTask godoc
//
// @Description	gets task
// @Param 		task_id 			query int 	 false 	"used for finding task by id"
// @Param 		task_name 			query string false 	"used for finding task by name"
// @Param 		task_description 	query string false 	"used for finding task by desc"
// @Accept		json
// @Produce		json
// @Success		200	{object}	Output	"task data"
// @Failure		400	{object}	Output	"error"
// @Router		/api/task/getUser 	[get]
func GetTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := 0
		if c.Query("task_id") != "" {
			id, _ = strconv.Atoi(c.Query("task_id"))
		}

		name := c.Query("task_name")
		description := c.Query("task_description")

		task := model.Task{
			ID:          uint(id),
			Name:        name,
			Description: description,
		}

		db.Find(&task)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  task,
			Error: "",
		})
	}
}

// CreateTask godoc
//
// @Description	create task
// @Param 		task body model.TaskCreate true "used to set task params"
// @Accept		json
// @Produce		json
// @Success		200	{object} 	Output "success response"
// @Failure		400	{object}	Output	"error"
// @Router		/api/task/createTask 	[post]
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

// UpdateTask godoc
//
// @Description	update task by id
// @Param 		task body model.TaskUpdate true "used to update task"
// @Accept		json
// @Produce		json
// @Success		200	{object} 	Output "success response"
// @Failure		400	{object}	Output "error"
// @Router		/api/task/updateTask 	[patch]
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

// DeliteTask godoc
//
// @Description	delete task by id
// @Param 		task body model.TaskDelete true "used to delete task"
// @Accept		json
// @Produce		json
// @Success		200	{object} Output
// @Failure		400	{object} Output	"error"
// @Router		/api/task/deliteTask 	[delete]
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
