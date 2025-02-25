package main

import (
	"io"
	"time-tracker/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func NewRouter(log io.Writer) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(), logger.New(logger.Config{
		Output: log,
	}))

	return app
}

func UserRouter(app fiber.Router, db *gorm.DB) {
	app.Get("/getUser", handler.GetUser(db))
	app.Get("/getUsers", handler.GetUsers(db))
	app.Get("/labourCosts", handler.LabourCosts(db))
	app.Post("/createUser", handler.CreateUser(db))
	app.Patch("/updateUser", handler.UpdateUser(db))
	app.Delete("/deleteUser", handler.DeliteUser(db))
}

func TaskRouter(app fiber.Router, db *gorm.DB) {
	app.Get("/getTask", handler.GetTask(db))
	app.Post("/createTask", handler.CreateTask(db))
	app.Patch("/updateTask", handler.UpdateTask(db))
	app.Delete("/deleteTask", handler.DeliteTask(db))
}

func TrackRouter(app fiber.Router, db *gorm.DB) {
	app.Get("/getTrack", handler.GetTrack(db))
	app.Post("/createTrack", handler.CreateTrack(db))
	app.Patch("/updateTrack", handler.UpdateTrack(db))
	app.Delete("/deleteTrack", handler.DeliteTrack(db))
}
