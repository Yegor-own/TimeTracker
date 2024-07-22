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

// GetTrack godoc
//
// @Description	gets track
// @Param 		track_id 		query int 	false 	"used for finding track by id"
// @Param 		user_id 		query int 	false 	"used for finding track by user id"
// @Param 		task_id 		query int 	false 	"used for finding track by task id"
// @Accept		json
// @Produce		json
// @Success		200	{object}	Output	"track data"
// @Failure		400	{object}	Output	"error"
// @Router		/api/track/getTrack 	[get]
func GetTrack(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := 0
		if c.Query("track_id") != "" {
			id, _ = strconv.Atoi(c.Query("track_id"))
		}

		userId := 0
		if c.Query("user_id") != "" {
			id, _ = strconv.Atoi(c.Query("user_id"))
		}

		taskId := 0
		if c.Query("task_id") != "" {
			id, _ = strconv.Atoi(c.Query("task_id"))
		}

		track := model.Track{
			ID:     uint(id),
			UserID: uint(userId),
			TaskID: uint(taskId),
		}

		db.Find(&track)

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  track,
			Error: "",
		})
	}
}

// CreateTrack godoc
//
// @Description	create track
// @Param 		track body model.TrackCreate true "used to set track"
// @Accept		json
// @Produce		json
// @Success		200	{object} 	Output "success response"
// @Failure		400	{object}	Output "error"
// @Router		/api/track/createTrack 	[post]
func CreateTrack(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TrackCreate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		start, err := time.Parse("2006-01-02 15:04:05", input.Start)
		if err != nil {
			return handleErrors(c, err)
		}

		stop, err := time.Parse("2006-01-02 15:04:05", input.Stop)
		if err != nil {
			return handleErrors(c, err)
		}

		track := model.Track{
			UserID: input.UserID,
			TaskID: input.TaskID,
			Start:  start,
			Stop:   stop,
		}

		res := db.Create(&track)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  track,
			Error: "",
		})
	}
}

// UpdateTrack godoc
//
// @Description	update track by id
// @Param 		track 	body model.TrackUpdate true "used to update track"
// @Accept		json
// @Produce		json
// @Success		200	{object} 	Output "success response"
// @Failure		400	{object}	Output "error"
// @Router		/api/track/updateTrack 	[patch]
func UpdateTrack(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TrackUpdate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		track := model.Track{
			ID: input.ID,
		}
		res := db.First(&track)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		if input.UserID != nil && *input.UserID != 0 {
			track.UserID = *input.UserID
		}
		if input.TaskID != nil && *input.TaskID != 0 {
			track.TaskID = *input.TaskID
		}
		if input.Start != nil && *input.Start != "" {
			track.Start, _ = time.Parse("2006-01-02 15:04:05", *input.Start)
		}
		if input.Stop != nil && *input.Stop != "" {
			track.Stop, _ = time.Parse("2006-01-02 15:04:05", *input.Stop)
		}

		res = db.Save(&track)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  track,
			Error: "",
		})
	}
}

// DeliteTrack godoc
//
// @Description	delete track by id
// @Param 		track body model.TrackDelete true "used to delete track"
// @Accept		json
// @Produce		json
// @Success		200	{object} Output
// @Failure		400	{object} Output	"error"
// @Router		/api/track/deliteTrack 	[delete]
func DeliteTrack(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TrackDelete
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}

		track := model.Track{
			ID: input.ID,
		}

		res := db.Delete(&track)
		if res.Error != nil {
			return handleErrors(c, err)
		}

		c.Status(http.StatusOK)
		return c.JSON(Output{
			Data:  fmt.Sprintf("track with ID:%v deleted successfuly", input.ID),
			Error: "",
		})
	}
}
