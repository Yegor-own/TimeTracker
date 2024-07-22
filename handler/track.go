package handler

import (
	"fmt"
	"net/http"
	"time"
	"time-tracker/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTrack(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.TrackCreate
		err := c.BodyParser(&input)
		if err != nil {
			return handleErrors(c, err)
		}
		start, err := time.Parse("2006-01-02 15:04:05", input.Start)
		stop, err := time.Parse("2006-01-02 15:04:05", input.Stop)

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
			Data:  fmt.Sprintf("user with ID:%v deleted successfuly", input.ID),
			Error: "",
		})
	}
}
